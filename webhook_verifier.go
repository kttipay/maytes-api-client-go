// Bundled with the generated SDK via bin/postprocess-go.sh.
//
// Lives at build/go-client/webhook_verifier.go after a `make generate-go`. The
// generator emits everything in package `maytesapiclient`, so this file lands
// in the same package and is importable as:
//
//   import maytes "github.com/maytes/api-client-go"
//   event, err := maytes.VerifyWebhookSignature(rawBody, header, secret)
//
// Verifies the `X-Maytes-Signature` header on an inbound webhook delivery and
// returns the parsed event. Merchants must never hand-roll this — see the four
// sharp edges below, each of which silently produces either a security hole or
// an outage:
//
//   1. RAW BODY. The signature covers the exact bytes we sent. Verifying
//      against re-marshalled JSON fails for any body whose key order or number
//      formatting differs after an unmarshal/marshal round-trip. This is the
//      single most common webhook integration bug.
//   2. MULTIPLE `v1=` ENTRIES. During a signing-secret rotation grace window
//      the header carries TWO signatures (current secret first, previous
//      second) and the delivery is valid if EITHER matches. A verifier that
//      reads only the first `v1=` works perfectly until the first rotation,
//      then rejects live traffic. See §7 of the webhook spec.
//   3. CONSTANT-TIME COMPARISON. A plain == on digests leaks how much of a
//      forged signature was correct, which is enough to recover a valid
//      signature one byte at a time.
//   4. REPLAY WINDOW. The `t=` timestamp is the anti-replay anchor; a captured
//      delivery stays valid forever without a freshness check.
//
// Verification and JSON unmarshalling are deliberately fused into ONE call: the
// function takes raw bytes and hands back the parsed event, so there is no
// intermediate state in which a caller can unmarshal first and verify second.
//
// Algorithm (webhook spec §6.1):
//   signed_payload = "<t>.<raw body>"
//   signature      = lowercase_hex( HMAC_SHA256(signing_secret, signed_payload) )

package maytesapiclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// DefaultWebhookToleranceSeconds is the default anti-replay window
// (webhook spec §6.1 advises ~5 min).
const DefaultWebhookToleranceSeconds = 300

// Reasons a delivery can fail verification. Log these; never echo one to the
// caller.
const (
	// WebhookReasonMalformedHeader — header absent or missing `t=` / `v1=`.
	WebhookReasonMalformedHeader = "malformed_header"
	// WebhookReasonTimestampOutOfTolerance — `t=` too far from now; possible replay.
	WebhookReasonTimestampOutOfTolerance = "timestamp_out_of_tolerance"
	// WebhookReasonNoMatchingSignature — wrong secret, or the body was modified.
	WebhookReasonNoMatchingSignature = "no_matching_signature"
	// WebhookReasonInvalidJSON — signature valid but body wasn't JSON.
	WebhookReasonInvalidJSON = "invalid_json"
)

// WebhookSignatureError reports that an inbound delivery is not authentic.
// Respond 400 and do NOT act on the body.
type WebhookSignatureError struct {
	// Reason is one of the WebhookReason* constants.
	Reason  string
	Message string
}

func (e *WebhookSignatureError) Error() string {
	return fmt.Sprintf("maytes webhook signature %s: %s", e.Reason, e.Message)
}

// MaytesWebhookEvent is a verified webhook event (webhook spec §6).
//
// Data is left as an open map because its shape varies per Type — switch on
// Type, then read the fields that event documents. Ignore unrecognised
// fields: additive changes do NOT bump APIVersion.
type MaytesWebhookEvent struct {
	// ID is a stable event id — identical across every retry. Deduplicate on it.
	ID string `json:"id"`
	// Type is e.g. "checkout.authorized", "checkout.voided", "webhook.test".
	Type string `json:"type"`
	// APIVersion is the envelope schema version, e.g. "2026-06". Bumped only
	// on breaking changes.
	APIVersion string `json:"api_version"`
	// CreatedAt is when the event occurred.
	CreatedAt string `json:"created_at"`
	// Data is the per-event payload.
	Data map[string]interface{} `json:"data"`
}

// VerifyWebhookSignature verifies an inbound Maytes webhook delivery and
// returns the parsed event.
//
// Pass the raw request body — the exact bytes, before any JSON unmarshal.
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    rawBody, err := io.ReadAll(r.Body)
//	    if err != nil {
//	        w.WriteHeader(http.StatusBadRequest)
//	        return
//	    }
//	    event, err := maytes.VerifyWebhookSignature(
//	        rawBody,
//	        r.Header.Get("X-Maytes-Signature"),
//	        os.Getenv("MAYTES_WEBHOOK_SECRET"),
//	    )
//	    if err != nil {
//	        w.WriteHeader(http.StatusBadRequest)
//	        return
//	    }
//	    // ... handle event.Type
//	    w.WriteHeader(http.StatusOK)
//	}
//
// On failure the error is always a *WebhookSignatureError.
func VerifyWebhookSignature(rawBody []byte, signatureHeader, signingSecret string) (*MaytesWebhookEvent, error) {
	return VerifyWebhookSignatureWithTolerance(rawBody, signatureHeader, signingSecret, DefaultWebhookToleranceSeconds)
}

// VerifyWebhookSignatureWithTolerance is VerifyWebhookSignature with an
// explicit anti-replay window.
//
// toleranceSeconds = 0 skips the freshness check entirely — TEST ONLY, for
// replaying a stored delivery in a fixture. Never 0 in production: it makes
// any captured delivery replayable forever.
func VerifyWebhookSignatureWithTolerance(
	rawBody []byte,
	signatureHeader, signingSecret string,
	toleranceSeconds int,
) (*MaytesWebhookEvent, error) {
	timestamp, signatures, err := parseWebhookSignatureHeader(signatureHeader)
	if err != nil {
		return nil, err
	}

	if toleranceSeconds > 0 {
		parsed, convErr := strconv.ParseInt(timestamp, 10, 64)
		if convErr != nil {
			return nil, &WebhookSignatureError{
				Reason:  WebhookReasonMalformedHeader,
				Message: "X-Maytes-Signature carries a non-numeric t= value.",
			}
		}
		ageSeconds := math.Abs(time.Since(time.Unix(parsed, 0)).Seconds())
		if ageSeconds > float64(toleranceSeconds) {
			return nil, &WebhookSignatureError{
				Reason: WebhookReasonTimestampOutOfTolerance,
				Message: fmt.Sprintf(
					"webhook timestamp is outside the %ds tolerance — possible replay",
					toleranceSeconds,
				),
			}
		}
	}

	mac := hmac.New(sha256.New, []byte(signingSecret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(rawBody)
	expected := hex.EncodeToString(mac.Sum(nil))

	// ANY match wins — during a rotation grace window the current secret's
	// signature and the previous secret's signature are both present.
	//
	// Compared as ASCII hex text rather than decoded bytes on purpose: a
	// malformed v1= would fail to decode, and hmac.Equal over the text is
	// equally constant-time.
	matched := false
	for _, candidate := range signatures {
		if hmac.Equal([]byte(candidate), []byte(expected)) {
			matched = true
			break
		}
	}
	if !matched {
		return nil, &WebhookSignatureError{
			Reason: WebhookReasonNoMatchingSignature,
			Message: "no v1 signature matched the computed HMAC — wrong signing secret, " +
				"or the body was modified in transit",
		}
	}

	var event MaytesWebhookEvent
	if err := json.Unmarshal(rawBody, &event); err != nil {
		return nil, &WebhookSignatureError{
			Reason:  WebhookReasonInvalidJSON,
			Message: "signature verified but the body is not valid JSON",
		}
	}
	if event.Data == nil {
		event.Data = map[string]interface{}{}
	}
	return &event, nil
}

func parseWebhookSignatureHeader(header string) (string, []string, error) {
	if strings.TrimSpace(header) == "" {
		return "", nil, &WebhookSignatureError{
			Reason:  WebhookReasonMalformedHeader,
			Message: "missing X-Maytes-Signature header",
		}
	}

	var timestamp string
	var signatures []string

	for _, part := range strings.Split(header, ",") {
		trimmed := strings.TrimSpace(part)
		switch {
		case strings.HasPrefix(trimmed, "t="):
			timestamp = strings.TrimPrefix(trimmed, "t=")
		case strings.HasPrefix(trimmed, "v1="):
			signatures = append(signatures, strings.TrimPrefix(trimmed, "v1="))
		}
		// Unknown schemes (a future v2=) are ignored, not fatal — forward
		// compatibility: we may add a scheme alongside v1 before removing v1.
	}

	if timestamp == "" || len(signatures) == 0 {
		return "", nil, &WebhookSignatureError{
			Reason:  WebhookReasonMalformedHeader,
			Message: "X-Maytes-Signature is malformed — expected `t=<unix>,v1=<hex>`",
		}
	}

	return timestamp, signatures, nil
}

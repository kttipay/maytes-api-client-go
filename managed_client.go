// Bundled with the generated SDK via bin/postprocess-go.sh.
//
// Lives at build/go-client/managed_client.go after a `make generate-go`. The
// generator emits everything in package `maytesapiclient`, so this file lands
// in the same package and is importable as:
//
//   import maytes "github.com/maytes/api-client-go"
//   api := maytes.CreateMaytesAPIClient(endpoint, clientID, clientSecret)
//
// Adds OAuth2 client_credentials credential management on top of the raw
// APIClient using a **reactive 401-retry** pattern:
//
//   1. First authed call mints a token via GetOAuthToken (a @auth([]) op) and
//      caches it.
//   2. Subsequent calls send the cached token as-is — no client-side expiry
//      math, no proactive refresh.
//   3. If the server returns 401, the bearerTransport refreshes the token and
//      retries the request once with the new Authorization header. The retry
//      uses a buffered copy of the request body so it can be re-sent verbatim.
//
// Concurrent 401s deduplicate via `RefreshIfStale(staleToken)`: only the first
// goroutine to take the lock mints; others see the cache already updated and
// use that token.

package maytesapiclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// OAuthTokenProvider caches an OAuth2 client_credentials token; refreshes only
// on demand. Goroutine-safe via a Mutex. An optional `initialToken` argument
// seeds the cache with a pre-minted JWT — used by the E2E test path to verify
// the wrapper's reactive 401-retry refreshes when the seeded token expires.
type OAuthTokenProvider struct {
	endpoint     string
	clientID     string
	clientSecret string

	mu       sync.Mutex
	token    string
	unauthed *APIClient
}

func NewOAuthTokenProvider(endpoint, clientID, clientSecret, initialToken string) *OAuthTokenProvider {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: endpoint}}
	return &OAuthTokenProvider{
		endpoint:     endpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
		token:        initialToken,
		unauthed:     NewAPIClient(cfg),
	}
}

func (p *OAuthTokenProvider) Get(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.token != "" {
		return p.token, nil
	}
	return p.mintLocked(ctx)
}

// RefreshIfStale mints a new token only if the cache still holds the stale
// token observed by the caller. If another goroutine refreshed first, returns
// that newer token instead — no second mint.
func (p *OAuthTokenProvider) RefreshIfStale(ctx context.Context, staleToken string) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.token != "" && p.token != staleToken {
		return p.token, nil
	}
	return p.mintLocked(ctx)
}

func (p *OAuthTokenProvider) mintLocked(ctx context.Context) (string, error) {
	body := *NewOAuthTokenRequest("client_credentials", p.clientID)
	body.SetClientSecret(p.clientSecret)
	resp, _, err := p.unauthed.DefaultAPI.GetOAuthToken(ctx).OAuthTokenRequest(body).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to mint OAuth token: %w", err)
	}
	p.token = resp.GetAccessToken()
	return p.token, nil
}

// bearerTransport injects a fresh Bearer token on each outbound request and,
// on a 401 response, refreshes the token and retries once. The two public
// endpoints (GetHealth, GetOAuthToken) are bypassed by URL path to avoid
// recursing into the token mint.
type bearerTransport struct {
	base     http.RoundTripper
	provider *OAuthTokenProvider
}

func (t *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path == "/api/health" || req.URL.Path == "/oauth/token" {
		return t.base.RoundTrip(req)
	}

	// Buffer the request body so we can replay it on retry. *bytes.Reader is
	// not naturally replayable; reading once and re-wrapping is the standard
	// pattern for retry middleware in Go.
	var bodyBytes []byte
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		_ = req.Body.Close()
		bodyBytes = b
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	token, err := t.provider.Get(req.Context())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := t.base.RoundTrip(req)
	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		return resp, err
	}

	// 401 — drain + close, refresh, retry once.
	_ = resp.Body.Close()
	newToken, err := t.provider.RefreshIfStale(req.Context(), token)
	if err != nil {
		return nil, err
	}

	retryReq := req.Clone(req.Context())
	if bodyBytes != nil {
		retryReq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}
	retryReq.Header.Set("Authorization", "Bearer "+newToken)
	return t.base.RoundTrip(retryReq)
}

// CreateMaytesAPIClient returns a fully managed APIClient. The OAuth token is
// minted on the first authed call, cached, and refreshed reactively when the
// server returns 401 (one retry per request).
//
//	api := maytesapiclient.CreateMaytesAPIClient(endpoint, clientID, clientSecret)
//	out, _, err := api.DefaultAPI.CreateCheckout(ctx).
//	    CreateCheckoutRequestContent(body).Execute()
func CreateMaytesAPIClient(endpoint, clientID, clientSecret string) *APIClient {
	return createWithInitialToken(endpoint, clientID, clientSecret, "")
}

// CreateMaytesAPIClientWithInitialToken is the SDK-test-only variant that
// seeds the cache with a pre-minted JWT. When that token expires server-side,
// the wrapper's reactive 401-retry kicks in and mints a fresh full-TTL token
// via /oauth/token. Production callers should use CreateMaytesAPIClient.
func CreateMaytesAPIClientWithInitialToken(endpoint, clientID, clientSecret, initialToken string) *APIClient {
	return createWithInitialToken(endpoint, clientID, clientSecret, initialToken)
}

func createWithInitialToken(endpoint, clientID, clientSecret, initialToken string) *APIClient {
	provider := NewOAuthTokenProvider(endpoint, clientID, clientSecret, initialToken)
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: endpoint}}
	cfg.HTTPClient = &http.Client{
		Transport: &bearerTransport{
			base:     http.DefaultTransport,
			provider: provider,
		},
	}
	return NewAPIClient(cfg)
}

# Maytes Go SDK

Official Go client for the [Maytes API](https://developers.maytes.co) — create,
capture and cancel checkouts, and verify inbound webhook signatures.

> **Generated package.** This repository is a build artefact, published from the
> Maytes API service definition. Pull requests here are not accepted — file
> issues instead, and fixes land upstream in the source model.

## Install

```bash
go get github.com/kttipay/maytes-api-client-go
```

## Usage

The client mints an OAuth2 `client_credentials` token on first use, caches it,
and — on a `401` — refreshes and retries the request once. There is no
client-side expiry maths to get wrong, and no proactive refresh to schedule.

```go
import maytes "github.com/kttipay/maytes-api-client-go"

client := maytes.CreateMaytesAPIClient(endpoint, clientID, clientSecret)
checkout, _, err := client.DefaultAPI.CreateCheckout(ctx).Execute()

event, err := maytes.VerifyWebhookSignature(rawBody, signatureHeader, signingSecret)
```

There is no Go package registry — `go get` resolves this module straight from
its repository, so the module path and the repo URL are the same string.

### Verifying webhooks

Verify **before** parsing the body, and pass the raw request bytes — not a
re-encoded object. Re-serialising changes the bytes and the signature will not
match. The event id is stable across retries, so deduplicate on it.

## Documentation

Full integration docs: <https://developers.maytes.co>

## License

MIT — see [LICENSE](LICENSE).

# foundry

Foundry is a small shared Go utility module for Penny Labs services.

## Current packages

- `pkg/api` — HTTP JSON rendering helpers.
- `pkg/apperr` — status-aware application errors with safe client-facing messages.
- `pkg/hash` — SHA and bcrypt helpers for token/key hashing.
- `pkg/logger` — shared logger setup.

## V1 roadmap role

Foundry is supporting infrastructure for the PennyOS v1 roadmap, not a standalone product surface.

For v1, it should remain boring and stable:

- do not add PennyOS product behavior here,
- do not add managed-sync/runtime/billing state here,
- keep service-specific contracts in `penny`, `management-api`, and generated OpenAPI clients,
- only change Foundry when a shared helper is already used by multiple services or a security/maintenance fix is needed.

## Validation

```bash
go test ./...
```

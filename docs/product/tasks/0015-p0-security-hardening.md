# 2026-02-13 - P0 Security Hardening

## Goal
Close critical P0 security gaps in the API before moving to extension/plugin work.

## What was done
1. Hardened `POST /voice/token` to ignore client-provided identity (`userId`).
2. Hardened `POST /voice/token` to validate channel existence and `voice` channel type.
3. Hardened `POST /voice/token` to enforce membership authorization before issuing LiveKit token.
4. Fixed invite usage accounting in `POST /invites/:code/join` to increment `uses_count` only when membership insert really happens.
5. Fixed invite usage accounting in `POST /invites/:code/join` to avoid consuming invite use for duplicate joins.
6. Added real CORS middleware with configurable allowed origins (`APP_ALLOWED_ORIGINS`).
7. Replaced permissive WebSocket `CheckOrigin=true` with allowlist enforcement.
8. Added admin RBAC checks for telemetry admin endpoints (`owner/admin` required).
9. Added unit tests for origin normalization/allowlist and CORS preflight/deny behavior.

## What is pending
1. Implement rate limiting baseline.
2. Convert Protected Community Mode from flag to enforceable policy set.
3. Add integration tests for end-to-end guardrails.

## Risks
1. Telemetry RBAC currently treats any community `owner/admin` as global admin access.
2. Remaining monolith handlers in `backend/cmd/api/main.go` still concentrate risk.

## Validation
1. `cd backend && go test ./...` -> passed.
2. `cd backend && go build ./cmd/api` -> passed.

## Notes
1. CORS and WS origin checks share the same origin normalization and allowlist helper.
2. Disallowed request origins now return `403 origin not allowed` instead of silently continuing.

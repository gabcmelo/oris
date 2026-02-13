# API Standards

## Contract stability
1. Preserve existing endpoint contracts unless explicitly approved.
2. If breaking changes are required, document migration path in ADR + README.

## REST conventions
1. JSON request/response.
2. Error payload format: `{ "error": "<message>" }`.
3. Auth with `Authorization: Bearer <token>`.

## Security baseline
1. Validate membership/role before protected actions.
2. Restrict sensitive endpoints by RBAC.
3. Validate WebSocket origin and token.

## Pagination
Current MVP endpoints are mostly small-list. Introduce consistent pagination when list size grows.

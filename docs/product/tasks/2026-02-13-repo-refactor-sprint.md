# 2026-02-13 - Repo Refactor Sprint Note

## What was done
1. Audited repository and created migration map.
2. Added monorepo standards docs and ADRs.
3. Added GitHub issue/PR templates and CI workflow.
4. Created standardized folder layout for backend/frontend/infra/docs/desktop.
5. Refactored backend auth into module layers (`domain/usecase/infra/transport`).
6. Reorganized frontend entry/app structure with compatibility wrappers.
7. Added standardized compose path at `infra/docker/docker-compose.yml`.

## Pending
1. Continue extracting remaining backend domains from `cmd/api/main.go`.
2. Apply backend hardening P0 items (voice/ws/cors/rbac/invite accounting).
3. Remove legacy compatibility layers after stabilization window.

## Risks
1. Remaining monolithic handlers in `main.go` can still cause coupling regressions.
2. Duplicate compose paths can drift if not kept synchronized.

## Validation
1. `go test ./...` (backend) passed.
2. `go build ./cmd/api` (backend) passed.
3. `npm run build` (frontend) passed.
4. `docker compose -f infra/docker/docker-compose.yml up -d --build` passed.
5. API smoke flow (auth + community + message) passed.

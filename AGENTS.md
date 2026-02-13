# AGENTS

## Purpose
Operational guide for humans and AI agents working in this repository.

## Repository structure conventions
1. Backend code goes under `backend/`.
2. Frontend code goes under `frontend/`.
3. Infra assets go under `infra/`.
4. Product/architecture docs go under `docs/`.
5. Desktop (post-MVP) planning/code goes under `desktop/`.

## Backend modularity rule
Each backend domain must follow:
`backend/internal/modules/<domain>/{domain,usecase,infra,transport}`

Definitions:
1. `domain`: entities, domain interfaces, domain errors.
2. `usecase`: business rules/application services.
3. `infra`: persistence/providers (postgres, cache, external services).
4. `transport`: HTTP handlers/DTO/route bindings.

## Shared is not a junk drawer
`backend/internal/shared` can contain only:
1. Cross-module DTOs with clear ownership.
2. Truly generic helpers that are domain-agnostic.
3. Constants used by multiple modules that do not belong to core config.

Do not put feature-specific business logic in `shared`.

## Handler vs usecase pattern
1. Handler is thin:
   - parse/validate request
   - call usecase
   - map response/error to HTTP
2. Usecase owns business rules and orchestration.
3. Infra owns SQL/storage/provider details.

## Testing standards
1. Unit tests near package when practical (`*_test.go`).
2. Integration tests under `backend/test/integration`.
3. Fixtures under `backend/test/fixtures`.
4. Test naming:
   - `Test<Thing>_<Condition>_<ExpectedResult>`
5. Priority coverage:
   - auth, invite flows, moderation, voice token guardrails, websocket access.

## Branch/commit/PR standards
1. Branch naming:
   - `feat/<scope>`
   - `fix/<scope>`
   - `chore/<scope>`
2. Commit style:
   - imperative and scoped: `backend: extract auth usecase`.
3. PR must include:
   - summary
   - test evidence
   - docs impact
   - migration notes (if DB/infrastructure changed)

## Local run commands
1. Backend test: `cd backend && go test ./...`
2. Backend build: `cd backend && go build ./cmd/api`
3. Frontend install/build:
   - `cd frontend && npm install`
   - `cd frontend && npm run build`
4. Full stack:
   - `docker compose -f infra/docker-compose.yml up -d --build`
5. Planned standardized stack path:
   - `docker compose -f infra/docker/docker-compose.yml up -d --build`

## Migrations
Current source of truth:
1. `database/migrations/*.sql`

Planned infra path:
1. `infra/docker/postgres/init.sql`
2. keep migrations compatible while transition is in progress.

## CI validation baseline
Before opening a PR:
1. Backend build/tests pass.
2. Frontend build passes.
3. Docs updated for behavior/structure changes.

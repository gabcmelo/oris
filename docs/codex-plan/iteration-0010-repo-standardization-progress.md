# Iteration 0010 Progress

## Focus
Monorepo standardization + architecture governance + compatibility-safe restructuring.

## Implemented
1. Repository governance:
   - `AGENTS.md`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SECURITY.md`
   - `.gitignore`, `.editorconfig`, `.env.example`, `LICENSE`
2. GitHub standards:
   - Issue templates (`bug`, `feature`, `task`, `question`)
   - `PULL_REQUEST_TEMPLATE.md`
   - Workflows: `ci.yml`, `release.yml` (placeholder), `security.yml` (placeholder)
   - `dependabot.yml`
3. Docs architecture/product/runbooks/api:
   - `docs/index.md`
   - ADRs 0001 and 0002
   - standards, diagrams, runbooks, API starter files
4. Backend alignment:
   - `internal/app/http/*` wrappers
   - `internal/core/config` and `internal/core/auth/jwt` wrappers
   - `auth` module split with `domain/usecase/infra/transport`
5. Frontend structure:
   - introduced `src/app`, `src/pages`, `src/features`, `src/entities`, `src/shared`, `src/styles`
   - compatibility wrappers kept in `src/main.jsx`, `src/App.jsx`, `src/styles.css`
6. Infra:
   - standardized compose path at `infra/docker/docker-compose.yml`
   - postgres init placeholder at `infra/docker/postgres/init.sql`
7. Desktop:
   - post-MVP placeholder docs under `desktop/`

## Validation
1. Backend:
   - `go test ./...`
   - `go build ./cmd/api`
2. Frontend:
   - `npm run build`
3. Infra:
   - `docker compose -f infra/docker/docker-compose.yml config`
   - `docker compose -f infra/docker/docker-compose.yml up -d --build`
   - health check and API version endpoint responding.

## Notes
1. Legacy compatibility paths were preserved to avoid runtime breakage.
2. Remaining backend domains beyond auth still need extraction from `cmd/api/main.go` and are tracked in TASK-0011.

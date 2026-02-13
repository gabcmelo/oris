# Oris

Build your own voice community platform. Open-source, self-hosted real-time infrastructure powered by WebRTC.

## System overview
1. Backend: Go + Gin + PostgreSQL (+ Redis support)
2. Frontend: React + Vite
3. Voice: LiveKit
4. Infra: Docker Compose
5. Telemetry: optional, opt-in only

## Repository structure (summary)
1. `backend/` API and domain modules
2. `frontend/` web client
3. `infra/` compose and container infrastructure
4. `docs/` architecture/product/runbooks/api
5. `.github/` contribution templates and workflows
6. `desktop/` post-MVP Electron plan

## Quickstart (dev)
1. Start stack (compat path):
   - `docker compose -f infra/docker-compose.yml up -d --build`
2. Start stack (standardized path):
   - `docker compose -f infra/docker/docker-compose.yml up -d --build`
3. Open:
   - Web: `http://localhost:5173`
   - API health: `http://localhost:8080/healthz`

## Local validation commands
1. Backend tests:
   - `cd backend && go test ./...`
2. Backend build:
   - `cd backend && go build ./cmd/api`
3. Frontend build:
   - `cd frontend && npm install --include=dev && npm run build`

## Important docs
1. Agent and collaboration standards: `AGENTS.md`
2. Architecture index: `docs/index.md`
3. ADRs:
   - `docs/architecture/decisions/0001-structure-monorepo.md`
   - `docs/architecture/decisions/0002-modular-backend.md`
4. Local runbook: `docs/runbooks/local-dev.md`
5. Contribution guide: `CONTRIBUTING.md`
6. Security policy: `SECURITY.md`

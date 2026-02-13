# Oris - Copilot Instructions

## 1. Product mission
Build an open-source, self-hosted communication platform similar to Discord, with strong safety defaults for kids, teens, and adults.

Primary audience:
1. Geek/dev communities
2. Gamers and streamers
3. Remote teams and home-office groups

Core values:
1. Safety first
2. Self-host simplicity
3. Open-source transparency
4. Fast iteration without breaking MVP stability

## 2. Current architecture (authoritative)
1. Frontend: React + Vite (web-first, PWA direction)
2. Backend: Go + Gin + PostgreSQL + Redis
3. Voice: LiveKit
4. Infra: Docker Compose
5. Telemetry: optional, opt-in only, OTLP path

Backend status:
1. Refactor in progress from single-file API into modular `backend/internal/*`
2. Keep endpoint contracts stable while refactoring

## 3. Non-negotiable guardrails
1. Do not remove existing MVP endpoints unless explicitly requested.
2. Keep API behavior backward compatible during refactors.
3. Telemetry must remain opt-in and without PII.
4. Prefer incremental changes over full rewrites.
5. Always validate with build + smoke tests after backend changes.

## 4. Security and safety priorities (P0)
When touching backend, prioritize these fixes and never regress them:
1. Voice token must validate channel membership and channel type.
2. WebSocket origin must be validated (no open `CheckOrigin=true` policy in production paths).
3. CORS must support web client origin config.
4. Invite accounting must not consume uses on duplicate membership joins.
5. Admin telemetry endpoints must require proper RBAC.

## 5. Code organization target
Use this structure for new code:
1. `backend/cmd/api/main.go` -> bootstrap only
2. `backend/internal/config` -> env and config loading
3. `backend/internal/http` -> router, middleware, handlers
4. `backend/internal/modules/<domain>` -> domain handlers/services/repositories
5. `backend/internal/platform` -> shared technical utilities
6. `backend/internal/realtime` -> ws hub and presence

## 6. Collaboration protocol (Codex here + Copilot there)
We are running parallel development streams. Follow this protocol:
1. Before starting, pick or create one task in `docs/tasks/TASK-XXXX-*.md`.
2. Work on one isolated scope per task.
3. Avoid editing unrelated files to reduce merge conflicts.
4. At finish, update the task status/checklist and add short result notes.
5. If touching the same area as another agent, prefer additive changes and small commits.

Recommended split:
1. Agent A: backend modularization and security hardening
2. Agent B: frontend UX flows and Electron/post-MVP client packaging

## 7. Definition of done (per change)
1. `go build ./...` passes for backend changes
2. Docker build passes for changed service(s)
3. Smoke flow still works:
   - auth -> community -> channel -> message
4. If auth/voice/ws changed, include targeted smoke checks for those paths
5. Update docs/task file for traceability

## 8. Output style for generated code
1. Keep code direct and practical.
2. Prefer explicit names over abstractions that hide behavior.
3. Add short comments only when logic is non-obvious.
4. Keep files UTF-8 and avoid unnecessary formatting churn.

# Oris - Copilot Instructions

**Quick Links:**
- Task system: `docs/product/tasks/taskboard.csv` (single source of truth)
- Full protocol: `AGENTS.md`
- Product vision: `docs/product/vision.md`
- Current task: Find `status=DOING` in taskboard.csv

---

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

**✅ Completed in task 0015 (P0 Security hardening):**
1. Voice token validates channel membership and channel type
2. WebSocket origin validation (no open `CheckOrigin=true` policy)
3. CORS with configurable web client origins (`APP_ALLOWED_ORIGINS`)
4. Invite accounting fixed (no duplicate use consumption)
5. Admin telemetry endpoints require RBAC (owner/admin)

**⚠️ Still to do:**
- Rate limiting baseline for critical endpoints
- Protected Community Mode operational rules (beyond flag)
- Integration tests for end-to-end guardrails

**Never regress:** All completed P0 security fixes must remain in place.

## 5. Code organization target
Use this structure for new code:
1. `backend/cmd/api/main.go` -> bootstrap only
2. `backend/internal/config` -> env and config loading
3. `backend/internal/http` -> router, middleware, handlers
4. `backend/internal/modules/<domain>` -> domain handlers/services/repositories
5. `backend/internal/platform` -> shared technical utilities
6. `backend/internal/realtime` -> ws hub and presence

## 6. Task execution protocol (MANDATORY)

This project follows **incremental development** with strict WIP limits.

### Golden Rule: WIP Limit = 1
**Only ONE task can be DOING at any time across the entire repo.**

### Before coding anything:
1. Read `docs/product/tasks/taskboard.csv` (single source of truth for all tasks)
2. Identify the ONE task with `status=DOING`
3. Open the task file: `docs/product/tasks/NNNN-task-name.md`
4. Confirm the task has:
   - Clear scope (includes/excludes)
   - Technical plan
   - Acceptance criteria
   - Explicit status `DOING` (both in file and CSV)

**If there is NO DOING task or MORE THAN ONE:**
- DO NOT CODE
- Ask user to clarify which task to work on

### During implementation:
1. Follow the task's scope strictly
2. If you discover missing subtasks, create new tasks in BACKLOG (don't expand current scope)
3. Work on isolated changes that can be reviewed quickly
4. Avoid editing unrelated files to reduce merge conflicts
5. Document progress in the task file

### Before marking DONE:
A task moves to DONE only when:
- [ ] Code implemented for that scope only
- [ ] Tests added or explicitly justified
- [ ] Validation commands + results recorded in task file
- [ ] `taskboard.csv` updated (status=DONE, updated_at=today)
- [ ] Docs updated if behavior changed

### Task file location:
- All tasks: `docs/product/tasks/NNNN-name.md` (4-digit numeric IDs)
- Taskboard: `docs/product/tasks/taskboard.csv`
- Template: `docs/product/tasks/0000-template.md`
- Evidence: `docs/product/tasks/logs/`

### Task sizing:
A single task should deliver ONE of:
- One screen/component
- One endpoint
- One flow (e.g., login flow)
- One migration
- One doc improvement
- One refactor

If a task is too large, split it into smaller tasks.

## 7. Definition of done (per task)
1. `go build ./...` passes for backend changes
2. `npm run build` passes for frontend changes
3. Docker build passes for changed service(s)
4. Smoke flow still works:
   - auth -> community -> channel -> message
5. If auth/voice/ws changed, include targeted smoke checks for those paths
6. Evidence recorded in task file (commands + outputs)
7. Task file updated with results
8. `taskboard.csv` updated (status=DONE, updated_at=YYYY-MM-DD)
9. Docs updated if behavior changed

## 8. Creating new tasks (when needed)

If you need to create a new task:
1. Find next sequential ID in `taskboard.csv` (last ID + 1)
2. Copy template: `docs/product/tasks/0000-template.md`
3. Rename to: `docs/product/tasks/NNNN-task-slug.md`
4. Fill all required sections:
   - Description (why this task exists)
   - Objective (measurable outcome)
   - Scope (includes/excludes)
   - Technical Plan (step-by-step)
   - Acceptance Criteria (checkboxes)
   - Status: BACKLOG
   - Owner: TBD or your name
   - Created At: today's date
5. Add row to `taskboard.csv`:
   ```csv
   NNNN,Task title,area,BACKLOG,P1,TBD,YYYY-MM-DD,YYYY-MM-DD,,,Optional notes
   ```

**Task areas:** backend | frontend | infra | docs | product  
**Priorities:** P0 (critical) | P1 (important) | P2 (nice-to-have)

## 9. Output style for generated code
1. Keep code direct and practical.
2. Prefer explicit names over abstractions that hide behavior.
3. Add short comments only when logic is non-obvious.
4. Keep files UTF-8 and avoid unnecessary formatting churn.

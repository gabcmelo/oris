# AGENTS

## Purpose
Operational guide for humans and AI agents working in this repository.

This repo is developed **incrementally**. The goal is not to auto-generate a full SaaS in one pass.
We build, review, test, and approve each small step before moving forward.

---

## Golden Rule: WIP Limit = 1
At any time, there must be **only one ACTIVE task** (Doing) across the entire repo.

Agents MUST NOT:
- start multiple features in parallel
- implement “full MVP” bundles (auth + chat + voice + integrations) in a single PR
- add features not explicitly requested in the ACTIVE task

Agents MUST:
- complete the ACTIVE task with evidence
- wait for approval (or a new ACTIVE task) before proceeding

---

## Execution Protocol (Mandatory)

### Before coding anything
1. Read `docs/product/vision.md` and `docs/product/roadmap.md`.
2. Read `docs/product/tasks/taskboard.csv` and identify:
   - exactly one row with `status=DOING`
3. Open the matching task file under `docs/product/tasks/<id>-*.md`.
4. Confirm the task has:
   - clear scope (includes/excludes)
   - technical plan
   - acceptance criteria
   - explicit status `DOING`

If there is no DOING task, or more than one DOING task:
- DO NOT CODE.
- Create/adjust tasks so that exactly one task is DOING.

### During implementation
- Follow the task’s scope strictly.
- If you discover missing subtasks, create new tasks (Backlog) instead of expanding scope.
- Prefer “small PRs” that can be reviewed in minutes.

### Before marking done
A task can be moved to DONE only if:
1. Code is implemented for that scope only
2. Tests are added or explicitly justified
3. Validation commands + results are recorded in the task file
4. `taskboard.csv` is updated
5. Docs are updated if behavior changed

---

## Task System (v2)

### Folder structure
- `docs/product/tasks/`
  - `taskboard.csv` (single source of truth for status)
  - `0001-some-task.md` (task cards)
  - `logs/` (optional: evidence logs / screenshots / notes)

### task IDs
All tasks must have an incremental numeric ID:
- `0001`, `0002`, `0003`, ...

Task file naming:
- `docs/product/tasks/<id>-<slug>.md`
Example:
- `docs/product/tasks/0007-voice-join-flow.md`

### Task lifecycle states
Allowed statuses in taskboard:
- `BACKLOG`
- `DOING`
- `DONE`
- `BLOCKED`
- `CANCELLED`

Rules:
- Only one `DOING` at a time
- A task cannot move to `DONE` without evidence + validation

---

## Taskboard CSV (Mandatory)

`docs/product/tasks/taskboard.csv` is the execution orchestrator.
It is updated in every PR that changes task status.

### CSV schema
Columns (fixed):
- `id`
- `title`
- `area` (backend|frontend|infra|docs|product)
- `status` (BACKLOG|DOING|DONE|BLOCKED|CANCELLED)
- `priority` (P0|P1|P2)
- `owner`
- `created_at` (YYYY-MM-DD)
- `updated_at` (YYYY-MM-DD)
- `depends_on` (comma-separated task ids)
- `pr` (link or PR number)
- `notes`

Example row:
`0007,Voice join UX baseline,frontend,DOING,P0,agent,2026-02-12,2026-02-12,,#123,"Implement join/leave UI only"`

### Agent behavior with the taskboard
- Agent MUST read `taskboard.csv` first.
- Agent MUST work only on the DOING task.
- Agent MUST update `taskboard.csv` when moving task status.
- Agent MUST NOT reorder priorities without explicit instruction.

---

## Task Card Standard (Mandatory)

All new tasks must follow the template below.
Tasks must be small enough to be completed and reviewed quickly.

### Task sizing rule
A single task should usually deliver **one** of:
- one screen
- one endpoint
- one flow (join/leave voice)
- one migration
- one doc improvement
- one refactor

If a task contains multiple flows/screens/endpoints:
- split it into smaller tasks.

### Official Task Template
```markdown
# [TASK TITLE]

## Description
Clear contextual explanation of the problem or need.
Explain why this task exists.

## Objective
Concrete and measurable outcome.

## Scope
### Includes
- Item
- Item

### Excludes
- Out of scope items
- Future improvements

## Technical Plan
Step-by-step executable technical plan.
Must be implementation-oriented.

## Acceptance Criteria
- [ ] Implemented exactly as scoped
- [ ] Unit tests added (when applicable)
- [ ] Integration tests added (when applicable)
- [ ] Docs updated (when behavior changes)
- [ ] Evidence recorded (commands + outputs)

## Evidence
Commands run + results (paste output or summary).
Example:
- `go test ./...` ✅
- `npm test` ✅

## Risks
- Possible side effects
- Infra/database impact
- Security impact

## Dependencies
- Related tasks
- External services
- Required approvals

## Status
BACKLOG | DOING | DONE | BLOCKED | CANCELLED

## Created At
YYYY-MM-DD

## Updated At
YYYY-MM-DD

## Owner
TBD

# ⚠️ DEPRECATED - Features converted to tasks

This folder structure has been **deprecated** as of 2026-02-12.

## Migration
Features are now tracked as tasks in the unified taskboard:
**`docs/product/tasks/taskboard.csv`**

## Rationale
Following the AGENTS.md principle of "WIP Limit = 1", all work items (whether features, epics, bugs, or improvements) are now tracked as tasks in a single taskboard for better execution control and incremental delivery.

## Migration mapping
- **FEAT-010-community-and-voice-mvp.md** → `docs/product/tasks/0023-community-voice-mvp.md`

## For new features
Instead of creating FEAT-* files:
1. Add a row to `docs/product/tasks/taskboard.csv`
2. Create a task file: `docs/product/tasks/NNNN-feature-name.md`
3. Use template: `docs/product/tasks/0000-template.md`
4. Set status as BACKLOG
5. Mark as DOING only when ready to start (respecting WIP Limit = 1)

## Template location
Use the unified task template:
**`docs/product/tasks/0000-template.md`**

---

**This folder will be removed in a future cleanup.**

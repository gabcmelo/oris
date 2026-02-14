# Tasks - Execution System

This folder contains the **single source of truth** for all work items in the Oris project.

## Quick Start
Before working on anything:
1. **Read** `taskboard.csv` to see all tasks
2. **Find** the ONE task with `status=DOING`
3. **Open** the task file `NNNN-task-name.md`
4. **Work** on that task only
5. **Update** `taskboard.csv` when status changes

## Golden Rule: WIP Limit = 1
**Only ONE task can be DOING at any time.**

This ensures:
- Focused, incremental progress
- Proper review and testing
- No half-implemented features
- Clear visibility of current work

## File structure

### taskboard.csv
The execution orchestrator. Contains all tasks with their status.

**Columns:**
- `id` - Numeric ID (0001, 0002, ...)
- `title` - Short descriptive title
- `area` - backend | frontend | infra | docs | product
- `status` - BACKLOG | DOING | DONE | BLOCKED | CANCELLED
- `priority` - P0 | P1 | P2
- `owner` - Who's responsible
- `created_at` - YYYY-MM-DD
- `updated_at` - YYYY-MM-DD
- `depends_on` - Comma-separated task IDs
- `pr` - PR number or link
- `notes` - Additional context

### Task files (NNNN-name.md)
One file per task, following the template in `0000-template.md`.

**Naming:** 
- Use 4-digit numeric ID: `0001-task-name.md`
- Keep slug descriptive and kebab-case

**Required sections:**
- Description - Why this task exists
- Objective - Measurable outcome
- Scope - Includes/Excludes
- Technical Plan - Implementation steps
- Acceptance Criteria - Checkboxes for completion
- Evidence - Validation commands + outputs
- Risks - Potential issues
- Dependencies - Related tasks/services
- Status - Current state
- Owner - Assignee
- Dates - Created/Updated

### logs/
Evidence storage for task validation:
- Command outputs
- Screenshots
- Test results
- Integration/smoke test logs

## Workflow

### Starting a new task
1. Check that NO task is currently DOING
2. Pick a BACKLOG task (usually highest priority)
3. Update `taskboard.csv`: change status to DOING
4. Read the task file completely
5. Follow the Technical Plan
6. Check off Acceptance Criteria as you go

### While working
- Stay within the defined Scope
- Document evidence in the task file or logs/
- If you find new work needed, create a NEW task in BACKLOG
- Never expand scope of current task

### Finishing a task
A task moves to DONE only when:
- ✅ All Acceptance Criteria checked
- ✅ Evidence recorded (commands + outputs)
- ✅ Tests pass (when applicable)
- ✅ Docs updated (when behavior changed)
- ✅ `taskboard.csv` updated with status=DONE

### Task states
- **BACKLOG** - Not started, waiting in queue
- **DOING** - Currently active (limit = 1)
- **DONE** - Completed with evidence
- **BLOCKED** - Waiting on dependency or approval
- **CANCELLED** - Abandoned or deprioritized

## Creating tasks

### Use the template
Copy `0000-template.md` and rename with next sequential ID.

### Task sizing
A task should deliver **ONE** of:
- One screen/component
- One endpoint
- One flow (e.g., login flow)
- One migration
- One doc section
- One refactor

If too large, split into smaller tasks.

### Increment ID
Find the highest ID in `taskboard.csv`, add 1.
Example: if last is 0026, next is 0027.

### Add to taskboard
Create a new row in `taskboard.csv`:
```csv
0027,Short task title,backend,BACKLOG,P1,TBD,2026-02-12,2026-02-12,,,Notes here
```

## Priority guidelines
- **P0** - Critical security, MVP blockers, production bugs
- **P1** - Important features, UX improvements, tech debt
- **P2** - Nice-to-have, optimizations, future work

## See also
- `AGENTS.md` - Complete execution protocol
- `docs/product/vision.md` - Product direction
- `docs/product/roadmap.md` - Strategic plan
- `0000-template.md` - Official task template

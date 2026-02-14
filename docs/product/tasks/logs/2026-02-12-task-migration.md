# Task System Migration - 2026-02-12

## What was done
Centralized all task documentation from multiple locations into a single, organized structure under `docs/product/tasks/`.

## Changes
### 1. Created unified taskboard.csv
- New file: `docs/product/tasks/taskboard.csv`
- Schema: `id,title,area,status,priority,owner,created_at,updated_at,depends_on,pr,notes`
- Consolidated 26 tasks from various sources

### 2. Migrated all tasks to docs/product/tasks/
**From docs/tasks/**:
- TASK-0001 through TASK-0014 migrated and renamed to 0001-*.md through 0014-*.md

**From docs/product/tasks/**:
- tasks.csv (old format) → taskboard.csv (new format)
- 2026-02-13-p0-security-hardening.md → 0015-p0-security-hardening.md
- backlog.md → integrated into taskboard.csv
- 2026-02-13-repo-refactor-sprint.md → removed (content integrated)

**From docs/product/features/**:
- FEAT-010-community-and-voice-mvp.md → 0023-community-voice-mvp.md

### 3. Updated template
- TASK-000-template.md → 0000-template.md
- Updated to match AGENTS.md official template structure
- Includes all required sections: Description, Objective, Scope, Technical Plan, Acceptance Criteria, Evidence, Risks, Dependencies, Status, Owner, Dates

### 4. Created infrastructure
- `docs/product/tasks/logs/` folder for evidence storage
- `docs/product/tasks/logs/README.md` with usage guidelines

## Current structure
```
docs/product/tasks/
├── taskboard.csv (single source of truth)
├── 0000-template.md (official template)
├── 0001-plan-governance.md
├── 0002-research-analysis-mapeamento.md
├── ... (all tasks 0001-0026)
└── logs/
    └── README.md
```

## Task status summary
From taskboard.csv:
- DONE: 15 tasks (0001-0010, 0012-0015)
- DOING: 1 task (0011 - Backend restructure and modularization) ✅ WIP Limit = 1
- BACKLOG: 11 tasks (0016-0026)

## Next steps
1. Review taskboard.csv for any duplicate tasks (0018-0022 may overlap with 0015)
2. Archive or remove old docs/tasks/ folder
3. Update docs/product/features/ to reference tasks or remove folder
4. Ensure all agents/developers read taskboard.csv before starting work

## Validation
- ✅ Only one DOING task (0011)
- ✅ All tasks follow numeric naming (0001-0026)
- ✅ Taskboard.csv has correct schema
- ✅ Template updated to AGENTS.md standard
- ✅ Evidence folder created with documentation

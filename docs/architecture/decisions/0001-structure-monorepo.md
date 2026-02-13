# ADR 0001 - Structure Monorepo

- Status: accepted
- Date: 2026-02-13

## Context
The project is evolving quickly with backend, frontend, infra, and docs changing together. We need consistency for parallel contributors/agents and predictable local development.

## Decision
Adopt a monorepo with clear top-level boundaries:
1. `backend/`
2. `frontend/`
3. `infra/`
4. `docs/`
5. `desktop/` (post-MVP)
6. `.github/` (collaboration and automation standards)

## Consequences
Positive:
1. Single source of truth for architecture and product docs.
2. Easier cross-layer changes and review.
3. Shared CI and contribution standards.

Tradeoffs:
1. Potentially larger PRs if boundaries are ignored.
2. Requires discipline on scope and task ownership.

## Notes
Maintain compatibility paths during transition to avoid breaking current build/runtime.

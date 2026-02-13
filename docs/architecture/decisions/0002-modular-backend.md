# ADR 0002 - Modular Backend by Domain

- Status: accepted
- Date: 2026-02-13

## Context
`backend/cmd/api/main.go` accumulated too much domain logic. This reduced maintainability, testability, and safe parallel development.

## Decision
Use domain modules with the structure:
`backend/internal/modules/<domain>/{domain,usecase,infra,transport}`

Guidelines:
1. `transport` parses requests and maps responses.
2. `usecase` enforces business rules.
3. `infra` handles SQL/external integrations.
4. `domain` defines entities/interfaces/errors.

## Consequences
Positive:
1. Better separation of concerns.
2. Easier incremental refactor and tests.
3. Safer ownership split across contributors.

Tradeoffs:
1. More files/packages to navigate.
2. Requires strict boundary discipline.

## Rollout
Incremental migration:
1. Bootstrap/router/middleware first.
2. Auth module as first complete example.
3. Move remaining domains in small validated steps.

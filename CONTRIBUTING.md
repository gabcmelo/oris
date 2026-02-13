# Contributing

## Development flow
1. Create a branch from latest main.
2. Pick or create a task file under `docs/tasks/`.
3. Implement in small, focused changes.
4. Run local validation:
   - `cd backend && go test ./...`
   - `cd backend && go build ./cmd/api`
   - `cd frontend && npm install && npm run build`
5. Update docs when behavior/architecture changes.
6. Open PR using the repository template.

## Coding expectations
1. Preserve API compatibility unless explicitly approved.
2. Keep handlers thin and business rules in usecases.
3. Avoid unrelated refactors in the same PR.
4. Keep shared modules minimal and intentional.

## Documentation expectations
1. Architecture decisions must go to `docs/architecture/decisions/`.
2. Feature decisions must be documented in `docs/product/features/`.
3. Sprint/task progress must be tracked in `docs/product/tasks/` and/or `docs/tasks/`.

# Code Style Standards

## General
1. Keep changes scoped and minimal.
2. Prefer explicit names and straightforward control flow.
3. Avoid dead code and temporary debug artifacts.

## Go
1. Run `gofmt` on changed files.
2. Keep handlers thin; move business logic to usecase layer.
3. Keep SQL in infra/repository packages whenever possible.

## Frontend
1. Organize by `app/pages/features/entities/shared`.
2. Keep components focused; avoid monolithic files when extending features.
3. Keep API calls centralized in shared API helpers over time.

# Runbook - Deploy (Self-host Basic)

## Objective
Deploy SafeGuild stack on a host with Docker Compose.

## Steps
1. Configure environment values in compose/env.
2. Run stack:
   - `docker compose -f infra/docker-compose.yml up -d --build`
3. Validate:
   - API health endpoint
   - frontend availability

## Upgrade
1. Linux host:
   - `bash scripts/upgrade.sh`
2. Windows host:
   - `powershell -ExecutionPolicy Bypass -File scripts/upgrade.ps1`

## Rollback note
Current scripts provide basic fallback but not full image-version rollback yet.

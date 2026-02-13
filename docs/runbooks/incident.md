# Runbook - Incident Response (Basic)

## Immediate triage
1. Identify affected component (api/web/db/livekit).
2. Capture logs and timestamps.
3. Assess user impact and severity.

## Quick diagnostics
1. `docker compose -f infra/docker-compose.yml ps`
2. `docker logs --tail 200 oris-api`
3. `docker logs --tail 200 oris-web`

## Mitigation options
1. Restart affected service.
2. Roll back to last known stable image/tag (manual until automated rollback is improved).
3. Communicate status and workaround to users.

## Post-incident
1. Document root cause.
2. Add regression test and/or guardrail.
3. Update runbooks/ADRs if process changed.

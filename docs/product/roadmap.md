# Oris Roadmap

## Document Purpose
1. This document defines phase sequencing.
2. This document includes a dated status snapshot for clarity.
3. Task-level execution priority lives in `docs/product/tasks/backlog.md`.

## Status Snapshot (2026-02-13)

Completed in code:
1. Local auth and session endpoints (`/auth/login`, `/auth/refresh`, `/auth/logout`).
2. Communities, channels, membership roles, moderation actions, audit logs.
3. Voice token endpoint with membership and channel-type guardrails.
4. WebSocket origin guard and HTTP CORS allowlist (`APP_ALLOWED_ORIGINS`).
5. Telemetry admin RBAC and invite usage accounting fix.

Partially complete:
1. Protected Community Mode exists as `safe_mode_enabled` flag, but no automatic operational policies yet.

Not complete:
1. Rate limiting baseline.
2. Anti-abuse and quarantine operational flows.
3. Integration test suite for guardrails and regressions.
4. Extensions v1 implementation (current endpoint is still stub).

## Execution Order

### Phase 0 - Foundations (active)
Outcome:
1. A developer can self-host Oris, run core community flows, and rely on minimum safety controls.

Remaining gate items:
1. Rate limiting on sensitive endpoints.
2. Protected Community Mode operational rules.
3. Integration tests for auth, invite, moderation, voice token, and websocket guardrails.

### Phase 1 - MVP
Outcome:
1. Oris becomes viable for real creator-led communities.

Core goals:
1. Onboarding and invite experience polish.
2. Presence and realtime UX resilience.
3. Voice quality and reconnect hardening.
4. Extensions v1 (webhook-first) only after Phase 0 gate is done.

### Phase 2 - Expansion
Outcome:
1. Expand beyond initial niche without changing core architecture.

Goals:
1. Targeted ecosystem plugins.
2. Better admin and observability UX.
3. Optional SSO path.

### Phase 3 - Scale and Hardening
Outcome:
1. Reliable operation at larger scale.

Goals:
1. Multi-node and orchestration path.
2. Advanced abuse mitigation options.
3. Security hardening and operational controls.

## Explicit Non-Goals Before PMF
1. Deep console platform integrations.
2. Feature parity race with large communication suites.
3. Marketplace-scale plugin surface before foundation maturity.

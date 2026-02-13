# SafeGuild / Voz do Povo - MVP + Post-MVP Consolidated Plan

## Goal
Deliver a practical MVP fast, self-hosted first, with safe defaults and clear path to desktop and observability.

## Decisions locked
1. MVP web-first now (React + Vite).
2. Electron right after MVP, not before.
3. Backend: Go + Gin.
4. Voice: LiveKit + coturn path.
5. Data: PostgreSQL + Redis.
6. Telemetry: opt-in only, OTLP, no PII.
7. Update strategy: stable channel + one-click upgrade scripts.

## MVP scope
1. Local auth with JWT.
2. Communities, members, text/voice channels.
3. Minimal text chat realtime.
4. Voice join flow via LiveKit token endpoint.
5. Moderation basic: kick, mute, ban.
6. Audit log + JSON/CSV export.
7. Safe mode default for minors context (no DM by default, expiring invites).
8. Docker Compose deploy local/VPS.
9. Version and upgrade-check endpoints.

## Immediate post-MVP
1. Electron desktop wrapper and installers.
2. Desktop auto-update stable channel.
3. Telemetry agent opt-in + admin status.
4. Better persistence layer (replace in-memory runtime state).

## What to reuse from market research

### P0 (use now)
1. Keep LiveKit as media backend abstraction baseline.
2. Include coturn support path for difficult networks.
3. Keep OpenTelemetry Collector profile in compose.
4. Keep licensing simple and transparent for core.
5. Launch sequence for technical trust first: GitHub + Docker + selfhosted communities.

### P1 (next month)
1. OIDC provider compatibility (Keycloak optional).
2. Backup/restore hardening (restic-style workflow).
3. Stronger governance surface: immutable-like audit behavior and better moderation workflows.

### P2 (later)
1. Optional federation experiments.
2. Optional bridges and search modules as add-ons.

## Anti-patterns to avoid
1. Shipping desktop before core reliability.
2. Promising full RTC depth before ops quality.
3. Mixing unclear license boundaries (open-core confusion) too early.
4. Opt-out telemetry by default.

## Acceptance checkpoints
1. Fresh machine setup in <= 30 min.
2. 5 concurrent users in same voice flow.
3. Moderation actions logged and exportable.
4. Telemetry disabled by default and explicit opt-in works.
5. Upgrade script performs backup and health check.

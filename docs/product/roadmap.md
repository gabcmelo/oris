# Oris Roadmap

This roadmap is outcome-driven. Features exist to achieve product outcomes, not to collect checkboxes.
Build your own voice community platform. Open-source, self-hosted real-time infrastructure for creators and gaming communities, powered by WebRTC.

## Phase 0 - Foundations (v0.1)

Outcome: a developer can run Oris locally and on a VPS, join voice, and moderate basic abuse.

Core:
1. Auth (local) + sessions.
2. Communities/servers, channels, roles/permissions (minimum viable).
3. Voice join/leave, speaking state, mute/deafen.
4. Web client (PWA-first) baseline UX.

Safety baseline:
1. Rate limiting (login, join, message, invite).
2. Basic moderation actions: mute/ban/kick.
3. Audit log (who did what, when).
4. Reporting (minimal flow).

Ops baseline:
1. Docker Compose one-command up.
2. Basic metrics + logs (at least API + signaling + media health indicators).
3. Minimal admin panel (or CLI) for initial operations.

Deliverable:
1. `docker compose up -d` working guide + `.env.example`.
2. Protected Community Mode configuration template (defaults).

## Phase 1 - MVP (v1.0)

Outcome: Oris is viable for real communities (creators/devs/gamers) with safety defaults.

Core improvements:
1. Better onboarding and invite flows.
2. Presence and realtime UX polish.
3. Resilient signaling & reconnect logic.
4. Voice quality tuning (codec/bitrate defaults, jitter handling where possible).

Safety by default:
1. Protected Community Mode as a first-class template.
2. Quarantine flow for new accounts.
3. Anti-raid presets (join storms, spam bursts).
4. Case management basics (reports, actions, notes).

Extensibility:
1. Oris Extensions v1 (events + scopes + auditing + rate limiting).
2. Plugin packaging + install/enable/disable.
3. Official plugins: Live Presence + Alerts (YouTube/Twitch) and Safety Pack.

Success metrics tracked:
1. Voice: join time, session success rate, TURN usage ratio, packet loss proxy.
2. Safety: raids blocked, false positives review rate.
3. Product: D7 retention, weekly active voice participants.

## Phase 2 - Growth (v1.x)

Outcome: lower-friction migration and richer niche value without bloating core.

1. Discord import tools (structure + roles mapping where feasible).
2. GitHub plugin (issues/PR/releases).
3. Steam presence-lite plugin (status/roles/events).
4. Better admin UI and moderation workflows.
5. Observability dashboards shipped by default (Grafana templates).
6. Optional SSO (OIDC) for orgs.

## Phase 3 - Scale & Hardening (v2.0)

Outcome: Oris can run reliably at larger scale and stricter environments.

1. Kubernetes deployment path (charts/manifests).
2. Multi-node scaling strategy (clear separation of control/media plane).
3. Advanced abuse mitigation options.
4. Optional E2EE advanced mode (with clear trade-offs and compatibility matrix).
5. Security hardening guides and automated checks.

## Explicit Non-Goals (MVP)

1. Deep integrations with console/platform ecosystems (Xbox/PlayStation/Nintendo/Epic).
2. All-in-one everything scope.
3. Partner-dependent features as core promises.

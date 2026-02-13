# Oris Extensions - Initial Spec (v1)

## Goals

Oris Extensions provides a secure way to extend Oris without bloating core.

Must provide:
1. Event-driven extension model.
2. Strong permission model (scopes).
3. Auditable actions.
4. Rate limiting and isolation.
5. Simple developer experience for building plugins.

Non-goals (v1):
1. A massive marketplace.
2. Untrusted arbitrary code execution on server without isolation.
3. Complex multi-language runtime support.

## Architecture Overview

Extension types (v1):
1. Webhook Extensions (recommended for MVP).
2. Worker Extensions (later).

MVP should implement Webhook Extensions first.

Webhook Extensions:
1. Oris emits signed events to HTTPS endpoints.
2. Extensions act via Oris REST API using scoped tokens.

Worker Extensions (later):
1. Controlled runtime hosted by Oris (sandboxed).
2. Higher complexity, delayed until hardening.

## Event Model

Events are emitted from Oris to extension endpoints.

### Delivery Guarantees

1. At-least-once delivery (webhooks may retry).
2. Event IDs are unique; extensions must be idempotent.
3. Retry policy: exponential backoff + max attempts.
4. Dead-letter: failures recorded for operator review.

### Security

1. Each webhook request is signed (HMAC) using extension secret.
2. Timestamp + nonce included to prevent replay.
3. Extensions must verify signature + freshness window.

### Event Envelope (Example)

```json
{
  "id": "evt_01HZY...ABC",
  "type": "voice.session.started",
  "occurred_at": "2026-02-12T22:00:00Z",
  "community_id": "com_...",
  "actor": {
    "user_id": "usr_...",
    "role": "member"
  },
  "data": {
    "channel_id": "chn_...",
    "session_id": "vs_..."
  }
}
```

## Scopes And Permissions

Extensions receive a scoped token (OAuth-like but internal).

Examples of scopes:
1. `community.read`
2. `community.write`
3. `members.read`
4. `members.moderate` (high risk)
5. `voice.read`
6. `voice.manage`
7. `notifications.send`
8. `reports.read`
9. `reports.write`
10. `audit.read`

Scopes must be:
1. Explicitly granted by operator/admin.
2. Visible in UI.
3. Logged whenever used in sensitive actions.

## Rate Limiting

1. Per-extension rate limits for webhook deliveries and API calls made by extension tokens.
2. Operator can set stricter limits.
3. Excess triggers audit + optional auto-disable.

## Audit And Observability

Every extension action produces an audit record:
1. `extension_id`
2. `action`
3. `scope_used`
4. `target`
5. `timestamp`
6. `result` (success/failure)
7. `correlation_id` (ties back to event id)

Operators must be able to:
1. See installed extensions.
2. Enable/disable.
3. View logs/errors.
4. Rotate secrets.

## MVP Event Catalog

Community & membership:
1. `member.joined`
2. `member.left`
3. `invite.created`
4. `invite.used`

Voice:
1. `voice.channel.joined`
2. `voice.channel.left`
3. `voice.session.started`
4. `voice.session.ended`

Moderation & safety:
1. `report.created`
2. `moderation.action.executed`
3. `security.raid.detected`
4. `security.quarantine.applied`

## Official MVP Plugins

### Plugin 1 - Live Presence + Alerts (YouTube/Twitch)

Goal: notify community when creator goes live and optionally create/join voice flows.

Minimal features:
1. Admin config: channel URL, notification channel, message template.
2. On live detection, post announcement in text channel.
3. Set live-now indicator (optional).
4. Optionally create/highlight a voice channel (Watch Party / Live Hangout).

Implementation approach (MVP-friendly):
1. External service polls YouTube/Twitch or uses supported webhooks.
2. Service calls Oris API with `notifications.send` and `voice.manage` if enabled.

Safety:
1. Avoid spamming; enforce cooldown windows.
2. Respect per-community config.

### Plugin 2 - Safety Pack

Goal: give communities immediate protection without requiring security expertise.

Minimal features:
1. Anti-raid presets.
2. Quarantine mode for new accounts.
3. Case management basics (enrich `report.created` with structured actions and notes).
4. Operator transparency on rules/signals, triggers, and overrides.

Data ethics / privacy:
1. Device/behavior signals must be minimally invasive.
2. Device/behavior signals must be transparent.
3. Device/behavior signals must be configurable.
4. Device/behavior signals must be logged.
5. Avoid unjustified profiling; prefer simple explainable rules in MVP.

## Implementation Checklist (MVP)

1. [ ] Extension registry table (installed extensions, secrets, scopes, status).
2. [ ] Webhook signer + verification docs.
3. [ ] Retry system + dead-letter storage.
4. [ ] Scoped tokens + middleware enforcement.
5. [ ] Rate limiting per extension.
6. [ ] Audit log integration.
7. [ ] Admin UI (minimal) for managing extensions.
8. [ ] Two official plugin implementations + docs.

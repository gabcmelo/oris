# Oris - Vision

## Document Purpose
1. This document defines long-term product direction.
2. This document is not a delivery status tracker.
3. Current implementation status lives in:
   - `docs/product/tasks/backlog.md`
   - `docs/product/roadmap.md` (status snapshot section)

## Product Thesis
Oris is a self-hosted, open-source voice community platform focused on:
1. Voice quality (low latency, stable sessions).
2. Safety by default (especially for communities with minors).
3. Extensibility without turning core into a bloated suite.

Oris does not compete on feature parity. Oris competes on reliability, safety, and control.

## Initial Audience
Primary entry point:
1. Creator-led gaming communities.

Planned expansion after product-market fit:
1. Developer communities.
2. Online communities with governance needs.
3. Startups and modern teams.

## Positioning
Build your own voice community platform.
Open-source, self-hosted, real-time voice infrastructure powered by WebRTC.

## MVP Principles
MVP is:
1. Voice-first.
2. Safety-first.
3. Operationally simple to self-host.

MVP is not:
1. A giant integrations catalog.
2. A general all-in-one communication suite.
3. A promise to support every niche from day one.

## Safety Principles
Safety is a product feature, not a setup burden.

Baseline requirements:
1. Role and permission model.
2. Rate limiting and anti-abuse controls.
3. Moderation actions and audit logs.
4. Protected Community Mode defaults for communities with minors.
5. Data minimization and transparent policies.

## Extensibility Principles
Extensions are part of the strategy, but come after foundation hardening.

MVP target:
1. Secure extension framework (events, scopes, audit, limits).
2. Very small official plugin set with immediate utility.

## Success Metrics
1. Voice quality:
   - Join time.
   - Session success rate.
   - Packet loss proxy.
2. Safety:
   - Moderation response time.
   - Raid and abuse mitigation effectiveness.
3. Product:
   - D7 retention.
   - Weekly active voice participants.
4. Operations:
   - Crash-free sessions.
   - Observability coverage.
   - Upgrade reliability.

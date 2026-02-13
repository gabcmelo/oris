# Oris - Vision

Oris is a self-hosted, open-source voice community platform focused on three pillars:

1. Excellent voice quality (low-latency, stable, production-grade).
2. Safety by default (especially for communities with minors).
3. Extensibility (plugins/extensions without turning Oris into a bloated all-in-one).

Oris is not trying to copy existing platforms feature-by-feature.
Oris aims to be the safest and most controllable foundation for real-time voice communities.

## Target Users

Primary audiences:
1. Streamers / YouTubers / creators (community + live presence + alerts).
2. Developer communities (events, GitHub integration, structured moderation).
3. Gaming communities (roles, events, presence, lightweight game integrations).
4. Startups/companies that need voice + moderation + governance.

## Product Positioning

Build your own voice community platform. Open-source, self-hosted real-time infrastructure powered by WebRTC.

The wedge (the why Oris wins):
1. Voice quality comparable to best-in-class.
2. Moderation + auditability taken seriously.
3. Defaults designed for safer communities (including minors).
4. A plugin system that is secure from day one.

## MVP Scope Principles

MVP is:
1. Voice-first.
2. Safety-first.
3. Extensible, but with few official plugins.

MVP is not:
1. A giant integration zoo (consoles/platforms/partner APIs).
2. All-in-one feature creep.
3. A promise of everything to everyone.

## Safety By Default

Safety is a product feature, not a configuration burden.

MVP must ship with:
1. Strong role/permission model.
2. Rate limits and anti-abuse mechanisms.
3. Moderation tools + audit logs.
4. A Protected Community Mode (recommended defaults for minors).

## Legal And Policy Baseline (Brazil / LGPD)

Oris may be deployed for communities with minors. Therefore:
1. A clear Privacy Policy template and Safety Policy must exist.
2. Data minimization is mandatory by design.
3. Any behavioral detection must be transparent, configurable, and avoid unjustified profiling.
4. Operators (server admins) must understand their responsibilities and retention choices.

Oris provides defaults and documentation, but operators remain responsible for deployment and governance.

## Extensibility (Oris Extensions)

Oris supports an extension system from day one, but the MVP ships with only:
1. A secure plugin framework (events + scopes + audit + rate limits).
2. Two official plugins that deliver immediate value.

### MVP Official Plugins

Live Presence + Alerts:
1. YouTube/Twitch go-live notifications.
2. Optional automatic voice channel creation for watch-party or post-live hangouts.
3. Visible live-now presence indicators.

Safety Pack:
1. Raid/flood protection presets.
2. Quarantine mode for new accounts.
3. Case management basics (reports, actions, notes).
4. Audit-friendly moderation exports.

## MVP Success Metrics

Oris is successful if it proves:
1. Voice quality: stable sessions, low join time, low packet loss.
2. Retention: D7 retention, active voice minutes per user.
3. Moderation efficiency: time to act, cases resolved, raids blocked.
4. Operational health: crash-free sessions, observability coverage, upgrade reliability.

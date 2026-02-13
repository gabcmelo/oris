# Usable Software Implementation Plan (Oris)

## 1. Product Focus (locked)

Build an open-source community communication tool that is:
1. Safe by default for kids, teens, and adults.
2. Practical for gamers, streamers, geek/dev communities, and home-office groups.
3. Easy to self-host (VPS/Docker) and easy to use for non-technical users.

Core value proposition:
- "Create your own safe community in minutes, with voice, text, moderation, and auditability."

## 2. What "minimally usable" means

A minimally usable release is not just "API works". It must be:
1. Onboarding clear in under 3 minutes.
2. Visual interface understandable without technical knowledge.
3. Core actions one-click discoverable:
   - create account
   - create/join community
   - join text channel
   - join voice channel
   - report/moderate abuse
4. Safe defaults enabled without manual tuning.
5. Stable enough for a small real group test (5-20 users).

## 3. Scope for Usable v0.1 (next implementation target)

### 3.1 Must-have product features (P0)
1. Auth with secure passwords + session handling.
2. Community creation and invite flow.
3. Text channel with realtime messages.
4. Voice channel join flow with clear UI states.
5. Basic moderation actions (kick/mute/ban).
6. Audit timeline for admin actions.
7. "Safe Community Mode" preset on by default.
8. Telemetry opt-in explicit toggle in admin panel.

### 3.2 Must-have UX/UI quality (P0)
1. Replace current raw console page with app layout:
   - left sidebar: communities
   - middle sidebar: channels
   - main panel: text/voice room
   - right panel: members/mod actions (contextual)
2. Clear visual hierarchy and accessible controls.
3. Empty states and guided actions (no dead-end screens).
4. Error messages human-readable.
5. Mobile-responsive baseline.

### 3.3 Must-have operational quality (P0)
1. Data persistence in PostgreSQL for all core entities.
2. Realtime websocket fanout stable after reconnect.
3. Docker compose startup and restart without data loss.
4. Health endpoint + basic logs.

## 4. Technical implementation phases

## Phase A - Reliability foundation (backend/data)
Goal: close backend gaps before UI polish.

Deliverables:
1. Replace in-memory stores with PostgreSQL repositories:
   - users
   - refresh_tokens
   - communities
   - community_members
   - channels
   - messages
   - invites
   - audit_log
   - telemetry_settings
2. Add DB transactions for moderation + audit log writes.
3. Ensure role checks run against DB state.
4. Add pagination to message/audit listing.

Acceptance:
- Restarting API does not lose users/communities/messages.
- Moderation actions persist and remain enforced.

## Phase B - Realtime and voice usability
Goal: realtime feels alive and voice flow is understandable.

Deliverables:
1. Stable WS channel subscriptions with reconnect strategy.
2. Message broadcast to all connected clients in same channel.
3. Voice state UI:
   - connecting
   - connected
   - muted/deafened
   - disconnected
4. LiveKit token endpoint verified and used by frontend voice client wiring.

Acceptance:
- Two browser sessions exchange messages in realtime.
- Voice join/leave states visible and consistent.

## Phase C - UX/UI redesign (minimum quality bar)
Goal: move from technical console to usable app.

Deliverables:
1. New design system tokens:
   - spacing scale
   - typography scale
   - color roles (no random hardcoded colors)
2. App shell + channel views.
3. Onboarding flow:
   - first login -> create or join community
4. Moderation UI panel integrated in member list.
5. Export audit action available in admin settings area.

Acceptance:
- A non-technical tester can complete core flow without guidance.

## Phase D - Safety defaults and trust surfaces
Goal: make safety visible and useful.

Deliverables:
1. Safe mode defaults:
   - DMs disabled by default
   - invite expiration prefilled
   - stricter default permissions
2. Reporting entrypoint in message/user context menu.
3. Transparency page in-app:
   - what is collected
   - what is not collected
   - telemetry opt-in status

Acceptance:
- Admin can verify and control safety/telemetry settings from UI.

## 5. Backlog by priority

### P0 (immediate)
1. DB persistence migration for backend core.
2. Realtime message flow hardening.
3. UI shell redesign and onboarding.
4. Safe mode + moderation UI integration.

### P1 (after usable v0.1)
1. Electron desktop wrapper with installer.
2. Better admin dashboard for performance metrics.
3. OIDC optional login.

### P2 (later)
1. Plugin/integration marketplace style model.
2. Federation experiments.
3. Advanced anti-abuse automation.

## 6. Usability test protocol (what to run with your group)

Test group: 5-10 people (mixed technical/non-technical).

Scenarios:
1. New user creates account and joins community from invite.
2. User sends messages and receives realtime updates.
3. User joins voice and toggles mute/deafen.
4. Moderator removes abusive user and checks audit entry.
5. Admin exports audit data.

Success thresholds:
1. 80% complete onboarding without help.
2. Realtime latency perceived as instant for text (<1s typical).
3. No data loss after service restart.
4. Moderation action reflected in <2s.

## 7. Definition of Done for "minimally usable"

Release is considered minimally usable only if all are true:
1. Core data persistence works across restart.
2. UI is navigable and understandable for non-technical users.
3. Text + voice + moderation work in one coherent flow.
4. Safety defaults are active and visible.
5. Setup docs allow clean install in <= 30 minutes.

## 8. Next execution order (recommended)

1. Backend DB persistence first (Phase A).
2. Realtime/voice stability (Phase B).
3. UI redesign + onboarding (Phase C).
4. Safety and trust surfaces (Phase D).

This order avoids polishing UI over unstable core behavior.

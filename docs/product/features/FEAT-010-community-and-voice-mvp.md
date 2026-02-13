# FEAT-010 - Community + Voice MVP

## Summary
Deliver a usable community experience with text channels, voice channel join flow, invites, and baseline moderation.

## Scope
1. Register/login/refresh/logout.
2. Create/list communities and channels.
3. Post/list messages in text channels.
4. Voice token issue and LiveKit join from UI.
5. Basic moderation (`kick`, `mute`, `ban`) and audit export.

## Out of scope
1. Federation.
2. Full desktop client packaging.
3. Advanced trust/safety ML moderation.

## Endpoints
1. `/api/v1/auth/*`
2. `/api/v1/communities*`
3. `/api/v1/channels/*`
4. `/api/v1/voice/token`
5. `/api/v1/ws/:channelId`

## Edge cases
1. Invalid/expired invite.
2. Muted/banned member trying to post message.
3. Voice token request for unauthorized channel (hardening task).

## Acceptance criteria
1. User can create community and channel in web UI.
2. Two users can exchange realtime messages.
3. Voice token is issued and room connection starts.
4. Moderation actions are persisted to audit log.

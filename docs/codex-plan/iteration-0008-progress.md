# Iteration 0008 Progress

## Implemented
1. Presence and voice realtime block finalized:
   - `GET /api/v1/channels/:channelId/presence`
   - WebSocket presence broadcast on connect/disconnect
   - LiveKit voice controls in frontend (join/leave/mute)
2. Web usability polish:
   - session persistence with `localStorage`
   - message UI with author name + timestamp + auto-scroll
   - improved voice panel and responsive layout
   - logout action from UI

## Verified
1. `docker compose build` and `up -d` for `api` and `web`.
2. Smoke flow API:
   - register/login
   - create community and channels
   - create/join invite
   - send message
   - generate voice token
   - moderation guard (self block) and moderation action
3. Frontend production build:
   - `npm run build` inside `oris-web`.

## Run targets
- Web: http://localhost:5173
- API: http://localhost:8080
- LiveKit: ws://localhost:7880

# Iteration 0007 Progress

## Implemented
1. Backend member listing endpoint for selected community.
2. Moderation safety guards:
   - cannot moderate yourself
   - cannot moderate same or higher role
3. Frontend additions:
   - invite creation and join by code
   - members panel
   - inline moderation actions per member

## Verified
1. Create invite -> success.
2. Join community via invite -> success.
3. Members endpoint -> returns joined users.
4. Self moderation attempt -> blocked.

## Current run targets
- Web: http://localhost:5173
- API: http://localhost:8080

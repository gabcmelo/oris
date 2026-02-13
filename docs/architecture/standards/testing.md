# Testing Standards

## Backend
1. Required baseline:
   - `go test ./...`
   - `go build ./cmd/api`
2. Integration checks for critical flows:
   - auth
   - invite join
   - moderation permissions
   - voice token authorization rules

## Frontend
1. Required baseline:
   - `npm install`
   - `npm run build`
2. Manual smoke:
   - login/register
   - create community/channel
   - send/receive messages
   - voice join flow

## Infra
1. Compose up/down should be reproducible.
2. Health endpoint and API version endpoint must respond.

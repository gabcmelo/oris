# Runbook - Local Development

## Prerequisites
1. Docker + Docker Compose
2. Go 1.22 (for local backend commands)
3. Node 20+ (for local frontend commands)

## Start full stack
1. `docker compose -f infra/docker-compose.yml up -d --build`
2. Open:
   - `http://localhost:5173`
   - `http://localhost:8080/healthz`

## Backend only
1. `cd backend`
2. `go test ./...`
3. `go build ./cmd/api`

## Frontend only
1. `cd frontend`
2. `npm install`
3. `npm run dev`

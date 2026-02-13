# Architecture Overview

## System summary
1. Web-first real-time communication platform.
2. Backend API in Go + Gin.
3. Frontend in React + Vite.
4. Voice through LiveKit.
5. Persistence in PostgreSQL; Redis available for realtime/cache scenarios.

## High-level components
1. `frontend`:
   - auth, community navigation, chat/voice controls.
2. `backend`:
   - auth, communities, channels, messages, invites, moderation, audit.
3. `infra`:
   - docker compose stack for local/dev self-host.
4. `telemetry-agent`:
   - opt-in metrics prototype.

## Architectural direction
1. Backend moves to modular domain-based structure.
2. Handlers stay thin; usecases own business rules.
3. DB and provider concerns live in infra layer.
4. Documentation-first governance for multi-agent collaboration.

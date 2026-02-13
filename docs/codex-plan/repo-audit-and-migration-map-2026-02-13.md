# Repo Audit + Migration Map - 2026-02-13

## 1) Auditoria rapida do estado atual

## Estrutura existente (real)
1. `backend/`
   - `cmd/api/main.go` (bootstrap + grande parte da logica de dominio ainda no mesmo arquivo)
   - `internal/config/config.go`
   - `internal/http/router/router.go`
   - `internal/http/middleware/auth.go`
   - `internal/platform/jwtutil/manager.go`
   - `internal/modules/auth/handler.go` (primeiro modulo extraido)
   - `Dockerfile`, `go.mod`, `go.sum`
2. `frontend/`
   - `src/App.jsx`, `src/main.jsx`, `src/styles.css`
   - `package.json`, `vite.config.js`, `Dockerfile`
3. `infra/`
   - `docker-compose.yml` (api/web/postgres/redis/livekit/telemetry)
4. `database/`
   - `migrations/0001_init.sql`
   - `migrations/0002_refresh_tokens.sql`
5. `docs/`
   - `codex-plan/*`, `tasks/*`, `gpt-researchs/*`
6. `.github/`
   - `copilot-instructions.md`
7. `telemetry-agent/`
   - `agent.py`, `Dockerfile`, `otel-collector-config.yaml`
8. `scripts/`
   - `upgrade.sh`, `upgrade.ps1`

## Onde estao os componentes criticos
1. Entrypoint backend:
   - `backend/cmd/api/main.go`
2. Registro de rotas HTTP:
   - `backend/internal/http/router/router.go`
3. Handlers ainda acoplados:
   - principal volume em `backend/cmd/api/main.go`
4. Modulo ja separado:
   - `backend/internal/modules/auth/handler.go`
5. Migrations:
   - `database/migrations/*.sql`
6. Build frontend:
   - `frontend/package.json` (scripts `dev/build/preview`)
7. Infra local:
   - `infra/docker-compose.yml`
8. CI:
   - inexistente (`.github/workflows` nao existe no estado atual)

## Pontos sensiveis mapeados
1. Go imports sensiveis:
   - modulo `safeguild/backend` deve ser preservado para evitar quebra em massa.
2. Build path backend:
   - Docker usa `go build -o safeguild-api ./cmd/api` (`backend/Dockerfile`).
3. Build path frontend:
   - Vite depende de `frontend/src/main.jsx` e `frontend/src/App.jsx`.
4. Infra path:
   - Compose atual referencia `../backend`, `../frontend`, `../database/migrations`, `../telemetry-agent`.
5. Contratos HTTP:
   - rotas publicas estao centralizadas no router e nao podem ser quebradas durante migracao.
6. Artefatos temporarios:
   - `backend/cmd/api/main.go.tmp` e `backend/cmd/api/main.go.bak` (limpar em etapa de hardening/repositorio).
7. Documentacao baseline:
   - `README.md` ainda contem nota desatualizada sobre estado in-memory.

## 2) Mapa de migracao origem -> destino

## Backend
1. `backend/cmd/api/main.go` -> manter apenas bootstrap (target)
   - logica de dominio deve migrar para `backend/internal/modules/*`.
2. `backend/internal/config/config.go` -> `backend/internal/core/config/` (fase seguinte)
3. `backend/internal/platform/jwtutil/manager.go` -> `backend/internal/core/auth/jwt/manager.go`
4. `backend/internal/http/middleware/auth.go` -> `backend/internal/app/http/middleware/auth.go`
5. `backend/internal/http/router/router.go` -> `backend/internal/app/http/router.go`
6. `backend/internal/modules/auth/handler.go` -> padrao final:
   - `backend/internal/modules/auth/domain/*`
   - `backend/internal/modules/auth/usecase/*`
   - `backend/internal/modules/auth/infra/*`
   - `backend/internal/modules/auth/transport/http/*`
7. Blocos ainda em `main.go` migrar por dominio para:
   - `backend/internal/modules/communities/*`
   - `backend/internal/modules/channels/*`
   - `backend/internal/modules/messages/*`
   - `backend/internal/modules/invites/*`
   - `backend/internal/modules/moderation/*`
   - `backend/internal/modules/audit/*`
   - `backend/internal/modules/voice/*`
   - `backend/internal/modules/telemetry/*`
   - `backend/internal/realtime/*` (hub/ws/presence)

## Frontend
1. `frontend/src/main.jsx` -> `frontend/src/app/main.jsx`
2. `frontend/src/App.jsx` -> split:
   - `frontend/src/app/App.jsx` (shell)
   - `frontend/src/pages/*`
   - `frontend/src/features/*`
   - `frontend/src/entities/*`
   - `frontend/src/shared/*`
3. `frontend/src/styles.css` -> `frontend/src/styles/global.css`
4. `frontend/vite.config.js` -> manter path de entrada compativel durante migracao (sem quebra de build).

## Infra
1. `infra/docker-compose.yml` -> `infra/docker/docker-compose.yml`
2. `database/migrations/*` -> alvo recomendado:
   - `infra/docker/postgres/migrations/*` (ou manter e mapear por volume)
3. `telemetry-agent/*` -> opcao A (manter local atual por baixo risco imediato),
   opcao B (mover para `infra/telemetry/` em etapa posterior).

## Docs e GitHub
1. `docs/codex-plan/*` + `docs/tasks/*` -> manter e complementar com:
   - `docs/architecture/*`
   - `docs/product/*`
   - `docs/runbooks/*`
   - `docs/api/*`
2. `.github/copilot-instructions.md` -> manter e adicionar:
   - `.github/ISSUE_TEMPLATE/*.yml`
   - `.github/PULL_REQUEST_TEMPLATE.md`
   - `.github/workflows/ci.yml`
3. Root files inexistentes para criar:
   - `AGENTS.md`
   - `CONTRIBUTING.md`
   - `CODE_OF_CONDUCT.md`
   - `SECURITY.md`
   - `.gitignore`
   - `.editorconfig`

## 3) Ordem recomendada (segura)
1. Criar estrutura alvo nova sem mover arquivos criticos.
2. Mover backend por dominio em pequenos passos com build/smoke por etapa.
3. Mover frontend por camada (`app/pages/features/entities/shared`) mantendo import paths funcionais.
4. Ajustar compose para novo path em `infra/docker`.
5. Adicionar padroes GitHub/docs/ADRs.
6. Limpeza final (remover `.bak/.tmp`, corrigir README e runbooks).

## 4) Checks de risco obrigatorios por etapa
1. Backend:
   - `go test ./...`
   - `go build ./cmd/api`
2. Frontend:
   - `npm install`
   - `npm run build`
3. Infra:
   - `docker compose -f infra/docker-compose.yml up -d --build` (antes da mudanca de path)
   - apos mover compose: validar novo caminho e atualizar docs.

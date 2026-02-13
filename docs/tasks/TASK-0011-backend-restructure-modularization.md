# TASK-0011-backend-restructure-modularization

- Status: in_progress
- Owner: Codex
- Objetivo: Reestruturar o backend monolitico em arquivo unico para arquitetura modular, mantendo compatibilidade do MVP.

## Escopo
1. Quebrar `main.go` em modulos por responsabilidade.
2. Introduzir camadas `http/service/repository/domain`.
3. Preservar contratos de API atuais.
4. Incluir hardening backend prioritario durante extracao.

## Entregaveis
1. Nova estrutura em `backend/internal/*`.
2. `main.go` reduzido para bootstrap.
3. Handlers separados por dominio.
4. Repositories com queries isoladas.
5. Smoke tests backend passando apos migracao.

## Checklist
- [x] Extrair config + bootstrap router
- [x] Extrair auth module (handler/service/repository)
- [ ] Extrair communities/channels/messages modules
- [ ] Extrair invites/moderation/audit modules
- [ ] Extrair voice/ws/telemetry modules
- [ ] Aplicar hardening P0 em voice/ws/admin telemetry/invites
- [ ] Validar compose build + smoke
- [ ] Atualizar documentacao tecnica

## Dependencias
1. Plano de acao: `docs/codex-plan/backend-restructure-action-plan-2026-02-13.md`

## Progresso atual
1. Fase 1 concluida:
   - `internal/config` criado para centralizar env/config.
   - `internal/http/router` criado para registrar todas as rotas do MVP.
   - `internal/http/middleware/auth` extraido do `main.go`.
   - `internal/platform/jwtutil` criado para gerar/validar JWT de auth.
2. Modulo `auth` extraido:
   - `internal/modules/auth/handler.go` com `Register/Login/Refresh/Logout/Me`.
   - composicao no bootstrap (`main.go`) agora usa handler dedicado de auth.
3. Validacao executada:
   - `go build ./...` no backend.
   - `docker compose -f infra/docker-compose.yml build api`.
   - smoke de auth (register/login/refresh/me) apos extracao.

## Resultado esperado
Backend com base tecnica escalavel para evolucao rapida de features (Electron, integracoes, compliance de seguranca) sem aumentar complexidade acidental.

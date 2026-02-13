# Backend Restructure Action Plan - 2026-02-13

## Objetivo
Retirar o acoplamento do `backend/cmd/api/main.go` (arquivo unico) e evoluir para uma arquitetura modular, testavel e segura, sem quebrar o MVP atual.

## Escopo desta reestruturacao
1. Separar inicializacao HTTP/rotas do dominio de negocio.
2. Introduzir camadas claras: `http`, `service`, `repository`, `domain`.
3. Manter compatibilidade total dos endpoints atuais.
4. Preparar base para testes unitarios e de integracao.

## Estrutura alvo (proposta)
1. `backend/cmd/api/main.go`
   - bootstrap apenas (config, db, router, run).
2. `backend/internal/config`
   - leitura e validacao de env vars.
3. `backend/internal/http`
   - router, middleware auth, handlers por contexto.
4. `backend/internal/service`
   - regras de negocio (auth, comunidade, moderacao, voice, telemetry).
5. `backend/internal/repository`
   - acesso a dados via pgx (queries isoladas).
6. `backend/internal/domain`
   - entidades, DTOs e erros de dominio.
7. `backend/internal/realtime`
   - ws hub/presence.
8. `backend/internal/platform`
   - utilitarios transversais (jwt, password, clock, ids).

## Plano de execucao (incremental)
1. Fase 1 - Bootstrap e router
   - extrair inicializacao do app para `internal/http/router.go`.
   - mover middleware auth e utilitarios JWT para `platform`.
2. Fase 2 - Modulos core
   - extrair `auth`, `communities`, `channels`, `messages` para handlers separados.
   - criar services e repositories minimos para esses modulos.
3. Fase 3 - Modulos de seguranca
   - extrair `moderation`, `audit`, `invites`, `voice`.
   - aplicar guard rails P0 (membership voice, origin ws, admin telemetry).
4. Fase 4 - Realtime e telemetry
   - mover hub/ws/presence para `internal/realtime`.
   - separar telemetry handlers e policy em modulo proprio.
5. Fase 5 - Limpeza e testes
   - reduzir `main.go` para bootstrap.
   - adicionar testes de regressao para auth/invite/moderation/voice/ws.

## Ordem de entrega recomendada
1. PR/logico 1: estrutura + router + auth middleware extraidos.
2. PR/logico 2: communities/channels/messages extraidos.
3. PR/logico 3: moderation/invites/voice/ws/telemetry extraidos + hardening P0.
4. PR/logico 4: testes e ajuste final de docs.

## Criterios de aceite
1. `backend/cmd/api/main.go` com no maximo ~120 linhas.
2. Nenhum endpoint removido ou contrato quebrado no MVP.
3. Build e smoke test existentes continuam passando.
4. Cobertura minima de testes para modulos criticos (auth, invite, voice guard).
5. Zero regressao funcional em fluxo: register/login -> comunidade -> canal -> chat -> voice token.

## Riscos
1. Regressao silenciosa por mover query inline para repository.
2. Aumento de tempo se tentar “reescrever tudo” de uma vez.
3. Conflito com evolucoes paralelas de frontend.

## Mitigacao
1. Reestruturacao por fases pequenas e validacao em cada fase.
2. Manter contratos HTTP identicos.
3. Rodar smoke scripts apos cada etapa.

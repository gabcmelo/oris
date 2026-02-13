# TASK-0014-monorepo-structure-and-standards

- Status: done
- Owner: Codex
- Objetivo: Aplicar estrutura monorepo padrao (backend/frontend/infra/docs/.github/desktop), criar padroes de contribuicao e manter compatibilidade de build/execucao.

## Escopo
1. Criar estrutura alvo de diretorios sem quebrar runtime.
2. Ajustar paths e wrappers de compatibilidade quando necessario.
3. Criar docs de arquitetura/produto/runbooks/api e ADRs.
4. Padronizar GitHub templates (issues/PR/workflow CI).
5. Atualizar README raiz e criar AGENTS.md.

## Checklist
- [x] Estrutura alvo criada (sem quebra)
- [x] Backend alinhado para `internal/app`, `internal/core` e modulo exemplo
- [x] Frontend reorganizado em `app/pages/features/entities/shared`
- [x] Infra padronizada em `infra/docker`
- [x] Docs e ADRs criados
- [x] GitHub templates + CI criados
- [x] Verificacao final de build/test/compose

## Dependencias
1. `docs/codex-plan/repo-audit-and-migration-map-2026-02-13.md`
2. `docs/codex-plan/backend-restructure-action-plan-2026-02-13.md`

## Resultado
1. Estrutura monorepo padrao criada com padroes de colaboracao e CI.
2. Backend com trilha `internal/app` + `internal/core` e modulo `auth` em `domain/usecase/infra/transport`.
3. Frontend reorganizado sem quebrar build, via camada de compatibilidade.
4. Novo compose padrao em `infra/docker/docker-compose.yml` validado em runtime.
5. Conjunto documental completo: AGENTS, ADRs, runbooks, templates de produto e templates GitHub.

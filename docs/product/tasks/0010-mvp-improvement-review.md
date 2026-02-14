# TASK-0010-mvp-improvement-review

- Status: done
- Owner: Codex
- Objetivo: Revisar tecnicamente o MVP atual e mapear melhorias priorizadas para evoluir para um produto utilizavel e seguro.

## Entregaveis
- Diagnostico com riscos tecnicos observados no estado atual
- Lista priorizada de melhorias (P0/P1/P2)
- Evidencias de validacao rapida para riscos criticos

## Checklist
- [x] Revisao backend (auth, voz, ws, moderacao, telemetry)
- [x] Revisao frontend (fluxo auth, runtime, API integration)
- [x] Revisao infra/scripts (compose, upgrade, docs)
- [x] Consolidacao de backlog de melhoria

## Resultado
### P0 (corrigir antes de abrir testes ampliados)
1. Endpoints de voz permitem emissao de token sem validar membership de canal/comunidade e aceitam `userId` arbitrario.
2. CORS/preflight ausente para fluxo web default (`5173` -> `8080`), quebrando navegador.
3. `ws` aceita qualquer origem (`CheckOrigin=true`) sem restricao.
4. Convite incrementa uso mesmo para usuario que ja era membro.
5. Endpoint de telemetry admin pode ser alterado por qualquer usuario autenticado.

### P1 (primeira semana apos P0)
1. Implementar refresh token no frontend + logout revogando refresh corretamente.
2. Bloquear mensagens em canais de voz.
3. Ajustar script de upgrade para rollback real (imagem/tag anterior + restore opcional de backup).
4. Atualizar README para refletir estado real do projeto.

### P2 (hardening)
1. Adicionar FKs e indices de performance nas tabelas principais.
2. Adicionar rate limiting e trilha de seguranca para auth/ws.
3. Introduzir testes automatizados API (auth, invite, moderation, voice guardrails).

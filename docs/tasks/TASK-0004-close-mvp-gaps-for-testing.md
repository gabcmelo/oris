# TASK-0004-close-mvp-gaps-for-testing

- Status: done
- Owner: Codex
- Objetivo: Fechar lacunas restantes do MVP para teste real imediato.

## Entregaveis
- Persistencia PostgreSQL para auth/comunidades/canais/mensagens/moderacao/auditoria/invites
- Refresh token persistente
- Telemetria opt-in persistida
- Validacao end-to-end apos restart

## Checklist
- [x] Refatorar backend para DB-first
- [x] Ajustar compose/env de banco
- [x] Validar build e fluxo API completo
- [x] Validar persistencia apos restart
- [x] Atualizar docs tecnicas

## Resultado
Fluxo principal persiste apos restart. API agora usa PostgreSQL como fonte de verdade para entidades do MVP.

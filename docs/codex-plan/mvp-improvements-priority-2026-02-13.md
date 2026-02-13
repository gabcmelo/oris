# MVP Improvements Priority - 2026-02-13

## Goal
Endurecer o MVP atual para teste real com publico, mantendo velocidade sem perder seguranca basica.

## P0 (24-48h)
1. Voice token hardening:
   - `POST /voice/token` deve ignorar `userId` do payload e usar somente usuario autenticado.
   - validar se `channelId` existe, eh do tipo `voice` e se usuario pertence a comunidade.
2. CORS para ambiente web default:
   - permitir origin configuravel (`APP_ALLOWED_ORIGINS`) e responder `OPTIONS`.
3. WebSocket origin guard:
   - validar `Origin` no upgrade do WS.
4. Invite accounting:
   - incrementar `uses_count` apenas quando novo membro entra de fato.
5. RBAC admin para telemetry:
   - limitar `/admin/telemetry/*` a papel `owner/admin`.

## P1 (3-5 dias)
1. Refresh token e sessao:
   - armazenar refresh token no frontend e renovar access token automaticamente.
   - logout deve revogar refresh token ativo.
2. Message/channel consistency:
   - bloquear envio/listagem de mensagens em canal `voice` (ou separar chat lateral explicitamente).
3. Upgrade script confiavel:
   - registrar versao/tag anterior e executar rollback real para imagem anterior.
4. Docs de operacao:
   - corrigir README para remover afirmacao de estado in-memory.

## P2 (1-2 semanas)
1. Banco:
   - adicionar foreign keys e indices para consultas de uso comum.
2. Seguranca:
   - rate-limit em login e endpoints sensiveis.
3. Qualidade:
   - testes automatizados de regressao para auth/invite/moderacao/voice/ws.

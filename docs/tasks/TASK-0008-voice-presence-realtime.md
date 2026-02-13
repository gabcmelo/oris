# TASK-0008-voice-presence-realtime

- Status: done
- Owner: Codex
- Objetivo: Evoluir para experiencia de comunidade mais completa com presenca em tempo real e controles de voz no frontend.

## Entregaveis
- Presenca por canal via WebSocket + endpoint de presenca
- Controles de voz no frontend (join/leave/mute)
- Ajuste de URL publica LiveKit para browser

## Checklist
- [x] Backend: eventos de presenca
- [x] Backend: endpoint de presenca
- [x] Frontend: integrar livekit-client
- [x] Frontend: controles de voz
- [x] Validacao build e fluxo

## Riscos
- Reproducao de audio remoto depende de permissao/dispositivo local do browser.

## Resultado
- Endpoint `GET /api/v1/channels/:channelId/presence` entregue.
- WebSocket atualizado para emitir `presence.updated` em connect/disconnect.
- Frontend com fluxo de voz LiveKit (join/leave/mute) e leitura de `serverUrl` publico.
- Validacoes executadas em container: build `api`/`web`, smoke de auth/comunidade/canais/moderacao/voz.

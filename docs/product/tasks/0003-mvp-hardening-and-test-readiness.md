# TASK-0003-mvp-hardening-and-test-readiness

- Status: done
- Owner: Codex
- Objetivo: Revisar documentos e complementar o MVP para ficar pronto para testes imediatos.

## Entregaveis
- Harden de auth (hash de senha)
- Token de voz LiveKit assinado
- Chat realtime via WebSocket
- Guia de teste atualizado

## Checklist
- [x] Revisar docs de plano e estado atual
- [x] Implementar melhorias no backend
- [x] Atualizar compose/env para LiveKit secret
- [x] Atualizar frontend para escuta realtime
- [x] Validar subida e fluxo basico
- [x] Documentar como testar

## Riscos
- Runtime ainda usa armazenamento em memoria no backend para entidades de negocio.

## Resultado
MVP ficou mais testavel hoje: auth com bcrypt, token de voz JWT assinado, websocket realtime por canal e stack validada em execucao local.

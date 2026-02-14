# Community and Voice MVP Feature

## Description
Entregar uma experiência de comunidade usável com canais de texto, fluxo de entrada em canais de voz, convites e moderação básica. Esta é uma task de alto nível que engloba múltiplas funcionalidades necessárias para o MVP.

## Objective
Permitir que usuários criem comunidades, canaisde texto e voz, troquem mensagens em tempo real, conectem-se a salas de voz via LiveKit, e tenham ferramentas básicas de moderação.

## Scope
### Includes
- Sistema completo de autenticação (register/login/refresh/logout)
- Criação e listagem de comunidades e canais
- Postagem e listagem de mensagens em canais de texto
- Emissão de token de voz e integração com LiveKit UI
- Moderação básica (kick, mute, ban) com auditoria

### Excludes
- Federação entre instâncias
- Empacotamento completo do cliente desktop
- Moderação avançada com ML/trust & safety
- Integrações com serviços externos
- Sistema de permissões granular avançado

## Technical Plan
1. Endpoints de autenticação (`/api/v1/auth/*`)
2. Endpoints de comunidades (`/api/v1/communities*`)
3. Endpoints de canais (`/api/v1/channels/*`)
4. Endpoint de token de voz (`/api/v1/voice/token`)
5. WebSocket para mensagens em tempo real (`/api/v1/ws/:channelId`)
6. Sistema de convites com validação
7. Sistema de moderação com log de auditoria
8. Interface web para todas as funcionalidades acima

## Acceptance Criteria
- [ ] Usuário pode registrar-se e fazer login via UI web
- [ ] Usuário pode criar comunidade e canais (texto e voz)
- [ ] Dois usuários podem trocar mensagens em tempo real
- [ ] Token de voz é emitido e conexão com sala LiveKit é estabelecida
- [ ] Ações de moderação são persistidas no log de auditoria
- [ ] Convites inválidos/expirados são rejeitados corretamente
- [ ] Membros mutados/banidos não podem postar mensagens
- [ ] Token de voz valida autorização para o canal

## Evidence
Esta é uma task épica/feature de alto nível. Evidências serão documentadas nas subtasks específicas.

## Risks
- Escopo muito amplo para uma única task - recomenda-se quebrar em subtasks menores
- Dependência de LiveKit para funcionalidade de voz
- Complexidade de sincronização em tempo real via WebSocket

## Dependencies
- Task 0004 (Close MVP gaps for testing)
- Task 0008 (Voice presence realtime)
- Infraestrutura LiveKit configurada
- Frontend React com Vite
- Backend Go com PostgreSQL e Redis

## Status
BACKLOG

## Created At
2026-02-12

## Updated At
2026-02-12

## Owner
TBD

## Notes
Esta task representa o MVP completo. Devido ao princípio de WIP Limit = 1 e necessidade de revisão incremental, recomenda-se quebrar em tasks menores e específicas:
- Auth flow
- Community/channel CRUD
- Text messaging + WebSocket
- Voice token + LiveKit integration
- Moderation system

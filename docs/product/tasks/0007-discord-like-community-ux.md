# TASK-0007-discord-like-community-ux

- Status: done
- Owner: Codex
- Objetivo: Avancar para experiencia mais proxima de um app de comunidade tipo Discord: membros, convites e moderacao contextual.

## Entregaveis
- Endpoint de membros por comunidade
- Regras de moderacao mais seguras (sem auto-acao e sem moderar cargo igual/superior)
- UI com painel de membros, acoes de moderacao e fluxo de convite/join

## Checklist
- [x] Adicionar `GET /communities/:communityId/members`
- [x] Endurecer regras de moderacao no backend
- [x] Adicionar bloco de convites no frontend
- [x] Adicionar painel de membros com acoes
- [x] Validar fluxo fim-a-fim

## Resultado
Aplicacao ficou mais utilizavel para testes de comunidade real: criacao de convite, entrada por codigo, visualizacao de membros e moderacao diretamente na UI.

# 🎯 Sistema de Tasks Centralizado - Guia Rápido

**Data:** 2026-02-12  
**Status:** ✅ Concluído

## 📍 Localização Central
**Tudo está agora em:** `docs/product/tasks/`

## 🚀 Como usar (5 passos)

### 1. Antes de começar qualquer trabalho
```powershell
# Abrir o taskboard
code docs/product/tasks/taskboard.csv
```

### 2. Verificar WIP Limit = 1
- Procurar por `status=DOING` no CSV
- **Deve haver APENAS UMA task DOING**
- Se não houver, escolher uma task BACKLOG e mudar para DOING

### 3. Abrir a task file
```powershell
# Exemplo: se a task DOING é 0011
code docs/product/tasks/0011-backend-restructure-modularization.md
```

### 4. Trabalhar na task
- Ler toda a task (Description, Objective, Scope, Technical Plan)
- Seguir o Technical Plan passo a passo
- Marcar Acceptance Criteria conforme completa
- **NÃO expandir o scope** - criar nova task se descobrir trabalho adicional

### 5. Finalizar a task
Antes de marcar como DONE:
- [ ] Todo o código implementado
- [ ] Testes passando (quando aplicável)
- [ ] Evidência documentada na seção Evidence
- [ ] `taskboard.csv` atualizado: `status=DONE`, `updated_at=YYYY-MM-DD`
- [ ] Docs atualizadas (se comportamento mudou)

## 📝 Criar nova task

### Passo 1: Gerar ID
```powershell
# Ver último ID usado
Import-Csv docs/product/tasks/taskboard.csv | Select-Object -Last 1 | Select-Object id
# Próximo ID = último ID + 1
```

### Passo 2: Criar arquivo
```powershell
# Exemplo: criar task 0027
Copy-Item docs/product/tasks/0000-template.md docs/product/tasks/0027-minha-task.md
code docs/product/tasks/0027-minha-task.md
```

### Passo 3: Preencher task
- Description: Por que essa task existe?
- Objective: O que será entregue?
- Scope > Includes: O que está no escopo
- Scope > Excludes: O que NÃO está no escopo (importante!)
- Technical Plan: Passos executáveis
- Acceptance Criteria: Lista de checagem (checkboxes)
- Status: BACKLOG
- Owner: TBD ou seu nome
- Created At: Data atual

### Passo 4: Adicionar ao taskboard
Adicionar linha no `taskboard.csv`:
```csv
0027,Nome curto da task,backend,BACKLOG,P1,TBD,2026-02-13,2026-02-13,,,Notas opcionais
```

**Campos:**
- `id`: Número sequencial (0027)
- `title`: Título curto e descritivo
- `area`: backend|frontend|infra|docs|product
- `status`: BACKLOG (para novas tasks)
- `priority`: P0 (crítico) | P1 (importante) | P2 (nice-to-have)
- `owner`: TBD ou nome
- `created_at`: YYYY-MM-DD
- `updated_at`: YYYY-MM-DD
- `depends_on`: IDs de tasks necessárias (ex: "0011,0015")
- `pr`: Número do PR quando criado
- `notes`: Contexto adicional

## 🔍 Status das tasks

```powershell
# Ver tasks DONE
Import-Csv docs/product/tasks/taskboard.csv | Where-Object { $_.status -eq "DONE" } | Select-Object id,title

# Ver tasks BACKLOG
Import-Csv docs/product/tasks/taskboard.csv | Where-Object { $_.status -eq "BACKLOG" } | Select-Object id,title,priority | Sort-Object priority

# Ver task atual (DOING)
Import-Csv docs/product/tasks/taskboard.csv | Where-Object { $_.status -eq "DOING" } | Select-Object id,title,owner
```

## ⚠️ Regras Obrigatórias

### Golden Rule: WIP Limit = 1
**Sempre apenas UMA task DOING por vez!**

Isso garante:
- ✅ Foco e qualidade
- ✅ Review completo antes de prosseguir
- ✅ Nenhuma feature pela metade
- ✅ Progresso visível e incremental

### Durante o trabalho
- ❌ NÃO expandir scope da task atual
- ❌ NÃO começar outra task em paralelo
- ❌ NÃO implementar "só mais essa feature rápida"
- ✅ Descobriu trabalho adicional? Criar nova task BACKLOG
- ✅ Task ficou muito grande? Quebrar em tasks menores
- ✅ Bloquado? Mudar status para BLOCKED com nota explicando

### Antes de marcar DONE
- ✅ Evidência documentada (comandos + outputs)
- ✅ `taskboard.csv` atualizado
- ✅ Testes passando (ou justificativa)
- ✅ Docs atualizadas

## 📂 Estrutura de arquivos

```
docs/product/tasks/
├── taskboard.csv              # 📊 SINGLE SOURCE OF TRUTH
├── README.md                  # 📖 Guia completo
├── 0000-template.md           # 📋 Template oficial
├── 0001-...-0026-*.md         # 📄 Task files
└── logs/                      # 📁 Evidence folder
    ├── README.md
    ├── 2026-02-12-task-migration.md
    └── ...
```

## 🔗 Links importantes

- **Taskboard**: [docs/product/tasks/taskboard.csv](docs/product/tasks/taskboard.csv)
- **Template**: [docs/product/tasks/0000-template.md](docs/product/tasks/0000-template.md)
- **README completo**: [docs/product/tasks/README.md](docs/product/tasks/README.md)
- **Protocolo AGENTS**: [AGENTS.md](../../AGENTS.md)
- **Product Vision**: [docs/product/vision.md](docs/product/vision.md)
- **Roadmap**: [docs/product/roadmap.md](docs/product/roadmap.md)

## 💡 Dicas

### Task muito grande?
Se uma task tem mais de 5-7 itens no Acceptance Criteria, provavelmente é grande demais. Quebre em tasks menores.

**Exemplo:**
❌ Task única: "Implementar sistema de autenticação completo"  
✅ Tasks granulares:
- 0027: Register endpoint + validação
- 0028: Login endpoint + JWT
- 0029: Refresh token flow
- 0030: Logout + revogação de token

### Descobriu bug durante o trabalho?
Não expanda o scope. Crie task separada:
```csv
0028,Fix: Validation error on empty email,backend,BACKLOG,P0,TBD,2026-02-13,2026-02-13,,,Found during task 0027
```

### Task bloqueada?
Mude status para BLOCKED e documente:
```csv
0029,Voice integration,backend,BLOCKED,P0,Dev,2026-02-13,2026-02-13,,,Waiting for LiveKit credentials
```

### Evidência visual?
Coloque screenshots em `logs/`:
```
docs/product/tasks/logs/0027-register-form.png
docs/product/tasks/logs/0027-validation-errors.png
```

E referencie na task:
```markdown
## Evidence
- ✅ `go test ./...` passed
- ✅ Integration test: `logs/0027-register-form.png`
- ✅ Validation behavior: `logs/0027-validation-errors.png`
```

## 🎓 Exemplo completo

### Ver task atual
```powershell
cat docs/product/tasks/taskboard.csv | Select-String "DOING"
```
Output: `0011,Backend restructure and modularization,backend,DOING,...`

### Abrir task
```powershell
code docs/product/tasks/0011-backend-restructure-modularization.md
```

### Trabalhar
- Seguir Technical Plan
- Marcar Acceptance Criteria

### Validar
```powershell
cd backend
go test ./...
go build ./cmd/api
```

### Documentar evidência
No arquivo `0011-backend-restructure-modularization.md`:
```markdown
## Evidence
- ✅ `go test ./...` - all tests passed
- ✅ `go build ./cmd/api` - build successful
- ✅ Docker compose up - services started
- ✅ Smoke test: auth endpoints working
```

### Finalizar
Atualizar `taskboard.csv`:
```csv
0011,Backend restructure and modularization,backend,DONE,P0,Codex,2026-01-25,2026-02-13,,,Refatoração concluída com sucesso
```

---

**Dúvidas?** Leia [docs/product/tasks/README.md](docs/product/tasks/README.md) ou [AGENTS.md](../../AGENTS.md)

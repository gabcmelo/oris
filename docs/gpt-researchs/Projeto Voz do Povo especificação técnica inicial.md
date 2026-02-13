# Projeto Voz do Povo: especificação técnica e apresentação visual para um sistema open-source de voz seguro e auto-hospedável

## Resumo executivo

“Voz do Povo” é um projeto open-source para comunicação por voz (e, opcionalmente, texto) que **qualquer pessoa ou empresa consegue auto-hospedar** em uma VPS (Linux) ou localmente via contêineres — com a mesma “nostalgia do servidor caseiro” (tipo “passa o endereço e entra”), porém com padrões modernos de segurança, privacidade e observabilidade. A proposta central é **reduzir dependência de plataformas centralizadas e fechadas** (como entity["company","Discord","chat and voice platform"]) e tornar as decisões de segurança auditáveis pela comunidade (código aberto + prática de engenharia sólida). citeturn1search0turn1search1

O caminho técnico mais realista (e com melhor custo/benefício para multiplataforma) é um modelo **web-first** com **PWA** instalável, somado a “wrappers”/apps para desktop e mobile quando necessário. PWAs exigem HTTPS para serem instaláveis e se baseiam em Service Workers para capacidades de cache/offline e comportamento em segundo plano. citeturn7search0turn7search3turn7search4

Para voz em tempo real, o núcleo de mídia deve ser **WebRTC**, que se apoia em um conjunto de protocolos:  
- **ICE** para atravessar NAT/firewalls, usando **STUN** e (quando necessário) **TURN**. citeturn14search2turn1search2turn1search3  
- **DTLS-SRTP** para estabelecer chaves e proteger mídia (SRTP) na rota de mídia. citeturn0search1turn0search2  
- Um **servidor de mídia (SFU/gateway)** para conferência multiponto (evita o custo explosivo de mesh puro). citeturn2search1turn2search14  

No plano de dados e estado, a combinação **PostgreSQL + Redis** é pragmática: PostgreSQL para persistência (usuários, servidores, canais, mensagens, auditoria) e Redis para presença/estado efêmero e fanout rápido (Pub/Sub/Streams). citeturn3search7turn3search3turn4search0turn4search3

O deploy deve nascer “na veia” com **Docker + Docker Compose** (uma única pasta, um `.env`, `docker compose up -d`) e ter caminho claro para **Kubernetes** quando o operador precisar de alta disponibilidade/escala. citeturn3search4turn3search0turn3search10turn3search6

A segurança real (especialmente para menores) não vem só do “ser open-source”, e sim de: criptografia correta, autenticação forte (OAuth2/OIDC + PKCE), isolamento, rate limits, telemetria e um kit de moderação/controles parentais por padrão. A **LGPD** exige “melhor interesse” para crianças e consentimento específico de responsável (com obrigações de transparência, minimização e linguagem acessível), então isso tem de aparecer como requisito de produto, não como “detalhe jurídico”. citeturn20view0

## Visão do produto e página de apresentação estilo entity["company","Canva","design platform"]

A página de apresentação (landing page) precisa funcionar como uma “apresentação vertical” (cada seção = um “slide”), com copy curto e acionável, e com CTAs explícitos: **Baixar**, **Rodar no Docker**, **Documentação**, **Comunidade**. A estética “Canva-like” aqui é menos sobre efeito visual e mais sobre **clareza narrativa + hierarquia + ícones**.

### Estrutura recomendada da homepage

**Hero (primeira dobra / 1 slide)**
- Título: “Comunicação por voz open-source. Você hospeda. Você controla.”
- Subtítulo (1 frase): “Servidor de voz seguro com WebRTC + contêineres — pronto para VPS, empresas, escolas e comunidades.”
- Botões (primários):
  - “Instalar via Docker Compose”
  - “Baixar app (Windows/macOS/Linux)”
- Botões (secundários):
  - “Docs”
  - “Ver código”
- Microcopy de confiança: “Sem lock-in. Sem telemetria forçada. Padrões abertos.”

**Slide “Como funciona em 30 segundos”**
- Bloco 1: “1) Suba seu servidor (1 comando)”
- Bloco 2: “2) Crie uma comunidade e convide”
- Bloco 3: “3) Entre por navegador/app e fale”
- Visual: mini-diagrama (cliente → servidor → voz) com termos simples: “Login”, “Sala”, “Voz”.

**Slide “Por que existe?” (posicionamento)**
- Transparência: “O código é auditável”
- Soberania: “Seus dados ficam com você”
- Segurança: “Criptografia e controles por padrão”
- Operação: “Observabilidade e logs pensados para produção”

**Slide “Recursos” (cards)**
- Voz com canais, push-to-talk, supressão de ruído (futuro)
- Convites e roles
- Painel admin (ban/mute/auditoria)
- Modo “Comunidade com menores” (defaults restritivos)

**Slide “Deploy”**
- “Docker Compose (recomendado)”
- “Kubernetes (para escala)”
- “Instaladores desktop”
- “Android/iOS (store/sideload)”

**Slide “Open-source de verdade”**
- Licença, governança (CONTRIBUTING, CODE_OF_CONDUCT, SECURITY)
- Roadmap público
- Política de vulnerabilidades

**Footer**
- Docs, comunidade, issues, segurança
- “Este projeto pode ser auto-hospedado; a responsabilidade de operação é do administrador do servidor.”

### Mockup de layout simples (wireframe textual)

```
[ NAV ]  Voz do Povo | Docs | Download | Comunidade | GitHub

[ HERO ]
H1: Comunicação por voz open‑source. Você hospeda. Você controla.
P: Suba seu próprio servidor de voz seguro com WebRTC + contêineres.
[ Botão ] Instalar via Docker Compose
[ Botão ] Baixar para Desktop
Links: Ver Docs • Ver Código

[ SLIDE: 30s ]
(1) docker compose up -d  ->  (2) Criar comunidade  ->  (3) Entrar e falar
Mini-diagrama + prints/ícones

[ SLIDE: Segurança ]
TLS + WebRTC seguro + Controles de moderação + Modo para menores
Checklist de “safe defaults”

[ SLIDE: Deploy ]
Docker Compose | Kubernetes | Instaladores | Mobile
Bloco “Requisitos mínimos” + botão “Guia de instalação”

[ SLIDE: Comunidade ]
Chamada para contribuir + link para issues + template de PR

[ FOOTER ]
Licença • Segurança • LGPD/GDPR • Status page (opcional)
```

image_group{"layout":"carousel","aspect_ratio":"16:9","query":["open source saas landing page hero section design","docker compose architecture diagram web app","webrtc architecture diagram sfu","admin dashboard moderation interface ui"],"num_per_query":1}

A escolha por PWA + “slides” de conteúdo facilita entregar “está em todo lugar” sem multiplicar custo de engenharia desde o dia 1. A instalabilidade e critérios típicos de PWA são documentados (HTTPS, manifest, critérios de engajamento em alguns navegadores). citeturn7search0turn7search1turn7search2

## Objetivos e público-alvo

### Objetivos de produto

O projeto deve perseguir objetivos mensuráveis e verificáveis:

- **Auto-hospedagem real**: qualquer pessoa roda em VPS ou máquina local sem “magia oculta” (um comando e um painel). Isso se apoia diretamente no modelo de composição de serviços do Docker Compose. citeturn3search0turn3search12  
- **Segurança auditável**: criptografia correta por padrão e trilha de auditoria (quem fez o quê), alinhada a recomendações do modelo de ameaça e arquitetura de segurança no ecossistema WebRTC. citeturn1search0turn1search1  
- **Compatibilidade**: navegador moderno (WebRTC) como “cliente universal”, com apps nativos/wrappers quando o caso exigir (push/background/ux). A API WebRTC é padronizada pelo entity["organization","W3C","web standards body"] e a suíte de protocolos é coordenada com o entity["organization","IETF","internet standards body"]. citeturn0search12turn0search8  
- **Escalabilidade pragmática**: suportar múltiplas salas (voice channels) por servidor via SFU/gateway e escalar horizontalmente quando necessário, com Redis para estado efêmero e PostgreSQL para consistência. citeturn2search14turn4search0turn3search7  

### Público-alvo prioritário

O projeto tem dois públicos principais, com necessidades diferentes:

**Comunidades (jovens, escolas, grupos de estudo, guildas, famílias)**  
Segurança, simplicidade e “modo protegido” são a prioridade: defaults conservadores, convites controlados, moderação visível, linguagem acessível. Na LGPD, o tratamento de dados de crianças/adolescentes deve observar “melhor interesse” e, para crianças, consentimento específico do responsável. citeturn20view0

**Empresas (comunicação interna, comunidades de clientes, times distribuídos)**  
Priorizam: SSO (OIDC), auditoria, logs e conformidade (trilha de acesso), deploy em infraestrutura própria, e observabilidade. OAuth2/OIDC são base padrão de mercado para autorização/autenticação modernas. citeturn5search0turn5search3

## Requisitos funcionais e não funcionais

### Requisitos funcionais

Para “parecer Discord” sem cair na armadilha de tentar **copiar tudo**, a lista deve ser curta no MVP e extensível por módulos:

1) **Identidade e acesso**
- Cadastro/login local + integração OAuth2/OIDC (opcional desde o MVP, mas arquitetura já preparada). citeturn5search0turn5search3turn16search0  
- Sessões com refresh tokens; revogação por dispositivo.

2) **Comunidades (“servidores”) e canais**
- Comunidades (tenant) com regras próprias.
- Canais de texto e de voz.
- Convites com expiração e limite de uso.

3) **Voz**
- Entrada/saída de canal de voz.
- Push-to-talk e mute/deafen.
- Lista de participantes e estado (falando / silenciado) em tempo real.

4) **Admin e moderação**
- Painel admin (ban/mute/kick), logs de auditoria.
- Ferramenta de denúncias (“report”) com workflow mínimo.

### Requisitos não funcionais

**Segurança e privacidade**
- API sempre em TLS (idealmente TLS 1.3). citeturn6search0  
- WebRTC com DTLS-SRTP e SRTP (criptografia de mídia e integridade). citeturn0search1turn0search2turn1search1  
- Modelo de ameaça explícito (WebRTC threat model) e “secure defaults”. citeturn1search0turn1search1  

**Escalabilidade**
- Não usar “mesh puro” por padrão em salas >3-4 pessoas; adotar SFU/gateway para multiponto.
- Capacidade de separar “sinalização” (back-end) de “mídia”.

**Compatibilidade**
- Web client com WebRTC e fallback de conectividade via ICE (STUN/TURN). citeturn14search2turn1search2turn1search3  
- Aplicativo instalável e “quase nativo” via PWA (HTTPS + manifest + service worker). citeturn7search0turn7search3  

**Operabilidade**
- Métricas, logs e tracing desde cedo (não como “depois a gente vê”). OpenTelemetry define um framework vendor-neutral para coletar/exportar sinais (traces/metrics/logs). citeturn8search5turn8search12  

## Arquitetura proposta e decisões técnicas

A arquitetura sugerida separa de forma rígida o que é **controle** (auth, permissões, mensagens, sinalização) do que é **mídia** (WebRTC/voz). Essa separação é o que viabiliza segurança e escala sem virar um “monólito impossível”.

### Fundamentos de WebRTC para voz

- **ICE** resolve conectividade em redes reais (NAT/firewall), e opera em conjunto com STUN e TURN conforme necessidade. citeturn14search2turn1search2turn1search3  
- **TURN** existe para casos em que a comunicação direta falha; ele define um relay intermediário controlado pelo host, essencial em redes restritivas. citeturn1search3turn1search7  
- **SRTP** provê confidencialidade, autenticação e proteção contra replay para tráfego RTP/RTCP em tempo real. citeturn0search2turn0search14  
- **DTLS-SRTP** define como estabelecer chaves SRTP via DTLS na rota de mídia, independente do canal de sinalização. citeturn0search1turn0search5  

### Diagrama de componentes

```mermaid
graph LR
  subgraph Clients
    Web[PWA/Web Client]
    Desktop[Desktop App]
    Mobile[Mobile App]
  end

  subgraph Edge
    Ingress[Reverse Proxy / Ingress]
    Static[Web Static / CDN-like]
  end

  subgraph ControlPlane["Control Plane (Auth + Signaling + Chat)"]
    API[Backend API (Rust/Go)]
    WS[WebSocket Gateway]
    Auth[OIDC/OAuth2 Provider (optional)]
  end

  subgraph Data
    PG[(PostgreSQL)]
    R[(Redis)]
  end

  subgraph MediaPlane["Media Plane (WebRTC Voice)"]
    SFU[Media Server (SFU/Gateway)]
    TURN[TURN/STUN Server]
  end

  subgraph Obs["Observability"]
    Metrics[Prometheus]
    Dash[Grafana]
    Errors[Sentry]
    OTel[OpenTelemetry Collector]
  end

  Web --> Ingress
  Desktop --> Ingress
  Mobile --> Ingress

  Ingress --> Static
  Ingress --> API
  Ingress --> WS

  API --> PG
  API --> R
  WS --> R

  Web -->|Signaling| WS
  Desktop -->|Signaling| WS
  Mobile -->|Signaling| WS
  WS --> SFU

  Web -->|ICE/STUN/TURN| TURN
  Desktop -->|ICE/STUN/TURN| TURN
  Mobile -->|ICE/STUN/TURN| TURN

  Web -->|DTLS-SRTP Media| SFU
  Desktop -->|DTLS-SRTP Media| SFU
  Mobile -->|DTLS-SRTP Media| SFU

  API --> OTel
  WS --> OTel
  SFU --> OTel
  OTel --> Metrics
  Metrics --> Dash
  API --> Errors
```

A separação “Control Plane vs Media Plane” é coerente com a forma como servidores WebRTC são descritos: o servidor “media” negocia WebRTC, troca mensagens (com JSON/transports) e roteia RTP/RTCP, enquanto a lógica de aplicação vive no entorno. citeturn2search3turn2search14turn2search1

### Sequência de conexão a um canal de voz

```mermaid
sequenceDiagram
  participant C as Client (Web/PWA/App)
  participant API as Backend API
  participant WS as WebSocket Signaling
  participant TURN as STUN/TURN
  participant SFU as Media Server (SFU/Gateway)
  participant DB as PostgreSQL
  participant R as Redis

  C->>API: Login (OIDC/OAuth2) / Session
  API->>DB: Load user + memberships
  API-->>C: Access token (JWT/opaque) + config

  C->>WS: Connect (token)
  WS->>R: Register presence
  WS-->>C: Presence OK

  C->>WS: JoinVoiceChannel(channel_id)
  WS->>API: Authorize join (role/permissions)
  API->>DB: Validate membership + permissions
  API-->>WS: Allowed + ICE servers + TURN creds
  WS-->>C: ICE + TURN + SFU room params

  C->>TURN: ICE gathering / connectivity checks
  C->>SFU: WebRTC Offer (via WS signaling)
  SFU-->>C: WebRTC Answer (via WS signaling)

  Note over C,SFU: DTLS handshake establishes SRTP keys
  Note over C,SFU: Encrypted audio flows via SRTP

  SFU->>R: Publish voice state (speaking/mute)
  R-->>WS: Fanout state update
  WS-->>C: Update participants UI
```

Os elementos ICE/STUN/TURN e DTLS-SRTP refletem diretamente os RFCs associados (ICE e dependências; DTLS-SRTP; SRTP). citeturn14search2turn1search2turn1search3turn0search1turn0search2

### Tabelas comparativas de escolhas centrais

#### Servidor de mídia (SFU/gateway): comparação prática

| Opção | Natureza | Pontos fortes | Atenções | Onde brilha no Voz do Povo |
|---|---|---|---|---|
| entity["organization","mediasoup","webrtc sfu library"] | SFU “baixo nível” e sinalização agnóstica | Flexível, focado em camada de mídia, integra bem com lógica própria | Requer “orquestração”/código de controle no app (comum em Node) | Produto custom, voz + features específicas |
| entity["organization","Janus","webrtc server by meetecho"] | Servidor WebRTC geral; troca mensagens (JSON) e faz relay de RTP/RTCP | Versátil e pluginável; bom para “salas” e gateways | Operação e tuning em Linux; Windows não é alvo do projeto upstream | MVP rápido de voz, arquitetura modular |
| entity["organization","Jitsi Videobridge","webrtc sfu component"] | SFU que roteia streams em conferências | Integrado ao ecossistema Jitsi; grande maturidade em conferência | Stack completa pode ser “grande” para voz-only; mais componentes | Quando você quer herdar features de meeting/ecossistema |

Fontes: descrição e escopo do Janus como servidor WebRTC geral e relay de mídia. citeturn2search3turn2search5 Descrição do mediasoup como SFU, minimalista e “signaling agnostic”. citeturn2search14 Arquitetura do ecossistema Jitsi (JVB como roteador de streams; Jitsi Meet como app WebRTC com React/React Native). citeturn2search1turn2search9

#### Backend: entity["organization","Rust","systems programming language"] vs entity["organization","Go","programming language"]

| Critério | Rust | Go |
|---|---|---|
| Segurança de memória | Sistema de ownership permite garantias sem GC | Usa GC (há docs oficiais sobre custos/ajuste do GC) |
| Produtividade no curto prazo | Curva de aprendizado mais íngreme | Time-to-market rápido para serviços HTTP/concurrency |
| Concurrency/IO | Excelente, porém mais “explícito” | Goroutines e channels são primitivos centrais |

Fontes: Rust “memory safety sem garbage collector” via ownership. citeturn12search2 Go: goroutines/channels e guia oficial sobre garbage collector. citeturn13search3turn13search0turn13search2

Recomendação direta (opinião prática): se o objetivo é **lançar e iterar rápido** (MVP em semanas), Go costuma reduzir atrito; se a prioridade é **alto desempenho com forte disciplina de segurança de memória** e você aguenta o custo inicial, Rust dá vantagens estruturais. As duas servem.

#### Clientes multiplataforma (desktop/mobile): entity["organization","Electron","desktop app framework"] vs entity["organization","Flutter","ui toolkit"] vs entity["organization","React Native","mobile app framework"]

| Critério | Electron | Flutter | React Native |
|---|---|---|---|
| Plataformas principais | Desktop (Windows/macOS/Linux) | Mobile + Web + Desktop | Mobile (Android/iOS) e extensão para “mais” |
| Modelo | Web tech + Chromium/Node embutidos | Toolkit multiplataforma com 1 codebase | Apps nativos com React |
| Melhor uso no projeto | App desktop rápido reutilizando web UI | Quando quer um “único app” e UX consistente | Quando a equipe é React-first no mobile |

Fontes: Electron: Chromium + Node embutidos e foco cross-platform desktop. citeturn12search0turn12search7 Flutter: single codebase para mobile/web/desktop. citeturn12search5turn12search8 React Native: criar apps nativos usando React. citeturn12search9turn12search1

## Modelo de deploy e distribuição

### Objetivo de instalação

O install “ideal” tem que ser **menos complexo do que rodar um servidor de Minecraft antigo**: se para subir o servidor o usuário precisa virar SRE, você perdeu o propósito.

O padrão ouro é:  
- **Server**: `docker compose up -d` (ou instalador one-click que faz isso por baixo). citeturn3search4turn3search8  
- **Client**: Web/PWA como default + desktop/mobile como complementos.

### Diagrama de deploy via Docker Compose

```mermaid
graph TD
  subgraph Host["VPS / Docker Host"]
    RP[reverse-proxy]
    WEB[web-frontend]
    API[api-backend]
    PG[(postgres)]
    REDIS[(redis)]
    SFU[media-server]
    TURN[turn-server]
    PROM[prometheus]
    GRAF[grafana]
  end

  Internet --> RP
  RP --> WEB
  RP --> API
  RP --> SFU

  API --> PG
  API --> REDIS
  SFU --> TURN

  API --> PROM
  SFU --> PROM
  PROM --> GRAF
```

Docker Compose define serviços, redes e volumes de um app multi-contêiner em um único arquivo e é a forma “humana” de orquestrar esse cenário no começo. citeturn3search0turn3search12

### Exemplo de `docker-compose.yml` (MVP)

```yaml
services:
  reverse-proxy:
    image: traefik:v3.1
    command:
      - --providers.docker=true
      - --entrypoints.web.address=:80
      - --entrypoints.websecure.address=:443
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro

  web:
    build: ./apps/web
    environment:
      - PUBLIC_API_BASE_URL=https://voz.local/api
    labels:
      - traefik.http.routers.web.rule=Host(`voz.local`)
      - traefik.http.routers.web.entrypoints=websecure
      - traefik.http.services.web.loadbalancer.server.port=3000

  api:
    build: ./apps/api
    environment:
      - DATABASE_URL=postgres://voz:voz@postgres:5432/voz
      - REDIS_URL=redis://redis:6379/0
      - JWT_ISSUER=voz.local
      - TURN_REST_SECRET=${TURN_REST_SECRET}
    depends_on:
      - postgres
      - redis
    labels:
      - traefik.http.routers.api.rule=Host(`voz.local`) && PathPrefix(`/api`)
      - traefik.http.routers.api.entrypoints=websecure
      - traefik.http.services.api.loadbalancer.server.port=8080

  media:
    image: meetecho/janus-gateway:latest
    # Alternativa: jitsi videobridge / mediasoup stack
    labels:
      - traefik.http.routers.media.rule=Host(`voz.local`) && PathPrefix(`/janus`)
      - traefik.http.routers.media.entrypoints=websecure

  turn:
    image: coturn/coturn:latest
    command: >
      -n --log-file=stdout
      --use-auth-secret
      --static-auth-secret=${TURN_REST_SECRET}
      --realm=voz.local
      --listening-port=3478
      --fingerprint
      --no-multicast-peers
    network_mode: "host"

  postgres:
    image: postgres:16
    environment:
      - POSTGRES_USER=voz
      - POSTGRES_PASSWORD=voz
      - POSTGRES_DB=voz
    volumes:
      - pgdata:/var/lib/postgresql/data

  redis:
    image: redis:7
    volumes:
      - redisdata:/data

volumes:
  pgdata:
  redisdata:
```

Observações: (1) Compose file/spec e sintaxe de serviços/builds são documentadas; (2) redes/portas UDP e faixas de mídia exigem atenção especial na prática (WebRTC é sensível a NAT, firewall e portas). citeturn3search0turn14search2turn1search3

### Exemplos de Dockerfile

Backend (exemplo “Rust API” multi-stage):

```dockerfile
# apps/api/Dockerfile
FROM rust:1.75 as builder
WORKDIR /app
COPY Cargo.toml Cargo.lock ./
COPY src ./src
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/target/release/voz-api /app/voz-api
EXPOSE 8080
ENTRYPOINT ["/app/voz-api"]
```

Frontend (exemplo “React/Next web”):

```dockerfile
# apps/web/Dockerfile
FROM node:20 as builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-slim
WORKDIR /app
COPY --from=builder /app ./
EXPOSE 3000
CMD ["npm","run","start"]
```

Docker documenta o modelo de construção via Dockerfile e o conjunto de instruções suportadas. citeturn3search1turn3search9

### Distribuição para Windows/macOS/Linux/Android/iOS

**Estratégia recomendada (prática, com menor risco):**
- Web/PWA como “cliente universal” (rápido de entregar e atualizar). citeturn7search9turn7search0  
- Desktop: Electron ou Flutter (se quiser unificar mais). citeturn12search0turn12search5  
- Mobile: React Native ou Flutter. citeturn12search1turn12search12  

**Canais de distribuição:**
- Desktop: instaladores (.exe/.msi no Windows; .dmg no macOS; AppImage/.deb no Linux) — idealmente gerados por pipeline CI.
- Android/iOS: loja (quando fizer sentido) + sideload para instâncias privadas.

**Kubernetes (fase v2)**  
Quando a operação exigir HA/escala: Deployments, Services e Ingress são os blocos básicos. Kubernetes é descrito como engine open-source de orquestração para automatizar deploy/scale/management, hospedado pela CNCF. citeturn3search10turn3search6turn3search2

## Dados, autenticação, criptografia e proteção de menores

### Banco de dados e cache: recomendações

- **entity["organization","PostgreSQL","relational database"]**: persistência forte e esquema claro; uso de `jsonb` para campos extensíveis e índices GIN quando fizer sentido. citeturn3search7turn3search3turn4search13turn4search9  
- **entity["organization","Redis","in-memory data store"]**: presença/estado em tempo real e fanout. Pub/Sub para eventos efêmeros e Streams quando você quiser persistência/garantias de entrega. citeturn4search0turn4search3  

### Esquema básico (DDL inicial)

```sql
-- Usuários
create table users (
  id uuid primary key,
  email text unique,
  username text unique not null,
  password_hash text,
  created_at timestamptz not null default now(),
  status text not null default 'active'
);

-- Comunidades (servidores)
create table communities (
  id uuid primary key,
  name text not null,
  owner_user_id uuid not null references users(id),
  created_at timestamptz not null default now()
);

create table community_members (
  community_id uuid not null references communities(id),
  user_id uuid not null references users(id),
  role text not null default 'member',
  joined_at timestamptz not null default now(),
  primary key (community_id, user_id)
);

-- Canais
create table channels (
  id uuid primary key,
  community_id uuid not null references communities(id),
  name text not null,
  type text not null check (type in ('text','voice')),
  created_at timestamptz not null default now()
);

-- Mensagens de texto
create table messages (
  id uuid primary key,
  channel_id uuid not null references channels(id),
  author_user_id uuid not null references users(id),
  content text,
  metadata jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now(),
  deleted_at timestamptz
);

create index messages_channel_time_idx
  on messages(channel_id, created_at desc);

create index messages_metadata_gin
  on messages using gin (metadata);

-- Sessões de voz (metadados)
create table voice_sessions (
  id uuid primary key,
  channel_id uuid not null references channels(id),
  started_at timestamptz not null default now(),
  ended_at timestamptz,
  media_backend text not null, -- "janus" | "jvb" | "mediasoup"
  details jsonb not null default '{}'::jsonb
);

create table voice_participants (
  session_id uuid not null references voice_sessions(id),
  user_id uuid not null references users(id),
  joined_at timestamptz not null default now(),
  left_at timestamptz,
  stats jsonb not null default '{}'::jsonb,
  primary key (session_id, user_id, joined_at)
);

-- Auditoria / moderação
create table audit_log (
  id uuid primary key,
  community_id uuid references communities(id),
  actor_user_id uuid references users(id),
  action_type text not null,
  target jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now()
);
```

PostgreSQL documenta a semântica de `CREATE TABLE`, tipos `json/jsonb` e criação de índices (incluindo GIN e índices em expressões). citeturn3search7turn3search3turn4search13turn4search1

### Autenticação e autorização

Modelo recomendado (flexível e “Enterprise-ready”):
- **OAuth 2.0 / OIDC** (SSO opcional) para login federado e integração com IdPs. citeturn5search0turn5search3  
- **Authorization Code + PKCE** para clientes “públicos” (mobile/desktop/web SPA), reduzindo risco de interceptação do authorization code. citeturn16search0  
- Tokens:
  - Preferir **tokens opacos** internamente (melhor controle/ revogação), ou JWT quando necessário, lembrando que JWT é um formato compacto de claims assinados/ protegidos. citeturn5search2turn5search10  
- Para Bearer tokens, TLS é obrigatório e os tokens devem ser protegidos em trânsito e armazenamento. citeturn5search1turn5search9  

### Criptografia: TLS, DTLS, SRTP e opção de E2EE

- **TLS 1.3** deve proteger API e tráfego de sinalização HTTP. citeturn6search0turn6search8  
- **DTLS-SRTP** estabelece chaves SRTP na rota de mídia; SRTP protege mídia com confidencialidade e integridade. citeturn0search1turn0search2turn0search14  
- **E2EE (ponta a ponta real)** em calls multiponto com SFU é possível, mas vem com trade-offs fortes: o servidor não deve ver mídia em claro, o que limita recursos como gravação server-side, moderação automatizada, mixagem etc. O entity["organization","W3C","web standards body"] possui especificações para acesso a mídia codificada (“Encoded Transform”), visando permitir criptografia local e cenários de “untrusted conferencing”. citeturn6search6turn6search2  
- A própria entity["organization","Mozilla","internet organization"] discute APIs de E2EE em navegadores e as diferenças de forma entre browsers (importante para compatibilidade real). citeturn6search14  

Recomendação objetiva: no MVP, entregue **cripto de transporte correta (DTLS-SRTP + TLS)**; em v2, ofereça E2EE como “modo avançado” com explicação clara do que quebra/limita. citeturn1search0turn1search1turn0search1turn0search2

### Proteção de menores e pontos legais

**LGPD (Brasil)**  
A LGPD determina que o tratamento de dados pessoais de crianças/adolescentes deve observar o “melhor interesse”; para crianças, exige consentimento específico e em destaque por ao menos um responsável, além de obrigações de transparência e esforço razoável de verificação de consentimento. citeturn20view0  

**ECA (Brasil)**  
O Estatuto define “criança” como pessoa até 12 anos incompletos e “adolescente” entre 12 e 18 anos, o que impacta diretamente políticas de age-gating e UX de consentimento. citeturn22search0  

**GDPR (UE)**  
O GDPR é o regulamento europeu de proteção de dados, aplicável quando houver oferta/uso por titulares na UE, com obrigações específicas (base legal, transparência, minimização, direitos do titular etc.). citeturn11search2turn11search6  

**Implicações diretas para o Voz do Povo (requisitos de produto, não só “jurídico”)**
- Modo “Servidor para menores” como template com defaults: desabilitar DMs por padrão, convites com expiração curta, permissões mais rígidas, logs de auditoria ativos, e linguagem acessível. (A exigência de informação clara e acessível está explicitada na LGPD para dados de crianças). citeturn20view0  
- Minimização: coletar o mínimo necessário e evitar “features de risco” (ex.: discovery pública por padrão) até haver maturidade de moderação.
- Transparência do administrador: o painel deve expor “o que coletamos” e “por quanto tempo guardamos”.

## Operação, observabilidade, CI/CD e governança open-source

### Logging, monitoring e observability

Sem observabilidade, “auto-hospedável” vira sinônimo de “meu servidor caiu e eu não sei por quê”.

Stack recomendada:

- **entity["organization","Prometheus","monitoring system"]** para métricas: há práticas oficiais de instrumentação e exposição. citeturn8search0turn8search3  
- **entity["organization","Grafana","dashboard and viz"] para dashboards (painéis são coleções de visualizações em uma visão rápida). citeturn8search14turn8search10  
- **entity["organization","OpenTelemetry","observability framework"] para correlação de traces/metrics/logs e export para backends. citeturn8search5turn8search12turn8search8  
- **entity["company","Sentry","error monitoring"] para erro/performance (captura explícita de errors e tracing). citeturn8search1turn8search7  

Mínimo viável de métricas (MVP):
- API: latência p95/p99, erro por rota, taxa de login, rate-limit hits
- Signaling: conexões WS ativas, reconexões, tempo de join de sala
- Media: sessões ativas, participantes, perda de pacotes (quando disponível), uso de TURN vs direto (proxy de custo)

### CI/CD e testes

- Workflows do entity["company","GitHub","code hosting platform"] Actions são definidos em YAML (sintaxe documentada). citeturn9search0turn9search4  
- Matrix strategy permite rodar jobs em combinações de SO/versão (ótimo para builds de instaladores e testes). citeturn9search2turn9search11  
- Para container images: publicar imagens via workflows “build and push” (docs e exemplos oficiais). citeturn9search5turn9search1  

Sugestão de pipeline (alto nível, prática):
- PR: lint + unit tests + SAST básico + build de containers (sem push)
- Main: build + push em registry + SBOM + assinatura de imagem
- Release tag: gerar instaladores desktop e artefatos mobile

Assinatura de contêiner (supply chain): **Sigstore/cosign** permite assinatura “keyless” via OIDC, reduzindo atrito. citeturn9search3turn9search13

### Governança open-source

Aqui a postura precisa ser “adulto na sala”. Se você quer confiança, você precisa de regras.

- Licença: use uma licença aprovada pela entity["organization","Open Source Initiative","open source license org"] (isso define o que é considerado “open source” em termos de revisão e conformidade com a Open Source Definition). citeturn10search0  
- Identificadores de licença: usar SPDX no repositório facilita automação e compliance; a entity["organization","SPDX","software license identifiers"] mantém a lista oficial de identificadores. citeturn10search1turn10search13  
- Segurança: publicar `SECURITY.md` com canal e processo de reporte é suportado/documentado pelo GitHub. citeturn10search2turn10search10  
- Código de conduta: entity["organization","Contributor Covenant","code of conduct project"] é um padrão amplamente adotado para estabelecer expectativas e reduzir abuso na comunidade. citeturn10search3  

**Recomendação de licença (opinião pragmática):**
- Se a meta é máxima adoção (inclusive comercial) e ecossistema de plugins: Apache-2.0 costuma ser o caminho “frictionless” (e a própria ASF publica a licença e SPDX ID). citeturn10search12  
- Se a meta é forçar que melhorias de quem “roda como serviço” voltem para a comunidade: considerar copyleft forte (ex.: AGPL). (Aqui a decisão é estratégia, não técnica.)

## Roadmap inicial, estimativas e checklist de lançamento

### Roadmap recomendado

**MVP (6–10 semanas, se foco for voz + base de comunidades)**
- Web/PWA: login, lista de comunidades, canais, entrar/sair de voz
- Backend: auth básico, permissões simples, signaling, presença
- Media: 1 SFU/gateway funcional + coturn opcional
- Admin mínimo: ban/mute + audit log
- Deploy: docker compose “production-ready” (env, volumes, reverse proxy)

**v1 (3–5 meses)**
- Apps desktop (Electron/Flutter) e mobile (RN/Flutter) com onboarding melhor
- SSO via OIDC + PKCE
- Moderação expandida (reporting, rate limits, proteção anti-spam)
- Observability completa + dashboards padrão
- Documentação “self-hosting” e hardening

**v2 (6–12 meses)**
- Kubernetes charts/manifests
- E2EE opcional (Encoded Transform) e matriz de compatibilidade por browser
- Federação entre instâncias (se “distribuído” virar requisito real)
- Extensibilidade por plugins

### Estimativa de recursos e custos (modelo, não preço exato)

Custos operacionais variam por provedor e região, mas os maiores drivers são previsíveis:

- **Banda** (especialmente quando TURN é usado como relay). TURN existe justamente para viabilizar conectividade quando direta falha. citeturn1search3turn1search7  
- **CPU/RAM** do SFU/gateway: depende de codec, número de usuários simultâneos e se há vídeo/gravação. Como referência de comportamento em escala, avaliações públicas do Jitsi Videobridge mostram distribuição de centenas/milhares de streams e alto uso de banda, com CPU relativamente moderado em alguns cenários. citeturn15search7  
- **Armazenamento** (PostgreSQL) + retenção de logs/auditoria.
- **Manutenção**: atualizações de segurança, triagem de issues, revisão de PRs, CI rodando.

Equipe mínima (realista) para não virar “projeto que morre”:
- 1 dev backend (Rust/Go)
- 1 dev frontend (web/PWA)
- 0.5–1 devops/infra (part-time) + cultura de observability
- 0.5 segurança/compliance (part-time, mas constante), principalmente por conta de menores e auto-hospedagem multi-tenant

### Checklist de lançamento (v0.1 público)

- Código
  - Testes de unidade essenciais (auth, permissões, signaling)
  - Threat model inicial (o que a plataforma protege e o que não protege) citeturn1search0turn1search1
- Segurança
  - TLS obrigatório (API e sites) citeturn6search0
  - Rate limiting em login/join/convites
  - `SECURITY.md` com disclosure e SLA de resposta citeturn10search2
- Deploy
  - `docker-compose.yml` de produção + `.env.example` citeturn3search0
  - Guia “subir em VPS” + checklist de portas/UDP (ICE/TURN) citeturn14search2turn1search3
- Observabilidade
  - Métricas Prometheus básicas + dashboard Grafana base citeturn8search0turn8search14
  - Captura de erros (Sentry ou opção self-host) citeturn8search7
- Open-source
  - Licença OSI + SPDX no repo citeturn10search0turn10search1
  - `CONTRIBUTING.md` e Code of Conduct (Contributor Covenant) citeturn10search3
- Menores e privacidade
  - Modo “instância para menores” + documentação de melhores práticas
  - Textos de transparência (o que coleta, por que, retenção) alinhados à LGPD (Art. 14) citeturn20view0
create extension if not exists "uuid-ossp";

create table if not exists users (
  id uuid primary key default uuid_generate_v4(),
  email text,
  username text unique not null,
  password_hash text not null,
  created_at timestamptz not null default now(),
  status text not null default 'active'
);

create table if not exists communities (
  id uuid primary key default uuid_generate_v4(),
  name text not null,
  owner_user_id uuid not null,
  safe_mode_enabled boolean not null default true,
  created_at timestamptz not null default now()
);

create table if not exists community_members (
  community_id uuid not null,
  user_id uuid not null,
  role text not null default 'member',
  muted boolean not null default false,
  banned boolean not null default false,
  joined_at timestamptz not null default now(),
  primary key (community_id, user_id)
);

create table if not exists channels (
  id uuid primary key default uuid_generate_v4(),
  community_id uuid not null,
  name text not null,
  type text not null check (type in ('text','voice')),
  created_at timestamptz not null default now()
);

create table if not exists messages (
  id uuid primary key default uuid_generate_v4(),
  channel_id uuid not null,
  author_user_id uuid not null,
  content text not null,
  created_at timestamptz not null default now(),
  deleted_at timestamptz
);

create table if not exists invites (
  code text primary key,
  community_id uuid not null,
  expires_at timestamptz not null,
  max_uses integer not null,
  uses_count integer not null default 0,
  created_at timestamptz not null default now()
);

create table if not exists audit_log (
  id uuid primary key default uuid_generate_v4(),
  community_id uuid not null,
  actor_user_id uuid not null,
  action_type text not null,
  target jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now()
);

create table if not exists telemetry_settings (
  id boolean primary key default true,
  enabled boolean not null default false,
  updated_at timestamptz not null default now()
);

insert into telemetry_settings (id, enabled) values (true, false)
on conflict (id) do nothing;

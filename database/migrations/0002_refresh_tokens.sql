create table if not exists refresh_tokens (
  token text primary key,
  user_id uuid not null,
  expires_at timestamptz not null,
  created_at timestamptz not null default now()
);

create index if not exists refresh_tokens_user_idx on refresh_tokens(user_id);
create index if not exists refresh_tokens_exp_idx on refresh_tokens(expires_at);

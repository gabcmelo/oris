#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

CURRENT_VERSION=${CURRENT_VERSION:-0.1.0}
TARGET_VERSION=${TARGET_VERSION:-0.1.0}

echo "[upgrade] current=$CURRENT_VERSION target=$TARGET_VERSION"

echo "[upgrade] backup postgres"
mkdir -p backups
BACKUP_FILE="backups/pg-$(date +%Y%m%d-%H%M%S).sql"
docker compose -f infra/docker-compose.yml exec -T postgres pg_dump -U safeguild safeguild > "$BACKUP_FILE"

echo "[upgrade] pulling latest images"
docker compose -f infra/docker-compose.yml pull

echo "[upgrade] restarting stack"
docker compose -f infra/docker-compose.yml up -d --build

echo "[upgrade] health check"
if ! curl -fsS http://localhost:8080/healthz >/dev/null; then
  echo "[upgrade] health check failed; rolling back"
  docker compose -f infra/docker-compose.yml down
  docker compose -f infra/docker-compose.yml up -d
  exit 1
fi

echo "[upgrade] done"

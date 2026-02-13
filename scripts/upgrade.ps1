$ErrorActionPreference = "Stop"

$currentVersion = if ($env:CURRENT_VERSION) { $env:CURRENT_VERSION } else { "0.1.0" }
$targetVersion = if ($env:TARGET_VERSION) { $env:TARGET_VERSION } else { "0.1.0" }

Write-Host "[upgrade] current=$currentVersion target=$targetVersion"

New-Item -ItemType Directory -Force -Path backups | Out-Null
$timestamp = Get-Date -Format "yyyyMMdd-HHmmss"
$backupFile = "backups/pg-$timestamp.sql"

Write-Host "[upgrade] backup postgres"
docker compose -f infra/docker-compose.yml exec -T postgres pg_dump -U oris oris | Out-File -FilePath $backupFile -Encoding utf8

Write-Host "[upgrade] pulling latest images"
docker compose -f infra/docker-compose.yml pull

Write-Host "[upgrade] restarting stack"
docker compose -f infra/docker-compose.yml up -d --build

Write-Host "[upgrade] health check"
try {
  $r = Invoke-WebRequest -Uri "http://localhost:8080/healthz" -UseBasicParsing
  if ($r.StatusCode -ne 200) { throw "health failed" }
} catch {
  Write-Host "[upgrade] health check failed; attempting rollback"
  docker compose -f infra/docker-compose.yml down
  docker compose -f infra/docker-compose.yml up -d
  throw
}

Write-Host "[upgrade] done"

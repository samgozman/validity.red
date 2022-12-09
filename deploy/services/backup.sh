#!/bin/sh

## Backup all databases to s3. This script is run by a cron job every day.
## Or you can run it manually by running via github actions.
## Provide the following secrets in github actions:
## - B2_APPLICATION_KEY_ID
## - B2_APPLICATION_KEY
## They are used to authorize the b2 cli. They will be applied first time from github actions.

# Create backup
docker exec validityred-postgres-1 bash -c 'pg_dump -h localhost --port 5432 --dbname $POSTGRES_DB -U $POSTGRES_USER > /backup/validityred.sql'
# Copy backup to BackBlaze
export B2_APPLICATION_KEY_ID="$(</validityred/b2-app-key-id.txt)"
export B2_APPLICATION_KEY="$(</validityred/b2-app-key.txt)"
b2 authorize-account
b2 sync "/backup/" "b2://validityred/postgres/"
# Delete backup from container
rm -f /backup/validityred.sql || true

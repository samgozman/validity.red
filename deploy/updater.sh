#!/bin/sh

# This script is used to fetch the latest version of the containers
# on "web" server and "services" server. And then restart the containers.
# It should be copied to the "web" server by the `web/web-config.yml`

# Copy .env files (created by the github "deploy" action) to the "services" server.
rsync -r ./calendars.env ./documents.env ./users.env root@10.0.1.1:/validityred
# Copy .env files for "db" server
rsync -r ./db.env updater@10.1.1.2:/validityred

# Update containers on the "db" server first
ssh root@10.1.1.2 "
  cd /validityred
  docker compose down || true
  curl -o docker-compose.yml https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/db/docker-compose.yml
  docker compose pull
  docker compose up --build -d
"

# Update containers on "services" server
ssh root@10.0.1.1 "
  cd /validityred
  docker compose down || true
  curl -o docker-compose.yml https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/services/docker-compose.yml
  docker compose pull
  docker compose up --build -d
"

# Then do the same in the "web" server
cd /validityred
docker compose down || true
curl -o docker-compose.yml https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/web/docker-compose.yml
docker compose pull
docker compose up --build -d

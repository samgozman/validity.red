#!/bin/sh

# This script is used to fetch the latest version of the containers
# on "web" server and "services" server. And then restart the containers.
# It should be copied to the "web" server by the `web/web-config.yml`

# Copy .env files (created by the github "deploy" action) to the "services" server.
rsync -r /validityred/calendars.env /validityred/documents.env /validityred/users.env /validityred/db.env root@10.0.1.1:/validityred

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

# Update nginx confiruation
curl -o nginx.conf https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/nginx.conf
cp -rf nginx.conf /etc/nginx/nginx.conf
systemctl reload nginx

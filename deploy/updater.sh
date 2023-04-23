#!/bin/sh

# This script is used to fetch the latest version of the containers
# and then restart of the containers. It should be copied to the server by the `deploy/cloud-config.yml`

# Then do the same in the "web" server
cd /validityred || exit 1
docker compose down || true
curl -o docker-compose.yml https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/docker-compose.yml
docker compose pull
docker compose up --build -d

# Update nginx confiruation
curl -o nginx.conf https://raw.githubusercontent.com/samgozman/validity.red/main/deploy/nginx.conf
cp -rf nginx.conf /etc/nginx/nginx.conf
systemctl reload nginx

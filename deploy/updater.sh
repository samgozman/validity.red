#!/bin/sh

# This script is used to fetch the latest version of the containers
# on "web" server and "services" server. And then restart the containers.
# It should be copied to the "web" server by the `web/web-config.yml`

# Copy .env files (created by the github "deploy" action) to the "services" server.
rsync -r ./calendars.env ./documents.env ./users.env updater@10.0.1.1:./PROJECT_DIRECTORY

# Update containers on "services" server first
ssh updater@10.0.1.1 "
  cd /validityred
  docker-compose down
  docker-compose pull
  docker-compose up --build -d
"

# Then do the same in the "web" server
cd /validityred
docker-compose down
docker-compose pull
docker-compose up --build -d

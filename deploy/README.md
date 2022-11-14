# Deploy instructions

## Current deployment scheme

```mermaid
graph LR
    A((frontend-service <br/> gateway-service <br/> redis DB <br/>))
    A-->B((user-service <br/> document-service <br/> calendar-service <br/>))
    B-->C((users_postgres <br/> documents_postgres <br/>))
```

In the current deployment setup there are only 3 servers used:

1. Public server - frontend-service, gateway-service, Redis DB
2. Private server - user-service, document-service, calendar-service
3. DB server - users_postgres, documents_postgres (database server)

The idea is that only the public server is exposed to the internet. The private server is only accessible from the public server and the Postgres DB server is only accessible from the private server.

## Preferred server configuration

Lowest cost option is to use a single server for all services. This is the easiest to setup and maintain. The main downside is that the server data will be more vulnerable. Besides, only DB server will need backups and persistent storage.

1. Public server: Ubuntu 22.04, 2 VCPU, 2 GB RAM, 20 GB SSD, IPv4, 1+ TB traffic
2. Private server: Ubuntu 22.04, 2 VCPU, 2 GB RAM, 20 GB SSD
3. DB server: Ubuntu 22.04, 2 VCPU, 2 GB RAM, 20 GB SSD, backups, persistent storage
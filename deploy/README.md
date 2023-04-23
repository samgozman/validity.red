# Deploy instructions

In the current deployment setup, there is only one server used.

The database is backed up to the s3 compatible storage - BackBlaze in this case, every day. BackBlaze bucket should be created manually as I see no use in creating it with Terraform with their current provider.

## Preferred server configuration

Public server: Ubuntu 22.04, 2 VCPU, 2 GB RAM, 20 GB SSD, IPv4, 1+ TB traffic

## Create servers with Terraform

1. Create `secrets.auto.tfvars` file in 'deploy' directory with the following content: `hcloud_token = "YOUR_HETZNER_TOKEN"`
2. Create SSH keys for the servers: `id_rsa` and `validityred_github`
3. Run `terraform init` and then `terraform apply` in 'deploy' directory

## Deploy services

Will be handled by CI/CD pipeline in GitHub Actions. But you can do it manually - just don't forget to set
environment variables in the 'deploy' directory (as described in the `.sample` files).

To-do list for github deployment:

- Create ENV variables in github Secrets section (see .sample files in deploy directory): DB_ENV_VARS, GATEWAY_ENV_VARS, SERVICES_CALENDARS_ENV_VARS, SERVICES_DOCUMENTS_ENV_VARS, SERVICES_USERS_ENV_VARS
- Create env `SENTRY_AUTH_TOKEN` for Sentry plugin for SPA
- Create env `B2_APPLICATION_KEY_ID` and `B2_APPLICATION_KEY` for BackBlaze backups for DB
- Run publish.yml workflow to create docker images
- Run deploy_services.yml workflow to deploy services
- Run deploy_spa.yml to build and deploy SPA
- Run backup_db.yml to create DB backup and to pass BackBlaze ENV variables to the server

## Add monitoring configuration

For monitoring, I use NewRelic and Sentry. Sentry integration is already configured in the services,
you just need to provide DSN to get it to work. As for the NewRelic, you need to install the agent on the server by official instructions.

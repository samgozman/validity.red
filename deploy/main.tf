terraform {
  required_providers {
    hcloud = {
      source = "hetznercloud/hcloud"
    }
  }
  required_version = ">= 1.3.4"
}

# Hetzner Cloud Provider documentation:
# https://registry.terraform.io/providers/hetznercloud/hcloud/latest/docs

#! Create .auto.tfvars file with the following content:
# hcloud_token = "<your_hetzner_api_key>"
variable "hcloud_token" {}

variable "os_type" {
  default = "ubuntu-22.04"
}

variable "ip_range_services" {
  default = "10.0.0.0/16"
}

variable "ip_range_db" {
  default = "10.1.0.0/16"
}

# TODO: get DC name from https://docs.hetzner.cloud/#datacenters
variable "datacenter" {
  default = "nbg1-dc3"
}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_ssh_key" "default" {
  name       = "hetzner_key"
  public_key = file("~/.ssh/id_rsa.pub")
}

## Firewall
# TODO: Create firewall

## VMs

resource "hcloud_server" "web" {
  count       = 1
  name        = "web-server-${count.index}"
  image       = var.os_type
  server_type = "cpx11"
  datacenter  = var.datacenter
  ssh_keys    = [hcloud_ssh_key.default.id]
  backups     = false
  public_net {
    ipv4_enabled = true
    ipv4 = hcloud_primary_ip.public.id
  }
  network {
    network_id = hcloud_network.service_network.id
    ip         = "10.0.1.0"
  }

  # TODO: cloud-init config
  # user_data = file("user_data.yml")
}

resource "hcloud_server" "services" {
  count       = 1
  name        = "service-server-${count.index}"
  image       = var.os_type
  server_type = "cpx11"
  datacenter  = var.datacenter
  ssh_keys    = [hcloud_ssh_key.default.id]
  backups     = false
  public_net {
    ipv4_enabled = false
    ipv6_enabled = false
  }
  network {
    network_id = hcloud_network.service_network.id
    ip         = "10.0.1.1"
  }
  network {
    network_id = hcloud_network.db_network.id
    ip         = "10.1.1.1"
  }
}

resource "hcloud_server" "db" {
  count              = 1
  name               = "db-server"
  image              = var.os_type
  server_type        = "cpx11"
  datacenter         = var.datacenter
  ssh_keys           = [hcloud_ssh_key.default.id]
  backups            = true
  # ! This will cause terraform to hung up on 'apply' or 'destroy' action once it's created.
  # ! If you really need to modify the server, you can do it manually in the Hetzner Cloud Console.
  delete_protection  = true
  rebuild_protection = true
  public_net {
    ipv4_enabled = false
    ipv6_enabled = false
  }
  network {
    network_id = hcloud_network.db_network.id
    ip         = "10.1.1.2"
  }
}

## Network

# Create private network
resource "hcloud_network" "service_network" {
  name     = "service_network"
  ip_range = var.ip_range_services
}
resource "hcloud_network" "db_network" {
  name     = "db_network"
  ip_range = var.ip_range_db
}
# Create subnet for private network
resource "hcloud_network_subnet" "service_network_subnet" {
  network_id   = hcloud_network.service_network.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = var.ip_range_services
}
resource "hcloud_network_subnet" "db_network_subnet" {
  network_id   = hcloud_network.db_network.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = var.ip_range_db
}

# Create public static IP address
resource "hcloud_primary_ip" "public" {
  name          = "primary_public_ip"
  datacenter    = var.datacenter
  type          = "ipv4"
  assignee_type = "server"
  auto_delete   = false
}

## Output

output "web_servers_ips" {
  value = {
    for server in hcloud_server.web :
    server.name => server.ipv4_address
  }
}

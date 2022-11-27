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

variable "datacenter" {
  default = "nbg1-dc3"
}

variable "location" {
  default = "nbg1"
}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_ssh_key" "default" {
  name       = "hetzner_key"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "hcloud_ssh_key" "github" {
  name       = "key_for_github"
  public_key = file("~/.ssh/validityred_github.pub")
}

## Firewall
resource "hcloud_firewall" "public_firewall" {
  name = "public_firewall"
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "80"
    source_ips = [
      "0.0.0.0/0",
      "::/0"
    ]
    description = "HTTP"
  }
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "443"
    source_ips = [
      "0.0.0.0/0",
      "::/0"
    ]
    description = "HTTPS"
  }
}

#! This is a temporary firewall.
#! It is used to allow SSH access to the server from the github actions.
#! This firevall should be removed after the server is configured with VPN.
#! SSH should be allowed only from the network IP range.
#! See https://github.com/samgozman/validity.red/issues/106
resource "hcloud_firewall" "ssh_firewall_public" {
  name = "ssh_firewall_public"
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "22"
    source_ips = [
      "0.0.0.0/0",
    ]
    description = "SSH"
  }
}

resource "hcloud_firewall" "ssh_firewall" {
  name = "ssh_firewall"
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "22"
    source_ips = [
      format("%s/32",hcloud_primary_ip.public.ip_address)
    ]
    description = "SSH"
  }
}

resource "hcloud_firewall" "db_firewall" {
  name = "db_firewall"
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "5432"
    source_ips = [
      "10.1.1.1/32",
      format("%s/32",hcloud_primary_ip.public.ip_address)
    ]
    description = "Postgres"
  }
}

## VMs

resource "hcloud_server" "web" {
  name        = "web-server"
  image       = var.os_type
  server_type = "cpx11"
  datacenter  = var.datacenter
  ssh_keys    = [
    hcloud_ssh_key.default.id,
    hcloud_ssh_key.github.id
  ]
  backups     = false
  public_net {
    ipv4_enabled = true
    ipv4 = hcloud_primary_ip.public.id
  }
  network {
    network_id = hcloud_network.service_network.id
    ip         = "10.0.1.0"
  }
  firewall_ids = [
    hcloud_firewall.public_firewall.id,
    hcloud_firewall.ssh_firewall.id,
    hcloud_firewall.ssh_firewall_public.id,
  ]
  user_data = file("web/web-config.yml")
}

resource "hcloud_server" "services" {
  name        = "service-server"
  image       = var.os_type
  server_type = "cpx11"
  datacenter  = var.datacenter
  ssh_keys    = [
    hcloud_ssh_key.default.id,
    hcloud_ssh_key.github.id
  ]
  backups     = false
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
  network {
    network_id = hcloud_network.service_network.id
    ip         = "10.0.1.1"
  }
  firewall_ids = [
    hcloud_firewall.ssh_firewall.id,
    hcloud_firewall.db_firewall.id,
  ]
  user_data = file("services/services-config.yml")
}

## Volumes

resource "hcloud_volume" "db_volume" {
  location  = var.location
  name      = "db_volume"
  size      = 10
  format    = "ext4"
  delete_protection = true
}

resource "hcloud_volume_attachment" "main" {
  volume_id = hcloud_volume.db_volume.id
  server_id = hcloud_server.services.id
  automount = true
}

## Network

# Create private network
resource "hcloud_network" "service_network" {
  name     = "service_network"
  ip_range = var.ip_range_services
}
# Create subnet for private network
resource "hcloud_network_subnet" "service_network_subnet" {
  network_id   = hcloud_network.service_network.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = var.ip_range_services
}

# Create public static IP address
resource "hcloud_primary_ip" "public" {
  name          = "primary_public_ip"
  datacenter    = var.datacenter
  type          = "ipv4"
  assignee_type = "server"
  auto_delete   = false
  delete_protection  = true
}

## Output

output "web_servers_ip" {
  value = hcloud_server.web.ipv4_address
}

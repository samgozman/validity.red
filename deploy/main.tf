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

## VMs

resource "hcloud_server" "validity" {
  name        = "validity-server"
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
  firewall_ids = [
    hcloud_firewall.public_firewall.id,
    hcloud_firewall.ssh_firewall.id,
    hcloud_firewall.ssh_firewall_public.id,
  ]
  user_data = file("cloud-config.yml")
}

## Network

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

output "servers_ip" {
  value = hcloud_server.validity.ipv4_address
}

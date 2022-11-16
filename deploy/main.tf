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

variable "ip_range" {
  default = "10.0.0.0/16"
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

	# cloud-init config
  # user_data = file("user_data.yml")
}

## Network

# Create private network
resource "hcloud_network" "hc_private" {
  name     = "hc_private"
  ip_range = var.ip_range
}
# Create subnet for private network
resource "hcloud_network_subnet" "hc_private_subnet" {
  network_id   = hcloud_network.hc_private.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = var.ip_range
}
# Attach private network to VMs 
#? Need to correct this
resource "hcloud_server_network" "web_network" {
  count     = 1
  server_id = hcloud_server.web[count.index].id
  subnet_id = hcloud_network_subnet.hc_private_subnet.id
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

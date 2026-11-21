terraform {
  required_providers {
    opnsense = {
      source = "mathieuHa/opnsense"
    }
  }
}

variable "name" {
  type = string
}

variable "local" {
  type = object({
    wan_ip     = string
    tunnel_ip  = string
    lan        = string
    wg_private = string
    wg_public  = string
  })
}

variable "peer" {
  type = object({
    wan_ip    = string
    tunnel_ip = string
    lan       = string
    wg_public = string
  })
}

resource "opnsense_firewall_filter" "wireguard_from_peer" {
  description = "WireGuard from ${var.name}'s peer"
  interface   = { interface = ["wan"] }
  filter = {
    action      = "pass"
    direction   = "in"
    protocol    = "UDP"
    source      = { net = var.peer.wan_ip }
    destination = { net = "(self)", port = "51820" }
  }
}

resource "opnsense_wireguard_client" "peer" {
  name           = "to-peer"
  public_key     = var.peer.wg_public
  server_address = var.peer.wan_ip
  server_port    = 51820
  tunnel_address = ["${var.peer.tunnel_ip}/32", var.peer.lan]
  keep_alive     = 25
}

resource "opnsense_wireguard_server" "this" {
  name           = "site-to-site"
  private_key    = var.local.wg_private
  public_key     = var.local.wg_public
  port           = 51820
  tunnel_address = ["${var.local.tunnel_ip}/24"]
  peers          = [opnsense_wireguard_client.peer.id]
}

resource "opnsense_wireguard_settings" "this" {
  enabled    = true
  depends_on = [opnsense_wireguard_server.this]
}

resource "opnsense_firewall_alias" "peer" {
  name        = "wg_peer"
  type        = "network"
  content     = ["${var.peer.tunnel_ip}/32", var.peer.lan]
  description = "Tunnel address and LAN of the other site"
}

resource "opnsense_firewall_filter" "tunnel" {
  description = "Traffic from the other site"
  interface   = { interface = [] }
  filter = {
    action    = "pass"
    direction = "in"
    protocol  = "any"
    source    = { net = opnsense_firewall_alias.peer.name }
  }
}

output "instance" {
  value = opnsense_wireguard_server.this.instance
}

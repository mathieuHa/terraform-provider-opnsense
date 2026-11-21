terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

resource "opnsense_firewall_alias" "admins" {
  name        = "admins"
  type        = "network"
  content     = ["10.0.50.0/24"]
  description = "Admin workstations"
}

resource "opnsense_firewall_filter" "admins_https" {
  description = "Admins to the web GUI"
  interface   = { interface = ["lan"] }
  filter = {
    action      = "pass"
    direction   = "in"
    protocol    = "TCP"
    source      = { net = opnsense_firewall_alias.admins.name }
    destination = { net = "(self)", port = "https" }
    log         = true
  }
}

resource "opnsense_unbound_host_override" "nas" {
  hostname    = "nas"
  domain      = "home.example"
  server      = "10.0.10.20"
  description = "NAS"
}

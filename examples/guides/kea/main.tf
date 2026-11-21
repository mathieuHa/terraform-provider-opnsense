terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

import {
  to = opnsense_kea_dhcpv4_settings.this
  id = "kea_dhcpv4_settings"
}

resource "opnsense_kea_dhcpv4_settings" "this" {
  enabled        = true
  interfaces     = ["lan"]
  valid_lifetime = 86400
}

resource "opnsense_kea_dhcpv4_subnet" "lan" {
  subnet       = "10.0.10.0/24"
  pools        = ["10.0.10.100 - 10.0.10.199"]
  auto_collect = false
  routers      = ["10.0.10.1"]
  dns_servers  = ["10.0.10.1"]
  domain_name  = "home.example"
  description  = "LAN"
  depends_on   = [opnsense_kea_dhcpv4_settings.this]
}

resource "opnsense_kea_dhcpv4_reservation" "printer" {
  subnet_id   = opnsense_kea_dhcpv4_subnet.lan.id
  ip_address  = "10.0.10.42"
  mac_address = "02:00:00:00:00:42"
  hostname    = "printer"
  description = "Office printer"
}

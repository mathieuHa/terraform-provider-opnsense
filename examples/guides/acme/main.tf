terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

variable "cloudflare_token" {
  type      = string
  sensitive = true
}

resource "opnsense_firmware_plugin" "acme" {
  name = "os-acme-client"
}

import {
  to = opnsense_acmeclient_settings.this
  id = "acmeclient_settings"
}

resource "opnsense_acmeclient_settings" "this" {
  enabled      = true
  auto_renewal = true
  depends_on   = [opnsense_firmware_plugin.acme]
}

resource "opnsense_acmeclient_account" "letsencrypt" {
  name       = "letsencrypt"
  email      = "admin@example.com"
  ca         = "letsencrypt"
  depends_on = [opnsense_firmware_plugin.acme]
}

resource "opnsense_acmeclient_validation" "cloudflare" {
  name         = "cloudflare"
  method       = "dns01"
  dns_service  = "dns_cf"
  dns_cf_token = var.cloudflare_token
  depends_on   = [opnsense_firmware_plugin.acme]
}

resource "opnsense_acmeclient_action" "restart_gui" {
  name       = "restart web GUI"
  type       = "configd_restart_gui"
  depends_on = [opnsense_firmware_plugin.acme]
}

resource "opnsense_acmeclient_certificate" "firewall" {
  name              = "fw.example.com"
  alt_names         = ["vpn.example.com"]
  account           = opnsense_acmeclient_account.letsencrypt.id
  validation_method = opnsense_acmeclient_validation.cloudflare.id
  key_length        = "key_ec256"
  restart_actions   = [opnsense_acmeclient_action.restart_gui.id]
}

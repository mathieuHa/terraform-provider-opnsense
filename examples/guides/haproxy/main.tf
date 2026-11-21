terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

resource "opnsense_firmware_plugin" "haproxy" {
  name = "os-haproxy"
}

import {
  to = opnsense_haproxy_settings.this
  id = "haproxy_settings"
}

resource "opnsense_haproxy_settings" "this" {
  enabled    = true
  depends_on = [opnsense_firmware_plugin.haproxy]
}

resource "opnsense_haproxy_server" "app" {
  for_each   = { app1 = "10.0.10.21", app2 = "10.0.10.22" }
  name       = each.key
  address    = each.value
  port       = 8080
  depends_on = [opnsense_firmware_plugin.haproxy]
}

resource "opnsense_haproxy_healthcheck" "http" {
  name       = "http"
  type       = "http"
  depends_on = [opnsense_firmware_plugin.haproxy]
}

resource "opnsense_haproxy_backend" "app" {
  name           = "app"
  mode           = "http"
  algorithm      = "roundrobin"
  linked_servers = [for s in opnsense_haproxy_server.app : s.id]
  health_check   = opnsense_haproxy_healthcheck.http.id
}

resource "opnsense_haproxy_backend" "fallback" {
  name                 = "fallback"
  mode                 = "http"
  health_check_enabled = false
  depends_on           = [opnsense_firmware_plugin.haproxy]
}

resource "opnsense_haproxy_acl" "app_host" {
  name       = "app_host"
  expression = "hdr"
  hdr        = "app.example.com"
  depends_on = [opnsense_firmware_plugin.haproxy]
}

resource "opnsense_haproxy_action" "to_app" {
  name        = "to_app"
  type        = "use_backend"
  linked_acls = [opnsense_haproxy_acl.app_host.id]
  use_backend = opnsense_haproxy_backend.app.id
}

resource "opnsense_haproxy_frontend" "http" {
  name            = "http"
  bind            = ["0.0.0.0:80"]
  mode            = "http"
  default_backend = opnsense_haproxy_backend.fallback.id
  linked_actions  = [opnsense_haproxy_action.to_app.id]
  depends_on      = [opnsense_haproxy_settings.this]
}

terraform {
  required_providers {
    opnsense = {
      source  = "mathieuHa/opnsense"
      version = "~> 0.37"
    }
  }
}

provider "opnsense" {}

variable "metrics_password" {
  type      = string
  sensitive = true
}

resource "opnsense_firmware_plugin" "caddy" {
  name = "os-caddy"
}

import {
  to = opnsense_caddy_settings.this
  id = "caddy_settings"
}

resource "opnsense_caddy_settings" "this" {
  enabled    = true
  tls_email  = "admin@example.com"
  depends_on = [opnsense_firmware_plugin.caddy]
}

resource "opnsense_caddy_reverse_proxy" "grafana" {
  from_domain = "grafana.example.com"
  description = "Grafana"
  depends_on  = [opnsense_caddy_settings.this]
}

resource "opnsense_caddy_handle" "grafana" {
  reverse   = opnsense_caddy_reverse_proxy.grafana.id
  to_domain = ["10.0.10.30"]
  to_port   = 3000
}

resource "opnsense_caddy_reverse_proxy" "metrics" {
  from_domain = "metrics.example.com"
  description = "Prometheus metrics"
  depends_on  = [opnsense_caddy_settings.this]
}

resource "opnsense_caddy_basic_auth" "prometheus" {
  user       = "prometheus"
  password   = var.metrics_password
  depends_on = [opnsense_firmware_plugin.caddy]
}

resource "opnsense_caddy_handle" "metrics" {
  reverse     = opnsense_caddy_reverse_proxy.metrics.id
  handle_path = "/metrics*"
  basicauth   = [opnsense_caddy_basic_auth.prometheus.id]
  to_domain   = ["10.0.10.31"]
  to_port     = 9100
}

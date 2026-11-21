import {
  to = opnsense_caddy_settings.this
  id = "caddy_settings"
}

resource "opnsense_caddy_settings" "this" {
  enabled   = true
  tls_email = "admin@example.com"
}

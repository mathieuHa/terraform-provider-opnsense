import {
  to = opnsense_tailscale_settings.this
  id = "tailscale_settings"
}

resource "opnsense_tailscale_settings" "this" {
  enabled              = true
  accept_subnet_routes = true
}

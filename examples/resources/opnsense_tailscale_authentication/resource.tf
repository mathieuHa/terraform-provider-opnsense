import {
  to = opnsense_tailscale_authentication.this
  id = "tailscale_authentication"
}

resource "opnsense_tailscale_authentication" "this" {
  login_server = "https://controlplane.tailscale.com"
  pre_auth_key = var.tailscale_auth_key
}

import {
  to = opnsense_acmeclient_settings.this
  id = "acmeclient_settings"
}

resource "opnsense_acmeclient_settings" "this" {
  enabled      = true
  auto_renewal = true
}

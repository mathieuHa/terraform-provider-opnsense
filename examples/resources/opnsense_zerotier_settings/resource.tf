import {
  to = opnsense_zerotier_settings.this
  id = "zerotier_settings"
}

resource "opnsense_zerotier_settings" "this" {
  enabled = true
}

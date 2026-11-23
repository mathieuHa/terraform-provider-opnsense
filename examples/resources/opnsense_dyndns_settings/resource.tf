import {
  to = opnsense_dyndns_settings.this
  id = "dyndns_settings"
}

resource "opnsense_dyndns_settings" "this" {
  enabled = true
}

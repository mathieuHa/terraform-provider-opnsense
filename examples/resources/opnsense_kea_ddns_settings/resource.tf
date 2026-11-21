import {
  to = opnsense_kea_ddns_settings.this
  id = "kea_ddns_settings"
}

resource "opnsense_kea_ddns_settings" "this" {
  enabled = true
}

import {
  to = opnsense_kea_dhcpv4_settings.this
  id = "kea_dhcpv4_settings"
}

resource "opnsense_kea_dhcpv4_settings" "this" {
  enabled        = true
  interfaces     = ["lan"]
  valid_lifetime = 4000
}

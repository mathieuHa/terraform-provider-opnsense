import {
  to = opnsense_kea_dhcpv6_settings.this
  id = "kea_dhcpv6_settings"
}

resource "opnsense_kea_dhcpv6_settings" "this" {
  enabled    = true
  interfaces = ["lan"]
}

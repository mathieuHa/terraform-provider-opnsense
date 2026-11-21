import {
  to = opnsense_interfaces_settings.this
  id = "interfaces_settings"
}

resource "opnsense_interfaces_settings" "this" {
  disable_checksum_offloading = true
  vlan_hw_filter              = "2"
}

import {
  to = opnsense_ipsec_settings.this
  id = "ipsec_settings"
}

resource "opnsense_ipsec_settings" "this" {
  enabled              = true
  make_before_break    = true
  passthrough_networks = ["192.168.1.0/24"]
}

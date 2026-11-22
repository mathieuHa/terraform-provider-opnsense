import {
  to = opnsense_firewall_outbound_nat_settings.this
  id = "firewall_outbound_nat_settings"
}

resource "opnsense_firewall_outbound_nat_settings" "this" {
  mode = "hybrid"
}

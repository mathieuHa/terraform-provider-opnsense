import {
  to = opnsense_ids_settings.this
  id = "ids_settings"
}

resource "opnsense_ids_settings" "this" {
  enabled       = true
  mode          = "pcap"
  interfaces    = ["wan"]
  home_networks = ["10.0.0.0/8"]
  syslog_alerts = true
}

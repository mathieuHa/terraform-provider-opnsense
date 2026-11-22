import {
  to = opnsense_netflow_settings.this
  id = "netflow_settings"
}

resource "opnsense_netflow_settings" "this" {
  capture_interfaces = ["lan"]
  capture_version    = "v9"
  capture_targets    = ["10.0.10.50:2055"]
}

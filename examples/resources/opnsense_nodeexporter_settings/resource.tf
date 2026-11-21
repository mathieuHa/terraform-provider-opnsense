import {
  to = opnsense_nodeexporter_settings.this
  id = "nodeexporter_settings"
}

resource "opnsense_nodeexporter_settings" "this" {
  enabled = true
  cpu     = true
  meminfo = true
  netdev  = true
}

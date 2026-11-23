import {
  to = opnsense_bind_settings.this
  id = "bind_settings"
}

resource "opnsense_bind_settings" "this" {
  enabled    = true
  port       = 53530
  forwarders = ["9.9.9.9"]
}

import {
  to = opnsense_haproxy_settings.this
  id = "haproxy_settings"
}

resource "opnsense_haproxy_settings" "this" {
  enabled = true
}

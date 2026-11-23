import {
  to = opnsense_bind_dnsbl_settings.this
  id = "bind_dnsbl_settings"
}

resource "opnsense_bind_dnsbl_settings" "this" {
  enabled         = true
  forcesafegoogle = true
}

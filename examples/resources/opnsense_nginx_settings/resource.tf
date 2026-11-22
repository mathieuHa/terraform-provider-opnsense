import {
  to = opnsense_nginx_settings.this
  id = "nginx_settings"
}

resource "opnsense_nginx_settings" "this" {
  enabled = true
}

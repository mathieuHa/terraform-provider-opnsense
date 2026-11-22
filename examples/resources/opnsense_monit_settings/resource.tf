import {
  to = opnsense_monit_settings.this
  id = "monit_settings"
}

resource "opnsense_monit_settings" "this" {
  enabled    = true
  interval   = 60
  mailserver = ["smtp.example.com"]
  port       = 587
  ssl        = true
}

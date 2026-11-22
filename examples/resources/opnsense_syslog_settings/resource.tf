import {
  to = opnsense_syslog_settings.this
  id = "syslog_settings"
}

resource "opnsense_syslog_settings" "this" {
  loglocal    = true
  maxpreserve = 14
}

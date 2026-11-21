import {
  to = opnsense_telegraf_general.this
  id = "telegraf_general"
}

resource "opnsense_telegraf_general" "this" {
  enabled  = true
  interval = 10
}

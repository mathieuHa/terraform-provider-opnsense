import {
  to = opnsense_telegraf_input.this
  id = "telegraf_input"
}

resource "opnsense_telegraf_input" "this" {
  cpu = true
  pf  = true
}

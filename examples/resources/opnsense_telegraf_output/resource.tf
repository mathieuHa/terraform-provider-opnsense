import {
  to = opnsense_telegraf_output.this
  id = "telegraf_output"
}

resource "opnsense_telegraf_output" "this" {
  influx_v2_enable       = true
  influx_v2_url          = "https://influx.example.com"
  influx_v2_token        = var.influx_token
  influx_v2_organization = "infra"
  influx_v2_bucket       = "opnsense"
}

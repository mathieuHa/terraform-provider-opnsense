import {
  to = opnsense_kea_ctrl_agent_settings.this
  id = "kea_ctrl_agent_settings"
}

resource "opnsense_kea_ctrl_agent_settings" "this" {
  enabled   = true
  http_host = "127.0.0.1"
  http_port = 8000
}

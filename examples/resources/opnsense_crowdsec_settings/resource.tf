import {
  to = opnsense_crowdsec_settings.this
  id = "crowdsec_settings"
}

resource "opnsense_crowdsec_settings" "this" {
  agent_enabled            = true
  lapi_enabled             = true
  firewall_bouncer_enabled = true
}

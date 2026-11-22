import {
  to = opnsense_systemhealth_settings.this
  id = "systemhealth_settings"
}

resource "opnsense_systemhealth_settings" "this" {
  enabled = true
}

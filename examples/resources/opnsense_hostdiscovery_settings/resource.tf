import {
  to = opnsense_hostdiscovery_settings.this
  id = "hostdiscovery_settings"
}

resource "opnsense_hostdiscovery_settings" "this" {
  enabled       = true
  skip_networks = ["10.0.50.0/24"]
}

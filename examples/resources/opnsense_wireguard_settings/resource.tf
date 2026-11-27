import {
  to = opnsense_wireguard_settings.settings
  id = "wireguard_settings"
}

resource "opnsense_wireguard_settings" "settings" {
  enabled = true
}

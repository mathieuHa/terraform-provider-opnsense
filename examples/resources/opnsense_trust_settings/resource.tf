import {
  to = opnsense_trust_settings.settings
  id = "trust_settings"
}

resource "opnsense_trust_settings" "settings" {
  store_intermediate_certs  = false
  install_crls              = false
  fetch_crls                = false
  enable_legacy_sect        = true
  enable_config_constraints = false
}

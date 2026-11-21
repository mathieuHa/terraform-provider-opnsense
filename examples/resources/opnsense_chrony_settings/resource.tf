import {
  to = opnsense_chrony_settings.this
  id = "chrony_settings"
}

resource "opnsense_chrony_settings" "this" {
  enabled          = true
  nts_client       = true
  peers            = ["time.cloudflare.com", "nts.netnod.se"]
  allowed_networks = ["10.0.0.0/8"]
}

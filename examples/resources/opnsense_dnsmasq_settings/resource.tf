import {
  to = opnsense_dnsmasq_settings.this
  id = "dnsmasq_settings"
}

resource "opnsense_dnsmasq_settings" "this" {
  enable      = true
  interface   = ["lan"]
  cache_size  = 10000
  dhcp_fqdn   = true
  dhcp_domain = "home.example"
}

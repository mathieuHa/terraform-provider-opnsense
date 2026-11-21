resource "opnsense_proxy_pac_match" "internal" {
  name       = "internal"
  match_type = "dns_domain_is"
  hostname   = ".corp.example.com"
}

resource "opnsense_proxy_pac_match" "office_hours" {
  name       = "office-hours"
  match_type = "time_range"
  time_from  = 8
  time_to    = 18
}

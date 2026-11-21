resource "opnsense_proxy_pac_proxy" "corp" {
  name = "corp"
  url  = "proxy.example.com:3128"
}

resource "opnsense_proxy_pac_match" "internal" {
  name       = "internal"
  match_type = "dns_domain_is"
  hostname   = ".corp.example.com"
}

resource "opnsense_proxy_pac_rule" "internal_via_corp" {
  description = "Internal sites through the corporate proxy"
  matches     = [opnsense_proxy_pac_match.internal.id]
  proxies     = [opnsense_proxy_pac_proxy.corp.id]
}

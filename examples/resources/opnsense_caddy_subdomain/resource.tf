resource "opnsense_caddy_reverse_proxy" "wildcard" {
  from_domain = "*.example.com"
  description = "wildcard"
}

resource "opnsense_caddy_subdomain" "grafana" {
  reverse     = opnsense_caddy_reverse_proxy.wildcard.id
  from_domain = "grafana.example.com"
}

resource "opnsense_caddy_subdomain" "grafana" {
  reverse     = opnsense_caddy_reverse_proxy.wildcard.id
  from_domain = "grafana.example.com"
}

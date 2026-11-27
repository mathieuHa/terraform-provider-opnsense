resource "opnsense_nginx_sni_hostname_map" "tls" {
  description = "TLS passthrough"
  entries = [
    { hostname = "app.example.com", upstream = opnsense_nginx_upstream.app.id },
  ]
}

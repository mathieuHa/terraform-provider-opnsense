resource "opnsense_nginx_stream_server" "tls_passthrough" {
  listen_address   = ["8443"]
  route_field      = "sni_upstream_map"
  sni_upstream_map = opnsense_nginx_sni_hostname_map.tls.id
}

resource "opnsense_nginx_http_server" "app" {
  servername           = ["app.example.com"]
  listen_https_address = ["443"]
  certificate          = opnsense_trust_cert.app.id
  https_only           = true
  locations            = [opnsense_nginx_location.root.id]
}

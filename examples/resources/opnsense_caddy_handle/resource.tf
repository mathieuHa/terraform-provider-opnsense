resource "opnsense_caddy_handle" "app" {
  reverse   = opnsense_caddy_reverse_proxy.app.id
  to_domain = ["10.0.10.20"]
  to_port   = 8080
}

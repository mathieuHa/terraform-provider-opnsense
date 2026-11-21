resource "opnsense_caddy_reverse_proxy" "app" {
  from_domain = "app.example.com"
  description = "app"
}

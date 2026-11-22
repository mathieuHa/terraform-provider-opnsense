resource "opnsense_nginx_upstream_server" "app1" {
  description = "app1"
  server      = "10.0.10.21"
  port        = 8080
  priority    = 1
}

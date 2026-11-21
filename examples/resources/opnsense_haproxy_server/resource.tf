resource "opnsense_haproxy_server" "web1" {
  name    = "web1"
  address = "10.0.10.20"
  port    = 8080
}

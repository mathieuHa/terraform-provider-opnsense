resource "opnsense_nginx_resolver" "local" {
  description = "local-resolver"
  address     = ["192.0.2.53"]
  valid       = 300
  ipv6_off    = true
}

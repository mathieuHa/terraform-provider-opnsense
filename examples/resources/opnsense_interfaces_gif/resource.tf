resource "opnsense_interfaces_gif" "site_b" {
  local_address         = "wan"
  remote_address        = "203.0.113.20"
  tunnel_local_address  = "10.255.0.1"
  tunnel_remote_address = "10.255.0.2"
  description           = "Tunnel to site B"
}

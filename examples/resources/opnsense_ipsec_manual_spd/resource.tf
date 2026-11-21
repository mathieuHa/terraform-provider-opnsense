resource "opnsense_ipsec_manual_spd" "extra_net" {
  request_id  = 100
  source      = "192.168.10.0/24"
  destination = "192.168.20.0/24"
  description = "Extra network through the route-based tunnel"
}

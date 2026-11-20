resource "opnsense_firewall_npt" "lan" {
  interface       = "wan"
  source_net      = "fd00:10::/64"
  destination_net = "2001:db8:10::/64"
  description     = "LAN ULA to public prefix"
}

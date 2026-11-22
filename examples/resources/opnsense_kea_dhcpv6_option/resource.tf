resource "opnsense_kea_dhcpv6_option" "dns" {
  code        = "23"
  encoding    = "ipv6-address"
  data        = "2001:db8::53"
  description = "DNS server"
}

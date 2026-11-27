resource "opnsense_kea_dhcpv6_pd_pool" "example" {
  subnet_id     = opnsense_kea_dhcpv6_subnet.lan.id
  prefix        = "2001:db8:100::"
  prefix_len    = "56"
  delegated_len = "64"

  description = "example PD pool"
}

resource "opnsense_kea_dhcpv6_subnet" "lan" {
  subnet      = "2001:db8::/64"
  interface   = "lan"
  description = "LAN IPv6"

  depends_on = [opnsense_kea_dhcpv6_settings.this]
}

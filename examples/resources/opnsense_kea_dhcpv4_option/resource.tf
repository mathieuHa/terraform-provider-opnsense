resource "opnsense_kea_dhcpv4_option" "ntp" {
  code        = "42"
  encoding    = "ipv4-address"
  data        = "10.0.10.1"
  description = "NTP server"
}

resource "opnsense_kea_dhcpv4_subnet" "lan" {
  subnet  = "10.0.10.0/24"
  options = [opnsense_kea_dhcpv4_option.ntp.id]
}

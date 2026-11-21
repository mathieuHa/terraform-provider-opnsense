resource "opnsense_dnsmasq_range" "lan" {
  interface     = "lan"
  start_address = "192.168.1.100"
  end_address   = "192.168.1.199"
  lease_time    = 86400
  tag           = opnsense_dnsmasq_tag.pxe.id
}

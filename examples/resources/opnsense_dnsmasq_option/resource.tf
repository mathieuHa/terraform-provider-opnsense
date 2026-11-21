resource "opnsense_dnsmasq_option" "dns" {
  option = "6"
  value  = "192.168.1.1"
}

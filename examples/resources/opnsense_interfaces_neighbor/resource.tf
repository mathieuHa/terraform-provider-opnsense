resource "opnsense_interfaces_neighbor" "printer" {
  mac_address = "02:00:00:00:00:42"
  ip_address  = "10.0.10.42"
  description = "printer"
}

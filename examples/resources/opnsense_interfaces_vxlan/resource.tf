resource "opnsense_interfaces_vxlan" "dc" {
  vni    = 100
  local  = "192.0.2.10"
  remote = "192.0.2.20"
}

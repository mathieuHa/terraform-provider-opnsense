resource "opnsense_firewall_group" "servers" {
  name        = "servers"
  members     = ["opt1", "opt2", "opt3"]
  description = "All server networks"
}

resource "opnsense_interfaces_lagg" "uplink" {
  members     = ["igb2", "igb3"]
  proto       = "lacp"
  hash        = ["l3", "l4"]
  description = "Uplink to the core switch"
}

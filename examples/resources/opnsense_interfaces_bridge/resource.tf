resource "opnsense_interfaces_bridge" "lab" {
  members     = ["opt2", "opt3"]
  description = "Lab segment"
}

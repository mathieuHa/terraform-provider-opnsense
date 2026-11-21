resource "opnsense_dhcrelay_destination" "central" {
  name   = "central-dhcp"
  server = ["10.0.0.67", "10.0.0.68"]
}

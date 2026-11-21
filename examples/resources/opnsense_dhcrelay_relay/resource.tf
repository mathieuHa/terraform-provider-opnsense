resource "opnsense_dhcrelay_destination" "central" {
  name   = "central-dhcp"
  server = ["10.0.0.67"]
}

resource "opnsense_dhcrelay_relay" "guests" {
  interface   = "opt2"
  destination = opnsense_dhcrelay_destination.central.id
}

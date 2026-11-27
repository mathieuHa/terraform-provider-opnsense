resource "opnsense_routing_gateway" "lan_router" {
  name      = "LAN_ROUTER"
  interface = "lan"
  gateway   = "192.168.1.254"
}

// Enabled with description
resource "opnsense_route" "one_route" {
  description = "Example route"
  gateway     = opnsense_routing_gateway.lan_router.name
  network     = "10.9.0.0/24"
}

// Disabled without description
resource "opnsense_route" "two_route" {
  enabled = false

  gateway = opnsense_routing_gateway.lan_router.name
  network = "10.10.0.0/24"
}

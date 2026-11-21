resource "opnsense_interfaces_loopback" "services" {
  description = "Service addresses"
}

resource "opnsense_interfaces_assignment" "services" {
  device      = opnsense_interfaces_loopback.services.device
  description = "SERVICES"
}

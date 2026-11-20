resource "opnsense_routing_gateway" "wan_primary" {
  name      = "WAN_PRIMARY"
  interface = "wan"
  gateway   = "203.0.113.1"

  monitor_disable = false
  monitor         = "1.1.1.1"
  latency_high    = 500
  loss_high       = 20
}

resource "opnsense_routing_gateway" "wan_backup" {
  name      = "WAN_BACKUP"
  interface = "opt1"
  gateway   = "198.51.100.1"
  priority  = 10
}

resource "opnsense_routing_gateway_group" "failover" {
  name    = "WAN_FAILOVER"
  trigger = "downloss"

  members = {
    (opnsense_routing_gateway.wan_primary.name) = 1
    (opnsense_routing_gateway.wan_backup.name)  = 2
  }
}

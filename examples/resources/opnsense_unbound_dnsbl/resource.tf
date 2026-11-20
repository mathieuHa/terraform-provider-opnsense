resource "opnsense_unbound_dnsbl" "servers" {
  description = "Servers: ads and trackers"
  type        = ["oisd1"]
  blocklists  = ["telemetry.example.com"]
  allowlists  = ["updates.example.com"]
  source_nets = ["10.0.118.0/24", "10.0.119.0/24"]
  nxdomain    = true
}

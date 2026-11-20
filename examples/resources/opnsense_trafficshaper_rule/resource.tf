resource "opnsense_trafficshaper_rule" "voip_out" {
  description = "VoIP upload"
  sequence    = 10
  interface   = "wan"
  direction   = "out"
  protocol    = "udp"
  dscp        = ["ef"]
  target      = opnsense_trafficshaper_queue.voip.id
}

resource "opnsense_trafficshaper_rule" "everything_else_out" {
  description = "Everything else upload"
  sequence    = 100
  interface   = "wan"
  direction   = "out"
  target      = opnsense_trafficshaper_queue.bulk.id
}

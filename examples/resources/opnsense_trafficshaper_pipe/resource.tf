resource "opnsense_trafficshaper_pipe" "wan_upload" {
  description      = "WAN upload"
  bandwidth        = 40
  bandwidth_metric = "Mbit"
  scheduler        = "fq_codel"
}

resource "opnsense_trafficshaper_queue" "voip" {
  description = "VoIP"
  pipe        = opnsense_trafficshaper_pipe.wan_upload.id
  weight      = 90
}

resource "opnsense_trafficshaper_queue" "bulk" {
  description = "Bulk"
  pipe        = opnsense_trafficshaper_pipe.wan_upload.id
  weight      = 10
}

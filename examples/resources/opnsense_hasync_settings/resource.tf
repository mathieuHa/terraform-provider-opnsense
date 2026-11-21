import {
  to = opnsense_hasync_settings.this
  id = "hasync_settings"
}

resource "opnsense_hasync_settings" "this" {
  pfsync_interface  = "opt3"
  pfsync_peer_ip    = "10.255.255.2"
  synchronize_to_ip = "10.255.255.2"
  username          = "hasync"
  password          = var.hasync_password
  sync_items        = ["aliases", "rules", "nat", "virtualip"]
}

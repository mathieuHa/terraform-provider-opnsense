resource "opnsense_zerotier_network" "office" {
  network_id  = "8056c2e21c000001"
  description = "office"
  enabled     = true
  depends_on  = [opnsense_zerotier_settings.this]
}

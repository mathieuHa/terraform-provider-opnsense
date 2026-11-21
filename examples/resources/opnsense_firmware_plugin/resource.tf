resource "opnsense_firmware_plugin" "squid" {
  name = "os-squid"
}

resource "opnsense_proxy_settings" "this" {
  enabled    = true
  depends_on = [opnsense_firmware_plugin.squid]
}

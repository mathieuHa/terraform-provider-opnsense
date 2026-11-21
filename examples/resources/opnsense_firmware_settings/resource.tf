import {
  to = opnsense_firmware_settings.this
  id = "firmware_settings"
}

resource "opnsense_firmware_settings" "this" {
  mirror = "https://pkg.opnsense.org"
}

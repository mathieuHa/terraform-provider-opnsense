import {
  to = opnsense_nut_settings.this
  id = "nut_settings"
}

resource "opnsense_nut_settings" "this" {
  general_enable         = true
  general_name           = "ups"
  account_admin_password = var.nut_admin_password
}

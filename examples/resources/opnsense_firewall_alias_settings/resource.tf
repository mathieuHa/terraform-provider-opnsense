import {
  to = opnsense_firewall_alias_settings.this
  id = "firewall_alias_settings"
}

resource "opnsense_firewall_alias_settings" "this" {
  geoip_url = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country-CSV&license_key=${var.maxmind_key}&suffix=zip"
}

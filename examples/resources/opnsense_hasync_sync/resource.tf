resource "opnsense_hasync_sync" "push" {
  triggers = {
    rules = join(",", [opnsense_firewall_filter.allow_https.id, opnsense_firewall_filter.icmp_echo.id])
  }
  depends_on = [opnsense_hasync_settings.this]
}

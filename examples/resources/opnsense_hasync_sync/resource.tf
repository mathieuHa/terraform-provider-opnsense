resource "opnsense_hasync_sync" "push" {
  triggers = {
    rules = join(",", [for r in opnsense_firewall_filter.all : r.id])
  }
  depends_on = [opnsense_hasync_settings.this]
}

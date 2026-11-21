resource "opnsense_auth_group" "readonly" {
  name       = "readonly"
  privileges = ["page-dashboard-all", "page-diagnostics-logs"]
}

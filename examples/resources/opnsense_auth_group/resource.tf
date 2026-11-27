resource "opnsense_auth_group" "readonly" {
  name       = "readonly"
  privileges = ["page-system-login-logout", "page-diagnostics-logs-system"]
}

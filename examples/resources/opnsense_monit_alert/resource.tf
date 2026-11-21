resource "opnsense_monit_alert" "ops" {
  recipient = "ops@example.com"
  events    = ["nonexist", "pid", "resource"]
}

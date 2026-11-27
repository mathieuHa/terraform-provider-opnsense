resource "opnsense_nginx_syslog_target" "central" {
  description = "Central syslog"
  host        = "192.0.2.10"
  port        = 514
  facility    = "local7"
  severity    = "warn"
}

resource "opnsense_syslog_destination" "loki" {
  hostname    = "10.0.116.20"
  transport   = "tcp4"
  port        = 6514
  program     = ["filterlog"]
  rfc5424     = true
  description = "Firewall log to Loki"
}

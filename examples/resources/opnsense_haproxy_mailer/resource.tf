resource "opnsense_haproxy_mailer" "ops" {
  name        = "ops"
  mailservers = ["192.0.2.25:25"]
  sender      = "haproxy@example.com"
  recipient   = "ops@example.com"
  loglevel    = "alert"
}

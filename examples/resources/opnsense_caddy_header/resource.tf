resource "opnsense_caddy_header" "real_ip" {
  header = "X-Real-IP"
  value  = "{remote_host}"
}

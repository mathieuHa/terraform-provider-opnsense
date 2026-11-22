resource "opnsense_nginx_ip_acl" "office" {
  description    = "office"
  default_action = "deny"
  entries = [
    { network = "10.0.0.0/8", action = "allow" },
    { network = "192.168.0.0/16", action = "allow" },
  ]
}

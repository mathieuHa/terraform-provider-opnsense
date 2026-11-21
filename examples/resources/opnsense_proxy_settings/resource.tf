import {
  to = opnsense_proxy_settings.this
  id = "proxy_settings"
}

resource "opnsense_proxy_settings" "this" {
  enabled              = true
  interfaces           = ["lan"]
  port                 = 3128
  acl_whitelist        = ["\\.example\\.com"]
  acl_blacklist        = ["ads\\."]
  parent_proxy_enabled = true
  parent_proxy_host    = "proxy.example.net"
  parent_proxy_port    = 8080
}

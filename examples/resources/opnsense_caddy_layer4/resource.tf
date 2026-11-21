resource "opnsense_caddy_layer4" "ssh" {
  type      = "global"
  matchers  = "ssh"
  from_port = 2222
  to_domain = ["10.0.10.30"]
  to_port   = 22
}

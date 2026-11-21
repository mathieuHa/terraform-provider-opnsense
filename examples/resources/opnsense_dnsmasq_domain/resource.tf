resource "opnsense_dnsmasq_domain" "consul" {
  domain = "service.consul"
  ip     = "10.0.118.10"
  port   = 8600
}

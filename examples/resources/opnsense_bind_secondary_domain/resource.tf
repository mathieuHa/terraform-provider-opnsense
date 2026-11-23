resource "opnsense_bind_secondary_domain" "partner" {
  domainname = "partner.example"
  primaryip  = ["192.0.2.53"]
}

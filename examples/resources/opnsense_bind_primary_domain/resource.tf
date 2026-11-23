resource "opnsense_bind_primary_domain" "home" {
  domainname = "home.example"
  allowquery = [opnsense_bind_acl.lan.id]
}

resource "opnsense_bind_record" "ns" {
  domain = opnsense_bind_primary_domain.home.id
  name   = "@"
  type   = "NS"
  value  = "ns.home.example."
}

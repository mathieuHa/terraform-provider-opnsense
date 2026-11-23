resource "opnsense_bind_record" "nas" {
  domain = opnsense_bind_primary_domain.home.id
  name   = "nas"
  type   = "A"
  value  = "10.0.10.20"
}

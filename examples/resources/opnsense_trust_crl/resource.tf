resource "opnsense_trust_ca" "internal" {
  description = "Internal CA"
  common_name = "internal-ca.example.com"
}

resource "opnsense_trust_cert" "lost_laptop" {
  description = "Laptop reported lost"
  caref       = opnsense_trust_ca.internal.ref_id
  cert_type   = "usr_cert"
  common_name = "laptop-42.example.com"
}

resource "opnsense_trust_crl" "internal" {
  ca          = opnsense_trust_ca.internal.ref_id
  description = "Internal CA revocations"
  revoked = [
    { certificate = opnsense_trust_cert.lost_laptop.ref_id, reason = "keyCompromise" },
  ]
}

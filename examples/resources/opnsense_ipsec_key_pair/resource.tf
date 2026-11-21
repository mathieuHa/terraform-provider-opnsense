resource "opnsense_ipsec_key_pair" "site_b" {
  name       = "site-b"
  public_key = file("${path.module}/site-b.pub.pem")
}

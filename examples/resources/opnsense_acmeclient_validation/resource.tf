resource "opnsense_acmeclient_validation" "cloudflare" {
  name         = "cloudflare"
  method       = "dns01"
  dns_service  = "dns_cf"
  dns_cf_token = var.cloudflare_token
}

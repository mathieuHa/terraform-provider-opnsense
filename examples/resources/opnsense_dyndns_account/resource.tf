resource "opnsense_dyndns_account" "cloudflare" {
  service   = "cloudflare"
  username  = "token"
  password  = var.cloudflare_token
  zone      = "example.com"
  hostnames = ["fw.example.com"]
  checkip   = "if"
  interface = "wan"
}

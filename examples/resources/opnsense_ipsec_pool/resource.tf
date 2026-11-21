resource "opnsense_ipsec_pool" "roadwarriors" {
  name      = "roadwarriors"
  addresses = "10.10.50.0/24"
  dns       = ["10.10.0.53"]
}

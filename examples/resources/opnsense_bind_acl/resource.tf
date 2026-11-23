resource "opnsense_bind_acl" "lan" {
  name     = "lan"
  networks = ["10.0.10.0/24"]
}

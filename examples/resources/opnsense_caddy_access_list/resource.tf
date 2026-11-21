resource "opnsense_caddy_access_list" "lan" {
  name       = "lan"
  client_ips = ["10.0.0.0/8"]
}

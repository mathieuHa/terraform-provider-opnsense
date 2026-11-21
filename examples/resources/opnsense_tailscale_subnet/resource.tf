resource "opnsense_tailscale_subnet" "lan" {
  subnet      = "10.0.10.0/24"
  description = "LAN"
}

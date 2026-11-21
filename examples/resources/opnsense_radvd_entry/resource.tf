resource "opnsense_radvd_entry" "lan" {
  interface = "lan"
  mode      = "assist"
  dnssl     = ["home.example"]
}

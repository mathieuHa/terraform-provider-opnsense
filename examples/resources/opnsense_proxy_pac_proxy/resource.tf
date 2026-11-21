resource "opnsense_proxy_pac_proxy" "corp" {
  name = "corp"
  url  = "proxy.example.com:3128"
}

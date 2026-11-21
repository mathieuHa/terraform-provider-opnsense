resource "opnsense_proxy_remote_blacklist" "ads" {
  filename    = "ads"
  url         = "https://lists.example.com/ads.txt"
  description = "Advertising domains"
}

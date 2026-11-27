resource "opnsense_nginx_limit_zone" "per_client" {
  description = "Per client address"
  size        = 10
  rate        = 20
  rate_unit   = "r/s"
}

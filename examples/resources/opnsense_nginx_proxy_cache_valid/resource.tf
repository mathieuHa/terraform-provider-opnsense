resource "opnsense_nginx_proxy_cache_valid" "ok_responses" {
  description = "cache-ok"
  code        = ["200", "301"]
  valid       = 10
}

resource "opnsense_nginx_cache_path" "static" {
  path     = "/var/cache/nginx/static"
  size     = 10
  inactive = 60
  max_size = 1
}

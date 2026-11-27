resource "opnsense_nginx_http_rewrite" "old_blog" {
  description = "old-blog"
  source      = "^/blog/(.*)$"
  destination = "/news/$1"
  flag        = "permanent"
}

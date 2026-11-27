resource "opnsense_haproxy_fcgi" "php" {
  name      = "php-fpm"
  docroot   = "/var/www/html"
  index     = "index.php"
  path_info = "^(/.+\\.php)(/.*)?$"
}

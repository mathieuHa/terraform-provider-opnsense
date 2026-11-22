resource "opnsense_nginx_location" "root" {
  description = "app root"
  urlpattern  = "/"
  upstream    = opnsense_nginx_upstream.app.id
  ip_acl      = opnsense_nginx_ip_acl.office.id
}

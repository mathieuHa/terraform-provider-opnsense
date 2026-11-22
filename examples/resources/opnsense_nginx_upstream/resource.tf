resource "opnsense_nginx_upstream" "app" {
  description   = "app"
  serverentries = [opnsense_nginx_upstream_server.app1.id]
}

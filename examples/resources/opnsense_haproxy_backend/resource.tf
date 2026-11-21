resource "opnsense_haproxy_backend" "web" {
  name           = "web"
  mode           = "http"
  algorithm      = "roundrobin"
  linked_servers = [opnsense_haproxy_server.web1.id]
  health_check   = opnsense_haproxy_healthcheck.http.id
}

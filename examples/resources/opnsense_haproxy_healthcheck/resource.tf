resource "opnsense_haproxy_healthcheck" "http" {
  name = "http"
  type = "http"
}

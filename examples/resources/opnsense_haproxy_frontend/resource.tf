resource "opnsense_haproxy_frontend" "https" {
  name             = "https"
  bind             = ["0.0.0.0:443"]
  mode             = "http"
  ssl_enabled      = true
  ssl_certificates = [opnsense_trust_cert.web.id]
  default_backend  = opnsense_haproxy_backend.web.id
  linked_actions   = [opnsense_haproxy_action.api.id]
}

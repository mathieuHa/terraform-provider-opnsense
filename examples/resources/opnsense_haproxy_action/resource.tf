resource "opnsense_haproxy_action" "api" {
  name        = "api"
  type        = "use_backend"
  linked_acls = [opnsense_haproxy_acl.api_host.id]
  use_backend = opnsense_haproxy_backend.web.id
}

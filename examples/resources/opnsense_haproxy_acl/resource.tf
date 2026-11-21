resource "opnsense_haproxy_acl" "api_host" {
  name       = "api-host"
  expression = "hdr"
  hdr        = "api.example.com"
}

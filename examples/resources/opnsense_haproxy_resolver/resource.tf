resource "opnsense_haproxy_resolver" "local" {
  name            = "local"
  nameservers     = ["192.0.2.53:53"]
  resolve_retries = 3
  timeout_resolve = "1s"
  timeout_retry   = "1s"
}

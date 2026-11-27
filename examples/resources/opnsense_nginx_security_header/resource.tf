resource "opnsense_nginx_security_header" "strict" {
  description          = "Strict headers"
  referrer             = "same-origin"
  content_type_options = true

  strict_transport_security_time = 31536000

  enable_csp              = true
  csp_default_src_enabled = true
  csp_default_src_self    = true
}

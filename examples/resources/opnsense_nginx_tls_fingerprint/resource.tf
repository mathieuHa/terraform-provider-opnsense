resource "opnsense_nginx_tls_fingerprint" "curl" {
  description = "curl on Debian"
  user_agent  = "curl/8.5.0"
  trusted     = true
  ciphers     = "TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256:TLS_AES_128_GCM_SHA256"
  curves      = "X25519:prime256v1:secp384r1"
}

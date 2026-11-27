variable "tls_crypt_key" {
  description = "Output of `openvpn --genkey secret`"
  type        = string
  sensitive   = true
}

resource "opnsense_openvpn_static_key" "example" {
  description = "tls-crypt key for example server"
  mode        = "crypt"
  key         = var.tls_crypt_key
}

resource "opnsense_nginx_credential" "alice" {
  username = "alice"
  password = var.alice_password
}

variable "alice_password" {
  type      = string
  sensitive = true
}

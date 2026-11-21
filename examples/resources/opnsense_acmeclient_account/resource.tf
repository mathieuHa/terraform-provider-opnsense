resource "opnsense_acmeclient_account" "le" {
  name  = "letsencrypt"
  email = "admin@example.com"
  ca    = "letsencrypt"
}

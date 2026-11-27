resource "opnsense_nginx_userlist" "staff" {
  name  = "staff"
  users = [opnsense_nginx_credential.alice.id]
}

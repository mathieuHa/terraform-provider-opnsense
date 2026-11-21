resource "opnsense_acmeclient_certificate" "fw" {
  name              = "fw.example.com"
  account           = opnsense_acmeclient_account.le.id
  validation_method = opnsense_acmeclient_validation.cloudflare.id
  restart_actions   = [opnsense_acmeclient_action.gui.id]
}

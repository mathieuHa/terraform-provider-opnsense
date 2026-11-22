resource "opnsense_auth_user" "automation" {
  name               = "automation"
  scrambled_password = true
}

resource "opnsense_auth_user_api_key" "automation" {
  username = opnsense_auth_user.automation.name
}

output "api_key" {
  value = opnsense_auth_user_api_key.automation.key
}

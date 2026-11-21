resource "opnsense_captiveportal_zone" "guest" {
  interfaces   = ["opt2"]
  description  = "guest wifi"
  auth_servers = ["Local Database"]
  template     = opnsense_captiveportal_template.guest.id
  idle_timeout = 60
}

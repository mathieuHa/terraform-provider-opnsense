resource "opnsense_haproxy_group" "admins" {
  name         = "admins"
  members      = [opnsense_haproxy_user.stats.id]
  add_userlist = true
}

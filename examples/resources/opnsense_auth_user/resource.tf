resource "opnsense_auth_user" "monitoring" {
  name               = "monitoring"
  full_name          = "Monitoring account"
  scrambled_password = true
  group_memberships  = [opnsense_auth_group.readonly.gid]
}

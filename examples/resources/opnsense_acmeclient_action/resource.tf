resource "opnsense_acmeclient_action" "gui" {
  name = "restart web GUI"
  type = "configd_restart_gui"
}

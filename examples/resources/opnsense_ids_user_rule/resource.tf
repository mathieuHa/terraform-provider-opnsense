resource "opnsense_ids_user_rule" "scanner" {
  description = "known scanner"
  source      = "203.0.113.7"
  action      = "drop"
}

resource "opnsense_ids_policy" "drop_sslbl" {
  description = "drop abuse.ch matches"
  rulesets    = [opnsense_ids_ruleset.sslbl.uuid]
  new_action  = "drop"
}

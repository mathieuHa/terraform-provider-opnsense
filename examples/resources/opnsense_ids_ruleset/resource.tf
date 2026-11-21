resource "opnsense_ids_ruleset" "sslbl" {
  filename = "abuse.ch.sslipblacklist.rules"
}

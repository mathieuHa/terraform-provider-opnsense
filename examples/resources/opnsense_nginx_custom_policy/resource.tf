resource "opnsense_nginx_custom_policy" "sqli" {
  name        = "sqli"
  naxsi_rules = [opnsense_nginx_naxsi_rule.sql_union.id]
  operator    = ">="
  value       = 8
  action      = "BLOCK"
}

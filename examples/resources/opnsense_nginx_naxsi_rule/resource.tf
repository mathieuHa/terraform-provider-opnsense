resource "opnsense_nginx_naxsi_rule" "sql_union" {
  description = "Block UNION SELECT in arguments"
  ruletype    = "main"
  message     = "SQL injection: UNION SELECT"
  identifier  = ["42000001"]
  match_value = "union select"
  score       = 8
  args        = true
  body        = true
}

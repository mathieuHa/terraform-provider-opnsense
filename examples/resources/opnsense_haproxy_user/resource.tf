resource "opnsense_haproxy_user" "stats" {
  name        = "stats"
  description = "Read access to the statistics page"
  password    = var.haproxy_stats_password
}

variable "haproxy_stats_password" {
  type      = string
  sensitive = true
}

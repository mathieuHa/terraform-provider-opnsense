resource "opnsense_nginx_limit_request_connection" "api" {
  description = "API rate limit"
  limit_zone  = [opnsense_nginx_limit_zone.per_client.id]
  burst       = 40
  nodelay     = true
}

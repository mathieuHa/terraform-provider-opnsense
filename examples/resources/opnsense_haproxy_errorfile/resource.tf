resource "opnsense_haproxy_errorfile" "maintenance" {
  name        = "maintenance"
  description = "Shown while the backends are down"
  code        = "x503"
  content     = <<-EOT
    HTTP/1.1 503 Service Unavailable
    Cache-Control: no-cache
    Connection: close
    Content-Type: text/html

    <html><body><h1>Down for maintenance</h1></body></html>
  EOT
}

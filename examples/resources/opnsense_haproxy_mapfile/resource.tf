resource "opnsense_haproxy_mapfile" "hosts" {
  name    = "hosts"
  type    = "str"
  content = <<-EOT
    app.example.com app_backend
    api.example.com api_backend
  EOT
}

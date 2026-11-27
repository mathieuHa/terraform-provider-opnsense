resource "opnsense_haproxy_lua" "hello" {
  name    = "hello"
  content = <<-EOT
    core.register_service("hello", "http", function(applet)
      local body = "hello\n"
      applet:set_status(200)
      applet:add_header("content-length", string.len(body))
      applet:start_response()
      applet:send(body)
    end)
  EOT
}

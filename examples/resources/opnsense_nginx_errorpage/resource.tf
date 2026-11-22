resource "opnsense_nginx_errorpage" "maintenance" {
  name        = "maintenance"
  statuscodes = ["502", "503", "504"]
  pagecontent = file("${path.module}/maintenance.html")
  response    = "503"
}

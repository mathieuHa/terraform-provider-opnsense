resource "opnsense_captiveportal_template" "guest" {
  name    = "guest"
  content = filebase64("${path.module}/guest-portal.zip")
}

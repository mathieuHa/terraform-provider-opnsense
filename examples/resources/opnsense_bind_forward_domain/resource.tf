resource "opnsense_bind_forward_domain" "corp" {
  domainname    = "corp.example"
  forwardserver = ["10.20.0.53"]
}

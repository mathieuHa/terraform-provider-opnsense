// Small example
resource "opnsense_dnsmasq_host" "test_sm" {
  hostname           = "test"
  ip_addresses       = ["10.8.2.1", "::10"]
  hardware_addresses = ["00:25:96:12:34:55"]
  client_id          = "01:02:f3"
}

// Full example
resource "opnsense_dnsmasq_host" "test_xl" {
  hostname           = "test-xl"
  domain             = "example.com"
  is_local_domain    = true
  alias_records      = ["alias.example.com"]
  cname_records      = ["cname.example.com"]
  ip_addresses       = ["10.8.2.2", "::11"]
  hardware_addresses = ["00:25:96:12:34:56"]
  client_id          = "01:02:f4"
  is_ignored         = false
  comment            = "This is a test host"
  description        = "Test host"
}

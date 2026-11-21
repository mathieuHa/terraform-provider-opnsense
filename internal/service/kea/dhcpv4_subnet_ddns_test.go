package kea_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccKeaDhcpv4SubnetResource_Ddns(t *testing.T) {
	a := "opnsense_kea_dhcpv4_subnet.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "opnsense_kea_dhcpv4_subnet" "test" {
  subnet                        = "192.168.201.0/24"
  pools                         = ["192.168.201.100 - 192.168.201.150"]
  valid_lifetime                = 3600
  allocator                     = "random"
  ddns_forward_zone             = "lan.example."
  ddns_reverse_zone             = "201.168.192.in-addr.arpa."
  ddns_dns_server               = "192.168.201.53"
  ddns_dns_port                 = 5353
  ddns_key_name                 = "tfacc-key"
  ddns_key_secret               = "dGZhY2Mtc2VjcmV0LWtleQ=="
  ddns_key_algorithm            = "hmac-sha256"
  ddns_override_client_update   = true
  ddns_update_on_renew          = true
  ddns_conflict_resolution_mode = "no-check-with-dhcid"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "valid_lifetime", "3600"),
					resource.TestCheckResourceAttr(a, "allocator", "random"),
					resource.TestCheckResourceAttr(a, "ddns_forward_zone", "lan.example."),
					resource.TestCheckResourceAttr(a, "ddns_dns_port", "5353"),
					resource.TestCheckResourceAttr(a, "ddns_key_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr(a, "ddns_update_on_renew", "true"),
					resource.TestCheckResourceAttr(a, "ddns_override_no_update", "false"),
				),
			},
			{
				ResourceName:      a,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
resource "opnsense_kea_dhcpv4_subnet" "test" {
  subnet = "192.168.201.0/24"
  pools  = ["192.168.201.100 - 192.168.201.150"]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "valid_lifetime", "-1"),
					resource.TestCheckResourceAttr(a, "allocator", ""),
					resource.TestCheckResourceAttr(a, "ddns_forward_zone", ""),
					resource.TestCheckResourceAttr(a, "ddns_key_secret", ""),
					resource.TestCheckResourceAttr(a, "ddns_dns_port", "-1"),
				),
			},
		},
	})
}

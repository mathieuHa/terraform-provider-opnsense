package unbound_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUnboundDnsblResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsblConfig(`["ads.example.com", "tracker.example.com"]`, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "blocklists.#", "2"),
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "source_nets.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_unbound_dnsbl.test", "source_nets.*", "192.0.2.0/24"),
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "type.#", "0"),
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "cache_ttl", "72000"),
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "nxdomain", "false"),
				),
			},
			{
				ResourceName:      "opnsense_unbound_dnsbl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsblConfig(`["ads.example.com"]`, true) + `
data "opnsense_unbound_dnsbl" "test" {
  id = opnsense_unbound_dnsbl.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "blocklists.#", "1"),
					resource.TestCheckResourceAttr("opnsense_unbound_dnsbl.test", "nxdomain", "true"),
					resource.TestCheckResourceAttr("data.opnsense_unbound_dnsbl.test", "allowlists.#", "1"),
				),
			},
		},
	})
}

func testAccDnsblConfig(blocklists string, nxdomain bool) string {
	return fmt.Sprintf(`
resource "opnsense_unbound_dnsbl" "test" {
  description = "tf-test-dnsbl"
  allowlists  = ["allowed.example.com"]
  blocklists  = %s
  source_nets = ["192.0.2.0/24"]
  nxdomain    = %t
}
`, blocklists, nxdomain)
}

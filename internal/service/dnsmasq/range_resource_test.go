package dnsmasq_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqRangeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dnsmasq_range\" \"test\" {\n  interface = \"wan\"\n  start_address = \"192.168.150.200\"\n  end_address = \"192.168.150.210\"\n  lease_time = 3600\n  domain = \"tf.example\"\n  description = \"tf-test-range\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "start_address", "192.168.150.200"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "end_address", "192.168.150.210"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "lease_time", "3600"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "domain", "tf.example"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "description", "tf-test-range"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_range.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_range.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_range\" \"test\" {\n  interface = \"wan\"\n  start_address = \"192.168.150.200\"\n  lease_time = 7200\n  mode = [\"static\"]\n  description = \"tf-test-range-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "start_address", "192.168.150.200"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "lease_time", "7200"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "mode.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_dnsmasq_range.test", "mode.*", "static"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "description", "tf-test-range-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_range\" \"test\" {\n  interface = \"wan\"\n  start_address = \"192.168.150.200\"\n  lease_time = 7200\n  mode = [\"static\"]\n  description = \"tf-test-range-2\"\n}\n\ndata \"opnsense_dnsmasq_range\" \"test\" {\n  id = opnsense_dnsmasq_range.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_range.test", "interface", "opnsense_dnsmasq_range.test", "interface"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_range.test", "start_address", "opnsense_dnsmasq_range.test", "start_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_range.test", "lease_time", "opnsense_dnsmasq_range.test", "lease_time"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_range.test", "description", "opnsense_dnsmasq_range.test", "description"),
				),
			},
		},
	})
}

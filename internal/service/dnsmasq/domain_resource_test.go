package dnsmasq_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqDomainResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dnsmasq_domain\" \"test\" {\n  domain = \"corp.example.com\"\n  ip = \"192.0.2.53\"\n  description = \"tf-test-domain\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "domain", "corp.example.com"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "ip", "192.0.2.53"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "description", "tf-test-domain"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_domain.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_domain.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_domain\" \"test\" {\n  domain = \"corp.example.com\"\n  ip = \"192.0.2.54\"\n  port = 5353\n  sequence = 5\n  description = \"tf-test-domain-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "domain", "corp.example.com"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "ip", "192.0.2.54"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "port", "5353"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "sequence", "5"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "description", "tf-test-domain-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_domain\" \"test\" {\n  domain = \"corp.example.com\"\n  ip = \"192.0.2.54\"\n  port = 5353\n  sequence = 5\n  description = \"tf-test-domain-2\"\n}\n\ndata \"opnsense_dnsmasq_domain\" \"test\" {\n  id = opnsense_dnsmasq_domain.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_domain.test", "domain", "opnsense_dnsmasq_domain.test", "domain"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_domain.test", "ip", "opnsense_dnsmasq_domain.test", "ip"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_domain.test", "port", "opnsense_dnsmasq_domain.test", "port"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_domain.test", "sequence", "opnsense_dnsmasq_domain.test", "sequence"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_domain.test", "description", "opnsense_dnsmasq_domain.test", "description"),
				),
			},
		},
	})
}

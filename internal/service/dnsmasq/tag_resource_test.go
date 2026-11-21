package dnsmasq_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqTagResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dnsmasq_tag\" \"test\" {\n  tag = \"tftesttag\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_tag.test", "tag", "tftesttag"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_tag.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_tag\" \"test\" {\n  tag = \"tftesttag2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_tag.test", "tag", "tftesttag2"),
				),
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_tag\" \"test\" {\n  tag = \"tftesttag2\"\n}\n\ndata \"opnsense_dnsmasq_tag\" \"test\" {\n  id = opnsense_dnsmasq_tag.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_tag.test", "tag", "opnsense_dnsmasq_tag.test", "tag"),
				),
			},
		},
	})
}

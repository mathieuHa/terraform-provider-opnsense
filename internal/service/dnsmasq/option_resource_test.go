package dnsmasq_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqOptionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dnsmasq_option\" \"test\" {\n  option = \"6\"\n  value = \"192.0.2.53\"\n  description = \"tf-test-option\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "option", "6"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "value", "192.0.2.53"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "description", "tf-test-option"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_option.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_option.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_option\" \"test\" {\n  option = \"6\"\n  value = \"192.0.2.53,192.0.2.54\"\n  force = true\n  description = \"tf-test-option-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "option", "6"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "value", "192.0.2.53,192.0.2.54"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "force", "true"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "description", "tf-test-option-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_option\" \"test\" {\n  option = \"6\"\n  value = \"192.0.2.53,192.0.2.54\"\n  force = true\n  description = \"tf-test-option-2\"\n}\n\ndata \"opnsense_dnsmasq_option\" \"test\" {\n  id = opnsense_dnsmasq_option.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_option.test", "option", "opnsense_dnsmasq_option.test", "option"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_option.test", "value", "opnsense_dnsmasq_option.test", "value"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_option.test", "force", "opnsense_dnsmasq_option.test", "force"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_option.test", "description", "opnsense_dnsmasq_option.test", "description"),
				),
			},
		},
	})
}

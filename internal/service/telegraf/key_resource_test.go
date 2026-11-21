package telegraf_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTelegrafKeyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/telegraf/general/get", "plugin os-telegraf") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_telegraf_key\" \"test\" {\n  name = \"site\"\n  value = \"lab\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_telegraf_key.test", "name", "site"),
					resource.TestCheckResourceAttr("opnsense_telegraf_key.test", "value", "lab"),
					resource.TestCheckResourceAttrSet("opnsense_telegraf_key.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_telegraf_key.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_telegraf_key\" \"test\" {\n  name = \"site\"\n  value = \"lab2\"\n  enabled = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_telegraf_key.test", "name", "site"),
					resource.TestCheckResourceAttr("opnsense_telegraf_key.test", "value", "lab2"),
					resource.TestCheckResourceAttr("opnsense_telegraf_key.test", "enabled", "false"),
				),
			},
			{
				Config: "\nresource \"opnsense_telegraf_key\" \"test\" {\n  name = \"site\"\n  value = \"lab2\"\n  enabled = false\n}\n\ndata \"opnsense_telegraf_key\" \"test\" {\n  id = opnsense_telegraf_key.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_telegraf_key.test", "name", "opnsense_telegraf_key.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_telegraf_key.test", "value", "opnsense_telegraf_key.test", "value"),
					resource.TestCheckResourceAttrPair("data.opnsense_telegraf_key.test", "enabled", "opnsense_telegraf_key.test", "enabled"),
				),
			},
		},
	})
}

package caddy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaddyHeaderResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_caddy_header\" \"test\" {\n  header = \"X-Tfacc\"\n  value = \"one\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "header", "X-Tfacc"),
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "value", "one"),
					resource.TestCheckResourceAttrSet("opnsense_caddy_header.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_caddy_header.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_caddy_header\" \"test\" {\n  header = \"X-Tfacc\"\n  value = \"two\"\n  direction = \"header_down\"\n  description = \"tfacc header\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "header", "X-Tfacc"),
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "value", "two"),
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "direction", "header_down"),
					resource.TestCheckResourceAttr("opnsense_caddy_header.test", "description", "tfacc header"),
				),
			},
			{
				Config: "\nresource \"opnsense_caddy_header\" \"test\" {\n  header = \"X-Tfacc\"\n  value = \"two\"\n  direction = \"header_down\"\n  description = \"tfacc header\"\n}\n\ndata \"opnsense_caddy_header\" \"test\" {\n  id = opnsense_caddy_header.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_header.test", "header", "opnsense_caddy_header.test", "header"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_header.test", "value", "opnsense_caddy_header.test", "value"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_header.test", "direction", "opnsense_caddy_header.test", "direction"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_header.test", "description", "opnsense_caddy_header.test", "description"),
				),
			},
		},
	})
}

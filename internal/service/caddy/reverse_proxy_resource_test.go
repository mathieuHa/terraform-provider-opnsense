package caddy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaddyReverseProxyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_caddy_reverse_proxy\" \"test\" {\n  from_domain = \"tfacc.example.com\"\n  description = \"tfacc domain\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "from_domain", "tfacc.example.com"),
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "description", "tfacc domain"),
					resource.TestCheckResourceAttrSet("opnsense_caddy_reverse_proxy.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_caddy_reverse_proxy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_caddy_reverse_proxy\" \"test\" {\n  from_domain = \"tfacc2.example.com\"\n  from_port = 8443\n  description = \"tfacc domain 2\"\n  access_log = true\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "from_domain", "tfacc2.example.com"),
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "from_port", "8443"),
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "description", "tfacc domain 2"),
					resource.TestCheckResourceAttr("opnsense_caddy_reverse_proxy.test", "access_log", "true"),
				),
			},
			{
				Config: "\nresource \"opnsense_caddy_reverse_proxy\" \"test\" {\n  from_domain = \"tfacc2.example.com\"\n  from_port = 8443\n  description = \"tfacc domain 2\"\n  access_log = true\n}\n\ndata \"opnsense_caddy_reverse_proxy\" \"test\" {\n  id = opnsense_caddy_reverse_proxy.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_reverse_proxy.test", "from_domain", "opnsense_caddy_reverse_proxy.test", "from_domain"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_reverse_proxy.test", "from_port", "opnsense_caddy_reverse_proxy.test", "from_port"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_reverse_proxy.test", "description", "opnsense_caddy_reverse_proxy.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_reverse_proxy.test", "access_log", "opnsense_caddy_reverse_proxy.test", "access_log"),
				),
			},
		},
	})
}

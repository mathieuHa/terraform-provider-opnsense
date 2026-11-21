package caddy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaddyBasicAuthResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_caddy_basic_auth\" \"test\" {\n  user = \"tfacc\"\n  password = \"tfacc-secret-1\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_basic_auth.test", "user", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_caddy_basic_auth.test", "password", "tfacc-secret-1"),
					resource.TestCheckResourceAttrSet("opnsense_caddy_basic_auth.test", "id"),
				),
			},
			{
				ResourceName:            "opnsense_caddy_basic_auth.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: "\nresource \"opnsense_caddy_basic_auth\" \"test\" {\n  user = \"tfacc2\"\n  password = \"tfacc-secret-2\"\n  description = \"tfacc user\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_basic_auth.test", "user", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_caddy_basic_auth.test", "password", "tfacc-secret-2"),
					resource.TestCheckResourceAttr("opnsense_caddy_basic_auth.test", "description", "tfacc user"),
				),
			},
			{
				Config: "\nresource \"opnsense_caddy_basic_auth\" \"test\" {\n  user = \"tfacc2\"\n  password = \"tfacc-secret-2\"\n  description = \"tfacc user\"\n}\n\ndata \"opnsense_caddy_basic_auth\" \"test\" {\n  id = opnsense_caddy_basic_auth.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_basic_auth.test", "user", "opnsense_caddy_basic_auth.test", "user"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_basic_auth.test", "description", "opnsense_caddy_basic_auth.test", "description"),
				),
			},
		},
	})
}

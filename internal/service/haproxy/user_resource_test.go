package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_user\" \"test\" {\n  name = \"tfacc\"\n  password = \"tfacc-secret-1\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_user.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_user.test", "password", "tfacc-secret-1"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_user.test", "id"),
				),
			},
			{
				ResourceName:            "opnsense_haproxy_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: "\nresource \"opnsense_haproxy_user\" \"test\" {\n  name = \"tfacc2\"\n  password = \"tfacc-secret-2\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_user.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_user.test", "password", "tfacc-secret-2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_user.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_user\" \"test\" {\n  name = \"tfacc2\"\n  password = \"tfacc-secret-2\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_user\" \"test\" {\n  id = opnsense_haproxy_user.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_user.test", "name", "opnsense_haproxy_user.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_user.test", "description", "opnsense_haproxy_user.test", "description"),
				),
			},
		},
	})
}

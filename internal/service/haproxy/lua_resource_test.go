package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyLuaResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_lua\" \"test\" {\n  name = \"tfacc\"\n  content = \"core.log(core.info, \\\"tfacc\\\")\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_lua.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_lua.test", "content", "core.log(core.info, \"tfacc\")"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_lua.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_lua.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_lua\" \"test\" {\n  name = \"tfacc2\"\n  content = \"core.log(core.info, \\\"tfacc2\\\")\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_lua.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_lua.test", "content", "core.log(core.info, \"tfacc2\")"),
					resource.TestCheckResourceAttr("opnsense_haproxy_lua.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_lua\" \"test\" {\n  name = \"tfacc2\"\n  content = \"core.log(core.info, \\\"tfacc2\\\")\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_lua\" \"test\" {\n  id = opnsense_haproxy_lua.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_lua.test", "name", "opnsense_haproxy_lua.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_lua.test", "content", "opnsense_haproxy_lua.test", "content"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_lua.test", "description", "opnsense_haproxy_lua.test", "description"),
				),
			},
		},
	})
}

package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyMapfileResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_mapfile\" \"test\" {\n  name = \"tfacc\"\n  content = \"tfacc.example.com backend1\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_mapfile.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mapfile.test", "content", "tfacc.example.com backend1"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_mapfile.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_mapfile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_mapfile\" \"test\" {\n  name = \"tfacc2\"\n  content = \"tfacc2.example.com backend2\"\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_mapfile.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mapfile.test", "content", "tfacc2.example.com backend2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_mapfile.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_mapfile\" \"test\" {\n  name = \"tfacc2\"\n  content = \"tfacc2.example.com backend2\"\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_mapfile\" \"test\" {\n  id = opnsense_haproxy_mapfile.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mapfile.test", "name", "opnsense_haproxy_mapfile.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mapfile.test", "content", "opnsense_haproxy_mapfile.test", "content"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_mapfile.test", "description", "opnsense_haproxy_mapfile.test", "description"),
				),
			},
		},
	})
}

package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyServerResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_server\" \"test\" {\n  name = \"tfacc\"\n  address = \"10.99.2.10\"\n  port = 8080\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "address", "10.99.2.10"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "port", "8080"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_server.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_server.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_server\" \"test\" {\n  name = \"tfacc2\"\n  address = \"10.99.2.11\"\n  port = 8081\n  description = \"tfacc\"\n  weight = 50\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "address", "10.99.2.11"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "port", "8081"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "description", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_server.test", "weight", "50"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_server\" \"test\" {\n  name = \"tfacc2\"\n  address = \"10.99.2.11\"\n  port = 8081\n  description = \"tfacc\"\n  weight = 50\n}\n\ndata \"opnsense_haproxy_server\" \"test\" {\n  id = opnsense_haproxy_server.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_server.test", "name", "opnsense_haproxy_server.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_server.test", "address", "opnsense_haproxy_server.test", "address"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_server.test", "port", "opnsense_haproxy_server.test", "port"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_server.test", "description", "opnsense_haproxy_server.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_server.test", "weight", "opnsense_haproxy_server.test", "weight"),
				),
			},
		},
	})
}

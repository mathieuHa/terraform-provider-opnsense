package haproxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHaproxyResolverResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/haproxy/settings/get", "plugin os-haproxy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_haproxy_resolver\" \"test\" {\n  name = \"tfacc\"\n  nameservers = [\"192.0.2.53:53\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_resolver.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_haproxy_resolver.test", "nameservers.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_resolver.test", "nameservers.*", "192.0.2.53:53"),
					resource.TestCheckResourceAttrSet("opnsense_haproxy_resolver.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_haproxy_resolver.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_haproxy_resolver\" \"test\" {\n  name = \"tfacc2\"\n  nameservers = [\"192.0.2.53:53\", \"192.0.2.54:53\"]\n  description = \"tfacc\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_haproxy_resolver.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_haproxy_resolver.test", "nameservers.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_resolver.test", "nameservers.*", "192.0.2.53:53"),
					resource.TestCheckTypeSetElemAttr("opnsense_haproxy_resolver.test", "nameservers.*", "192.0.2.54:53"),
					resource.TestCheckResourceAttr("opnsense_haproxy_resolver.test", "description", "tfacc"),
				),
			},
			{
				Config: "\nresource \"opnsense_haproxy_resolver\" \"test\" {\n  name = \"tfacc2\"\n  nameservers = [\"192.0.2.53:53\", \"192.0.2.54:53\"]\n  description = \"tfacc\"\n}\n\ndata \"opnsense_haproxy_resolver\" \"test\" {\n  id = opnsense_haproxy_resolver.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_resolver.test", "name", "opnsense_haproxy_resolver.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_haproxy_resolver.test", "description", "opnsense_haproxy_resolver.test", "description"),
				),
			},
		},
	})
}

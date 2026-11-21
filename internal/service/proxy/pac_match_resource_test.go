package proxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProxyPacMatchResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_proxy_pac_match\" \"test\" {\n  name = \"tfacc\"\n  match_type = \"hostname_matches\"\n  hostname = \"*.example.com\"\n  description = \"tf-acc pac match\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "match_type", "hostname_matches"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "hostname", "*.example.com"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "description", "tf-acc pac match"),
					resource.TestCheckResourceAttrSet("opnsense_proxy_pac_match.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_proxy_pac_match.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_proxy_pac_match\" \"test\" {\n  name = \"tfacc\"\n  match_type = \"time_range\"\n  time_from = 8\n  time_to = 18\n  negate = true\n  description = \"tf-acc pac match 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "match_type", "time_range"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "time_from", "8"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "time_to", "18"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "negate", "true"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_match.test", "description", "tf-acc pac match 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_proxy_pac_match\" \"test\" {\n  name = \"tfacc\"\n  match_type = \"time_range\"\n  time_from = 8\n  time_to = 18\n  negate = true\n  description = \"tf-acc pac match 2\"\n}\n\ndata \"opnsense_proxy_pac_match\" \"test\" {\n  id = opnsense_proxy_pac_match.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "name", "opnsense_proxy_pac_match.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "match_type", "opnsense_proxy_pac_match.test", "match_type"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "time_from", "opnsense_proxy_pac_match.test", "time_from"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "time_to", "opnsense_proxy_pac_match.test", "time_to"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "negate", "opnsense_proxy_pac_match.test", "negate"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_match.test", "description", "opnsense_proxy_pac_match.test", "description"),
				),
			},
		},
	})
}

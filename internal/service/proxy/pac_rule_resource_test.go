package proxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProxyPacRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "resource \"opnsense_proxy_pac_proxy\" \"p\" {\n  name = \"tfacc-rule\"\n  url  = \"proxy.example.com:3128\"\n}\n\nresource \"opnsense_proxy_pac_match\" \"m\" {\n  name       = \"tfacc-rule\"\n  match_type = \"dns_domain_is\"\n  hostname   = \".example.com\"\n}\n\nresource \"opnsense_proxy_pac_rule\" \"test\" {\n  matches = [\"${opnsense_proxy_pac_match.m.id}\"]\n  proxies = [\"${opnsense_proxy_pac_proxy.p.id}\"]\n  description = \"tf-acc pac rule\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "matches.#", "1"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "proxies.#", "1"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "description", "tf-acc pac rule"),
					resource.TestCheckResourceAttrSet("opnsense_proxy_pac_rule.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_proxy_pac_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "resource \"opnsense_proxy_pac_proxy\" \"p\" {\n  name = \"tfacc-rule\"\n  url  = \"proxy.example.com:3128\"\n}\n\nresource \"opnsense_proxy_pac_match\" \"m\" {\n  name       = \"tfacc-rule\"\n  match_type = \"dns_domain_is\"\n  hostname   = \".example.com\"\n}\n\nresource \"opnsense_proxy_pac_rule\" \"test\" {\n  matches = [\"${opnsense_proxy_pac_match.m.id}\"]\n  proxies = [\"${opnsense_proxy_pac_proxy.p.id}\"]\n  match_type = \"unless\"\n  enabled = false\n  description = \"tf-acc pac rule 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "matches.#", "1"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "proxies.#", "1"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "match_type", "unless"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_rule.test", "description", "tf-acc pac rule 2"),
				),
			},
			{
				Config: "resource \"opnsense_proxy_pac_proxy\" \"p\" {\n  name = \"tfacc-rule\"\n  url  = \"proxy.example.com:3128\"\n}\n\nresource \"opnsense_proxy_pac_match\" \"m\" {\n  name       = \"tfacc-rule\"\n  match_type = \"dns_domain_is\"\n  hostname   = \".example.com\"\n}\n\nresource \"opnsense_proxy_pac_rule\" \"test\" {\n  matches = [\"${opnsense_proxy_pac_match.m.id}\"]\n  proxies = [\"${opnsense_proxy_pac_proxy.p.id}\"]\n  match_type = \"unless\"\n  enabled = false\n  description = \"tf-acc pac rule 2\"\n}\n\ndata \"opnsense_proxy_pac_rule\" \"test\" {\n  id = opnsense_proxy_pac_rule.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_rule.test", "match_type", "opnsense_proxy_pac_rule.test", "match_type"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_rule.test", "enabled", "opnsense_proxy_pac_rule.test", "enabled"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_rule.test", "description", "opnsense_proxy_pac_rule.test", "description"),
				),
			},
		},
	})
}

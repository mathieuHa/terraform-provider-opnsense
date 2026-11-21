package proxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProxyPacProxyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_proxy_pac_proxy\" \"test\" {\n  name = \"tfacc\"\n  url = \"proxy.example.com:3128\"\n  description = \"tf-acc pac proxy\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "url", "proxy.example.com:3128"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "description", "tf-acc pac proxy"),
					resource.TestCheckResourceAttrSet("opnsense_proxy_pac_proxy.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_proxy_pac_proxy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_proxy_pac_proxy\" \"test\" {\n  name = \"tfacc2\"\n  proxy_type = \"SOCKS5\"\n  url = \"proxy.example.com:1080\"\n  description = \"tf-acc pac proxy 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "proxy_type", "SOCKS5"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "url", "proxy.example.com:1080"),
					resource.TestCheckResourceAttr("opnsense_proxy_pac_proxy.test", "description", "tf-acc pac proxy 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_proxy_pac_proxy\" \"test\" {\n  name = \"tfacc2\"\n  proxy_type = \"SOCKS5\"\n  url = \"proxy.example.com:1080\"\n  description = \"tf-acc pac proxy 2\"\n}\n\ndata \"opnsense_proxy_pac_proxy\" \"test\" {\n  id = opnsense_proxy_pac_proxy.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_proxy.test", "name", "opnsense_proxy_pac_proxy.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_proxy.test", "proxy_type", "opnsense_proxy_pac_proxy.test", "proxy_type"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_proxy.test", "url", "opnsense_proxy_pac_proxy.test", "url"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_pac_proxy.test", "description", "opnsense_proxy_pac_proxy.test", "description"),
				),
			},
		},
	})
}

package proxy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProxyRemoteBlacklistResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/proxy/settings/get", "requires the os-squid plugin") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_proxy_remote_blacklist\" \"test\" {\n  filename = \"tfacc\"\n  url = \"http://192.0.2.10/list.txt\"\n  description = \"tf acc remote list\"\n  enabled = false\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "filename", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "url", "http://192.0.2.10/list.txt"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "description", "tf acc remote list"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "enabled", "false"),
					resource.TestCheckResourceAttrSet("opnsense_proxy_remote_blacklist.test", "id"),
				),
			},
			{
				ResourceName:            "opnsense_proxy_remote_blacklist.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: "\nresource \"opnsense_proxy_remote_blacklist\" \"test\" {\n  filename = \"tfacc\"\n  url = \"http://192.0.2.11/list.txt\"\n  description = \"tf acc remote list 2\"\n  enabled = false\n  ssl_no_verify = true\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "filename", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "url", "http://192.0.2.11/list.txt"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "description", "tf acc remote list 2"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "enabled", "false"),
					resource.TestCheckResourceAttr("opnsense_proxy_remote_blacklist.test", "ssl_no_verify", "true"),
				),
			},
			{
				Config: "\nresource \"opnsense_proxy_remote_blacklist\" \"test\" {\n  filename = \"tfacc\"\n  url = \"http://192.0.2.11/list.txt\"\n  description = \"tf acc remote list 2\"\n  enabled = false\n  ssl_no_verify = true\n}\n\ndata \"opnsense_proxy_remote_blacklist\" \"test\" {\n  id = opnsense_proxy_remote_blacklist.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_remote_blacklist.test", "filename", "opnsense_proxy_remote_blacklist.test", "filename"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_remote_blacklist.test", "url", "opnsense_proxy_remote_blacklist.test", "url"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_remote_blacklist.test", "description", "opnsense_proxy_remote_blacklist.test", "description"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_remote_blacklist.test", "enabled", "opnsense_proxy_remote_blacklist.test", "enabled"),
					resource.TestCheckResourceAttrPair("data.opnsense_proxy_remote_blacklist.test", "ssl_no_verify", "opnsense_proxy_remote_blacklist.test", "ssl_no_verify"),
				),
			},
		},
	})
}

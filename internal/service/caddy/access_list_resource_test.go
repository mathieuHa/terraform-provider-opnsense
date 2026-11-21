package caddy_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaddyAccessListResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.EndpointPreCheck(t, "/caddy/general/get", "plugin os-caddy") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_caddy_access_list\" \"test\" {\n  name = \"tfacc\"\n  client_ips = [\"10.0.0.0/8\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "client_ips.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_caddy_access_list.test", "client_ips.*", "10.0.0.0/8"),
					resource.TestCheckResourceAttrSet("opnsense_caddy_access_list.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_caddy_access_list.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_caddy_access_list\" \"test\" {\n  name = \"tfacc2\"\n  client_ips = [\"10.0.0.0/8\", \"192.168.0.0/16\"]\n  invert = true\n  http_response_code = 403\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "client_ips.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_caddy_access_list.test", "client_ips.*", "10.0.0.0/8"),
					resource.TestCheckTypeSetElemAttr("opnsense_caddy_access_list.test", "client_ips.*", "192.168.0.0/16"),
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "invert", "true"),
					resource.TestCheckResourceAttr("opnsense_caddy_access_list.test", "http_response_code", "403"),
				),
			},
			{
				Config: "\nresource \"opnsense_caddy_access_list\" \"test\" {\n  name = \"tfacc2\"\n  client_ips = [\"10.0.0.0/8\", \"192.168.0.0/16\"]\n  invert = true\n  http_response_code = 403\n}\n\ndata \"opnsense_caddy_access_list\" \"test\" {\n  id = opnsense_caddy_access_list.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_access_list.test", "name", "opnsense_caddy_access_list.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_access_list.test", "invert", "opnsense_caddy_access_list.test", "invert"),
					resource.TestCheckResourceAttrPair("data.opnsense_caddy_access_list.test", "http_response_code", "opnsense_caddy_access_list.test", "http_response_code"),
				),
			},
		},
	})
}

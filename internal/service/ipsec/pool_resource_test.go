package ipsec_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIpsecPoolResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_ipsec_pool\" \"test\" {\n  name = \"tfacc\"\n  addresses = \"10.199.0.0/24\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_pool.test", "name", "tfacc"),
					resource.TestCheckResourceAttr("opnsense_ipsec_pool.test", "addresses", "10.199.0.0/24"),
					resource.TestCheckResourceAttrSet("opnsense_ipsec_pool.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_ipsec_pool.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_ipsec_pool\" \"test\" {\n  name = \"tfacc2\"\n  addresses = \"10.199.1.0/24\"\n  dns = [\"10.0.0.53\"]\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_pool.test", "name", "tfacc2"),
					resource.TestCheckResourceAttr("opnsense_ipsec_pool.test", "addresses", "10.199.1.0/24"),
					resource.TestCheckResourceAttr("opnsense_ipsec_pool.test", "dns.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_ipsec_pool.test", "dns.*", "10.0.0.53"),
				),
			},
			{
				Config: "\nresource \"opnsense_ipsec_pool\" \"test\" {\n  name = \"tfacc2\"\n  addresses = \"10.199.1.0/24\"\n  dns = [\"10.0.0.53\"]\n}\n\ndata \"opnsense_ipsec_pool\" \"test\" {\n  id = opnsense_ipsec_pool.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_pool.test", "name", "opnsense_ipsec_pool.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_pool.test", "addresses", "opnsense_ipsec_pool.test", "addresses"),
				),
			},
		},
	})
}

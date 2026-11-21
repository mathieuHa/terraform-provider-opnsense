package ipsec_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIpsecManualSpdResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_ipsec_manual_spd\" \"test\" {\n  request_id = 100\n  source = \"10.198.0.0/24\"\n  destination = \"10.197.0.0/24\"\n  description = \"tf-acc spd\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "request_id", "100"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "source", "10.198.0.0/24"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "destination", "10.197.0.0/24"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "description", "tf-acc spd"),
					resource.TestCheckResourceAttrSet("opnsense_ipsec_manual_spd.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_ipsec_manual_spd.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_ipsec_manual_spd\" \"test\" {\n  request_id = 101\n  source = \"10.198.1.0/24\"\n  protocol = \"ah\"\n  description = \"tf-acc spd 2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "request_id", "101"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "source", "10.198.1.0/24"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "protocol", "ah"),
					resource.TestCheckResourceAttr("opnsense_ipsec_manual_spd.test", "description", "tf-acc spd 2"),
				),
			},
			{
				Config: "\nresource \"opnsense_ipsec_manual_spd\" \"test\" {\n  request_id = 101\n  source = \"10.198.1.0/24\"\n  protocol = \"ah\"\n  description = \"tf-acc spd 2\"\n}\n\ndata \"opnsense_ipsec_manual_spd\" \"test\" {\n  id = opnsense_ipsec_manual_spd.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_manual_spd.test", "request_id", "opnsense_ipsec_manual_spd.test", "request_id"),
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_manual_spd.test", "source", "opnsense_ipsec_manual_spd.test", "source"),
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_manual_spd.test", "protocol", "opnsense_ipsec_manual_spd.test", "protocol"),
					resource.TestCheckResourceAttrPair("data.opnsense_ipsec_manual_spd.test", "description", "opnsense_ipsec_manual_spd.test", "description"),
				),
			},
		},
	})
}

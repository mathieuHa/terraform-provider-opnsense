package dnsmasq_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqBootResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: "\nresource \"opnsense_dnsmasq_boot\" \"test\" {\n  filename = \"pxelinux.0\"\n  server_address = \"192.0.2.69\"\n  description = \"tf-test-boot\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "filename", "pxelinux.0"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "server_address", "192.0.2.69"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "description", "tf-test-boot"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_boot.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_boot.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_boot\" \"test\" {\n  filename = \"ipxe.efi\"\n  server_address = \"192.0.2.70\"\n  server_name = \"tftp\"\n  description = \"tf-test-boot-2\"\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "filename", "ipxe.efi"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "server_address", "192.0.2.70"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "server_name", "tftp"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "description", "tf-test-boot-2"),
				),
			},
			{
				Config: "\nresource \"opnsense_dnsmasq_boot\" \"test\" {\n  filename = \"ipxe.efi\"\n  server_address = \"192.0.2.70\"\n  server_name = \"tftp\"\n  description = \"tf-test-boot-2\"\n}\n\ndata \"opnsense_dnsmasq_boot\" \"test\" {\n  id = opnsense_dnsmasq_boot.test.id\n}\n",
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_boot.test", "filename", "opnsense_dnsmasq_boot.test", "filename"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_boot.test", "server_address", "opnsense_dnsmasq_boot.test", "server_address"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_boot.test", "server_name", "opnsense_dnsmasq_boot.test", "server_name"),
					resource.TestCheckResourceAttrPair("data.opnsense_dnsmasq_boot.test", "description", "opnsense_dnsmasq_boot.test", "description"),
				),
			},
		},
	})
}

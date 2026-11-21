package openvpn_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOpenvpnInstanceResource_AutoVPNID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOpenvpnInstanceAutoVPNIDConfig("tf-acc-auto-vpnid"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("opnsense_openvpn_instance.test", "vpn_id"),
					resource.TestCheckResourceAttrSet("opnsense_openvpn_instance.test", "id"),
				),
			},
			{
				Config: testAccOpenvpnInstanceAutoVPNIDConfig("tf-acc-auto-vpnid-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_openvpn_instance.test", "description", "tf-acc-auto-vpnid-updated"),
					resource.TestCheckResourceAttrSet("opnsense_openvpn_instance.test", "vpn_id"),
				),
			},
		},
	})
}

func testAccOpenvpnInstanceAutoVPNIDConfig(description string) string {
	return `
resource "opnsense_trust_ca" "test" {
  description = "tf-acc-auto-vpnid"
  common_name = "tf-acc-auto-vpnid"
}

resource "opnsense_openvpn_instance" "test" {
  certificate_authority = opnsense_trust_ca.test.ref_id
  description        = "` + description + `"
  role               = "client"
  dev_type           = "tun"
  protocol           = "udp"
  remote             = ["192.0.2.10:1194"]
  verify_client_cert = "none"
}
`
}

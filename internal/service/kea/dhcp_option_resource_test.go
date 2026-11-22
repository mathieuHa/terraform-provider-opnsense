package kea_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const keaOptionConfig = `
resource "opnsense_kea_dhcpv4_option" "wpad" {
  code        = "252"
  encoding    = "hex"
  data        = "%[1]s"
  description = "tfacc wpad"
}

resource "opnsense_kea_dhcpv4_option" "ntp" {
  code        = "42"
  encoding    = "ipv4-address"
  data        = "192.0.2.123"
  force       = true
  description = "tfacc ntp"
}

resource "opnsense_kea_dhcpv4_subnet" "test" {
  subnet      = "192.168.202.0/24"
  pools       = ["192.168.202.100 - 192.168.202.150"]
  options     = [%[2]s]
  description = "tfacc options subnet"
}

resource "opnsense_kea_dhcpv4_reservation" "test" {
  subnet_id   = opnsense_kea_dhcpv4_subnet.test.id
  ip_address  = "192.168.202.42"
  mac_address = "02:00:00:00:02:42"
  options     = [opnsense_kea_dhcpv4_option.ntp.id]
  description = "tfacc options reservation"
}
`

func TestAccKeaDhcpv4OptionResource(t *testing.T) {
	both := "opnsense_kea_dhcpv4_option.wpad.id, opnsense_kea_dhcpv4_option.ntp.id"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(keaOptionConfig, "687474703a2f2f777061642e6578616d706c652f777061642e646174", both),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_option.wpad", "code", "252"),
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_option.ntp", "force", "true"),
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_subnet.test", "options.#", "2"),
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_reservation.test", "options.#", "1"),
				),
			},
			{
				ResourceName:      "opnsense_kea_dhcpv4_option.wpad",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(keaOptionConfig, "687474703a2f2f777061642e6578616d706c652f70726f78792e706163", "opnsense_kea_dhcpv4_option.wpad.id"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_option.wpad", "data", "687474703a2f2f777061642e6578616d706c652f70726f78792e706163"),
					resource.TestCheckResourceAttr("opnsense_kea_dhcpv4_subnet.test", "options.#", "1"),
				),
			},
		},
	})
}

func TestAccKeaDhcpv6OptionResource(t *testing.T) {
	a := "opnsense_kea_dhcpv6_option.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "opnsense_kea_dhcpv6_option" "test" {
  code        = "23"
  encoding    = "ipv6-address"
  data        = "2001:db8::53"
  description = "tfacc dns"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "code", "23"),
					resource.TestCheckResourceAttr(a, "encoding", "ipv6-address"),
				),
			},
			{
				ResourceName:      a,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

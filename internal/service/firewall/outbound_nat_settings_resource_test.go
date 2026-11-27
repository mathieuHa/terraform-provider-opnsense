package firewall_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFirewallOutboundNatSettingsResource(t *testing.T) {
	acctest.AccPreCheck(t)

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Firewall().OutboundNatConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_firewall_outbound_nat_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_firewall_outbound_nat_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "firewall_outbound_nat_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_firewall_outbound_nat_settings" "test" {
  mode = "hybrid"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "mode", "hybrid"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_firewall_outbound_nat_settings" "test" {
  mode = %q
}
`, cur.Monad.General.Mode.String()),
			},
		},
	})
}

func TestAccFirewallAliasSettingsResource(t *testing.T) {
	acctest.AccPreCheck(t)

	c := opnsense.NewClient(acctest.Client())
	cur, err := c.Firewall().AliasConfigGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := "opnsense_firewall_alias_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_firewall_alias_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "firewall_alias_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_firewall_alias_settings" "test" {
  geoip_url = "https://geoip.example.com/GeoLite2-Country-CSV.zip"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(a, "geoip_url", "https://geoip.example.com/GeoLite2-Country-CSV.zip"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_firewall_alias_settings" "test" {
  geoip_url = %q
}
`, cur.Monad.Geoip.GeoipUrl),
			},
		},
	})
}

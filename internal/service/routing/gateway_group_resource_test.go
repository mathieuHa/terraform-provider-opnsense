package routing_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func gatewayGroupPreCheck(t *testing.T) {
	acctest.EndpointPreCheck(t, "/routing/group_settings/get", "gateway groups require OPNsense 26.7 or later")
}

const testAccGatewayGroupGateways = `
resource "opnsense_routing_gateway" "primary" {
  name        = "tf_test_gw_primary"
  interface   = "wan"
  gateway     = "192.0.2.1"
  far_gateway = true
}

resource "opnsense_routing_gateway" "backup" {
  name        = "tf_test_gw_backup"
  interface   = "wan"
  gateway     = "192.0.2.2"
  far_gateway = true
}
`

func TestAccRoutingGatewayGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { gatewayGroupPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGatewayGroupGateways + `
resource "opnsense_routing_gateway_group" "test" {
  name = "tf_test_failover"
  members = {
    (opnsense_routing_gateway.primary.name) = 1
    (opnsense_routing_gateway.backup.name)  = 2
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "name", "tf_test_failover"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "members.%", "2"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "members.tf_test_gw_primary", "1"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "members.tf_test_gw_backup", "2"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "trigger", "down"),
					resource.TestCheckNoResourceAttr("opnsense_routing_gateway_group.test", "pool_options"),
					resource.TestCheckResourceAttrSet("opnsense_routing_gateway_group.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_routing_gateway_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGatewayGroupGateways + `
resource "opnsense_routing_gateway_group" "test" {
  name         = "tf_test_failover"
  description  = "load balanced"
  trigger      = "downloss"
  pool_options = "round-robin sticky-address"
  members = {
    (opnsense_routing_gateway.primary.name) = 1
    (opnsense_routing_gateway.backup.name)  = 1
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "members.tf_test_gw_backup", "1"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "trigger", "downloss"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "pool_options", "round-robin sticky-address"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway_group.test", "description", "load balanced"),
				),
			},
			{
				ResourceName:      "opnsense_routing_gateway_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRoutingGatewayGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { gatewayGroupPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGatewayGroupGateways + `
resource "opnsense_routing_gateway_group" "test" {
  name = "tf_test_failover_ds"
  members = {
    (opnsense_routing_gateway.primary.name) = 1
    (opnsense_routing_gateway.backup.name)  = 3
  }
}

data "opnsense_routing_gateway_group" "test" {
  id = opnsense_routing_gateway_group.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.opnsense_routing_gateway_group.test", "name", "tf_test_failover_ds"),
					resource.TestCheckResourceAttr("data.opnsense_routing_gateway_group.test", "members.tf_test_gw_backup", "3"),
				),
			},
		},
	})
}

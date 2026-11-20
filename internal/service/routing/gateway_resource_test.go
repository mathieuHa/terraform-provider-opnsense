package routing_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRoutingGatewayResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoutingGatewayConfig("192.0.2.1", 255),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "name", "tf_test_gw"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "enabled", "true"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "ip_protocol", "inet"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "gateway", "192.0.2.1"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "far_gateway", "true"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "monitor_disable", "true"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "priority", "255"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "latency_low", "-1"),
					resource.TestCheckResourceAttrSet("opnsense_routing_gateway.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_routing_gateway.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRoutingGatewayConfig("192.0.2.2", 10),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "gateway", "192.0.2.2"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "priority", "10"),
				),
			},
		},
	})
}

func TestAccRoutingGatewayResource_Monitoring(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoutingGatewayMonitoringConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "monitor_disable", "false"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "monitor", "192.0.2.53"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "latency_low", "200"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "latency_high", "500"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "loss_low", "10"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "loss_high", "20"),
					resource.TestCheckResourceAttr("opnsense_routing_gateway.test", "description", "monitored test gateway"),
				),
			},
			{
				ResourceName:      "opnsense_routing_gateway.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRoutingGatewayDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoutingGatewayConfig("192.0.2.1", 255) + `
data "opnsense_routing_gateway" "test" {
  id = opnsense_routing_gateway.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_routing_gateway.test", "name", "opnsense_routing_gateway.test", "name"),
					resource.TestCheckResourceAttrPair("data.opnsense_routing_gateway.test", "gateway", "opnsense_routing_gateway.test", "gateway"),
					resource.TestCheckResourceAttr("data.opnsense_routing_gateway.test", "priority", "255"),
				),
			},
		},
	})
}

func testAccRoutingGatewayConfig(address string, priority int) string {
	return fmt.Sprintf(`
resource "opnsense_routing_gateway" "test" {
  name        = "tf_test_gw"
  interface   = "wan"
  gateway     = %[1]q
  far_gateway = true
  priority    = %[2]d
}
`, address, priority)
}

func testAccRoutingGatewayMonitoringConfig() string {
	return `
resource "opnsense_routing_gateway" "test" {
  name            = "tf_test_gw_mon"
  description     = "monitored test gateway"
  interface       = "wan"
  gateway         = "192.0.2.1"
  far_gateway     = true
  monitor_disable = false
  monitor         = "192.0.2.53"
  latency_low     = 200
  latency_high    = 500
  loss_low        = 10
  loss_high       = 20
}
`
}

package routes_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	testRouteGateway   = "tf_test_route_gw"
	testRouteGatewayV6 = "WAN_DHCP6"
)

func TestAccRouteResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouteResourceConfig(testRouteGateway, "192.0.2.0/24"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_route.test", "enabled", "true"),
					resource.TestCheckResourceAttr("opnsense_route.test", "gateway", testRouteGateway),
					resource.TestCheckResourceAttr("opnsense_route.test", "network", "192.0.2.0/24"),
					resource.TestCheckResourceAttrSet("opnsense_route.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRouteResourceConfig(testRouteGateway, "192.0.2.128/25"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_route.test", "network", "192.0.2.128/25"),
				),
			},
		},
	})
}

func TestAccRouteResource_Disabled(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouteResourceConfigDisabled(testRouteGateway, "192.0.2.0/24"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_route.test", "enabled", "false"),
					resource.TestCheckResourceAttrSet("opnsense_route.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRouteResource_WithDescription(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouteResourceConfigWithDescription(testRouteGateway, "192.0.2.0/24", "Test static route"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_route.test", "description", "Test static route"),
					resource.TestCheckResourceAttrSet("opnsense_route.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRouteResource_IPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouteResourceConfig(testRouteGatewayV6, "2001:db8::/32"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_route.test", "network", "2001:db8::/32"),
					resource.TestCheckResourceAttrSet("opnsense_route.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func routeGateway(gateway string) (string, string) {
	if gateway != testRouteGateway {
		return "", fmt.Sprintf("%q", gateway)
	}
	return fmt.Sprintf(`
resource "opnsense_routing_gateway" "test" {
  name        = %q
  interface   = "wan"
  gateway     = "198.51.100.1"
  far_gateway = true
}
`, gateway), "opnsense_routing_gateway.test.name"
}

func testAccRouteResourceConfig(gateway, network string) string {
	pre, ref := routeGateway(gateway)
	return pre + fmt.Sprintf(`
resource "opnsense_route" "test" {
  gateway = %[1]s
  network = %[2]q
}
`, ref, network)
}

func testAccRouteResourceConfigDisabled(gateway, network string) string {
	pre, ref := routeGateway(gateway)
	return pre + fmt.Sprintf(`
resource "opnsense_route" "test" {
  enabled = false
  gateway = %[1]s
  network = %[2]q
}
`, ref, network)
}

func testAccRouteResourceConfigWithDescription(gateway, network, description string) string {
	pre, ref := routeGateway(gateway)
	return pre + fmt.Sprintf(`
resource "opnsense_route" "test" {
  gateway     = %[1]s
  network     = %[2]q
  description = %[3]q
}
`, ref, network, description)
}

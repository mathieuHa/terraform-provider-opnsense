package trafficshaper_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTrafficShaperPipeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPipeConfig(50, `scheduler = "fq_codel"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "bandwidth", "50"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "bandwidth_metric", "Mbit"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "scheduler", "fq_codel"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "mask", "none"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "enabled", "true"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "queue", "-1"),
					resource.TestCheckResourceAttrSet("opnsense_trafficshaper_pipe.test", "number"),
					resource.TestCheckResourceAttrSet("opnsense_trafficshaper_pipe.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_trafficshaper_pipe.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPipeConfig(80, `mask = "dst-ip"`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "bandwidth", "80"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_pipe.test", "mask", "dst-ip"),
					resource.TestCheckNoResourceAttr("opnsense_trafficshaper_pipe.test", "scheduler"),
				),
			},
		},
	})
}

func TestAccTrafficShaperQueueAndRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccShaperChainConfig(30, `["192.0.2.0/24"]`, `[]`, "opnsense_trafficshaper_queue.test.id"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("opnsense_trafficshaper_queue.test", "pipe", "opnsense_trafficshaper_pipe.test", "id"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_queue.test", "weight", "30"),
					resource.TestCheckResourceAttrSet("opnsense_trafficshaper_queue.test", "number"),
					resource.TestCheckResourceAttrPair("opnsense_trafficshaper_rule.test", "target", "opnsense_trafficshaper_queue.test", "id"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "interface", "wan"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "protocol", "ip"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "source.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_trafficshaper_rule.test", "source.*", "any"),
					resource.TestCheckTypeSetElemAttr("opnsense_trafficshaper_rule.test", "destination.*", "192.0.2.0/24"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "destination_port", "any"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "dscp.#", "0"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "direction", "out"),
				),
			},
			{
				ResourceName:      "opnsense_trafficshaper_queue.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "opnsense_trafficshaper_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccShaperChainConfig(60, `["192.0.2.0/24", "198.51.100.0/24"]`, `["ef", "af41"]`, "opnsense_trafficshaper_pipe.test.id"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_trafficshaper_queue.test", "weight", "60"),
					resource.TestCheckResourceAttrPair("opnsense_trafficshaper_rule.test", "target", "opnsense_trafficshaper_pipe.test", "id"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "destination.#", "2"),
					resource.TestCheckResourceAttr("opnsense_trafficshaper_rule.test", "dscp.#", "2"),
					resource.TestCheckTypeSetElemAttr("opnsense_trafficshaper_rule.test", "dscp.*", "ef"),
				),
			},
		},
	})
}

func TestAccTrafficShaperDataSources(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccShaperChainConfig(30, `["192.0.2.0/24"]`, `[]`, "opnsense_trafficshaper_queue.test.id") + `
data "opnsense_trafficshaper_pipe" "test" {
  id = opnsense_trafficshaper_pipe.test.id
}

data "opnsense_trafficshaper_queue" "test" {
  id = opnsense_trafficshaper_queue.test.id
}

data "opnsense_trafficshaper_rule" "test" {
  id = opnsense_trafficshaper_rule.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.opnsense_trafficshaper_pipe.test", "number", "opnsense_trafficshaper_pipe.test", "number"),
					resource.TestCheckResourceAttr("data.opnsense_trafficshaper_pipe.test", "bandwidth", "100"),
					resource.TestCheckResourceAttr("data.opnsense_trafficshaper_queue.test", "weight", "30"),
					resource.TestCheckResourceAttrPair("data.opnsense_trafficshaper_rule.test", "target", "opnsense_trafficshaper_queue.test", "id"),
				),
			},
		},
	})
}

func testAccPipeConfig(bandwidth int, extra string) string {
	return fmt.Sprintf(`
resource "opnsense_trafficshaper_pipe" "test" {
  description      = "tf-test-pipe"
  bandwidth        = %d
  bandwidth_metric = "Mbit"
  %s
}
`, bandwidth, extra)
}

func testAccShaperChainConfig(weight int, destination, dscp, target string) string {
	return fmt.Sprintf(`
resource "opnsense_trafficshaper_pipe" "test" {
  description      = "tf-test-chain-pipe"
  bandwidth        = 100
  bandwidth_metric = "Mbit"
  scheduler        = "fq_codel"
}

resource "opnsense_trafficshaper_queue" "test" {
  description = "tf-test-chain-queue"
  pipe        = opnsense_trafficshaper_pipe.test.id
  weight      = %d
}

resource "opnsense_trafficshaper_rule" "test" {
  description = "tf-test-chain-rule"
  sequence    = 10
  interface   = "wan"
  destination = %s
  dscp        = %s
  direction   = "out"
  target      = %s
}
`, weight, destination, dscp, target)
}

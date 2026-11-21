package core_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccFirmwareSettingsResource(t *testing.T) {
	acctest.AccPreCheck(t)
	cur, err := opnsense.NewClient(acctest.Client()).Core().FirmwareSettingsGet(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	mirror := cur.Monad.Mirror
	a := "opnsense_firmware_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             `resource "opnsense_firmware_settings" "test" {}`,
				ResourceName:       a,
				ImportState:        true,
				ImportStateId:      "firmware_settings",
				ImportStatePersist: true,
			},
			{
				Config: `
resource "opnsense_firmware_settings" "test" {
  mirror = "https://pkg.opnsense.org"
}
`,
				Check: resource.TestCheckResourceAttr(a, "mirror", "https://pkg.opnsense.org"),
			},
			{
				Config: fmt.Sprintf(`
resource "opnsense_firmware_settings" "test" {
  mirror = %q
}
`, mirror),
				Check: resource.TestCheckResourceAttr(a, "mirror", mirror),
			},
		},
	})
}

func TestAccFirmwarePluginResource(t *testing.T) {
	pkg := os.Getenv("OPNSENSE_TEST_FIRMWARE_PLUGIN")
	if pkg == "" {
		t.Skip("set OPNSENSE_TEST_FIRMWARE_PLUGIN to a plugin that is not installed (e.g. os-iperf); the firewall must reach its mirror")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			info, err := opnsense.NewClient(acctest.Client()).Core().FirmwareInfo(context.Background())
			if err != nil {
				return err
			}
			for _, p := range info.Packages {
				if p.Name == pkg && p.Installed == "1" {
					return fmt.Errorf("%s still installed", pkg)
				}
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "opnsense_firmware_plugin" "test" { name = %q }`, pkg),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_firmware_plugin.test", "id", pkg),
					resource.TestCheckResourceAttrSet("opnsense_firmware_plugin.test", "version"),
				),
			},
			{
				ResourceName:      "opnsense_firmware_plugin.test",
				ImportState:       true,
				ImportStateId:     pkg,
				ImportStateVerify: true,
			},
		},
	})
}

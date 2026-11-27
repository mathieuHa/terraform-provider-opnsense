package trust_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccTrustKeepMaterialConfig(suffix string) string {
	return fmt.Sprintf(`
resource "opnsense_trust_ca" "test" {
  description = "tfacc keep material CA%[1]s"
  common_name = "keep-material-ca.example.com"
  country     = "NL"
}

resource "opnsense_trust_cert" "test" {
  description = "tfacc keep material cert%[1]s"
  caref       = opnsense_trust_ca.test.ref_id
  cert_type   = "server_cert"
  common_name = "keep-material.example.com"
  country     = "NL"
}
`, suffix)
}

type material struct{ caCrt, caPrv, crt, prv string }

func readMaterial(s *terraform.State, m *material) error {
	c := opnsense.NewClient(acctest.Client())
	ca, err := c.Trust().GetCa(context.Background(), s.RootModule().Resources["opnsense_trust_ca.test"].Primary.ID)
	if err != nil {
		return err
	}
	cert, err := c.Trust().GetCert(context.Background(), s.RootModule().Resources["opnsense_trust_cert.test"].Primary.ID)
	if err != nil {
		return err
	}
	*m = material{ca.Crt, ca.Prv, cert.Crt, cert.Prv}
	return nil
}

func TestAccTrustUpdateKeepsKeyMaterial(t *testing.T) {
	var before, after material
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrustKeepMaterialConfig(""),
				Check:  func(s *terraform.State) error { return readMaterial(s, &before) },
			},
			{
				Config: testAccTrustKeepMaterialConfig(" updated"),
				Check: func(s *terraform.State) error {
					if err := readMaterial(s, &after); err != nil {
						return err
					}
					if after != before {
						return fmt.Errorf("updating the descriptions regenerated the CA or the certificate")
					}
					return nil
				},
			},
			{
				ResourceName:    "opnsense_trust_ca.test",
				Config:          testAccTrustKeepMaterialConfig(" updated"),
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				ImportPlanChecks: resource.ImportPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("opnsense_trust_ca.test", plancheck.ResourceActionNoop)},
				},
			},
			{
				ResourceName:    "opnsense_trust_cert.test",
				Config:          testAccTrustKeepMaterialConfig(" updated"),
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				ImportPlanChecks: resource.ImportPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction("opnsense_trust_cert.test", plancheck.ResourceActionNoop)},
				},
			},
		},
	})
}

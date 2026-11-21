package trust_test

import (
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTrustResource_DefaultCountryAndImport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "opnsense_trust_ca" "root" {
  description = "tf-acc-default-country"
  common_name = "tf-acc-default-country"
}

resource "opnsense_trust_ca" "imported" {
  description = "tf-acc-imported"
  action      = "existing"
  crt         = opnsense_trust_ca.root.crt
}

resource "opnsense_trust_cert" "leaf" {
  description = "tf-acc-leaf"
  caref       = opnsense_trust_ca.root.ref_id
  common_name = "tf-acc-leaf"
}

resource "opnsense_trust_cert" "imported" {
  description = "tf-acc-leaf-imported"
  action      = "import"
  crt         = opnsense_trust_cert.leaf.crt
  prv         = opnsense_trust_cert.leaf.prv
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_trust_ca.root", "country", "NL"),
					resource.TestCheckResourceAttrSet("opnsense_trust_ca.imported", "ref_id"),
					resource.TestCheckResourceAttrPair("opnsense_trust_ca.imported", "crt", "opnsense_trust_ca.root", "crt"),
					resource.TestCheckResourceAttr("opnsense_trust_ca.imported", "common_name", "tf-acc-default-country"),
					resource.TestCheckResourceAttrSet("opnsense_trust_cert.imported", "ref_id"),
					resource.TestCheckResourceAttrPair("opnsense_trust_cert.imported", "crt", "opnsense_trust_cert.leaf", "crt"),
				),
			},
		},
	})
}

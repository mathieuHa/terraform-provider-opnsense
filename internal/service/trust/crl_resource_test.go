package trust_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const crlConfig = `
resource "opnsense_trust_ca" "crl" {
  description = "tfacc CRL CA"
  action      = "internal"
  key_type    = "2048"
  digest      = "sha256"
  lifetime    = "3650"
  country     = "US"
  common_name = "tfacc-crl-ca.example.com"
}

resource "opnsense_trust_cert" "laptop" {
  description = "tfacc laptop"
  action      = "internal"
  caref       = opnsense_trust_ca.crl.ref_id
  key_type    = "2048"
  digest      = "sha256"
  cert_type   = "usr_cert"
  lifetime    = "397"
  country     = "US"
  common_name = "tfacc-laptop"
}

resource "opnsense_trust_cert" "phone" {
  description = "tfacc phone"
  action      = "internal"
  caref       = opnsense_trust_ca.crl.ref_id
  key_type    = "2048"
  digest      = "sha256"
  cert_type   = "usr_cert"
  lifetime    = "397"
  country     = "US"
  common_name = "tfacc-phone"
}

resource "opnsense_trust_crl" "test" {
  ca          = opnsense_trust_ca.crl.ref_id
  description = "tfacc CRL"
  lifetime    = 30
  revoked     = [%s]
}
`

func TestAccTrustCrlResource(t *testing.T) {
	c := "opnsense_trust_crl.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(crlConfig, `{ certificate = opnsense_trust_cert.laptop.ref_id, reason = "keyCompromise" }`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(c, "id", "opnsense_trust_ca.crl", "ref_id"),
					resource.TestCheckResourceAttr(c, "method", "internal"),
					resource.TestCheckResourceAttr(c, "revoked.#", "1"),
					resource.TestCheckResourceAttr(c, "revoked.0.reason", "keyCompromise"),
					resource.TestMatchResourceAttr(c, "pem", regexp.MustCompile(`BEGIN X509 CRL`)),
				),
			},
			{
				ResourceName:            c,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"import_pem"},
			},
			{
				Config: fmt.Sprintf(crlConfig, `{ certificate = opnsense_trust_cert.laptop.ref_id, reason = "keyCompromise" }, { certificate = opnsense_trust_cert.phone.ref_id, reason = "superseded" }`),
				Check:  resource.TestCheckResourceAttr(c, "revoked.#", "2"),
			},
			{
				Config: fmt.Sprintf(crlConfig, ""),
				Check:  resource.TestCheckResourceAttr(c, "revoked.#", "0"),
			},
		},
	})
}

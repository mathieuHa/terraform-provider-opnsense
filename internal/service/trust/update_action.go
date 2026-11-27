package trust

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/trust"
)

func keepCaMaterial(ca *trust.Ca) {
	ca.Action = api.SelectedMap("existing")
	if ca.CrtPayload == "" {
		ca.CrtPayload = pemText(ca.Crt)
		ca.PrvPayload = pemText(ca.Prv)
	}
	ca.Crt = ""
	ca.Prv = ""
}

func keepCertMaterial(cert *trust.Cert) {
	cert.Action = api.SelectedMap("manual")
}

package kea

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func joinStaticRoutes(pieces []string) string {
	return strings.Join(pieces, ",")
}

func parseStaticRoutes(s string) []staticRouteModel {
	fields := strings.FieldsFunc(s, func(r rune) bool { return r == ',' || r == ';' })
	routes := []staticRouteModel{}
	for i := 0; i+1 < len(fields); i += 2 {
		routes = append(routes, staticRouteModel{
			DestinationIp: types.StringValue(strings.TrimSpace(fields[i])),
			RouterIp:      types.StringValue(strings.TrimSpace(fields[i+1])),
		})
	}
	return routes
}

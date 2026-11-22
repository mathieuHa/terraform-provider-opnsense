package dnsmasq

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHostResource,
		newBootResource,
		newDomainResource,
		newOptionResource,
		newRangeResource,
		newTagResource,
		newSettingsResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newClientDataSource,
		newBootDataSource,
		newDomainDataSource,
		newOptionDataSource,
		newRangeDataSource,
		newTagDataSource,
	}
}

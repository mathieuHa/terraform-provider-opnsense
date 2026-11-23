package bind

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newSettingsResource,
		newDnsblSettingsResource,
		newAclResource,
		newPrimaryDomainResource,
		newSecondaryDomainResource,
		newForwardDomainResource,
		newRecordResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newAclDataSource,
		newPrimaryDomainDataSource,
		newSecondaryDomainDataSource,
		newForwardDomainDataSource,
		newRecordDataSource,
	}
}

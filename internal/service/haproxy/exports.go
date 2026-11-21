package haproxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newSettingsResource,
		newFrontendResource,
		newBackendResource,
		newServerResource,
		newHealthcheckResource,
		newAclResource,
		newActionResource,
		newLuaResource,
		newFcgiResource,
		newErrorfileResource,
		newMapfileResource,
		newCpuResource,
		newGroupResource,
		newUserResource,
		newResolverResource,
		newMailerResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newFrontendDataSource,
		newBackendDataSource,
		newServerDataSource,
		newHealthcheckDataSource,
		newAclDataSource,
		newActionDataSource,
		newLuaDataSource,
		newFcgiDataSource,
		newErrorfileDataSource,
		newMapfileDataSource,
		newCpuDataSource,
		newGroupDataSource,
		newUserDataSource,
		newResolverDataSource,
		newMailerDataSource,
	}
}

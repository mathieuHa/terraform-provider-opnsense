package caddy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newSettingsResource,
		newReverseProxyResource,
		newSubdomainResource,
		newHandleResource,
		newAccessListResource,
		newBasicAuthResource,
		newHeaderResource,
		newLayer4Resource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newReverseProxyDataSource,
		newSubdomainDataSource,
		newHandleDataSource,
		newAccessListDataSource,
		newBasicAuthDataSource,
		newHeaderDataSource,
		newLayer4DataSource,
	}
}

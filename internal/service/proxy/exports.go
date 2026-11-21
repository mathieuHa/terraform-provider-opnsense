package proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newSettingsResource,
		newRemoteBlacklistResource,
		newPacProxyResource,
		newPacMatchResource,
		newPacRuleResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newRemoteBlacklistDataSource,
		newPacProxyDataSource,
		newPacMatchDataSource,
		newPacRuleDataSource,
	}
}

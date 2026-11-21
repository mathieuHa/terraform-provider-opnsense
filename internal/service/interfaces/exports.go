package interfaces

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newBridgeResource,
		newVipResource,
		newVlanResource,
		newAssignmentResource,
		newLoopbackResource,
		newLaggResource,
		newGreResource,
		newGifResource,
		newVxlanResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newBridgeDataSource,
		newVipDataSource,
		newVlanDataSource,
		newLoopbackDataSource,
		newLaggDataSource,
		newGreDataSource,
		newGifDataSource,
		newVxlanDataSource,
		newOverviewInterfaceDataSource,
		newOverviewAllDataSource,
	}
}

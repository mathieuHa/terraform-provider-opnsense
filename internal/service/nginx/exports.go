package nginx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newSettingsResource,
		newUpstreamServerResource,
		newUpstreamResource,
		newLocationResource,
		newHttpServerResource,
		newStreamServerResource,
		newCredentialResource,
		newUserlistResource,
		newHttpRewriteResource,
		newNaxsiRuleResource,
		newCustomPolicyResource,
		newSecurityHeaderResource,
		newLimitZoneResource,
		newLimitRequestConnectionResource,
		newCachePathResource,
		newProxyCacheValidResource,
		newResolverResource,
		newSyslogTargetResource,
		newTlsFingerprintResource,
		newErrorpageResource,
		newIpAclResource,
		newSniHostnameMapResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newUpstreamServerDataSource,
		newUpstreamDataSource,
		newLocationDataSource,
		newHttpServerDataSource,
		newStreamServerDataSource,
		newCredentialDataSource,
		newUserlistDataSource,
		newHttpRewriteDataSource,
		newNaxsiRuleDataSource,
		newCustomPolicyDataSource,
		newSecurityHeaderDataSource,
		newLimitZoneDataSource,
		newLimitRequestConnectionDataSource,
		newCachePathDataSource,
		newProxyCacheValidDataSource,
		newResolverDataSource,
		newSyslogTargetDataSource,
		newTlsFingerprintDataSource,
		newErrorpageDataSource,
		newIpAclDataSource,
		newSniHostnameMapDataSource,
	}
}

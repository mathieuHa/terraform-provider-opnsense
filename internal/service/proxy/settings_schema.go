package proxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/proxy"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type settingsResourceModel struct {
	Enabled                      types.Bool   `tfsdk:"enabled"`
	ErrorPages                   types.String `tfsdk:"error_pages"`
	IcpPort                      types.Int64  `tfsdk:"icp_port"`
	AlternateDnsServers          types.Set    `tfsdk:"alternate_dns_servers"`
	ForwardedForHandling         types.String `tfsdk:"forwarded_for_handling"`
	UriWhitespaceHandling        types.String `tfsdk:"uri_whitespace_handling"`
	EnablePinger                 types.Bool   `tfsdk:"enable_pinger"`
	UseViaHeader                 types.Bool   `tfsdk:"use_via_header"`
	SuppressVersion              types.Bool   `tfsdk:"suppress_version"`
	ConnectTimeout               types.Int64  `tfsdk:"connect_timeout"`
	VisibleEmail                 types.String `tfsdk:"visible_email"`
	VisibleHostname              types.String `tfsdk:"visible_hostname"`
	AccessLog                    types.Bool   `tfsdk:"access_log"`
	StoreLog                     types.Bool   `tfsdk:"store_log"`
	LogTarget                    types.String `tfsdk:"log_target"`
	IgnoreLogAcl                 types.Set    `tfsdk:"ignore_log_acl"`
	CacheEnabled                 types.Bool   `tfsdk:"cache_enabled"`
	CacheDirectory               types.String `tfsdk:"cache_directory"`
	CacheMemory                  types.Int64  `tfsdk:"cache_memory"`
	CacheMaxObjectSize           types.Int64  `tfsdk:"cache_max_object_size"`
	CacheMaxObjectSizeInMemory   types.Int64  `tfsdk:"cache_max_object_size_in_memory"`
	CacheMemoryMode              types.String `tfsdk:"cache_memory_mode"`
	CacheSize                    types.Int64  `tfsdk:"cache_size"`
	CacheSwapTimeout             types.Int64  `tfsdk:"cache_swap_timeout"`
	CacheMaxSwapRate             types.Int64  `tfsdk:"cache_max_swap_rate"`
	CacheSlotSize                types.Int64  `tfsdk:"cache_slot_size"`
	CacheLinuxPackages           types.Bool   `tfsdk:"cache_linux_packages"`
	CacheWindowsUpdates          types.Bool   `tfsdk:"cache_windows_updates"`
	TrafficEnabled               types.Bool   `tfsdk:"traffic_enabled"`
	MaxDownloadSize              types.Int64  `tfsdk:"max_download_size"`
	MaxUploadSize                types.Int64  `tfsdk:"max_upload_size"`
	OverallBandwidth             types.Int64  `tfsdk:"overall_bandwidth"`
	PerHostBandwidth             types.Int64  `tfsdk:"per_host_bandwidth"`
	ParentProxyEnabled           types.Bool   `tfsdk:"parent_proxy_enabled"`
	ParentProxyHost              types.String `tfsdk:"parent_proxy_host"`
	ParentProxyAuth              types.Bool   `tfsdk:"parent_proxy_auth"`
	ParentProxyUser              types.String `tfsdk:"parent_proxy_user"`
	ParentProxyPassword          types.String `tfsdk:"parent_proxy_password"`
	ParentProxyPort              types.Int64  `tfsdk:"parent_proxy_port"`
	ParentProxyLocalDomains      types.Set    `tfsdk:"parent_proxy_local_domains"`
	ParentProxyLocalIps          types.Set    `tfsdk:"parent_proxy_local_ips"`
	Interfaces                   types.Set    `tfsdk:"interfaces"`
	Port                         types.Int64  `tfsdk:"port"`
	SslBumpPort                  types.Int64  `tfsdk:"ssl_bump_port"`
	SslBump                      types.Bool   `tfsdk:"ssl_bump"`
	SslUrlOnly                   types.Bool   `tfsdk:"ssl_url_only"`
	SslCertificate               types.String `tfsdk:"ssl_certificate"`
	SslNoBumpSites               types.Set    `tfsdk:"ssl_no_bump_sites"`
	Workers                      types.Int64  `tfsdk:"workers"`
	SslCrtdStorageMaxSize        types.Int64  `tfsdk:"ssl_crtd_storage_max_size"`
	SslCrtdChildren              types.Int64  `tfsdk:"ssl_crtd_children"`
	SnmpEnabled                  types.Bool   `tfsdk:"snmp_enabled"`
	SnmpPort                     types.Int64  `tfsdk:"snmp_port"`
	SnmpPassword                 types.String `tfsdk:"snmp_password"`
	FtpInterfaces                types.Set    `tfsdk:"ftp_interfaces"`
	FtpPort                      types.Int64  `tfsdk:"ftp_port"`
	FtpTransparentMode           types.Bool   `tfsdk:"ftp_transparent_mode"`
	AclForInterfaceSubnets       types.Bool   `tfsdk:"acl_for_interface_subnets"`
	TransparentMode              types.Bool   `tfsdk:"transparent_mode"`
	AclAllowedSubnets            types.Set    `tfsdk:"acl_allowed_subnets"`
	AclUnrestricted              types.Set    `tfsdk:"acl_unrestricted"`
	AclBannedHosts               types.Set    `tfsdk:"acl_banned_hosts"`
	AclAllowWhitelistBannedHosts types.Bool   `tfsdk:"acl_allow_whitelist_banned_hosts"`
	AclWhitelist                 types.Set    `tfsdk:"acl_whitelist"`
	AclBlacklist                 types.Set    `tfsdk:"acl_blacklist"`
	AclBrowser                   types.Set    `tfsdk:"acl_browser"`
	AclMimeType                  types.Set    `tfsdk:"acl_mime_type"`
	AclGoogleApps                types.String `tfsdk:"acl_google_apps"`
	AclYoutube                   types.String `tfsdk:"acl_youtube"`
	AclSafePorts                 types.Set    `tfsdk:"acl_safe_ports"`
	AclSslPorts                  types.Set    `tfsdk:"acl_ssl_ports"`
	IcapEnabled                  types.Bool   `tfsdk:"icap_enabled"`
	IcapRequestUrl               types.String `tfsdk:"icap_request_url"`
	IcapResponseUrl              types.String `tfsdk:"icap_response_url"`
	IcapSendClientIp             types.Bool   `tfsdk:"icap_send_client_ip"`
	IcapSendUsername             types.Bool   `tfsdk:"icap_send_username"`
	IcapEncodeUsername           types.Bool   `tfsdk:"icap_encode_username"`
	IcapUsernameHeader           types.String `tfsdk:"icap_username_header"`
	IcapEnablePreview            types.Bool   `tfsdk:"icap_enable_preview"`
	IcapPreviewSize              types.Int64  `tfsdk:"icap_preview_size"`
	IcapOptionsTtl               types.Int64  `tfsdk:"icap_options_ttl"`
	IcapExclude                  types.Set    `tfsdk:"icap_exclude"`
	AuthMethod                   types.Set    `tfsdk:"auth_method"`
	AuthEnforceGroup             types.Set    `tfsdk:"auth_enforce_group"`
	AuthRealm                    types.String `tfsdk:"auth_realm"`
	AuthCredentialsTtl           types.Int64  `tfsdk:"auth_credentials_ttl"`
	AuthChildren                 types.Int64  `tfsdk:"auth_children"`

	Id types.String `tfsdk:"id"`
}

func settingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the Squid web proxy.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"error_pages": schema.StringAttribute{
			MarkdownDescription: "Error page theme: `opnsense`, `squid` or `custom`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"icp_port": schema.Int64Attribute{
			MarkdownDescription: "ICP port, empty to disable ICP.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"alternate_dns_servers": schema.SetAttribute{
			MarkdownDescription: "DNS servers used by the proxy instead of the system resolvers.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"forwarded_for_handling": schema.StringAttribute{
			MarkdownDescription: "Handling of the X-Forwarded-For header: `on`, `off`, `transparent`, `truncate` or `delete`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"uri_whitespace_handling": schema.StringAttribute{
			MarkdownDescription: "Handling of whitespace in URIs: `strip`, `deny`, `allow`, `encode` or `chop`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"enable_pinger": schema.BoolAttribute{
			MarkdownDescription: "Enable the ICMP pinger used to measure the distance to origin servers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"use_via_header": schema.BoolAttribute{
			MarkdownDescription: "Add a Via header to forwarded requests.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"suppress_version": schema.BoolAttribute{
			MarkdownDescription: "Hide the Squid version in error pages and headers.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"connect_timeout": schema.Int64Attribute{
			MarkdownDescription: "Timeout in seconds for connections to origin servers.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"visible_email": schema.StringAttribute{
			MarkdownDescription: "Administrator e-mail shown on error pages.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"visible_hostname": schema.StringAttribute{
			MarkdownDescription: "Hostname shown on error pages.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"access_log": schema.BoolAttribute{
			MarkdownDescription: "Write the access log.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"store_log": schema.BoolAttribute{
			MarkdownDescription: "Write the store log.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"log_target": schema.StringAttribute{
			MarkdownDescription: "Where access logs go: empty for the local file, `syslog` or `file_extendend`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ignore_log_acl": schema.SetAttribute{
			MarkdownDescription: "Networks whose requests are not logged.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"cache_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the local disk cache.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cache_directory": schema.StringAttribute{
			MarkdownDescription: "Directory of the disk cache.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"cache_memory": schema.Int64Attribute{
			MarkdownDescription: "Memory cache size in MB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_max_object_size": schema.Int64Attribute{
			MarkdownDescription: "Largest object stored on disk, in MB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_max_object_size_in_memory": schema.Int64Attribute{
			MarkdownDescription: "Largest object kept in memory, in KB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_memory_mode": schema.StringAttribute{
			MarkdownDescription: "Which objects are kept in memory: `always`, `disk` or `network`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"cache_size": schema.Int64Attribute{
			MarkdownDescription: "Disk cache size in MB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_swap_timeout": schema.Int64Attribute{
			MarkdownDescription: "Swap timeout of the rock store, in milliseconds.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_max_swap_rate": schema.Int64Attribute{
			MarkdownDescription: "Maximum swap operations per second of the rock store.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_slot_size": schema.Int64Attribute{
			MarkdownDescription: "Slot size of the rock store, in bytes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"cache_linux_packages": schema.BoolAttribute{
			MarkdownDescription: "Cache Linux package downloads.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"cache_windows_updates": schema.BoolAttribute{
			MarkdownDescription: "Cache Windows updates.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"traffic_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable traffic management.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"max_download_size": schema.Int64Attribute{
			MarkdownDescription: "Maximum download size in KB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"max_upload_size": schema.Int64Attribute{
			MarkdownDescription: "Maximum upload size in KB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"overall_bandwidth": schema.Int64Attribute{
			MarkdownDescription: "Overall bandwidth limit in Kbit/s.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"per_host_bandwidth": schema.Int64Attribute{
			MarkdownDescription: "Per-host bandwidth limit in Kbit/s.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"parent_proxy_enabled": schema.BoolAttribute{
			MarkdownDescription: "Forward requests to a parent proxy.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_host": schema.StringAttribute{
			MarkdownDescription: "Parent proxy host.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_auth": schema.BoolAttribute{
			MarkdownDescription: "Authenticate to the parent proxy.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_user": schema.StringAttribute{
			MarkdownDescription: "Parent proxy user name.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_password": schema.StringAttribute{
			MarkdownDescription: "Parent proxy password.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_port": schema.Int64Attribute{
			MarkdownDescription: "Parent proxy port.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"parent_proxy_local_domains": schema.SetAttribute{
			MarkdownDescription: "Domains reached directly instead of through the parent proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"parent_proxy_local_ips": schema.SetAttribute{
			MarkdownDescription: "Addresses reached directly instead of through the parent proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"interfaces": schema.SetAttribute{
			MarkdownDescription: "Interfaces the proxy listens on, e.g. `[\"lan\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "Proxy port.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ssl_bump_port": schema.Int64Attribute{
			MarkdownDescription: "Port of the transparent HTTPS proxy.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ssl_bump": schema.BoolAttribute{
			MarkdownDescription: "Enable SSL inspection.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ssl_url_only": schema.BoolAttribute{
			MarkdownDescription: "Only log and filter on the SNI host name, without decrypting.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"ssl_certificate": schema.StringAttribute{
			MarkdownDescription: "Reference ID of the CA used for SSL inspection.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ssl_no_bump_sites": schema.SetAttribute{
			MarkdownDescription: "Sites never inspected.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"workers": schema.Int64Attribute{
			MarkdownDescription: "Number of Squid worker processes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ssl_crtd_storage_max_size": schema.Int64Attribute{
			MarkdownDescription: "Size of the generated certificate cache, in MB.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ssl_crtd_children": schema.Int64Attribute{
			MarkdownDescription: "Number of certificate generator processes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"snmp_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable the SNMP agent.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"snmp_port": schema.Int64Attribute{
			MarkdownDescription: "SNMP agent port.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"snmp_password": schema.StringAttribute{
			MarkdownDescription: "SNMP community.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ftp_interfaces": schema.SetAttribute{
			MarkdownDescription: "Interfaces of the FTP proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"ftp_port": schema.Int64Attribute{
			MarkdownDescription: "FTP proxy port.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"ftp_transparent_mode": schema.BoolAttribute{
			MarkdownDescription: "Transparent FTP proxy.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"acl_for_interface_subnets": schema.BoolAttribute{
			MarkdownDescription: "Allow the subnets of the listening interfaces automatically.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"transparent_mode": schema.BoolAttribute{
			MarkdownDescription: "Enable the transparent HTTP proxy.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"acl_allowed_subnets": schema.SetAttribute{
			MarkdownDescription: "Additional subnets allowed to use the proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_unrestricted": schema.SetAttribute{
			MarkdownDescription: "Addresses that bypass all ACLs.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_banned_hosts": schema.SetAttribute{
			MarkdownDescription: "Addresses denied access to the proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_allow_whitelist_banned_hosts": schema.BoolAttribute{
			MarkdownDescription: "Let banned hosts still reach whitelisted sites.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"acl_whitelist": schema.SetAttribute{
			MarkdownDescription: "Sites always allowed (regular expressions).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_blacklist": schema.SetAttribute{
			MarkdownDescription: "Sites always denied (regular expressions).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_browser": schema.SetAttribute{
			MarkdownDescription: "Denied user agents (regular expressions).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_mime_type": schema.SetAttribute{
			MarkdownDescription: "Denied MIME types.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_google_apps": schema.StringAttribute{
			MarkdownDescription: "Domains allowed to use Google services (X-GoogApps-Allowed-Domains).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"acl_youtube": schema.StringAttribute{
			MarkdownDescription: "YouTube restriction: `strict` or `moderate`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"acl_safe_ports": schema.SetAttribute{
			MarkdownDescription: "Allowed destination ports, as `port:description` or `from-to:description`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"acl_ssl_ports": schema.SetAttribute{
			MarkdownDescription: "Allowed CONNECT ports, as `port:description`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"icap_enabled": schema.BoolAttribute{
			MarkdownDescription: "Enable ICAP.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"icap_request_url": schema.StringAttribute{
			MarkdownDescription: "ICAP request modification URL.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"icap_response_url": schema.StringAttribute{
			MarkdownDescription: "ICAP response modification URL.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"icap_send_client_ip": schema.BoolAttribute{
			MarkdownDescription: "Send the client address to the ICAP server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"icap_send_username": schema.BoolAttribute{
			MarkdownDescription: "Send the user name to the ICAP server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"icap_encode_username": schema.BoolAttribute{
			MarkdownDescription: "Base64-encode the user name sent to the ICAP server.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"icap_username_header": schema.StringAttribute{
			MarkdownDescription: "Header carrying the user name.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"icap_enable_preview": schema.BoolAttribute{
			MarkdownDescription: "Enable ICAP preview.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"icap_preview_size": schema.Int64Attribute{
			MarkdownDescription: "ICAP preview size in bytes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"icap_options_ttl": schema.Int64Attribute{
			MarkdownDescription: "TTL of the ICAP OPTIONS response, in seconds.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"icap_exclude": schema.SetAttribute{
			MarkdownDescription: "Domains not sent to the ICAP server.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"auth_method": schema.SetAttribute{
			MarkdownDescription: "Authentication servers, e.g. `[\"Local Database\"]`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"auth_enforce_group": schema.SetAttribute{
			MarkdownDescription: "Only members of these local groups may use the proxy.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"auth_realm": schema.StringAttribute{
			MarkdownDescription: "Authentication realm.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"auth_credentials_ttl": schema.Int64Attribute{
			MarkdownDescription: "Credential cache lifetime, in minutes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"auth_children": schema.Int64Attribute{
			MarkdownDescription: "Number of authentication processes.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(0)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `proxy_settings`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func settingsBool(b types.Bool) string {
	if b.IsNull() || b.IsUnknown() {
		return ""
	}
	return tools.BoolToString(b.ValueBool())
}

func settingsInt(i types.Int64) string {
	if i.IsNull() || i.IsUnknown() {
		return ""
	}
	return tools.Int64ToString(i.ValueInt64())
}

func settingsList(s types.Set) []string {
	out := tools.SetToStringSlice(s)
	if out == nil {
		return []string{}
	}
	return out
}

func settingsNonEmpty(in []string) []string {
	out := []string{}
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func convertSettingsSchemaToStruct(d *settingsResourceModel) *proxy.Settings {
	return &proxy.Settings{
		General: proxy.SettingsGeneral{
			Enabled:               settingsBool(d.Enabled),
			ErrorPages:            api.SelectedMap(d.ErrorPages.ValueString()),
			IcpPort:               settingsInt(d.IcpPort),
			AlternateDnsServers:   api.SelectedMapList(settingsList(d.AlternateDnsServers)),
			ForwardedForHandling:  api.SelectedMap(d.ForwardedForHandling.ValueString()),
			UriWhitespaceHandling: api.SelectedMap(d.UriWhitespaceHandling.ValueString()),
			EnablePinger:          settingsBool(d.EnablePinger),
			UseViaHeader:          settingsBool(d.UseViaHeader),
			SuppressVersion:       settingsBool(d.SuppressVersion),
			ConnectTimeout:        settingsInt(d.ConnectTimeout),
			VisibleEmail:          d.VisibleEmail.ValueString(),
			VisibleHostname:       d.VisibleHostname.ValueString(),
			Logging: proxy.SettingsGeneralLogging{
				LogTarget:    api.SelectedMap(d.LogTarget.ValueString()),
				IgnoreLogAcl: api.SelectedMapList(settingsList(d.IgnoreLogAcl)),
				Enable: proxy.SettingsGeneralLoggingEnable{
					AccessLog: settingsBool(d.AccessLog),
					StoreLog:  settingsBool(d.StoreLog),
				},
			},
			Cache: proxy.SettingsGeneralCache{
				Local: proxy.SettingsGeneralCacheLocal{
					CacheEnabled:               settingsBool(d.CacheEnabled),
					CacheDirectory:             d.CacheDirectory.ValueString(),
					CacheMemory:                settingsInt(d.CacheMemory),
					CacheMaxObjectSize:         settingsInt(d.CacheMaxObjectSize),
					CacheMaxObjectSizeInMemory: settingsInt(d.CacheMaxObjectSizeInMemory),
					CacheMemoryMode:            api.SelectedMap(d.CacheMemoryMode.ValueString()),
					CacheSize:                  settingsInt(d.CacheSize),
					CacheSwapTimeout:           settingsInt(d.CacheSwapTimeout),
					CacheMaxSwapRate:           settingsInt(d.CacheMaxSwapRate),
					CacheSlotSize:              settingsInt(d.CacheSlotSize),
					CacheLinuxPackages:         settingsBool(d.CacheLinuxPackages),
					CacheWindowsUpdates:        settingsBool(d.CacheWindowsUpdates),
				},
			},
			Traffic: proxy.SettingsGeneralTraffic{
				TrafficEnabled:   settingsBool(d.TrafficEnabled),
				MaxDownloadSize:  settingsInt(d.MaxDownloadSize),
				MaxUploadSize:    settingsInt(d.MaxUploadSize),
				OverallBandwidth: settingsInt(d.OverallBandwidth),
				PerHostBandwidth: settingsInt(d.PerHostBandwidth),
			},
			Parentproxy: proxy.SettingsGeneralParentproxy{
				ParentProxyEnabled:      settingsBool(d.ParentProxyEnabled),
				ParentProxyHost:         d.ParentProxyHost.ValueString(),
				ParentProxyAuth:         settingsBool(d.ParentProxyAuth),
				ParentProxyUser:         d.ParentProxyUser.ValueString(),
				ParentProxyPassword:     d.ParentProxyPassword.ValueString(),
				ParentProxyPort:         settingsInt(d.ParentProxyPort),
				ParentProxyLocalDomains: api.SelectedMapList(settingsList(d.ParentProxyLocalDomains)),
				ParentProxyLocalIps:     api.SelectedMapList(settingsList(d.ParentProxyLocalIps)),
			},
		},
		Forward: proxy.SettingsForward{
			Interfaces:             api.SelectedMapList(settingsList(d.Interfaces)),
			Port:                   settingsInt(d.Port),
			SslBumpPort:            settingsInt(d.SslBumpPort),
			SslBump:                settingsBool(d.SslBump),
			SslUrlOnly:             settingsBool(d.SslUrlOnly),
			SslCertificate:         api.SelectedMap(d.SslCertificate.ValueString()),
			SslNoBumpSites:         api.SelectedMapList(settingsList(d.SslNoBumpSites)),
			Workers:                settingsInt(d.Workers),
			SslCrtdStorageMaxSize:  settingsInt(d.SslCrtdStorageMaxSize),
			SslCrtdChildren:        settingsInt(d.SslCrtdChildren),
			SnmpEnabled:            settingsBool(d.SnmpEnabled),
			SnmpPort:               settingsInt(d.SnmpPort),
			SnmpPassword:           d.SnmpPassword.ValueString(),
			FtpInterfaces:          api.SelectedMapList(settingsList(d.FtpInterfaces)),
			FtpPort:                settingsInt(d.FtpPort),
			FtpTransparentMode:     settingsBool(d.FtpTransparentMode),
			AclForInterfaceSubnets: settingsBool(d.AclForInterfaceSubnets),
			TransparentMode:        settingsBool(d.TransparentMode),
			Acl: proxy.SettingsForwardAcl{
				AclAllowedSubnets:            api.SelectedMapList(settingsList(d.AclAllowedSubnets)),
				AclUnrestricted:              api.SelectedMapList(settingsList(d.AclUnrestricted)),
				AclBannedHosts:               api.SelectedMapList(settingsList(d.AclBannedHosts)),
				AclAllowWhitelistBannedHosts: settingsBool(d.AclAllowWhitelistBannedHosts),
				AclWhitelist:                 api.SelectedMapList(settingsList(d.AclWhitelist)),
				AclBlacklist:                 api.SelectedMapList(settingsList(d.AclBlacklist)),
				AclBrowser:                   api.SelectedMapList(settingsList(d.AclBrowser)),
				AclMimeType:                  api.SelectedMapList(settingsList(d.AclMimeType)),
				AclGoogleApps:                d.AclGoogleApps.ValueString(),
				AclYoutube:                   api.SelectedMap(d.AclYoutube.ValueString()),
				AclSafePorts:                 api.SelectedMapList(settingsList(d.AclSafePorts)),
				AclSslPorts:                  api.SelectedMapList(settingsList(d.AclSslPorts)),
			},
			Icap: proxy.SettingsForwardIcap{
				IcapEnabled:        settingsBool(d.IcapEnabled),
				IcapRequestUrl:     d.IcapRequestUrl.ValueString(),
				IcapResponseUrl:    d.IcapResponseUrl.ValueString(),
				IcapSendClientIp:   settingsBool(d.IcapSendClientIp),
				IcapSendUsername:   settingsBool(d.IcapSendUsername),
				IcapEncodeUsername: settingsBool(d.IcapEncodeUsername),
				IcapUsernameHeader: d.IcapUsernameHeader.ValueString(),
				IcapEnablePreview:  settingsBool(d.IcapEnablePreview),
				IcapPreviewSize:    settingsInt(d.IcapPreviewSize),
				IcapOptionsTtl:     settingsInt(d.IcapOptionsTtl),
				IcapExclude:        api.SelectedMapList(settingsList(d.IcapExclude)),
			},
			Authentication: proxy.SettingsForwardAuthentication{
				AuthMethod:         api.SelectedMapList(settingsList(d.AuthMethod)),
				AuthEnforceGroup:   api.SelectedMapList(settingsList(d.AuthEnforceGroup)),
				AuthRealm:          d.AuthRealm.ValueString(),
				AuthCredentialsTtl: settingsInt(d.AuthCredentialsTtl),
				AuthChildren:       settingsInt(d.AuthChildren),
			},
		},
	}
}

func convertSettingsStructToSchema(s *proxy.Settings) *settingsResourceModel {
	return &settingsResourceModel{
		Enabled:                      types.BoolValue(tools.StringToBool(s.General.Enabled)),
		ErrorPages:                   types.StringValue(s.General.ErrorPages.String()),
		IcpPort:                      tools.StringToInt64Null(s.General.IcpPort),
		AlternateDnsServers:          tools.StringSliceToSet(settingsNonEmpty(s.General.AlternateDnsServers)),
		ForwardedForHandling:         types.StringValue(s.General.ForwardedForHandling.String()),
		UriWhitespaceHandling:        types.StringValue(s.General.UriWhitespaceHandling.String()),
		EnablePinger:                 types.BoolValue(tools.StringToBool(s.General.EnablePinger)),
		UseViaHeader:                 types.BoolValue(tools.StringToBool(s.General.UseViaHeader)),
		SuppressVersion:              types.BoolValue(tools.StringToBool(s.General.SuppressVersion)),
		ConnectTimeout:               tools.StringToInt64Null(s.General.ConnectTimeout),
		VisibleEmail:                 types.StringValue(s.General.VisibleEmail),
		VisibleHostname:              types.StringValue(s.General.VisibleHostname),
		AccessLog:                    types.BoolValue(tools.StringToBool(s.General.Logging.Enable.AccessLog)),
		StoreLog:                     types.BoolValue(tools.StringToBool(s.General.Logging.Enable.StoreLog)),
		LogTarget:                    types.StringValue(s.General.Logging.LogTarget.String()),
		IgnoreLogAcl:                 tools.StringSliceToSet(settingsNonEmpty(s.General.Logging.IgnoreLogAcl)),
		CacheEnabled:                 types.BoolValue(tools.StringToBool(s.General.Cache.Local.CacheEnabled)),
		CacheDirectory:               types.StringValue(s.General.Cache.Local.CacheDirectory),
		CacheMemory:                  tools.StringToInt64Null(s.General.Cache.Local.CacheMemory),
		CacheMaxObjectSize:           tools.StringToInt64Null(s.General.Cache.Local.CacheMaxObjectSize),
		CacheMaxObjectSizeInMemory:   tools.StringToInt64Null(s.General.Cache.Local.CacheMaxObjectSizeInMemory),
		CacheMemoryMode:              types.StringValue(s.General.Cache.Local.CacheMemoryMode.String()),
		CacheSize:                    tools.StringToInt64Null(s.General.Cache.Local.CacheSize),
		CacheSwapTimeout:             tools.StringToInt64Null(s.General.Cache.Local.CacheSwapTimeout),
		CacheMaxSwapRate:             tools.StringToInt64Null(s.General.Cache.Local.CacheMaxSwapRate),
		CacheSlotSize:                tools.StringToInt64Null(s.General.Cache.Local.CacheSlotSize),
		CacheLinuxPackages:           types.BoolValue(tools.StringToBool(s.General.Cache.Local.CacheLinuxPackages)),
		CacheWindowsUpdates:          types.BoolValue(tools.StringToBool(s.General.Cache.Local.CacheWindowsUpdates)),
		TrafficEnabled:               types.BoolValue(tools.StringToBool(s.General.Traffic.TrafficEnabled)),
		MaxDownloadSize:              tools.StringToInt64Null(s.General.Traffic.MaxDownloadSize),
		MaxUploadSize:                tools.StringToInt64Null(s.General.Traffic.MaxUploadSize),
		OverallBandwidth:             tools.StringToInt64Null(s.General.Traffic.OverallBandwidth),
		PerHostBandwidth:             tools.StringToInt64Null(s.General.Traffic.PerHostBandwidth),
		ParentProxyEnabled:           types.BoolValue(tools.StringToBool(s.General.Parentproxy.ParentProxyEnabled)),
		ParentProxyHost:              types.StringValue(s.General.Parentproxy.ParentProxyHost),
		ParentProxyAuth:              types.BoolValue(tools.StringToBool(s.General.Parentproxy.ParentProxyAuth)),
		ParentProxyUser:              types.StringValue(s.General.Parentproxy.ParentProxyUser),
		ParentProxyPassword:          types.StringValue(s.General.Parentproxy.ParentProxyPassword),
		ParentProxyPort:              tools.StringToInt64Null(s.General.Parentproxy.ParentProxyPort),
		ParentProxyLocalDomains:      tools.StringSliceToSet(settingsNonEmpty(s.General.Parentproxy.ParentProxyLocalDomains)),
		ParentProxyLocalIps:          tools.StringSliceToSet(settingsNonEmpty(s.General.Parentproxy.ParentProxyLocalIps)),
		Interfaces:                   tools.StringSliceToSet(settingsNonEmpty(s.Forward.Interfaces)),
		Port:                         tools.StringToInt64Null(s.Forward.Port),
		SslBumpPort:                  tools.StringToInt64Null(s.Forward.SslBumpPort),
		SslBump:                      types.BoolValue(tools.StringToBool(s.Forward.SslBump)),
		SslUrlOnly:                   types.BoolValue(tools.StringToBool(s.Forward.SslUrlOnly)),
		SslCertificate:               types.StringValue(s.Forward.SslCertificate.String()),
		SslNoBumpSites:               tools.StringSliceToSet(settingsNonEmpty(s.Forward.SslNoBumpSites)),
		Workers:                      tools.StringToInt64Null(s.Forward.Workers),
		SslCrtdStorageMaxSize:        tools.StringToInt64Null(s.Forward.SslCrtdStorageMaxSize),
		SslCrtdChildren:              tools.StringToInt64Null(s.Forward.SslCrtdChildren),
		SnmpEnabled:                  types.BoolValue(tools.StringToBool(s.Forward.SnmpEnabled)),
		SnmpPort:                     tools.StringToInt64Null(s.Forward.SnmpPort),
		SnmpPassword:                 types.StringValue(s.Forward.SnmpPassword),
		FtpInterfaces:                tools.StringSliceToSet(settingsNonEmpty(s.Forward.FtpInterfaces)),
		FtpPort:                      tools.StringToInt64Null(s.Forward.FtpPort),
		FtpTransparentMode:           types.BoolValue(tools.StringToBool(s.Forward.FtpTransparentMode)),
		AclForInterfaceSubnets:       types.BoolValue(tools.StringToBool(s.Forward.AclForInterfaceSubnets)),
		TransparentMode:              types.BoolValue(tools.StringToBool(s.Forward.TransparentMode)),
		AclAllowedSubnets:            tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclAllowedSubnets)),
		AclUnrestricted:              tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclUnrestricted)),
		AclBannedHosts:               tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclBannedHosts)),
		AclAllowWhitelistBannedHosts: types.BoolValue(tools.StringToBool(s.Forward.Acl.AclAllowWhitelistBannedHosts)),
		AclWhitelist:                 tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclWhitelist)),
		AclBlacklist:                 tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclBlacklist)),
		AclBrowser:                   tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclBrowser)),
		AclMimeType:                  tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclMimeType)),
		AclGoogleApps:                types.StringValue(s.Forward.Acl.AclGoogleApps),
		AclYoutube:                   types.StringValue(s.Forward.Acl.AclYoutube.String()),
		AclSafePorts:                 tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclSafePorts)),
		AclSslPorts:                  tools.StringSliceToSet(settingsNonEmpty(s.Forward.Acl.AclSslPorts)),
		IcapEnabled:                  types.BoolValue(tools.StringToBool(s.Forward.Icap.IcapEnabled)),
		IcapRequestUrl:               types.StringValue(s.Forward.Icap.IcapRequestUrl),
		IcapResponseUrl:              types.StringValue(s.Forward.Icap.IcapResponseUrl),
		IcapSendClientIp:             types.BoolValue(tools.StringToBool(s.Forward.Icap.IcapSendClientIp)),
		IcapSendUsername:             types.BoolValue(tools.StringToBool(s.Forward.Icap.IcapSendUsername)),
		IcapEncodeUsername:           types.BoolValue(tools.StringToBool(s.Forward.Icap.IcapEncodeUsername)),
		IcapUsernameHeader:           types.StringValue(s.Forward.Icap.IcapUsernameHeader),
		IcapEnablePreview:            types.BoolValue(tools.StringToBool(s.Forward.Icap.IcapEnablePreview)),
		IcapPreviewSize:              tools.StringToInt64Null(s.Forward.Icap.IcapPreviewSize),
		IcapOptionsTtl:               tools.StringToInt64Null(s.Forward.Icap.IcapOptionsTtl),
		IcapExclude:                  tools.StringSliceToSet(settingsNonEmpty(s.Forward.Icap.IcapExclude)),
		AuthMethod:                   tools.StringSliceToSet(settingsNonEmpty(s.Forward.Authentication.AuthMethod)),
		AuthEnforceGroup:             tools.StringSliceToSet(settingsNonEmpty(s.Forward.Authentication.AuthEnforceGroup)),
		AuthRealm:                    types.StringValue(s.Forward.Authentication.AuthRealm),
		AuthCredentialsTtl:           tools.StringToInt64Null(s.Forward.Authentication.AuthCredentialsTtl),
		AuthChildren:                 tools.StringToInt64Null(s.Forward.Authentication.AuthChildren),
	}
}

package acmeclient

import (
	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type validationResourceModel struct {
	Enabled                     types.Bool   `tfsdk:"enabled"`
	Name                        types.String `tfsdk:"name"`
	Description                 types.String `tfsdk:"description"`
	Method                      types.String `tfsdk:"method"`
	HttpService                 types.String `tfsdk:"http_service"`
	HttpOpnAutodiscovery        types.Bool   `tfsdk:"http_opn_autodiscovery"`
	HttpOpnInterface            types.String `tfsdk:"http_opn_interface"`
	HttpOpnIpaddresses          types.Set    `tfsdk:"http_opn_ipaddresses"`
	HttpHaproxyInject           types.Bool   `tfsdk:"http_haproxy_inject"`
	HttpHaproxyFrontends        types.Set    `tfsdk:"http_haproxy_frontends"`
	TlsalpnService              types.String `tfsdk:"tlsalpn_service"`
	TlsalpnAcmeAutodiscovery    types.Bool   `tfsdk:"tlsalpn_acme_autodiscovery"`
	TlsalpnAcmeInterface        types.String `tfsdk:"tlsalpn_acme_interface"`
	TlsalpnAcmeIpaddresses      types.Set    `tfsdk:"tlsalpn_acme_ipaddresses"`
	DNSService                  types.String `tfsdk:"dns_service"`
	DNSSleep                    types.Int64  `tfsdk:"dns_sleep"`
	DNSActive24ApiKey           types.String `tfsdk:"dns_active24_api_key"`
	DNSActive24ApiSecret        types.String `tfsdk:"dns_active24_api_secret"`
	DNSAdKey                    types.String `tfsdk:"dns_ad_key"`
	DNSAliKey                   types.String `tfsdk:"dns_ali_key"`
	DNSAliSecret                types.String `tfsdk:"dns_ali_secret"`
	DNSAutodnsUser              types.String `tfsdk:"dns_autodns_user"`
	DNSAutodnsPassword          types.String `tfsdk:"dns_autodns_password"`
	DNSAutodnsContext           types.String `tfsdk:"dns_autodns_context"`
	DNSAwsID                    types.String `tfsdk:"dns_aws_id"`
	DNSAwsSecret                types.String `tfsdk:"dns_aws_secret"`
	DNSAzurednsSubscriptionid   types.String `tfsdk:"dns_azuredns_subscriptionid"`
	DNSAzurednsTenantid         types.String `tfsdk:"dns_azuredns_tenantid"`
	DNSAzurednsAppid            types.String `tfsdk:"dns_azuredns_appid"`
	DNSAzurednsClientsecret     types.String `tfsdk:"dns_azuredns_clientsecret"`
	DNSAzurednsManagedidentity  types.Bool   `tfsdk:"dns_azuredns_managedidentity"`
	DNSBhostedUser              types.String `tfsdk:"dns_bhosted_user"`
	DNSBhostedPassword          types.String `tfsdk:"dns_bhosted_password"`
	DNSBhostedTTL               types.String `tfsdk:"dns_bhosted_ttl"`
	DNSBhostedSld               types.String `tfsdk:"dns_bhosted_sld"`
	DNSBhostedTld               types.String `tfsdk:"dns_bhosted_tld"`
	DNSBunnyApiKey              types.String `tfsdk:"dns_bunny_api_key"`
	DNSCfEmail                  types.String `tfsdk:"dns_cf_email"`
	DNSCfKey                    types.String `tfsdk:"dns_cf_key"`
	DNSCfToken                  types.String `tfsdk:"dns_cf_token"`
	DNSCfAccountID              types.String `tfsdk:"dns_cf_account_id"`
	DNSCfZoneID                 types.String `tfsdk:"dns_cf_zone_id"`
	DNSCloudnsAuthID            types.String `tfsdk:"dns_cloudns_auth_id"`
	DNSCloudnsSubAuthID         types.String `tfsdk:"dns_cloudns_sub_auth_id"`
	DNSCloudnsAuthPassword      types.String `tfsdk:"dns_cloudns_auth_password"`
	DNSCxKey                    types.String `tfsdk:"dns_cx_key"`
	DNSCxSecret                 types.String `tfsdk:"dns_cx_secret"`
	DNSCyonUser                 types.String `tfsdk:"dns_cyon_user"`
	DNSCyonPassword             types.String `tfsdk:"dns_cyon_password"`
	DNSDaKey                    types.String `tfsdk:"dns_da_key"`
	DNSDaInsecure               types.Bool   `tfsdk:"dns_da_insecure"`
	DNSDdnssToken               types.String `tfsdk:"dns_ddnss_token"`
	DNSDgonKey                  types.String `tfsdk:"dns_dgon_key"`
	DNSDnsexitAuthUser          types.String `tfsdk:"dns_dnsexit_auth_user"`
	DNSDnsexitAuthPass          types.String `tfsdk:"dns_dnsexit_auth_pass"`
	DNSDnsexitApi               types.String `tfsdk:"dns_dnsexit_api"`
	DNSDnshomePassword          types.String `tfsdk:"dns_dnshome_password"`
	DNSDnshomeSubdomain         types.String `tfsdk:"dns_dnshome_subdomain"`
	DNSDnsimpleToken            types.String `tfsdk:"dns_dnsimple_token"`
	DNSDnsservicesUser          types.String `tfsdk:"dns_dnsservices_user"`
	DNSDnsservicesPassword      types.String `tfsdk:"dns_dnsservices_password"`
	DNSDoapiToken               types.String `tfsdk:"dns_doapi_token"`
	DNSDoPid                    types.String `tfsdk:"dns_do_pid"`
	DNSDoPassword               types.String `tfsdk:"dns_do_password"`
	DNSDomeneshopToken          types.String `tfsdk:"dns_domeneshop_token"`
	DNSDomeneshopSecret         types.String `tfsdk:"dns_domeneshop_secret"`
	DNSDpID                     types.String `tfsdk:"dns_dp_id"`
	DNSDpKey                    types.String `tfsdk:"dns_dp_key"`
	DNSDhKey                    types.String `tfsdk:"dns_dh_key"`
	DNSDuckdnsToken             types.String `tfsdk:"dns_duckdns_token"`
	DNSDynCustomer              types.String `tfsdk:"dns_dyn_customer"`
	DNSDynUser                  types.String `tfsdk:"dns_dyn_user"`
	DNSDynPassword              types.String `tfsdk:"dns_dyn_password"`
	DNSDynuClientid             types.String `tfsdk:"dns_dynu_clientid"`
	DNSDynuSecret               types.String `tfsdk:"dns_dynu_secret"`
	DNSFreednsUser              types.String `tfsdk:"dns_freedns_user"`
	DNSFreednsPassword          types.String `tfsdk:"dns_freedns_password"`
	DNSFreemyipToken            types.String `tfsdk:"dns_freemyip_token"`
	DNSFornexApiKey             types.String `tfsdk:"dns_fornex_api_key"`
	DNSGandiLivednsKey          types.String `tfsdk:"dns_gandi_livedns_key"`
	DNSGandiLivednsToken        types.String `tfsdk:"dns_gandi_livedns_token"`
	DNSGcloudKey                types.String `tfsdk:"dns_gcloud_key"`
	DNSGcoreToken               types.String `tfsdk:"dns_gcore_token"`
	DNSGoogledomainsAccessToken types.String `tfsdk:"dns_googledomains_access_token"`
	DNSGoogledomainsZone        types.String `tfsdk:"dns_googledomains_zone"`
	DNSGdKey                    types.String `tfsdk:"dns_gd_key"`
	DNSGdSecret                 types.String `tfsdk:"dns_gd_secret"`
	DNSHostingdeServer          types.String `tfsdk:"dns_hostingde_server"`
	DNSHostingdeApiKey          types.String `tfsdk:"dns_hostingde_api_key"`
	DNSHostingerToken           types.String `tfsdk:"dns_hostinger_token"`
	DNSHeUser                   types.String `tfsdk:"dns_he_user"`
	DNSHePassword               types.String `tfsdk:"dns_he_password"`
	DNSHeDdnsKey                types.String `tfsdk:"dns_he_ddns_key"`
	DNSInfobloxCredentials      types.String `tfsdk:"dns_infoblox_credentials"`
	DNSInfobloxServer           types.String `tfsdk:"dns_infoblox_server"`
	DNSInwxUser                 types.String `tfsdk:"dns_inwx_user"`
	DNSInwsPassword             types.String `tfsdk:"dns_inws_password"`
	DNSInwxPassword             types.String `tfsdk:"dns_inwx_password"`
	DNSInwxSharedSecret         types.String `tfsdk:"dns_inwx_shared_secret"`
	DNSIonosPrefix              types.String `tfsdk:"dns_ionos_prefix"`
	DNSIonosSecret              types.String `tfsdk:"dns_ionos_secret"`
	DNSIpv64Token               types.String `tfsdk:"dns_ipv64_token"`
	DNSIspconfigUser            types.String `tfsdk:"dns_ispconfig_user"`
	DNSIspconfigPassword        types.String `tfsdk:"dns_ispconfig_password"`
	DNSIspconfigApi             types.String `tfsdk:"dns_ispconfig_api"`
	DNSIspconfigInsecure        types.Bool   `tfsdk:"dns_ispconfig_insecure"`
	DNSJdID                     types.String `tfsdk:"dns_jd_id"`
	DNSJdRegion                 types.String `tfsdk:"dns_jd_region"`
	DNSJdSecret                 types.String `tfsdk:"dns_jd_secret"`
	DNSJokerUsername            types.String `tfsdk:"dns_joker_username"`
	DNSJokerPassword            types.String `tfsdk:"dns_joker_password"`
	DNSKinghostUsername         types.String `tfsdk:"dns_kinghost_username"`
	DNSKinghostPassword         types.String `tfsdk:"dns_kinghost_password"`
	DNSKnotServer               types.String `tfsdk:"dns_knot_server"`
	DNSKnotKey                  types.String `tfsdk:"dns_knot_key"`
	DNSLexiconProvider          types.String `tfsdk:"dns_lexicon_provider"`
	DNSLexiconUser              types.String `tfsdk:"dns_lexicon_user"`
	DNSLexiconToken             types.String `tfsdk:"dns_lexicon_token"`
	DNSLimacityApikey           types.String `tfsdk:"dns_limacity_apikey"`
	DNSLinodeKey                types.String `tfsdk:"dns_linode_key"`
	DNSLinodeV4Key              types.String `tfsdk:"dns_linode_v4_key"`
	DNSLoopiaApi                types.String `tfsdk:"dns_loopia_api"`
	DNSLoopiaUser               types.String `tfsdk:"dns_loopia_user"`
	DNSLoopiaPassword           types.String `tfsdk:"dns_loopia_password"`
	DNSLuaEmail                 types.String `tfsdk:"dns_lua_email"`
	DNSLuaKey                   types.String `tfsdk:"dns_lua_key"`
	DNSMiabUser                 types.String `tfsdk:"dns_miab_user"`
	DNSMiabPassword             types.String `tfsdk:"dns_miab_password"`
	DNSMiabServer               types.String `tfsdk:"dns_miab_server"`
	DNSMeKey                    types.String `tfsdk:"dns_me_key"`
	DNSMeSecret                 types.String `tfsdk:"dns_me_secret"`
	DNSMydnsjpMasterid          types.String `tfsdk:"dns_mydnsjp_masterid"`
	DNSMydnsjpPassword          types.String `tfsdk:"dns_mydnsjp_password"`
	DNSMylocToken               types.String `tfsdk:"dns_myloc_token"`
	DNSMythicBeastsKey          types.String `tfsdk:"dns_mythic_beasts_key"`
	DNSMythicBeastsSecret       types.String `tfsdk:"dns_mythic_beasts_secret"`
	DNSNamecheapUser            types.String `tfsdk:"dns_namecheap_user"`
	DNSNamecheapApi             types.String `tfsdk:"dns_namecheap_api"`
	DNSNamecheapSourceip        types.String `tfsdk:"dns_namecheap_sourceip"`
	DNSNamecomUser              types.String `tfsdk:"dns_namecom_user"`
	DNSNamecomToken             types.String `tfsdk:"dns_namecom_token"`
	DNSNamesiloKey              types.String `tfsdk:"dns_namesilo_key"`
	DNSNederhostKey             types.String `tfsdk:"dns_nederhost_key"`
	DNSNetcupCid                types.String `tfsdk:"dns_netcup_cid"`
	DNSNetcupKey                types.String `tfsdk:"dns_netcup_key"`
	DNSNetcupPw                 types.String `tfsdk:"dns_netcup_pw"`
	DNSNjallaToken              types.String `tfsdk:"dns_njalla_token"`
	DNSNsoneKey                 types.String `tfsdk:"dns_nsone_key"`
	DNSNsupdateServer           types.String `tfsdk:"dns_nsupdate_server"`
	DNSNsupdateZone             types.String `tfsdk:"dns_nsupdate_zone"`
	DNSNsupdateKey              types.String `tfsdk:"dns_nsupdate_key"`
	DNSOciCliUser               types.String `tfsdk:"dns_oci_cli_user"`
	DNSOciCliTenancy            types.String `tfsdk:"dns_oci_cli_tenancy"`
	DNSOciCliRegion             types.String `tfsdk:"dns_oci_cli_region"`
	DNSOciCliKey                types.String `tfsdk:"dns_oci_cli_key"`
	DNSOnlineKey                types.String `tfsdk:"dns_online_key"`
	DNSOpenproviderRestUsername types.String `tfsdk:"dns_openprovider_rest_username"`
	DNSOpenproviderRestPassword types.String `tfsdk:"dns_openprovider_rest_password"`
	DNSOpnsenseHost             types.String `tfsdk:"dns_opnsense_host"`
	DNSOpnsensePort             types.String `tfsdk:"dns_opnsense_port"`
	DNSOpnsenseKey              types.String `tfsdk:"dns_opnsense_key"`
	DNSOpnsenseToken            types.String `tfsdk:"dns_opnsense_token"`
	DNSOpnsenseInsecure         types.Bool   `tfsdk:"dns_opnsense_insecure"`
	DNSOvhAppKey                types.String `tfsdk:"dns_ovh_app_key"`
	DNSOvhAppSecret             types.String `tfsdk:"dns_ovh_app_secret"`
	DNSOvhConsumerKey           types.String `tfsdk:"dns_ovh_consumer_key"`
	DNSOvhEndpoint              types.String `tfsdk:"dns_ovh_endpoint"`
	DNSPleskxmlUser             types.String `tfsdk:"dns_pleskxml_user"`
	DNSPleskxmlPass             types.String `tfsdk:"dns_pleskxml_pass"`
	DNSPleskxmlUri              types.String `tfsdk:"dns_pleskxml_uri"`
	DNSPdnsURL                  types.String `tfsdk:"dns_pdns_url"`
	DNSPdnsServerid             types.String `tfsdk:"dns_pdns_serverid"`
	DNSPdnsToken                types.String `tfsdk:"dns_pdns_token"`
	DNSPorkbunKey               types.String `tfsdk:"dns_porkbun_key"`
	DNSPorkbunSecret            types.String `tfsdk:"dns_porkbun_secret"`
	DNSSlKey                    types.String `tfsdk:"dns_sl_key"`
	DNSSlApiver                 types.String `tfsdk:"dns_sl_apiver"`
	DNSSlTokenLifetime          types.String `tfsdk:"dns_sl_token_lifetime"`
	DNSSlAccountID              types.String `tfsdk:"dns_sl_account_id"`
	DNSSlProjectName            types.String `tfsdk:"dns_sl_project_name"`
	DNSSlLoginName              types.String `tfsdk:"dns_sl_login_name"`
	DNSSlPassword               types.String `tfsdk:"dns_sl_password"`
	DNSSelfhostApikey           types.String `tfsdk:"dns_selfhost_apikey"`
	DNSSelfhostMap              types.String `tfsdk:"dns_selfhost_map"`
	DNSServercowUsername        types.String `tfsdk:"dns_servercow_username"`
	DNSServercowPassword        types.String `tfsdk:"dns_servercow_password"`
	DNSSimplyApiKey             types.String `tfsdk:"dns_simply_api_key"`
	DNSSimplyAccountName        types.String `tfsdk:"dns_simply_account_name"`
	DNSSpaceshipApiKey          types.String `tfsdk:"dns_spaceship_api_key"`
	DNSSpaceshipApiSecret       types.String `tfsdk:"dns_spaceship_api_secret"`
	DNSSpaceshipRootDomain      types.String `tfsdk:"dns_spaceship_root_domain"`
	DNSTechnitiumToken          types.String `tfsdk:"dns_technitium_token"`
	DNSTechnitiumHostname       types.String `tfsdk:"dns_technitium_hostname"`
	DNSTransipUsername          types.String `tfsdk:"dns_transip_username"`
	DNSTransipKey               types.String `tfsdk:"dns_transip_key"`
	DNSTransipTokenGlobalKey    types.Bool   `tfsdk:"dns_transip_token_global_key"`
	DNSTimewebToken             types.String `tfsdk:"dns_timeweb_token"`
	DNSUdrUser                  types.String `tfsdk:"dns_udr_user"`
	DNSUdrPassword              types.String `tfsdk:"dns_udr_password"`
	DNSUnoKey                   types.String `tfsdk:"dns_uno_key"`
	DNSUnoUser                  types.String `tfsdk:"dns_uno_user"`
	DNSVscaleKey                types.String `tfsdk:"dns_vscale_key"`
	DNSVultrKey                 types.String `tfsdk:"dns_vultr_key"`
	DNSYandexToken              types.String `tfsdk:"dns_yandex_token"`
	DNSZiloreKey                types.String `tfsdk:"dns_zilore_key"`
	DNSZmKey                    types.String `tfsdk:"dns_zm_key"`
	DNSGdnsdkUser               types.String `tfsdk:"dns_gdnsdk_user"`
	DNSGdnsdkPassword           types.String `tfsdk:"dns_gdnsdk_password"`
	DNSAcmednsUser              types.String `tfsdk:"dns_acmedns_user"`
	DNSAcmednsPassword          types.String `tfsdk:"dns_acmedns_password"`
	DNSAcmednsSubdomain         types.String `tfsdk:"dns_acmedns_subdomain"`
	DNSAcmednsUpdateurl         types.String `tfsdk:"dns_acmedns_updateurl"`
	DNSAcmednsBaseurl           types.String `tfsdk:"dns_acmedns_baseurl"`
	DNSAcmeproxyEndpoint        types.String `tfsdk:"dns_acmeproxy_endpoint"`
	DNSAcmeproxyUsername        types.String `tfsdk:"dns_acmeproxy_username"`
	DNSAcmeproxyPassword        types.String `tfsdk:"dns_acmeproxy_password"`
	DNSVariomediaKey            types.String `tfsdk:"dns_variomedia_key"`
	DNSSchlundtechUser          types.String `tfsdk:"dns_schlundtech_user"`
	DNSSchlundtechPassword      types.String `tfsdk:"dns_schlundtech_password"`
	DNSEasydnsApitoken          types.String `tfsdk:"dns_easydns_apitoken"`
	DNSEasydnsApikey            types.String `tfsdk:"dns_easydns_apikey"`
	DNSEdgecenterApikey         types.String `tfsdk:"dns_edgecenter_apikey"`
	DNSEuservUser               types.String `tfsdk:"dns_euserv_user"`
	DNSEuservPassword           types.String `tfsdk:"dns_euserv_password"`
	DNSLeasewebKey              types.String `tfsdk:"dns_leaseweb_key"`
	DNSCnUser                   types.String `tfsdk:"dns_cn_user"`
	DNSCnPassword               types.String `tfsdk:"dns_cn_password"`
	DNSArvanToken               types.String `tfsdk:"dns_arvan_token"`
	DNSArtfilesUsername         types.String `tfsdk:"dns_artfiles_username"`
	DNSArtfilesPassword         types.String `tfsdk:"dns_artfiles_password"`
	DNSHetznerToken             types.String `tfsdk:"dns_hetzner_token"`
	DNSHetznercloudToken        types.String `tfsdk:"dns_hetznercloud_token"`
	DNSHexonetLogin             types.String `tfsdk:"dns_hexonet_login"`
	DNSHexonetPassword          types.String `tfsdk:"dns_hexonet_password"`
	DNS1984hostingUser          types.String `tfsdk:"dns_1984hosting_user"`
	DNS1984hostingPassword      types.String `tfsdk:"dns_1984hosting_password"`
	DNS1984hostingTotpSecret    types.String `tfsdk:"dns_1984hosting_totp_secret"`
	DNSKasLogin                 types.String `tfsdk:"dns_kas_login"`
	DNSKasAuthdata              types.String `tfsdk:"dns_kas_authdata"`
	DNSKasAuthtype              types.String `tfsdk:"dns_kas_authtype"`
	DNSDnsmintApiKey            types.String `tfsdk:"dns_dnsmint_api_key"`
	DNSDesecToken               types.String `tfsdk:"dns_desec_token"`
	DNSDesecName                types.String `tfsdk:"dns_desec_name"`
	DNSInfomaniakToken          types.String `tfsdk:"dns_infomaniak_token"`
	DNSZoneUsername             types.String `tfsdk:"dns_zone_username"`
	DNSZoneKey                  types.String `tfsdk:"dns_zone_key"`
	DNSDynv6Token               types.String `tfsdk:"dns_dynv6_token"`
	DNSCpanelUser               types.String `tfsdk:"dns_cpanel_user"`
	DNSCpanelToken              types.String `tfsdk:"dns_cpanel_token"`
	DNSCpanelHostname           types.String `tfsdk:"dns_cpanel_hostname"`
	DNSRegruUsername            types.String `tfsdk:"dns_regru_username"`
	DNSRegruPassword            types.String `tfsdk:"dns_regru_password"`
	DNSNicUsername              types.String `tfsdk:"dns_nic_username"`
	DNSNicPassword              types.String `tfsdk:"dns_nic_password"`
	DNSNicClient                types.String `tfsdk:"dns_nic_client"`
	DNSNicSecret                types.String `tfsdk:"dns_nic_secret"`
	DNSWebsupportApiKey         types.String `tfsdk:"dns_websupport_api_key"`
	DNSWebsupportApiSecret      types.String `tfsdk:"dns_websupport_api_secret"`
	DNSWedosUsername            types.String `tfsdk:"dns_wedos_username"`
	DNSWedosWapipass            types.String `tfsdk:"dns_wedos_wapipass"`
	DNSWorld4youUsername        types.String `tfsdk:"dns_world4you_username"`
	DNSWorld4youPassword        types.String `tfsdk:"dns_world4you_password"`
	DNSAuroraKey                types.String `tfsdk:"dns_aurora_key"`
	DNSAuroraSecret             types.String `tfsdk:"dns_aurora_secret"`
	DNSConohaUser               types.String `tfsdk:"dns_conoha_user"`
	DNSConohaPassword           types.String `tfsdk:"dns_conoha_password"`
	DNSConohaTenantid           types.String `tfsdk:"dns_conoha_tenantid"`
	DNSConohaIdapi              types.String `tfsdk:"dns_conoha_idapi"`
	DNSConstellixKey            types.String `tfsdk:"dns_constellix_key"`
	DNSConstellixSecret         types.String `tfsdk:"dns_constellix_secret"`
	DNSExoscaleKey              types.String `tfsdk:"dns_exoscale_key"`
	DNSExoscaleSecret           types.String `tfsdk:"dns_exoscale_secret"`
	DNSInternetbsKey            types.String `tfsdk:"dns_internetbs_key"`
	DNSInternetbsPassword       types.String `tfsdk:"dns_internetbs_password"`
	DNSPointhqKey               types.String `tfsdk:"dns_pointhq_key"`
	DNSPointhqEmail             types.String `tfsdk:"dns_pointhq_email"`
	DNSRackspaceUser            types.String `tfsdk:"dns_rackspace_user"`
	DNSRackspaceKey             types.String `tfsdk:"dns_rackspace_key"`
	DNSRage4Token               types.String `tfsdk:"dns_rage4_token"`
	DNSRage4User                types.String `tfsdk:"dns_rage4_user"`
	DNSRcode0ApiToken           types.String `tfsdk:"dns_rcode0_api_token"`
	DNSRcode0URL                types.String `tfsdk:"dns_rcode0_url"`
	DNSRcode0TTL                types.Int64  `tfsdk:"dns_rcode0_ttl"`
	DNSMijnhostApiKey           types.String `tfsdk:"dns_mijnhost_api_key"`
	DNSScalewayToken            types.String `tfsdk:"dns_scaleway_token"`
	DNSZoneeditID               types.String `tfsdk:"dns_zoneedit_id"`
	DNSZoneeditToken            types.String `tfsdk:"dns_zoneedit_token"`

	Id types.String `tfsdk:"id"`
}

func validationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "ACME challenge types (**Services ▸ ACME Client ▸ Challenge Types**, plugin `os-acme-client`): HTTP-01, TLS-ALPN-01 or DNS-01 with one of the acme.sh DNS providers. Only the attributes of the chosen `dns_service` are used.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this validation. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this validation.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this validation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"method": schema.StringAttribute{
				MarkdownDescription: "Set the ACME challenge type. You'll have to add configuration for the selected challenge type below. One of `http01`, `dns01`, `tlsalpn01`. Defaults to `dns01`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("dns01"),
				Validators:          []validator.String{stringvalidator.OneOf("http01", "dns01", "tlsalpn01")},
			},
			"http_service": schema.StringAttribute{
				MarkdownDescription: "HTTP Service. One of `opnsense`, `haproxy`. Defaults to `opnsense`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("opnsense"),
				Validators:          []validator.String{stringvalidator.OneOf("opnsense", "haproxy")},
			},
			"http_opn_autodiscovery": schema.BoolAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose this option to let OPNsense try to auto-discover these IP addresses. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_opn_interface": schema.StringAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose the interface where this IP address is currently configured. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"http_opn_ipaddresses": schema.SetAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to one or more official IP addresses. Enter the all of these IP addresses here. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with these IP addresses.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"http_haproxy_inject": schema.BoolAttribute{
				MarkdownDescription: "Automatically inject config into the local HAProxy instance to let it serve acme challanges without service interruption. Of course, adding the configuration requires a short restart of the HAProxy service. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_haproxy_frontends": schema.SetAttribute{
				MarkdownDescription: "Choose the local HAProxy frontends. They will automatically be configured to redirect acme challenges to the internal acme client. The HAProxy service will automatically be restarted if a certificate was renewed. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tlsalpn_service": schema.StringAttribute{
				MarkdownDescription: "TLS-ALPN Service. One of `acme`. Defaults to `acme`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("acme"),
				Validators:          []validator.String{stringvalidator.OneOf("acme")},
			},
			"tlsalpn_acme_autodiscovery": schema.BoolAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose this option to let OPNsense try to auto-discover these IP addresses. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"tlsalpn_acme_interface": schema.StringAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose the interface where this IP address is currently configured. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tlsalpn_acme_ipaddresses": schema.SetAttribute{
				MarkdownDescription: "The FQDN's used in your certificate must currently point to one or more official IP addresses. Enter the all of these IP addresses here. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with these IP addresses.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall. Defaults to `[]`.",
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"dns_service": schema.StringAttribute{
				MarkdownDescription: "DNS Service. One of `dns_1984hosting`, `dns_acmedns`, `dns_acmeproxy`, `dns_active24`, `dns_ad`, `dns_ali`, `dns_kas`, `dns_arvan`, `dns_artfiles`, `dns_aurora`, `dns_autodns`, `dns_aws`, `dns_azure`, `dns_bhosted`, `dns_bunny`, `dns_cloudns`, `dns_cf`, `dns_cx`, `dns_cn`, `dns_conoha`, `dns_constellix`, `dns_cpanel`, `dns_cyon`, `dns_ddnss`, `dns_desec`, `dns_dgon`, `dns_da`, `dns_dnsexit`, `dns_dnshome`, `dns_dnsimple`, `dns_dnsmint`, `dns_dnsservices`, `dns_domeneshop`, `dns_me`, `dns_dp`, `dns_doapi`, `dns_do`, `dns_dreamhost`, `dns_duckdns`, `dns_dyn`, `dns_dynu`, `dns_dynv6`, `dns_easydns`, `dns_edgecenter`, `dns_euserv`, `dns_exoscale`, `dns_fornex`, `dns_freedns`, `dns_freemyip`, `dns_gandi_livedns`, `dns_gcore`, `dns_gd`, `dns_gcloud`, `dns_googledomains`, `dns_gdnsdk`, `dns_hetzner`, `dns_hetznercloud`, `dns_hexonet`, `dns_hostingde`, `dns_hostinger`, `dns_he`, `dns_he_ddns`, `dns_infoblox`, `dns_infomaniak`, `dns_internetbs`, `dns_inwx`, `dns_ionos`, `dns_ipv64`, `dns_ispconfig`, `dns_jd`, `dns_joker`, `dns_kinghost`, `dns_knot`, `dns_leaseweb`, `dns_lexicon`, `dns_limacity`, `dns_linode`, `dns_linode_v4`, `dns_loopia`, `dns_lua`, `dns_miab`, `dns_mijnhost`, `dns_mydnsjp`, `dns_myloc`, `dns_mythic_beasts`, `dns_namecom`, `dns_namecheap`, `dns_namesilo`, `dns_nederhost`, `dns_netcup`, `dns_nic`, `dns_njalla`, `dns_nsone`, `dns_nsupdate`, `dns_online`, `dns_openprovider_rest`, `dns_opnsense`, `dns_oci`, `dns_ovh`, `dns_pdns`, `dns_pleskxml`, `dns_pointhq`, `dns_porkbun`, `dns_rackspace`, `dns_rage4`, `dns_rcode0`, `dns_regru`, `dns_scaleway`, `dns_schlundtech`, `dns_selectel`, `dns_selfhost`, `dns_servercow`, `dns_simply`, `dns_spaceship`, `dns_technitium`, `dns_timeweb`, `dns_transip`, `dns_udr`, `dns_unoeuro`, `dns_variomedia`, `dns_vscale`, `dns_vultr`, `dns_websupport`, `dns_wedos`, `dns_world4you`, `dns_yandex`, `dns_zilore`, `dns_zone`, `dns_zoneedit`, `dns_zonomi`. Defaults to `dns_freedns`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("dns_freedns"),
				Validators:          []validator.String{stringvalidator.OneOf("dns_1984hosting", "dns_acmedns", "dns_acmeproxy", "dns_active24", "dns_ad", "dns_ali", "dns_kas", "dns_arvan", "dns_artfiles", "dns_aurora", "dns_autodns", "dns_aws", "dns_azure", "dns_bhosted", "dns_bunny", "dns_cloudns", "dns_cf", "dns_cx", "dns_cn", "dns_conoha", "dns_constellix", "dns_cpanel", "dns_cyon", "dns_ddnss", "dns_desec", "dns_dgon", "dns_da", "dns_dnsexit", "dns_dnshome", "dns_dnsimple", "dns_dnsmint", "dns_dnsservices", "dns_domeneshop", "dns_me", "dns_dp", "dns_doapi", "dns_do", "dns_dreamhost", "dns_duckdns", "dns_dyn", "dns_dynu", "dns_dynv6", "dns_easydns", "dns_edgecenter", "dns_euserv", "dns_exoscale", "dns_fornex", "dns_freedns", "dns_freemyip", "dns_gandi_livedns", "dns_gcore", "dns_gd", "dns_gcloud", "dns_googledomains", "dns_gdnsdk", "dns_hetzner", "dns_hetznercloud", "dns_hexonet", "dns_hostingde", "dns_hostinger", "dns_he", "dns_he_ddns", "dns_infoblox", "dns_infomaniak", "dns_internetbs", "dns_inwx", "dns_ionos", "dns_ipv64", "dns_ispconfig", "dns_jd", "dns_joker", "dns_kinghost", "dns_knot", "dns_leaseweb", "dns_lexicon", "dns_limacity", "dns_linode", "dns_linode_v4", "dns_loopia", "dns_lua", "dns_miab", "dns_mijnhost", "dns_mydnsjp", "dns_myloc", "dns_mythic_beasts", "dns_namecom", "dns_namecheap", "dns_namesilo", "dns_nederhost", "dns_netcup", "dns_nic", "dns_njalla", "dns_nsone", "dns_nsupdate", "dns_online", "dns_openprovider_rest", "dns_opnsense", "dns_oci", "dns_ovh", "dns_pdns", "dns_pleskxml", "dns_pointhq", "dns_porkbun", "dns_rackspace", "dns_rage4", "dns_rcode0", "dns_regru", "dns_scaleway", "dns_schlundtech", "dns_selectel", "dns_selfhost", "dns_servercow", "dns_simply", "dns_spaceship", "dns_technitium", "dns_timeweb", "dns_transip", "dns_udr", "dns_unoeuro", "dns_variomedia", "dns_vscale", "dns_vultr", "dns_websupport", "dns_wedos", "dns_world4you", "dns_yandex", "dns_zilore", "dns_zone", "dns_zoneedit", "dns_zonomi")},
			},
			"dns_sleep": schema.Int64Attribute{
				MarkdownDescription: "The time in seconds to wait for all the TXT records to take effect after adding them to the DNS API. Defaults to 0 seconds, which causes Acme Client to check public DNS services every 10 seconds for up to 20 minutes. If set to a non-zero value, a fixed DNS sleep time will be used and the local DNS servers will be queried instead. A DNS sleep time of 120 seconds or more is recommended for some DNS APIs. Defaults to `0`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"dns_active24_api_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_active24_api_secret": schema.StringAttribute{
				MarkdownDescription: "API Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ad_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ali_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ali_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_autodns_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_autodns_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_autodns_context": schema.StringAttribute{
				MarkdownDescription: "Context. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_aws_id": schema.StringAttribute{
				MarkdownDescription: "AWS ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_aws_secret": schema.StringAttribute{
				MarkdownDescription: "AWS Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_azuredns_subscriptionid": schema.StringAttribute{
				MarkdownDescription: "Subscription ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_azuredns_tenantid": schema.StringAttribute{
				MarkdownDescription: "Tenant ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_azuredns_appid": schema.StringAttribute{
				MarkdownDescription: "APP ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_azuredns_clientsecret": schema.StringAttribute{
				MarkdownDescription: "Client Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_azuredns_managedidentity": schema.BoolAttribute{
				MarkdownDescription: "When System Assigned Managed Identity is enabled the Tenant ID, APP ID and Client Secret settings are ignored by the acme client. Access tokens are obtained using the Azure Instance Metadata Service for the System Assigned Managed Identity. See documentation. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dns_bhosted_user": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_bhosted_password": schema.StringAttribute{
				MarkdownDescription: "API password (MD5 hash like bHosted.nl web services example). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_bhosted_ttl": schema.StringAttribute{
				MarkdownDescription: "Optional override TTL for TXT record (default: 300). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_bhosted_sld": schema.StringAttribute{
				MarkdownDescription: "Optional override (useful for multi-part TLDs like co.uk). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_bhosted_tld": schema.StringAttribute{
				MarkdownDescription: "Optional override (useful for multi-part TLDs like co.uk). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_bunny_api_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cf_email": schema.StringAttribute{
				MarkdownDescription: "Use your CF account with full privileges. This is less secure, a API token should be used instead (see below). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cf_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cf_token": schema.StringAttribute{
				MarkdownDescription: "The token needs \"Read\" access to Zone.Zone and \"Edit\" access to Zone.DNS across all zones from an account. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cf_account_id": schema.StringAttribute{
				MarkdownDescription: "Can be found in the URI after logging into the Cloudflare dashboard. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cf_zone_id": schema.StringAttribute{
				MarkdownDescription: "Note that specifying a Zone ID will limit this configuration to a single domain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cloudns_auth_id": schema.StringAttribute{
				MarkdownDescription: "Auth ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cloudns_sub_auth_id": schema.StringAttribute{
				MarkdownDescription: "Sub Auth ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cloudns_auth_password": schema.StringAttribute{
				MarkdownDescription: "Auth Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cx_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cx_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cyon_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cyon_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_da_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_da_insecure": schema.BoolAttribute{
				MarkdownDescription: "Uncheck this box if you have a valid SSL certificate for your DirectAdmin installation. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dns_ddnss_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dgon_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dnsexit_auth_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dnsexit_auth_pass": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dnsexit_api": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dnshome_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dnshome_subdomain": schema.StringAttribute{
				MarkdownDescription: "Subdomain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dnsimple_token": schema.StringAttribute{
				MarkdownDescription: "Note that this is the account token not the user token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dnsservices_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dnsservices_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_doapi_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_do_pid": schema.StringAttribute{
				MarkdownDescription: "Partner ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_do_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_domeneshop_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_domeneshop_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dp_id": schema.StringAttribute{
				MarkdownDescription: "ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dp_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dh_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_duckdns_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dyn_customer": schema.StringAttribute{
				MarkdownDescription: "Customer. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dyn_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dyn_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dynu_clientid": schema.StringAttribute{
				MarkdownDescription: "Client ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_dynu_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_freedns_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_freedns_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_freemyip_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_fornex_api_key": schema.StringAttribute{
				MarkdownDescription: "ApiKey. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gandi_livedns_key": schema.StringAttribute{
				MarkdownDescription: "The API key is the previous mechanism that was replaced with Personal Access Tokens. API keys should no longer be used. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gandi_livedns_token": schema.StringAttribute{
				MarkdownDescription: "Personal Access Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gcloud_key": schema.StringAttribute{
				MarkdownDescription: "Provide a service account key in JSON format for your Google Cloud account. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gcore_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_googledomains_access_token": schema.StringAttribute{
				MarkdownDescription: "Access Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_googledomains_zone": schema.StringAttribute{
				MarkdownDescription: "Zone. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_gd_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gd_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_hostingde_server": schema.StringAttribute{
				MarkdownDescription: "Enter the API endpoint, e.g. https://secure.hosting.de or https://partner.http.net. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_hostingde_api_key": schema.StringAttribute{
				MarkdownDescription: "API key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_hostinger_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_he_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_he_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_he_ddns_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_infoblox_credentials": schema.StringAttribute{
				MarkdownDescription: "Use the following format: username:password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_infoblox_server": schema.StringAttribute{
				MarkdownDescription: "Enter either the IP address or FQDN of your Infoblox appliance. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_inwx_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_inws_password": schema.StringAttribute{
				MarkdownDescription: "dns_inws_password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_inwx_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_inwx_shared_secret": schema.StringAttribute{
				MarkdownDescription: "When 2FA is enabled, the Shared Secret must be provided. Note that this feature requires the package oath-toolkit, which must be installed manually. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ionos_prefix": schema.StringAttribute{
				MarkdownDescription: "Prefix. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_ionos_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ipv64_token": schema.StringAttribute{
				MarkdownDescription: "IPv64.net API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ispconfig_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_ispconfig_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ispconfig_api": schema.StringAttribute{
				MarkdownDescription: "API URL. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_ispconfig_insecure": schema.BoolAttribute{
				MarkdownDescription: "Uncheck this box if you have a valid SSL certificate for your ISPConfig installation. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dns_jd_id": schema.StringAttribute{
				MarkdownDescription: "API key id. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_jd_region": schema.StringAttribute{
				MarkdownDescription: "JD Cloud region. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_jd_secret": schema.StringAttribute{
				MarkdownDescription: "API key secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_joker_username": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_joker_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_kinghost_username": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_kinghost_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_knot_server": schema.StringAttribute{
				MarkdownDescription: "Knot server address, like: dns.example.com. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_knot_key": schema.StringAttribute{
				MarkdownDescription: "Specify the location of the generated TSIG Key inside the TSIG file using grep and cut, example: grep \\# /etc/knot/acme.key | cut -d' ' -f2. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_lexicon_provider": schema.StringAttribute{
				MarkdownDescription: "Provider. One of `aliyun`, `aurora`, `auto`, `azure`, `cloudflare`, `cloudns`, `cloudxns`, `conoha`, `constellix`, `digitalocean`, `dinahosting`, `directadmin`, `dnsimple`, `dnsmadeeasy`, `dnspark`, `dnspod`, `dreamhost`, `easydns`, `easyname`, `exoscale`, `gandi`, `gehirn`, `glesys`, `godaddy`, `googleclouddns`, `gratisdns`, `henet`, `hetzner`, `hover`, `infoblox`, `internetbs`, `inwx`, `linode`, `linode4`, `localzone`, `luadns`, `memset`, `namecheap`, `namesilo`, `netcup`, `nfsn`, `nsone`, `onapp`, `online`, `ovh`, `plesk`, `pointhq`, `powerdns`, `rackspace`, `rage4`, `route53`, `safedns`, `sakuracloud`, `softlayer`, `subreg`, `transip`, `vultr`, `yandex`, `zeit`, `zilore`, `zonomi`. Defaults to `cloudflare`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("cloudflare"),
				Validators:          []validator.String{stringvalidator.OneOf("", "aliyun", "aurora", "auto", "azure", "cloudflare", "cloudns", "cloudxns", "conoha", "constellix", "digitalocean", "dinahosting", "directadmin", "dnsimple", "dnsmadeeasy", "dnspark", "dnspod", "dreamhost", "easydns", "easyname", "exoscale", "gandi", "gehirn", "glesys", "godaddy", "googleclouddns", "gratisdns", "henet", "hetzner", "hover", "infoblox", "internetbs", "inwx", "linode", "linode4", "localzone", "luadns", "memset", "namecheap", "namesilo", "netcup", "nfsn", "nsone", "onapp", "online", "ovh", "plesk", "pointhq", "powerdns", "rackspace", "rage4", "route53", "safedns", "sakuracloud", "softlayer", "subreg", "transip", "vultr", "yandex", "zeit", "zilore", "zonomi")},
			},
			"dns_lexicon_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_lexicon_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_limacity_apikey": schema.StringAttribute{
				MarkdownDescription: "API-Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_linode_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_linode_v4_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_loopia_api": schema.StringAttribute{
				MarkdownDescription: "API Endpoint. Defaults to `https://api.loopia.se/RPCSERV`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("https://api.loopia.se/RPCSERV"),
			},
			"dns_loopia_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_loopia_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_lua_email": schema.StringAttribute{
				MarkdownDescription: "E-Mail. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_lua_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_miab_user": schema.StringAttribute{
				MarkdownDescription: "User with admin priviledges. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_miab_password": schema.StringAttribute{
				MarkdownDescription: "Password for admin user. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_miab_server": schema.StringAttribute{
				MarkdownDescription: "MailinaBox Server FQDN. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_me_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_me_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_mydnsjp_masterid": schema.StringAttribute{
				MarkdownDescription: "Master ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_mydnsjp_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_myloc_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_mythic_beasts_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_mythic_beasts_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_namecheap_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_namecheap_api": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_namecheap_sourceip": schema.StringAttribute{
				MarkdownDescription: "Source IP. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_namecom_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_namecom_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_namesilo_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_nederhost_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_netcup_cid": schema.StringAttribute{
				MarkdownDescription: "Customer number. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_netcup_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_netcup_pw": schema.StringAttribute{
				MarkdownDescription: "API Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_njalla_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_nsone_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_nsupdate_server": schema.StringAttribute{
				MarkdownDescription: "Server (FQDN). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_nsupdate_zone": schema.StringAttribute{
				MarkdownDescription: "Optionally set the name of the hosted zone (e.g. example.com) as some DNS Providers require, like dyn.com's 'Standard DNS'. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_nsupdate_key": schema.StringAttribute{
				MarkdownDescription: "Requires the whole key file in a format that is compatible with nsupdate. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_oci_cli_user": schema.StringAttribute{
				MarkdownDescription: "OCID of the user calling the API. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_oci_cli_tenancy": schema.StringAttribute{
				MarkdownDescription: "OCID of your tenancy. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_oci_cli_region": schema.StringAttribute{
				MarkdownDescription: "Your OCI region. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_oci_cli_key": schema.StringAttribute{
				MarkdownDescription: "Please refer to the acme.sh documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_online_key": schema.StringAttribute{
				MarkdownDescription: "API key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_openprovider_rest_username": schema.StringAttribute{
				MarkdownDescription: "Account Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_openprovider_rest_password": schema.StringAttribute{
				MarkdownDescription: "Account Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_opnsense_host": schema.StringAttribute{
				MarkdownDescription: "OPNSense Server (FQDN). Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"dns_opnsense_port": schema.StringAttribute{
				MarkdownDescription: "OPNSense Server Port. Defaults to `443`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("443"),
			},
			"dns_opnsense_key": schema.StringAttribute{
				MarkdownDescription: "User API key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_opnsense_token": schema.StringAttribute{
				MarkdownDescription: "User API token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_opnsense_insecure": schema.BoolAttribute{
				MarkdownDescription: "Uncheck this box if you have a valid certificate for your OPNsense interface. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dns_ovh_app_key": schema.StringAttribute{
				MarkdownDescription: "Application Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ovh_app_secret": schema.StringAttribute{
				MarkdownDescription: "Application Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ovh_consumer_key": schema.StringAttribute{
				MarkdownDescription: "Consumer Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_ovh_endpoint": schema.StringAttribute{
				MarkdownDescription: "Specify the OVH endpoint, i.e. ovh-eu, ovh-ca, kimsufi-eu, etc. Please refer to the acme.sh documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_pleskxml_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_pleskxml_pass": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_pleskxml_uri": schema.StringAttribute{
				MarkdownDescription: "Plesk XML URIs often look similar to: https://my-plesk-site.com:8443/enterprise/control/agent.php. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_pdns_url": schema.StringAttribute{
				MarkdownDescription: "Specify the URL for your PowerDNS server, i.e. http://ns.example.com:8081. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_pdns_serverid": schema.StringAttribute{
				MarkdownDescription: "Specify the Server ID of your PowerDNS server, i.e. localhost. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_pdns_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_porkbun_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_porkbun_secret": schema.StringAttribute{
				MarkdownDescription: "Secret API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_sl_key": schema.StringAttribute{
				MarkdownDescription: "API key (deprecated). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_sl_apiver": schema.StringAttribute{
				MarkdownDescription: "API Version. One of `v1`, `v2`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "v1", "v2")},
			},
			"dns_sl_token_lifetime": schema.StringAttribute{
				MarkdownDescription: "Token lifetime in minutes (0-1440). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_sl_account_id": schema.StringAttribute{
				MarkdownDescription: "The account number can be found on the Selectel control panel. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_sl_project_name": schema.StringAttribute{
				MarkdownDescription: "Project name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_sl_login_name": schema.StringAttribute{
				MarkdownDescription: "The service username can be found in the Selectel control panel. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_sl_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_selfhost_apikey": schema.StringAttribute{
				MarkdownDescription: "Create a DNS-01 API key in the Selfhost portal (menu \"Entwickler\"). The legacy username/password authentication is no longer supported by Selfhost. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_selfhost_map": schema.StringAttribute{
				MarkdownDescription: "Please create the TXT record with subdomain _acme-challenge first, than you can get the ID from the edit page. Up to two RIDs per fulldomain are supported but at least one must be set, e.g. _acme-challenge.domain.net:RID:RID2. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_servercow_username": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_servercow_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_simply_api_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_simply_account_name": schema.StringAttribute{
				MarkdownDescription: "Account Name. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_spaceship_api_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_spaceship_api_secret": schema.StringAttribute{
				MarkdownDescription: "API Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_spaceship_root_domain": schema.StringAttribute{
				MarkdownDescription: "Root Domain (Optional). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_technitium_token": schema.StringAttribute{
				MarkdownDescription: "Generate a token from Technitium user and grant rights to the necessary zones. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_technitium_hostname": schema.StringAttribute{
				MarkdownDescription: "The full server URL, for example https://technitium:5380. The port number is optional.NOTE: Remove the trailing / if it is present. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_transip_username": schema.StringAttribute{
				MarkdownDescription: "Your TransIP username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_transip_key": schema.StringAttribute{
				MarkdownDescription: "Requires the whole key file in a format that is compatible with TransIP. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_transip_token_global_key": schema.BoolAttribute{
				MarkdownDescription: "When using a global key, IP whitelisting is disabled. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dns_timeweb_token": schema.StringAttribute{
				MarkdownDescription: "Timeweb Cloud API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_udr_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_udr_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_uno_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_uno_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_vscale_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_vultr_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_yandex_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_zilore_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_zm_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_gdnsdk_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_gdnsdk_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_acmedns_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmedns_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_acmedns_subdomain": schema.StringAttribute{
				MarkdownDescription: "Subdomain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmedns_updateurl": schema.StringAttribute{
				MarkdownDescription: "dns_acmedns_updateurl. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmedns_baseurl": schema.StringAttribute{
				MarkdownDescription: "Specify the custom ACME DNS URL, i.e. https://auth.acme-dns.io:443 (optional). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmeproxy_endpoint": schema.StringAttribute{
				MarkdownDescription: "Specify the acmeproxy endpoint URL, i.e. https://acmeproxy.example.com/. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmeproxy_username": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_acmeproxy_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_variomedia_key": schema.StringAttribute{
				MarkdownDescription: "You need to obtain your API key from variomedia's customer support. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_schlundtech_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_schlundtech_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_easydns_apitoken": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_easydns_apikey": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_edgecenter_apikey": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_euserv_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_euserv_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_leaseweb_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cn_user": schema.StringAttribute{
				MarkdownDescription: "API User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cn_password": schema.StringAttribute{
				MarkdownDescription: "API Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_arvan_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_artfiles_username": schema.StringAttribute{
				MarkdownDescription: "API Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_artfiles_password": schema.StringAttribute{
				MarkdownDescription: "API Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_hetzner_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_hetznercloud_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_hexonet_login": schema.StringAttribute{
				MarkdownDescription: "This should be a combination of the username and role ID. For example, when the username is \"myuser\" and role ID is \"testrole\", then \"myuser!testrole\" must be used here. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_hexonet_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_1984hosting_user": schema.StringAttribute{
				MarkdownDescription: "API User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_1984hosting_password": schema.StringAttribute{
				MarkdownDescription: "API Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_1984hosting_totp_secret": schema.StringAttribute{
				MarkdownDescription: "When 2FA is enabled, the Base32 TOTP secret must be provided. Note that this feature requires the package oath-toolkit, which must be installed manually. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_kas_login": schema.StringAttribute{
				MarkdownDescription: "KAS Login. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_kas_authdata": schema.StringAttribute{
				MarkdownDescription: "KAS Authdata. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_kas_authtype": schema.StringAttribute{
				MarkdownDescription: "KAS Authtype. One of `plain`, `sha1`. Defaults to `plain`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("plain"),
				Validators:          []validator.String{stringvalidator.OneOf("", "plain", "sha1")},
			},
			"dns_dnsmint_api_key": schema.StringAttribute{
				MarkdownDescription: "DNSMint API key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_desec_token": schema.StringAttribute{
				MarkdownDescription: "deSEC token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_desec_name": schema.StringAttribute{
				MarkdownDescription: "deSEC domain. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_infomaniak_token": schema.StringAttribute{
				MarkdownDescription: "API Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_zone_username": schema.StringAttribute{
				MarkdownDescription: "Zone Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_zone_key": schema.StringAttribute{
				MarkdownDescription: "Can be generated under ZoneID account management. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_dynv6_token": schema.StringAttribute{
				MarkdownDescription: "HTTP Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cpanel_user": schema.StringAttribute{
				MarkdownDescription: "cPanel Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_cpanel_token": schema.StringAttribute{
				MarkdownDescription: "Please refer to the cPanel documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_cpanel_hostname": schema.StringAttribute{
				MarkdownDescription: "EX: https://hostname:2083. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_regru_username": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_regru_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_nic_username": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_nic_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_nic_client": schema.StringAttribute{
				MarkdownDescription: "ClientId. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_nic_secret": schema.StringAttribute{
				MarkdownDescription: "ClientSecret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_websupport_api_key": schema.StringAttribute{
				MarkdownDescription: "Identifier. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_websupport_api_secret": schema.StringAttribute{
				MarkdownDescription: "Secret key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_wedos_username": schema.StringAttribute{
				MarkdownDescription: "First activate the WAPI interface in WEDOS administration and set the WAPI password. Set the protocol to JSON and add the public IP addresses of this OPNsense instance to the allowed IP addresses. Note that requests from any other IP are rejected, so whitelist the address the server actually uses for outgoing connections — on dual-stack OPNsense this is usually the IPv6 address. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_wedos_wapipass": schema.StringAttribute{
				MarkdownDescription: "WAPI Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_world4you_username": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_world4you_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_aurora_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_aurora_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_conoha_user": schema.StringAttribute{
				MarkdownDescription: "User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_conoha_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_conoha_tenantid": schema.StringAttribute{
				MarkdownDescription: "Tenant ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_conoha_idapi": schema.StringAttribute{
				MarkdownDescription: "Identity Service API. Defaults to `https://identity.xxxx.conoha.io/v2.0`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("https://identity.xxxx.conoha.io/v2.0"),
			},
			"dns_constellix_key": schema.StringAttribute{
				MarkdownDescription: "Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_constellix_secret": schema.StringAttribute{
				MarkdownDescription: "Secret. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_exoscale_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_exoscale_secret": schema.StringAttribute{
				MarkdownDescription: "Secret Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_internetbs_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_internetbs_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_pointhq_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_pointhq_email": schema.StringAttribute{
				MarkdownDescription: "E-Mail. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_rackspace_user": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_rackspace_key": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_rage4_token": schema.StringAttribute{
				MarkdownDescription: "Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_rage4_user": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_rcode0_api_token": schema.StringAttribute{
				MarkdownDescription: "The ACME API token, it can be generated at https://my.rcodezero.at/enableapi. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_rcode0_url": schema.StringAttribute{
				MarkdownDescription: "The URL of the RcodeZero API. Usually there is no need to change this value. Defaults to `https://my.rcodezero.at`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("https://my.rcodezero.at"),
			},
			"dns_rcode0_ttl": schema.Int64Attribute{
				MarkdownDescription: "The TTL in seconds for the TXT record that is created for validation. Defaults to `60`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
			},
			"dns_mijnhost_api_key": schema.StringAttribute{
				MarkdownDescription: "Please refer to the API documentation for further information. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_scaleway_token": schema.StringAttribute{
				MarkdownDescription: "API Key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"dns_zoneedit_id": schema.StringAttribute{
				MarkdownDescription: "ZoneEdit ID. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dns_zoneedit_token": schema.StringAttribute{
				MarkdownDescription: "ZoneEdit Token. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the ACME challenge type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func validationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "ACME challenge types (**Services ▸ ACME Client ▸ Challenge Types**, plugin `os-acme-client`): HTTP-01, TLS-ALPN-01 or DNS-01 with one of the acme.sh DNS providers. Only the attributes of the chosen `dns_service` are used.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME challenge type.",
				Required:            true,
			},
			"enabled":                        dschema.BoolAttribute{MarkdownDescription: "Enable this validation.", Computed: true},
			"name":                           dschema.StringAttribute{MarkdownDescription: "Name to identify this validation.", Computed: true},
			"description":                    dschema.StringAttribute{MarkdownDescription: "Description for this validation.", Computed: true},
			"method":                         dschema.StringAttribute{MarkdownDescription: "Set the ACME challenge type. You'll have to add configuration for the selected challenge type below. One of `http01`, `dns01`, `tlsalpn01`.", Computed: true},
			"http_service":                   dschema.StringAttribute{MarkdownDescription: "HTTP Service. One of `opnsense`, `haproxy`.", Computed: true},
			"http_opn_autodiscovery":         dschema.BoolAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose this option to let OPNsense try to auto-discover these IP addresses. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true},
			"http_opn_interface":             dschema.StringAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose the interface where this IP address is currently configured. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true},
			"http_opn_ipaddresses":           dschema.SetAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to one or more official IP addresses. Enter the all of these IP addresses here. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with these IP addresses.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true, ElementType: types.StringType},
			"http_haproxy_inject":            dschema.BoolAttribute{MarkdownDescription: "Automatically inject config into the local HAProxy instance to let it serve acme challanges without service interruption. Of course, adding the configuration requires a short restart of the HAProxy service.", Computed: true},
			"http_haproxy_frontends":         dschema.SetAttribute{MarkdownDescription: "Choose the local HAProxy frontends. They will automatically be configured to redirect acme challenges to the internal acme client. The HAProxy service will automatically be restarted if a certificate was renewed.", Computed: true, ElementType: types.StringType},
			"tlsalpn_service":                dschema.StringAttribute{MarkdownDescription: "TLS-ALPN Service. One of `acme`.", Computed: true},
			"tlsalpn_acme_autodiscovery":     dschema.BoolAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose this option to let OPNsense try to auto-discover these IP addresses. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true},
			"tlsalpn_acme_interface":         dschema.StringAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to an official IP address. Choose the interface where this IP address is currently configured. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with this IP address.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true},
			"tlsalpn_acme_ipaddresses":       dschema.SetAttribute{MarkdownDescription: "The FQDN's used in your certificate must currently point to one or more official IP addresses. Enter the all of these IP addresses here. OPNsense will automatically create a temporary port forward to allow the ACME validation to succeed. This will lead to a short downtime of the service that is normally used with these IP addresses.NOTE:This will ONLY work if the official IP addresses are LOCALLY configured on your OPNsense firewall.", Computed: true, ElementType: types.StringType},
			"dns_service":                    dschema.StringAttribute{MarkdownDescription: "DNS Service. One of `dns_1984hosting`, `dns_acmedns`, `dns_acmeproxy`, `dns_active24`, `dns_ad`, `dns_ali`, `dns_kas`, `dns_arvan`, `dns_artfiles`, `dns_aurora`, `dns_autodns`, `dns_aws`, `dns_azure`, `dns_bhosted`, `dns_bunny`, `dns_cloudns`, `dns_cf`, `dns_cx`, `dns_cn`, `dns_conoha`, `dns_constellix`, `dns_cpanel`, `dns_cyon`, `dns_ddnss`, `dns_desec`, `dns_dgon`, `dns_da`, `dns_dnsexit`, `dns_dnshome`, `dns_dnsimple`, `dns_dnsmint`, `dns_dnsservices`, `dns_domeneshop`, `dns_me`, `dns_dp`, `dns_doapi`, `dns_do`, `dns_dreamhost`, `dns_duckdns`, `dns_dyn`, `dns_dynu`, `dns_dynv6`, `dns_easydns`, `dns_edgecenter`, `dns_euserv`, `dns_exoscale`, `dns_fornex`, `dns_freedns`, `dns_freemyip`, `dns_gandi_livedns`, `dns_gcore`, `dns_gd`, `dns_gcloud`, `dns_googledomains`, `dns_gdnsdk`, `dns_hetzner`, `dns_hetznercloud`, `dns_hexonet`, `dns_hostingde`, `dns_hostinger`, `dns_he`, `dns_he_ddns`, `dns_infoblox`, `dns_infomaniak`, `dns_internetbs`, `dns_inwx`, `dns_ionos`, `dns_ipv64`, `dns_ispconfig`, `dns_jd`, `dns_joker`, `dns_kinghost`, `dns_knot`, `dns_leaseweb`, `dns_lexicon`, `dns_limacity`, `dns_linode`, `dns_linode_v4`, `dns_loopia`, `dns_lua`, `dns_miab`, `dns_mijnhost`, `dns_mydnsjp`, `dns_myloc`, `dns_mythic_beasts`, `dns_namecom`, `dns_namecheap`, `dns_namesilo`, `dns_nederhost`, `dns_netcup`, `dns_nic`, `dns_njalla`, `dns_nsone`, `dns_nsupdate`, `dns_online`, `dns_openprovider_rest`, `dns_opnsense`, `dns_oci`, `dns_ovh`, `dns_pdns`, `dns_pleskxml`, `dns_pointhq`, `dns_porkbun`, `dns_rackspace`, `dns_rage4`, `dns_rcode0`, `dns_regru`, `dns_scaleway`, `dns_schlundtech`, `dns_selectel`, `dns_selfhost`, `dns_servercow`, `dns_simply`, `dns_spaceship`, `dns_technitium`, `dns_timeweb`, `dns_transip`, `dns_udr`, `dns_unoeuro`, `dns_variomedia`, `dns_vscale`, `dns_vultr`, `dns_websupport`, `dns_wedos`, `dns_world4you`, `dns_yandex`, `dns_zilore`, `dns_zone`, `dns_zoneedit`, `dns_zonomi`.", Computed: true},
			"dns_sleep":                      dschema.Int64Attribute{MarkdownDescription: "The time in seconds to wait for all the TXT records to take effect after adding them to the DNS API. Defaults to 0 seconds, which causes Acme Client to check public DNS services every 10 seconds for up to 20 minutes. If set to a non-zero value, a fixed DNS sleep time will be used and the local DNS servers will be queried instead. A DNS sleep time of 120 seconds or more is recommended for some DNS APIs.", Computed: true},
			"dns_active24_api_key":           dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_active24_api_secret":        dschema.StringAttribute{MarkdownDescription: "API Secret.", Computed: true, Sensitive: true},
			"dns_ad_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_ali_key":                    dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_ali_secret":                 dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_autodns_user":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_autodns_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_autodns_context":            dschema.StringAttribute{MarkdownDescription: "Context.", Computed: true},
			"dns_aws_id":                     dschema.StringAttribute{MarkdownDescription: "AWS ID.", Computed: true},
			"dns_aws_secret":                 dschema.StringAttribute{MarkdownDescription: "AWS Secret.", Computed: true, Sensitive: true},
			"dns_azuredns_subscriptionid":    dschema.StringAttribute{MarkdownDescription: "Subscription ID.", Computed: true},
			"dns_azuredns_tenantid":          dschema.StringAttribute{MarkdownDescription: "Tenant ID.", Computed: true},
			"dns_azuredns_appid":             dschema.StringAttribute{MarkdownDescription: "APP ID.", Computed: true},
			"dns_azuredns_clientsecret":      dschema.StringAttribute{MarkdownDescription: "Client Secret.", Computed: true, Sensitive: true},
			"dns_azuredns_managedidentity":   dschema.BoolAttribute{MarkdownDescription: "When System Assigned Managed Identity is enabled the Tenant ID, APP ID and Client Secret settings are ignored by the acme client. Access tokens are obtained using the Azure Instance Metadata Service for the System Assigned Managed Identity. See documentation.", Computed: true},
			"dns_bhosted_user":               dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_bhosted_password":           dschema.StringAttribute{MarkdownDescription: "API password (MD5 hash like bHosted.nl web services example).", Computed: true, Sensitive: true},
			"dns_bhosted_ttl":                dschema.StringAttribute{MarkdownDescription: "Optional override TTL for TXT record (default: 300).", Computed: true},
			"dns_bhosted_sld":                dschema.StringAttribute{MarkdownDescription: "Optional override (useful for multi-part TLDs like co.uk).", Computed: true},
			"dns_bhosted_tld":                dschema.StringAttribute{MarkdownDescription: "Optional override (useful for multi-part TLDs like co.uk).", Computed: true},
			"dns_bunny_api_key":              dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_cf_email":                   dschema.StringAttribute{MarkdownDescription: "Use your CF account with full privileges. This is less secure, a API token should be used instead (see below).", Computed: true},
			"dns_cf_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_cf_token":                   dschema.StringAttribute{MarkdownDescription: "The token needs \"Read\" access to Zone.Zone and \"Edit\" access to Zone.DNS across all zones from an account.", Computed: true, Sensitive: true},
			"dns_cf_account_id":              dschema.StringAttribute{MarkdownDescription: "Can be found in the URI after logging into the Cloudflare dashboard.", Computed: true},
			"dns_cf_zone_id":                 dschema.StringAttribute{MarkdownDescription: "Note that specifying a Zone ID will limit this configuration to a single domain.", Computed: true},
			"dns_cloudns_auth_id":            dschema.StringAttribute{MarkdownDescription: "Auth ID.", Computed: true},
			"dns_cloudns_sub_auth_id":        dschema.StringAttribute{MarkdownDescription: "Sub Auth ID.", Computed: true},
			"dns_cloudns_auth_password":      dschema.StringAttribute{MarkdownDescription: "Auth Password.", Computed: true, Sensitive: true},
			"dns_cx_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_cx_secret":                  dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_cyon_user":                  dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_cyon_password":              dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_da_key":                     dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_da_insecure":                dschema.BoolAttribute{MarkdownDescription: "Uncheck this box if you have a valid SSL certificate for your DirectAdmin installation.", Computed: true},
			"dns_ddnss_token":                dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_dgon_key":                   dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_dnsexit_auth_user":          dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_dnsexit_auth_pass":          dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_dnsexit_api":                dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true},
			"dns_dnshome_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_dnshome_subdomain":          dschema.StringAttribute{MarkdownDescription: "Subdomain.", Computed: true},
			"dns_dnsimple_token":             dschema.StringAttribute{MarkdownDescription: "Note that this is the account token not the user token.", Computed: true, Sensitive: true},
			"dns_dnsservices_user":           dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_dnsservices_password":       dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_doapi_token":                dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_do_pid":                     dschema.StringAttribute{MarkdownDescription: "Partner ID.", Computed: true},
			"dns_do_password":                dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_domeneshop_token":           dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_domeneshop_secret":          dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_dp_id":                      dschema.StringAttribute{MarkdownDescription: "ID.", Computed: true},
			"dns_dp_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_dh_key":                     dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_duckdns_token":              dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_dyn_customer":               dschema.StringAttribute{MarkdownDescription: "Customer.", Computed: true},
			"dns_dyn_user":                   dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_dyn_password":               dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_dynu_clientid":              dschema.StringAttribute{MarkdownDescription: "Client ID.", Computed: true},
			"dns_dynu_secret":                dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_freedns_user":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_freedns_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_freemyip_token":             dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_fornex_api_key":             dschema.StringAttribute{MarkdownDescription: "ApiKey.", Computed: true, Sensitive: true},
			"dns_gandi_livedns_key":          dschema.StringAttribute{MarkdownDescription: "The API key is the previous mechanism that was replaced with Personal Access Tokens. API keys should no longer be used.", Computed: true, Sensitive: true},
			"dns_gandi_livedns_token":        dschema.StringAttribute{MarkdownDescription: "Personal Access Token.", Computed: true, Sensitive: true},
			"dns_gcloud_key":                 dschema.StringAttribute{MarkdownDescription: "Provide a service account key in JSON format for your Google Cloud account.", Computed: true, Sensitive: true},
			"dns_gcore_token":                dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_googledomains_access_token": dschema.StringAttribute{MarkdownDescription: "Access Token.", Computed: true, Sensitive: true},
			"dns_googledomains_zone":         dschema.StringAttribute{MarkdownDescription: "Zone.", Computed: true},
			"dns_gd_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_gd_secret":                  dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_hostingde_server":           dschema.StringAttribute{MarkdownDescription: "Enter the API endpoint, e.g. https://secure.hosting.de or https://partner.http.net.", Computed: true},
			"dns_hostingde_api_key":          dschema.StringAttribute{MarkdownDescription: "API key.", Computed: true, Sensitive: true},
			"dns_hostinger_token":            dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_he_user":                    dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_he_password":                dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_he_ddns_key":                dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_infoblox_credentials":       dschema.StringAttribute{MarkdownDescription: "Use the following format: username:password.", Computed: true},
			"dns_infoblox_server":            dschema.StringAttribute{MarkdownDescription: "Enter either the IP address or FQDN of your Infoblox appliance.", Computed: true},
			"dns_inwx_user":                  dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_inws_password":              dschema.StringAttribute{MarkdownDescription: "dns_inws_password.", Computed: true, Sensitive: true},
			"dns_inwx_password":              dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_inwx_shared_secret":         dschema.StringAttribute{MarkdownDescription: "When 2FA is enabled, the Shared Secret must be provided. Note that this feature requires the package oath-toolkit, which must be installed manually.", Computed: true, Sensitive: true},
			"dns_ionos_prefix":               dschema.StringAttribute{MarkdownDescription: "Prefix.", Computed: true},
			"dns_ionos_secret":               dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_ipv64_token":                dschema.StringAttribute{MarkdownDescription: "IPv64.net API Token.", Computed: true, Sensitive: true},
			"dns_ispconfig_user":             dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_ispconfig_password":         dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_ispconfig_api":              dschema.StringAttribute{MarkdownDescription: "API URL.", Computed: true},
			"dns_ispconfig_insecure":         dschema.BoolAttribute{MarkdownDescription: "Uncheck this box if you have a valid SSL certificate for your ISPConfig installation.", Computed: true},
			"dns_jd_id":                      dschema.StringAttribute{MarkdownDescription: "API key id.", Computed: true},
			"dns_jd_region":                  dschema.StringAttribute{MarkdownDescription: "JD Cloud region.", Computed: true},
			"dns_jd_secret":                  dschema.StringAttribute{MarkdownDescription: "API key secret.", Computed: true, Sensitive: true},
			"dns_joker_username":             dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_joker_password":             dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_kinghost_username":          dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_kinghost_password":          dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_knot_server":                dschema.StringAttribute{MarkdownDescription: "Knot server address, like: dns.example.com.", Computed: true},
			"dns_knot_key":                   dschema.StringAttribute{MarkdownDescription: "Specify the location of the generated TSIG Key inside the TSIG file using grep and cut, example: grep \\# /etc/knot/acme.key | cut -d' ' -f2.", Computed: true, Sensitive: true},
			"dns_lexicon_provider":           dschema.StringAttribute{MarkdownDescription: "Provider. One of `aliyun`, `aurora`, `auto`, `azure`, `cloudflare`, `cloudns`, `cloudxns`, `conoha`, `constellix`, `digitalocean`, `dinahosting`, `directadmin`, `dnsimple`, `dnsmadeeasy`, `dnspark`, `dnspod`, `dreamhost`, `easydns`, `easyname`, `exoscale`, `gandi`, `gehirn`, `glesys`, `godaddy`, `googleclouddns`, `gratisdns`, `henet`, `hetzner`, `hover`, `infoblox`, `internetbs`, `inwx`, `linode`, `linode4`, `localzone`, `luadns`, `memset`, `namecheap`, `namesilo`, `netcup`, `nfsn`, `nsone`, `onapp`, `online`, `ovh`, `plesk`, `pointhq`, `powerdns`, `rackspace`, `rage4`, `route53`, `safedns`, `sakuracloud`, `softlayer`, `subreg`, `transip`, `vultr`, `yandex`, `zeit`, `zilore`, `zonomi`.", Computed: true},
			"dns_lexicon_user":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_lexicon_token":              dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_limacity_apikey":            dschema.StringAttribute{MarkdownDescription: "API-Key.", Computed: true, Sensitive: true},
			"dns_linode_key":                 dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_linode_v4_key":              dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_loopia_api":                 dschema.StringAttribute{MarkdownDescription: "API Endpoint.", Computed: true},
			"dns_loopia_user":                dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_loopia_password":            dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_lua_email":                  dschema.StringAttribute{MarkdownDescription: "E-Mail.", Computed: true},
			"dns_lua_key":                    dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_miab_user":                  dschema.StringAttribute{MarkdownDescription: "User with admin priviledges.", Computed: true},
			"dns_miab_password":              dschema.StringAttribute{MarkdownDescription: "Password for admin user.", Computed: true, Sensitive: true},
			"dns_miab_server":                dschema.StringAttribute{MarkdownDescription: "MailinaBox Server FQDN.", Computed: true},
			"dns_me_key":                     dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_me_secret":                  dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_mydnsjp_masterid":           dschema.StringAttribute{MarkdownDescription: "Master ID.", Computed: true},
			"dns_mydnsjp_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_myloc_token":                dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_mythic_beasts_key":          dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_mythic_beasts_secret":       dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_namecheap_user":             dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_namecheap_api":              dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true},
			"dns_namecheap_sourceip":         dschema.StringAttribute{MarkdownDescription: "Source IP.", Computed: true},
			"dns_namecom_user":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_namecom_token":              dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_namesilo_key":               dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_nederhost_key":              dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true},
			"dns_netcup_cid":                 dschema.StringAttribute{MarkdownDescription: "Customer number.", Computed: true},
			"dns_netcup_key":                 dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_netcup_pw":                  dschema.StringAttribute{MarkdownDescription: "API Password.", Computed: true},
			"dns_njalla_token":               dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_nsone_key":                  dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_nsupdate_server":            dschema.StringAttribute{MarkdownDescription: "Server (FQDN).", Computed: true},
			"dns_nsupdate_zone":              dschema.StringAttribute{MarkdownDescription: "Optionally set the name of the hosted zone (e.g. example.com) as some DNS Providers require, like dyn.com's 'Standard DNS'.", Computed: true},
			"dns_nsupdate_key":               dschema.StringAttribute{MarkdownDescription: "Requires the whole key file in a format that is compatible with nsupdate.", Computed: true, Sensitive: true},
			"dns_oci_cli_user":               dschema.StringAttribute{MarkdownDescription: "OCID of the user calling the API.", Computed: true},
			"dns_oci_cli_tenancy":            dschema.StringAttribute{MarkdownDescription: "OCID of your tenancy.", Computed: true},
			"dns_oci_cli_region":             dschema.StringAttribute{MarkdownDescription: "Your OCI region.", Computed: true},
			"dns_oci_cli_key":                dschema.StringAttribute{MarkdownDescription: "Please refer to the acme.sh documentation for further information.", Computed: true, Sensitive: true},
			"dns_online_key":                 dschema.StringAttribute{MarkdownDescription: "API key.", Computed: true, Sensitive: true},
			"dns_openprovider_rest_username": dschema.StringAttribute{MarkdownDescription: "Account Username.", Computed: true},
			"dns_openprovider_rest_password": dschema.StringAttribute{MarkdownDescription: "Account Password.", Computed: true, Sensitive: true},
			"dns_opnsense_host":              dschema.StringAttribute{MarkdownDescription: "OPNSense Server (FQDN).", Computed: true},
			"dns_opnsense_port":              dschema.StringAttribute{MarkdownDescription: "OPNSense Server Port.", Computed: true},
			"dns_opnsense_key":               dschema.StringAttribute{MarkdownDescription: "User API key.", Computed: true, Sensitive: true},
			"dns_opnsense_token":             dschema.StringAttribute{MarkdownDescription: "User API token.", Computed: true, Sensitive: true},
			"dns_opnsense_insecure":          dschema.BoolAttribute{MarkdownDescription: "Uncheck this box if you have a valid certificate for your OPNsense interface.", Computed: true},
			"dns_ovh_app_key":                dschema.StringAttribute{MarkdownDescription: "Application Key.", Computed: true, Sensitive: true},
			"dns_ovh_app_secret":             dschema.StringAttribute{MarkdownDescription: "Application Secret.", Computed: true, Sensitive: true},
			"dns_ovh_consumer_key":           dschema.StringAttribute{MarkdownDescription: "Consumer Key.", Computed: true, Sensitive: true},
			"dns_ovh_endpoint":               dschema.StringAttribute{MarkdownDescription: "Specify the OVH endpoint, i.e. ovh-eu, ovh-ca, kimsufi-eu, etc. Please refer to the acme.sh documentation for further information.", Computed: true},
			"dns_pleskxml_user":              dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_pleskxml_pass":              dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_pleskxml_uri":               dschema.StringAttribute{MarkdownDescription: "Plesk XML URIs often look similar to: https://my-plesk-site.com:8443/enterprise/control/agent.php.", Computed: true},
			"dns_pdns_url":                   dschema.StringAttribute{MarkdownDescription: "Specify the URL for your PowerDNS server, i.e. http://ns.example.com:8081.", Computed: true},
			"dns_pdns_serverid":              dschema.StringAttribute{MarkdownDescription: "Specify the Server ID of your PowerDNS server, i.e. localhost.", Computed: true},
			"dns_pdns_token":                 dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_porkbun_key":                dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_porkbun_secret":             dschema.StringAttribute{MarkdownDescription: "Secret API Key.", Computed: true, Sensitive: true},
			"dns_sl_key":                     dschema.StringAttribute{MarkdownDescription: "API key (deprecated).", Computed: true, Sensitive: true},
			"dns_sl_apiver":                  dschema.StringAttribute{MarkdownDescription: "API Version. One of `v1`, `v2`.", Computed: true},
			"dns_sl_token_lifetime":          dschema.StringAttribute{MarkdownDescription: "Token lifetime in minutes (0-1440).", Computed: true, Sensitive: true},
			"dns_sl_account_id":              dschema.StringAttribute{MarkdownDescription: "The account number can be found on the Selectel control panel.", Computed: true},
			"dns_sl_project_name":            dschema.StringAttribute{MarkdownDescription: "Project name.", Computed: true},
			"dns_sl_login_name":              dschema.StringAttribute{MarkdownDescription: "The service username can be found in the Selectel control panel.", Computed: true},
			"dns_sl_password":                dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_selfhost_apikey":            dschema.StringAttribute{MarkdownDescription: "Create a DNS-01 API key in the Selfhost portal (menu \"Entwickler\"). The legacy username/password authentication is no longer supported by Selfhost.", Computed: true, Sensitive: true},
			"dns_selfhost_map":               dschema.StringAttribute{MarkdownDescription: "Please create the TXT record with subdomain _acme-challenge first, than you can get the ID from the edit page. Up to two RIDs per fulldomain are supported but at least one must be set, e.g. _acme-challenge.domain.net:RID:RID2.", Computed: true},
			"dns_servercow_username":         dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_servercow_password":         dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_simply_api_key":             dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_simply_account_name":        dschema.StringAttribute{MarkdownDescription: "Account Name.", Computed: true},
			"dns_spaceship_api_key":          dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_spaceship_api_secret":       dschema.StringAttribute{MarkdownDescription: "API Secret.", Computed: true, Sensitive: true},
			"dns_spaceship_root_domain":      dschema.StringAttribute{MarkdownDescription: "Root Domain (Optional).", Computed: true},
			"dns_technitium_token":           dschema.StringAttribute{MarkdownDescription: "Generate a token from Technitium user and grant rights to the necessary zones.", Computed: true, Sensitive: true},
			"dns_technitium_hostname":        dschema.StringAttribute{MarkdownDescription: "The full server URL, for example https://technitium:5380. The port number is optional.NOTE: Remove the trailing / if it is present.", Computed: true},
			"dns_transip_username":           dschema.StringAttribute{MarkdownDescription: "Your TransIP username.", Computed: true},
			"dns_transip_key":                dschema.StringAttribute{MarkdownDescription: "Requires the whole key file in a format that is compatible with TransIP.", Computed: true, Sensitive: true},
			"dns_transip_token_global_key":   dschema.BoolAttribute{MarkdownDescription: "When using a global key, IP whitelisting is disabled.", Computed: true},
			"dns_timeweb_token":              dschema.StringAttribute{MarkdownDescription: "Timeweb Cloud API Token.", Computed: true, Sensitive: true},
			"dns_udr_user":                   dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_udr_password":               dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_uno_key":                    dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_uno_user":                   dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_vscale_key":                 dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_vultr_key":                  dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_yandex_token":               dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_zilore_key":                 dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_zm_key":                     dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_gdnsdk_user":                dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_gdnsdk_password":            dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_acmedns_user":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_acmedns_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_acmedns_subdomain":          dschema.StringAttribute{MarkdownDescription: "Subdomain.", Computed: true},
			"dns_acmedns_updateurl":          dschema.StringAttribute{MarkdownDescription: "dns_acmedns_updateurl.", Computed: true},
			"dns_acmedns_baseurl":            dschema.StringAttribute{MarkdownDescription: "Specify the custom ACME DNS URL, i.e. https://auth.acme-dns.io:443 (optional).", Computed: true},
			"dns_acmeproxy_endpoint":         dschema.StringAttribute{MarkdownDescription: "Specify the acmeproxy endpoint URL, i.e. https://acmeproxy.example.com/.", Computed: true},
			"dns_acmeproxy_username":         dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_acmeproxy_password":         dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_variomedia_key":             dschema.StringAttribute{MarkdownDescription: "You need to obtain your API key from variomedia's customer support.", Computed: true, Sensitive: true},
			"dns_schlundtech_user":           dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_schlundtech_password":       dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_easydns_apitoken":           dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_easydns_apikey":             dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_edgecenter_apikey":          dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_euserv_user":                dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_euserv_password":            dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_leaseweb_key":               dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_cn_user":                    dschema.StringAttribute{MarkdownDescription: "API User.", Computed: true},
			"dns_cn_password":                dschema.StringAttribute{MarkdownDescription: "API Password.", Computed: true, Sensitive: true},
			"dns_arvan_token":                dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_artfiles_username":          dschema.StringAttribute{MarkdownDescription: "API Username.", Computed: true},
			"dns_artfiles_password":          dschema.StringAttribute{MarkdownDescription: "API Password.", Computed: true, Sensitive: true},
			"dns_hetzner_token":              dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_hetznercloud_token":         dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_hexonet_login":              dschema.StringAttribute{MarkdownDescription: "This should be a combination of the username and role ID. For example, when the username is \"myuser\" and role ID is \"testrole\", then \"myuser!testrole\" must be used here.", Computed: true},
			"dns_hexonet_password":           dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_1984hosting_user":           dschema.StringAttribute{MarkdownDescription: "API User.", Computed: true},
			"dns_1984hosting_password":       dschema.StringAttribute{MarkdownDescription: "API Password.", Computed: true, Sensitive: true},
			"dns_1984hosting_totp_secret":    dschema.StringAttribute{MarkdownDescription: "When 2FA is enabled, the Base32 TOTP secret must be provided. Note that this feature requires the package oath-toolkit, which must be installed manually.", Computed: true, Sensitive: true},
			"dns_kas_login":                  dschema.StringAttribute{MarkdownDescription: "KAS Login.", Computed: true},
			"dns_kas_authdata":               dschema.StringAttribute{MarkdownDescription: "KAS Authdata.", Computed: true},
			"dns_kas_authtype":               dschema.StringAttribute{MarkdownDescription: "KAS Authtype. One of `plain`, `sha1`.", Computed: true},
			"dns_dnsmint_api_key":            dschema.StringAttribute{MarkdownDescription: "DNSMint API key.", Computed: true, Sensitive: true},
			"dns_desec_token":                dschema.StringAttribute{MarkdownDescription: "deSEC token.", Computed: true, Sensitive: true},
			"dns_desec_name":                 dschema.StringAttribute{MarkdownDescription: "deSEC domain.", Computed: true},
			"dns_infomaniak_token":           dschema.StringAttribute{MarkdownDescription: "API Token.", Computed: true, Sensitive: true},
			"dns_zone_username":              dschema.StringAttribute{MarkdownDescription: "Zone Username.", Computed: true},
			"dns_zone_key":                   dschema.StringAttribute{MarkdownDescription: "Can be generated under ZoneID account management.", Computed: true, Sensitive: true},
			"dns_dynv6_token":                dschema.StringAttribute{MarkdownDescription: "HTTP Token.", Computed: true, Sensitive: true},
			"dns_cpanel_user":                dschema.StringAttribute{MarkdownDescription: "cPanel Username.", Computed: true},
			"dns_cpanel_token":               dschema.StringAttribute{MarkdownDescription: "Please refer to the cPanel documentation for further information.", Computed: true, Sensitive: true},
			"dns_cpanel_hostname":            dschema.StringAttribute{MarkdownDescription: "EX: https://hostname:2083.", Computed: true},
			"dns_regru_username":             dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_regru_password":             dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_nic_username":               dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_nic_password":               dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_nic_client":                 dschema.StringAttribute{MarkdownDescription: "ClientId.", Computed: true},
			"dns_nic_secret":                 dschema.StringAttribute{MarkdownDescription: "ClientSecret.", Computed: true, Sensitive: true},
			"dns_websupport_api_key":         dschema.StringAttribute{MarkdownDescription: "Identifier.", Computed: true, Sensitive: true},
			"dns_websupport_api_secret":      dschema.StringAttribute{MarkdownDescription: "Secret key.", Computed: true, Sensitive: true},
			"dns_wedos_username":             dschema.StringAttribute{MarkdownDescription: "First activate the WAPI interface in WEDOS administration and set the WAPI password. Set the protocol to JSON and add the public IP addresses of this OPNsense instance to the allowed IP addresses. Note that requests from any other IP are rejected, so whitelist the address the server actually uses for outgoing connections — on dual-stack OPNsense this is usually the IPv6 address.", Computed: true},
			"dns_wedos_wapipass":             dschema.StringAttribute{MarkdownDescription: "WAPI Password.", Computed: true, Sensitive: true},
			"dns_world4you_username":         dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_world4you_password":         dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_aurora_key":                 dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_aurora_secret":              dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_conoha_user":                dschema.StringAttribute{MarkdownDescription: "User.", Computed: true},
			"dns_conoha_password":            dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_conoha_tenantid":            dschema.StringAttribute{MarkdownDescription: "Tenant ID.", Computed: true},
			"dns_conoha_idapi":               dschema.StringAttribute{MarkdownDescription: "Identity Service API.", Computed: true},
			"dns_constellix_key":             dschema.StringAttribute{MarkdownDescription: "Key.", Computed: true, Sensitive: true},
			"dns_constellix_secret":          dschema.StringAttribute{MarkdownDescription: "Secret.", Computed: true, Sensitive: true},
			"dns_exoscale_key":               dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_exoscale_secret":            dschema.StringAttribute{MarkdownDescription: "Secret Key.", Computed: true, Sensitive: true},
			"dns_internetbs_key":             dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_internetbs_password":        dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"dns_pointhq_key":                dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_pointhq_email":              dschema.StringAttribute{MarkdownDescription: "E-Mail.", Computed: true},
			"dns_rackspace_user":             dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_rackspace_key":              dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_rage4_token":                dschema.StringAttribute{MarkdownDescription: "Token.", Computed: true, Sensitive: true},
			"dns_rage4_user":                 dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"dns_rcode0_api_token":           dschema.StringAttribute{MarkdownDescription: "The ACME API token, it can be generated at https://my.rcodezero.at/enableapi.", Computed: true, Sensitive: true},
			"dns_rcode0_url":                 dschema.StringAttribute{MarkdownDescription: "The URL of the RcodeZero API. Usually there is no need to change this value.", Computed: true},
			"dns_rcode0_ttl":                 dschema.Int64Attribute{MarkdownDescription: "The TTL in seconds for the TXT record that is created for validation.", Computed: true},
			"dns_mijnhost_api_key":           dschema.StringAttribute{MarkdownDescription: "Please refer to the API documentation for further information.", Computed: true, Sensitive: true},
			"dns_scaleway_token":             dschema.StringAttribute{MarkdownDescription: "API Key.", Computed: true, Sensitive: true},
			"dns_zoneedit_id":                dschema.StringAttribute{MarkdownDescription: "ZoneEdit ID.", Computed: true},
			"dns_zoneedit_token":             dschema.StringAttribute{MarkdownDescription: "ZoneEdit Token.", Computed: true, Sensitive: true},
		},
	}
}

func convertValidationSchemaToStruct(d *validationResourceModel) (*acmeclient.Validation, error) {
	return &acmeclient.Validation{
		Enabled:                     tools.BoolToString(d.Enabled.ValueBool()),
		Name:                        d.Name.ValueString(),
		Description:                 d.Description.ValueString(),
		Method:                      api.SelectedMap(d.Method.ValueString()),
		HttpService:                 api.SelectedMap(d.HttpService.ValueString()),
		HttpOpnAutodiscovery:        tools.BoolToString(d.HttpOpnAutodiscovery.ValueBool()),
		HttpOpnInterface:            api.SelectedMap(d.HttpOpnInterface.ValueString()),
		HttpOpnIpaddresses:          api.SelectedMapList(setToSlice(d.HttpOpnIpaddresses)),
		HttpHaproxyInject:           tools.BoolToString(d.HttpHaproxyInject.ValueBool()),
		HttpHaproxyFrontends:        api.SelectedMapList(setToSlice(d.HttpHaproxyFrontends)),
		TlsalpnService:              api.SelectedMap(d.TlsalpnService.ValueString()),
		TlsalpnAcmeAutodiscovery:    tools.BoolToString(d.TlsalpnAcmeAutodiscovery.ValueBool()),
		TlsalpnAcmeInterface:        api.SelectedMap(d.TlsalpnAcmeInterface.ValueString()),
		TlsalpnAcmeIpaddresses:      api.SelectedMapList(setToSlice(d.TlsalpnAcmeIpaddresses)),
		DnsService:                  api.SelectedMap(d.DNSService.ValueString()),
		DnsSleep:                    tools.Int64ToString(d.DNSSleep.ValueInt64()),
		DnsActive24ApiKey:           d.DNSActive24ApiKey.ValueString(),
		DnsActive24ApiSecret:        d.DNSActive24ApiSecret.ValueString(),
		DnsAdKey:                    d.DNSAdKey.ValueString(),
		DnsAliKey:                   d.DNSAliKey.ValueString(),
		DnsAliSecret:                d.DNSAliSecret.ValueString(),
		DnsAutodnsUser:              d.DNSAutodnsUser.ValueString(),
		DnsAutodnsPassword:          d.DNSAutodnsPassword.ValueString(),
		DnsAutodnsContext:           d.DNSAutodnsContext.ValueString(),
		DnsAwsId:                    d.DNSAwsID.ValueString(),
		DnsAwsSecret:                d.DNSAwsSecret.ValueString(),
		DnsAzurednsSubscriptionid:   d.DNSAzurednsSubscriptionid.ValueString(),
		DnsAzurednsTenantid:         d.DNSAzurednsTenantid.ValueString(),
		DnsAzurednsAppid:            d.DNSAzurednsAppid.ValueString(),
		DnsAzurednsClientsecret:     d.DNSAzurednsClientsecret.ValueString(),
		DnsAzurednsManagedidentity:  tools.BoolToString(d.DNSAzurednsManagedidentity.ValueBool()),
		DnsBhostedUser:              d.DNSBhostedUser.ValueString(),
		DnsBhostedPassword:          d.DNSBhostedPassword.ValueString(),
		DnsBhostedTtl:               d.DNSBhostedTTL.ValueString(),
		DnsBhostedSld:               d.DNSBhostedSld.ValueString(),
		DnsBhostedTld:               d.DNSBhostedTld.ValueString(),
		DnsBunnyApiKey:              d.DNSBunnyApiKey.ValueString(),
		DnsCfEmail:                  d.DNSCfEmail.ValueString(),
		DnsCfKey:                    d.DNSCfKey.ValueString(),
		DnsCfToken:                  d.DNSCfToken.ValueString(),
		DnsCfAccountId:              d.DNSCfAccountID.ValueString(),
		DnsCfZoneId:                 d.DNSCfZoneID.ValueString(),
		DnsCloudnsAuthId:            d.DNSCloudnsAuthID.ValueString(),
		DnsCloudnsSubAuthId:         d.DNSCloudnsSubAuthID.ValueString(),
		DnsCloudnsAuthPassword:      d.DNSCloudnsAuthPassword.ValueString(),
		DnsCxKey:                    d.DNSCxKey.ValueString(),
		DnsCxSecret:                 d.DNSCxSecret.ValueString(),
		DnsCyonUser:                 d.DNSCyonUser.ValueString(),
		DnsCyonPassword:             d.DNSCyonPassword.ValueString(),
		DnsDaKey:                    d.DNSDaKey.ValueString(),
		DnsDaInsecure:               tools.BoolToString(d.DNSDaInsecure.ValueBool()),
		DnsDdnssToken:               d.DNSDdnssToken.ValueString(),
		DnsDgonKey:                  d.DNSDgonKey.ValueString(),
		DnsDnsexitAuthUser:          d.DNSDnsexitAuthUser.ValueString(),
		DnsDnsexitAuthPass:          d.DNSDnsexitAuthPass.ValueString(),
		DnsDnsexitApi:               d.DNSDnsexitApi.ValueString(),
		DnsDnshomePassword:          d.DNSDnshomePassword.ValueString(),
		DnsDnshomeSubdomain:         d.DNSDnshomeSubdomain.ValueString(),
		DnsDnsimpleToken:            d.DNSDnsimpleToken.ValueString(),
		DnsDnsservicesUser:          d.DNSDnsservicesUser.ValueString(),
		DnsDnsservicesPassword:      d.DNSDnsservicesPassword.ValueString(),
		DnsDoapiToken:               d.DNSDoapiToken.ValueString(),
		DnsDoPid:                    d.DNSDoPid.ValueString(),
		DnsDoPassword:               d.DNSDoPassword.ValueString(),
		DnsDomeneshopToken:          d.DNSDomeneshopToken.ValueString(),
		DnsDomeneshopSecret:         d.DNSDomeneshopSecret.ValueString(),
		DnsDpId:                     d.DNSDpID.ValueString(),
		DnsDpKey:                    d.DNSDpKey.ValueString(),
		DnsDhKey:                    d.DNSDhKey.ValueString(),
		DnsDuckdnsToken:             d.DNSDuckdnsToken.ValueString(),
		DnsDynCustomer:              d.DNSDynCustomer.ValueString(),
		DnsDynUser:                  d.DNSDynUser.ValueString(),
		DnsDynPassword:              d.DNSDynPassword.ValueString(),
		DnsDynuClientid:             d.DNSDynuClientid.ValueString(),
		DnsDynuSecret:               d.DNSDynuSecret.ValueString(),
		DnsFreednsUser:              d.DNSFreednsUser.ValueString(),
		DnsFreednsPassword:          d.DNSFreednsPassword.ValueString(),
		DnsFreemyipToken:            d.DNSFreemyipToken.ValueString(),
		DnsFornexApiKey:             d.DNSFornexApiKey.ValueString(),
		DnsGandiLivednsKey:          d.DNSGandiLivednsKey.ValueString(),
		DnsGandiLivednsToken:        d.DNSGandiLivednsToken.ValueString(),
		DnsGcloudKey:                d.DNSGcloudKey.ValueString(),
		DnsGcoreToken:               d.DNSGcoreToken.ValueString(),
		DnsGoogledomainsAccessToken: d.DNSGoogledomainsAccessToken.ValueString(),
		DnsGoogledomainsZone:        d.DNSGoogledomainsZone.ValueString(),
		DnsGdKey:                    d.DNSGdKey.ValueString(),
		DnsGdSecret:                 d.DNSGdSecret.ValueString(),
		DnsHostingdeServer:          d.DNSHostingdeServer.ValueString(),
		DnsHostingdeApiKey:          d.DNSHostingdeApiKey.ValueString(),
		DnsHostingerToken:           d.DNSHostingerToken.ValueString(),
		DnsHeUser:                   d.DNSHeUser.ValueString(),
		DnsHePassword:               d.DNSHePassword.ValueString(),
		DnsHeDdnsKey:                d.DNSHeDdnsKey.ValueString(),
		DnsInfobloxCredentials:      d.DNSInfobloxCredentials.ValueString(),
		DnsInfobloxServer:           d.DNSInfobloxServer.ValueString(),
		DnsInwxUser:                 d.DNSInwxUser.ValueString(),
		DnsInwsPassword:             d.DNSInwsPassword.ValueString(),
		DnsInwxPassword:             d.DNSInwxPassword.ValueString(),
		DnsInwxSharedSecret:         d.DNSInwxSharedSecret.ValueString(),
		DnsIonosPrefix:              d.DNSIonosPrefix.ValueString(),
		DnsIonosSecret:              d.DNSIonosSecret.ValueString(),
		DnsIpv64Token:               d.DNSIpv64Token.ValueString(),
		DnsIspconfigUser:            d.DNSIspconfigUser.ValueString(),
		DnsIspconfigPassword:        d.DNSIspconfigPassword.ValueString(),
		DnsIspconfigApi:             d.DNSIspconfigApi.ValueString(),
		DnsIspconfigInsecure:        tools.BoolToString(d.DNSIspconfigInsecure.ValueBool()),
		DnsJdId:                     d.DNSJdID.ValueString(),
		DnsJdRegion:                 d.DNSJdRegion.ValueString(),
		DnsJdSecret:                 d.DNSJdSecret.ValueString(),
		DnsJokerUsername:            d.DNSJokerUsername.ValueString(),
		DnsJokerPassword:            d.DNSJokerPassword.ValueString(),
		DnsKinghostUsername:         d.DNSKinghostUsername.ValueString(),
		DnsKinghostPassword:         d.DNSKinghostPassword.ValueString(),
		DnsKnotServer:               d.DNSKnotServer.ValueString(),
		DnsKnotKey:                  d.DNSKnotKey.ValueString(),
		DnsLexiconProvider:          api.SelectedMap(d.DNSLexiconProvider.ValueString()),
		DnsLexiconUser:              d.DNSLexiconUser.ValueString(),
		DnsLexiconToken:             d.DNSLexiconToken.ValueString(),
		DnsLimacityApikey:           d.DNSLimacityApikey.ValueString(),
		DnsLinodeKey:                d.DNSLinodeKey.ValueString(),
		DnsLinodeV4Key:              d.DNSLinodeV4Key.ValueString(),
		DnsLoopiaApi:                d.DNSLoopiaApi.ValueString(),
		DnsLoopiaUser:               d.DNSLoopiaUser.ValueString(),
		DnsLoopiaPassword:           d.DNSLoopiaPassword.ValueString(),
		DnsLuaEmail:                 d.DNSLuaEmail.ValueString(),
		DnsLuaKey:                   d.DNSLuaKey.ValueString(),
		DnsMiabUser:                 d.DNSMiabUser.ValueString(),
		DnsMiabPassword:             d.DNSMiabPassword.ValueString(),
		DnsMiabServer:               d.DNSMiabServer.ValueString(),
		DnsMeKey:                    d.DNSMeKey.ValueString(),
		DnsMeSecret:                 d.DNSMeSecret.ValueString(),
		DnsMydnsjpMasterid:          d.DNSMydnsjpMasterid.ValueString(),
		DnsMydnsjpPassword:          d.DNSMydnsjpPassword.ValueString(),
		DnsMylocToken:               d.DNSMylocToken.ValueString(),
		DnsMythicBeastsKey:          d.DNSMythicBeastsKey.ValueString(),
		DnsMythicBeastsSecret:       d.DNSMythicBeastsSecret.ValueString(),
		DnsNamecheapUser:            d.DNSNamecheapUser.ValueString(),
		DnsNamecheapApi:             d.DNSNamecheapApi.ValueString(),
		DnsNamecheapSourceip:        d.DNSNamecheapSourceip.ValueString(),
		DnsNamecomUser:              d.DNSNamecomUser.ValueString(),
		DnsNamecomToken:             d.DNSNamecomToken.ValueString(),
		DnsNamesiloKey:              d.DNSNamesiloKey.ValueString(),
		DnsNederhostKey:             d.DNSNederhostKey.ValueString(),
		DnsNetcupCid:                d.DNSNetcupCid.ValueString(),
		DnsNetcupKey:                d.DNSNetcupKey.ValueString(),
		DnsNetcupPw:                 d.DNSNetcupPw.ValueString(),
		DnsNjallaToken:              d.DNSNjallaToken.ValueString(),
		DnsNsoneKey:                 d.DNSNsoneKey.ValueString(),
		DnsNsupdateServer:           d.DNSNsupdateServer.ValueString(),
		DnsNsupdateZone:             d.DNSNsupdateZone.ValueString(),
		DnsNsupdateKey:              d.DNSNsupdateKey.ValueString(),
		DnsOciCliUser:               d.DNSOciCliUser.ValueString(),
		DnsOciCliTenancy:            d.DNSOciCliTenancy.ValueString(),
		DnsOciCliRegion:             d.DNSOciCliRegion.ValueString(),
		DnsOciCliKey:                d.DNSOciCliKey.ValueString(),
		DnsOnlineKey:                d.DNSOnlineKey.ValueString(),
		DnsOpenproviderRestUsername: d.DNSOpenproviderRestUsername.ValueString(),
		DnsOpenproviderRestPassword: d.DNSOpenproviderRestPassword.ValueString(),
		DnsOpnsenseHost:             d.DNSOpnsenseHost.ValueString(),
		DnsOpnsensePort:             d.DNSOpnsensePort.ValueString(),
		DnsOpnsenseKey:              d.DNSOpnsenseKey.ValueString(),
		DnsOpnsenseToken:            d.DNSOpnsenseToken.ValueString(),
		DnsOpnsenseInsecure:         tools.BoolToString(d.DNSOpnsenseInsecure.ValueBool()),
		DnsOvhAppKey:                d.DNSOvhAppKey.ValueString(),
		DnsOvhAppSecret:             d.DNSOvhAppSecret.ValueString(),
		DnsOvhConsumerKey:           d.DNSOvhConsumerKey.ValueString(),
		DnsOvhEndpoint:              d.DNSOvhEndpoint.ValueString(),
		DnsPleskxmlUser:             d.DNSPleskxmlUser.ValueString(),
		DnsPleskxmlPass:             d.DNSPleskxmlPass.ValueString(),
		DnsPleskxmlUri:              d.DNSPleskxmlUri.ValueString(),
		DnsPdnsUrl:                  d.DNSPdnsURL.ValueString(),
		DnsPdnsServerid:             d.DNSPdnsServerid.ValueString(),
		DnsPdnsToken:                d.DNSPdnsToken.ValueString(),
		DnsPorkbunKey:               d.DNSPorkbunKey.ValueString(),
		DnsPorkbunSecret:            d.DNSPorkbunSecret.ValueString(),
		DnsSlKey:                    d.DNSSlKey.ValueString(),
		DnsSlApiver:                 api.SelectedMap(d.DNSSlApiver.ValueString()),
		DnsSlTokenLifetime:          d.DNSSlTokenLifetime.ValueString(),
		DnsSlAccountId:              d.DNSSlAccountID.ValueString(),
		DnsSlProjectName:            d.DNSSlProjectName.ValueString(),
		DnsSlLoginName:              d.DNSSlLoginName.ValueString(),
		DnsSlPassword:               d.DNSSlPassword.ValueString(),
		DnsSelfhostApikey:           d.DNSSelfhostApikey.ValueString(),
		DnsSelfhostMap:              d.DNSSelfhostMap.ValueString(),
		DnsServercowUsername:        d.DNSServercowUsername.ValueString(),
		DnsServercowPassword:        d.DNSServercowPassword.ValueString(),
		DnsSimplyApiKey:             d.DNSSimplyApiKey.ValueString(),
		DnsSimplyAccountName:        d.DNSSimplyAccountName.ValueString(),
		DnsSpaceshipApiKey:          d.DNSSpaceshipApiKey.ValueString(),
		DnsSpaceshipApiSecret:       d.DNSSpaceshipApiSecret.ValueString(),
		DnsSpaceshipRootDomain:      d.DNSSpaceshipRootDomain.ValueString(),
		DnsTechnitiumToken:          d.DNSTechnitiumToken.ValueString(),
		DnsTechnitiumHostname:       d.DNSTechnitiumHostname.ValueString(),
		DnsTransipUsername:          d.DNSTransipUsername.ValueString(),
		DnsTransipKey:               d.DNSTransipKey.ValueString(),
		DnsTransipTokenGlobalKey:    tools.BoolToString(d.DNSTransipTokenGlobalKey.ValueBool()),
		DnsTimewebToken:             d.DNSTimewebToken.ValueString(),
		DnsUdrUser:                  d.DNSUdrUser.ValueString(),
		DnsUdrPassword:              d.DNSUdrPassword.ValueString(),
		DnsUnoKey:                   d.DNSUnoKey.ValueString(),
		DnsUnoUser:                  d.DNSUnoUser.ValueString(),
		DnsVscaleKey:                d.DNSVscaleKey.ValueString(),
		DnsVultrKey:                 d.DNSVultrKey.ValueString(),
		DnsYandexToken:              d.DNSYandexToken.ValueString(),
		DnsZiloreKey:                d.DNSZiloreKey.ValueString(),
		DnsZmKey:                    d.DNSZmKey.ValueString(),
		DnsGdnsdkUser:               d.DNSGdnsdkUser.ValueString(),
		DnsGdnsdkPassword:           d.DNSGdnsdkPassword.ValueString(),
		DnsAcmednsUser:              d.DNSAcmednsUser.ValueString(),
		DnsAcmednsPassword:          d.DNSAcmednsPassword.ValueString(),
		DnsAcmednsSubdomain:         d.DNSAcmednsSubdomain.ValueString(),
		DnsAcmednsUpdateurl:         d.DNSAcmednsUpdateurl.ValueString(),
		DnsAcmednsBaseurl:           d.DNSAcmednsBaseurl.ValueString(),
		DnsAcmeproxyEndpoint:        d.DNSAcmeproxyEndpoint.ValueString(),
		DnsAcmeproxyUsername:        d.DNSAcmeproxyUsername.ValueString(),
		DnsAcmeproxyPassword:        d.DNSAcmeproxyPassword.ValueString(),
		DnsVariomediaKey:            d.DNSVariomediaKey.ValueString(),
		DnsSchlundtechUser:          d.DNSSchlundtechUser.ValueString(),
		DnsSchlundtechPassword:      d.DNSSchlundtechPassword.ValueString(),
		DnsEasydnsApitoken:          d.DNSEasydnsApitoken.ValueString(),
		DnsEasydnsApikey:            d.DNSEasydnsApikey.ValueString(),
		DnsEdgecenterApikey:         d.DNSEdgecenterApikey.ValueString(),
		DnsEuservUser:               d.DNSEuservUser.ValueString(),
		DnsEuservPassword:           d.DNSEuservPassword.ValueString(),
		DnsLeasewebKey:              d.DNSLeasewebKey.ValueString(),
		DnsCnUser:                   d.DNSCnUser.ValueString(),
		DnsCnPassword:               d.DNSCnPassword.ValueString(),
		DnsArvanToken:               d.DNSArvanToken.ValueString(),
		DnsArtfilesUsername:         d.DNSArtfilesUsername.ValueString(),
		DnsArtfilesPassword:         d.DNSArtfilesPassword.ValueString(),
		DnsHetznerToken:             d.DNSHetznerToken.ValueString(),
		DnsHetznercloudToken:        d.DNSHetznercloudToken.ValueString(),
		DnsHexonetLogin:             d.DNSHexonetLogin.ValueString(),
		DnsHexonetPassword:          d.DNSHexonetPassword.ValueString(),
		Dns1984hostingUser:          d.DNS1984hostingUser.ValueString(),
		Dns1984hostingPassword:      d.DNS1984hostingPassword.ValueString(),
		Dns1984hostingTotpSecret:    d.DNS1984hostingTotpSecret.ValueString(),
		DnsKasLogin:                 d.DNSKasLogin.ValueString(),
		DnsKasAuthdata:              d.DNSKasAuthdata.ValueString(),
		DnsKasAuthtype:              api.SelectedMap(d.DNSKasAuthtype.ValueString()),
		DnsDnsmintApiKey:            d.DNSDnsmintApiKey.ValueString(),
		DnsDesecToken:               d.DNSDesecToken.ValueString(),
		DnsDesecName:                d.DNSDesecName.ValueString(),
		DnsInfomaniakToken:          d.DNSInfomaniakToken.ValueString(),
		DnsZoneUsername:             d.DNSZoneUsername.ValueString(),
		DnsZoneKey:                  d.DNSZoneKey.ValueString(),
		DnsDynv6Token:               d.DNSDynv6Token.ValueString(),
		DnsCpanelUser:               d.DNSCpanelUser.ValueString(),
		DnsCpanelToken:              d.DNSCpanelToken.ValueString(),
		DnsCpanelHostname:           d.DNSCpanelHostname.ValueString(),
		DnsRegruUsername:            d.DNSRegruUsername.ValueString(),
		DnsRegruPassword:            d.DNSRegruPassword.ValueString(),
		DnsNicUsername:              d.DNSNicUsername.ValueString(),
		DnsNicPassword:              d.DNSNicPassword.ValueString(),
		DnsNicClient:                d.DNSNicClient.ValueString(),
		DnsNicSecret:                d.DNSNicSecret.ValueString(),
		DnsWebsupportApiKey:         d.DNSWebsupportApiKey.ValueString(),
		DnsWebsupportApiSecret:      d.DNSWebsupportApiSecret.ValueString(),
		DnsWedosUsername:            d.DNSWedosUsername.ValueString(),
		DnsWedosWapipass:            d.DNSWedosWapipass.ValueString(),
		DnsWorld4youUsername:        d.DNSWorld4youUsername.ValueString(),
		DnsWorld4youPassword:        d.DNSWorld4youPassword.ValueString(),
		DnsAuroraKey:                d.DNSAuroraKey.ValueString(),
		DnsAuroraSecret:             d.DNSAuroraSecret.ValueString(),
		DnsConohaUser:               d.DNSConohaUser.ValueString(),
		DnsConohaPassword:           d.DNSConohaPassword.ValueString(),
		DnsConohaTenantid:           d.DNSConohaTenantid.ValueString(),
		DnsConohaIdapi:              d.DNSConohaIdapi.ValueString(),
		DnsConstellixKey:            d.DNSConstellixKey.ValueString(),
		DnsConstellixSecret:         d.DNSConstellixSecret.ValueString(),
		DnsExoscaleKey:              d.DNSExoscaleKey.ValueString(),
		DnsExoscaleSecret:           d.DNSExoscaleSecret.ValueString(),
		DnsInternetbsKey:            d.DNSInternetbsKey.ValueString(),
		DnsInternetbsPassword:       d.DNSInternetbsPassword.ValueString(),
		DnsPointhqKey:               d.DNSPointhqKey.ValueString(),
		DnsPointhqEmail:             d.DNSPointhqEmail.ValueString(),
		DnsRackspaceUser:            d.DNSRackspaceUser.ValueString(),
		DnsRackspaceKey:             d.DNSRackspaceKey.ValueString(),
		DnsRage4Token:               d.DNSRage4Token.ValueString(),
		DnsRage4User:                d.DNSRage4User.ValueString(),
		DnsRcode0ApiToken:           d.DNSRcode0ApiToken.ValueString(),
		DnsRcode0Url:                d.DNSRcode0URL.ValueString(),
		DnsRcode0Ttl:                tools.Int64ToString(d.DNSRcode0TTL.ValueInt64()),
		DnsMijnhostApiKey:           d.DNSMijnhostApiKey.ValueString(),
		DnsScalewayToken:            d.DNSScalewayToken.ValueString(),
		DnsZoneeditId:               d.DNSZoneeditID.ValueString(),
		DnsZoneeditToken:            d.DNSZoneeditToken.ValueString(),
	}, nil
}

func convertValidationStructToSchema(d *acmeclient.Validation) (*validationResourceModel, error) {
	return &validationResourceModel{
		Enabled:                     types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                        types.StringValue(d.Name),
		Description:                 types.StringValue(d.Description),
		Method:                      types.StringValue(d.Method.String()),
		HttpService:                 types.StringValue(d.HttpService.String()),
		HttpOpnAutodiscovery:        types.BoolValue(tools.StringToBool(d.HttpOpnAutodiscovery)),
		HttpOpnInterface:            types.StringValue(d.HttpOpnInterface.String()),
		HttpOpnIpaddresses:          tools.StringSliceToSet(d.HttpOpnIpaddresses),
		HttpHaproxyInject:           types.BoolValue(tools.StringToBool(d.HttpHaproxyInject)),
		HttpHaproxyFrontends:        tools.StringSliceToSet(d.HttpHaproxyFrontends),
		TlsalpnService:              types.StringValue(d.TlsalpnService.String()),
		TlsalpnAcmeAutodiscovery:    types.BoolValue(tools.StringToBool(d.TlsalpnAcmeAutodiscovery)),
		TlsalpnAcmeInterface:        types.StringValue(d.TlsalpnAcmeInterface.String()),
		TlsalpnAcmeIpaddresses:      tools.StringSliceToSet(d.TlsalpnAcmeIpaddresses),
		DNSService:                  types.StringValue(d.DnsService.String()),
		DNSSleep:                    types.Int64Value(tools.StringToInt64(d.DnsSleep)),
		DNSActive24ApiKey:           types.StringValue(d.DnsActive24ApiKey),
		DNSActive24ApiSecret:        types.StringValue(d.DnsActive24ApiSecret),
		DNSAdKey:                    types.StringValue(d.DnsAdKey),
		DNSAliKey:                   types.StringValue(d.DnsAliKey),
		DNSAliSecret:                types.StringValue(d.DnsAliSecret),
		DNSAutodnsUser:              types.StringValue(d.DnsAutodnsUser),
		DNSAutodnsPassword:          types.StringValue(d.DnsAutodnsPassword),
		DNSAutodnsContext:           types.StringValue(d.DnsAutodnsContext),
		DNSAwsID:                    types.StringValue(d.DnsAwsId),
		DNSAwsSecret:                types.StringValue(d.DnsAwsSecret),
		DNSAzurednsSubscriptionid:   types.StringValue(d.DnsAzurednsSubscriptionid),
		DNSAzurednsTenantid:         types.StringValue(d.DnsAzurednsTenantid),
		DNSAzurednsAppid:            types.StringValue(d.DnsAzurednsAppid),
		DNSAzurednsClientsecret:     types.StringValue(d.DnsAzurednsClientsecret),
		DNSAzurednsManagedidentity:  types.BoolValue(tools.StringToBool(d.DnsAzurednsManagedidentity)),
		DNSBhostedUser:              types.StringValue(d.DnsBhostedUser),
		DNSBhostedPassword:          types.StringValue(d.DnsBhostedPassword),
		DNSBhostedTTL:               types.StringValue(d.DnsBhostedTtl),
		DNSBhostedSld:               types.StringValue(d.DnsBhostedSld),
		DNSBhostedTld:               types.StringValue(d.DnsBhostedTld),
		DNSBunnyApiKey:              types.StringValue(d.DnsBunnyApiKey),
		DNSCfEmail:                  types.StringValue(d.DnsCfEmail),
		DNSCfKey:                    types.StringValue(d.DnsCfKey),
		DNSCfToken:                  types.StringValue(d.DnsCfToken),
		DNSCfAccountID:              types.StringValue(d.DnsCfAccountId),
		DNSCfZoneID:                 types.StringValue(d.DnsCfZoneId),
		DNSCloudnsAuthID:            types.StringValue(d.DnsCloudnsAuthId),
		DNSCloudnsSubAuthID:         types.StringValue(d.DnsCloudnsSubAuthId),
		DNSCloudnsAuthPassword:      types.StringValue(d.DnsCloudnsAuthPassword),
		DNSCxKey:                    types.StringValue(d.DnsCxKey),
		DNSCxSecret:                 types.StringValue(d.DnsCxSecret),
		DNSCyonUser:                 types.StringValue(d.DnsCyonUser),
		DNSCyonPassword:             types.StringValue(d.DnsCyonPassword),
		DNSDaKey:                    types.StringValue(d.DnsDaKey),
		DNSDaInsecure:               types.BoolValue(tools.StringToBool(d.DnsDaInsecure)),
		DNSDdnssToken:               types.StringValue(d.DnsDdnssToken),
		DNSDgonKey:                  types.StringValue(d.DnsDgonKey),
		DNSDnsexitAuthUser:          types.StringValue(d.DnsDnsexitAuthUser),
		DNSDnsexitAuthPass:          types.StringValue(d.DnsDnsexitAuthPass),
		DNSDnsexitApi:               types.StringValue(d.DnsDnsexitApi),
		DNSDnshomePassword:          types.StringValue(d.DnsDnshomePassword),
		DNSDnshomeSubdomain:         types.StringValue(d.DnsDnshomeSubdomain),
		DNSDnsimpleToken:            types.StringValue(d.DnsDnsimpleToken),
		DNSDnsservicesUser:          types.StringValue(d.DnsDnsservicesUser),
		DNSDnsservicesPassword:      types.StringValue(d.DnsDnsservicesPassword),
		DNSDoapiToken:               types.StringValue(d.DnsDoapiToken),
		DNSDoPid:                    types.StringValue(d.DnsDoPid),
		DNSDoPassword:               types.StringValue(d.DnsDoPassword),
		DNSDomeneshopToken:          types.StringValue(d.DnsDomeneshopToken),
		DNSDomeneshopSecret:         types.StringValue(d.DnsDomeneshopSecret),
		DNSDpID:                     types.StringValue(d.DnsDpId),
		DNSDpKey:                    types.StringValue(d.DnsDpKey),
		DNSDhKey:                    types.StringValue(d.DnsDhKey),
		DNSDuckdnsToken:             types.StringValue(d.DnsDuckdnsToken),
		DNSDynCustomer:              types.StringValue(d.DnsDynCustomer),
		DNSDynUser:                  types.StringValue(d.DnsDynUser),
		DNSDynPassword:              types.StringValue(d.DnsDynPassword),
		DNSDynuClientid:             types.StringValue(d.DnsDynuClientid),
		DNSDynuSecret:               types.StringValue(d.DnsDynuSecret),
		DNSFreednsUser:              types.StringValue(d.DnsFreednsUser),
		DNSFreednsPassword:          types.StringValue(d.DnsFreednsPassword),
		DNSFreemyipToken:            types.StringValue(d.DnsFreemyipToken),
		DNSFornexApiKey:             types.StringValue(d.DnsFornexApiKey),
		DNSGandiLivednsKey:          types.StringValue(d.DnsGandiLivednsKey),
		DNSGandiLivednsToken:        types.StringValue(d.DnsGandiLivednsToken),
		DNSGcloudKey:                types.StringValue(d.DnsGcloudKey),
		DNSGcoreToken:               types.StringValue(d.DnsGcoreToken),
		DNSGoogledomainsAccessToken: types.StringValue(d.DnsGoogledomainsAccessToken),
		DNSGoogledomainsZone:        types.StringValue(d.DnsGoogledomainsZone),
		DNSGdKey:                    types.StringValue(d.DnsGdKey),
		DNSGdSecret:                 types.StringValue(d.DnsGdSecret),
		DNSHostingdeServer:          types.StringValue(d.DnsHostingdeServer),
		DNSHostingdeApiKey:          types.StringValue(d.DnsHostingdeApiKey),
		DNSHostingerToken:           types.StringValue(d.DnsHostingerToken),
		DNSHeUser:                   types.StringValue(d.DnsHeUser),
		DNSHePassword:               types.StringValue(d.DnsHePassword),
		DNSHeDdnsKey:                types.StringValue(d.DnsHeDdnsKey),
		DNSInfobloxCredentials:      types.StringValue(d.DnsInfobloxCredentials),
		DNSInfobloxServer:           types.StringValue(d.DnsInfobloxServer),
		DNSInwxUser:                 types.StringValue(d.DnsInwxUser),
		DNSInwsPassword:             types.StringValue(d.DnsInwsPassword),
		DNSInwxPassword:             types.StringValue(d.DnsInwxPassword),
		DNSInwxSharedSecret:         types.StringValue(d.DnsInwxSharedSecret),
		DNSIonosPrefix:              types.StringValue(d.DnsIonosPrefix),
		DNSIonosSecret:              types.StringValue(d.DnsIonosSecret),
		DNSIpv64Token:               types.StringValue(d.DnsIpv64Token),
		DNSIspconfigUser:            types.StringValue(d.DnsIspconfigUser),
		DNSIspconfigPassword:        types.StringValue(d.DnsIspconfigPassword),
		DNSIspconfigApi:             types.StringValue(d.DnsIspconfigApi),
		DNSIspconfigInsecure:        types.BoolValue(tools.StringToBool(d.DnsIspconfigInsecure)),
		DNSJdID:                     types.StringValue(d.DnsJdId),
		DNSJdRegion:                 types.StringValue(d.DnsJdRegion),
		DNSJdSecret:                 types.StringValue(d.DnsJdSecret),
		DNSJokerUsername:            types.StringValue(d.DnsJokerUsername),
		DNSJokerPassword:            types.StringValue(d.DnsJokerPassword),
		DNSKinghostUsername:         types.StringValue(d.DnsKinghostUsername),
		DNSKinghostPassword:         types.StringValue(d.DnsKinghostPassword),
		DNSKnotServer:               types.StringValue(d.DnsKnotServer),
		DNSKnotKey:                  types.StringValue(d.DnsKnotKey),
		DNSLexiconProvider:          types.StringValue(d.DnsLexiconProvider.String()),
		DNSLexiconUser:              types.StringValue(d.DnsLexiconUser),
		DNSLexiconToken:             types.StringValue(d.DnsLexiconToken),
		DNSLimacityApikey:           types.StringValue(d.DnsLimacityApikey),
		DNSLinodeKey:                types.StringValue(d.DnsLinodeKey),
		DNSLinodeV4Key:              types.StringValue(d.DnsLinodeV4Key),
		DNSLoopiaApi:                types.StringValue(d.DnsLoopiaApi),
		DNSLoopiaUser:               types.StringValue(d.DnsLoopiaUser),
		DNSLoopiaPassword:           types.StringValue(d.DnsLoopiaPassword),
		DNSLuaEmail:                 types.StringValue(d.DnsLuaEmail),
		DNSLuaKey:                   types.StringValue(d.DnsLuaKey),
		DNSMiabUser:                 types.StringValue(d.DnsMiabUser),
		DNSMiabPassword:             types.StringValue(d.DnsMiabPassword),
		DNSMiabServer:               types.StringValue(d.DnsMiabServer),
		DNSMeKey:                    types.StringValue(d.DnsMeKey),
		DNSMeSecret:                 types.StringValue(d.DnsMeSecret),
		DNSMydnsjpMasterid:          types.StringValue(d.DnsMydnsjpMasterid),
		DNSMydnsjpPassword:          types.StringValue(d.DnsMydnsjpPassword),
		DNSMylocToken:               types.StringValue(d.DnsMylocToken),
		DNSMythicBeastsKey:          types.StringValue(d.DnsMythicBeastsKey),
		DNSMythicBeastsSecret:       types.StringValue(d.DnsMythicBeastsSecret),
		DNSNamecheapUser:            types.StringValue(d.DnsNamecheapUser),
		DNSNamecheapApi:             types.StringValue(d.DnsNamecheapApi),
		DNSNamecheapSourceip:        types.StringValue(d.DnsNamecheapSourceip),
		DNSNamecomUser:              types.StringValue(d.DnsNamecomUser),
		DNSNamecomToken:             types.StringValue(d.DnsNamecomToken),
		DNSNamesiloKey:              types.StringValue(d.DnsNamesiloKey),
		DNSNederhostKey:             types.StringValue(d.DnsNederhostKey),
		DNSNetcupCid:                types.StringValue(d.DnsNetcupCid),
		DNSNetcupKey:                types.StringValue(d.DnsNetcupKey),
		DNSNetcupPw:                 types.StringValue(d.DnsNetcupPw),
		DNSNjallaToken:              types.StringValue(d.DnsNjallaToken),
		DNSNsoneKey:                 types.StringValue(d.DnsNsoneKey),
		DNSNsupdateServer:           types.StringValue(d.DnsNsupdateServer),
		DNSNsupdateZone:             types.StringValue(d.DnsNsupdateZone),
		DNSNsupdateKey:              types.StringValue(d.DnsNsupdateKey),
		DNSOciCliUser:               types.StringValue(d.DnsOciCliUser),
		DNSOciCliTenancy:            types.StringValue(d.DnsOciCliTenancy),
		DNSOciCliRegion:             types.StringValue(d.DnsOciCliRegion),
		DNSOciCliKey:                types.StringValue(d.DnsOciCliKey),
		DNSOnlineKey:                types.StringValue(d.DnsOnlineKey),
		DNSOpenproviderRestUsername: types.StringValue(d.DnsOpenproviderRestUsername),
		DNSOpenproviderRestPassword: types.StringValue(d.DnsOpenproviderRestPassword),
		DNSOpnsenseHost:             types.StringValue(d.DnsOpnsenseHost),
		DNSOpnsensePort:             types.StringValue(d.DnsOpnsensePort),
		DNSOpnsenseKey:              types.StringValue(d.DnsOpnsenseKey),
		DNSOpnsenseToken:            types.StringValue(d.DnsOpnsenseToken),
		DNSOpnsenseInsecure:         types.BoolValue(tools.StringToBool(d.DnsOpnsenseInsecure)),
		DNSOvhAppKey:                types.StringValue(d.DnsOvhAppKey),
		DNSOvhAppSecret:             types.StringValue(d.DnsOvhAppSecret),
		DNSOvhConsumerKey:           types.StringValue(d.DnsOvhConsumerKey),
		DNSOvhEndpoint:              types.StringValue(d.DnsOvhEndpoint),
		DNSPleskxmlUser:             types.StringValue(d.DnsPleskxmlUser),
		DNSPleskxmlPass:             types.StringValue(d.DnsPleskxmlPass),
		DNSPleskxmlUri:              types.StringValue(d.DnsPleskxmlUri),
		DNSPdnsURL:                  types.StringValue(d.DnsPdnsUrl),
		DNSPdnsServerid:             types.StringValue(d.DnsPdnsServerid),
		DNSPdnsToken:                types.StringValue(d.DnsPdnsToken),
		DNSPorkbunKey:               types.StringValue(d.DnsPorkbunKey),
		DNSPorkbunSecret:            types.StringValue(d.DnsPorkbunSecret),
		DNSSlKey:                    types.StringValue(d.DnsSlKey),
		DNSSlApiver:                 types.StringValue(d.DnsSlApiver.String()),
		DNSSlTokenLifetime:          types.StringValue(d.DnsSlTokenLifetime),
		DNSSlAccountID:              types.StringValue(d.DnsSlAccountId),
		DNSSlProjectName:            types.StringValue(d.DnsSlProjectName),
		DNSSlLoginName:              types.StringValue(d.DnsSlLoginName),
		DNSSlPassword:               types.StringValue(d.DnsSlPassword),
		DNSSelfhostApikey:           types.StringValue(d.DnsSelfhostApikey),
		DNSSelfhostMap:              types.StringValue(d.DnsSelfhostMap),
		DNSServercowUsername:        types.StringValue(d.DnsServercowUsername),
		DNSServercowPassword:        types.StringValue(d.DnsServercowPassword),
		DNSSimplyApiKey:             types.StringValue(d.DnsSimplyApiKey),
		DNSSimplyAccountName:        types.StringValue(d.DnsSimplyAccountName),
		DNSSpaceshipApiKey:          types.StringValue(d.DnsSpaceshipApiKey),
		DNSSpaceshipApiSecret:       types.StringValue(d.DnsSpaceshipApiSecret),
		DNSSpaceshipRootDomain:      types.StringValue(d.DnsSpaceshipRootDomain),
		DNSTechnitiumToken:          types.StringValue(d.DnsTechnitiumToken),
		DNSTechnitiumHostname:       types.StringValue(d.DnsTechnitiumHostname),
		DNSTransipUsername:          types.StringValue(d.DnsTransipUsername),
		DNSTransipKey:               types.StringValue(d.DnsTransipKey),
		DNSTransipTokenGlobalKey:    types.BoolValue(tools.StringToBool(d.DnsTransipTokenGlobalKey)),
		DNSTimewebToken:             types.StringValue(d.DnsTimewebToken),
		DNSUdrUser:                  types.StringValue(d.DnsUdrUser),
		DNSUdrPassword:              types.StringValue(d.DnsUdrPassword),
		DNSUnoKey:                   types.StringValue(d.DnsUnoKey),
		DNSUnoUser:                  types.StringValue(d.DnsUnoUser),
		DNSVscaleKey:                types.StringValue(d.DnsVscaleKey),
		DNSVultrKey:                 types.StringValue(d.DnsVultrKey),
		DNSYandexToken:              types.StringValue(d.DnsYandexToken),
		DNSZiloreKey:                types.StringValue(d.DnsZiloreKey),
		DNSZmKey:                    types.StringValue(d.DnsZmKey),
		DNSGdnsdkUser:               types.StringValue(d.DnsGdnsdkUser),
		DNSGdnsdkPassword:           types.StringValue(d.DnsGdnsdkPassword),
		DNSAcmednsUser:              types.StringValue(d.DnsAcmednsUser),
		DNSAcmednsPassword:          types.StringValue(d.DnsAcmednsPassword),
		DNSAcmednsSubdomain:         types.StringValue(d.DnsAcmednsSubdomain),
		DNSAcmednsUpdateurl:         types.StringValue(d.DnsAcmednsUpdateurl),
		DNSAcmednsBaseurl:           types.StringValue(d.DnsAcmednsBaseurl),
		DNSAcmeproxyEndpoint:        types.StringValue(d.DnsAcmeproxyEndpoint),
		DNSAcmeproxyUsername:        types.StringValue(d.DnsAcmeproxyUsername),
		DNSAcmeproxyPassword:        types.StringValue(d.DnsAcmeproxyPassword),
		DNSVariomediaKey:            types.StringValue(d.DnsVariomediaKey),
		DNSSchlundtechUser:          types.StringValue(d.DnsSchlundtechUser),
		DNSSchlundtechPassword:      types.StringValue(d.DnsSchlundtechPassword),
		DNSEasydnsApitoken:          types.StringValue(d.DnsEasydnsApitoken),
		DNSEasydnsApikey:            types.StringValue(d.DnsEasydnsApikey),
		DNSEdgecenterApikey:         types.StringValue(d.DnsEdgecenterApikey),
		DNSEuservUser:               types.StringValue(d.DnsEuservUser),
		DNSEuservPassword:           types.StringValue(d.DnsEuservPassword),
		DNSLeasewebKey:              types.StringValue(d.DnsLeasewebKey),
		DNSCnUser:                   types.StringValue(d.DnsCnUser),
		DNSCnPassword:               types.StringValue(d.DnsCnPassword),
		DNSArvanToken:               types.StringValue(d.DnsArvanToken),
		DNSArtfilesUsername:         types.StringValue(d.DnsArtfilesUsername),
		DNSArtfilesPassword:         types.StringValue(d.DnsArtfilesPassword),
		DNSHetznerToken:             types.StringValue(d.DnsHetznerToken),
		DNSHetznercloudToken:        types.StringValue(d.DnsHetznercloudToken),
		DNSHexonetLogin:             types.StringValue(d.DnsHexonetLogin),
		DNSHexonetPassword:          types.StringValue(d.DnsHexonetPassword),
		DNS1984hostingUser:          types.StringValue(d.Dns1984hostingUser),
		DNS1984hostingPassword:      types.StringValue(d.Dns1984hostingPassword),
		DNS1984hostingTotpSecret:    types.StringValue(d.Dns1984hostingTotpSecret),
		DNSKasLogin:                 types.StringValue(d.DnsKasLogin),
		DNSKasAuthdata:              types.StringValue(d.DnsKasAuthdata),
		DNSKasAuthtype:              types.StringValue(d.DnsKasAuthtype.String()),
		DNSDnsmintApiKey:            types.StringValue(d.DnsDnsmintApiKey),
		DNSDesecToken:               types.StringValue(d.DnsDesecToken),
		DNSDesecName:                types.StringValue(d.DnsDesecName),
		DNSInfomaniakToken:          types.StringValue(d.DnsInfomaniakToken),
		DNSZoneUsername:             types.StringValue(d.DnsZoneUsername),
		DNSZoneKey:                  types.StringValue(d.DnsZoneKey),
		DNSDynv6Token:               types.StringValue(d.DnsDynv6Token),
		DNSCpanelUser:               types.StringValue(d.DnsCpanelUser),
		DNSCpanelToken:              types.StringValue(d.DnsCpanelToken),
		DNSCpanelHostname:           types.StringValue(d.DnsCpanelHostname),
		DNSRegruUsername:            types.StringValue(d.DnsRegruUsername),
		DNSRegruPassword:            types.StringValue(d.DnsRegruPassword),
		DNSNicUsername:              types.StringValue(d.DnsNicUsername),
		DNSNicPassword:              types.StringValue(d.DnsNicPassword),
		DNSNicClient:                types.StringValue(d.DnsNicClient),
		DNSNicSecret:                types.StringValue(d.DnsNicSecret),
		DNSWebsupportApiKey:         types.StringValue(d.DnsWebsupportApiKey),
		DNSWebsupportApiSecret:      types.StringValue(d.DnsWebsupportApiSecret),
		DNSWedosUsername:            types.StringValue(d.DnsWedosUsername),
		DNSWedosWapipass:            types.StringValue(d.DnsWedosWapipass),
		DNSWorld4youUsername:        types.StringValue(d.DnsWorld4youUsername),
		DNSWorld4youPassword:        types.StringValue(d.DnsWorld4youPassword),
		DNSAuroraKey:                types.StringValue(d.DnsAuroraKey),
		DNSAuroraSecret:             types.StringValue(d.DnsAuroraSecret),
		DNSConohaUser:               types.StringValue(d.DnsConohaUser),
		DNSConohaPassword:           types.StringValue(d.DnsConohaPassword),
		DNSConohaTenantid:           types.StringValue(d.DnsConohaTenantid),
		DNSConohaIdapi:              types.StringValue(d.DnsConohaIdapi),
		DNSConstellixKey:            types.StringValue(d.DnsConstellixKey),
		DNSConstellixSecret:         types.StringValue(d.DnsConstellixSecret),
		DNSExoscaleKey:              types.StringValue(d.DnsExoscaleKey),
		DNSExoscaleSecret:           types.StringValue(d.DnsExoscaleSecret),
		DNSInternetbsKey:            types.StringValue(d.DnsInternetbsKey),
		DNSInternetbsPassword:       types.StringValue(d.DnsInternetbsPassword),
		DNSPointhqKey:               types.StringValue(d.DnsPointhqKey),
		DNSPointhqEmail:             types.StringValue(d.DnsPointhqEmail),
		DNSRackspaceUser:            types.StringValue(d.DnsRackspaceUser),
		DNSRackspaceKey:             types.StringValue(d.DnsRackspaceKey),
		DNSRage4Token:               types.StringValue(d.DnsRage4Token),
		DNSRage4User:                types.StringValue(d.DnsRage4User),
		DNSRcode0ApiToken:           types.StringValue(d.DnsRcode0ApiToken),
		DNSRcode0URL:                types.StringValue(d.DnsRcode0Url),
		DNSRcode0TTL:                types.Int64Value(tools.StringToInt64(d.DnsRcode0Ttl)),
		DNSMijnhostApiKey:           types.StringValue(d.DnsMijnhostApiKey),
		DNSScalewayToken:            types.StringValue(d.DnsScalewayToken),
		DNSZoneeditID:               types.StringValue(d.DnsZoneeditId),
		DNSZoneeditToken:            types.StringValue(d.DnsZoneeditToken),
	}, nil
}

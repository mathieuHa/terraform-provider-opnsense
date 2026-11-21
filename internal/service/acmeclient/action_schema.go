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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type actionResourceModel struct {
	Enabled                       types.Bool   `tfsdk:"enabled"`
	Name                          types.String `tfsdk:"name"`
	Description                   types.String `tfsdk:"description"`
	Type                          types.String `tfsdk:"type"`
	SftpHost                      types.String `tfsdk:"sftp_host"`
	SftpHostKey                   types.String `tfsdk:"sftp_host_key"`
	SftpPort                      types.Int64  `tfsdk:"sftp_port"`
	SftpUser                      types.String `tfsdk:"sftp_user"`
	SftpIdentityType              types.String `tfsdk:"sftp_identity_type"`
	SftpRemotePath                types.String `tfsdk:"sftp_remote_path"`
	SftpChgrp                     types.String `tfsdk:"sftp_chgrp"`
	SftpChmod                     types.String `tfsdk:"sftp_chmod"`
	SftpChmodKey                  types.String `tfsdk:"sftp_chmod_key"`
	SftpModtime                   types.Bool   `tfsdk:"sftp_modtime"`
	SftpFilenameCert              types.String `tfsdk:"sftp_filename_cert"`
	SftpFilenameKey               types.String `tfsdk:"sftp_filename_key"`
	SftpFilenameCa                types.String `tfsdk:"sftp_filename_ca"`
	SftpFilenameFullchain         types.String `tfsdk:"sftp_filename_fullchain"`
	RemoteSshHost                 types.String `tfsdk:"remote_ssh_host"`
	RemoteSshHostKey              types.String `tfsdk:"remote_ssh_host_key"`
	RemoteSshPort                 types.Int64  `tfsdk:"remote_ssh_port"`
	RemoteSshUser                 types.String `tfsdk:"remote_ssh_user"`
	RemoteSshIdentityType         types.String `tfsdk:"remote_ssh_identity_type"`
	RemoteSshCommand              types.String `tfsdk:"remote_ssh_command"`
	Configd                       types.String `tfsdk:"configd"`
	ConfigdGenericCommand         types.String `tfsdk:"configd_generic_command"`
	AcmeSynologyDsmHostname       types.String `tfsdk:"acme_synology_dsm_hostname"`
	AcmeSynologyDsmPort           types.Int64  `tfsdk:"acme_synology_dsm_port"`
	AcmeSynologyDsmScheme         types.String `tfsdk:"acme_synology_dsm_scheme"`
	AcmeSynologyDsmUsername       types.String `tfsdk:"acme_synology_dsm_username"`
	AcmeSynologyDsmPassword       types.String `tfsdk:"acme_synology_dsm_password"`
	AcmeSynologyDsmCreate         types.Bool   `tfsdk:"acme_synology_dsm_create"`
	AcmeSynologyDsmDeviceid       types.String `tfsdk:"acme_synology_dsm_deviceid"`
	AcmeSynologyDsmDevicename     types.String `tfsdk:"acme_synology_dsm_devicename"`
	AcmeSynologyDsmOtpcode        types.String `tfsdk:"acme_synology_dsm_otpcode"`
	AcmeFritzboxURL               types.String `tfsdk:"acme_fritzbox_url"`
	AcmeFritzboxUsername          types.String `tfsdk:"acme_fritzbox_username"`
	AcmeFritzboxPassword          types.String `tfsdk:"acme_fritzbox_password"`
	AcmePanosUsername             types.String `tfsdk:"acme_panos_username"`
	AcmePanosPassword             types.String `tfsdk:"acme_panos_password"`
	AcmePanosHost                 types.String `tfsdk:"acme_panos_host"`
	AcmeProxmoxveUser             types.String `tfsdk:"acme_proxmoxve_user"`
	AcmeProxmoxveServer           types.String `tfsdk:"acme_proxmoxve_server"`
	AcmeProxmoxvePort             types.Int64  `tfsdk:"acme_proxmoxve_port"`
	AcmeProxmoxveNodename         types.String `tfsdk:"acme_proxmoxve_nodename"`
	AcmeProxmoxveRealm            types.String `tfsdk:"acme_proxmoxve_realm"`
	AcmeProxmoxveTokenid          types.String `tfsdk:"acme_proxmoxve_tokenid"`
	AcmeProxmoxveTokenkey         types.String `tfsdk:"acme_proxmoxve_tokenkey"`
	AcmeProxmoxbsUser             types.String `tfsdk:"acme_proxmoxbs_user"`
	AcmeProxmoxbsServer           types.String `tfsdk:"acme_proxmoxbs_server"`
	AcmeProxmoxbsPort             types.Int64  `tfsdk:"acme_proxmoxbs_port"`
	AcmeProxmoxbsNodename         types.String `tfsdk:"acme_proxmoxbs_nodename"`
	AcmeProxmoxbsRealm            types.String `tfsdk:"acme_proxmoxbs_realm"`
	AcmeProxmoxbsTokenid          types.String `tfsdk:"acme_proxmoxbs_tokenid"`
	AcmeProxmoxbsTokenkey         types.String `tfsdk:"acme_proxmoxbs_tokenkey"`
	AcmeRuckusHost                types.String `tfsdk:"acme_ruckus_host"`
	AcmeRuckusUser                types.String `tfsdk:"acme_ruckus_user"`
	AcmeRuckusPass                types.String `tfsdk:"acme_ruckus_pass"`
	AcmeTruenasApikey             types.String `tfsdk:"acme_truenas_apikey"`
	AcmeTruenasHostname           types.String `tfsdk:"acme_truenas_hostname"`
	AcmeTruenasScheme             types.String `tfsdk:"acme_truenas_scheme"`
	AcmeTruenasWsApikey           types.String `tfsdk:"acme_truenas_ws_apikey"`
	AcmeTruenasWsHostname         types.String `tfsdk:"acme_truenas_ws_hostname"`
	AcmeTruenasWsProtocol         types.String `tfsdk:"acme_truenas_ws_protocol"`
	AcmeTruenasWebsocatApikey     types.String `tfsdk:"acme_truenas_websocat_apikey"`
	AcmeTruenasWebsocatHostname   types.String `tfsdk:"acme_truenas_websocat_hostname"`
	AcmeTruenasWebsocatProtocol   types.String `tfsdk:"acme_truenas_websocat_protocol"`
	AcmeTruenasWebsocatPort       types.Int64  `tfsdk:"acme_truenas_websocat_port"`
	AcmeTruenasWebsocatUpdateFtp  types.String `tfsdk:"acme_truenas_websocat_update_ftp"`
	AcmeTruenasWebsocatUpdateApps types.String `tfsdk:"acme_truenas_websocat_update_apps"`
	AcmeTruenaswsApikey           types.String `tfsdk:"acme_truenasws_apikey"`
	AcmeTruenaswsHostname         types.String `tfsdk:"acme_truenasws_hostname"`
	AcmeTruenaswsProtocol         types.String `tfsdk:"acme_truenasws_protocol"`
	AcmeUnifiKeystore             types.String `tfsdk:"acme_unifi_keystore"`
	AcmeVaultURL                  types.String `tfsdk:"acme_vault_url"`
	AcmeVaultPrefix               types.String `tfsdk:"acme_vault_prefix"`
	AcmeVaultToken                types.String `tfsdk:"acme_vault_token"`
	AcmeVaultKvv2                 types.Bool   `tfsdk:"acme_vault_kvv2"`
	AcmeZyxelGs1900Host           types.String `tfsdk:"acme_zyxel_gs1900_host"`
	AcmeZyxelGs1900User           types.String `tfsdk:"acme_zyxel_gs1900_user"`
	AcmeZyxelGs1900Password       types.String `tfsdk:"acme_zyxel_gs1900_password"`
	AcmeZyxelGs1900Insecure       types.Bool   `tfsdk:"acme_zyxel_gs1900_insecure"`
	AcmeZyxelGs1900Reboot         types.Bool   `tfsdk:"acme_zyxel_gs1900_reboot"`
	AcmeJetkvmHost                types.String `tfsdk:"acme_jetkvm_host"`
	AcmeJetkvmUser                types.String `tfsdk:"acme_jetkvm_user"`
	AcmeJetkvmPort                types.Int64  `tfsdk:"acme_jetkvm_port"`
	AcmeJetkvmReboot              types.Bool   `tfsdk:"acme_jetkvm_reboot"`

	Id types.String `tfsdk:"id"`
}

func actionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "ACME automations (**Services ▸ ACME Client ▸ Automations**, plugin `os-acme-client`): commands run after a certificate is issued or renewed, e.g. restart the web GUI or upload the certificate over SFTP.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this automation. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this automation.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this automation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Pre-defined commands for this automation. One of `configd_restart_gui`, `configd_restart_haproxy`, `configd_restart_nginx`, `configd_reload_caddy`, `configd_upload_sftp`, `configd_remote_ssh`, `acme_fritzbox`, `acme_panos`, `acme_proxmoxbs`, `acme_proxmoxve`, `acme_ruckus`, `acme_vault`, `acme_synology_dsm`, `acme_truenas`, `acme_truenas_ws`, `acme_truenas_websocat`, `acme_zyxel_gs1900`, `acme_unifi`, `acme_jetkvm`, `configd_generic`.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("configd_restart_gui", "configd_restart_haproxy", "configd_restart_nginx", "configd_reload_caddy", "configd_upload_sftp", "configd_remote_ssh", "acme_fritzbox", "acme_panos", "acme_proxmoxbs", "acme_proxmoxve", "acme_ruckus", "acme_vault", "acme_synology_dsm", "acme_truenas", "acme_truenas_ws", "acme_truenas_websocat", "acme_zyxel_gs1900", "acme_unifi", "acme_jetkvm", "configd_generic")},
			},
			"sftp_host": schema.StringAttribute{
				MarkdownDescription: "IP address or hostname of the SFTP server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_host_key": schema.StringAttribute{
				MarkdownDescription: "SFTP server host key, formatted as in 'known_hosts'. Leave blank to auto accept host key on first connect (not as secure as specifying it). Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_port": schema.Int64Attribute{
				MarkdownDescription: "SFTP server port. Leave blank to use default \"22\". Defaults to `22`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(22),
			},
			"sftp_user": schema.StringAttribute{
				MarkdownDescription: "The username to login to the SFTP server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_identity_type": schema.StringAttribute{
				MarkdownDescription: "The type of identify to present to the SFTP server for authorization. Select 'none' to use default \"ECDSA\". One of `ecdsa`, `rsa`, `ed25519`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "ecdsa", "rsa", "ed25519")},
			},
			"sftp_remote_path": schema.StringAttribute{
				MarkdownDescription: "Path on the SFTP server to change to after login. The path can be absolute or relative to home and must exist. Leave blank to not change path after login. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_chgrp": schema.StringAttribute{
				MarkdownDescription: "Unix group id to apply to all uploaded files. Leave blank to not change the group. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_chmod": schema.StringAttribute{
				MarkdownDescription: "Unix permission to apply to uploaded public keys. Leave blank to use default \"0440\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_chmod_key": schema.StringAttribute{
				MarkdownDescription: "Unix permission to apply to uploaded private keys. Leave blank to use default \"0400\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_modtime": schema.BoolAttribute{
				MarkdownDescription: "Preserves modification times from the source file. Note that this is not supported by all SFTP servers and may cause the upload to fail, e.g. on VMware. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sftp_filename_cert": schema.StringAttribute{
				MarkdownDescription: "Name template for the public certificate. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/cert.pem\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_filename_key": schema.StringAttribute{
				MarkdownDescription: "Name template for the certificate's private key. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/key.pem\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_filename_ca": schema.StringAttribute{
				MarkdownDescription: "Name template for the public certificate chain file. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/ca.pem\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"sftp_filename_fullchain": schema.StringAttribute{
				MarkdownDescription: "Name template for the public certificate fullchain file (cert + ca). Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/fullchain.pem\". Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"remote_ssh_host": schema.StringAttribute{
				MarkdownDescription: "IP address or hostname of the SSH server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"remote_ssh_host_key": schema.StringAttribute{
				MarkdownDescription: "remote_ssh_host_key. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"remote_ssh_port": schema.Int64Attribute{
				MarkdownDescription: "SSH server port. Leave blank to use default \"22\". Defaults to `22`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(22),
			},
			"remote_ssh_user": schema.StringAttribute{
				MarkdownDescription: "The username to login to the SSH server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"remote_ssh_identity_type": schema.StringAttribute{
				MarkdownDescription: "The type of identify to present to the SSH server for authorization. Select 'none' to use default \"ECDSA\". One of `ecdsa`, `rsa`, `ed25519`. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Validators:          []validator.String{stringvalidator.OneOf("", "ecdsa", "rsa", "ed25519")},
			},
			"remote_ssh_command": schema.StringAttribute{
				MarkdownDescription: "The command to execute on the SSH server. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"configd": schema.StringAttribute{
				MarkdownDescription: "configd. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"configd_generic_command": schema.StringAttribute{
				MarkdownDescription: "Select a pre-defined system command which should be run. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_synology_dsm_hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname of IP address of the Synology DSM, i.e. synology.example.com or 192.168.0.1. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_synology_dsm_port": schema.Int64Attribute{
				MarkdownDescription: "Port that will be used when connecting to Synology DSM. Defaults to `5000`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(5000),
			},
			"acme_synology_dsm_scheme": schema.StringAttribute{
				MarkdownDescription: "Connection scheme that will be used when uploading certificates to Synology DSM. One of `http`, `https`. Defaults to `http`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
				Validators:          []validator.String{stringvalidator.OneOf("", "http", "https")},
			},
			"acme_synology_dsm_username": schema.StringAttribute{
				MarkdownDescription: "Username to login, must be an administrator. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_synology_dsm_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_synology_dsm_create": schema.BoolAttribute{
				MarkdownDescription: "This option ensures that a new certificate is created in Synology DSM if it does not exist yet. If unchecked only existing certificates will be updated. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"acme_synology_dsm_deviceid": schema.StringAttribute{
				MarkdownDescription: "If Synology DSM has OTP enabled, then the device ID has to be provided so that no OTP is required when running the automation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_synology_dsm_devicename": schema.StringAttribute{
				MarkdownDescription: "If Synology DSM has OTP enabled, then the device name has to be provided so that no OTP is required when running the automation. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_synology_dsm_otpcode": schema.StringAttribute{
				MarkdownDescription: "If Synology DSM has OTP enabled, then a OTP may be required. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_fritzbox_url": schema.StringAttribute{
				MarkdownDescription: "URL of the router, i.e. https://fritzbox.example.com. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_fritzbox_username": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_fritzbox_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_panos_username": schema.StringAttribute{
				MarkdownDescription: "Username. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_panos_password": schema.StringAttribute{
				MarkdownDescription: "Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_panos_host": schema.StringAttribute{
				MarkdownDescription: "Host. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_proxmoxve_user": schema.StringAttribute{
				MarkdownDescription: "The user who owns the API key. Defaults to root. Defaults to `root`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("root"),
			},
			"acme_proxmoxve_server": schema.StringAttribute{
				MarkdownDescription: "The hostname of the proxmox ve node. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_proxmoxve_port": schema.Int64Attribute{
				MarkdownDescription: "The port number the management interface is on. Defaults to 8006. Defaults to `8006`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(8006),
			},
			"acme_proxmoxve_nodename": schema.StringAttribute{
				MarkdownDescription: "The name of the node we will be connecting to. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_proxmoxve_realm": schema.StringAttribute{
				MarkdownDescription: "The authentication realm the user authenticates with. Defaults to pam. Defaults to `pam`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("pam"),
			},
			"acme_proxmoxve_tokenid": schema.StringAttribute{
				MarkdownDescription: "The name of the API token created for the user account. Defaults to acme. Defaults to `acme`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("acme"),
				Sensitive:           true,
			},
			"acme_proxmoxve_tokenkey": schema.StringAttribute{
				MarkdownDescription: "The API token. Required. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_proxmoxbs_user": schema.StringAttribute{
				MarkdownDescription: "The user who owns the API key. Defaults to root. Defaults to `root`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("root"),
			},
			"acme_proxmoxbs_server": schema.StringAttribute{
				MarkdownDescription: "The hostname of the proxmox BS node. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_proxmoxbs_port": schema.Int64Attribute{
				MarkdownDescription: "The port number the management interface is on. Defaults to 8007. Defaults to `8007`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(8007),
			},
			"acme_proxmoxbs_nodename": schema.StringAttribute{
				MarkdownDescription: "acme_proxmoxbs_nodename. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"acme_proxmoxbs_realm": schema.StringAttribute{
				MarkdownDescription: "The authentication realm the user authenticates with. Defaults to pam. Defaults to `pam`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("pam"),
			},
			"acme_proxmoxbs_tokenid": schema.StringAttribute{
				MarkdownDescription: "The name of the API token created for the user account. Defaults to acme. Defaults to `acme`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("acme"),
				Sensitive:           true,
			},
			"acme_proxmoxbs_tokenkey": schema.StringAttribute{
				MarkdownDescription: "The API token. Required. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_ruckus_host": schema.StringAttribute{
				MarkdownDescription: "Ruckus Host. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_ruckus_user": schema.StringAttribute{
				MarkdownDescription: "Ruckus User. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_ruckus_pass": schema.StringAttribute{
				MarkdownDescription: "Ruckus Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_truenas_apikey": schema.StringAttribute{
				MarkdownDescription: "API key generated in the TrueNAS web UI. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_truenas_hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP address of TrueNAS Server. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"acme_truenas_scheme": schema.StringAttribute{
				MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Core Server. One of `http`, `https`. Defaults to `http`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
				Validators:          []validator.String{stringvalidator.OneOf("", "http", "https")},
			},
			"acme_truenas_ws_apikey": schema.StringAttribute{
				MarkdownDescription: "API key generated in the TrueNAS web UI. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_truenas_ws_hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP address of TrueNAS Server. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"acme_truenas_ws_protocol": schema.StringAttribute{
				MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Server. One of `ws`, `wss`. Defaults to `ws`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ws"),
				Validators:          []validator.String{stringvalidator.OneOf("", "ws", "wss")},
			},
			"acme_truenas_websocat_apikey": schema.StringAttribute{
				MarkdownDescription: "API key generated in the TrueNAS web UI. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_truenas_websocat_hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP address of TrueNAS Server. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"acme_truenas_websocat_protocol": schema.StringAttribute{
				MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Server. One of `ws`, `wss`. Defaults to `ws`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ws"),
				Validators:          []validator.String{stringvalidator.OneOf("", "ws", "wss")},
			},
			"acme_truenas_websocat_port": schema.Int64Attribute{
				MarkdownDescription: "Port of TrueNAS HTTP Server. Defaults to `80`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(80),
			},
			"acme_truenas_websocat_update_ftp": schema.StringAttribute{
				MarkdownDescription: "Upload certificate to TrueNAS FTP Server. One of `no`, `yes`. Defaults to `no`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("no"),
				Validators:          []validator.String{stringvalidator.OneOf("", "no", "yes")},
			},
			"acme_truenas_websocat_update_apps": schema.StringAttribute{
				MarkdownDescription: "Upload certificate to TrueNAS Server ix-Apps. One of `no`, `yes`. Defaults to `no`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("no"),
				Validators:          []validator.String{stringvalidator.OneOf("", "no", "yes")},
			},
			"acme_truenasws_apikey": schema.StringAttribute{
				MarkdownDescription: "acme_truenasws_apikey. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_truenasws_hostname": schema.StringAttribute{
				MarkdownDescription: "acme_truenasws_hostname. Defaults to `localhost`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("localhost"),
			},
			"acme_truenasws_protocol": schema.StringAttribute{
				MarkdownDescription: "acme_truenasws_protocol. One of `ws`, `wss`. Defaults to `ws`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ws"),
				Validators:          []validator.String{stringvalidator.OneOf("", "ws", "wss")},
			},
			"acme_unifi_keystore": schema.StringAttribute{
				MarkdownDescription: "Path to the Unifi keystore file in the local filesystem, i.e. /usr/local/share/java/unifi/data/keystore. Defaults to `/usr/local/share/java/unifi/data/keystore`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/usr/local/share/java/unifi/data/keystore"),
				Sensitive:           true,
			},
			"acme_vault_url": schema.StringAttribute{
				MarkdownDescription: "URL of the Vault, i.e. http://vault.example.com:8200. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_vault_prefix": schema.StringAttribute{
				MarkdownDescription: "This specifies the prefix path in Vault. If you select KV v2 you need to add .../data/... between the secret-mount-path and the path. Example: v1 prefix path: secret/acme, v2 prefix path: secret/data/acme. Defaults to `acme`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("acme"),
			},
			"acme_vault_token": schema.StringAttribute{
				MarkdownDescription: "This specifies the Vault token to authenticate with. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_vault_kvv2": schema.BoolAttribute{
				MarkdownDescription: "If checked version 2 of the kv store will be used, otherwise version 1. If you use v2 please consider the comment in the field \"Vault Prefix\". Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"acme_zyxel_gs1900_host": schema.StringAttribute{
				MarkdownDescription: "Hostname or IP adress of a Zyxel GS1900 series switch. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_zyxel_gs1900_user": schema.StringAttribute{
				MarkdownDescription: "A user configured as an adminstrator. Defaults to admin. Defaults to `admin`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("admin"),
			},
			"acme_zyxel_gs1900_password": schema.StringAttribute{
				MarkdownDescription: "Switch Admin Password. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				Sensitive:           true,
			},
			"acme_zyxel_gs1900_insecure": schema.BoolAttribute{
				MarkdownDescription: "If checked the deployment will be run in insecure mode and will skip certificate validation when connecting to the switch. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"acme_zyxel_gs1900_reboot": schema.BoolAttribute{
				MarkdownDescription: "If checked the switch will be rebooted once the certificate has been successfully installed. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"acme_jetkvm_host": schema.StringAttribute{
				MarkdownDescription: "IP address or hostname of the JetKVM device. Required first: the device must already have \"HTTPS Mode\" set to \"Custom\" in its own web UI (Settings > Network) -- this automation only writes the certificate/key files (and optionally reboots); it does not switch HTTPS mode for you, and the deploy will fail if it isn't already set (check the ACME Client log for the exact reason). Also requires SSH access to already work non-interactively as the configured user (JetKVM only supports key-based SSH authentication -- enable \"Developer Mode\" on the device and paste this firewall's SSH public key, e.g. from /root/.ssh/id_rsa.pub, into its \"SSH Keys\" field under Settings > Advanced). Leaving this blank does not fail validation -- it falls back to the certificate's own domain name, which is almost never the device's address, so set it explicitly. Defaults to ``.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"acme_jetkvm_user": schema.StringAttribute{
				MarkdownDescription: "The username to login to the JetKVM device via SSH. JetKVM only supports the \"root\" account for SSH access. Defaults to \"root\". Defaults to `root`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("root"),
			},
			"acme_jetkvm_port": schema.Int64Attribute{
				MarkdownDescription: "SSH server port on the JetKVM device. Defaults to \"22\". Defaults to `22`.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(22),
			},
			"acme_jetkvm_reboot": schema.BoolAttribute{
				MarkdownDescription: "JetKVM does not hot-reload a new \"Custom\" certificate -- a full device reboot is required to apply it, which briefly drops any active KVM-over-IP session. If checked (default), the device is rebooted automatically right after the certificate is uploaded, so an unattended renewal actually takes effect. Uncheck to upload only and apply/verify manually instead. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the ACME automation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func actionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "ACME automations (**Services ▸ ACME Client ▸ Automations**, plugin `os-acme-client`): commands run after a certificate is issued or renewed, e.g. restart the web GUI or upload the certificate over SFTP.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME automation.",
				Required:            true,
			},
			"enabled":                           dschema.BoolAttribute{MarkdownDescription: "Enable this automation.", Computed: true},
			"name":                              dschema.StringAttribute{MarkdownDescription: "Name to identify this automation.", Computed: true},
			"description":                       dschema.StringAttribute{MarkdownDescription: "Description for this automation.", Computed: true},
			"type":                              dschema.StringAttribute{MarkdownDescription: "Pre-defined commands for this automation. One of `configd_restart_gui`, `configd_restart_haproxy`, `configd_restart_nginx`, `configd_reload_caddy`, `configd_upload_sftp`, `configd_remote_ssh`, `acme_fritzbox`, `acme_panos`, `acme_proxmoxbs`, `acme_proxmoxve`, `acme_ruckus`, `acme_vault`, `acme_synology_dsm`, `acme_truenas`, `acme_truenas_ws`, `acme_truenas_websocat`, `acme_zyxel_gs1900`, `acme_unifi`, `acme_jetkvm`, `configd_generic`.", Computed: true},
			"sftp_host":                         dschema.StringAttribute{MarkdownDescription: "IP address or hostname of the SFTP server.", Computed: true},
			"sftp_host_key":                     dschema.StringAttribute{MarkdownDescription: "SFTP server host key, formatted as in 'known_hosts'. Leave blank to auto accept host key on first connect (not as secure as specifying it).", Computed: true},
			"sftp_port":                         dschema.Int64Attribute{MarkdownDescription: "SFTP server port. Leave blank to use default \"22\".", Computed: true},
			"sftp_user":                         dschema.StringAttribute{MarkdownDescription: "The username to login to the SFTP server.", Computed: true},
			"sftp_identity_type":                dschema.StringAttribute{MarkdownDescription: "The type of identify to present to the SFTP server for authorization. Select 'none' to use default \"ECDSA\". One of `ecdsa`, `rsa`, `ed25519`.", Computed: true},
			"sftp_remote_path":                  dschema.StringAttribute{MarkdownDescription: "Path on the SFTP server to change to after login. The path can be absolute or relative to home and must exist. Leave blank to not change path after login.", Computed: true},
			"sftp_chgrp":                        dschema.StringAttribute{MarkdownDescription: "Unix group id to apply to all uploaded files. Leave blank to not change the group.", Computed: true},
			"sftp_chmod":                        dschema.StringAttribute{MarkdownDescription: "Unix permission to apply to uploaded public keys. Leave blank to use default \"0440\".", Computed: true},
			"sftp_chmod_key":                    dschema.StringAttribute{MarkdownDescription: "Unix permission to apply to uploaded private keys. Leave blank to use default \"0400\".", Computed: true},
			"sftp_modtime":                      dschema.BoolAttribute{MarkdownDescription: "Preserves modification times from the source file. Note that this is not supported by all SFTP servers and may cause the upload to fail, e.g. on VMware.", Computed: true},
			"sftp_filename_cert":                dschema.StringAttribute{MarkdownDescription: "Name template for the public certificate. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/cert.pem\".", Computed: true},
			"sftp_filename_key":                 dschema.StringAttribute{MarkdownDescription: "Name template for the certificate's private key. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/key.pem\".", Computed: true},
			"sftp_filename_ca":                  dschema.StringAttribute{MarkdownDescription: "Name template for the public certificate chain file. Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/ca.pem\".", Computed: true},
			"sftp_filename_fullchain":           dschema.StringAttribute{MarkdownDescription: "Name template for the public certificate fullchain file (cert + ca). Placeholders \"{{name}}\" and \"%s\" are replaced by the name of the certificate being uploaded. Leave blank to use default \"{{name}}/fullchain.pem\".", Computed: true},
			"remote_ssh_host":                   dschema.StringAttribute{MarkdownDescription: "IP address or hostname of the SSH server.", Computed: true},
			"remote_ssh_host_key":               dschema.StringAttribute{MarkdownDescription: "remote_ssh_host_key.", Computed: true},
			"remote_ssh_port":                   dschema.Int64Attribute{MarkdownDescription: "SSH server port. Leave blank to use default \"22\".", Computed: true},
			"remote_ssh_user":                   dschema.StringAttribute{MarkdownDescription: "The username to login to the SSH server.", Computed: true},
			"remote_ssh_identity_type":          dschema.StringAttribute{MarkdownDescription: "The type of identify to present to the SSH server for authorization. Select 'none' to use default \"ECDSA\". One of `ecdsa`, `rsa`, `ed25519`.", Computed: true},
			"remote_ssh_command":                dschema.StringAttribute{MarkdownDescription: "The command to execute on the SSH server.", Computed: true},
			"configd":                           dschema.StringAttribute{MarkdownDescription: "configd.", Computed: true},
			"configd_generic_command":           dschema.StringAttribute{MarkdownDescription: "Select a pre-defined system command which should be run.", Computed: true},
			"acme_synology_dsm_hostname":        dschema.StringAttribute{MarkdownDescription: "Hostname of IP address of the Synology DSM, i.e. synology.example.com or 192.168.0.1.", Computed: true},
			"acme_synology_dsm_port":            dschema.Int64Attribute{MarkdownDescription: "Port that will be used when connecting to Synology DSM.", Computed: true},
			"acme_synology_dsm_scheme":          dschema.StringAttribute{MarkdownDescription: "Connection scheme that will be used when uploading certificates to Synology DSM. One of `http`, `https`.", Computed: true},
			"acme_synology_dsm_username":        dschema.StringAttribute{MarkdownDescription: "Username to login, must be an administrator.", Computed: true},
			"acme_synology_dsm_password":        dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"acme_synology_dsm_create":          dschema.BoolAttribute{MarkdownDescription: "This option ensures that a new certificate is created in Synology DSM if it does not exist yet. If unchecked only existing certificates will be updated.", Computed: true},
			"acme_synology_dsm_deviceid":        dschema.StringAttribute{MarkdownDescription: "If Synology DSM has OTP enabled, then the device ID has to be provided so that no OTP is required when running the automation.", Computed: true},
			"acme_synology_dsm_devicename":      dschema.StringAttribute{MarkdownDescription: "If Synology DSM has OTP enabled, then the device name has to be provided so that no OTP is required when running the automation.", Computed: true},
			"acme_synology_dsm_otpcode":         dschema.StringAttribute{MarkdownDescription: "If Synology DSM has OTP enabled, then a OTP may be required.", Computed: true},
			"acme_fritzbox_url":                 dschema.StringAttribute{MarkdownDescription: "URL of the router, i.e. https://fritzbox.example.com.", Computed: true},
			"acme_fritzbox_username":            dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"acme_fritzbox_password":            dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"acme_panos_username":               dschema.StringAttribute{MarkdownDescription: "Username.", Computed: true},
			"acme_panos_password":               dschema.StringAttribute{MarkdownDescription: "Password.", Computed: true, Sensitive: true},
			"acme_panos_host":                   dschema.StringAttribute{MarkdownDescription: "Host.", Computed: true},
			"acme_proxmoxve_user":               dschema.StringAttribute{MarkdownDescription: "The user who owns the API key. Defaults to root.", Computed: true},
			"acme_proxmoxve_server":             dschema.StringAttribute{MarkdownDescription: "The hostname of the proxmox ve node.", Computed: true},
			"acme_proxmoxve_port":               dschema.Int64Attribute{MarkdownDescription: "The port number the management interface is on. Defaults to 8006.", Computed: true},
			"acme_proxmoxve_nodename":           dschema.StringAttribute{MarkdownDescription: "The name of the node we will be connecting to.", Computed: true},
			"acme_proxmoxve_realm":              dschema.StringAttribute{MarkdownDescription: "The authentication realm the user authenticates with. Defaults to pam.", Computed: true},
			"acme_proxmoxve_tokenid":            dschema.StringAttribute{MarkdownDescription: "The name of the API token created for the user account. Defaults to acme.", Computed: true, Sensitive: true},
			"acme_proxmoxve_tokenkey":           dschema.StringAttribute{MarkdownDescription: "The API token. Required.", Computed: true, Sensitive: true},
			"acme_proxmoxbs_user":               dschema.StringAttribute{MarkdownDescription: "The user who owns the API key. Defaults to root.", Computed: true},
			"acme_proxmoxbs_server":             dschema.StringAttribute{MarkdownDescription: "The hostname of the proxmox BS node.", Computed: true},
			"acme_proxmoxbs_port":               dschema.Int64Attribute{MarkdownDescription: "The port number the management interface is on. Defaults to 8007.", Computed: true},
			"acme_proxmoxbs_nodename":           dschema.StringAttribute{MarkdownDescription: "acme_proxmoxbs_nodename.", Computed: true},
			"acme_proxmoxbs_realm":              dschema.StringAttribute{MarkdownDescription: "The authentication realm the user authenticates with. Defaults to pam.", Computed: true},
			"acme_proxmoxbs_tokenid":            dschema.StringAttribute{MarkdownDescription: "The name of the API token created for the user account. Defaults to acme.", Computed: true, Sensitive: true},
			"acme_proxmoxbs_tokenkey":           dschema.StringAttribute{MarkdownDescription: "The API token. Required.", Computed: true, Sensitive: true},
			"acme_ruckus_host":                  dschema.StringAttribute{MarkdownDescription: "Ruckus Host.", Computed: true},
			"acme_ruckus_user":                  dschema.StringAttribute{MarkdownDescription: "Ruckus User.", Computed: true},
			"acme_ruckus_pass":                  dschema.StringAttribute{MarkdownDescription: "Ruckus Password.", Computed: true, Sensitive: true},
			"acme_truenas_apikey":               dschema.StringAttribute{MarkdownDescription: "API key generated in the TrueNAS web UI.", Computed: true, Sensitive: true},
			"acme_truenas_hostname":             dschema.StringAttribute{MarkdownDescription: "Hostname or IP address of TrueNAS Server.", Computed: true},
			"acme_truenas_scheme":               dschema.StringAttribute{MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Core Server. One of `http`, `https`.", Computed: true},
			"acme_truenas_ws_apikey":            dschema.StringAttribute{MarkdownDescription: "API key generated in the TrueNAS web UI.", Computed: true, Sensitive: true},
			"acme_truenas_ws_hostname":          dschema.StringAttribute{MarkdownDescription: "Hostname or IP address of TrueNAS Server.", Computed: true},
			"acme_truenas_ws_protocol":          dschema.StringAttribute{MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Server. One of `ws`, `wss`.", Computed: true},
			"acme_truenas_websocat_apikey":      dschema.StringAttribute{MarkdownDescription: "API key generated in the TrueNAS web UI.", Computed: true, Sensitive: true},
			"acme_truenas_websocat_hostname":    dschema.StringAttribute{MarkdownDescription: "Hostname or IP address of TrueNAS Server.", Computed: true},
			"acme_truenas_websocat_protocol":    dschema.StringAttribute{MarkdownDescription: "Connection scheme that will be used when uploading certificates to TrueNAS Server. One of `ws`, `wss`.", Computed: true},
			"acme_truenas_websocat_port":        dschema.Int64Attribute{MarkdownDescription: "Port of TrueNAS HTTP Server.", Computed: true},
			"acme_truenas_websocat_update_ftp":  dschema.StringAttribute{MarkdownDescription: "Upload certificate to TrueNAS FTP Server. One of `no`, `yes`.", Computed: true},
			"acme_truenas_websocat_update_apps": dschema.StringAttribute{MarkdownDescription: "Upload certificate to TrueNAS Server ix-Apps. One of `no`, `yes`.", Computed: true},
			"acme_truenasws_apikey":             dschema.StringAttribute{MarkdownDescription: "acme_truenasws_apikey.", Computed: true, Sensitive: true},
			"acme_truenasws_hostname":           dschema.StringAttribute{MarkdownDescription: "acme_truenasws_hostname.", Computed: true},
			"acme_truenasws_protocol":           dschema.StringAttribute{MarkdownDescription: "acme_truenasws_protocol. One of `ws`, `wss`.", Computed: true},
			"acme_unifi_keystore":               dschema.StringAttribute{MarkdownDescription: "Path to the Unifi keystore file in the local filesystem, i.e. /usr/local/share/java/unifi/data/keystore.", Computed: true, Sensitive: true},
			"acme_vault_url":                    dschema.StringAttribute{MarkdownDescription: "URL of the Vault, i.e. http://vault.example.com:8200.", Computed: true},
			"acme_vault_prefix":                 dschema.StringAttribute{MarkdownDescription: "This specifies the prefix path in Vault. If you select KV v2 you need to add .../data/... between the secret-mount-path and the path. Example: v1 prefix path: secret/acme, v2 prefix path: secret/data/acme.", Computed: true},
			"acme_vault_token":                  dschema.StringAttribute{MarkdownDescription: "This specifies the Vault token to authenticate with.", Computed: true, Sensitive: true},
			"acme_vault_kvv2":                   dschema.BoolAttribute{MarkdownDescription: "If checked version 2 of the kv store will be used, otherwise version 1. If you use v2 please consider the comment in the field \"Vault Prefix\".", Computed: true},
			"acme_zyxel_gs1900_host":            dschema.StringAttribute{MarkdownDescription: "Hostname or IP adress of a Zyxel GS1900 series switch.", Computed: true},
			"acme_zyxel_gs1900_user":            dschema.StringAttribute{MarkdownDescription: "A user configured as an adminstrator. Defaults to admin.", Computed: true},
			"acme_zyxel_gs1900_password":        dschema.StringAttribute{MarkdownDescription: "Switch Admin Password.", Computed: true, Sensitive: true},
			"acme_zyxel_gs1900_insecure":        dschema.BoolAttribute{MarkdownDescription: "If checked the deployment will be run in insecure mode and will skip certificate validation when connecting to the switch.", Computed: true},
			"acme_zyxel_gs1900_reboot":          dschema.BoolAttribute{MarkdownDescription: "If checked the switch will be rebooted once the certificate has been successfully installed.", Computed: true},
			"acme_jetkvm_host":                  dschema.StringAttribute{MarkdownDescription: "IP address or hostname of the JetKVM device. Required first: the device must already have \"HTTPS Mode\" set to \"Custom\" in its own web UI (Settings > Network) -- this automation only writes the certificate/key files (and optionally reboots); it does not switch HTTPS mode for you, and the deploy will fail if it isn't already set (check the ACME Client log for the exact reason). Also requires SSH access to already work non-interactively as the configured user (JetKVM only supports key-based SSH authentication -- enable \"Developer Mode\" on the device and paste this firewall's SSH public key, e.g. from /root/.ssh/id_rsa.pub, into its \"SSH Keys\" field under Settings > Advanced). Leaving this blank does not fail validation -- it falls back to the certificate's own domain name, which is almost never the device's address, so set it explicitly.", Computed: true},
			"acme_jetkvm_user":                  dschema.StringAttribute{MarkdownDescription: "The username to login to the JetKVM device via SSH. JetKVM only supports the \"root\" account for SSH access. Defaults to \"root\".", Computed: true},
			"acme_jetkvm_port":                  dschema.Int64Attribute{MarkdownDescription: "SSH server port on the JetKVM device. Defaults to \"22\".", Computed: true},
			"acme_jetkvm_reboot":                dschema.BoolAttribute{MarkdownDescription: "JetKVM does not hot-reload a new \"Custom\" certificate -- a full device reboot is required to apply it, which briefly drops any active KVM-over-IP session. If checked (default), the device is rebooted automatically right after the certificate is uploaded, so an unattended renewal actually takes effect. Uncheck to upload only and apply/verify manually instead.", Computed: true},
		},
	}
}

func convertActionSchemaToStruct(d *actionResourceModel) (*acmeclient.Action, error) {
	return &acmeclient.Action{
		Enabled:                       tools.BoolToString(d.Enabled.ValueBool()),
		Name:                          d.Name.ValueString(),
		Description:                   d.Description.ValueString(),
		Type:                          api.SelectedMap(d.Type.ValueString()),
		SftpHost:                      d.SftpHost.ValueString(),
		SftpHostKey:                   d.SftpHostKey.ValueString(),
		SftpPort:                      tools.Int64ToString(d.SftpPort.ValueInt64()),
		SftpUser:                      d.SftpUser.ValueString(),
		SftpIdentityType:              api.SelectedMap(d.SftpIdentityType.ValueString()),
		SftpRemotePath:                d.SftpRemotePath.ValueString(),
		SftpChgrp:                     d.SftpChgrp.ValueString(),
		SftpChmod:                     d.SftpChmod.ValueString(),
		SftpChmodKey:                  d.SftpChmodKey.ValueString(),
		SftpModtime:                   tools.BoolToString(d.SftpModtime.ValueBool()),
		SftpFilenameCert:              d.SftpFilenameCert.ValueString(),
		SftpFilenameKey:               d.SftpFilenameKey.ValueString(),
		SftpFilenameCa:                d.SftpFilenameCa.ValueString(),
		SftpFilenameFullchain:         d.SftpFilenameFullchain.ValueString(),
		RemoteSshHost:                 d.RemoteSshHost.ValueString(),
		RemoteSshHostKey:              d.RemoteSshHostKey.ValueString(),
		RemoteSshPort:                 tools.Int64ToString(d.RemoteSshPort.ValueInt64()),
		RemoteSshUser:                 d.RemoteSshUser.ValueString(),
		RemoteSshIdentityType:         api.SelectedMap(d.RemoteSshIdentityType.ValueString()),
		RemoteSshCommand:              d.RemoteSshCommand.ValueString(),
		Configd:                       api.SelectedMap(d.Configd.ValueString()),
		ConfigdGenericCommand:         api.SelectedMap(d.ConfigdGenericCommand.ValueString()),
		AcmeSynologyDsmHostname:       d.AcmeSynologyDsmHostname.ValueString(),
		AcmeSynologyDsmPort:           tools.Int64ToString(d.AcmeSynologyDsmPort.ValueInt64()),
		AcmeSynologyDsmScheme:         api.SelectedMap(d.AcmeSynologyDsmScheme.ValueString()),
		AcmeSynologyDsmUsername:       d.AcmeSynologyDsmUsername.ValueString(),
		AcmeSynologyDsmPassword:       d.AcmeSynologyDsmPassword.ValueString(),
		AcmeSynologyDsmCreate:         tools.BoolToString(d.AcmeSynologyDsmCreate.ValueBool()),
		AcmeSynologyDsmDeviceid:       d.AcmeSynologyDsmDeviceid.ValueString(),
		AcmeSynologyDsmDevicename:     d.AcmeSynologyDsmDevicename.ValueString(),
		AcmeSynologyDsmOtpcode:        d.AcmeSynologyDsmOtpcode.ValueString(),
		AcmeFritzboxUrl:               d.AcmeFritzboxURL.ValueString(),
		AcmeFritzboxUsername:          d.AcmeFritzboxUsername.ValueString(),
		AcmeFritzboxPassword:          d.AcmeFritzboxPassword.ValueString(),
		AcmePanosUsername:             d.AcmePanosUsername.ValueString(),
		AcmePanosPassword:             d.AcmePanosPassword.ValueString(),
		AcmePanosHost:                 d.AcmePanosHost.ValueString(),
		AcmeProxmoxveUser:             d.AcmeProxmoxveUser.ValueString(),
		AcmeProxmoxveServer:           d.AcmeProxmoxveServer.ValueString(),
		AcmeProxmoxvePort:             tools.Int64ToString(d.AcmeProxmoxvePort.ValueInt64()),
		AcmeProxmoxveNodename:         d.AcmeProxmoxveNodename.ValueString(),
		AcmeProxmoxveRealm:            d.AcmeProxmoxveRealm.ValueString(),
		AcmeProxmoxveTokenid:          d.AcmeProxmoxveTokenid.ValueString(),
		AcmeProxmoxveTokenkey:         d.AcmeProxmoxveTokenkey.ValueString(),
		AcmeProxmoxbsUser:             d.AcmeProxmoxbsUser.ValueString(),
		AcmeProxmoxbsServer:           d.AcmeProxmoxbsServer.ValueString(),
		AcmeProxmoxbsPort:             tools.Int64ToString(d.AcmeProxmoxbsPort.ValueInt64()),
		AcmeProxmoxbsNodename:         d.AcmeProxmoxbsNodename.ValueString(),
		AcmeProxmoxbsRealm:            d.AcmeProxmoxbsRealm.ValueString(),
		AcmeProxmoxbsTokenid:          d.AcmeProxmoxbsTokenid.ValueString(),
		AcmeProxmoxbsTokenkey:         d.AcmeProxmoxbsTokenkey.ValueString(),
		AcmeRuckusHost:                d.AcmeRuckusHost.ValueString(),
		AcmeRuckusUser:                d.AcmeRuckusUser.ValueString(),
		AcmeRuckusPass:                d.AcmeRuckusPass.ValueString(),
		AcmeTruenasApikey:             d.AcmeTruenasApikey.ValueString(),
		AcmeTruenasHostname:           d.AcmeTruenasHostname.ValueString(),
		AcmeTruenasScheme:             api.SelectedMap(d.AcmeTruenasScheme.ValueString()),
		AcmeTruenasWsApikey:           d.AcmeTruenasWsApikey.ValueString(),
		AcmeTruenasWsHostname:         d.AcmeTruenasWsHostname.ValueString(),
		AcmeTruenasWsProtocol:         api.SelectedMap(d.AcmeTruenasWsProtocol.ValueString()),
		AcmeTruenasWebsocatApikey:     d.AcmeTruenasWebsocatApikey.ValueString(),
		AcmeTruenasWebsocatHostname:   d.AcmeTruenasWebsocatHostname.ValueString(),
		AcmeTruenasWebsocatProtocol:   api.SelectedMap(d.AcmeTruenasWebsocatProtocol.ValueString()),
		AcmeTruenasWebsocatPort:       tools.Int64ToString(d.AcmeTruenasWebsocatPort.ValueInt64()),
		AcmeTruenasWebsocatUpdateFtp:  api.SelectedMap(d.AcmeTruenasWebsocatUpdateFtp.ValueString()),
		AcmeTruenasWebsocatUpdateApps: api.SelectedMap(d.AcmeTruenasWebsocatUpdateApps.ValueString()),
		AcmeTruenaswsApikey:           d.AcmeTruenaswsApikey.ValueString(),
		AcmeTruenaswsHostname:         d.AcmeTruenaswsHostname.ValueString(),
		AcmeTruenaswsProtocol:         api.SelectedMap(d.AcmeTruenaswsProtocol.ValueString()),
		AcmeUnifiKeystore:             d.AcmeUnifiKeystore.ValueString(),
		AcmeVaultUrl:                  d.AcmeVaultURL.ValueString(),
		AcmeVaultPrefix:               d.AcmeVaultPrefix.ValueString(),
		AcmeVaultToken:                d.AcmeVaultToken.ValueString(),
		AcmeVaultKvv2:                 tools.BoolToString(d.AcmeVaultKvv2.ValueBool()),
		AcmeZyxelGs1900Host:           d.AcmeZyxelGs1900Host.ValueString(),
		AcmeZyxelGs1900User:           d.AcmeZyxelGs1900User.ValueString(),
		AcmeZyxelGs1900Password:       d.AcmeZyxelGs1900Password.ValueString(),
		AcmeZyxelGs1900Insecure:       tools.BoolToString(d.AcmeZyxelGs1900Insecure.ValueBool()),
		AcmeZyxelGs1900Reboot:         tools.BoolToString(d.AcmeZyxelGs1900Reboot.ValueBool()),
		AcmeJetkvmHost:                d.AcmeJetkvmHost.ValueString(),
		AcmeJetkvmUser:                d.AcmeJetkvmUser.ValueString(),
		AcmeJetkvmPort:                tools.Int64ToString(d.AcmeJetkvmPort.ValueInt64()),
		AcmeJetkvmReboot:              tools.BoolToString(d.AcmeJetkvmReboot.ValueBool()),
	}, nil
}

func convertActionStructToSchema(d *acmeclient.Action) (*actionResourceModel, error) {
	return &actionResourceModel{
		Enabled:                       types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:                          types.StringValue(d.Name),
		Description:                   types.StringValue(d.Description),
		Type:                          types.StringValue(d.Type.String()),
		SftpHost:                      types.StringValue(d.SftpHost),
		SftpHostKey:                   types.StringValue(d.SftpHostKey),
		SftpPort:                      types.Int64Value(tools.StringToInt64(d.SftpPort)),
		SftpUser:                      types.StringValue(d.SftpUser),
		SftpIdentityType:              types.StringValue(d.SftpIdentityType.String()),
		SftpRemotePath:                types.StringValue(d.SftpRemotePath),
		SftpChgrp:                     types.StringValue(d.SftpChgrp),
		SftpChmod:                     types.StringValue(d.SftpChmod),
		SftpChmodKey:                  types.StringValue(d.SftpChmodKey),
		SftpModtime:                   types.BoolValue(tools.StringToBool(d.SftpModtime)),
		SftpFilenameCert:              types.StringValue(d.SftpFilenameCert),
		SftpFilenameKey:               types.StringValue(d.SftpFilenameKey),
		SftpFilenameCa:                types.StringValue(d.SftpFilenameCa),
		SftpFilenameFullchain:         types.StringValue(d.SftpFilenameFullchain),
		RemoteSshHost:                 types.StringValue(d.RemoteSshHost),
		RemoteSshHostKey:              types.StringValue(d.RemoteSshHostKey),
		RemoteSshPort:                 types.Int64Value(tools.StringToInt64(d.RemoteSshPort)),
		RemoteSshUser:                 types.StringValue(d.RemoteSshUser),
		RemoteSshIdentityType:         types.StringValue(d.RemoteSshIdentityType.String()),
		RemoteSshCommand:              types.StringValue(d.RemoteSshCommand),
		Configd:                       types.StringValue(d.Configd.String()),
		ConfigdGenericCommand:         types.StringValue(d.ConfigdGenericCommand.String()),
		AcmeSynologyDsmHostname:       types.StringValue(d.AcmeSynologyDsmHostname),
		AcmeSynologyDsmPort:           types.Int64Value(tools.StringToInt64(d.AcmeSynologyDsmPort)),
		AcmeSynologyDsmScheme:         types.StringValue(d.AcmeSynologyDsmScheme.String()),
		AcmeSynologyDsmUsername:       types.StringValue(d.AcmeSynologyDsmUsername),
		AcmeSynologyDsmPassword:       types.StringValue(d.AcmeSynologyDsmPassword),
		AcmeSynologyDsmCreate:         types.BoolValue(tools.StringToBool(d.AcmeSynologyDsmCreate)),
		AcmeSynologyDsmDeviceid:       types.StringValue(d.AcmeSynologyDsmDeviceid),
		AcmeSynologyDsmDevicename:     types.StringValue(d.AcmeSynologyDsmDevicename),
		AcmeSynologyDsmOtpcode:        types.StringValue(d.AcmeSynologyDsmOtpcode),
		AcmeFritzboxURL:               types.StringValue(d.AcmeFritzboxUrl),
		AcmeFritzboxUsername:          types.StringValue(d.AcmeFritzboxUsername),
		AcmeFritzboxPassword:          types.StringValue(d.AcmeFritzboxPassword),
		AcmePanosUsername:             types.StringValue(d.AcmePanosUsername),
		AcmePanosPassword:             types.StringValue(d.AcmePanosPassword),
		AcmePanosHost:                 types.StringValue(d.AcmePanosHost),
		AcmeProxmoxveUser:             types.StringValue(d.AcmeProxmoxveUser),
		AcmeProxmoxveServer:           types.StringValue(d.AcmeProxmoxveServer),
		AcmeProxmoxvePort:             types.Int64Value(tools.StringToInt64(d.AcmeProxmoxvePort)),
		AcmeProxmoxveNodename:         types.StringValue(d.AcmeProxmoxveNodename),
		AcmeProxmoxveRealm:            types.StringValue(d.AcmeProxmoxveRealm),
		AcmeProxmoxveTokenid:          types.StringValue(d.AcmeProxmoxveTokenid),
		AcmeProxmoxveTokenkey:         types.StringValue(d.AcmeProxmoxveTokenkey),
		AcmeProxmoxbsUser:             types.StringValue(d.AcmeProxmoxbsUser),
		AcmeProxmoxbsServer:           types.StringValue(d.AcmeProxmoxbsServer),
		AcmeProxmoxbsPort:             types.Int64Value(tools.StringToInt64(d.AcmeProxmoxbsPort)),
		AcmeProxmoxbsNodename:         types.StringValue(d.AcmeProxmoxbsNodename),
		AcmeProxmoxbsRealm:            types.StringValue(d.AcmeProxmoxbsRealm),
		AcmeProxmoxbsTokenid:          types.StringValue(d.AcmeProxmoxbsTokenid),
		AcmeProxmoxbsTokenkey:         types.StringValue(d.AcmeProxmoxbsTokenkey),
		AcmeRuckusHost:                types.StringValue(d.AcmeRuckusHost),
		AcmeRuckusUser:                types.StringValue(d.AcmeRuckusUser),
		AcmeRuckusPass:                types.StringValue(d.AcmeRuckusPass),
		AcmeTruenasApikey:             types.StringValue(d.AcmeTruenasApikey),
		AcmeTruenasHostname:           types.StringValue(d.AcmeTruenasHostname),
		AcmeTruenasScheme:             types.StringValue(d.AcmeTruenasScheme.String()),
		AcmeTruenasWsApikey:           types.StringValue(d.AcmeTruenasWsApikey),
		AcmeTruenasWsHostname:         types.StringValue(d.AcmeTruenasWsHostname),
		AcmeTruenasWsProtocol:         types.StringValue(d.AcmeTruenasWsProtocol.String()),
		AcmeTruenasWebsocatApikey:     types.StringValue(d.AcmeTruenasWebsocatApikey),
		AcmeTruenasWebsocatHostname:   types.StringValue(d.AcmeTruenasWebsocatHostname),
		AcmeTruenasWebsocatProtocol:   types.StringValue(d.AcmeTruenasWebsocatProtocol.String()),
		AcmeTruenasWebsocatPort:       types.Int64Value(tools.StringToInt64(d.AcmeTruenasWebsocatPort)),
		AcmeTruenasWebsocatUpdateFtp:  types.StringValue(d.AcmeTruenasWebsocatUpdateFtp.String()),
		AcmeTruenasWebsocatUpdateApps: types.StringValue(d.AcmeTruenasWebsocatUpdateApps.String()),
		AcmeTruenaswsApikey:           types.StringValue(d.AcmeTruenaswsApikey),
		AcmeTruenaswsHostname:         types.StringValue(d.AcmeTruenaswsHostname),
		AcmeTruenaswsProtocol:         types.StringValue(d.AcmeTruenaswsProtocol.String()),
		AcmeUnifiKeystore:             types.StringValue(d.AcmeUnifiKeystore),
		AcmeVaultURL:                  types.StringValue(d.AcmeVaultUrl),
		AcmeVaultPrefix:               types.StringValue(d.AcmeVaultPrefix),
		AcmeVaultToken:                types.StringValue(d.AcmeVaultToken),
		AcmeVaultKvv2:                 types.BoolValue(tools.StringToBool(d.AcmeVaultKvv2)),
		AcmeZyxelGs1900Host:           types.StringValue(d.AcmeZyxelGs1900Host),
		AcmeZyxelGs1900User:           types.StringValue(d.AcmeZyxelGs1900User),
		AcmeZyxelGs1900Password:       types.StringValue(d.AcmeZyxelGs1900Password),
		AcmeZyxelGs1900Insecure:       types.BoolValue(tools.StringToBool(d.AcmeZyxelGs1900Insecure)),
		AcmeZyxelGs1900Reboot:         types.BoolValue(tools.StringToBool(d.AcmeZyxelGs1900Reboot)),
		AcmeJetkvmHost:                types.StringValue(d.AcmeJetkvmHost),
		AcmeJetkvmUser:                types.StringValue(d.AcmeJetkvmUser),
		AcmeJetkvmPort:                types.Int64Value(tools.StringToInt64(d.AcmeJetkvmPort)),
		AcmeJetkvmReboot:              types.BoolValue(tools.StringToBool(d.AcmeJetkvmReboot)),
	}, nil
}

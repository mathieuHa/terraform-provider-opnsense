package telegraf

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/telegraf"
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

type outputSettingsResourceModel struct {
	InfluxEnable                    types.Bool   `tfsdk:"influx_enable"`
	InfluxUrl                       types.String `tfsdk:"influx_url"`
	InfluxDatabase                  types.String `tfsdk:"influx_database"`
	InfluxTimeout                   types.Int64  `tfsdk:"influx_timeout"`
	InfluxNamePrefix                types.String `tfsdk:"influx_name_prefix"`
	InfluxUsername                  types.String `tfsdk:"influx_username"`
	InfluxPassword                  types.String `tfsdk:"influx_password"`
	InfluxInsecureSkipVerify        types.Bool   `tfsdk:"influx_insecure_skip_verify"`
	GraphiteEnable                  types.Bool   `tfsdk:"graphite_enable"`
	GraphiteServer                  types.String `tfsdk:"graphite_server"`
	GraphitePrefix                  types.String `tfsdk:"graphite_prefix"`
	GraphiteTemplate                types.String `tfsdk:"graphite_template"`
	GraphiteSslDisable              types.Bool   `tfsdk:"graphite_ssl_disable"`
	GraphiteVerify                  types.Bool   `tfsdk:"graphite_verify"`
	GraphiteTagsupport              types.Bool   `tfsdk:"graphite_tagsupport"`
	GraylogEnable                   types.Bool   `tfsdk:"graylog_enable"`
	GraylogServer                   types.String `tfsdk:"graylog_server"`
	ElasticEnable                   types.Bool   `tfsdk:"elastic_enable"`
	ElasticUrl                      types.String `tfsdk:"elastic_url"`
	ElasticUsername                 types.String `tfsdk:"elastic_username"`
	ElasticPassword                 types.String `tfsdk:"elastic_password"`
	ElasticTimeout                  types.Int64  `tfsdk:"elastic_timeout"`
	ElasticIndexname                types.String `tfsdk:"elastic_indexname"`
	PrometheusEnable                types.Bool   `tfsdk:"prometheus_enable"`
	PrometheusListen                types.Int64  `tfsdk:"prometheus_listen"`
	PrometheusExclude               types.Set    `tfsdk:"prometheus_exclude"`
	PrometheusStringaslabel         types.Bool   `tfsdk:"prometheus_stringaslabel"`
	InfluxV2Enable                  types.Bool   `tfsdk:"influx_v2_enable"`
	InfluxV2Url                     types.String `tfsdk:"influx_v2_url"`
	InfluxV2Token                   types.String `tfsdk:"influx_v2_token"`
	InfluxV2Organization            types.String `tfsdk:"influx_v2_organization"`
	InfluxV2Bucket                  types.String `tfsdk:"influx_v2_bucket"`
	InfluxV2InsecureSkipVerify      types.Bool   `tfsdk:"influx_v2_insecure_skip_verify"`
	InfluxV2Timeout                 types.Int64  `tfsdk:"influx_v2_timeout"`
	DatadogEnable                   types.Bool   `tfsdk:"datadog_enable"`
	DatadogUrl                      types.String `tfsdk:"datadog_url"`
	DatadogApikey                   types.String `tfsdk:"datadog_apikey"`
	MqttEnable                      types.Bool   `tfsdk:"mqtt_enable"`
	MqttTopicPrefix                 types.String `tfsdk:"mqtt_topic_prefix"`
	MqttTopic                       types.String `tfsdk:"mqtt_topic"`
	MqttServers                     types.Set    `tfsdk:"mqtt_servers"`
	MqttInsecureSkipVerify          types.Bool   `tfsdk:"mqtt_insecure_skip_verify"`
	MqttClientId                    types.String `tfsdk:"mqtt_client_id"`
	MqttQos                         types.String `tfsdk:"mqtt_qos"`
	MqttRetain                      types.Bool   `tfsdk:"mqtt_retain"`
	MqttTimeout                     types.Int64  `tfsdk:"mqtt_timeout"`
	MqttUsername                    types.String `tfsdk:"mqtt_username"`
	MqttPassword                    types.String `tfsdk:"mqtt_password"`
	MqttLayout                      types.String `tfsdk:"mqtt_layout"`
	MqttDataFormat                  types.String `tfsdk:"mqtt_data_format"`
	OpentelemetryEnable             types.Bool   `tfsdk:"opentelemetry_enable"`
	OpentelemetryServer             types.String `tfsdk:"opentelemetry_server"`
	OpentelemetryInsecureSkipVerify types.Bool   `tfsdk:"opentelemetry_insecure_skip_verify"`
	OpentelemetryTimeout            types.Int64  `tfsdk:"opentelemetry_timeout"`
	OpentelemetryCompression        types.String `tfsdk:"opentelemetry_compression"`
	OpentelemetryHeaderName         types.String `tfsdk:"opentelemetry_header_name"`
	OpentelemetryHeaderValue        types.String `tfsdk:"opentelemetry_header_value"`

	Id types.String `tfsdk:"id"`
}

func outputSettingsAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"influx_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable InfluxDB as output. Format is without square brackets, just like http://192.168.0.1:8086.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"influx_url": schema.StringAttribute{
			MarkdownDescription: "Set the URL where metrics shoud be sent to.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_database": schema.StringAttribute{
			MarkdownDescription: "Set the name of the database on InfluxDB.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_timeout": schema.Int64Attribute{
			MarkdownDescription: "Write timeout (for the InfluxDB client), formatted as a string. If not provided, will default to 5s. 0s means no timeout (not recommended). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"influx_name_prefix": schema.StringAttribute{
			MarkdownDescription: "Specifies a prefix to attach to the measurement name, formatted as a string.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_username": schema.StringAttribute{
			MarkdownDescription: "Set the username for authentication.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_password": schema.StringAttribute{
			MarkdownDescription: "Set the password for authentication.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_insecure_skip_verify": schema.BoolAttribute{
			MarkdownDescription: "This will skip chain and host verification.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graphite_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable Graphite as output. Format is without square brackets, just like localhost:2003.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graphite_server": schema.StringAttribute{
			MarkdownDescription: "Set the IP and port where metrics shoud be sent to. Format: IP:port.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"graphite_prefix": schema.StringAttribute{
			MarkdownDescription: "Set an optional prefix to the metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"graphite_template": schema.StringAttribute{
			MarkdownDescription: "Use a specific output template. For usage see https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"graphite_ssl_disable": schema.BoolAttribute{
			MarkdownDescription: "This will disable SSL/TLS connection to Graphite host. Default is encrypted connection.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graphite_verify": schema.BoolAttribute{
			MarkdownDescription: "This will enable verification of a secure connection to Graphite. Default is disabled for compatibility reasons.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graphite_tagsupport": schema.BoolAttribute{
			MarkdownDescription: "This will enable support for tags.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graylog_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable Graylog as output. Format is without square brackets, just like localhost:12201.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"graylog_server": schema.StringAttribute{
			MarkdownDescription: "Set the IP and port where metrics shoud be sent to.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"elastic_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable Elasticsearch as output. Format is without square brackets, just like http://192.168.0.1:9200.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"elastic_url": schema.StringAttribute{
			MarkdownDescription: "Set the URL where metrics shoud be sent to.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"elastic_username": schema.StringAttribute{
			MarkdownDescription: "Optional HTTP basic authentication details for Elasticsearch.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"elastic_password": schema.StringAttribute{
			MarkdownDescription: "Optional HTTP basic authentication details for Elasticsearch.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"elastic_timeout": schema.Int64Attribute{
			MarkdownDescription: "Write timeout (for the Elasticsearch client), formatted as a string. If not provided, will default to 5s. 0s means no timeout (not recommended). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"elastic_indexname": schema.StringAttribute{
			MarkdownDescription: "Set the index name.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"prometheus_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable Prometheus as output.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"prometheus_listen": schema.Int64Attribute{
			MarkdownDescription: "Port to listen on. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"prometheus_exclude": schema.SetAttribute{
			MarkdownDescription: "Collectors to enable, valid entries are \"gocollector\" and \"process\". If unset, both are enabled.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"prometheus_stringaslabel": schema.BoolAttribute{
			MarkdownDescription: "Send string metrics as Prometheus labels.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable InfluxDB v2 as output. Format is without square brackets, just like http://192.168.0.1:9999.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_url": schema.StringAttribute{
			MarkdownDescription: "Set the URL where metrics shoud be sent to.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_token": schema.StringAttribute{
			MarkdownDescription: "Influx v2 authentication token.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_organization": schema.StringAttribute{
			MarkdownDescription: "Set the name of the organization in Influx v2.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_bucket": schema.StringAttribute{
			MarkdownDescription: "Set the name of the bucket in Influx v2.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_insecure_skip_verify": schema.BoolAttribute{
			MarkdownDescription: "This will skip chain and host verification.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"influx_v2_timeout": schema.Int64Attribute{
			MarkdownDescription: "Flush timeout for the v2 Telegraf output in seconds, formatted as a string. If not provided, will default to 5. 0 means no timeout, but is not recommended. Increase if you get timeout errors. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"datadog_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable Datadog output.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"datadog_url": schema.StringAttribute{
			MarkdownDescription: "Set the URL where metrics shoud be sent to. Format is without square brackets, just like https://app.datadoghq.com/api/v1/series.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"datadog_apikey": schema.StringAttribute{
			MarkdownDescription: "Set the API Key for accessing Datadog.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable writes to a MQTT Broker acting as a mqtt Producer.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"mqtt_topic_prefix": schema.StringAttribute{
			MarkdownDescription: "MQTT Topic Prefix - Will be overidden by the Topic if defined below.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_topic": schema.StringAttribute{
			MarkdownDescription: "MQTT Topic for produced messages. if left blank, will default to [ mqtt_topic_prefix/{{ .Hostname }}/{{ .PluginName }} ].",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_servers": schema.SetAttribute{
			MarkdownDescription: "URL of MQTT brokers. Format is without square brackets, just like localhost:8083 or mqtts://server.net:8883. The same credentials will be used when defining multiple brokers.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
		},
		"mqtt_insecure_skip_verify": schema.BoolAttribute{
			MarkdownDescription: "Use TLS, but skip chain and host verification.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"mqtt_client_id": schema.StringAttribute{
			MarkdownDescription: "Client ID, if not set a random ID is generated.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_qos": schema.StringAttribute{
			MarkdownDescription: "QoS policy for messages. 0 = at most once, 1 = at least once, 2 = exactly once. Defaults to 2. One of `0`, `1`, `2`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_retain": schema.BoolAttribute{
			MarkdownDescription: "When selected, metrics will have the RETAIN flag set.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"mqtt_timeout": schema.Int64Attribute{
			MarkdownDescription: "Timeout for write operations. Default is 5s. `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"mqtt_username": schema.StringAttribute{
			MarkdownDescription: "Set the username for authentication.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_password": schema.StringAttribute{
			MarkdownDescription: "Set the password for authentication.",
			Optional:            true,
			Computed:            true,
			Sensitive:           true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_layout": schema.StringAttribute{
			MarkdownDescription: "Data layout to output. Defaults to \"non-batch\". One of `non-batch`, `batch`, `field`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"mqtt_data_format": schema.StringAttribute{
			MarkdownDescription: "Data format to output. Defaults to \"influx\". One of `carbon2`, `cloudevents`, `csv`, `graphite`, `influx`, `json`, `msgpack`, `nowmetric`, `prometheus`, `prometheusremotewrite`, `splunkmetric`, `wavefront`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_enable": schema.BoolAttribute{
			MarkdownDescription: "This will enable OpenTelemetry as output.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_server": schema.StringAttribute{
			MarkdownDescription: "For gRPC, set the address and port, e.g. 192.168.0.1:4317. For HTTP, use a full URL, e.g. https://collector.example.com/v1/metrics.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_insecure_skip_verify": schema.BoolAttribute{
			MarkdownDescription: "This will skip chain and host verification.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_timeout": schema.Int64Attribute{
			MarkdownDescription: "Write timeout, formatted as a string. If not provided, will default to 5s. 0s means no timeout (not recommended). `-1` when unset; set `-1` to clear it.",
			Optional:            true,
			Computed:            true,
			Validators:          []validator.Int64{int64validator.AtLeast(-1)},
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"opentelemetry_compression": schema.StringAttribute{
			MarkdownDescription: "Set the compression mode. One of `gzip`, `none`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_header_name": schema.StringAttribute{
			MarkdownDescription: "Name of an additional HTTP header to send with requests, e.g. Authorization.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"opentelemetry_header_value": schema.StringAttribute{
			MarkdownDescription: "Value for the custom header, e.g. Bearer your-token-here.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Always `telegraf_output`.",
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func convertOutputSettingsSchemaToStruct(d *outputSettingsResourceModel) *telegraf.OutputSettings {
	return &telegraf.OutputSettings{
		InfluxEnable:                    settingsBool(d.InfluxEnable),
		InfluxUrl:                       d.InfluxUrl.ValueString(),
		InfluxDatabase:                  d.InfluxDatabase.ValueString(),
		InfluxTimeout:                   settingsInt(d.InfluxTimeout),
		InfluxNamePrefix:                d.InfluxNamePrefix.ValueString(),
		InfluxUsername:                  d.InfluxUsername.ValueString(),
		InfluxPassword:                  d.InfluxPassword.ValueString(),
		InfluxInsecureSkipVerify:        settingsBool(d.InfluxInsecureSkipVerify),
		GraphiteEnable:                  settingsBool(d.GraphiteEnable),
		GraphiteServer:                  d.GraphiteServer.ValueString(),
		GraphitePrefix:                  d.GraphitePrefix.ValueString(),
		GraphiteTemplate:                d.GraphiteTemplate.ValueString(),
		GraphiteSslDisable:              settingsBool(d.GraphiteSslDisable),
		GraphiteVerify:                  settingsBool(d.GraphiteVerify),
		GraphiteTagsupport:              settingsBool(d.GraphiteTagsupport),
		GraylogEnable:                   settingsBool(d.GraylogEnable),
		GraylogServer:                   d.GraylogServer.ValueString(),
		ElasticEnable:                   settingsBool(d.ElasticEnable),
		ElasticUrl:                      d.ElasticUrl.ValueString(),
		ElasticUsername:                 d.ElasticUsername.ValueString(),
		ElasticPassword:                 d.ElasticPassword.ValueString(),
		ElasticTimeout:                  settingsInt(d.ElasticTimeout),
		ElasticIndexname:                d.ElasticIndexname.ValueString(),
		PrometheusEnable:                settingsBool(d.PrometheusEnable),
		PrometheusListen:                settingsInt(d.PrometheusListen),
		PrometheusExclude:               api.SelectedMapList(settingsList(d.PrometheusExclude)),
		PrometheusStringaslabel:         settingsBool(d.PrometheusStringaslabel),
		InfluxV2Enable:                  settingsBool(d.InfluxV2Enable),
		InfluxV2Url:                     d.InfluxV2Url.ValueString(),
		InfluxV2Token:                   d.InfluxV2Token.ValueString(),
		InfluxV2Organization:            d.InfluxV2Organization.ValueString(),
		InfluxV2Bucket:                  d.InfluxV2Bucket.ValueString(),
		InfluxV2InsecureSkipVerify:      settingsBool(d.InfluxV2InsecureSkipVerify),
		InfluxV2Timeout:                 settingsInt(d.InfluxV2Timeout),
		DatadogEnable:                   settingsBool(d.DatadogEnable),
		DatadogUrl:                      d.DatadogUrl.ValueString(),
		DatadogApikey:                   d.DatadogApikey.ValueString(),
		MqttEnable:                      settingsBool(d.MqttEnable),
		MqttTopicPrefix:                 d.MqttTopicPrefix.ValueString(),
		MqttTopic:                       d.MqttTopic.ValueString(),
		MqttServers:                     api.SelectedMapList(settingsList(d.MqttServers)),
		MqttInsecureSkipVerify:          settingsBool(d.MqttInsecureSkipVerify),
		MqttClientId:                    d.MqttClientId.ValueString(),
		MqttQos:                         api.SelectedMap(d.MqttQos.ValueString()),
		MqttRetain:                      settingsBool(d.MqttRetain),
		MqttTimeout:                     settingsInt(d.MqttTimeout),
		MqttUsername:                    d.MqttUsername.ValueString(),
		MqttPassword:                    d.MqttPassword.ValueString(),
		MqttLayout:                      api.SelectedMap(d.MqttLayout.ValueString()),
		MqttDataFormat:                  api.SelectedMap(d.MqttDataFormat.ValueString()),
		OpentelemetryEnable:             settingsBool(d.OpentelemetryEnable),
		OpentelemetryServer:             d.OpentelemetryServer.ValueString(),
		OpentelemetryInsecureSkipVerify: settingsBool(d.OpentelemetryInsecureSkipVerify),
		OpentelemetryTimeout:            settingsInt(d.OpentelemetryTimeout),
		OpentelemetryCompression:        api.SelectedMap(d.OpentelemetryCompression.ValueString()),
		OpentelemetryHeaderName:         d.OpentelemetryHeaderName.ValueString(),
		OpentelemetryHeaderValue:        d.OpentelemetryHeaderValue.ValueString(),
	}
}

func convertOutputSettingsStructToSchema(s *telegraf.OutputSettings) *outputSettingsResourceModel {
	return &outputSettingsResourceModel{
		InfluxEnable:                    types.BoolValue(tools.StringToBool(s.InfluxEnable)),
		InfluxUrl:                       types.StringValue(s.InfluxUrl),
		InfluxDatabase:                  types.StringValue(s.InfluxDatabase),
		InfluxTimeout:                   settingsReadInt(s.InfluxTimeout),
		InfluxNamePrefix:                types.StringValue(s.InfluxNamePrefix),
		InfluxUsername:                  types.StringValue(s.InfluxUsername),
		InfluxPassword:                  types.StringValue(s.InfluxPassword),
		InfluxInsecureSkipVerify:        types.BoolValue(tools.StringToBool(s.InfluxInsecureSkipVerify)),
		GraphiteEnable:                  types.BoolValue(tools.StringToBool(s.GraphiteEnable)),
		GraphiteServer:                  types.StringValue(s.GraphiteServer),
		GraphitePrefix:                  types.StringValue(s.GraphitePrefix),
		GraphiteTemplate:                types.StringValue(s.GraphiteTemplate),
		GraphiteSslDisable:              types.BoolValue(tools.StringToBool(s.GraphiteSslDisable)),
		GraphiteVerify:                  types.BoolValue(tools.StringToBool(s.GraphiteVerify)),
		GraphiteTagsupport:              types.BoolValue(tools.StringToBool(s.GraphiteTagsupport)),
		GraylogEnable:                   types.BoolValue(tools.StringToBool(s.GraylogEnable)),
		GraylogServer:                   types.StringValue(s.GraylogServer),
		ElasticEnable:                   types.BoolValue(tools.StringToBool(s.ElasticEnable)),
		ElasticUrl:                      types.StringValue(s.ElasticUrl),
		ElasticUsername:                 types.StringValue(s.ElasticUsername),
		ElasticPassword:                 types.StringValue(s.ElasticPassword),
		ElasticTimeout:                  settingsReadInt(s.ElasticTimeout),
		ElasticIndexname:                types.StringValue(s.ElasticIndexname),
		PrometheusEnable:                types.BoolValue(tools.StringToBool(s.PrometheusEnable)),
		PrometheusListen:                settingsReadInt(s.PrometheusListen),
		PrometheusExclude:               tools.StringSliceToSet(settingsNonEmpty(s.PrometheusExclude)),
		PrometheusStringaslabel:         types.BoolValue(tools.StringToBool(s.PrometheusStringaslabel)),
		InfluxV2Enable:                  types.BoolValue(tools.StringToBool(s.InfluxV2Enable)),
		InfluxV2Url:                     types.StringValue(s.InfluxV2Url),
		InfluxV2Token:                   types.StringValue(s.InfluxV2Token),
		InfluxV2Organization:            types.StringValue(s.InfluxV2Organization),
		InfluxV2Bucket:                  types.StringValue(s.InfluxV2Bucket),
		InfluxV2InsecureSkipVerify:      types.BoolValue(tools.StringToBool(s.InfluxV2InsecureSkipVerify)),
		InfluxV2Timeout:                 settingsReadInt(s.InfluxV2Timeout),
		DatadogEnable:                   types.BoolValue(tools.StringToBool(s.DatadogEnable)),
		DatadogUrl:                      types.StringValue(s.DatadogUrl),
		DatadogApikey:                   types.StringValue(s.DatadogApikey),
		MqttEnable:                      types.BoolValue(tools.StringToBool(s.MqttEnable)),
		MqttTopicPrefix:                 types.StringValue(s.MqttTopicPrefix),
		MqttTopic:                       types.StringValue(s.MqttTopic),
		MqttServers:                     tools.StringSliceToSet(settingsNonEmpty(s.MqttServers)),
		MqttInsecureSkipVerify:          types.BoolValue(tools.StringToBool(s.MqttInsecureSkipVerify)),
		MqttClientId:                    types.StringValue(s.MqttClientId),
		MqttQos:                         types.StringValue(s.MqttQos.String()),
		MqttRetain:                      types.BoolValue(tools.StringToBool(s.MqttRetain)),
		MqttTimeout:                     settingsReadInt(s.MqttTimeout),
		MqttUsername:                    types.StringValue(s.MqttUsername),
		MqttPassword:                    types.StringValue(s.MqttPassword),
		MqttLayout:                      types.StringValue(s.MqttLayout.String()),
		MqttDataFormat:                  types.StringValue(s.MqttDataFormat.String()),
		OpentelemetryEnable:             types.BoolValue(tools.StringToBool(s.OpentelemetryEnable)),
		OpentelemetryServer:             types.StringValue(s.OpentelemetryServer),
		OpentelemetryInsecureSkipVerify: types.BoolValue(tools.StringToBool(s.OpentelemetryInsecureSkipVerify)),
		OpentelemetryTimeout:            settingsReadInt(s.OpentelemetryTimeout),
		OpentelemetryCompression:        types.StringValue(s.OpentelemetryCompression.String()),
		OpentelemetryHeaderName:         types.StringValue(s.OpentelemetryHeaderName),
		OpentelemetryHeaderValue:        types.StringValue(s.OpentelemetryHeaderValue),
	}
}

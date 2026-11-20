package routing

import (
	"context"
	"fmt"
	"sort"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/routing"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const gatewayGroupMaxTier = 5

type gatewayGroupResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Members     types.Map    `tfsdk:"members"`
	Trigger     types.String `tfsdk:"trigger"`
	PoolOptions types.String `tfsdk:"pool_options"`
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func gatewayGroupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Gateway groups combine gateways into tiers for failover and load balancing. Policy rules can use a group wherever they accept a gateway. Requires OPNsense 26.7 or later.",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Group name. Only `a-zA-Z0-9_-`, up to 32 characters. Policy rules refer to the group by this name.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(gatewayNameRegexp, "must only contain a-zA-Z0-9_- and be at most 32 characters"),
				},
			},
			"members": schema.MapAttribute{
				MarkdownDescription: "Gateways in the group, keyed by gateway name, with their tier (`1` to `5`) as value. Lower tiers are used first; gateways sharing a tier are load-balanced.",
				Required:            true,
				ElementType:         types.Int64Type,
				Validators: []validator.Map{
					mapvalidator.SizeAtLeast(1),
					mapvalidator.ValueInt64sAre(int64validator.Between(1, gatewayGroupMaxTier)),
				},
			},
			"trigger": schema.StringAttribute{
				MarkdownDescription: "When a member is considered down, one of `down` (member down), `downloss` (packet loss), `downlatency` (high latency) or `downlosslatency` (packet loss or high latency). Defaults to `down`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("down"),
				Validators: []validator.String{
					stringvalidator.OneOf("down", "downloss", "downlatency", "downlosslatency"),
				},
			},
			"pool_options": schema.StringAttribute{
				MarkdownDescription: "Load-balancing behaviour between gateways of the same tier, one of `round-robin` or `round-robin sticky-address`. Leave unset for the OPNsense default.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("round-robin", "round-robin sticky-address"),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the gateway group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func gatewayGroupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gateway groups combine gateways into tiers for failover and load balancing. Requires OPNsense 26.7 or later.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the gateway group.",
				Required:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Group name.",
				Computed:            true,
			},
			"members": dschema.MapAttribute{
				MarkdownDescription: "Gateways in the group, keyed by gateway name, with their tier as value.",
				Computed:            true,
				ElementType:         types.Int64Type,
			},
			"trigger": dschema.StringAttribute{
				MarkdownDescription: "When a member is considered down.",
				Computed:            true,
			},
			"pool_options": dschema.StringAttribute{
				MarkdownDescription: "Load-balancing behaviour between gateways of the same tier.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Optional description here for your reference (not parsed).",
				Computed:            true,
			},
		},
	}
}

func convertGatewayGroupSchemaToStruct(d *gatewayGroupResourceModel) (*routing.GatewayGroup, error) {
	members := map[string]int64{}
	if diags := d.Members.ElementsAs(context.Background(), &members, false); diags.HasError() {
		return nil, fmt.Errorf("unable to read members: %v", diags)
	}

	tiers := make([]api.SelectedMapList, gatewayGroupMaxTier)
	names := make([]string, 0, len(members))
	for name := range members {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		tier := members[name]
		if tier < 1 || tier > gatewayGroupMaxTier {
			return nil, fmt.Errorf("gateway %q has tier %d, must be between 1 and %d", name, tier, gatewayGroupMaxTier)
		}
		tiers[tier-1] = append(tiers[tier-1], name)
	}
	for i := range tiers {
		if tiers[i] == nil {
			tiers[i] = api.SelectedMapList{}
		}
	}

	return &routing.GatewayGroup{
		Name:        d.Name.ValueString(),
		Tier1:       tiers[0],
		Tier2:       tiers[1],
		Tier3:       tiers[2],
		Tier4:       tiers[3],
		Tier5:       tiers[4],
		Trigger:     api.SelectedMap(d.Trigger.ValueString()),
		PoolOptions: api.SelectedMap(d.PoolOptions.ValueString()),
		Description: d.Description.ValueString(),
	}, nil
}

func convertGatewayGroupStructToSchema(d *routing.GatewayGroup) (*gatewayGroupResourceModel, error) {
	members := map[string]int64{}
	for i, tier := range []api.SelectedMapList{d.Tier1, d.Tier2, d.Tier3, d.Tier4, d.Tier5} {
		for _, name := range tier {
			if name == "" {
				continue
			}
			if _, seen := members[name]; !seen {
				members[name] = int64(i + 1)
			}
		}
	}

	membersValue, diags := types.MapValueFrom(context.Background(), types.Int64Type, members)
	if diags.HasError() {
		return nil, fmt.Errorf("unable to build members: %v", diags)
	}

	return &gatewayGroupResourceModel{
		Name:        types.StringValue(d.Name),
		Members:     membersValue,
		Trigger:     types.StringValue(d.Trigger.String()),
		PoolOptions: tools.StringOrNull(d.PoolOptions.String()),
		Description: tools.StringOrNull(d.Description),
	}, nil
}

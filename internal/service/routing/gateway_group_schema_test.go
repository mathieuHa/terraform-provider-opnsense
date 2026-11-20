package routing

import (
	"context"
	"reflect"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/routing"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestConvertGatewayGroupSchemaToStruct_GroupsMembersByTier(t *testing.T) {
	members, _ := types.MapValueFrom(context.Background(), types.Int64Type, map[string]int64{
		"wan_b": 1,
		"wan_a": 1,
		"lte":   3,
	})

	got, err := convertGatewayGroupSchemaToStruct(&gatewayGroupResourceModel{
		Name:    types.StringValue("failover"),
		Members: members,
		Trigger: types.StringValue("down"),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got.Tier1, api.SelectedMapList{"wan_a", "wan_b"}) {
		t.Errorf("Tier1 = %v, want [wan_a wan_b]", got.Tier1)
	}
	if !reflect.DeepEqual(got.Tier3, api.SelectedMapList{"lte"}) {
		t.Errorf("Tier3 = %v, want [lte]", got.Tier3)
	}
	for i, tier := range []api.SelectedMapList{got.Tier2, got.Tier4, got.Tier5} {
		if tier == nil || len(tier) != 0 {
			t.Errorf("empty tier %d = %#v, want empty non-nil list so the field is cleared", i, tier)
		}
	}
}

func TestConvertGatewayGroupStructToSchema_FlattensTiers(t *testing.T) {
	got, err := convertGatewayGroupStructToSchema(&routing.GatewayGroup{
		Name:    "failover",
		Tier1:   api.SelectedMapList{"wan_a"},
		Tier2:   api.SelectedMapList{"wan_b", ""},
		Tier5:   api.SelectedMapList{"lte"},
		Trigger: api.SelectedMap("downloss"),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	members := map[string]int64{}
	got.Members.ElementsAs(context.Background(), &members, false)
	want := map[string]int64{"wan_a": 1, "wan_b": 2, "lte": 5}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("members = %v, want %v", members, want)
	}
	if !got.PoolOptions.IsNull() {
		t.Errorf("PoolOptions = %v, want null for the OPNsense default", got.PoolOptions)
	}
}

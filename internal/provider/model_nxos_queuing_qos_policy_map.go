// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type QueuingQOSPolicyMap struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	MatchType types.String `tfsdk:"match_type"`
}

func (data QueuingQOSPolicyMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/queuing/p/name-[%s]", data.Name.ValueString())
}

func (data QueuingQOSPolicyMap) getClassName() string {
	return "ipqosPMapInst"
}

func (data QueuingQOSPolicyMap) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}
	if (!data.MatchType.IsUnknown() && !data.MatchType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"matchType", data.MatchType.ValueString())
	}

	return nxos.Body{body}
}

func (data *QueuingQOSPolicyMap) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
	if !data.MatchType.IsNull() || all {
		data.MatchType = types.StringValue(res.Get(data.getClassName() + ".attributes.matchType").String())
	} else {
		data.MatchType = types.StringNull()
	}
}

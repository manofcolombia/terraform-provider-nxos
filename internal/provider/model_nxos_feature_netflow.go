// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FeatureNetflow struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	AdminState types.String `tfsdk:"admin_state"`
}

func (data FeatureNetflow) getDn() string {
	return "sys/fm/netflow"
}

func (data FeatureNetflow) getClassName() string {
	return "fmNetflow"
}

func (data FeatureNetflow) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}

	return nxos.Body{body}
}

func (data *FeatureNetflow) fromBody(res gjson.Result, all bool) {
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
}

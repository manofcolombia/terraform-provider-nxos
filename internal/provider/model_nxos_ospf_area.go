// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type OSPFArea struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	InstanceName       types.String `tfsdk:"instance_name"`
	VrfName            types.String `tfsdk:"vrf_name"`
	AreaId             types.String `tfsdk:"area_id"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	Cost               types.Int64  `tfsdk:"cost"`
	Type               types.String `tfsdk:"type"`
}

func (data OSPFArea) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/area-[%s]", data.InstanceName.ValueString(), data.VrfName.ValueString(), data.AreaId.ValueString())
}

func (data OSPFArea) getClassName() string {
	return "ospfArea"
}

func (data OSPFArea) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.AreaId.IsUnknown() && !data.AreaId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.AreaId.ValueString())
	}
	if (!data.AuthenticationType.IsUnknown() && !data.AuthenticationType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authType", data.AuthenticationType.ValueString())
	}
	if (!data.Cost.IsUnknown() && !data.Cost.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cost", strconv.FormatInt(data.Cost.ValueInt64(), 10))
	}
	if (!data.Type.IsUnknown() && !data.Type.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.Type.ValueString())
	}

	return nxos.Body{body}
}

func (data *OSPFArea) fromBody(res gjson.Result, all bool) {
	if !data.AreaId.IsNull() || all {
		data.AreaId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.AreaId = types.StringNull()
	}
	if !data.AuthenticationType.IsNull() || all {
		data.AuthenticationType = types.StringValue(res.Get(data.getClassName() + ".attributes.authType").String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if !data.Cost.IsNull() || all {
		data.Cost = types.Int64Value(res.Get(data.getClassName() + ".attributes.cost").Int())
	} else {
		data.Cost = types.Int64Null()
	}
	if !data.Type.IsNull() || all {
		data.Type = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.Type = types.StringNull()
	}
}

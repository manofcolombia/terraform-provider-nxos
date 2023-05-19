// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LoopbackInterface struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	InterfaceId types.String `tfsdk:"interface_id"`
	AdminState  types.String `tfsdk:"admin_state"`
	Description types.String `tfsdk:"description"`
}

func (data LoopbackInterface) getDn() string {
	return fmt.Sprintf("sys/intf/lb-[%s]", data.InterfaceId.ValueString())
}

func (data LoopbackInterface) getClassName() string {
	return "l3LbRtdIf"
}

func (data LoopbackInterface) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.Description.IsUnknown() && !data.Description.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"descr", data.Description.ValueString())
	}

	return nxos.Body{body}
}

func (data *LoopbackInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
	if !data.Description.IsNull() || all {
		data.Description = types.StringValue(res.Get(data.getClassName() + ".attributes.descr").String())
	} else {
		data.Description = types.StringNull()
	}
}

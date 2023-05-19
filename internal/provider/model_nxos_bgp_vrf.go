// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPVRF struct {
	Device   types.String `tfsdk:"device"`
	Dn       types.String `tfsdk:"id"`
	Asn      types.String `tfsdk:"asn"`
	Name     types.String `tfsdk:"name"`
	RouterId types.String `tfsdk:"router_id"`
}

func (data BGPVRF) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]", data.Name.ValueString())
}

func (data BGPVRF) getClassName() string {
	return "bgpDom"
}

func (data BGPVRF) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}
	if (!data.RouterId.IsUnknown() && !data.RouterId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rtrId", data.RouterId.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPVRF) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
	if !data.RouterId.IsNull() || all {
		data.RouterId = types.StringValue(res.Get(data.getClassName() + ".attributes.rtrId").String())
	} else {
		data.RouterId = types.StringNull()
	}
}

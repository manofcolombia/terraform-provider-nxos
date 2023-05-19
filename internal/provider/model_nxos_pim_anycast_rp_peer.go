// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PIMAnycastRPPeer struct {
	Device       types.String `tfsdk:"device"`
	Dn           types.String `tfsdk:"id"`
	VrfName      types.String `tfsdk:"vrf_name"`
	Address      types.String `tfsdk:"address"`
	RpSetAddress types.String `tfsdk:"rp_set_address"`
}

func (data PIMAnycastRPPeer) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/acastrpfunc/peer-[%s]-peer-[%s]", data.VrfName.ValueString(), data.Address.ValueString(), data.RpSetAddress.ValueString())
}

func (data PIMAnycastRPPeer) getClassName() string {
	return "pimAcastRPPeer"
}

func (data PIMAnycastRPPeer) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"addr", data.Address.ValueString())
	}
	if (!data.RpSetAddress.IsUnknown() && !data.RpSetAddress.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rpSetAddr", data.RpSetAddress.ValueString())
	}

	return nxos.Body{body}
}

func (data *PIMAnycastRPPeer) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.addr").String())
	} else {
		data.Address = types.StringNull()
	}
	if !data.RpSetAddress.IsNull() || all {
		data.RpSetAddress = types.StringValue(res.Get(data.getClassName() + ".attributes.rpSetAddr").String())
	} else {
		data.RpSetAddress = types.StringNull()
	}
}

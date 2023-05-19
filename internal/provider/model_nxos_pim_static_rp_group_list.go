// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PIMStaticRPGroupList struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	VrfName   types.String `tfsdk:"vrf_name"`
	RpAddress types.String `tfsdk:"rp_address"`
	Address   types.String `tfsdk:"address"`
	Bidir     types.Bool   `tfsdk:"bidir"`
	Override  types.Bool   `tfsdk:"override"`
}

func (data PIMStaticRPGroupList) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/staticrp/rp-[%s]/rpgrplist-[%s]", data.VrfName.ValueString(), data.RpAddress.ValueString(), data.Address.ValueString())
}

func (data PIMStaticRPGroupList) getClassName() string {
	return "pimRPGrpList"
}

func (data PIMStaticRPGroupList) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"grpListName", data.Address.ValueString())
	}
	if (!data.Bidir.IsUnknown() && !data.Bidir.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bidir", strconv.FormatBool(data.Bidir.ValueBool()))
	}
	if (!data.Override.IsUnknown() && !data.Override.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"override", strconv.FormatBool(data.Override.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *PIMStaticRPGroupList) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.grpListName").String())
	} else {
		data.Address = types.StringNull()
	}
	if !data.Bidir.IsNull() || all {
		data.Bidir = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.bidir").String()))
	} else {
		data.Bidir = types.BoolNull()
	}
	if !data.Override.IsNull() || all {
		data.Override = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.override").String()))
	} else {
		data.Override = types.BoolNull()
	}
}

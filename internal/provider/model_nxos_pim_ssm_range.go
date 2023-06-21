// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PIMSSMRange struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	VrfName    types.String `tfsdk:"vrf_name"`
	GroupList1 types.String `tfsdk:"group_list_1"`
	GroupList2 types.String `tfsdk:"group_list_2"`
	GroupList3 types.String `tfsdk:"group_list_3"`
	GroupList4 types.String `tfsdk:"group_list_4"`
	PrefixList types.String `tfsdk:"prefix_list"`
	RouteMap   types.String `tfsdk:"route_map"`
	SsmNone    types.Bool   `tfsdk:"ssm_none"`
}

func (data PIMSSMRange) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/ssm/range", data.VrfName.ValueString())
}

func (data PIMSSMRange) getClassName() string {
	return "pimSSMRangeP"
}

func (data PIMSSMRange) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.GroupList1.IsUnknown() && !data.GroupList1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"grpList", data.GroupList1.ValueString())
	}
	if (!data.GroupList2.IsUnknown() && !data.GroupList2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"grpList1", data.GroupList2.ValueString())
	}
	if (!data.GroupList3.IsUnknown() && !data.GroupList3.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"grpList2", data.GroupList3.ValueString())
	}
	if (!data.GroupList4.IsUnknown() && !data.GroupList4.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"grpList3", data.GroupList4.ValueString())
	}
	if (!data.PrefixList.IsUnknown() && !data.PrefixList.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pfxList", data.PrefixList.ValueString())
	}
	if (!data.RouteMap.IsUnknown() && !data.RouteMap.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rtMap", data.RouteMap.ValueString())
	}
	if (!data.SsmNone.IsUnknown() && !data.SsmNone.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ssmNone", strconv.FormatBool(data.SsmNone.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *PIMSSMRange) fromBody(res gjson.Result, all bool) {
	if !data.GroupList1.IsNull() || all {
		data.GroupList1 = types.StringValue(res.Get(data.getClassName() + ".attributes.grpList").String())
	} else {
		data.GroupList1 = types.StringNull()
	}
	if !data.GroupList2.IsNull() || all {
		data.GroupList2 = types.StringValue(res.Get(data.getClassName() + ".attributes.grpList1").String())
	} else {
		data.GroupList2 = types.StringNull()
	}
	if !data.GroupList3.IsNull() || all {
		data.GroupList3 = types.StringValue(res.Get(data.getClassName() + ".attributes.grpList2").String())
	} else {
		data.GroupList3 = types.StringNull()
	}
	if !data.GroupList4.IsNull() || all {
		data.GroupList4 = types.StringValue(res.Get(data.getClassName() + ".attributes.grpList3").String())
	} else {
		data.GroupList4 = types.StringNull()
	}
	if !data.PrefixList.IsNull() || all {
		data.PrefixList = types.StringValue(res.Get(data.getClassName() + ".attributes.pfxList").String())
	} else {
		data.PrefixList = types.StringNull()
	}
	if !data.RouteMap.IsNull() || all {
		data.RouteMap = types.StringValue(res.Get(data.getClassName() + ".attributes.rtMap").String())
	} else {
		data.RouteMap = types.StringNull()
	}
	if !data.SsmNone.IsNull() || all {
		data.SsmNone = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.ssmNone").String()))
	} else {
		data.SsmNone = types.BoolNull()
	}
}

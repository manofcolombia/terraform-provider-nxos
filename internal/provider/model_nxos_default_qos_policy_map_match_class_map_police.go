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

type DefaultQOSPolicyMapMatchClassMapPolice struct {
	Device               types.String `tfsdk:"device"`
	Dn                   types.String `tfsdk:"id"`
	PolicyMapName        types.String `tfsdk:"policy_map_name"`
	ClassMapName         types.String `tfsdk:"class_map_name"`
	BcRate               types.Int64  `tfsdk:"bc_rate"`
	BcUnit               types.String `tfsdk:"bc_unit"`
	BeRate               types.Int64  `tfsdk:"be_rate"`
	BeUnit               types.String `tfsdk:"be_unit"`
	CirRate              types.Int64  `tfsdk:"cir_rate"`
	CirUnit              types.String `tfsdk:"cir_unit"`
	ConformAction        types.String `tfsdk:"conform_action"`
	ConformSetCos        types.Int64  `tfsdk:"conform_set_cos"`
	ConformSetDscp       types.Int64  `tfsdk:"conform_set_dscp"`
	ConformSetPrecedence types.String `tfsdk:"conform_set_precedence"`
	ConformSetQosGroup   types.Int64  `tfsdk:"conform_set_qos_group"`
	ExceedAction         types.String `tfsdk:"exceed_action"`
	ExceedSetCos         types.Int64  `tfsdk:"exceed_set_cos"`
	ExceedSetDscp        types.Int64  `tfsdk:"exceed_set_dscp"`
	ExceedSetPrecedence  types.String `tfsdk:"exceed_set_precedence"`
	ExceedSetQosGroup    types.Int64  `tfsdk:"exceed_set_qos_group"`
	PirRate              types.Int64  `tfsdk:"pir_rate"`
	PirUnit              types.String `tfsdk:"pir_unit"`
	ViolateAction        types.String `tfsdk:"violate_action"`
	ViolateSetCos        types.Int64  `tfsdk:"violate_set_cos"`
	ViolateSetDscp       types.Int64  `tfsdk:"violate_set_dscp"`
	ViolateSetPrecedence types.String `tfsdk:"violate_set_precedence"`
	ViolateSetQosGroup   types.Int64  `tfsdk:"violate_set_qos_group"`
}

func (data DefaultQOSPolicyMapMatchClassMapPolice) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/p/name-[%s]/cmap-[%s]/police", data.PolicyMapName.ValueString(), data.ClassMapName.ValueString())
}

func (data DefaultQOSPolicyMapMatchClassMapPolice) getClassName() string {
	return "ipqosPolice"
}

func (data DefaultQOSPolicyMapMatchClassMapPolice) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.BcRate.IsUnknown() && !data.BcRate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bcRate", strconv.FormatInt(data.BcRate.ValueInt64(), 10))
	}
	if (!data.BcUnit.IsUnknown() && !data.BcUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bcUnit", data.BcUnit.ValueString())
	}
	if (!data.BeRate.IsUnknown() && !data.BeRate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"beRate", strconv.FormatInt(data.BeRate.ValueInt64(), 10))
	}
	if (!data.BeUnit.IsUnknown() && !data.BeUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"beUnit", data.BeUnit.ValueString())
	}
	if (!data.CirRate.IsUnknown() && !data.CirRate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cirRate", strconv.FormatInt(data.CirRate.ValueInt64(), 10))
	}
	if (!data.CirUnit.IsUnknown() && !data.CirUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cirUnit", data.CirUnit.ValueString())
	}
	if (!data.ConformAction.IsUnknown() && !data.ConformAction.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"conformAction", data.ConformAction.ValueString())
	}
	if (!data.ConformSetCos.IsUnknown() && !data.ConformSetCos.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"conformSetCosTransmit", strconv.FormatInt(data.ConformSetCos.ValueInt64(), 10))
	}
	if (!data.ConformSetDscp.IsUnknown() && !data.ConformSetDscp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"conformSetDscpTransmit", strconv.FormatInt(data.ConformSetDscp.ValueInt64(), 10))
	}
	if (!data.ConformSetPrecedence.IsUnknown() && !data.ConformSetPrecedence.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"conformSetPrecTransmit", data.ConformSetPrecedence.ValueString())
	}
	if (!data.ConformSetQosGroup.IsUnknown() && !data.ConformSetQosGroup.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"conformSetQosGrpTransmit", strconv.FormatInt(data.ConformSetQosGroup.ValueInt64(), 10))
	}
	if (!data.ExceedAction.IsUnknown() && !data.ExceedAction.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"exceedAction", data.ExceedAction.ValueString())
	}
	if (!data.ExceedSetCos.IsUnknown() && !data.ExceedSetCos.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"exceedSetCosTransmit", strconv.FormatInt(data.ExceedSetCos.ValueInt64(), 10))
	}
	if (!data.ExceedSetDscp.IsUnknown() && !data.ExceedSetDscp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"exceedSetDscpTransmit", strconv.FormatInt(data.ExceedSetDscp.ValueInt64(), 10))
	}
	if (!data.ExceedSetPrecedence.IsUnknown() && !data.ExceedSetPrecedence.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"exceedSetPrecTransmit", data.ExceedSetPrecedence.ValueString())
	}
	if (!data.ExceedSetQosGroup.IsUnknown() && !data.ExceedSetQosGroup.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"exceedSetQosGrpTransmit", strconv.FormatInt(data.ExceedSetQosGroup.ValueInt64(), 10))
	}
	if (!data.PirRate.IsUnknown() && !data.PirRate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pirRate", strconv.FormatInt(data.PirRate.ValueInt64(), 10))
	}
	if (!data.PirUnit.IsUnknown() && !data.PirUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pirUnit", data.PirUnit.ValueString())
	}
	if (!data.ViolateAction.IsUnknown() && !data.ViolateAction.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"violateAction", data.ViolateAction.ValueString())
	}
	if (!data.ViolateSetCos.IsUnknown() && !data.ViolateSetCos.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"violateSetCosTransmit", strconv.FormatInt(data.ViolateSetCos.ValueInt64(), 10))
	}
	if (!data.ViolateSetDscp.IsUnknown() && !data.ViolateSetDscp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"violateSetDscpTransmit", strconv.FormatInt(data.ViolateSetDscp.ValueInt64(), 10))
	}
	if (!data.ViolateSetPrecedence.IsUnknown() && !data.ViolateSetPrecedence.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"violateSetPrecTransmit", data.ViolateSetPrecedence.ValueString())
	}
	if (!data.ViolateSetQosGroup.IsUnknown() && !data.ViolateSetQosGroup.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"violateSetQosGrpTransmit", strconv.FormatInt(data.ViolateSetQosGroup.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *DefaultQOSPolicyMapMatchClassMapPolice) fromBody(res gjson.Result, all bool) {
	if !data.BcRate.IsNull() || all {
		data.BcRate = types.Int64Value(res.Get(data.getClassName() + ".attributes.bcRate").Int())
	} else {
		data.BcRate = types.Int64Null()
	}
	if !data.BcUnit.IsNull() || all {
		data.BcUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.bcUnit").String())
	} else {
		data.BcUnit = types.StringNull()
	}
	if !data.BeRate.IsNull() || all {
		data.BeRate = types.Int64Value(res.Get(data.getClassName() + ".attributes.beRate").Int())
	} else {
		data.BeRate = types.Int64Null()
	}
	if !data.BeUnit.IsNull() || all {
		data.BeUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.beUnit").String())
	} else {
		data.BeUnit = types.StringNull()
	}
	if !data.CirRate.IsNull() || all {
		data.CirRate = types.Int64Value(res.Get(data.getClassName() + ".attributes.cirRate").Int())
	} else {
		data.CirRate = types.Int64Null()
	}
	if !data.CirUnit.IsNull() || all {
		data.CirUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.cirUnit").String())
	} else {
		data.CirUnit = types.StringNull()
	}
	if !data.ConformAction.IsNull() || all {
		data.ConformAction = types.StringValue(res.Get(data.getClassName() + ".attributes.conformAction").String())
	} else {
		data.ConformAction = types.StringNull()
	}
	if !data.ConformSetCos.IsNull() || all {
		data.ConformSetCos = types.Int64Value(res.Get(data.getClassName() + ".attributes.conformSetCosTransmit").Int())
	} else {
		data.ConformSetCos = types.Int64Null()
	}
	if !data.ConformSetDscp.IsNull() || all {
		data.ConformSetDscp = types.Int64Value(res.Get(data.getClassName() + ".attributes.conformSetDscpTransmit").Int())
	} else {
		data.ConformSetDscp = types.Int64Null()
	}
	if !data.ConformSetPrecedence.IsNull() || all {
		data.ConformSetPrecedence = types.StringValue(res.Get(data.getClassName() + ".attributes.conformSetPrecTransmit").String())
	} else {
		data.ConformSetPrecedence = types.StringNull()
	}
	if !data.ConformSetQosGroup.IsNull() || all {
		data.ConformSetQosGroup = types.Int64Value(res.Get(data.getClassName() + ".attributes.conformSetQosGrpTransmit").Int())
	} else {
		data.ConformSetQosGroup = types.Int64Null()
	}
	if !data.ExceedAction.IsNull() || all {
		data.ExceedAction = types.StringValue(res.Get(data.getClassName() + ".attributes.exceedAction").String())
	} else {
		data.ExceedAction = types.StringNull()
	}
	if !data.ExceedSetCos.IsNull() || all {
		data.ExceedSetCos = types.Int64Value(res.Get(data.getClassName() + ".attributes.exceedSetCosTransmit").Int())
	} else {
		data.ExceedSetCos = types.Int64Null()
	}
	if !data.ExceedSetDscp.IsNull() || all {
		data.ExceedSetDscp = types.Int64Value(res.Get(data.getClassName() + ".attributes.exceedSetDscpTransmit").Int())
	} else {
		data.ExceedSetDscp = types.Int64Null()
	}
	if !data.ExceedSetPrecedence.IsNull() || all {
		data.ExceedSetPrecedence = types.StringValue(res.Get(data.getClassName() + ".attributes.exceedSetPrecTransmit").String())
	} else {
		data.ExceedSetPrecedence = types.StringNull()
	}
	if !data.ExceedSetQosGroup.IsNull() || all {
		data.ExceedSetQosGroup = types.Int64Value(res.Get(data.getClassName() + ".attributes.exceedSetQosGrpTransmit").Int())
	} else {
		data.ExceedSetQosGroup = types.Int64Null()
	}
	if !data.PirRate.IsNull() || all {
		data.PirRate = types.Int64Value(res.Get(data.getClassName() + ".attributes.pirRate").Int())
	} else {
		data.PirRate = types.Int64Null()
	}
	if !data.PirUnit.IsNull() || all {
		data.PirUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.pirUnit").String())
	} else {
		data.PirUnit = types.StringNull()
	}
	if !data.ViolateAction.IsNull() || all {
		data.ViolateAction = types.StringValue(res.Get(data.getClassName() + ".attributes.violateAction").String())
	} else {
		data.ViolateAction = types.StringNull()
	}
	if !data.ViolateSetCos.IsNull() || all {
		data.ViolateSetCos = types.Int64Value(res.Get(data.getClassName() + ".attributes.violateSetCosTransmit").Int())
	} else {
		data.ViolateSetCos = types.Int64Null()
	}
	if !data.ViolateSetDscp.IsNull() || all {
		data.ViolateSetDscp = types.Int64Value(res.Get(data.getClassName() + ".attributes.violateSetDscpTransmit").Int())
	} else {
		data.ViolateSetDscp = types.Int64Null()
	}
	if !data.ViolateSetPrecedence.IsNull() || all {
		data.ViolateSetPrecedence = types.StringValue(res.Get(data.getClassName() + ".attributes.violateSetPrecTransmit").String())
	} else {
		data.ViolateSetPrecedence = types.StringNull()
	}
	if !data.ViolateSetQosGroup.IsNull() || all {
		data.ViolateSetQosGroup = types.Int64Value(res.Get(data.getClassName() + ".attributes.violateSetQosGrpTransmit").Int())
	} else {
		data.ViolateSetQosGroup = types.Int64Null()
	}
}

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

type OSPFAuthentication struct {
	Device           types.String `tfsdk:"device"`
	Dn               types.String `tfsdk:"id"`
	InstanceName     types.String `tfsdk:"instance_name"`
	VrfName          types.String `tfsdk:"vrf_name"`
	InterfaceId      types.String `tfsdk:"interface_id"`
	Key              types.String `tfsdk:"key"`
	KeyId            types.Int64  `tfsdk:"key_id"`
	KeySecureMode    types.Bool   `tfsdk:"key_secure_mode"`
	Keychain         types.String `tfsdk:"keychain"`
	Md5Key           types.String `tfsdk:"md5_key"`
	Md5KeySecureMode types.Bool   `tfsdk:"md5_key_secure_mode"`
	Type             types.String `tfsdk:"type"`
}

func (data OSPFAuthentication) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/if-[%s]/authnew", data.InstanceName.ValueString(), data.VrfName.ValueString(), data.InterfaceId.ValueString())
}

func (data OSPFAuthentication) getClassName() string {
	return "ospfAuthNewP"
}

func (data OSPFAuthentication) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Key.IsUnknown() && !data.Key.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"key", data.Key.ValueString())
	}
	if (!data.KeyId.IsUnknown() && !data.KeyId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"keyId", strconv.FormatInt(data.KeyId.ValueInt64(), 10))
	}
	if (!data.KeySecureMode.IsUnknown() && !data.KeySecureMode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"keySecureMode", strconv.FormatBool(data.KeySecureMode.ValueBool()))
	}
	if (!data.Keychain.IsUnknown() && !data.Keychain.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"keychain", data.Keychain.ValueString())
	}
	if (!data.Md5Key.IsUnknown() && !data.Md5Key.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"md5key", data.Md5Key.ValueString())
	}
	if (!data.Md5KeySecureMode.IsUnknown() && !data.Md5KeySecureMode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"md5keySecureMode", strconv.FormatBool(data.Md5KeySecureMode.ValueBool()))
	}
	if (!data.Type.IsUnknown() && !data.Type.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.Type.ValueString())
	}

	return nxos.Body{body}
}

func (data *OSPFAuthentication) fromBody(res gjson.Result, all bool) {
	if !data.KeyId.IsNull() || all {
		data.KeyId = types.Int64Value(res.Get(data.getClassName() + ".attributes.keyId").Int())
	} else {
		data.KeyId = types.Int64Null()
	}
	if !data.KeySecureMode.IsNull() || all {
		data.KeySecureMode = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.keySecureMode").String()))
	} else {
		data.KeySecureMode = types.BoolNull()
	}
	if !data.Keychain.IsNull() || all {
		data.Keychain = types.StringValue(res.Get(data.getClassName() + ".attributes.keychain").String())
	} else {
		data.Keychain = types.StringNull()
	}
	if !data.Md5KeySecureMode.IsNull() || all {
		data.Md5KeySecureMode = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.md5keySecureMode").String()))
	} else {
		data.Md5KeySecureMode = types.BoolNull()
	}
	if !data.Type.IsNull() || all {
		data.Type = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.Type = types.StringNull()
	}
}

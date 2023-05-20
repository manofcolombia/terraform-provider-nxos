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

type PortChannelInterface struct {
	Device              types.String `tfsdk:"device"`
	Dn                  types.String `tfsdk:"id"`
	InterfaceId         types.String `tfsdk:"interface_id"`
	PortChannelMode     types.String `tfsdk:"port_channel_mode"`
	MinimumLinks        types.Int64  `tfsdk:"minimum_links"`
	MaximumLinks        types.Int64  `tfsdk:"maximum_links"`
	SuspendIndividual   types.String `tfsdk:"suspend_individual"`
	AccessVlan          types.String `tfsdk:"access_vlan"`
	AdminState          types.String `tfsdk:"admin_state"`
	AutoNegotiation     types.String `tfsdk:"auto_negotiation"`
	Bandwidth           types.Int64  `tfsdk:"bandwidth"`
	Delay               types.Int64  `tfsdk:"delay"`
	Description         types.String `tfsdk:"description"`
	Duplex              types.String `tfsdk:"duplex"`
	Layer               types.String `tfsdk:"layer"`
	LinkLogging         types.String `tfsdk:"link_logging"`
	Medium              types.String `tfsdk:"medium"`
	Mode                types.String `tfsdk:"mode"`
	Mtu                 types.Int64  `tfsdk:"mtu"`
	NativeVlan          types.String `tfsdk:"native_vlan"`
	Speed               types.String `tfsdk:"speed"`
	TrunkVlans          types.String `tfsdk:"trunk_vlans"`
	UserConfiguredFlags types.String `tfsdk:"user_configured_flags"`
}

func (data PortChannelInterface) getDn() string {
	return fmt.Sprintf("sys/intf/aggr-[%s]", data.InterfaceId.ValueString())
}

func (data PortChannelInterface) getClassName() string {
	return "pcAggrIf"
}

func (data PortChannelInterface) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.PortChannelMode.IsUnknown() && !data.PortChannelMode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pcMode", data.PortChannelMode.ValueString())
	}
	if (!data.MinimumLinks.IsUnknown() && !data.MinimumLinks.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"minLinks", strconv.FormatInt(data.MinimumLinks.ValueInt64(), 10))
	}
	if (!data.MaximumLinks.IsUnknown() && !data.MaximumLinks.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxLinks", strconv.FormatInt(data.MaximumLinks.ValueInt64(), 10))
	}
	if (!data.SuspendIndividual.IsUnknown() && !data.SuspendIndividual.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"suspIndividual", data.SuspendIndividual.ValueString())
	}
	if (!data.AccessVlan.IsUnknown() && !data.AccessVlan.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"accessVlan", data.AccessVlan.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.AutoNegotiation.IsUnknown() && !data.AutoNegotiation.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"autoNeg", data.AutoNegotiation.ValueString())
	}
	if (!data.Bandwidth.IsUnknown() && !data.Bandwidth.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bw", strconv.FormatInt(data.Bandwidth.ValueInt64(), 10))
	}
	if (!data.Delay.IsUnknown() && !data.Delay.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"delay", strconv.FormatInt(data.Delay.ValueInt64(), 10))
	}
	if (!data.Description.IsUnknown() && !data.Description.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"descr", data.Description.ValueString())
	}
	if (!data.Duplex.IsUnknown() && !data.Duplex.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"duplex", data.Duplex.ValueString())
	}
	if (!data.Layer.IsUnknown() && !data.Layer.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"layer", data.Layer.ValueString())
	}
	if (!data.LinkLogging.IsUnknown() && !data.LinkLogging.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkLog", data.LinkLogging.ValueString())
	}
	if (!data.Medium.IsUnknown() && !data.Medium.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"medium", data.Medium.ValueString())
	}
	if (!data.Mode.IsUnknown() && !data.Mode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mode", data.Mode.ValueString())
	}
	if (!data.Mtu.IsUnknown() && !data.Mtu.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if (!data.NativeVlan.IsUnknown() && !data.NativeVlan.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nativeVlan", data.NativeVlan.ValueString())
	}
	if (!data.Speed.IsUnknown() && !data.Speed.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"speed", data.Speed.ValueString())
	}
	if (!data.TrunkVlans.IsUnknown() && !data.TrunkVlans.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"trunkVlans", data.TrunkVlans.ValueString())
	}
	if (!data.UserConfiguredFlags.IsUnknown() && !data.UserConfiguredFlags.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"userCfgdFlags", data.UserConfiguredFlags.ValueString())
	}

	return nxos.Body{body}
}

func (data *PortChannelInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.PortChannelMode.IsNull() || all {
		data.PortChannelMode = types.StringValue(res.Get(data.getClassName() + ".attributes.pcMode").String())
	} else {
		data.PortChannelMode = types.StringNull()
	}
	if !data.MinimumLinks.IsNull() || all {
		data.MinimumLinks = types.Int64Value(res.Get(data.getClassName() + ".attributes.minLinks").Int())
	} else {
		data.MinimumLinks = types.Int64Null()
	}
	if !data.MaximumLinks.IsNull() || all {
		data.MaximumLinks = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxLinks").Int())
	} else {
		data.MaximumLinks = types.Int64Null()
	}
	if !data.SuspendIndividual.IsNull() || all {
		data.SuspendIndividual = types.StringValue(res.Get(data.getClassName() + ".attributes.suspIndividual").String())
	} else {
		data.SuspendIndividual = types.StringNull()
	}
	if !data.AccessVlan.IsNull() || all {
		data.AccessVlan = types.StringValue(res.Get(data.getClassName() + ".attributes.accessVlan").String())
	} else {
		data.AccessVlan = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
	if !data.AutoNegotiation.IsNull() || all {
		data.AutoNegotiation = types.StringValue(res.Get(data.getClassName() + ".attributes.autoNeg").String())
	} else {
		data.AutoNegotiation = types.StringNull()
	}
	if !data.Bandwidth.IsNull() || all {
		data.Bandwidth = types.Int64Value(res.Get(data.getClassName() + ".attributes.bw").Int())
	} else {
		data.Bandwidth = types.Int64Null()
	}
	if !data.Delay.IsNull() || all {
		data.Delay = types.Int64Value(res.Get(data.getClassName() + ".attributes.delay").Int())
	} else {
		data.Delay = types.Int64Null()
	}
	if !data.Description.IsNull() || all {
		data.Description = types.StringValue(res.Get(data.getClassName() + ".attributes.descr").String())
	} else {
		data.Description = types.StringNull()
	}
	if !data.Duplex.IsNull() || all {
		data.Duplex = types.StringValue(res.Get(data.getClassName() + ".attributes.duplex").String())
	} else {
		data.Duplex = types.StringNull()
	}
	if !data.Layer.IsNull() || all {
		data.Layer = types.StringValue(res.Get(data.getClassName() + ".attributes.layer").String())
	} else {
		data.Layer = types.StringNull()
	}
	if !data.LinkLogging.IsNull() || all {
		data.LinkLogging = types.StringValue(res.Get(data.getClassName() + ".attributes.linkLog").String())
	} else {
		data.LinkLogging = types.StringNull()
	}
	if !data.Medium.IsNull() || all {
		data.Medium = types.StringValue(res.Get(data.getClassName() + ".attributes.medium").String())
	} else {
		data.Medium = types.StringNull()
	}
	if !data.Mode.IsNull() || all {
		data.Mode = types.StringValue(res.Get(data.getClassName() + ".attributes.mode").String())
	} else {
		data.Mode = types.StringNull()
	}
	if !data.Mtu.IsNull() || all {
		data.Mtu = types.Int64Value(res.Get(data.getClassName() + ".attributes.mtu").Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if !data.NativeVlan.IsNull() || all {
		data.NativeVlan = types.StringValue(res.Get(data.getClassName() + ".attributes.nativeVlan").String())
	} else {
		data.NativeVlan = types.StringNull()
	}
	if !data.Speed.IsNull() || all {
		data.Speed = types.StringValue(res.Get(data.getClassName() + ".attributes.speed").String())
	} else {
		data.Speed = types.StringNull()
	}
	if !data.TrunkVlans.IsNull() || all {
		data.TrunkVlans = types.StringValue(res.Get(data.getClassName() + ".attributes.trunkVlans").String())
	} else {
		data.TrunkVlans = types.StringNull()
	}
	if !data.UserConfiguredFlags.IsNull() || all {
		data.UserConfiguredFlags = types.StringValue(res.Get(data.getClassName() + ".attributes.userCfgdFlags").String())
	} else {
		data.UserConfiguredFlags = types.StringNull()
	}
}

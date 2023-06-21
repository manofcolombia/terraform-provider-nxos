// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type NVEInterface struct {
	Device                        types.String `tfsdk:"device"`
	Dn                            types.String `tfsdk:"id"`
	AdminState                    types.String `tfsdk:"admin_state"`
	AdvertiseVirtualMac           types.Bool   `tfsdk:"advertise_virtual_mac"`
	HoldDownTime                  types.Int64  `tfsdk:"hold_down_time"`
	HostReachabilityProtocol      types.String `tfsdk:"host_reachability_protocol"`
	IngressReplicationProtocolBgp types.Bool   `tfsdk:"ingress_replication_protocol_bgp"`
	MulticastGroupL2              types.String `tfsdk:"multicast_group_l2"`
	MulticastGroupL3              types.String `tfsdk:"multicast_group_l3"`
	MultisiteSourceInterface      types.String `tfsdk:"multisite_source_interface"`
	SourceInterface               types.String `tfsdk:"source_interface"`
	SuppressArp                   types.Bool   `tfsdk:"suppress_arp"`
	SuppressMacRoute              types.Bool   `tfsdk:"suppress_mac_route"`
}

func (data NVEInterface) getDn() string {
	return "sys/eps/epId-[1]"
}

func (data NVEInterface) getClassName() string {
	return "nvoEp"
}

func (data NVEInterface) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.AdvertiseVirtualMac.IsUnknown() && !data.AdvertiseVirtualMac.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advertiseVmac", strconv.FormatBool(data.AdvertiseVirtualMac.ValueBool()))
	}
	if (!data.HoldDownTime.IsUnknown() && !data.HoldDownTime.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"holdDownTime", strconv.FormatInt(data.HoldDownTime.ValueInt64(), 10))
	}
	if (!data.HostReachabilityProtocol.IsUnknown() && !data.HostReachabilityProtocol.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"hostReach", data.HostReachabilityProtocol.ValueString())
	}
	if (!data.IngressReplicationProtocolBgp.IsUnknown() && !data.IngressReplicationProtocolBgp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ingressReplProtoBGP", strconv.FormatBool(data.IngressReplicationProtocolBgp.ValueBool()))
	}
	if (!data.MulticastGroupL2.IsUnknown() && !data.MulticastGroupL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mcastGroupL2", data.MulticastGroupL2.ValueString())
	}
	if (!data.MulticastGroupL3.IsUnknown() && !data.MulticastGroupL3.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mcastGroupL3", data.MulticastGroupL3.ValueString())
	}
	if (!data.MultisiteSourceInterface.IsUnknown() && !data.MultisiteSourceInterface.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"multisiteBordergwInterface", data.MultisiteSourceInterface.ValueString())
	}
	if (!data.SourceInterface.IsUnknown() && !data.SourceInterface.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"sourceInterface", data.SourceInterface.ValueString())
	}
	if (!data.SuppressArp.IsUnknown() && !data.SuppressArp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"suppressARP", strconv.FormatBool(data.SuppressArp.ValueBool()))
	}
	if (!data.SuppressMacRoute.IsUnknown() && !data.SuppressMacRoute.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"suppressMacRoute", strconv.FormatBool(data.SuppressMacRoute.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *NVEInterface) fromBody(res gjson.Result, all bool) {
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
	if !data.AdvertiseVirtualMac.IsNull() || all {
		data.AdvertiseVirtualMac = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.advertiseVmac").String()))
	} else {
		data.AdvertiseVirtualMac = types.BoolNull()
	}
	if !data.HoldDownTime.IsNull() || all {
		data.HoldDownTime = types.Int64Value(res.Get(data.getClassName() + ".attributes.holdDownTime").Int())
	} else {
		data.HoldDownTime = types.Int64Null()
	}
	if !data.HostReachabilityProtocol.IsNull() || all {
		data.HostReachabilityProtocol = types.StringValue(res.Get(data.getClassName() + ".attributes.hostReach").String())
	} else {
		data.HostReachabilityProtocol = types.StringNull()
	}
	if !data.IngressReplicationProtocolBgp.IsNull() || all {
		data.IngressReplicationProtocolBgp = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.ingressReplProtoBGP").String()))
	} else {
		data.IngressReplicationProtocolBgp = types.BoolNull()
	}
	if !data.MulticastGroupL2.IsNull() || all {
		data.MulticastGroupL2 = types.StringValue(res.Get(data.getClassName() + ".attributes.mcastGroupL2").String())
	} else {
		data.MulticastGroupL2 = types.StringNull()
	}
	if !data.MulticastGroupL3.IsNull() || all {
		data.MulticastGroupL3 = types.StringValue(res.Get(data.getClassName() + ".attributes.mcastGroupL3").String())
	} else {
		data.MulticastGroupL3 = types.StringNull()
	}
	if !data.MultisiteSourceInterface.IsNull() || all {
		data.MultisiteSourceInterface = types.StringValue(res.Get(data.getClassName() + ".attributes.multisiteBordergwInterface").String())
	} else {
		data.MultisiteSourceInterface = types.StringNull()
	}
	if !data.SourceInterface.IsNull() || all {
		data.SourceInterface = types.StringValue(res.Get(data.getClassName() + ".attributes.sourceInterface").String())
	} else {
		data.SourceInterface = types.StringNull()
	}
	if !data.SuppressArp.IsNull() || all {
		data.SuppressArp = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.suppressARP").String()))
	} else {
		data.SuppressArp = types.BoolNull()
	}
	if !data.SuppressMacRoute.IsNull() || all {
		data.SuppressMacRoute = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.suppressMacRoute").String()))
	} else {
		data.SuppressMacRoute = types.BoolNull()
	}
}

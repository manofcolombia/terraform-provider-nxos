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

type BGPPeer struct {
	Device          types.String `tfsdk:"device"`
	Dn              types.String `tfsdk:"id"`
	Asn             types.String `tfsdk:"asn"`
	Vrf             types.String `tfsdk:"vrf"`
	Address         types.String `tfsdk:"address"`
	RemoteAsn       types.String `tfsdk:"remote_asn"`
	Description     types.String `tfsdk:"description"`
	PeerTemplate    types.String `tfsdk:"peer_template"`
	PeerType        types.String `tfsdk:"peer_type"`
	SourceInterface types.String `tfsdk:"source_interface"`
	HoldTime        types.Int64  `tfsdk:"hold_time"`
	Keepalive       types.Int64  `tfsdk:"keepalive"`
}

func (data BGPPeer) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]", data.Vrf.ValueString(), data.Address.ValueString())
}

func (data BGPPeer) getClassName() string {
	return "bgpPeer"
}

func (data BGPPeer) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"addr", data.Address.ValueString())
	}
	if (!data.RemoteAsn.IsUnknown() && !data.RemoteAsn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"asn", data.RemoteAsn.ValueString())
	}
	if (!data.Description.IsUnknown() && !data.Description.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Description.ValueString())
	}
	if (!data.PeerTemplate.IsUnknown() && !data.PeerTemplate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"peerImp", data.PeerTemplate.ValueString())
	}
	if (!data.PeerType.IsUnknown() && !data.PeerType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"peerType", data.PeerType.ValueString())
	}
	if (!data.SourceInterface.IsUnknown() && !data.SourceInterface.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcIf", data.SourceInterface.ValueString())
	}
	if (!data.HoldTime.IsUnknown() && !data.HoldTime.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"holdIntvl", strconv.FormatInt(data.HoldTime.ValueInt64(), 10))
	}
	if (!data.Keepalive.IsUnknown() && !data.Keepalive.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"kaIntvl", strconv.FormatInt(data.Keepalive.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *BGPPeer) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.addr").String())
	} else {
		data.Address = types.StringNull()
	}
	if !data.RemoteAsn.IsNull() || all {
		data.RemoteAsn = types.StringValue(res.Get(data.getClassName() + ".attributes.asn").String())
	} else {
		data.RemoteAsn = types.StringNull()
	}
	if !data.Description.IsNull() || all {
		data.Description = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Description = types.StringNull()
	}
	if !data.PeerTemplate.IsNull() || all {
		data.PeerTemplate = types.StringValue(res.Get(data.getClassName() + ".attributes.peerImp").String())
	} else {
		data.PeerTemplate = types.StringNull()
	}
	if !data.PeerType.IsNull() || all {
		data.PeerType = types.StringValue(res.Get(data.getClassName() + ".attributes.peerType").String())
	} else {
		data.PeerType = types.StringNull()
	}
	if !data.SourceInterface.IsNull() || all {
		data.SourceInterface = types.StringValue(res.Get(data.getClassName() + ".attributes.srcIf").String())
	} else {
		data.SourceInterface = types.StringNull()
	}
	if !data.HoldTime.IsNull() || all {
		data.HoldTime = types.Int64Value(res.Get(data.getClassName() + ".attributes.holdIntvl").Int())
	} else {
		data.HoldTime = types.Int64Null()
	}
	if !data.Keepalive.IsNull() || all {
		data.Keepalive = types.Int64Value(res.Get(data.getClassName() + ".attributes.kaIntvl").Int())
	} else {
		data.Keepalive = types.Int64Null()
	}
}

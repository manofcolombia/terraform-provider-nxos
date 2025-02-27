// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	EbgpMultihopTtl types.Int64  `tfsdk:"ebgp_multihop_ttl"`
	PeerControl     types.String `tfsdk:"peer_control"`
	PasswordType    types.String `tfsdk:"password_type"`
	Password        types.String `tfsdk:"password"`
}

func (data BGPPeer) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]", data.Vrf.ValueString(), data.Address.ValueString())
}

func (data BGPPeer) getClassName() string {
	return "bgpPeer"
}

func (data BGPPeer) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
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
	if (!data.EbgpMultihopTtl.IsUnknown() && !data.EbgpMultihopTtl.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ttl", strconv.FormatInt(data.EbgpMultihopTtl.ValueInt64(), 10))
	}
	if (!data.PeerControl.IsUnknown() && !data.PeerControl.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ctrl", data.PeerControl.ValueString())
	}
	if (!data.PasswordType.IsUnknown() && !data.PasswordType.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"passwdType", data.PasswordType.ValueString())
	}
	if (!data.Password.IsUnknown() && !data.Password.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"password", data.Password.ValueString())
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
	if !data.EbgpMultihopTtl.IsNull() || all {
		data.EbgpMultihopTtl = types.Int64Value(res.Get(data.getClassName() + ".attributes.ttl").Int())
	} else {
		data.EbgpMultihopTtl = types.Int64Null()
	}
	if !data.PeerControl.IsNull() || all {
		data.PeerControl = types.StringValue(res.Get(data.getClassName() + ".attributes.ctrl").String())
	} else {
		data.PeerControl = types.StringNull()
	}
	if !data.PasswordType.IsNull() || all {
		data.PasswordType = types.StringValue(res.Get(data.getClassName() + ".attributes.passwdType").String())
	} else {
		data.PasswordType = types.StringNull()
	}
}

func (data BGPPeer) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *BGPPeer) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/bgp/inst/dom-[%s]/peer-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.Address = types.StringValue(matches[2])
}

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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPPeerTemplate struct {
	Device          types.String `tfsdk:"device"`
	Dn              types.String `tfsdk:"id"`
	Asn             types.String `tfsdk:"asn"`
	TemplateName    types.String `tfsdk:"template_name"`
	RemoteAsn       types.String `tfsdk:"remote_asn"`
	Description     types.String `tfsdk:"description"`
	PeerType        types.String `tfsdk:"peer_type"`
	SourceInterface types.String `tfsdk:"source_interface"`
}

func (data BGPPeerTemplate) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[default]/peercont-[%s]", data.TemplateName.ValueString())
}

func (data BGPPeerTemplate) getClassName() string {
	return "bgpPeerCont"
}

func (data BGPPeerTemplate) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.TemplateName.IsUnknown() && !data.TemplateName.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.TemplateName.ValueString())
	}
	if (!data.RemoteAsn.IsUnknown() && !data.RemoteAsn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"asn", data.RemoteAsn.ValueString())
	}
	if (!data.Description.IsUnknown() && !data.Description.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"desc", data.Description.ValueString())
	}
	if (!data.PeerType.IsUnknown() && !data.PeerType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"peerType", data.PeerType.ValueString())
	}
	if (!data.SourceInterface.IsUnknown() && !data.SourceInterface.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcIf", data.SourceInterface.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPPeerTemplate) fromBody(res gjson.Result, all bool) {
	if !data.TemplateName.IsNull() || all {
		data.TemplateName = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.TemplateName = types.StringNull()
	}
	if !data.RemoteAsn.IsNull() || all {
		data.RemoteAsn = types.StringValue(res.Get(data.getClassName() + ".attributes.asn").String())
	} else {
		data.RemoteAsn = types.StringNull()
	}
	if !data.Description.IsNull() || all {
		data.Description = types.StringValue(res.Get(data.getClassName() + ".attributes.desc").String())
	} else {
		data.Description = types.StringNull()
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
}

func (data BGPPeerTemplate) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *BGPPeerTemplate) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/bgp/inst/dom-[default]/peercont-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.TemplateName = types.StringValue(matches[1])
}

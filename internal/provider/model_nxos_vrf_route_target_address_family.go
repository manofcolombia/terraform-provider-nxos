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

type VRFRouteTargetAddressFamily struct {
	Device                   types.String `tfsdk:"device"`
	Dn                       types.String `tfsdk:"id"`
	Vrf                      types.String `tfsdk:"vrf"`
	AddressFamily            types.String `tfsdk:"address_family"`
	RouteTargetAddressFamily types.String `tfsdk:"route_target_address_family"`
}

func (data VRFRouteTargetAddressFamily) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]", data.Vrf.ValueString(), data.AddressFamily.ValueString(), data.RouteTargetAddressFamily.ValueString())
}

func (data VRFRouteTargetAddressFamily) getClassName() string {
	return "rtctrlAfCtrl"
}

func (data VRFRouteTargetAddressFamily) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.RouteTargetAddressFamily.IsUnknown() && !data.RouteTargetAddressFamily.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.RouteTargetAddressFamily.ValueString())
	}

	return nxos.Body{body}
}

func (data *VRFRouteTargetAddressFamily) fromBody(res gjson.Result, all bool) {
	if !data.RouteTargetAddressFamily.IsNull() || all {
		data.RouteTargetAddressFamily = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.RouteTargetAddressFamily = types.StringNull()
	}
}

func (data VRFRouteTargetAddressFamily) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *VRFRouteTargetAddressFamily) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.AddressFamily = types.StringValue(matches[2])
	data.RouteTargetAddressFamily = types.StringValue(matches[3])
}

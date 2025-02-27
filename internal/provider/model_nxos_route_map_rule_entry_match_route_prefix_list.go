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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouteMapRuleEntryMatchRoutePrefixList struct {
	Device       types.String `tfsdk:"device"`
	Dn           types.String `tfsdk:"id"`
	RuleName     types.String `tfsdk:"rule_name"`
	Order        types.Int64  `tfsdk:"order"`
	PrefixListDn types.String `tfsdk:"prefix_list_dn"`
}

func (data RouteMapRuleEntryMatchRoutePrefixList) getDn() string {
	return fmt.Sprintf("sys/rpm/rtmap-[%s]/ent-[%v]/mrtdst/rsrtDstAtt-[%s]", data.RuleName.ValueString(), data.Order.ValueInt64(), data.PrefixListDn.ValueString())
}

func (data RouteMapRuleEntryMatchRoutePrefixList) getClassName() string {
	return "rtmapRsRtDstAtt"
}

func (data RouteMapRuleEntryMatchRoutePrefixList) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.PrefixListDn.IsUnknown() && !data.PrefixListDn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tDn", data.PrefixListDn.ValueString())
	}

	return nxos.Body{body}
}

func (data *RouteMapRuleEntryMatchRoutePrefixList) fromBody(res gjson.Result, all bool) {
	if !data.PrefixListDn.IsNull() || all {
		data.PrefixListDn = types.StringValue(res.Get(data.getClassName() + ".attributes.tDn").String())
	} else {
		data.PrefixListDn = types.StringNull()
	}
}

func (data RouteMapRuleEntryMatchRoutePrefixList) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *RouteMapRuleEntryMatchRoutePrefixList) getIdsFromDn() {
	var RuleName string
	var Order int64
	var PrefixListDn string
	fmt.Sscanf(data.Dn.ValueString(), "sys/rpm/rtmap-[%s]/ent-[%v]/mrtdst/rsrtDstAtt-[%s]", &RuleName, &Order, &PrefixListDn)
	data.RuleName = types.StringValue(RuleName)
	data.Order = types.Int64Value(Order)
	data.PrefixListDn = types.StringValue(PrefixListDn)
}

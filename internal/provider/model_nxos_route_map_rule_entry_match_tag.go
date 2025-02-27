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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouteMapRuleEntryMatchTag struct {
	Device   types.String `tfsdk:"device"`
	Dn       types.String `tfsdk:"id"`
	RuleName types.String `tfsdk:"rule_name"`
	Order    types.Int64  `tfsdk:"order"`
	Tag      types.Int64  `tfsdk:"tag"`
}

func (data RouteMapRuleEntryMatchTag) getDn() string {
	return fmt.Sprintf("sys/rpm/rtmap-[%s]/ent-[%v]/mrttag-[%v]", data.RuleName.ValueString(), data.Order.ValueInt64(), data.Tag.ValueInt64())
}

func (data RouteMapRuleEntryMatchTag) getClassName() string {
	return "rtmapMatchRtTag"
}

func (data RouteMapRuleEntryMatchTag) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Tag.IsUnknown() && !data.Tag.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tag", strconv.FormatInt(data.Tag.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *RouteMapRuleEntryMatchTag) fromBody(res gjson.Result, all bool) {
	if !data.Tag.IsNull() || all {
		data.Tag = types.Int64Value(res.Get(data.getClassName() + ".attributes.tag").Int())
	} else {
		data.Tag = types.Int64Null()
	}
}

func (data RouteMapRuleEntryMatchTag) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *RouteMapRuleEntryMatchTag) getIdsFromDn() {
	var RuleName string
	var Order int64
	var Tag int64
	fmt.Sscanf(data.Dn.ValueString(), "sys/rpm/rtmap-[%s]/ent-[%v]/mrttag-[%v]", &RuleName, &Order, &Tag)
	data.RuleName = types.StringValue(RuleName)
	data.Order = types.Int64Value(Order)
	data.Tag = types.Int64Value(Tag)
}

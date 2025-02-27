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

type DefaultQOSPolicyMap struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	MatchType types.String `tfsdk:"match_type"`
}

func (data DefaultQOSPolicyMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/p/name-[%s]", data.Name.ValueString())
}

func (data DefaultQOSPolicyMap) getClassName() string {
	return "ipqosPMapInst"
}

func (data DefaultQOSPolicyMap) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}
	if (!data.MatchType.IsUnknown() && !data.MatchType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"matchType", data.MatchType.ValueString())
	}

	return nxos.Body{body}
}

func (data *DefaultQOSPolicyMap) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
	if !data.MatchType.IsNull() || all {
		data.MatchType = types.StringValue(res.Get(data.getClassName() + ".attributes.matchType").String())
	} else {
		data.MatchType = types.StringNull()
	}
}

func (data DefaultQOSPolicyMap) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *DefaultQOSPolicyMap) getIdsFromDn() {
	var Name string
	fmt.Sscanf(data.Dn.ValueString(), "sys/ipqos/dflt/p/name-[%s]", &Name)
	data.Name = types.StringValue(Name)
}

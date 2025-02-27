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

type EVPNVNI struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	Encap              types.String `tfsdk:"encap"`
	RouteDistinguisher types.String `tfsdk:"route_distinguisher"`
}

func (data EVPNVNI) getDn() string {
	return fmt.Sprintf("sys/evpn/bdevi-[%s]", data.Encap.ValueString())
}

func (data EVPNVNI) getClassName() string {
	return "rtctrlBDEvi"
}

func (data EVPNVNI) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Encap.IsUnknown() && !data.Encap.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"encap", data.Encap.ValueString())
	}
	if (!data.RouteDistinguisher.IsUnknown() && !data.RouteDistinguisher.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rd", data.RouteDistinguisher.ValueString())
	}

	return nxos.Body{body}
}

func (data *EVPNVNI) fromBody(res gjson.Result, all bool) {
	if !data.Encap.IsNull() || all {
		data.Encap = types.StringValue(res.Get(data.getClassName() + ".attributes.encap").String())
	} else {
		data.Encap = types.StringNull()
	}
	if !data.RouteDistinguisher.IsNull() || all {
		data.RouteDistinguisher = types.StringValue(res.Get(data.getClassName() + ".attributes.rd").String())
	} else {
		data.RouteDistinguisher = types.StringNull()
	}
}

func (data EVPNVNI) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *EVPNVNI) getIdsFromDn() {
	var Encap string
	fmt.Sscanf(data.Dn.ValueString(), "sys/evpn/bdevi-[%s]", &Encap)
	data.Encap = types.StringValue(Encap)
}

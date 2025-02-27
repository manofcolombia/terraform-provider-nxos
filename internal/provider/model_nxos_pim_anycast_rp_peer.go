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

type PIMAnycastRPPeer struct {
	Device       types.String `tfsdk:"device"`
	Dn           types.String `tfsdk:"id"`
	VrfName      types.String `tfsdk:"vrf_name"`
	Address      types.String `tfsdk:"address"`
	RpSetAddress types.String `tfsdk:"rp_set_address"`
}

func (data PIMAnycastRPPeer) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/acastrpfunc/peer-[%s]-peer-[%s]", data.VrfName.ValueString(), data.Address.ValueString(), data.RpSetAddress.ValueString())
}

func (data PIMAnycastRPPeer) getClassName() string {
	return "pimAcastRPPeer"
}

func (data PIMAnycastRPPeer) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"addr", data.Address.ValueString())
	}
	if (!data.RpSetAddress.IsUnknown() && !data.RpSetAddress.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rpSetAddr", data.RpSetAddress.ValueString())
	}

	return nxos.Body{body}
}

func (data *PIMAnycastRPPeer) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.addr").String())
	} else {
		data.Address = types.StringNull()
	}
	if !data.RpSetAddress.IsNull() || all {
		data.RpSetAddress = types.StringValue(res.Get(data.getClassName() + ".attributes.rpSetAddr").String())
	} else {
		data.RpSetAddress = types.StringNull()
	}
}

func (data PIMAnycastRPPeer) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *PIMAnycastRPPeer) getIdsFromDn() {
	var VrfName string
	var Address string
	var RpSetAddress string
	fmt.Sscanf(data.Dn.ValueString(), "sys/pim/inst/dom-[%s]/acastrpfunc/peer-[%s]-peer-[%s]", &VrfName, &Address, &RpSetAddress)
	data.VrfName = types.StringValue(VrfName)
	data.Address = types.StringValue(Address)
	data.RpSetAddress = types.StringValue(RpSetAddress)
}

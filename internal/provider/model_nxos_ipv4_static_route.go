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

type IPv4StaticRoute struct {
	Device   types.String              `tfsdk:"device"`
	Dn       types.String              `tfsdk:"id"`
	VrfName  types.String              `tfsdk:"vrf_name"`
	Prefix   types.String              `tfsdk:"prefix"`
	NextHops []IPv4StaticRouteNextHops `tfsdk:"next_hops"`
}

type IPv4StaticRouteNextHops struct {
	InterfaceId types.String `tfsdk:"interface_id"`
	Address     types.String `tfsdk:"address"`
	VrfName     types.String `tfsdk:"vrf_name"`
	Description types.String `tfsdk:"description"`
	Object      types.Int64  `tfsdk:"object"`
	Preference  types.Int64  `tfsdk:"preference"`
	Tag         types.Int64  `tfsdk:"tag"`
}

func (data IPv4StaticRoute) getDn() string {
	return fmt.Sprintf("sys/ipv4/inst/dom-[%s]/rt-[%s]", data.VrfName.ValueString(), data.Prefix.ValueString())
}

func (data IPv4StaticRouteNextHops) getRn() string {
	return fmt.Sprintf("nh-[%s]-addr-[%s]-vrf-[%s]", data.InterfaceId.ValueString(), data.Address.ValueString(), data.VrfName.ValueString())
}

func (data IPv4StaticRoute) getClassName() string {
	return "ipv4Route"
}

func (data IPv4StaticRoute) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Prefix.IsUnknown() && !data.Prefix.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"prefix", data.Prefix.ValueString())
	}
	var attrs string
	for _, child := range data.NextHops {
		attrs = ""
		if (!child.InterfaceId.IsUnknown() && !child.InterfaceId.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "nhIf", child.InterfaceId.ValueString())
		}
		if (!child.Address.IsUnknown() && !child.Address.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "nhAddr", child.Address.ValueString())
		}
		if (!child.VrfName.IsUnknown() && !child.VrfName.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "nhVrf", child.VrfName.ValueString())
		}
		if (!child.Description.IsUnknown() && !child.Description.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "descr", child.Description.ValueString())
		}
		if (!child.Object.IsUnknown() && !child.Object.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "object", strconv.FormatInt(child.Object.ValueInt64(), 10))
		}
		if (!child.Preference.IsUnknown() && !child.Preference.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "pref", strconv.FormatInt(child.Preference.ValueInt64(), 10))
		}
		if (!child.Tag.IsUnknown() && !child.Tag.IsNull()) || true {
			attrs, _ = sjson.Set(attrs, "tag", strconv.FormatInt(child.Tag.ValueInt64(), 10))
		}
		body, _ = sjson.SetRaw(body, data.getClassName()+".children.-1.ipv4Nexthop.attributes", attrs)
	}

	return nxos.Body{body}
}

func (data *IPv4StaticRoute) fromBody(res gjson.Result, all bool) {
	if !data.Prefix.IsNull() || all {
		data.Prefix = types.StringValue(res.Get(data.getClassName() + ".attributes.prefix").String())
	} else {
		data.Prefix = types.StringNull()
	}
	if all {
		res.Get(data.getClassName() + ".children").ForEach(
			func(_, v gjson.Result) bool {
				v.ForEach(
					func(classname, value gjson.Result) bool {
						if classname.String() == "ipv4Nexthop" {
							var child IPv4StaticRouteNextHops
							child.InterfaceId = types.StringValue(value.Get("attributes.nhIf").String())
							child.Address = types.StringValue(value.Get("attributes.nhAddr").String())
							child.VrfName = types.StringValue(value.Get("attributes.nhVrf").String())
							child.Description = types.StringValue(value.Get("attributes.descr").String())
							child.Object = types.Int64Value(value.Get("attributes.object").Int())
							child.Preference = types.Int64Value(value.Get("attributes.pref").Int())
							child.Tag = types.Int64Value(value.Get("attributes.tag").Int())
							data.NextHops = append(data.NextHops, child)
						}
						return true
					},
				)
				return true
			},
		)
	} else {
		for c := range data.NextHops {
			var r gjson.Result
			res.Get(data.getClassName() + ".children").ForEach(
				func(_, v gjson.Result) bool {
					key := v.Get("ipv4Nexthop.attributes.rn").String()
					if key == data.NextHops[c].getRn() {
						r = v
						return false
					}
					return true
				},
			)
			if !data.NextHops[c].InterfaceId.IsNull() || all {
				data.NextHops[c].InterfaceId = types.StringValue(r.Get("ipv4Nexthop.attributes.nhIf").String())
			} else {
				data.NextHops[c].InterfaceId = types.StringNull()
			}
			if !data.NextHops[c].Address.IsNull() || all {
				data.NextHops[c].Address = types.StringValue(r.Get("ipv4Nexthop.attributes.nhAddr").String())
			} else {
				data.NextHops[c].Address = types.StringNull()
			}
			if !data.NextHops[c].VrfName.IsNull() || all {
				data.NextHops[c].VrfName = types.StringValue(r.Get("ipv4Nexthop.attributes.nhVrf").String())
			} else {
				data.NextHops[c].VrfName = types.StringNull()
			}
			if !data.NextHops[c].Description.IsNull() || all {
				data.NextHops[c].Description = types.StringValue(r.Get("ipv4Nexthop.attributes.descr").String())
			} else {
				data.NextHops[c].Description = types.StringNull()
			}
			if !data.NextHops[c].Object.IsNull() || all {
				data.NextHops[c].Object = types.Int64Value(r.Get("ipv4Nexthop.attributes.object").Int())
			} else {
				data.NextHops[c].Object = types.Int64Null()
			}
			if !data.NextHops[c].Preference.IsNull() || all {
				data.NextHops[c].Preference = types.Int64Value(r.Get("ipv4Nexthop.attributes.pref").Int())
			} else {
				data.NextHops[c].Preference = types.Int64Null()
			}
			if !data.NextHops[c].Tag.IsNull() || all {
				data.NextHops[c].Tag = types.Int64Value(r.Get("ipv4Nexthop.attributes.tag").Int())
			} else {
				data.NextHops[c].Tag = types.Int64Null()
			}
		}
	}
}

func (data IPv4StaticRoute) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *IPv4StaticRoute) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ipv4/inst/dom-[%s]/rt-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.VrfName = types.StringValue(matches[1])
	data.Prefix = types.StringValue(matches[2])
}

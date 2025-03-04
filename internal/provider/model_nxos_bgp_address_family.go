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

type BGPAddressFamily struct {
	Device                            types.String `tfsdk:"device"`
	Dn                                types.String `tfsdk:"id"`
	Asn                               types.String `tfsdk:"asn"`
	Vrf                               types.String `tfsdk:"vrf"`
	AddressFamily                     types.String `tfsdk:"address_family"`
	CriticalNexthopTimeout            types.String `tfsdk:"critical_nexthop_timeout"`
	NonCriticalNexthopTimeout         types.String `tfsdk:"non_critical_nexthop_timeout"`
	AdvertiseL2vpnEvpn                types.String `tfsdk:"advertise_l2vpn_evpn"`
	AdvertisePhysicalIpForType5Routes types.String `tfsdk:"advertise_physical_ip_for_type5_routes"`
	MaxEcmpPaths                      types.Int64  `tfsdk:"max_ecmp_paths"`
	MaxExternalEcmpPaths              types.Int64  `tfsdk:"max_external_ecmp_paths"`
	MaxExternalInternalEcmpPaths      types.Int64  `tfsdk:"max_external_internal_ecmp_paths"`
	MaxLocalEcmpPaths                 types.Int64  `tfsdk:"max_local_ecmp_paths"`
	MaxMixedEcmpPaths                 types.Int64  `tfsdk:"max_mixed_ecmp_paths"`
	DefaultInformationOriginate       types.String `tfsdk:"default_information_originate"`
	NextHopRouteMapName               types.String `tfsdk:"next_hop_route_map_name"`
	PrefixPriority                    types.String `tfsdk:"prefix_priority"`
	RetainRtAll                       types.String `tfsdk:"retain_rt_all"`
	AdvertiseOnlyActiveRoutes         types.String `tfsdk:"advertise_only_active_routes"`
	TableMapRouteMapName              types.String `tfsdk:"table_map_route_map_name"`
	VniEthernetTag                    types.String `tfsdk:"vni_ethernet_tag"`
	WaitIgpConverged                  types.String `tfsdk:"wait_igp_converged"`
}

func (data BGPAddressFamily) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/af-[%s]", data.Vrf.ValueString(), data.AddressFamily.ValueString())
}

func (data BGPAddressFamily) getClassName() string {
	return "bgpDomAf"
}

func (data BGPAddressFamily) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.AddressFamily.IsUnknown() && !data.AddressFamily.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.AddressFamily.ValueString())
	}
	if (!data.CriticalNexthopTimeout.IsUnknown() && !data.CriticalNexthopTimeout.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"critNhTimeout", data.CriticalNexthopTimeout.ValueString())
	}
	if (!data.NonCriticalNexthopTimeout.IsUnknown() && !data.NonCriticalNexthopTimeout.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nonCritNhTimeout", data.NonCriticalNexthopTimeout.ValueString())
	}
	if (!data.AdvertiseL2vpnEvpn.IsUnknown() && !data.AdvertiseL2vpnEvpn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advertL2vpnEvpn", data.AdvertiseL2vpnEvpn.ValueString())
	}
	if (!data.AdvertisePhysicalIpForType5Routes.IsUnknown() && !data.AdvertisePhysicalIpForType5Routes.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advPip", data.AdvertisePhysicalIpForType5Routes.ValueString())
	}
	if (!data.MaxEcmpPaths.IsUnknown() && !data.MaxEcmpPaths.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxEcmp", strconv.FormatInt(data.MaxEcmpPaths.ValueInt64(), 10))
	}
	if (!data.MaxExternalEcmpPaths.IsUnknown() && !data.MaxExternalEcmpPaths.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxExtEcmp", strconv.FormatInt(data.MaxExternalEcmpPaths.ValueInt64(), 10))
	}
	if (!data.MaxExternalInternalEcmpPaths.IsUnknown() && !data.MaxExternalInternalEcmpPaths.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxExtIntEcmp", strconv.FormatInt(data.MaxExternalInternalEcmpPaths.ValueInt64(), 10))
	}
	if (!data.MaxLocalEcmpPaths.IsUnknown() && !data.MaxLocalEcmpPaths.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxLclEcmp", strconv.FormatInt(data.MaxLocalEcmpPaths.ValueInt64(), 10))
	}
	if (!data.MaxMixedEcmpPaths.IsUnknown() && !data.MaxMixedEcmpPaths.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"maxMxdEcmp", strconv.FormatInt(data.MaxMixedEcmpPaths.ValueInt64(), 10))
	}
	if (!data.DefaultInformationOriginate.IsUnknown() && !data.DefaultInformationOriginate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"defInfOriginate", data.DefaultInformationOriginate.ValueString())
	}
	if (!data.NextHopRouteMapName.IsUnknown() && !data.NextHopRouteMapName.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nhRtMap", data.NextHopRouteMapName.ValueString())
	}
	if (!data.PrefixPriority.IsUnknown() && !data.PrefixPriority.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"prfxPriority", data.PrefixPriority.ValueString())
	}
	if (!data.RetainRtAll.IsUnknown() && !data.RetainRtAll.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"retainRttAll", data.RetainRtAll.ValueString())
	}
	if (!data.AdvertiseOnlyActiveRoutes.IsUnknown() && !data.AdvertiseOnlyActiveRoutes.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"supprInactive", data.AdvertiseOnlyActiveRoutes.ValueString())
	}
	if (!data.TableMapRouteMapName.IsUnknown() && !data.TableMapRouteMapName.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tblMap", data.TableMapRouteMapName.ValueString())
	}
	if (!data.VniEthernetTag.IsUnknown() && !data.VniEthernetTag.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"vniEthTag", data.VniEthernetTag.ValueString())
	}
	if (!data.WaitIgpConverged.IsUnknown() && !data.WaitIgpConverged.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"waitIgpConv", data.WaitIgpConverged.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPAddressFamily) fromBody(res gjson.Result, all bool) {
	if !data.AddressFamily.IsNull() || all {
		data.AddressFamily = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.AddressFamily = types.StringNull()
	}
	if !data.CriticalNexthopTimeout.IsNull() || all {
		data.CriticalNexthopTimeout = types.StringValue(res.Get(data.getClassName() + ".attributes.critNhTimeout").String())
	} else {
		data.CriticalNexthopTimeout = types.StringNull()
	}
	if !data.NonCriticalNexthopTimeout.IsNull() || all {
		data.NonCriticalNexthopTimeout = types.StringValue(res.Get(data.getClassName() + ".attributes.nonCritNhTimeout").String())
	} else {
		data.NonCriticalNexthopTimeout = types.StringNull()
	}
	if !data.AdvertiseL2vpnEvpn.IsNull() || all {
		data.AdvertiseL2vpnEvpn = types.StringValue(res.Get(data.getClassName() + ".attributes.advertL2vpnEvpn").String())
	} else {
		data.AdvertiseL2vpnEvpn = types.StringNull()
	}
	if !data.AdvertisePhysicalIpForType5Routes.IsNull() || all {
		data.AdvertisePhysicalIpForType5Routes = types.StringValue(res.Get(data.getClassName() + ".attributes.advPip").String())
	} else {
		data.AdvertisePhysicalIpForType5Routes = types.StringNull()
	}
	if !data.MaxEcmpPaths.IsNull() || all {
		data.MaxEcmpPaths = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxEcmp").Int())
	} else {
		data.MaxEcmpPaths = types.Int64Null()
	}
	if !data.MaxExternalEcmpPaths.IsNull() || all {
		data.MaxExternalEcmpPaths = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxExtEcmp").Int())
	} else {
		data.MaxExternalEcmpPaths = types.Int64Null()
	}
	if !data.MaxExternalInternalEcmpPaths.IsNull() || all {
		data.MaxExternalInternalEcmpPaths = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxExtIntEcmp").Int())
	} else {
		data.MaxExternalInternalEcmpPaths = types.Int64Null()
	}
	if !data.MaxLocalEcmpPaths.IsNull() || all {
		data.MaxLocalEcmpPaths = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxLclEcmp").Int())
	} else {
		data.MaxLocalEcmpPaths = types.Int64Null()
	}
	if !data.MaxMixedEcmpPaths.IsNull() || all {
		data.MaxMixedEcmpPaths = types.Int64Value(res.Get(data.getClassName() + ".attributes.maxMxdEcmp").Int())
	} else {
		data.MaxMixedEcmpPaths = types.Int64Null()
	}
	if !data.DefaultInformationOriginate.IsNull() || all {
		data.DefaultInformationOriginate = types.StringValue(res.Get(data.getClassName() + ".attributes.defInfOriginate").String())
	} else {
		data.DefaultInformationOriginate = types.StringNull()
	}
	if !data.NextHopRouteMapName.IsNull() || all {
		data.NextHopRouteMapName = types.StringValue(res.Get(data.getClassName() + ".attributes.nhRtMap").String())
	} else {
		data.NextHopRouteMapName = types.StringNull()
	}
	if !data.PrefixPriority.IsNull() || all {
		data.PrefixPriority = types.StringValue(res.Get(data.getClassName() + ".attributes.prfxPriority").String())
	} else {
		data.PrefixPriority = types.StringNull()
	}
	if !data.RetainRtAll.IsNull() || all {
		data.RetainRtAll = types.StringValue(res.Get(data.getClassName() + ".attributes.retainRttAll").String())
	} else {
		data.RetainRtAll = types.StringNull()
	}
	if !data.AdvertiseOnlyActiveRoutes.IsNull() || all {
		data.AdvertiseOnlyActiveRoutes = types.StringValue(res.Get(data.getClassName() + ".attributes.supprInactive").String())
	} else {
		data.AdvertiseOnlyActiveRoutes = types.StringNull()
	}
	if !data.TableMapRouteMapName.IsNull() || all {
		data.TableMapRouteMapName = types.StringValue(res.Get(data.getClassName() + ".attributes.tblMap").String())
	} else {
		data.TableMapRouteMapName = types.StringNull()
	}
	if !data.VniEthernetTag.IsNull() || all {
		data.VniEthernetTag = types.StringValue(res.Get(data.getClassName() + ".attributes.vniEthTag").String())
	} else {
		data.VniEthernetTag = types.StringNull()
	}
	if !data.WaitIgpConverged.IsNull() || all {
		data.WaitIgpConverged = types.StringValue(res.Get(data.getClassName() + ".attributes.waitIgpConv").String())
	} else {
		data.WaitIgpConverged = types.StringNull()
	}
}

func (data BGPAddressFamily) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *BGPAddressFamily) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/bgp/inst/dom-[%s]/af-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.AddressFamily = types.StringValue(matches[2])
}

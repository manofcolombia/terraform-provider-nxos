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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNxosVRFRouteTarget(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosVRFRouteTargetPrerequisitesConfig + testAccDataSourceNxosVRFRouteTargetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_vrf_route_target.test", "route_target", "route-target:as2-nn2:2:2"),
				),
			},
		},
	})
}

const testAccDataSourceNxosVRFRouteTargetPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/inst-[VRF1]"
  class_name = "l3Inst"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ipv4/inst/dom-[VRF1]"
  class_name = "ipv4Dom"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]"
  class_name = "rtctrlDom"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]"
  class_name = "rtctrlDomAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]"
  class_name = "rtctrlAfCtrl"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq6" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]/rttp-[import]"
  class_name = "rtctrlRttP"
  content = {
      type = "import"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

const testAccDataSourceNxosVRFRouteTargetConfig = `

resource "nxos_vrf_route_target" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, nxos_rest.PreReq6, ]
}

data "nxos_vrf_route_target" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
  depends_on = [nxos_vrf_route_target.test]
}
`

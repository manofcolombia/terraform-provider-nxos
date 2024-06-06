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

func TestAccNxosSubinterfaceVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosSubinterfaceVRFPrerequisitesConfig + testAccNxosSubinterfaceVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "interface_id", "eth1/10.124"),
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "vrf_dn", "sys/inst-VRF123"),
				),
			},
			{
				ResourceName:  "nxos_subinterface_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/intf/encrtd-[eth1/10.124]/rtvrfMbr",
			},
		},
	})
}

const testAccNxosSubinterfaceVRFPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      id = "eth1/10"
      layer = "Layer3"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/encrtd-[eth1/10.124]"
  class_name = "l3EncRtdIf"
  content = {
      id = "eth1/10.124"
  }
}

`

func testAccNxosSubinterfaceVRFConfig_minimum() string {
	return `
	resource "nxos_subinterface_vrf" "test" {
		interface_id = "eth1/10.124"
		vrf_dn = "sys/inst-VRF123"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosSubinterfaceVRFConfig_all() string {
	return `
	resource "nxos_subinterface_vrf" "test" {
		interface_id = "eth1/10.124"
		vrf_dn = "sys/inst-VRF123"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

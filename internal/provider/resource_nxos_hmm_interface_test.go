// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosHMMInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosHMMInterfacePrerequisitesConfig + testAccNxosHMMInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_hmm_interface.test", "interface_id", "vlan10"),
					resource.TestCheckResourceAttr("nxos_hmm_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_hmm_interface.test", "mode", "anycastGW"),
				),
			},
			{
				ResourceName:  "nxos_hmm_interface.test",
				ImportState:   true,
				ImportStateId: "sys/hmm/fwdinst/if-[vlan10]",
			},
		},
	})
}

const testAccNxosHMMInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ifvlan"
  class_name = "fmInterfaceVlan"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/svi-[vlan10]"
  class_name = "sviIf"
  content = {
      id = "vlan10"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/fm/hmm"
  class_name = "fmHmm"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/hmm"
  class_name = "hmmEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/hmm/fwdinst"
  class_name = "hmmFwdInst"
  content = {
      adminSt = "enabled"
      amac = "20:20:00:00:10:10"
  }
  depends_on = [nxos_rest.PreReq1, nxos_rest.PreReq4, ]
}

`

func testAccNxosHMMInterfaceConfig_minimum() string {
	return `
	resource "nxos_hmm_interface" "test" {
		interface_id = "vlan10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

func testAccNxosHMMInterfaceConfig_all() string {
	return `
	resource "nxos_hmm_interface" "test" {
		interface_id = "vlan10"
		admin_state = "enabled"
		mode = "anycastGW"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

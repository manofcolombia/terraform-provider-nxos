// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosVPCInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVPCInstancePrerequisitesConfig + testAccNxosVPCInstanceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vpc_instance.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_vpc_instance.test",
				ImportState:   true,
				ImportStateId: "sys/vpc/inst",
			},
		},
	})
}

const testAccNxosVPCInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/vpc"
  class_name = "fmVpc"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

`

func testAccNxosVPCInstanceConfig_minimum() string {
	return `
	resource "nxos_vpc_instance" "test" {
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosVPCInstanceConfig_all() string {
	return `
	resource "nxos_vpc_instance" "test" {
		admin_state = "enabled"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

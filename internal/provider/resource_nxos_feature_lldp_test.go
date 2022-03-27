// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureLLDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureLLDPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_lldp.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_lldp.test",
				ImportState:   true,
				ImportStateId: "sys/fm/lldp",
			},
		},
	})
}

func testAccNxosFeatureLLDPConfig_minimum() string {
	return `
	resource "nxos_feature_lldp" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureLLDPConfig_all() string {
	return `
	resource "nxos_feature_lldp" "test" {
		admin_state = "enabled"
	}
	`
}

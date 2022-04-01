// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureHSRP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureHSRPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_hsrp.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_hsrp.test",
				ImportState:   true,
				ImportStateId: "sys/fm/hsrp",
			},
		},
	})
}

func testAccNxosFeatureHSRPConfig_minimum() string {
	return `
	resource "nxos_feature_hsrp" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureHSRPConfig_all() string {
	return `
	resource "nxos_feature_hsrp" "test" {
		admin_state = "enabled"
	}
	`
}

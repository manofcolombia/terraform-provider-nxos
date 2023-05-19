// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosRouteMapRuleEntryMatchRoutePrefixList(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosRouteMapRuleEntryMatchRoutePrefixListPrerequisitesConfig + testAccNxosRouteMapRuleEntryMatchRoutePrefixListConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_route_map_rule_entry_match_route_prefix_list.test", "rule_name", "RULE1"),
					resource.TestCheckResourceAttr("nxos_route_map_rule_entry_match_route_prefix_list.test", "order", "10"),
					resource.TestCheckResourceAttr("nxos_route_map_rule_entry_match_route_prefix_list.test", "prefix_list_dn", "sys/rpm/pfxlistv4-[LIST1]"),
				),
			},
			{
				ResourceName:  "nxos_route_map_rule_entry_match_route_prefix_list.test",
				ImportState:   true,
				ImportStateId: "sys/rpm/rtmap-[RULE1]/ent-[10]/mrtdst/rsrtDstAtt-[sys/rpm/pfxlistv4-[LIST1]]",
			},
		},
	})
}

const testAccNxosRouteMapRuleEntryMatchRoutePrefixListPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/rpm/rtmap-[RULE1]"
  class_name = "rtmapRule"
  content = {
      name = "RULE1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/rpm/rtmap-[RULE1]/ent-[10]"
  class_name = "rtmapEntry"
  content = {
      order = "10"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/rpm/rtmap-[RULE1]/ent-[10]/mrtdst"
  class_name = "rtmapMatchRtDst"
  depends_on = [nxos_rest.PreReq1, ]
}

`

func testAccNxosRouteMapRuleEntryMatchRoutePrefixListConfig_minimum() string {
	return `
	resource "nxos_route_map_rule_entry_match_route_prefix_list" "test" {
		rule_name = "RULE1"
		order = 10
		prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosRouteMapRuleEntryMatchRoutePrefixListConfig_all() string {
	return `
	resource "nxos_route_map_rule_entry_match_route_prefix_list" "test" {
		rule_name = "RULE1"
		order = 10
		prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

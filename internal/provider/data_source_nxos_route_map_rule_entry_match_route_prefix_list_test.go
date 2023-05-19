// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosRouteMapRuleEntryMatchRoutePrefixList(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosRouteMapRuleEntryMatchRoutePrefixListPrerequisitesConfig + testAccDataSourceNxosRouteMapRuleEntryMatchRoutePrefixListConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_route_map_rule_entry_match_route_prefix_list.test", "prefix_list_dn", "sys/rpm/pfxlistv4-[LIST1]"),
				),
			},
		},
	})
}

const testAccDataSourceNxosRouteMapRuleEntryMatchRoutePrefixListPrerequisitesConfig = `
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

const testAccDataSourceNxosRouteMapRuleEntryMatchRoutePrefixListConfig = `

resource "nxos_route_map_rule_entry_match_route_prefix_list" "test" {
  rule_name = "RULE1"
  order = 10
  prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_route_map_rule_entry_match_route_prefix_list" "test" {
  rule_name = "RULE1"
  order = 10
  prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
  depends_on = [nxos_route_map_rule_entry_match_route_prefix_list.test]
}
`

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMStaticRPPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPIMStaticRPPolicyPrerequisitesConfig + testAccDataSourceNxosPIMStaticRPPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_static_rp_policy.test", "name", "RP"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMStaticRPPolicyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosPIMStaticRPPolicyConfig = `

resource "nxos_pim_static_rp_policy" "test" {
  vrf_name = "default"
  name = "RP"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_pim_static_rp_policy" "test" {
  vrf_name = "default"
  depends_on = [nxos_pim_static_rp_policy.test]
}
`

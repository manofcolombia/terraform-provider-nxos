// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPGracefulRestart(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPGracefulRestartPrerequisitesConfig + testAccNxosBGPGracefulRestartConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_graceful_restart.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_graceful_restart.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_graceful_restart.test", "restart_interval", "240"),
					resource.TestCheckResourceAttr("nxos_bgp_graceful_restart.test", "stale_interval", "1800"),
				),
			},
			{
				ResourceName:  "nxos_bgp_graceful_restart.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/gr",
			},
		},
	})
}

const testAccNxosBGPGracefulRestartPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosBGPGracefulRestartConfig_minimum() string {
	return `
	resource "nxos_bgp_graceful_restart" "test" {
		asn = "65001"
		vrf = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosBGPGracefulRestartConfig_all() string {
	return `
	resource "nxos_bgp_graceful_restart" "test" {
		asn = "65001"
		vrf = "default"
		restart_interval = 240
		stale_interval = 1800
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

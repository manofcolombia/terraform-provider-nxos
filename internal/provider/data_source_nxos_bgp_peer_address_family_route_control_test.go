// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosBGPPeerAddressFamilyRouteControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPPeerAddressFamilyRouteControlPrerequisitesConfig + testAccDataSourceNxosBGPPeerAddressFamilyRouteControlConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_address_family_route_control.test", "direction", "in"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_address_family_route_control.test", "route_map_name", "ROUTE_MAP1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPPeerAddressFamilyRouteControlPrerequisitesConfig = `
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

resource "nxos_rest" "PreReq4" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]"
  class_name = "bgpPeer"
  content = {
      addr = "192.168.0.1"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]"
  class_name = "bgpPeerAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

const testAccDataSourceNxosBGPPeerAddressFamilyRouteControlConfig = `

resource "nxos_bgp_peer_address_family_route_control" "test" {
  asn = "65001"
  vrf = "default"
  address = "192.168.0.1"
  address_family = "ipv4-ucast"
  direction = "in"
  route_map_name = "ROUTE_MAP1"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
}

data "nxos_bgp_peer_address_family_route_control" "test" {
  asn = "65001"
  vrf = "default"
  address = "192.168.0.1"
  address_family = "ipv4-ucast"
  direction = "in"
  depends_on = [nxos_bgp_peer_address_family_route_control.test]
}
`

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosOSPFAuthentication(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosOSPFAuthenticationPrerequisitesConfig + testAccDataSourceNxosOSPFAuthenticationConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ospf_authentication.test", "key_id", "1"),
					resource.TestCheckResourceAttr("data.nxos_ospf_authentication.test", "key_secure_mode", "false"),
					resource.TestCheckResourceAttr("data.nxos_ospf_authentication.test", "keychain", "mykeychain"),
					resource.TestCheckResourceAttr("data.nxos_ospf_authentication.test", "md5_key_secure_mode", "false"),
					resource.TestCheckResourceAttr("data.nxos_ospf_authentication.test", "type", "none"),
				),
			},
		},
	})
}

const testAccDataSourceNxosOSPFAuthenticationPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ospf"
  class_name = "fmOspf"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]"
  class_name = "ospfDom"
  content = {
      name = "VRF1"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/intf/phys-[eth1/10]/rtvrfMbr"
  class_name = "nwRtVrfMbr"
  content = {
      tDn = "sys/inst-VRF1"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

resource "nxos_rest" "PreReq6" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]/if-[eth1/10]"
  class_name = "ospfIf"
  content = {
      id = "eth1/10"
  }
  depends_on = [nxos_rest.PreReq5, ]
}

`

const testAccDataSourceNxosOSPFAuthenticationConfig = `

resource "nxos_ospf_authentication" "test" {
  instance_name = "OSPF1"
  vrf_name = "VRF1"
  interface_id = "eth1/10"
  key = "0 mykey"
  key_id = 1
  key_secure_mode = false
  keychain = "mykeychain"
  md5_key = "0 mymd5key"
  md5_key_secure_mode = false
  type = "none"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, nxos_rest.PreReq6, ]
}

data "nxos_ospf_authentication" "test" {
  instance_name = "OSPF1"
  vrf_name = "VRF1"
  interface_id = "eth1/10"
  depends_on = [nxos_ospf_authentication.test]
}
`

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPortChannelInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPortChannelInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "interface_id", "po1"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "port_channel_mode", "active"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "minimum_links", "2"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "maximum_links", "10"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "suspend_individual", "disable"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "access_vlan", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "admin_state", "up"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "auto_negotiation", "on"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "delay", "10"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "duplex", "auto"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "layer", "Layer3"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "link_logging", "enable"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "medium", "broadcast"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "mode", "access"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "mtu", "1500"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "native_vlan", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "speed", "auto"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "trunk_vlans", "1-4094"),
					resource.TestCheckResourceAttr("data.nxos_port_channel_interface.test", "user_configured_flags", "admin_layer,admin_mtu,admin_state"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPortChannelInterfaceConfig = `

resource "nxos_port_channel_interface" "test" {
  interface_id = "po1"
  port_channel_mode = "active"
  minimum_links = 2
  maximum_links = 10
  suspend_individual = "disable"
  access_vlan = "unknown"
  admin_state = "up"
  auto_negotiation = "on"
  bandwidth = 1000
  delay = 10
  description = "My Description"
  duplex = "auto"
  layer = "Layer3"
  link_logging = "enable"
  medium = "broadcast"
  mode = "access"
  mtu = 1500
  native_vlan = "unknown"
  speed = "auto"
  trunk_vlans = "1-4094"
  user_configured_flags = "admin_layer,admin_mtu,admin_state"
}

data "nxos_port_channel_interface" "test" {
  interface_id = "po1"
  depends_on = [nxos_port_channel_interface.test]
}
`

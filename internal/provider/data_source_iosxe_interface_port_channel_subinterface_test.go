// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxeInterfacePortChannelSubinterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfacePortChannelSubinterfacePrerequisitesConfig + testAccDataSourceIosxeInterfacePortChannelSubinterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_interface_port_channel_subinterface.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxe_interface_port_channel_subinterface.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("data.iosxe_interface_port_channel_subinterface.test", "vrf_forwarding", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxe_interface_port_channel_subinterface.test", "ipv4_address", "192.0.2.2"),
					resource.TestCheckResourceAttr("data.iosxe_interface_port_channel_subinterface.test", "ipv4_address_mask", "255.255.255.0"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeInterfacePortChannelSubinterfacePrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      name = "VRF1"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1/address-family"
  delete = false
  attributes = {
      ipv4 = ""
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
  path = "Cisco-IOS-XE-native:native/interface/Port-channel=10"
  attributes = {
      name = "10"
      switchport = "false"
  }
}

`

const testAccDataSourceIosxeInterfacePortChannelSubinterfaceConfig = `

resource "iosxe_interface_port_channel_subinterface" "test" {
  name = "10.666"
  description = "My Interface Description"
  shutdown = false
  vrf_forwarding = "VRF1"
  ipv4_address = "192.0.2.2"
  ipv4_address_mask = "255.255.255.0"
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]
}

data "iosxe_interface_port_channel_subinterface" "test" {
  name = "10.666"
  depends_on = [iosxe_interface_port_channel_subinterface.test]
}
`

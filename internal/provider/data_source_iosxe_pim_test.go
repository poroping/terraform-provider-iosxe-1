// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxePIM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxePIMPrerequisitesConfig + testAccDataSourceIosxePIMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "autorp", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "autorp_listener", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "bsr_candidate_loopback", "100"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "bsr_candidate_mask", "30"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "bsr_candidate_priority", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "bsr_candidate_accept_rp_candidate", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "ssm_range", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "ssm_default", "true"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_addresses.0.access_list", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_addresses.0.rp_address", "10.10.10.10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_addresses.0.override", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_addresses.0.bidir", "false"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_candidates.0.interface", "Loopback100"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_candidates.0.group_list", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_candidates.0.interval", "100"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_candidates.0.priority", "10"),
					resource.TestCheckResourceAttr("data.iosxe_pim.test", "rp_candidates.0.bidir", "false"),
				),
			},
		},
	})
}

const testAccDataSourceIosxePIMPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
  attributes = {
      name = "100"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/interface/Loopback=100/ip/address/primary"
  attributes = {
      address = "200.200.200.200"
      mask = "255.255.255.255"
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

`

const testAccDataSourceIosxePIMConfig = `

resource "iosxe_pim" "test" {
  autorp = false
  autorp_listener = false
  bsr_candidate_loopback = 100
  bsr_candidate_mask = 30
  bsr_candidate_priority = 10
  bsr_candidate_accept_rp_candidate = "10"
  ssm_range = "10"
  ssm_default = true
  rp_addresses = [{
    access_list = "10"
    rp_address = "10.10.10.10"
    override = false
    bidir = false
  }]
  rp_candidates = [{
    interface = "Loopback100"
    group_list = "10"
    interval = 100
    priority = 10
    bidir = false
  }]
  depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
}

data "iosxe_pim" "test" {
  depends_on = [iosxe_pim.test]
}
`

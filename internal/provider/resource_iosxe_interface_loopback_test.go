// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeInterfaceLoopback(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceLoopbackPrerequisitesConfig + testAccIosxeInterfaceLoopbackConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "name", "100"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "vrf_forwarding", "VRF1"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ipv4_address", "200.1.1.1"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "ipv4_address_mask", "255.255.255.255"),
					resource.TestCheckResourceAttr("iosxe_interface_loopback.test", "pim_sparse_mode", "true"),
				),
			},
			{
				ResourceName:  "iosxe_interface_loopback.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Loopback=100",
			},
		},
	})
}

const testAccIosxeInterfaceLoopbackPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
  delete = false
  attributes = {
      name = "VRF1"
  }
}

resource "iosxe_restconf" "PreReq1" {
  path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1/address-family"
  attributes = {
      ipv4 = ""
  }
  depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxeInterfaceLoopbackConfig_minimum() string {
	return `
	resource "iosxe_interface_loopback" "test" {
		name = 100
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

func testAccIosxeInterfaceLoopbackConfig_all() string {
	return `
	resource "iosxe_interface_loopback" "test" {
		name = 100
		description = "My Interface Description"
		shutdown = false
		vrf_forwarding = "VRF1"
		ipv4_address = "200.1.1.1"
		ipv4_address_mask = "255.255.255.255"
		pim_sparse_mode = true
  		depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]
	}
	`
}

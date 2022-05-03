//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeVLANConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeVLANConfigurationConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_vlan_configuration.test", "vlan_id", "123"),
					resource.TestCheckResourceAttr("iosxe_vlan_configuration.test", "evpn_instance", "123"),
					resource.TestCheckResourceAttr("iosxe_vlan_configuration.test", "evpn_instance_vni", "10123"),
				),
			},
			{
				ResourceName:  "iosxe_vlan_configuration.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:configuration=123",
			},
		},
	})
}

func testAccIosxeVLANConfigurationConfig_minimum() string {
	return `
	resource "iosxe_vlan_configuration" "test" {
		vlan_id = 123
	}
	`
}

func testAccIosxeVLANConfigurationConfig_all() string {
	return `
	resource "iosxe_vlan_configuration" "test" {
		vlan_id = 123
		evpn_instance = 123
		evpn_instance_vni = 10123
	}
	`
}

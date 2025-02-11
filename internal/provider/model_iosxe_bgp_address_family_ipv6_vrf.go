// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPAddressFamilyIPv6VRF struct {
	Device types.String                  `tfsdk:"device"`
	Id     types.String                  `tfsdk:"id"`
	Asn    types.String                  `tfsdk:"asn"`
	AfName types.String                  `tfsdk:"af_name"`
	Vrfs   []BGPAddressFamilyIPv6VRFVrfs `tfsdk:"vrfs"`
}
type BGPAddressFamilyIPv6VRFVrfs struct {
	Name                  types.String `tfsdk:"name"`
	AdvertiseL2vpnEvpn    types.Bool   `tfsdk:"advertise_l2vpn_evpn"`
	RedistributeConnected types.Bool   `tfsdk:"redistribute_connected"`
	RedistributeStatic    types.Bool   `tfsdk:"redistribute_static"`
}

func (data BGPAddressFamilyIPv6VRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv6=%s", data.Asn.Value, data.AfName.Value)
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyIPv6VRF) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyIPv6VRF) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.Null && !data.AfName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.Value)
	}
	if len(data.Vrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf", []interface{}{})
		for index, item := range data.Vrfs {
			if !item.Name.Null && !item.Name.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"name", item.Name.Value)
			}
			if !item.AdvertiseL2vpnEvpn.Null && !item.AdvertiseL2vpnEvpn.Unknown {
				if item.AdvertiseL2vpnEvpn.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.advertise.l2vpn.evpn", map[string]string{})
				}
			}
			if !item.RedistributeConnected.Null && !item.RedistributeConnected.Unknown {
				if item.RedistributeConnected.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.redistribute-v6.connected", map[string]string{})
				}
			}
			if !item.RedistributeStatic.Null && !item.RedistributeStatic.Unknown {
				if item.RedistributeStatic.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.redistribute-v6.static", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *BGPAddressFamilyIPv6VRF) updateFromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "af-name"); value.Exists() {
		data.AfName.Value = value.String()
	} else {
		data.AfName.Null = true
	}
	for i := range data.Vrfs {
		key := data.Vrfs[i].Name.Value
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.#(name==\"" + key + "\")." + "name"); value.Exists() {
			data.Vrfs[i].Name.Value = value.String()
		} else {
			data.Vrfs[i].Name.Null = true
		}
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.#(name==\"" + key + "\")." + "ipv6-unicast.advertise.l2vpn.evpn"); value.Exists() {
			data.Vrfs[i].AdvertiseL2vpnEvpn.Value = true
		} else {
			data.Vrfs[i].AdvertiseL2vpnEvpn.Value = false
		}
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.#(name==\"" + key + "\")." + "ipv6-unicast.redistribute-v6.connected"); value.Exists() {
			data.Vrfs[i].RedistributeConnected.Value = true
		} else {
			data.Vrfs[i].RedistributeConnected.Value = false
		}
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.#(name==\"" + key + "\")." + "ipv6-unicast.redistribute-v6.static"); value.Exists() {
			data.Vrfs[i].RedistributeStatic.Value = true
		} else {
			data.Vrfs[i].RedistributeStatic.Value = false
		}
	}
}

func (data *BGPAddressFamilyIPv6VRF) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf"); value.Exists() {
		data.Vrfs = make([]BGPAddressFamilyIPv6VRFVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv6VRFVrfs{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name.Value = cValue.String()
			}
			if cValue := v.Get("ipv6-unicast.advertise.l2vpn.evpn"); cValue.Exists() {
				item.AdvertiseL2vpnEvpn.Value = true
			}
			if cValue := v.Get("ipv6-unicast.redistribute-v6.connected"); cValue.Exists() {
				item.RedistributeConnected.Value = true
			}
			if cValue := v.Get("ipv6-unicast.redistribute-v6.static"); cValue.Exists() {
				item.RedistributeStatic.Value = true
			}
			data.Vrfs = append(data.Vrfs, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv6VRF) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Asn.Unknown {
		data.Asn.Unknown = false
		data.Asn.Null = true
	}
	if data.AfName.Unknown {
		data.AfName.Unknown = false
		data.AfName.Null = true
	}
	for i := range data.Vrfs {
		if data.Vrfs[i].Name.Unknown {
			data.Vrfs[i].Name.Unknown = false
			data.Vrfs[i].Name.Null = true
		}
		if data.Vrfs[i].AdvertiseL2vpnEvpn.Unknown {
			data.Vrfs[i].AdvertiseL2vpnEvpn.Unknown = false
			data.Vrfs[i].AdvertiseL2vpnEvpn.Null = true
		}
		if data.Vrfs[i].RedistributeConnected.Unknown {
			data.Vrfs[i].RedistributeConnected.Unknown = false
			data.Vrfs[i].RedistributeConnected.Null = true
		}
		if data.Vrfs[i].RedistributeStatic.Unknown {
			data.Vrfs[i].RedistributeStatic.Unknown = false
			data.Vrfs[i].RedistributeStatic.Null = true
		}
	}
}

func (data *BGPAddressFamilyIPv6VRF) getDeletedListItems(state BGPAddressFamilyIPv6VRF) []string {
	deletedListItems := make([]string, 0)
	for _, i := range state.Vrfs {
		if reflect.ValueOf(i.Name.Value).IsZero() {
			continue
		}
		found := false
		for _, j := range data.Vrfs {
			if i.Name.Value == j.Name.Value {
				found = true
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, state.getPath()+"/vrf="+i.Name.Value)
		}
	}
	return deletedListItems
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfaceNVE struct {
	Device                      types.String          `tfsdk:"device"`
	Id                          types.String          `tfsdk:"id"`
	Name                        types.Int64           `tfsdk:"name"`
	Description                 types.String          `tfsdk:"description"`
	Shutdown                    types.Bool            `tfsdk:"shutdown"`
	HostReachabilityProtocolBgp types.Bool            `tfsdk:"host_reachability_protocol_bgp"`
	SourceInterfaceLoopback     types.Int64           `tfsdk:"source_interface_loopback"`
	VniVrfs                     []InterfaceNVEVniVrfs `tfsdk:"vni_vrfs"`
	Vnis                        []InterfaceNVEVnis    `tfsdk:"vnis"`
}
type InterfaceNVEVniVrfs struct {
	VniRange types.String `tfsdk:"vni_range"`
	Vrf      types.String `tfsdk:"vrf"`
}
type InterfaceNVEVnis struct {
	VniRange           types.String `tfsdk:"vni_range"`
	Ipv4MulticastGroup types.String `tfsdk:"ipv4_multicast_group"`
	IngressReplication types.Bool   `tfsdk:"ingress_replication"`
}

func (data InterfaceNVE) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/nve=%v", data.Name.Value)
}

func (data InterfaceNVE) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.Value, 10))
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.HostReachabilityProtocolBgp.Null && !data.HostReachabilityProtocolBgp.Unknown {
		if data.HostReachabilityProtocolBgp.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host-reachability.protocol.bgp", map[string]string{})
		}
	}
	if !data.SourceInterfaceLoopback.Null && !data.SourceInterfaceLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface.Loopback", strconv.FormatInt(data.SourceInterfaceLoopback.Value, 10))
	}
	if len(data.VniVrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni", []interface{}{})
		for index, item := range data.VniVrfs {
			if !item.VniRange.Null && !item.VniRange.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
			}
			if !item.Vrf.Null && !item.Vrf.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.Value)
			}
		}
	}
	if len(data.Vnis) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", []interface{}{})
		for index, item := range data.Vnis {
			if !item.VniRange.Null && !item.VniRange.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
			}
			if !item.Ipv4MulticastGroup.Null && !item.Ipv4MulticastGroup.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"mcast-group.multicast-group-min", item.Ipv4MulticastGroup.Value)
			}
			if !item.IngressReplication.Null && !item.IngressReplication.Unknown {
				if item.IngressReplication.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"ir-cp-config.ingress-replication", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *InterfaceNVE) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "host-reachability.protocol.bgp"); value.Exists() {
		data.HostReachabilityProtocolBgp.Value = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "source-interface.Loopback"); value.Exists() {
		data.SourceInterfaceLoopback.Value = value.Int()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member-in-one-line.member.vni"); value.Exists() {
		data.VniVrfs = make([]InterfaceNVEVniVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVniVrfs{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf.Value = cValue.String()
			}
			data.VniVrfs = append(data.VniVrfs, item)
			return true
		})
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.vni"); value.Exists() {
		data.Vnis = make([]InterfaceNVEVnis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVnis{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
			}
			if cValue := v.Get("mcast-group.multicast-group-min"); cValue.Exists() {
				item.Ipv4MulticastGroup.Value = cValue.String()
			}
			if cValue := v.Get("ir-cp-config.ingress-replication"); cValue.Exists() {
				item.IngressReplication.Value = true
			}
			data.Vnis = append(data.Vnis, item)
			return true
		})
	}
}

func (data *InterfaceNVE) fromPlan(plan InterfaceNVE) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
	sort.SliceStable(data.VniVrfs, func(i, j int) bool {
		for ii := range plan.VniVrfs {
			if plan.VniVrfs[ii].VniRange.Value == data.VniVrfs[i].VniRange.Value {
				return true
			}
			if plan.VniVrfs[ii].VniRange.Value == data.VniVrfs[j].VniRange.Value {
				return false
			}
		}
		return false
	})
	sort.SliceStable(data.Vnis, func(i, j int) bool {
		for ii := range plan.Vnis {
			if plan.Vnis[ii].VniRange.Value == data.Vnis[i].VniRange.Value {
				return true
			}
			if plan.Vnis[ii].VniRange.Value == data.Vnis[j].VniRange.Value {
				return false
			}
		}
		return false
	})
}

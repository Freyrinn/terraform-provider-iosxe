// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfaceNVE struct {
	Device                                         types.String                           `tfsdk:"device"`
	Id                                             types.String                           `tfsdk:"id"`
	Name                                           types.Int64                            `tfsdk:"name"`
	Description                                    types.String                           `tfsdk:"description"`
	Shutdown                                       types.Bool                             `tfsdk:"shutdown"`
	HostReachabilityProtocolBgp                    types.Bool                             `tfsdk:"host_reachability_protocol_bgp"`
	SourceInterfaceInterfaceChoiceLoopbackLoopback types.Int64                            `tfsdk:"source_interface_loopback"`
	MemberInOneLineMemberVni                       []InterfaceNVEMemberInOneLineMemberVni `tfsdk:"vni_vrfs"`
	MemberVni                                      []InterfaceNVEMemberVni                `tfsdk:"vnis"`
}
type InterfaceNVEMemberInOneLineMemberVni struct {
	VniRange types.String `tfsdk:"vni_range"`
	Vrf      types.String `tfsdk:"vrf"`
}
type InterfaceNVEMemberVni struct {
	VniRange                     types.String `tfsdk:"vni_range"`
	McastGroupMulticastGroupMin  types.String `tfsdk:"ipv4_multicast_group"`
	IrCpConfigIngressReplication types.Bool   `tfsdk:"ingress_replication"`
}

func (data InterfaceNVE) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/nve=%v", data.Name.Value)
}

func (data InterfaceNVE) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni", []interface{}{})
	for index, item := range data.MemberInOneLineMemberVni {
		if !item.VniRange.Null && !item.VniRange.Unknown {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
		}
		if !item.Vrf.Null && !item.Vrf.Unknown {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.Value)
		}
	}
	body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", []interface{}{})
	for index, item := range data.MemberVni {
		if !item.VniRange.Null && !item.VniRange.Unknown {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.Value)
		}
		if !item.McastGroupMulticastGroupMin.Null && !item.McastGroupMulticastGroupMin.Unknown {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"mcast-group.multicast-group-min", item.McastGroupMulticastGroupMin.Value)
		}
		if !item.IrCpConfigIngressReplication.Null && !item.IrCpConfigIngressReplication.Unknown {
			if item.IrCpConfigIngressReplication.Value {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"ir-cp-config.ingress-replication", map[string]string{})
			}
		}
	}
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
	if !data.SourceInterfaceInterfaceChoiceLoopbackLoopback.Null && !data.SourceInterfaceInterfaceChoiceLoopbackLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface.Loopback", strconv.FormatInt(data.SourceInterfaceInterfaceChoiceLoopbackLoopback.Value, 10))
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
		data.SourceInterfaceInterfaceChoiceLoopbackLoopback.Value = value.Int()
	}
	data.MemberInOneLineMemberVni = make([]InterfaceNVEMemberInOneLineMemberVni, 0)
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member-in-one-line.member.vni"); value.Exists() {
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEMemberInOneLineMemberVni{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf.Value = cValue.String()
			}
			data.MemberInOneLineMemberVni = append(data.MemberInOneLineMemberVni, item)
			return true
		})
	}
	data.MemberVni = make([]InterfaceNVEMemberVni, 0)
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.vni"); value.Exists() {
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEMemberVni{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange.Value = cValue.String()
			}
			if cValue := v.Get("mcast-group.multicast-group-min"); cValue.Exists() {
				item.McastGroupMulticastGroupMin.Value = cValue.String()
			}
			if cValue := v.Get("ir-cp-config.ingress-replication"); cValue.Exists() {
				item.IrCpConfigIngressReplication.Value = true
			}
			data.MemberVni = append(data.MemberVni, item)
			return true
		})
	}
}

func (data *InterfaceNVE) fromPlan(plan InterfaceNVE) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
}

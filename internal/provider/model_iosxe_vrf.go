// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VRF struct {
	Device            types.String           `tfsdk:"device"`
	Id                types.String           `tfsdk:"id"`
	Name              types.String           `tfsdk:"name"`
	Description       types.String           `tfsdk:"description"`
	Rd                types.String           `tfsdk:"rd"`
	AddressFamilyIpv4 types.Bool             `tfsdk:"address_family_ipv4"`
	AddressFamilyIpv6 types.Bool             `tfsdk:"address_family_ipv6"`
	VpnId             types.String           `tfsdk:"vpn_id"`
	RouteTargetImport []VRFRouteTargetImport `tfsdk:"route_target_import"`
	RouteTargetExport []VRFRouteTargetExport `tfsdk:"route_target_export"`
}
type VRFRouteTargetImport struct {
	Value     types.String `tfsdk:"value"`
	Stitching types.Bool   `tfsdk:"stitching"`
}
type VRFRouteTargetExport struct {
	Value     types.String `tfsdk:"value"`
	Stitching types.Bool   `tfsdk:"stitching"`
}

func (data VRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vrf/definition=%s", data.Name.Value)
}

// if last path element has a key -> remove it
func (data VRF) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data VRF) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Rd.Null && !data.Rd.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"rd", data.Rd.Value)
	}
	if !data.AddressFamilyIpv4.Null && !data.AddressFamilyIpv4.Unknown {
		if data.AddressFamilyIpv4.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address-family.ipv4", map[string]string{})
		}
	}
	if !data.AddressFamilyIpv6.Null && !data.AddressFamilyIpv6.Unknown {
		if data.AddressFamilyIpv6.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"address-family.ipv6", map[string]string{})
		}
	}
	if !data.VpnId.Null && !data.VpnId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vpn.id", data.VpnId.Value)
	}
	if len(data.RouteTargetImport) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.import", []interface{}{})
		for index, item := range data.RouteTargetImport {
			if !item.Value.Null && !item.Value.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.import"+"."+strconv.Itoa(index)+"."+"asn-ip", item.Value.Value)
			}
			if !item.Stitching.Null && !item.Stitching.Unknown {
				if item.Stitching.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.import"+"."+strconv.Itoa(index)+"."+"stitching", map[string]string{})
				}
			}
		}
	}
	if len(data.RouteTargetExport) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.export", []interface{}{})
		for index, item := range data.RouteTargetExport {
			if !item.Value.Null && !item.Value.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.export"+"."+strconv.Itoa(index)+"."+"asn-ip", item.Value.Value)
			}
			if !item.Stitching.Null && !item.Stitching.Unknown {
				if item.Stitching.Value {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.export"+"."+strconv.Itoa(index)+"."+"stitching", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *VRF) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "rd"); value.Exists() {
		data.Rd.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "address-family.ipv4"); value.Exists() {
		data.AddressFamilyIpv4.Value = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "address-family.ipv6"); value.Exists() {
		data.AddressFamilyIpv6.Value = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vpn.id"); value.Exists() {
		data.VpnId.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.import"); value.Exists() {
		data.RouteTargetImport = make([]VRFRouteTargetImport, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VRFRouteTargetImport{}
			if cValue := v.Get("asn-ip"); cValue.Exists() {
				item.Value.Value = cValue.String()
			}
			if cValue := v.Get("stitching"); cValue.Exists() {
				item.Stitching.Value = true
			}
			data.RouteTargetImport = append(data.RouteTargetImport, item)
			return true
		})
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.export"); value.Exists() {
		data.RouteTargetExport = make([]VRFRouteTargetExport, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VRFRouteTargetExport{}
			if cValue := v.Get("asn-ip"); cValue.Exists() {
				item.Value.Value = cValue.String()
			}
			if cValue := v.Get("stitching"); cValue.Exists() {
				item.Stitching.Value = true
			}
			data.RouteTargetExport = append(data.RouteTargetExport, item)
			return true
		})
	}
}

func (data *VRF) fromPlan(plan VRF) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
	sort.SliceStable(data.RouteTargetImport, func(i, j int) bool {
		for ii := range plan.RouteTargetImport {
			if plan.RouteTargetImport[ii].Value.Value == data.RouteTargetImport[i].Value.Value {
				return true
			}
			if plan.RouteTargetImport[ii].Value.Value == data.RouteTargetImport[j].Value.Value {
				return false
			}
		}
		return false
	})
	sort.SliceStable(data.RouteTargetExport, func(i, j int) bool {
		for ii := range plan.RouteTargetExport {
			if plan.RouteTargetExport[ii].Value.Value == data.RouteTargetExport[i].Value.Value {
				return true
			}
			if plan.RouteTargetExport[ii].Value.Value == data.RouteTargetExport[j].Value.Value {
				return false
			}
		}
		return false
	})
}

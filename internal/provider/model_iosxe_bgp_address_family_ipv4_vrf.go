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

type BGPAddressFamilyIPv4VRF struct {
	Device types.String                 `tfsdk:"device"`
	Id     types.String                 `tfsdk:"id"`
	YangId types.String                 `tfsdk:"asn"`
	AfName types.String                 `tfsdk:"af_name"`
	Vrf    []BGPAddressFamilyIPv4VRFVrf `tfsdk:"vrfs"`
}
type BGPAddressFamilyIPv4VRFVrf struct {
	Name types.String `tfsdk:"name"`
}

func (data BGPAddressFamilyIPv4VRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv4=%s", data.YangId.Value, data.AfName.Value)
}

func (data BGPAddressFamilyIPv4VRF) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.Null && !data.AfName.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.Value)
	}
	if len(data.Vrf) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf", []interface{}{})
		for index, item := range data.Vrf {
			if !item.Name.Null && !item.Name.Unknown {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"name", item.Name.Value)
			}
		}
	}
	return body
}

func (data *BGPAddressFamilyIPv4VRF) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf"); value.Exists() {
		data.Vrf = make([]BGPAddressFamilyIPv4VRFVrf, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv4VRFVrf{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name.Value = cValue.String()
			}
			data.Vrf = append(data.Vrf, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv4VRF) fromPlan(plan BGPAddressFamilyIPv4VRF) {
	data.Device = plan.Device
	data.YangId.Value = plan.YangId.Value
	data.AfName.Value = plan.AfName.Value
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"regexp"
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

func (data *VRF) updateFromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "name"); value.Exists() {
		data.Name.Value = value.String()
	} else {
		data.Name.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "rd"); value.Exists() {
		data.Rd.Value = value.String()
	} else {
		data.Rd.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "address-family.ipv4"); value.Exists() {
		data.AddressFamilyIpv4.Value = true
	} else {
		data.AddressFamilyIpv4.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "address-family.ipv6"); value.Exists() {
		data.AddressFamilyIpv6.Value = true
	} else {
		data.AddressFamilyIpv6.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vpn.id"); value.Exists() {
		data.VpnId.Value = value.String()
	} else {
		data.VpnId.Null = true
	}
	for i := range data.RouteTargetImport {
		key := data.RouteTargetImport[i].Value.Value
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.import.#(asn-ip==\"" + key + "\")." + "asn-ip"); value.Exists() {
			data.RouteTargetImport[i].Value.Value = value.String()
		} else {
			data.RouteTargetImport[i].Value.Null = true
		}
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.import.#(asn-ip==\"" + key + "\")." + "stitching"); value.Exists() {
			data.RouteTargetImport[i].Stitching.Value = true
		} else {
			data.RouteTargetImport[i].Stitching.Value = false
		}
	}
	for i := range data.RouteTargetExport {
		key := data.RouteTargetExport[i].Value.Value
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.export.#(asn-ip==\"" + key + "\")." + "asn-ip"); value.Exists() {
			data.RouteTargetExport[i].Value.Value = value.String()
		} else {
			data.RouteTargetExport[i].Value.Null = true
		}
		if value := res.Get(helpers.LastElement(data.getPath()) + "." + "route-target.export.#(asn-ip==\"" + key + "\")." + "stitching"); value.Exists() {
			data.RouteTargetExport[i].Stitching.Value = true
		} else {
			data.RouteTargetExport[i].Stitching.Value = false
		}
	}
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

func (data *VRF) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Name.Unknown {
		data.Name.Unknown = false
		data.Name.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
	if data.Rd.Unknown {
		data.Rd.Unknown = false
		data.Rd.Null = true
	}
	if data.AddressFamilyIpv4.Unknown {
		data.AddressFamilyIpv4.Unknown = false
		data.AddressFamilyIpv4.Null = true
	}
	if data.AddressFamilyIpv6.Unknown {
		data.AddressFamilyIpv6.Unknown = false
		data.AddressFamilyIpv6.Null = true
	}
	if data.VpnId.Unknown {
		data.VpnId.Unknown = false
		data.VpnId.Null = true
	}
	for i := range data.RouteTargetImport {
		if data.RouteTargetImport[i].Value.Unknown {
			data.RouteTargetImport[i].Value.Unknown = false
			data.RouteTargetImport[i].Value.Null = true
		}
		if data.RouteTargetImport[i].Stitching.Unknown {
			data.RouteTargetImport[i].Stitching.Unknown = false
			data.RouteTargetImport[i].Stitching.Null = true
		}
	}
	for i := range data.RouteTargetExport {
		if data.RouteTargetExport[i].Value.Unknown {
			data.RouteTargetExport[i].Value.Unknown = false
			data.RouteTargetExport[i].Value.Null = true
		}
		if data.RouteTargetExport[i].Stitching.Unknown {
			data.RouteTargetExport[i].Stitching.Unknown = false
			data.RouteTargetExport[i].Stitching.Null = true
		}
	}
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPL2VPNEVPNNeighbor struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	Activate             types.Bool   `tfsdk:"activate"`
	SendCommunity        types.String `tfsdk:"send_community"`
	RouteReflectorClient types.Bool   `tfsdk:"route_reflector_client"`
}

func (data BGPL2VPNEVPNNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/l2vpn=evpn/l2vpn-evpn/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.Value)), url.QueryEscape(fmt.Sprintf("%v", data.Ip.Value)))
}

// if last path element has a key -> remove it
func (data BGPL2VPNEVPNNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPL2VPNEVPNNeighbor) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.Null && !data.Ip.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.Value)
	}
	if !data.Activate.Null && !data.Activate.Unknown {
		if data.Activate.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"activate", map[string]string{})
		}
	}
	if !data.SendCommunity.Null && !data.SendCommunity.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"send-community.send-community-where", data.SendCommunity.Value)
	}
	if !data.RouteReflectorClient.Null && !data.RouteReflectorClient.Unknown {
		if data.RouteReflectorClient.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-reflector-client", map[string]string{})
		}
	}
	return body
}

func (data *BGPL2VPNEVPNNeighbor) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.Ip.Value = value.String()
	} else {
		data.Ip.Null = true
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate.Value = true
	} else {
		data.Activate.Value = false
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity.Value = value.String()
	} else {
		data.SendCommunity.Null = true
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient.Value = true
	} else {
		data.RouteReflectorClient.Value = false
	}
}

func (data *BGPL2VPNEVPNNeighbor) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate.Value = true
		data.Activate.Null = false
	} else {
		data.Activate.Value = false
		data.Activate.Null = false
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity.Value = value.String()
		data.SendCommunity.Null = false
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient.Value = true
		data.RouteReflectorClient.Null = false
	} else {
		data.RouteReflectorClient.Value = false
		data.RouteReflectorClient.Null = false
	}
}

func (data *BGPL2VPNEVPNNeighbor) setUnknownValues() {
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
	if data.Ip.Unknown {
		data.Ip.Unknown = false
		data.Ip.Null = true
	}
	if data.Activate.Unknown {
		data.Activate.Unknown = false
		data.Activate.Null = true
	}
	if data.SendCommunity.Unknown {
		data.SendCommunity.Unknown = false
		data.SendCommunity.Null = true
	}
	if data.RouteReflectorClient.Unknown {
		data.RouteReflectorClient.Unknown = false
		data.RouteReflectorClient.Null = true
	}
}

func (data *BGPL2VPNEVPNNeighbor) getDeletedListItems(state BGPL2VPNEVPNNeighbor) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

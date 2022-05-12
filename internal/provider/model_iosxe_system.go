// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type System struct {
	Device             types.String `tfsdk:"device"`
	Id                 types.String `tfsdk:"id"`
	Hostname           types.String `tfsdk:"hostname"`
	Ipv6UnicastRouting types.Bool   `tfsdk:"ipv6_unicast_routing"`
}

func (data System) getPath() string {
	return "Cisco-IOS-XE-native:native"
}

// if last path element has a key -> remove it
func (data System) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data System) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Hostname.Null && !data.Hostname.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"hostname", data.Hostname.Value)
	}
	if !data.Ipv6UnicastRouting.Null && !data.Ipv6UnicastRouting.Unknown {
		if data.Ipv6UnicastRouting.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.unicast-routing", map[string]string{})
		}
	}
	return body
}

func (data *System) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "hostname"); value.Exists() {
		data.Hostname.Value = value.String()
	} else {
		data.Hostname.Null = true
	}
	if value := res.Get(prefix + "ipv6.unicast-routing"); value.Exists() {
		data.Ipv6UnicastRouting.Value = true
	} else {
		data.Ipv6UnicastRouting.Value = false
	}
}

func (data *System) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "hostname"); value.Exists() {
		data.Hostname.Value = value.String()
		data.Hostname.Null = false
	}
	if value := res.Get(prefix + "ipv6.unicast-routing"); value.Exists() {
		data.Ipv6UnicastRouting.Value = true
		data.Ipv6UnicastRouting.Null = false
	} else {
		data.Ipv6UnicastRouting.Value = false
		data.Ipv6UnicastRouting.Null = false
	}
}

func (data *System) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Hostname.Unknown {
		data.Hostname.Unknown = false
		data.Hostname.Null = true
	}
	if data.Ipv6UnicastRouting.Unknown {
		data.Ipv6UnicastRouting.Unknown = false
		data.Ipv6UnicastRouting.Null = true
	}
}

func (data *System) getDeletedListItems(state System) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

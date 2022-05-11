// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfaceEthernet struct {
	Device                   types.String `tfsdk:"device"`
	Id                       types.String `tfsdk:"id"`
	Type                     types.String `tfsdk:"type"`
	Name                     types.String `tfsdk:"name"`
	MediaType                types.String `tfsdk:"media_type"`
	Switchport               types.Bool   `tfsdk:"switchport"`
	Description              types.String `tfsdk:"description"`
	Shutdown                 types.Bool   `tfsdk:"shutdown"`
	VrfForwarding            types.String `tfsdk:"vrf_forwarding"`
	Ipv4Address              types.String `tfsdk:"ipv4_address"`
	Ipv4AddressMask          types.String `tfsdk:"ipv4_address_mask"`
	Unnumbered               types.String `tfsdk:"unnumbered"`
	EncapsulationDot1qVlanId types.Int64  `tfsdk:"encapsulation_dot1q_vlan_id"`
}

func (data InterfaceEthernet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v", url.QueryEscape(fmt.Sprintf("%v", data.Type.Value)), url.QueryEscape(fmt.Sprintf("%v", data.Name.Value)))
}

// if last path element has a key -> remove it
func (data InterfaceEthernet) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceEthernet) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.Null && !data.Name.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.Value)
	}
	if !data.MediaType.Null && !data.MediaType.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"media-type", data.MediaType.Value)
	}
	if !data.Switchport.Null && !data.Switchport.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"switchport-conf.switchport", data.Switchport.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.VrfForwarding.Null && !data.VrfForwarding.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf.forwarding", data.VrfForwarding.Value)
	}
	if !data.Ipv4Address.Null && !data.Ipv4Address.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.address", data.Ipv4Address.Value)
	}
	if !data.Ipv4AddressMask.Null && !data.Ipv4AddressMask.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.mask", data.Ipv4AddressMask.Value)
	}
	if !data.Unnumbered.Null && !data.Unnumbered.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.unnumbered", data.Unnumbered.Value)
	}
	if !data.EncapsulationDot1qVlanId.Null && !data.EncapsulationDot1qVlanId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"encapsulation.dot1Q.vlan-id", strconv.FormatInt(data.EncapsulationDot1qVlanId.Value, 10))
	}
	return body
}

func (data *InterfaceEthernet) updateFromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "name"); value.Exists() {
		data.Name.Value = value.String()
	} else {
		data.Name.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "media-type"); value.Exists() {
		data.MediaType.Value = value.String()
	} else {
		data.MediaType.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "switchport-conf.switchport"); value.Exists() {
		data.Switchport.Value = value.Bool()
	} else {
		data.Switchport.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	} else {
		data.Shutdown.Value = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding.Value = value.String()
	} else {
		data.VrfForwarding.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address.Value = value.String()
	} else {
		data.Ipv4Address.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask.Value = value.String()
	} else {
		data.Ipv4AddressMask.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.unnumbered"); value.Exists() {
		data.Unnumbered.Value = value.String()
	} else {
		data.Unnumbered.Null = true
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId.Value = value.Int()
	} else {
		data.EncapsulationDot1qVlanId.Null = true
	}
}

func (data *InterfaceEthernet) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "media-type"); value.Exists() {
		data.MediaType.Value = value.String()
		data.MediaType.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "switchport-conf.switchport"); value.Exists() {
		data.Switchport.Value = value.Bool()
		data.Switchport.Null = false
	} else {
		data.Switchport.Value = false
		data.Switchport.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "description"); value.Exists() {
		data.Description.Value = value.String()
		data.Description.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding.Value = value.String()
		data.VrfForwarding.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address.Value = value.String()
		data.Ipv4Address.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask.Value = value.String()
		data.Ipv4AddressMask.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "ip.unnumbered"); value.Exists() {
		data.Unnumbered.Value = value.String()
		data.Unnumbered.Null = false
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "encapsulation.dot1Q.vlan-id"); value.Exists() {
		data.EncapsulationDot1qVlanId.Value = value.Int()
		data.EncapsulationDot1qVlanId.Null = false
	}
}

func (data *InterfaceEthernet) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.Type.Unknown {
		data.Type.Unknown = false
		data.Type.Null = true
	}
	if data.Name.Unknown {
		data.Name.Unknown = false
		data.Name.Null = true
	}
	if data.MediaType.Unknown {
		data.MediaType.Unknown = false
		data.MediaType.Null = true
	}
	if data.Switchport.Unknown {
		data.Switchport.Unknown = false
		data.Switchport.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
	if data.VrfForwarding.Unknown {
		data.VrfForwarding.Unknown = false
		data.VrfForwarding.Null = true
	}
	if data.Ipv4Address.Unknown {
		data.Ipv4Address.Unknown = false
		data.Ipv4Address.Null = true
	}
	if data.Ipv4AddressMask.Unknown {
		data.Ipv4AddressMask.Unknown = false
		data.Ipv4AddressMask.Null = true
	}
	if data.Unnumbered.Unknown {
		data.Unnumbered.Unknown = false
		data.Unnumbered.Null = true
	}
	if data.EncapsulationDot1qVlanId.Unknown {
		data.EncapsulationDot1qVlanId.Unknown = false
		data.EncapsulationDot1qVlanId.Null = true
	}
}

func (data *InterfaceEthernet) getDeletedListItems(state InterfaceEthernet) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

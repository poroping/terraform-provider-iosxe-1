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

type VLANConfiguration struct {
	Device          types.String `tfsdk:"device"`
	Id              types.String `tfsdk:"id"`
	VlanId          types.String `tfsdk:"vlan_id"`
	Vni             types.Int64  `tfsdk:"vni"`
	AccessVfi       types.String `tfsdk:"access_vfi"`
	EvpnInstance    types.Int64  `tfsdk:"evpn_instance"`
	EvpnInstanceVni types.Int64  `tfsdk:"evpn_instance_vni"`
}

func (data VLANConfiguration) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:configuration-entry=%v", data.VlanId.Value)
}

func (data VLANConfiguration) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.VlanId.Null && !data.VlanId.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-id", data.VlanId.Value)
	}
	if !data.Vni.Null && !data.Vni.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", strconv.FormatInt(data.Vni.Value, 10))
	}
	if !data.AccessVfi.Null && !data.AccessVfi.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.access-vfi", data.AccessVfi.Value)
	}
	if !data.EvpnInstance.Null && !data.EvpnInstance.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.evpn-instance", strconv.FormatInt(data.EvpnInstance.Value, 10))
	}
	if !data.EvpnInstanceVni.Null && !data.EvpnInstanceVni.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.evpn-instance.vni", strconv.FormatInt(data.EvpnInstanceVni.Value, 10))
	}
	return body
}

func (data *VLANConfiguration) fromBody(res gjson.Result) {
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.vni"); value.Exists() {
		data.Vni.Value = value.Int()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.access-vfi"); value.Exists() {
		data.AccessVfi.Value = value.String()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.evpn-instance"); value.Exists() {
		data.EvpnInstance.Value = value.Int()
	}
	if value := res.Get(helpers.LastElement(data.getPath()) + "." + "member.evpn-instance.vni"); value.Exists() {
		data.EvpnInstanceVni.Value = value.Int()
	}
}

func (data *VLANConfiguration) fromPlan(plan VLANConfiguration) {
	data.Device = plan.Device
	data.VlanId.Value = plan.VlanId.Value
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type EVPN struct {
	Device                  types.String `tfsdk:"device"`
	Id                      types.String `tfsdk:"id"`
	ReplicationTypeIngress  types.Bool   `tfsdk:"replication_type_ingress"`
	ReplicationTypeStatic   types.Bool   `tfsdk:"replication_type_static"`
	ReplicationTypeP2mp     types.Bool   `tfsdk:"replication_type_p2mp"`
	ReplicationTypeMp2mp    types.Bool   `tfsdk:"replication_type_mp2mp"`
	MacDuplicationLimit     types.Int64  `tfsdk:"mac_duplication_limit"`
	MacDuplicationTime      types.Int64  `tfsdk:"mac_duplication_time"`
	IpDuplicationLimit      types.Int64  `tfsdk:"ip_duplication_limit"`
	IpDuplicationTime       types.Int64  `tfsdk:"ip_duplication_time"`
	RouterIdLoopback        types.Int64  `tfsdk:"router_id_loopback"`
	DefaultGatewayAdvertise types.Bool   `tfsdk:"default_gateway_advertise"`
	LoggingPeerState        types.Bool   `tfsdk:"logging_peer_state"`
	RouteTargetAutoVni      types.Bool   `tfsdk:"route_target_auto_vni"`
}

func (data EVPN) getPath() string {
	return "Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn"
}

// if last path element has a key -> remove it
func (data EVPN) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data EVPN) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.ReplicationTypeIngress.Null && !data.ReplicationTypeIngress.Unknown {
		if data.ReplicationTypeIngress.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"replication-type.ingress", map[string]string{})
		}
	}
	if !data.ReplicationTypeStatic.Null && !data.ReplicationTypeStatic.Unknown {
		if data.ReplicationTypeStatic.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"replication-type.static", map[string]string{})
		}
	}
	if !data.ReplicationTypeP2mp.Null && !data.ReplicationTypeP2mp.Unknown {
		if data.ReplicationTypeP2mp.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"replication-type.p2mp", map[string]string{})
		}
	}
	if !data.ReplicationTypeMp2mp.Null && !data.ReplicationTypeMp2mp.Unknown {
		if data.ReplicationTypeMp2mp.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"replication-type.mp2mp", map[string]string{})
		}
	}
	if !data.MacDuplicationLimit.Null && !data.MacDuplicationLimit.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mac.duplication.limit", strconv.FormatInt(data.MacDuplicationLimit.Value, 10))
	}
	if !data.MacDuplicationTime.Null && !data.MacDuplicationTime.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"mac.duplication.time", strconv.FormatInt(data.MacDuplicationTime.Value, 10))
	}
	if !data.IpDuplicationLimit.Null && !data.IpDuplicationLimit.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.duplication.limit", strconv.FormatInt(data.IpDuplicationLimit.Value, 10))
	}
	if !data.IpDuplicationTime.Null && !data.IpDuplicationTime.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.duplication.time", strconv.FormatInt(data.IpDuplicationTime.Value, 10))
	}
	if !data.RouterIdLoopback.Null && !data.RouterIdLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"router-id.interface.Loopback", strconv.FormatInt(data.RouterIdLoopback.Value, 10))
	}
	if !data.DefaultGatewayAdvertise.Null && !data.DefaultGatewayAdvertise.Unknown {
		if data.DefaultGatewayAdvertise.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"default-gateway.advertise", map[string]string{})
		}
	}
	if !data.LoggingPeerState.Null && !data.LoggingPeerState.Unknown {
		if data.LoggingPeerState.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"logging.peer.state", map[string]string{})
		}
	}
	if !data.RouteTargetAutoVni.Null && !data.RouteTargetAutoVni.Unknown {
		if data.RouteTargetAutoVni.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-target.auto.vni", map[string]string{})
		}
	}
	return body
}

func (data *EVPN) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "replication-type.ingress"); value.Exists() {
		data.ReplicationTypeIngress.Value = true
	} else {
		data.ReplicationTypeIngress.Value = false
	}
	if value := res.Get(prefix + "replication-type.static"); value.Exists() {
		data.ReplicationTypeStatic.Value = true
	} else {
		data.ReplicationTypeStatic.Value = false
	}
	if value := res.Get(prefix + "replication-type.p2mp"); value.Exists() {
		data.ReplicationTypeP2mp.Value = true
	} else {
		data.ReplicationTypeP2mp.Value = false
	}
	if value := res.Get(prefix + "replication-type.mp2mp"); value.Exists() {
		data.ReplicationTypeMp2mp.Value = true
	} else {
		data.ReplicationTypeMp2mp.Value = false
	}
	if value := res.Get(prefix + "mac.duplication.limit"); value.Exists() {
		data.MacDuplicationLimit.Value = value.Int()
	} else {
		data.MacDuplicationLimit.Null = true
	}
	if value := res.Get(prefix + "mac.duplication.time"); value.Exists() {
		data.MacDuplicationTime.Value = value.Int()
	} else {
		data.MacDuplicationTime.Null = true
	}
	if value := res.Get(prefix + "ip.duplication.limit"); value.Exists() {
		data.IpDuplicationLimit.Value = value.Int()
	} else {
		data.IpDuplicationLimit.Null = true
	}
	if value := res.Get(prefix + "ip.duplication.time"); value.Exists() {
		data.IpDuplicationTime.Value = value.Int()
	} else {
		data.IpDuplicationTime.Null = true
	}
	if value := res.Get(prefix + "router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback.Value = value.Int()
	} else {
		data.RouterIdLoopback.Null = true
	}
	if value := res.Get(prefix + "default-gateway.advertise"); value.Exists() {
		data.DefaultGatewayAdvertise.Value = true
	} else {
		data.DefaultGatewayAdvertise.Value = false
	}
	if value := res.Get(prefix + "logging.peer.state"); value.Exists() {
		data.LoggingPeerState.Value = true
	} else {
		data.LoggingPeerState.Value = false
	}
	if value := res.Get(prefix + "route-target.auto.vni"); value.Exists() {
		data.RouteTargetAutoVni.Value = true
	} else {
		data.RouteTargetAutoVni.Value = false
	}
}

func (data *EVPN) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "replication-type.ingress"); value.Exists() {
		data.ReplicationTypeIngress.Value = true
		data.ReplicationTypeIngress.Null = false
	} else {
		data.ReplicationTypeIngress.Value = false
		data.ReplicationTypeIngress.Null = false
	}
	if value := res.Get(prefix + "replication-type.static"); value.Exists() {
		data.ReplicationTypeStatic.Value = true
		data.ReplicationTypeStatic.Null = false
	} else {
		data.ReplicationTypeStatic.Value = false
		data.ReplicationTypeStatic.Null = false
	}
	if value := res.Get(prefix + "replication-type.p2mp"); value.Exists() {
		data.ReplicationTypeP2mp.Value = true
		data.ReplicationTypeP2mp.Null = false
	} else {
		data.ReplicationTypeP2mp.Value = false
		data.ReplicationTypeP2mp.Null = false
	}
	if value := res.Get(prefix + "replication-type.mp2mp"); value.Exists() {
		data.ReplicationTypeMp2mp.Value = true
		data.ReplicationTypeMp2mp.Null = false
	} else {
		data.ReplicationTypeMp2mp.Value = false
		data.ReplicationTypeMp2mp.Null = false
	}
	if value := res.Get(prefix + "mac.duplication.limit"); value.Exists() {
		data.MacDuplicationLimit.Value = value.Int()
		data.MacDuplicationLimit.Null = false
	}
	if value := res.Get(prefix + "mac.duplication.time"); value.Exists() {
		data.MacDuplicationTime.Value = value.Int()
		data.MacDuplicationTime.Null = false
	}
	if value := res.Get(prefix + "ip.duplication.limit"); value.Exists() {
		data.IpDuplicationLimit.Value = value.Int()
		data.IpDuplicationLimit.Null = false
	}
	if value := res.Get(prefix + "ip.duplication.time"); value.Exists() {
		data.IpDuplicationTime.Value = value.Int()
		data.IpDuplicationTime.Null = false
	}
	if value := res.Get(prefix + "router-id.interface.Loopback"); value.Exists() {
		data.RouterIdLoopback.Value = value.Int()
		data.RouterIdLoopback.Null = false
	}
	if value := res.Get(prefix + "default-gateway.advertise"); value.Exists() {
		data.DefaultGatewayAdvertise.Value = true
		data.DefaultGatewayAdvertise.Null = false
	} else {
		data.DefaultGatewayAdvertise.Value = false
		data.DefaultGatewayAdvertise.Null = false
	}
	if value := res.Get(prefix + "logging.peer.state"); value.Exists() {
		data.LoggingPeerState.Value = true
		data.LoggingPeerState.Null = false
	} else {
		data.LoggingPeerState.Value = false
		data.LoggingPeerState.Null = false
	}
	if value := res.Get(prefix + "route-target.auto.vni"); value.Exists() {
		data.RouteTargetAutoVni.Value = true
		data.RouteTargetAutoVni.Null = false
	} else {
		data.RouteTargetAutoVni.Value = false
		data.RouteTargetAutoVni.Null = false
	}
}

func (data *EVPN) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	if data.ReplicationTypeIngress.Unknown {
		data.ReplicationTypeIngress.Unknown = false
		data.ReplicationTypeIngress.Null = true
	}
	if data.ReplicationTypeStatic.Unknown {
		data.ReplicationTypeStatic.Unknown = false
		data.ReplicationTypeStatic.Null = true
	}
	if data.ReplicationTypeP2mp.Unknown {
		data.ReplicationTypeP2mp.Unknown = false
		data.ReplicationTypeP2mp.Null = true
	}
	if data.ReplicationTypeMp2mp.Unknown {
		data.ReplicationTypeMp2mp.Unknown = false
		data.ReplicationTypeMp2mp.Null = true
	}
	if data.MacDuplicationLimit.Unknown {
		data.MacDuplicationLimit.Unknown = false
		data.MacDuplicationLimit.Null = true
	}
	if data.MacDuplicationTime.Unknown {
		data.MacDuplicationTime.Unknown = false
		data.MacDuplicationTime.Null = true
	}
	if data.IpDuplicationLimit.Unknown {
		data.IpDuplicationLimit.Unknown = false
		data.IpDuplicationLimit.Null = true
	}
	if data.IpDuplicationTime.Unknown {
		data.IpDuplicationTime.Unknown = false
		data.IpDuplicationTime.Null = true
	}
	if data.RouterIdLoopback.Unknown {
		data.RouterIdLoopback.Unknown = false
		data.RouterIdLoopback.Null = true
	}
	if data.DefaultGatewayAdvertise.Unknown {
		data.DefaultGatewayAdvertise.Unknown = false
		data.DefaultGatewayAdvertise.Null = true
	}
	if data.LoggingPeerState.Unknown {
		data.LoggingPeerState.Unknown = false
		data.LoggingPeerState.Null = true
	}
	if data.RouteTargetAutoVni.Unknown {
		data.RouteTargetAutoVni.Unknown = false
		data.RouteTargetAutoVni.Null = true
	}
}

func (data *EVPN) getDeletedListItems(state EVPN) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

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

type BGPNeighbor struct {
	Device               types.String `tfsdk:"device"`
	Id                   types.String `tfsdk:"id"`
	Asn                  types.String `tfsdk:"asn"`
	Ip                   types.String `tfsdk:"ip"`
	RemoteAs             types.String `tfsdk:"remote_as"`
	Description          types.String `tfsdk:"description"`
	Shutdown             types.Bool   `tfsdk:"shutdown"`
	UpdateSourceLoopback types.Int64  `tfsdk:"update_source_loopback"`
}

func (data BGPNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.Value)), url.QueryEscape(fmt.Sprintf("%v", data.Ip.Value)))
}

// if last path element has a key -> remove it
func (data BGPNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPNeighbor) toBody() string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.Null && !data.Ip.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.Value)
	}
	if !data.RemoteAs.Null && !data.RemoteAs.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"remote-as", data.RemoteAs.Value)
	}
	if !data.Description.Null && !data.Description.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.Value)
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.UpdateSourceLoopback.Null && !data.UpdateSourceLoopback.Unknown {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"update-source.interface.Loopback", strconv.FormatInt(data.UpdateSourceLoopback.Value, 10))
	}
	return body
}

func (data *BGPNeighbor) updateFromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() {
		data.Ip.Value = value.String()
	} else {
		data.Ip.Null = true
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() {
		data.RemoteAs.Value = value.String()
	} else {
		data.RemoteAs.Null = true
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
	} else {
		data.Description.Null = true
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	} else {
		data.Shutdown.Value = false
	}
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() {
		data.UpdateSourceLoopback.Value = value.Int()
	} else {
		data.UpdateSourceLoopback.Null = true
	}
}

func (data *BGPNeighbor) fromBody(res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() {
		data.RemoteAs.Value = value.String()
		data.RemoteAs.Null = false
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description.Value = value.String()
		data.Description.Null = false
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown.Value = true
		data.Shutdown.Null = false
	} else {
		data.Shutdown.Value = false
		data.Shutdown.Null = false
	}
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() {
		data.UpdateSourceLoopback.Value = value.Int()
		data.UpdateSourceLoopback.Null = false
	}
}

func (data *BGPNeighbor) setUnknownValues() {
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
	if data.RemoteAs.Unknown {
		data.RemoteAs.Unknown = false
		data.RemoteAs.Null = true
	}
	if data.Description.Unknown {
		data.Description.Unknown = false
		data.Description.Null = true
	}
	if data.Shutdown.Unknown {
		data.Shutdown.Unknown = false
		data.Shutdown.Null = true
	}
	if data.UpdateSourceLoopback.Unknown {
		data.UpdateSourceLoopback.Unknown = false
		data.UpdateSourceLoopback.Null = true
	}
}

func (data *BGPNeighbor) getDeletedListItems(state BGPNeighbor) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

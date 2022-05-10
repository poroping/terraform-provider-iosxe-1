// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceBGPIPv4UnicastVRFNeighborType struct{}

func (t dataSourceBGPIPv4UnicastVRFNeighborType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP IPv4 Unicast VRF Neighbor configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"asn": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"vrf": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"ip": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
			"remote_as": {
				MarkdownDescription: "Specify a BGP peer-group remote-as",
				Type:                types.StringType,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: "Neighbor specific description",
				Type:                types.StringType,
				Computed:            true,
			},
			"shutdown": {
				MarkdownDescription: "Administratively shut down this neighbor",
				Type:                types.BoolType,
				Computed:            true,
			},
			"update_source_loopback": {
				MarkdownDescription: "Loopback interface",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"activate": {
				MarkdownDescription: "Enable the address family for this neighbor",
				Type:                types.BoolType,
				Computed:            true,
			},
			"send_community": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Computed:            true,
			},
			"route_reflector_client": {
				MarkdownDescription: "Configure a neighbor as Route Reflector client",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceBGPIPv4UnicastVRFNeighborType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPIPv4UnicastVRFNeighbor{
		provider: provider,
	}, diags
}

type dataSourceBGPIPv4UnicastVRFNeighbor struct {
	provider provider
}

func (d dataSourceBGPIPv4UnicastVRFNeighbor) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BGPIPv4UnicastVRFNeighbor

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		state = BGPIPv4UnicastVRFNeighbor{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.fromBody(res.Res)
	}

	state.Id = types.String{Value: config.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

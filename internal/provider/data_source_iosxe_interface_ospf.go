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

type dataSourceInterfaceOSPFType struct{}

func (t dataSourceInterfaceOSPFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Interface OSPF configuration.",

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
			"type": {
				MarkdownDescription: "Interface type",
				Type:                types.StringType,
				Required:            true,
			},
			"name": {
				MarkdownDescription: "",
				Type:                types.StringType,
				Required:            true,
			},
			"cost": {
				MarkdownDescription: "Route cost of this interface",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"dead_interval": {
				MarkdownDescription: "Interval after which a neighbor is declared dead",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_interval": {
				MarkdownDescription: "Time between HELLO packets",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"mtu_ignore": {
				MarkdownDescription: "Ignores the MTU in DBD packets",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_type_broadcast": {
				MarkdownDescription: "Specify OSPF broadcast multi-access network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_type_non_broadcast": {
				MarkdownDescription: "Specify OSPF NBMA network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_type_point_to_multipoint": {
				MarkdownDescription: "Specify OSPF point-to-multipoint network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_type_point_to_point": {
				MarkdownDescription: "Specify OSPF point-to-point network",
				Type:                types.BoolType,
				Computed:            true,
			},
			"priority": {
				MarkdownDescription: "Router priority",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceInterfaceOSPFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceInterfaceOSPF{
		provider: provider,
	}, diags
}

type dataSourceInterfaceOSPF struct {
	provider provider
}

func (d dataSourceInterfaceOSPF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config InterfaceOSPF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = InterfaceOSPF{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(res.Res)
	}

	config.Id = types.String{Value: config.getPath()}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

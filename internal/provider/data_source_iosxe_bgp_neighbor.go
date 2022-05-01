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

type dataSourceBGPNeighborType struct{}

func (t dataSourceBGPNeighborType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the BGP Neighbor configuration.",

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
		},
	}, nil
}

func (t dataSourceBGPNeighborType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPNeighbor{
		provider: provider,
	}, diags
}

type dataSourceBGPNeighbor struct {
	provider provider
}

func (d dataSourceBGPNeighbor) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BGPNeighbor

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.provider.clients[config.Device.Value].GetData(config.getPath())
	if res.StatusCode != 404 {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		state.fromBody(res.Res)
	}
	state.fromPlan(config)
	state.Id.Value = config.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

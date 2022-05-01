// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
	"github.com/netascode/terraform-provider-iosxe/internal/provider/helpers"
)

type resourceInterfaceVLANType struct{}

func (t resourceInterfaceVLANType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Interface VLAN configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"name": {
				MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 4094).String,
				Type:                types.Int64Type,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4094),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"autostate": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable auto-state determination for VLAN").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Interface specific description").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 200, `.*`),
				},
			},
			"shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown the selected interface").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"vrf_forwarding": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure forwarding table").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"ipv4_address": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"ipv4_address_mask": {
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"pim_sparse_mode": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable PIM sparse-mode operation").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (t resourceInterfaceVLANType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceInterfaceVLAN{
		provider: provider,
	}, diags
}

type resourceInterfaceVLAN struct {
	provider provider
}

func (r resourceInterfaceVLAN) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state InterfaceVLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody()

	res, _ := r.provider.clients[plan.Device.Value].GetData(plan.getPath(), restconf.Query("depth", "1"))
	if res.StatusCode < 200 || res.StatusCode > 299 {
		_, err := r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s", err))
			return
		}
	} else {
		_, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PATCH), got error: %s", err))
			return
		}
	}

	// Read object
	res, err := r.provider.clients[plan.Device.Value].GetData(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res.Res)
	state.fromPlan(plan)
	state.Id.Value = plan.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceInterfaceVLAN) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state, newState InterfaceVLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.Value))

	res, err := r.provider.clients[state.Device.Value].GetData(state.Id.Value)
	if res.StatusCode != 404 {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		newState.fromBody(res.Res)
	}
	newState.fromPlan(state)
	newState.Id = state.Id

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.Value))

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}

func (r resourceInterfaceVLAN) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state InterfaceVLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.Value))

	// Update object
	body := plan.toBody()
	res, _ := r.provider.clients[plan.Device.Value].GetData(plan.getPath(), restconf.Query("depth", "1"))
	if res.StatusCode < 200 || res.StatusCode > 299 {
		_, err := r.provider.clients[plan.Device.Value].PutData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	} else {
		_, err := r.provider.clients[plan.Device.Value].PatchData(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PATCH), got error: %s", err))
			return
		}
	}

	// Read object
	res, err := r.provider.clients[plan.Device.Value].GetData(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res.Res)
	state.fromPlan(plan)
	state.Id.Value = plan.getPath()

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceInterfaceVLAN) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state InterfaceVLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.Value))

	_, err := r.provider.clients[state.Device.Value].DeleteData(state.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceInterfaceVLAN) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}

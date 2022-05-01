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

type resourceOSPFType struct{}

func (t resourceOSPFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the OSPF configuration.",

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
			"process_id": {
				MarkdownDescription: helpers.NewAttributeDescription("Process ID").AddIntegerRangeDescription(1, 65535).String,
				Type:                types.Int64Type,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 65535),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"bfd_all_interfaces": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable BFD on all interfaces").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"default_information_originate": {
				MarkdownDescription: helpers.NewAttributeDescription("Distribute a default route").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"default_information_originate_always": {
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"default_metric": {
				MarkdownDescription: helpers.NewAttributeDescription("Set metric of redistributed routes").AddIntegerRangeDescription(1, 16777214).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 16777214),
				},
			},
			"distance": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 255),
				},
			},
			"domain_tag": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF domain-tag").AddIntegerRangeDescription(1, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 4294967295),
				},
			},
			"mpls_ldp_autoconfig": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure LDP automatic configuration").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"mpls_ldp_sync": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure LDP-IGP Synchronization").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"neighbor": {
				MarkdownDescription: helpers.NewAttributeDescription("Specify a neighbor router").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: helpers.NewAttributeDescription("Neighbor address").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"priority": {
						MarkdownDescription: helpers.NewAttributeDescription("OSPF priority of non-broadcast neighbor").AddIntegerRangeDescription(0, 255).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(0, 255),
						},
					},
					"cost": {
						MarkdownDescription: helpers.NewAttributeDescription("OSPF cost for point-to-multipoint neighbor").AddIntegerRangeDescription(1, 65535).String,
						Type:                types.Int64Type,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.IntegerRangeValidator(1, 65535),
						},
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"network": {
				MarkdownDescription: helpers.NewAttributeDescription("Enable routing on an IP network").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: helpers.NewAttributeDescription("Network number").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"wildcard": {
						MarkdownDescription: helpers.NewAttributeDescription("OSPF wild card bits").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"area": {
						MarkdownDescription: helpers.NewAttributeDescription("Set the OSPF area ID").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
			"priority": {
				MarkdownDescription: helpers.NewAttributeDescription("OSPF topology priority").AddIntegerRangeDescription(0, 127).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 127),
				},
			},
			"router_id": {
				MarkdownDescription: helpers.NewAttributeDescription("Override configured router identifier (peers will reset)").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
				},
			},
			"shutdown": {
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown the OSPF protocol under the current instance").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
			},
			"summary_address": {
				MarkdownDescription: helpers.NewAttributeDescription("Configure IP address summaries").String,
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip": {
						MarkdownDescription: helpers.NewAttributeDescription("IP summary address").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
					"mask": {
						MarkdownDescription: helpers.NewAttributeDescription("Summary mask").String,
						Type:                types.StringType,
						Optional:            true,
						Computed:            true,
						Validators: []tfsdk.AttributeValidator{
							helpers.StringPatternValidator(0, 0, `(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?`),
						},
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t resourceOSPFType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceOSPF{
		provider: provider,
	}, diags
}

type resourceOSPF struct {
	provider provider
}

func (r resourceOSPF) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state OSPF

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

func (r resourceOSPF) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state, newState OSPF

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

func (r resourceOSPF) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state OSPF

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

func (r resourceOSPF) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state OSPF

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

func (r resourceOSPF) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}

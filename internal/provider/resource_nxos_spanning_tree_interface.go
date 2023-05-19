// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SpanningTreeInterfaceResource{}
var _ resource.ResourceWithImportState = &SpanningTreeInterfaceResource{}

func NewSpanningTreeInterfaceResource() resource.Resource {
	return &SpanningTreeInterfaceResource{}
}

type SpanningTreeInterfaceResource struct {
	data *NxosProviderData
}

func (r *SpanningTreeInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spanning_tree_interface"
}

func (r *SpanningTreeInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the Spanning Tree interface configuration.", "stpIf", "Discovery%20Protocols/stp:If/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The administrative state of the object or policy.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("enabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"bpdu_filter": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BPDU filter mode.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "enable", "disable"),
				},
			},
			"bpdu_guard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BPDU guard mode.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "enable", "disable"),
				},
			},
			"cost": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port path cost.").AddIntegerRangeDescription(0, 200000000).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 200000000),
				},
			},
			"ctrl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface controls.").AddStringEnumDescription("unspecified", "bpdu-guard", "bpdu-filter").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "bpdu-guard", "bpdu-filter"),
				},
			},
			"guard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Guard mode.").AddStringEnumDescription("default", "root", "loop", "none").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "root", "loop", "none"),
				},
			},
			"link_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Link type.").AddStringEnumDescription("auto", "p2p", "shared").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("auto"),
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "p2p", "shared"),
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port mode.").AddStringEnumDescription("default", "edge", "network", "normal", "trunk").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "edge", "network", "normal", "trunk"),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port priority.").AddIntegerRangeDescription(0, 224).AddDefaultValueDescription("128").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128),
				Validators: []validator.Int64{
					int64validator.Between(0, 224),
				},
			},
		},
	}
}

func (r *SpanningTreeInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *SpanningTreeInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SpanningTreeInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SpanningTreeInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SpanningTreeInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
	queries = append(queries, nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	res, err := r.data.client.GetDn(state.Dn.ValueString(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res, false)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *SpanningTreeInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SpanningTreeInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SpanningTreeInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SpanningTreeInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	res, err := r.data.client.DeleteDn(state.Dn.ValueString(), nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "1" && errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *SpanningTreeInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

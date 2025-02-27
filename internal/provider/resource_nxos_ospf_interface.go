// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &OSPFInterfaceResource{}
var _ resource.ResourceWithImportState = &OSPFInterfaceResource{}

func NewOSPFInterfaceResource() resource.Resource {
	return &OSPFInterfaceResource{}
}

type OSPFInterfaceResource struct {
	clients map[string]*nxos.Client
}

func (r *OSPFInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ospf_interface"
}

func (r *OSPFInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the OSPF interface configuration.", "ospfIf", "Routing%20and%20Forwarding/ospf:If/").AddParents("ospf_vrf").AddChildren("ospf_authentication").String,

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
			"instance_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF instance name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"advertise_secondaries": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise secondary IP addresses.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"area": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Area identifier to which a network or interface belongs in IPv4 address format.").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"bfd": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bidirectional Forwarding Detection (BFD).").AddStringEnumDescription("unspecified", "enabled", "disabled").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "enabled", "disabled"),
				},
			},
			"cost": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies the cost of interface.").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"dead_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dead interval, interval after which router declares that neighbor as down.").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval, interval between hello packets that OSPF sends on the interface.").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"network_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network type.").AddStringEnumDescription("unspecified", "p2p", "bcast").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "p2p", "bcast"),
				},
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passive interface control. Interface can be configured as passive or non-passive.").AddStringEnumDescription("unspecified", "enabled", "disabled").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "enabled", "disabled"),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Priority, used in determining the designated router on this network.").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
		},
	}
}

func (r *OSPFInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *OSPFInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OSPFInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	client, ok := r.clients[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	// Post object
	body := plan.toBody(false)
	_, err := client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *OSPFInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OSPFInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	client, ok := r.clients[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
	res, err := client.GetDn(state.Dn.ValueString(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}
	state.fromBody(res, imp)
	if imp {
		state.getIdsFromDn()
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *OSPFInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OSPFInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	client, ok := r.clients[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	body := plan.toBody(false)

	_, err := client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *OSPFInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OSPFInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	client, ok := r.clients[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	body := state.toDeleteBody()

	if len(body.Str) > 0 {
		_, err := client.Post(state.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := client.DeleteDn(state.Dn.ValueString())
		if err != nil {
			errCode := res.Get("imdata.0.error.attributes.code").Str
			// Ignore errors of type "Cannot delete object"
			if errCode != "1" && errCode != "107" {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *OSPFInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

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
var _ resource.Resource = &ISISVRFResource{}
var _ resource.ResourceWithImportState = &ISISVRFResource{}

func NewISISVRFResource() resource.Resource {
	return &ISISVRFResource{}
}

type ISISVRFResource struct {
	clients map[string]*nxos.Client
}

func (r *ISISVRFResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_vrf"
}

func (r *ISISVRFResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the IS-IS VRF configuration.", "isisDom", "Routing%20and%20Forwarding/isis:Dom/").AddParents("isis_instance").AddReferences("vrf").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS instance name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative state.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("enabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"authentication_check_l1": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-1.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"authentication_check_l2": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-2.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"authentication_key_l1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-1.").String,
				Optional:            true,
			},
			"authentication_key_l2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-2.").String,
				Optional:            true,
			},
			"authentication_type_l1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-1.").AddStringEnumDescription("clear", "md5", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "md5", "unknown"),
				},
			},
			"authentication_type_l2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-2.").AddStringEnumDescription("clear", "md5", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "md5", "unknown"),
				},
			},
			"bandwidth_reference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost metric.").AddIntegerRangeDescription(0, 4294967295).AddDefaultValueDescription("40000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(40000),
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"banwidth_reference_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bandwidth reference unit.").AddStringEnumDescription("mbps", "gbps").AddDefaultValueDescription("mbps").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("mbps"),
				Validators: []validator.String{
					stringvalidator.OneOf("mbps", "gbps"),
				},
			},
			"is_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS domain type.").AddStringEnumDescription("l1", "l2", "l12").AddDefaultValueDescription("l12").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("l12"),
				Validators: []validator.String{
					stringvalidator.OneOf("l1", "l2", "l12"),
				},
			},
			"metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS metric type.").AddStringEnumDescription("narrow", "wide", "transition").AddDefaultValueDescription("wide").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("wide"),
				Validators: []validator.String{
					stringvalidator.OneOf("narrow", "wide", "transition"),
				},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352 bytes.").AddIntegerRangeDescription(256, 4352).AddDefaultValueDescription("1492").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1492),
				Validators: []validator.Int64{
					int64validator.Between(256, 4352),
				},
			},
			"net": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Holds IS-IS domain NET (address) value.").String,
				Optional:            true,
			},
			"passive_default": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Domain passive-interface default level.").AddStringEnumDescription("l1", "l2", "l12", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("l1", "l2", "l12", "unknown"),
				},
			},
		},
	}
}

func (r *ISISVRFResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *ISISVRFResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ISISVRF

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

func (r *ISISVRFResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ISISVRF

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

func (r *ISISVRFResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ISISVRF

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

func (r *ISISVRFResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ISISVRF

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

func (r *ISISVRFResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

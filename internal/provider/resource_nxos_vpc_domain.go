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
var _ resource.Resource = &VPCDomainResource{}
var _ resource.ResourceWithImportState = &VPCDomainResource{}

func NewVPCDomainResource() resource.Resource {
	return &VPCDomainResource{}
}

type VPCDomainResource struct {
	clients map[string]*nxos.Client
}

func (r *VPCDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_domain"
}

func (r *VPCDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the vPC domain configuration.", "vpcDom", "System/vpc:Dom/").AddParents("vpc_instance").AddChildren("vpc_interface", "vpc_keepalive").String,

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
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC suspend locally.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("enabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"domain_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Domain id.").AddIntegerRangeDescription(1, 1000).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1000),
				},
			},
			"auto_recovery": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Auto Recovery.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"auto_recovery_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Auto Recovery interval.").AddIntegerRangeDescription(60, 3600).AddDefaultValueDescription("240").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(240),
				Validators: []validator.Int64{
					int64validator.Between(60, 3600),
				},
			},
			"delay_restore_orphan_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Delay restore for orphan ports.").AddIntegerRangeDescription(0, 300).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 300),
				},
			},
			"delay_restore_svi": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Delay restore for SVI.").AddIntegerRangeDescription(1, 3600).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"delay_restore_vpc": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Delay restore for vPC links.").AddIntegerRangeDescription(1, 3600).AddDefaultValueDescription("30").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("DSCP.").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"fast_convergence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fast Convergence.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"graceful_consistency_check": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Graceful Type-1 Consistency Check.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("enabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"l3_peer_router": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("L3 Peer Router.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"l3_peer_router_syslog": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("L3 Peer Router Syslog.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"l3_peer_router_syslog_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("L3 Peer Router Syslog Interval.").AddIntegerRangeDescription(1, 3600).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3600),
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"peer_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer Gateway.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"peer_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC peer IP address.").String,
				Optional:            true,
			},
			"peer_switch": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC pair switches.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"role_priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Role priority.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("32667").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(32667),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"sys_mac": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("System MAC.").String,
				Optional:            true,
			},
			"system_priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("System priority.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("32667").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(32667),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"track": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tracking object to suspend vPC if object goes down.").AddIntegerRangeDescription(0, 500).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 500),
				},
			},
			"virtual_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC virtual IP address (vIP).").String,
				Optional:            true,
			},
		},
	}
}

func (r *VPCDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *VPCDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPCDomain

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

func (r *VPCDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPCDomain

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *VPCDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VPCDomain

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

func (r *VPCDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPCDomain

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

func (r *VPCDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

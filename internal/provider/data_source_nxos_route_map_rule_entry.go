// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouteMapRuleEntryDataSource{}
	_ datasource.DataSourceWithConfigure = &RouteMapRuleEntryDataSource{}
)

func NewRouteMapRuleEntryDataSource() datasource.DataSource {
	return &RouteMapRuleEntryDataSource{}
}

type RouteMapRuleEntryDataSource struct {
	clients map[string]*nxos.Client
}

func (d *RouteMapRuleEntryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_map_rule_entry"
}

func (d *RouteMapRuleEntryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read a Route-Map Rule Entry configuration.", "rtmapEntry", "Routing%20and%20Forwarding/rtmap:Entry/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"rule_name": schema.StringAttribute{
				MarkdownDescription: "Route Map rule name.",
				Required:            true,
			},
			"order": schema.Int64Attribute{
				MarkdownDescription: "Route-Map Rule Entry order.",
				Required:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Route-Map Rule Entry action.",
				Computed:            true,
			},
		},
	}
}

func (d *RouteMapRuleEntryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *RouteMapRuleEntryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouteMapRuleEntry

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	queries := []func(*nxos.Req){}
	res, err := d.clients[config.Device.ValueString()].GetDn(config.getDn(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(res, true)
	config.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

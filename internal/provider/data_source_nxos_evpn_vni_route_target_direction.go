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
	_ datasource.DataSource              = &EVPNVNIRouteTargetDirectionDataSource{}
	_ datasource.DataSourceWithConfigure = &EVPNVNIRouteTargetDirectionDataSource{}
)

func NewEVPNVNIRouteTargetDirectionDataSource() datasource.DataSource {
	return &EVPNVNIRouteTargetDirectionDataSource{}
}

type EVPNVNIRouteTargetDirectionDataSource struct {
	clients map[string]*nxos.Client
}

func (d *EVPNVNIRouteTargetDirectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_evpn_vni_route_target_direction"
}

func (d *EVPNVNIRouteTargetDirectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read a EVPN VNI Route Target direction.", "rtctrlRttP", "Routing%20and%20Forwarding/rtctrl:RttP/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"encap": schema.StringAttribute{
				MarkdownDescription: "Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Required:            true,
			},
			"direction": schema.StringAttribute{
				MarkdownDescription: "Route Target direction.",
				Required:            true,
			},
		},
	}
}

func (d *EVPNVNIRouteTargetDirectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *EVPNVNIRouteTargetDirectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config EVPNVNIRouteTargetDirection

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

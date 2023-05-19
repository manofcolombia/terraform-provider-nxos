// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OSPFAuthenticationDataSource{}
	_ datasource.DataSourceWithConfigure = &OSPFAuthenticationDataSource{}
)

func NewOSPFAuthenticationDataSource() datasource.DataSource {
	return &OSPFAuthenticationDataSource{}
}

type OSPFAuthenticationDataSource struct {
	data *NxosProviderData
}

func (d *OSPFAuthenticationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ospf_authentication"
}

func (d *OSPFAuthenticationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the OSPF authentication configuration.", "ospfAuthNewP", "Routing%20and%20Forwarding/ospf:AuthNewP/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: "OSPF instance name.",
				Required:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Required:            true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Key used for authentication.",
				Computed:            true,
			},
			"key_id": schema.Int64Attribute{
				MarkdownDescription: "Key ID used for authentication.",
				Computed:            true,
			},
			"key_secure_mode": schema.BoolAttribute{
				MarkdownDescription: "Encrypted authentication key or plain text key.",
				Computed:            true,
			},
			"keychain": schema.StringAttribute{
				MarkdownDescription: "Authentication keychain.",
				Computed:            true,
			},
			"md5_key": schema.StringAttribute{
				MarkdownDescription: "Key used for md5 authentication.",
				Computed:            true,
			},
			"md5_key_secure_mode": schema.BoolAttribute{
				MarkdownDescription: "Encrypted authentication md5 key or plain text key.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Authentication type.",
				Computed:            true,
			},
		},
	}
}

func (d *OSPFAuthenticationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *OSPFAuthenticationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OSPFAuthentication

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	queries := []func(*nxos.Req){nxos.OverrideUrl(d.data.devices[config.Device.ValueString()])}
	res, err := d.data.client.GetDn(config.getDn(), queries...)
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

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
	_ datasource.DataSource              = &PIMSSMRangeDataSource{}
	_ datasource.DataSourceWithConfigure = &PIMSSMRangeDataSource{}
)

func NewPIMSSMRangeDataSource() datasource.DataSource {
	return &PIMSSMRangeDataSource{}
}

type PIMSSMRangeDataSource struct {
	data *NxosProviderData
}

func (d *PIMSSMRangeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pim_ssm_range"
}

func (d *PIMSSMRangeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the PIM SSM range configuration.", "pimSSMRangeP", "Layer%203/pim:SSMRangeP/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"group_list_1": schema.StringAttribute{
				MarkdownDescription: "Group list 1.",
				Computed:            true,
			},
			"group_list_2": schema.StringAttribute{
				MarkdownDescription: "Group list 2.",
				Computed:            true,
			},
			"group_list_3": schema.StringAttribute{
				MarkdownDescription: "Group list 3.",
				Computed:            true,
			},
			"group_list_4": schema.StringAttribute{
				MarkdownDescription: "Group list 4.",
				Computed:            true,
			},
			"prefix_list": schema.StringAttribute{
				MarkdownDescription: "Prefix list name.",
				Computed:            true,
			},
			"route_map": schema.StringAttribute{
				MarkdownDescription: "Route map name.",
				Computed:            true,
			},
			"ssm_none": schema.BoolAttribute{
				MarkdownDescription: "Exclude standard SSM range (232.0.0.0/8).",
				Computed:            true,
			},
		},
	}
}

func (d *PIMSSMRangeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *PIMSSMRangeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PIMSSMRange

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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type dataSourceSubinterfaceVRFType struct{}

func (t dataSourceSubinterfaceVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read a subinterface VRF association.", "nwRtVrfMbr", "Routing%20and%20Forwarding/nw:RtVrfMbr/").String,

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"interface_id": {
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.",
				Type:                types.StringType,
				Required:            true,
			},
			"vrf_dn": {
				MarkdownDescription: "DN of VRF. For example: `sys/inst-VRF1`.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceSubinterfaceVRFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceSubinterfaceVRF{
		provider: provider,
	}, diags
}

type dataSourceSubinterfaceVRF struct {
	provider provider
}

func (d dataSourceSubinterfaceVRF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state SubinterfaceVRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn(), nxos.OverrideUrl(d.provider.devices[config.Device.Value]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

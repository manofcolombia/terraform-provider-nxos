// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceDHCPRelayAddressType struct{}

func (t dataSourceDHCPRelayAddressType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read a DHCP relay address.",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"interface_id": {
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Type:                types.StringType,
				Required:            true,
			},
			"vrf": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"address": {
				MarkdownDescription: "IPv4 or IPv6 address.",
				Type:                types.StringType,
				Required:            true,
			},
		},
	}, nil
}

func (t dataSourceDHCPRelayAddressType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceDHCPRelayAddress{
		provider: provider,
	}, diags
}

type dataSourceDHCPRelayAddress struct {
	provider provider
}

func (d dataSourceDHCPRelayAddress) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state DHCPRelayAddress

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn())

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

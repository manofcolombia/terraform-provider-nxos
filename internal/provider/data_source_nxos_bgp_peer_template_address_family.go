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

type dataSourceBGPPeerTemplateAddressFamilyType struct{}

func (t dataSourceBGPPeerTemplateAddressFamilyType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer template address family configuration.", "bgpPeerAf", "Routing%20and%20Forwarding/bgp:PeerAf/").String,

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
			"asn": {
				MarkdownDescription: "Autonomous system number.",
				Type:                types.StringType,
				Required:            true,
			},
			"template_name": {
				MarkdownDescription: "Peer template name.",
				Type:                types.StringType,
				Required:            true,
			},
			"address_family": {
				MarkdownDescription: "Address Family.",
				Type:                types.StringType,
				Required:            true,
			},
			"control": {
				MarkdownDescription: "Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`, `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats:\n  - Single value. Example: `nh-self`\n  - Multiple values (comma-separated). Example: `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.",
				Type:                types.StringType,
				Computed:            true,
			},
			"send_community_extended": {
				MarkdownDescription: "Send-community extended.",
				Type:                types.StringType,
				Computed:            true,
			},
			"send_community_standard": {
				MarkdownDescription: "Send-community standard.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceBGPPeerTemplateAddressFamilyType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPPeerTemplateAddressFamily{
		provider: provider,
	}, diags
}

type dataSourceBGPPeerTemplateAddressFamily struct {
	provider provider
}

func (d dataSourceBGPPeerTemplateAddressFamily) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BGPPeerTemplateAddressFamily

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

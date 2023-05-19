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
	_ datasource.DataSource              = &IPv4AccessListEntryDataSource{}
	_ datasource.DataSourceWithConfigure = &IPv4AccessListEntryDataSource{}
)

func NewIPv4AccessListEntryDataSource() datasource.DataSource {
	return &IPv4AccessListEntryDataSource{}
}

type IPv4AccessListEntryDataSource struct {
	data *NxosProviderData
}

func (d *IPv4AccessListEntryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_access_list_entry"
}

func (d *IPv4AccessListEntryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read IPv4 Access List Entries.", "ipv4aclACE", "Security%20and%20Policing/ipv4acl:ACE/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Access list name.",
				Required:            true,
			},
			"sequence_number": schema.Int64Attribute{
				MarkdownDescription: "Sequence number.",
				Required:            true,
			},
			"ack": schema.BoolAttribute{
				MarkdownDescription: "Match TCP ACK flag.",
				Computed:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action.",
				Computed:            true,
			},
			"dscp": schema.Int64Attribute{
				MarkdownDescription: "Match DSCP.",
				Computed:            true,
			},
			"destination_address_group": schema.StringAttribute{
				MarkdownDescription: "Destination address group.",
				Computed:            true,
			},
			"destination_port_1": schema.StringAttribute{
				MarkdownDescription: "First destination port number or name.",
				Computed:            true,
			},
			"destination_port_2": schema.StringAttribute{
				MarkdownDescription: "Second destination port number or name.",
				Computed:            true,
			},
			"destination_port_group": schema.StringAttribute{
				MarkdownDescription: "Destination port group.",
				Computed:            true,
			},
			"destination_port_mask": schema.StringAttribute{
				MarkdownDescription: "Destination port mask number or name.",
				Computed:            true,
			},
			"destination_port_operator": schema.StringAttribute{
				MarkdownDescription: "Destination port operator.",
				Computed:            true,
			},
			"destination_prefix": schema.StringAttribute{
				MarkdownDescription: "Destination prefix.",
				Computed:            true,
			},
			"destination_prefix_length": schema.StringAttribute{
				MarkdownDescription: "Destination prefix length.",
				Computed:            true,
			},
			"destination_prefix_mask": schema.StringAttribute{
				MarkdownDescription: "Destination prefix mask.",
				Computed:            true,
			},
			"est": schema.BoolAttribute{
				MarkdownDescription: "Match TCP EST flag.",
				Computed:            true,
			},
			"fin": schema.BoolAttribute{
				MarkdownDescription: "Match TCP FIN flag.",
				Computed:            true,
			},
			"fragment": schema.BoolAttribute{
				MarkdownDescription: "Match non-initial fragment.",
				Computed:            true,
			},
			"http_option_type": schema.StringAttribute{
				MarkdownDescription: "HTTP option method.",
				Computed:            true,
			},
			"icmp_code": schema.Int64Attribute{
				MarkdownDescription: "ICMP code.",
				Computed:            true,
			},
			"icmp_type": schema.Int64Attribute{
				MarkdownDescription: "ICMP type.",
				Computed:            true,
			},
			"logging": schema.BoolAttribute{
				MarkdownDescription: "Log matches against ACL entry.",
				Computed:            true,
			},
			"packet_length_1": schema.StringAttribute{
				MarkdownDescription: "First packet length. Either `invalid` or a number between 19 and 9210.",
				Computed:            true,
			},
			"packet_length_2": schema.StringAttribute{
				MarkdownDescription: "Second packet length. Either `invalid` or a number between 19 and 9210.",
				Computed:            true,
			},
			"packet_length_operator": schema.StringAttribute{
				MarkdownDescription: "Packet length operator.",
				Computed:            true,
			},
			"precedence": schema.StringAttribute{
				MarkdownDescription: "Precedence. Either `unspecified` or a number between 0 and 7.",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Protocol name or number.",
				Computed:            true,
			},
			"protocol_mask": schema.StringAttribute{
				MarkdownDescription: "Protocol mask name or number.",
				Computed:            true,
			},
			"psh": schema.BoolAttribute{
				MarkdownDescription: "Match TCP PSH flag.",
				Computed:            true,
			},
			"redirect": schema.StringAttribute{
				MarkdownDescription: "Redirect action.",
				Computed:            true,
			},
			"remark": schema.StringAttribute{
				MarkdownDescription: "ACL comment.",
				Computed:            true,
			},
			"rev": schema.BoolAttribute{
				MarkdownDescription: "Match TCP REV flag.",
				Computed:            true,
			},
			"rst": schema.BoolAttribute{
				MarkdownDescription: "Match TCP RST flag.",
				Computed:            true,
			},
			"source_address_group": schema.StringAttribute{
				MarkdownDescription: "Source address group.",
				Computed:            true,
			},
			"source_port_1": schema.StringAttribute{
				MarkdownDescription: "First source port name or number.",
				Computed:            true,
			},
			"source_port_2": schema.StringAttribute{
				MarkdownDescription: "Second source port name or number.",
				Computed:            true,
			},
			"source_port_group": schema.StringAttribute{
				MarkdownDescription: "Source port group.",
				Computed:            true,
			},
			"source_port_mask": schema.StringAttribute{
				MarkdownDescription: "Source port mask name or number.",
				Computed:            true,
			},
			"source_port_operator": schema.StringAttribute{
				MarkdownDescription: "Source port operator.",
				Computed:            true,
			},
			"source_prefix": schema.StringAttribute{
				MarkdownDescription: "Source prefix.",
				Computed:            true,
			},
			"source_prefix_length": schema.StringAttribute{
				MarkdownDescription: "Source prefix length.",
				Computed:            true,
			},
			"source_prefix_mask": schema.StringAttribute{
				MarkdownDescription: "Source prefix mask.",
				Computed:            true,
			},
			"syn": schema.BoolAttribute{
				MarkdownDescription: "Match TCP SYN flag.",
				Computed:            true,
			},
			"time_range": schema.StringAttribute{
				MarkdownDescription: "Time range name.",
				Computed:            true,
			},
			"ttl": schema.Int64Attribute{
				MarkdownDescription: "TTL.",
				Computed:            true,
			},
			"urg": schema.BoolAttribute{
				MarkdownDescription: "Match TCP URG flag.",
				Computed:            true,
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: "VLAN ID.",
				Computed:            true,
			},
			"vni": schema.StringAttribute{
				MarkdownDescription: "NVE VNI ID. Either `invalid` or a number between 0 and 16777216.",
				Computed:            true,
			},
		},
	}
}

func (d *IPv4AccessListEntryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *IPv4AccessListEntryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPv4AccessListEntry

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

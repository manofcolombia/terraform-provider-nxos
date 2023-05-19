// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type IPv4PrefixListRule struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
}

func (data IPv4PrefixListRule) getDn() string {
	return fmt.Sprintf("sys/rpm/pfxlistv4-[%s]", data.Name.ValueString())
}

func (data IPv4PrefixListRule) getClassName() string {
	return "rtpfxRuleV4"
}

func (data IPv4PrefixListRule) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}

	return nxos.Body{body}
}

func (data *IPv4PrefixListRule) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
}

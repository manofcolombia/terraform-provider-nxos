// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type PIM struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data PIM) getDn() string {
	return "sys/pim"
}

func (data PIM) getClassName() string {
	return "pimEntity"
}

func (data PIM) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PIM) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *PIM) fromPlan(plan PIM) {
	data.Device = plan.Device
}

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

type PIMInterface struct {
	Device        types.String `tfsdk:"device"`
	Dn            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"vrf_name"`
	Id            types.String `tfsdk:"interface_id"`
	AdminSt       types.String `tfsdk:"admin_state"`
	BfdInst       types.String `tfsdk:"bfd"`
	DrPrio        types.Int64  `tfsdk:"dr_priority"`
	Passive       types.Bool   `tfsdk:"passive"`
	PimSparseMode types.Bool   `tfsdk:"sparse_mode"`
}

func (data PIMInterface) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/if-[%s]", data.Name.Value, data.Id.Value)
}

func (data PIMInterface) getClassName() string {
	return "pimIf"
}

func (data PIMInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("adminSt", data.AdminSt.Value).
		Set("bfdInst", data.BfdInst.Value).
		Set("drPrio", strconv.FormatInt(data.DrPrio.Value, 10)).
		Set("passive", strconv.FormatBool(data.Passive.Value)).
		Set("pimSparseMode", strconv.FormatBool(data.PimSparseMode.Value))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PIMInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
	data.BfdInst.Value = res.Get("*.attributes.bfdInst").String()
	data.DrPrio.Value = res.Get("*.attributes.drPrio").Int()
	data.Passive.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.passive").String())
	data.PimSparseMode.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.pimSparseMode").String())
}

func (data *PIMInterface) fromPlan(plan PIMInterface) {
	data.Device = plan.Device
	data.Name.Value = plan.Name.Value
}

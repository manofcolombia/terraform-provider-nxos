//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Device types.String `tfsdk:"device"`
	Dn types.String `tfsdk:"id"`
{{- range .Attributes}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- range .ChildClasses}}
{{- if eq .Type "single"}}
{{- range .Attributes}}
    {{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- else if eq .Type "list"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}

{{range .ChildClasses}}
{{if eq .Type "list"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
}
{{end}}
{{end}}

func (data {{camelCase .Name}}) getDn() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Dn}}"{{range .Attributes}}{{if .Id}}, data.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})
{{- else}}
	return "{{.Dn}}"
{{- end}}
}

{{range .ChildClasses}}
{{if eq .Type "list"}}
func (data {{$name}}{{toGoName .TfName}}) getRn() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Rn}}"{{range .Attributes}}{{if .Id}}, data.{{toGoName .TfName}}.Value{{.Type}}(){{end}}{{end}})
{{- else}}
	return "{{.Rn}}"
{{- end}}
}
{{end}}
{{end}}

func (data {{camelCase .Name}}) getClassName() string {
	return "{{.ClassName}}"
}

func (data {{camelCase .Name}}) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (statusReplace){
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	{{- range .Attributes}}
	{{- if not .ReferenceOnly}}
	if (!data.{{toGoName .TfName}}.IsUnknown() && !data.{{toGoName .TfName}}.IsNull()) || {{not .OmitEmptyValue}} {
		{{- if eq .Type "Int64"}}
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", strconv.FormatInt(data.{{toGoName .TfName}}.ValueInt64(), 10))
		{{- else if eq .Type "Bool"}}
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", strconv.FormatBool(data.{{toGoName .TfName}}.ValueBool()))
		{{- else if eq .Type "String"}}
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", data.{{toGoName .TfName}}.ValueString())
		{{- end}}
	}
	{{- end}}
	{{- end}}

	{{- if .ChildClasses}}
	var attrs string
	{{- end}}
	{{- range .ChildClasses}}
	{{- $childClassName := .ClassName }}
	{{- if eq .Type "single"}}
	attrs = ""
	{{- range .Attributes}}
	if (!data.{{toGoName .TfName}}.IsUnknown() && !data.{{toGoName .TfName}}.IsNull()) || {{not .OmitEmptyValue}} {
		{{- if eq .Type "Int64"}}
		attrs, _ = sjson.Set(attrs, "{{.NxosName}}", strconv.FormatInt(data.{{toGoName .TfName}}.ValueInt64(), 10))
		{{- else if eq .Type "Bool"}}
		attrs, _ = sjson.Set(attrs, "{{.NxosName}}", strconv.FormatBool(data.{{toGoName .TfName}}.ValueBool()))
		{{- else if eq .Type "String"}}
		attrs, _ = sjson.Set(attrs, "{{.NxosName}}", data.{{toGoName .TfName}}.ValueString())
		{{- end}}
	}
	{{- end}}
	body, _ = sjson.SetRaw(body, data.getClassName()+".children.-1.{{$childClassName}}.attributes", attrs)
	{{- else if eq .Type "list"}}
	for _, child := range data.{{toGoName .TfName}} {
		attrs = ""
		{{- range .Attributes}}
		if (!child.{{toGoName .TfName}}.IsUnknown() && !child.{{toGoName .TfName}}.IsNull()) || {{not .OmitEmptyValue}} {
			{{- if eq .Type "Int64"}}
			attrs, _ = sjson.Set(attrs, "{{.NxosName}}", strconv.FormatInt(child.{{toGoName .TfName}}.ValueInt64(), 10))
			{{- else if eq .Type "Bool"}}
			attrs, _ = sjson.Set(attrs, "{{.NxosName}}", strconv.FormatBool(child.{{toGoName .TfName}}.ValueBool()))
			{{- else if eq .Type "String"}}
			attrs, _ = sjson.Set(attrs, "{{.NxosName}}", child.{{toGoName .TfName}}.ValueString())
			{{- end}}
		}
		{{- end}}
		body, _ = sjson.SetRaw(body, data.getClassName()+".children.-1.{{$childClassName}}.attributes", attrs)
	}
	{{- end}}
	{{- end}}

	return nxos.Body{body}
}

func (data *{{camelCase .Name}}) fromBody(res gjson.Result, all bool) {
	{{- range .Attributes}}
	{{- if and (not .ReferenceOnly) (not .WriteOnly)}}
	if !data.{{toGoName .TfName}}.IsNull() || all {
		{{- if eq .Type "Int64"}}
		data.{{toGoName .TfName}} = types.Int64Value(res.Get(data.getClassName()+".attributes.{{.NxosName}}").Int())
		{{- else if eq .Type "Bool"}}
		data.{{toGoName .TfName}} = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName()+".attributes.{{.NxosName}}").String()))
		{{- else if eq .Type "String"}}
		data.{{toGoName .TfName}} = types.StringValue(res.Get(data.getClassName()+".attributes.{{.NxosName}}").String())
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
	}
	{{- end}}
	{{- end}}

	{{- range .ChildClasses}}
	{{- $childClassName := .ClassName }}
	{{- $childRn := .Rn }}
	{{- $list := (toGoName .TfName)}}
	{{- if eq .Type "single"}}
	var r gjson.Result
	res.Get(data.getClassName() + ".children").ForEach(
		func(_, v gjson.Result) bool {
			key := v.Get("{{$childClassName}}.attributes.rn").String()
			if key == "{{$childRn}}" {
				r = v
				return false
			}
			return true
		},
	)
	{{- range .Attributes}}
	{{- if and (not .ReferenceOnly) (not .WriteOnly)}}
	if !data.{{toGoName .TfName}}.IsNull() || all {
		{{- if eq .Type "Int64"}}
		data.{{toGoName .TfName}} = types.Int64Value(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").Int())
		{{- else if eq .Type "Bool"}}
		data.{{toGoName .TfName}} = types.BoolValue(helpers.ParseNxosBoolean(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").String()))
		{{- else if eq .Type "String"}}
		data.{{toGoName .TfName}} = types.StringValue(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").String())
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
	}
	{{- end}}
	{{- end}}
	{{- else if eq .Type "list"}}
	if all {
		res.Get(data.getClassName() + ".children").ForEach(
			func(_, v gjson.Result) bool {
				v.ForEach(
					func(classname, value gjson.Result) bool {
						if classname.String() == "{{$childClassName}}" {
							var child {{$name}}{{toGoName .TfName}}
							{{- range .Attributes}}
							{{- if and (not .ReferenceOnly) (not .WriteOnly)}}
							{{- if eq .Type "Int64"}}
							child.{{toGoName .TfName}} = types.Int64Value(value.Get("attributes.{{.NxosName}}").Int())
							{{- else if eq .Type "Bool"}}
							child.{{toGoName .TfName}} = types.BoolValue(helpers.ParseNxosBoolean(value.Get("attributes.{{.NxosName}}").String()))
							{{- else if eq .Type "String"}}
							child.{{toGoName .TfName}} = types.StringValue(value.Get("attributes.{{.NxosName}}").String())
							{{- end}}
							{{- end}}
							{{- end}}
							data.{{$list}} = append(data.{{$list}}, child)
						}
						return true
					},
				)
				return true
			},
		)
	} else {
		for c := range data.{{toGoName .TfName}} {
			var r gjson.Result
			res.Get(data.getClassName() + ".children").ForEach(
				func(_, v gjson.Result) bool {
					key := v.Get("{{$childClassName}}.attributes.rn").String()
					if key == data.{{$list}}[c].getRn() {
						r = v
						return false
					}
					return true
				},
			)
			{{- range .Attributes}}
			{{- if and (not .ReferenceOnly) (not .WriteOnly)}}
			if !data.{{$list}}[c].{{toGoName .TfName}}.IsNull() || all {
				{{- if eq .Type "Int64"}}
				data.{{$list}}[c].{{toGoName .TfName}} = types.Int64Value(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").Int())
				{{- else if eq .Type "Bool"}}
				data.{{$list}}[c].{{toGoName .TfName}} = types.BoolValue(helpers.ParseNxosBoolean(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").String()))
				{{- else if eq .Type "String"}}
				data.{{$list}}[c].{{toGoName .TfName}} = types.StringValue(r.Get("{{$childClassName}}.attributes.{{.NxosName}}").String())
				{{- end}}
			} else {
				data.{{$list}}[c].{{toGoName .TfName}} = types.{{.Type}}Null()
			}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
}

func (data {{camelCase .Name}}) toDeleteBody() nxos.Body {
	body := ""

	{{- range .Attributes}}
	{{- if and (not .ReferenceOnly) (.DeleteValue)}}
	{{- if eq .Type "Int64"}}
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", strconv.FormatInt({{ .DeleteValue}}, 10))
	{{- else if eq .Type "Bool"}}
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", strconv.FormatBool({{ .DeleteValue}}))
	{{- else if eq .Type "String"}}
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"{{.NxosName}}", "{{ .DeleteValue}}")
	{{- end}}
	{{- end}}
	{{- end}}

	return nxos.Body{body}
}

{{if hasId .Attributes}}
func (data *{{camelCase .Name}}) getIdsFromDn() {
{{- range .Attributes}}
{{- if .Id}}
	var {{toGoName .TfName}} {{if eq .Type "Int64"}}int64{{else}}string{{end}}
{{- end}}
{{- end}}
	fmt.Sscanf(data.Dn.ValueString(), "{{.Dn}}"{{range .Attributes}}{{if .Id}}, &{{toGoName .TfName}}{{end}}{{end}})
{{- range .Attributes}}
{{- if .Id}}
	data.{{toGoName .TfName}} = types.{{.Type}}Value({{toGoName .TfName}})
{{- end}}
{{- end}}
}
{{end}}

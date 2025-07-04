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

//template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type ActiveDirectoryGroupsByDomain struct {
	JoinPointId types.String                          `tfsdk:"join_point_id"`
	Domain      types.String                          `tfsdk:"domain"`
	Filter      types.String                          `tfsdk:"filter"`
	SidFilter   types.String                          `tfsdk:"sid_filter"`
	TypeFilter  types.String                          `tfsdk:"type_filter"`
	Groups      []ActiveDirectoryGroupsByDomainGroups `tfsdk:"groups"`
}

type ActiveDirectoryGroupsByDomainGroups struct {
	Name types.String `tfsdk:"name"`
	Sid  types.String `tfsdk:"sid"`
	Type types.String `tfsdk:"type"`
}

//template:end types

//template:begin getPath
func (data ActiveDirectoryGroupsByDomain) getPath() string {
	return fmt.Sprintf("/ers/config/activedirectory/%v/getGroupsByDomain", url.QueryEscape(data.JoinPointId.ValueString()))
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data ActiveDirectoryGroupsByDomain) toBody(ctx context.Context, state ActiveDirectoryGroupsByDomain) string {
	body := ""
	if !data.Domain.IsNull() {
		body, _ = sjson.Set(body, "Domain", data.Domain.ValueString())
	}
	if !data.Filter.IsNull() {
		body, _ = sjson.Set(body, "filter", data.Filter.ValueString())
	}
	if !data.SidFilter.IsNull() {
		body, _ = sjson.Set(body, "sidFilter", data.SidFilter.ValueString())
	}
	if !data.TypeFilter.IsNull() {
		body, _ = sjson.Set(body, "typeFilter", data.TypeFilter.ValueString())
	}
	if len(data.Groups) > 0 {
		body, _ = sjson.Set(body, "ERSActiveDirectoryGroups.groups", []interface{}{})
		for _, item := range data.Groups {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Sid.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sid", item.Sid.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ERSActiveDirectoryGroups.groups.-1", itemBody)
		}
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *ActiveDirectoryGroupsByDomain) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("Domain"); value.Exists() && value.Type != gjson.Null {
		data.Domain = types.StringValue(value.String())
	} else {
		data.Domain = types.StringNull()
	}
	if value := res.Get("filter"); value.Exists() && value.Type != gjson.Null {
		data.Filter = types.StringValue(value.String())
	} else {
		data.Filter = types.StringNull()
	}
	if value := res.Get("sidFilter"); value.Exists() && value.Type != gjson.Null {
		data.SidFilter = types.StringValue(value.String())
	} else {
		data.SidFilter = types.StringNull()
	}
	if value := res.Get("typeFilter"); value.Exists() && value.Type != gjson.Null {
		data.TypeFilter = types.StringValue(value.String())
	} else {
		data.TypeFilter = types.StringNull()
	}
	if value := res.Get("ERSActiveDirectoryGroups.groups"); value.Exists() {
		data.Groups = make([]ActiveDirectoryGroupsByDomainGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ActiveDirectoryGroupsByDomainGroups{}
			if cValue := v.Get("name"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("sid"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Sid = types.StringValue(cValue.String())
			} else {
				item.Sid = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() && cValue.Type != gjson.Null {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			data.Groups = append(data.Groups, item)
			return true
		})
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *ActiveDirectoryGroupsByDomain) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("Domain"); value.Exists() && !data.Domain.IsNull() {
		data.Domain = types.StringValue(value.String())
	} else {
		data.Domain = types.StringNull()
	}
	if value := res.Get("filter"); value.Exists() && !data.Filter.IsNull() {
		data.Filter = types.StringValue(value.String())
	} else {
		data.Filter = types.StringNull()
	}
	if value := res.Get("sidFilter"); value.Exists() && !data.SidFilter.IsNull() {
		data.SidFilter = types.StringValue(value.String())
	} else {
		data.SidFilter = types.StringNull()
	}
	if value := res.Get("typeFilter"); value.Exists() && !data.TypeFilter.IsNull() {
		data.TypeFilter = types.StringValue(value.String())
	} else {
		data.TypeFilter = types.StringNull()
	}
	for i := range data.Groups {
		keys := [...]string{"name", "sid", "type"}
		keyValues := [...]string{data.Groups[i].Name.ValueString(), data.Groups[i].Sid.ValueString(), data.Groups[i].Type.ValueString()}

		var r gjson.Result
		res.Get("ERSActiveDirectoryGroups.groups").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.Groups[i].Name.IsNull() {
			data.Groups[i].Name = types.StringValue(value.String())
		} else {
			data.Groups[i].Name = types.StringNull()
		}
		if value := r.Get("sid"); value.Exists() && !data.Groups[i].Sid.IsNull() {
			data.Groups[i].Sid = types.StringValue(value.String())
		} else {
			data.Groups[i].Sid = types.StringNull()
		}
		if value := r.Get("type"); value.Exists() && !data.Groups[i].Type.IsNull() {
			data.Groups[i].Type = types.StringValue(value.String())
		} else {
			data.Groups[i].Type = types.StringNull()
		}
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *ActiveDirectoryGroupsByDomain) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Domain.IsNull() {
		return false
	}
	if !data.Filter.IsNull() {
		return false
	}
	if !data.SidFilter.IsNull() {
		return false
	}
	if !data.TypeFilter.IsNull() {
		return false
	}
	if len(data.Groups) > 0 {
		return false
	}
	return true
}

//template:end isNull

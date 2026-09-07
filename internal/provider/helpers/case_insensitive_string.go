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

package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// CaseInsensitiveStringType is a custom string type for attributes backing the
// definition flag `case_insensitive: true`. It is used for Required/non-Computed
// String attributes where the API always normalizes casing server-side (e.g. ISE
// always returns MAC addresses in uppercase, regardless of the case submitted).
//
// Terraform core requires a provider's planned value for a Required, non-Computed
// attribute to be byte-identical to the config value, so a plan modifier can never
// rewrite it (doing so produces a "Provider produced invalid plan" error). Instead,
// CaseInsensitiveStringType/CaseInsensitiveStringValue implement semantic equality:
// a case-only difference between the prior state and a freshly-read API value (during
// refresh) or between state and the planned value (during the post-apply consistency
// check) is treated as "no change", so Terraform keeps the practitioner's original
// casing in state rather than drifting to the API's casing.
var (
	_ basetypes.StringTypable                    = CaseInsensitiveStringType{}
	_ basetypes.StringValuableWithSemanticEquals = CaseInsensitiveStringValue{}
)

type CaseInsensitiveStringType struct {
	basetypes.StringType
}

func (t CaseInsensitiveStringType) Equal(o attr.Type) bool {
	other, ok := o.(CaseInsensitiveStringType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t CaseInsensitiveStringType) String() string {
	return "helpers.CaseInsensitiveStringType"
}

func (t CaseInsensitiveStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return CaseInsensitiveStringValue{StringValue: in}, nil
}

func (t CaseInsensitiveStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	stringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}
	stringValuable, diags := t.ValueFromString(ctx, stringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}
	return stringValuable, nil
}

func (t CaseInsensitiveStringType) ValueType(ctx context.Context) attr.Value {
	return CaseInsensitiveStringValue{}
}

// CaseInsensitiveStringValue is the value type produced by CaseInsensitiveStringType.
type CaseInsensitiveStringValue struct {
	basetypes.StringValue
}

func (v CaseInsensitiveStringValue) Equal(o attr.Value) bool {
	other, ok := o.(CaseInsensitiveStringValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

func (v CaseInsensitiveStringValue) Type(ctx context.Context) attr.Type {
	return CaseInsensitiveStringType{}
}

// StringSemanticEquals reports two values as equal when they are identical ignoring
// case, so the API's canonical casing does not perpetually drift against a
// differently-cased value the practitioner supplied in configuration.
func (v CaseInsensitiveStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newValue, ok := newValuable.(CaseInsensitiveStringValue)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			fmt.Sprintf("expected value type of helpers.CaseInsensitiveStringValue but got %T", newValuable),
		)
		return false, diags
	}
	if v.IsNull() || v.IsUnknown() || newValue.IsNull() || newValue.IsUnknown() {
		return v.StringValue.Equal(newValue.StringValue), diags
	}
	return strings.EqualFold(v.ValueString(), newValue.ValueString()), diags
}

// NewCaseInsensitiveStringNull returns a null CaseInsensitiveStringValue.
func NewCaseInsensitiveStringNull() CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{StringValue: basetypes.NewStringNull()}
}

// NewCaseInsensitiveStringValue returns a known CaseInsensitiveStringValue holding v.
func NewCaseInsensitiveStringValue(v string) CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{StringValue: basetypes.NewStringValue(v)}
}

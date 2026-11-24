// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates.
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type boolEnabledWhen struct{ parent string }

// BoolEnabledWhen preserves an omitted switch while its parent remains enabled.
// Disabled parents make their dependent switches effectively false.
func BoolEnabledWhen(parent string) planmodifier.Bool { return boolEnabledWhen{parent: parent} }

func (m boolEnabledWhen) Description(context.Context) string {
	return fmt.Sprintf("Can only be enabled when %s is true.", m.parent)
}
func (m boolEnabledWhen) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }

func (m boolEnabledWhen) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var parent types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root(m.parent), &parent)...)
	if !resp.Diagnostics.HasError() && parent.IsUnknown() && !req.State.Raw.IsNull() {
		var configured types.Bool
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root(m.parent), &configured)...)
		if !resp.Diagnostics.HasError() && configured.IsNull() {
			resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(m.parent), &parent)...)
		}
	}
	if resp.Diagnostics.HasError() || parent.IsUnknown() || parent.IsNull() {
		return
	}
	if !parent.ValueBool() {
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && req.ConfigValue.ValueBool() {
			resp.Diagnostics.AddAttributeError(req.Path, "Inactive setting", fmt.Sprintf("This setting can only be true when %s is true.", m.parent))
		} else if req.ConfigValue.IsNull() {
			resp.PlanValue = types.BoolValue(false)
		}
		return
	}
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.StateValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	var previous types.Bool
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(m.parent), &previous)...)
	if !resp.Diagnostics.HasError() && parent.Equal(previous) {
		resp.PlanValue = req.StateValue
	}
}

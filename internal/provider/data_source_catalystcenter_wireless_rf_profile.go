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

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessRFProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRFProfileDataSource{}
)

func NewWirelessRFProfileDataSource() datasource.DataSource {
	return &WirelessRFProfileDataSource{}
}

type WirelessRFProfileDataSource struct {
	client *cc.Client
}

func (d *WirelessRFProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profile"
}

func (d *WirelessRFProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless RF Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"rf_profile_name": schema.StringAttribute{
				MarkdownDescription: "RF Profile Name",
				Computed:            true,
			},
			"default_rf_profile": schema.BoolAttribute{
				MarkdownDescription: "is Default Rf Profile",
				Computed:            true,
			},
			"enable_radio_type_a": schema.BoolAttribute{
				MarkdownDescription: "True if 5 GHz radio band is enabled in the RF Profile, else False",
				Computed:            true,
			},
			"enable_radio_type_b": schema.BoolAttribute{
				MarkdownDescription: "True if 2.4 GHz radio band is enabled in the RF Profile, else False",
				Computed:            true,
			},
			"enable_radio_type6_g_hz": schema.BoolAttribute{
				MarkdownDescription: "True if 6 GHz radio band is enabled in the RF Profile, else False",
				Computed:            true,
			},
			"radio_type_a_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Parent profile of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - DCA channels of 5 GHz radio band passed in comma separated format without any spaces",
				Computed:            true,
			},
			"radio_type_a_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Data rates of 5 GHz radio band passed in comma separated format without any spaces",
				Computed:            true,
			},
			"radio_type_a_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values",
				Computed:            true,
			},
			"radio_type_a_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Power threshold of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - RX-SOP threshold of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Minimum power level of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Maximum power level of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_channel_width": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Channel Width",
				Computed:            true,
			},
			"radio_type_a_preamble_puncture": schema.BoolAttribute{
				MarkdownDescription: "Radio TypeA Properties - Enable or Disable Preamble Puncturing",
				Computed:            true,
			},
			"radio_type_a_zero_wait_dfs_enable": schema.BoolAttribute{
				MarkdownDescription: "Radio TypeA Properties - Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions",
				Computed:            true,
			},
			"radio_type_a_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Custom RX-SOP threshold of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Client Limit of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_aware": schema.BoolAttribute{
				MarkdownDescription: "Radio TypeA Properties - Client Aware of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_select": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Client Select(%) of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_reset": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Client Reset(%) of 5 GHz radio band",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Coverage Hole Detection Client Level",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Coverage Hole Detection Data Rssi Threshold",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Coverage Hole Detection Voice Rssi Threshold",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Coverage Hole Detection Exception Level(%)",
				Computed:            true,
			},
			"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Radio TypeA Properties - Dot11ax Non SRG OBSS PD",
				Computed:            true,
			},
			"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Dot11ax Non SRG OBSS PD Max Threshold",
				Computed:            true,
			},
			"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Radio TypeA Properties - Dot11ax SRG OBSS PD",
				Computed:            true,
			},
			"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Dot11ax SRG OBSS PD Min Threshold",
				Computed:            true,
			},
			"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Dot11ax SRG OBSS PD Max Threshold",
				Computed:            true,
			},
			"radio_type_b_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Parent Profile",
				Computed:            true,
			},
			"radio_type_b_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Radio Channels",
				Computed:            true,
			},
			"radio_type_b_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Data Rates",
				Computed:            true,
			},
			"radio_type_b_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Mandatory Data Rates",
				Computed:            true,
			},
			"radio_type_b_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Power Threshold V1",
				Computed:            true,
			},
			"radio_type_b_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Rx Sop Threshold",
				Computed:            true,
			},
			"radio_type_b_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Min Power Level",
				Computed:            true,
			},
			"radio_type_b_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Max Power Level",
				Computed:            true,
			},
			"radio_type_c_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Parent Profile",
				Computed:            true,
			},
			"radio_type_c_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Radio Channels",
				Computed:            true,
			},
			"radio_type_c_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Data Rates",
				Computed:            true,
			},
			"radio_type_c_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Mandatory Data Rates",
				Computed:            true,
			},
			"radio_type_c_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Power Threshold V1",
				Computed:            true,
			},
			"radio_type_c_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Rx Sop Threshold",
				Computed:            true,
			},
			"radio_type_c_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Min Power Level",
				Computed:            true,
			},
			"radio_type_c_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Max Power Level",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessRFProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessRFProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessRFProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

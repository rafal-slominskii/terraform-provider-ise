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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessRFProfile struct {
	Id                                                                         types.String `tfsdk:"id"`
	RfProfileName                                                              types.String `tfsdk:"rf_profile_name"`
	DefaultRfProfile                                                           types.Bool   `tfsdk:"default_rf_profile"`
	EnableRadioTypeA                                                           types.Bool   `tfsdk:"enable_radio_type_a"`
	EnableRadioTypeB                                                           types.Bool   `tfsdk:"enable_radio_type_b"`
	EnableRadioType6GHz                                                        types.Bool   `tfsdk:"enable_radio_type6_g_hz"`
	RadioTypeAParentProfile                                                    types.String `tfsdk:"radio_type_a_parent_profile"`
	RadioTypeARadioChannels                                                    types.String `tfsdk:"radio_type_a_radio_channels"`
	RadioTypeADataRates                                                        types.String `tfsdk:"radio_type_a_data_rates"`
	RadioTypeAMandatoryDataRates                                               types.String `tfsdk:"radio_type_a_mandatory_data_rates"`
	RadioTypeAPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_a_power_threshold_v1"`
	RadioTypeARxSopThreshold                                                   types.String `tfsdk:"radio_type_a_rx_sop_threshold"`
	RadioTypeAMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_a_min_power_level"`
	RadioTypeAMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_a_max_power_level"`
	RadioTypeAChannelWidth                                                     types.String `tfsdk:"radio_type_a_channel_width"`
	RadioTypeAPreamblePuncture                                                 types.Bool   `tfsdk:"radio_type_a_preamble_puncture"`
	RadioTypeAZeroWaitDfsEnable                                                types.Bool   `tfsdk:"radio_type_a_zero_wait_dfs_enable"`
	RadioTypeACustomRxSopThreshold                                             types.Int64  `tfsdk:"radio_type_a_custom_rx_sop_threshold"`
	RadioTypeAMaxRadioClients                                                  types.Int64  `tfsdk:"radio_type_a_max_radio_clients"`
	RadioTypeAFraPropertiesClientAware                                         types.Bool   `tfsdk:"radio_type_a_fra_properties_client_aware"`
	RadioTypeAFraPropertiesClientSelect                                        types.Int64  `tfsdk:"radio_type_a_fra_properties_client_select"`
	RadioTypeAFraPropertiesClientReset                                         types.Int64  `tfsdk:"radio_type_a_fra_properties_client_reset"`
	RadioTypeACoverageHoleDetectionPropertiesChdClientLevel                    types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_client_level"`
	RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold              types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold"`
	RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold             types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold"`
	RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel                 types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_exception_level"`
	RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect             types.Bool   `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect"`
	RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect                types.Bool   `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold    types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold    types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold"`
	RadioTypeBParentProfile                                                    types.String `tfsdk:"radio_type_b_parent_profile"`
	RadioTypeBRadioChannels                                                    types.String `tfsdk:"radio_type_b_radio_channels"`
	RadioTypeBDataRates                                                        types.String `tfsdk:"radio_type_b_data_rates"`
	RadioTypeBMandatoryDataRates                                               types.String `tfsdk:"radio_type_b_mandatory_data_rates"`
	RadioTypeBPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_b_power_threshold_v1"`
	RadioTypeBRxSopThreshold                                                   types.String `tfsdk:"radio_type_b_rx_sop_threshold"`
	RadioTypeBMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_b_min_power_level"`
	RadioTypeBMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_b_max_power_level"`
	RadioTypeCParentProfile                                                    types.String `tfsdk:"radio_type_c_parent_profile"`
	RadioTypeCRadioChannels                                                    types.String `tfsdk:"radio_type_c_radio_channels"`
	RadioTypeCDataRates                                                        types.String `tfsdk:"radio_type_c_data_rates"`
	RadioTypeCMandatoryDataRates                                               types.String `tfsdk:"radio_type_c_mandatory_data_rates"`
	RadioTypeCPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_c_power_threshold_v1"`
	RadioTypeCRxSopThreshold                                                   types.String `tfsdk:"radio_type_c_rx_sop_threshold"`
	RadioTypeCMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_c_min_power_level"`
	RadioTypeCMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_c_max_power_level"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessRFProfile) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/rfProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessRFProfile) toBody(ctx context.Context, state WirelessRFProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.RfProfileName.IsNull() {
		body, _ = sjson.Set(body, "rfProfileName", data.RfProfileName.ValueString())
	}
	if !data.DefaultRfProfile.IsNull() {
		body, _ = sjson.Set(body, "defaultRfProfile", data.DefaultRfProfile.ValueBool())
	}
	if !data.EnableRadioTypeA.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeA", data.EnableRadioTypeA.ValueBool())
	}
	if !data.EnableRadioTypeB.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeB", data.EnableRadioTypeB.ValueBool())
	}
	if !data.EnableRadioType6GHz.IsNull() {
		body, _ = sjson.Set(body, "enableRadioType6GHz", data.EnableRadioType6GHz.ValueBool())
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.parentProfile", data.RadioTypeAParentProfile.ValueString())
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.radioChannels", data.RadioTypeARadioChannels.ValueString())
	}
	if !data.RadioTypeADataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.dataRates", data.RadioTypeADataRates.ValueString())
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.mandatoryDataRates", data.RadioTypeAMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeAPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.powerThresholdV1", data.RadioTypeAPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.rxSopThreshold", data.RadioTypeARxSopThreshold.ValueString())
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.minPowerLevel", data.RadioTypeAMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.maxPowerLevel", data.RadioTypeAMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeAChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.channelWidth", data.RadioTypeAChannelWidth.ValueString())
	}
	if !data.RadioTypeAPreamblePuncture.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.preamblePuncture", data.RadioTypeAPreamblePuncture.ValueBool())
	}
	if !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.zeroWaitDfsEnable", data.RadioTypeAZeroWaitDfsEnable.ValueBool())
	}
	if !data.RadioTypeACustomRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.customRxSopThreshold", data.RadioTypeACustomRxSopThreshold.ValueInt64())
	}
	if !data.RadioTypeAMaxRadioClients.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.maxRadioClients", data.RadioTypeAMaxRadioClients.ValueInt64())
	}
	if !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientAware", data.RadioTypeAFraPropertiesClientAware.ValueBool())
	}
	if !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientSelect", data.RadioTypeAFraPropertiesClientSelect.ValueInt64())
	}
	if !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientReset", data.RadioTypeAFraPropertiesClientReset.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel", data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold", data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold", data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel", data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect", data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold", data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.parentProfile", data.RadioTypeBParentProfile.ValueString())
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.radioChannels", data.RadioTypeBRadioChannels.ValueString())
	}
	if !data.RadioTypeBDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.dataRates", data.RadioTypeBDataRates.ValueString())
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.mandatoryDataRates", data.RadioTypeBMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeBPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.powerThresholdV1", data.RadioTypeBPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.rxSopThreshold", data.RadioTypeBRxSopThreshold.ValueString())
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.minPowerLevel", data.RadioTypeBMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.maxPowerLevel", data.RadioTypeBMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.parentProfile", data.RadioTypeCParentProfile.ValueString())
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.radioChannels", data.RadioTypeCRadioChannels.ValueString())
	}
	if !data.RadioTypeCDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.dataRates", data.RadioTypeCDataRates.ValueString())
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.mandatoryDataRates", data.RadioTypeCMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeCPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.powerThresholdV1", data.RadioTypeCPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.rxSopThreshold", data.RadioTypeCRxSopThreshold.ValueString())
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.minPowerLevel", data.RadioTypeCMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.maxPowerLevel", data.RadioTypeCMaxPowerLevel.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessRFProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rfProfileName"); value.Exists() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("defaultRfProfile"); value.Exists() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("enableRadioTypeA"); value.Exists() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("enableRadioTypeB"); value.Exists() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("enableRadioType6GHz"); value.Exists() {
		data.EnableRadioType6GHz = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioType6GHz = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.parentProfile"); value.Exists() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.radioChannels"); value.Exists() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.dataRates"); value.Exists() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeAPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.channelWidth"); value.Exists() {
		data.RadioTypeAChannelWidth = types.StringValue(value.String())
	} else {
		data.RadioTypeAChannelWidth = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.preamblePuncture"); value.Exists() {
		data.RadioTypeAPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.zeroWaitDfsEnable"); value.Exists() {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.customRxSopThreshold"); value.Exists() {
		data.RadioTypeACustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.maxRadioClients"); value.Exists() {
		data.RadioTypeAMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientAware"); value.Exists() {
		data.RadioTypeAFraPropertiesClientAware = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAFraPropertiesClientAware = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientSelect"); value.Exists() {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientReset"); value.Exists() {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.parentProfile"); value.Exists() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.radioChannels"); value.Exists() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.dataRates"); value.Exists() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeBPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.parentProfile"); value.Exists() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.radioChannels"); value.Exists() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.dataRates"); value.Exists() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeCPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessRFProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rfProfileName"); value.Exists() && !data.RfProfileName.IsNull() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("defaultRfProfile"); value.Exists() && !data.DefaultRfProfile.IsNull() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("enableRadioTypeA"); value.Exists() && !data.EnableRadioTypeA.IsNull() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("enableRadioTypeB"); value.Exists() && !data.EnableRadioTypeB.IsNull() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("enableRadioType6GHz"); value.Exists() && !data.EnableRadioType6GHz.IsNull() {
		data.EnableRadioType6GHz = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioType6GHz = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.parentProfile"); value.Exists() && !data.RadioTypeAParentProfile.IsNull() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.radioChannels"); value.Exists() && !data.RadioTypeARadioChannels.IsNull() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.dataRates"); value.Exists() && !data.RadioTypeADataRates.IsNull() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeAMandatoryDataRates.IsNull() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeAPowerThresholdV1.IsNull() {
		data.RadioTypeAPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeARxSopThreshold.IsNull() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.minPowerLevel"); value.Exists() && !data.RadioTypeAMinPowerLevel.IsNull() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeAMaxPowerLevel.IsNull() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.channelWidth"); value.Exists() && !data.RadioTypeAChannelWidth.IsNull() {
		data.RadioTypeAChannelWidth = types.StringValue(value.String())
	} else {
		data.RadioTypeAChannelWidth = types.StringNull()
	}
	if value := res.Get("radioTypeAProperties.preamblePuncture"); value.Exists() && !data.RadioTypeAPreamblePuncture.IsNull() {
		data.RadioTypeAPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.zeroWaitDfsEnable"); value.Exists() && !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.customRxSopThreshold"); value.Exists() && !data.RadioTypeACustomRxSopThreshold.IsNull() {
		data.RadioTypeACustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.maxRadioClients"); value.Exists() && !data.RadioTypeAMaxRadioClients.IsNull() {
		data.RadioTypeAMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientAware"); value.Exists() && !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		data.RadioTypeAFraPropertiesClientAware = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAFraPropertiesClientAware = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientSelect"); value.Exists() && !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.fraPropertiesA.clientReset"); value.Exists() && !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.parentProfile"); value.Exists() && !data.RadioTypeBParentProfile.IsNull() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.radioChannels"); value.Exists() && !data.RadioTypeBRadioChannels.IsNull() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.dataRates"); value.Exists() && !data.RadioTypeBDataRates.IsNull() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeBMandatoryDataRates.IsNull() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeBPowerThresholdV1.IsNull() {
		data.RadioTypeBPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeBRxSopThreshold.IsNull() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.minPowerLevel"); value.Exists() && !data.RadioTypeBMinPowerLevel.IsNull() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeBMaxPowerLevel.IsNull() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.parentProfile"); value.Exists() && !data.RadioTypeCParentProfile.IsNull() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.radioChannels"); value.Exists() && !data.RadioTypeCRadioChannels.IsNull() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.dataRates"); value.Exists() && !data.RadioTypeCDataRates.IsNull() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeCMandatoryDataRates.IsNull() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeCPowerThresholdV1.IsNull() {
		data.RadioTypeCPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeCRxSopThreshold.IsNull() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.minPowerLevel"); value.Exists() && !data.RadioTypeCMinPowerLevel.IsNull() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeCMaxPowerLevel.IsNull() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessRFProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.RfProfileName.IsNull() {
		return false
	}
	if !data.DefaultRfProfile.IsNull() {
		return false
	}
	if !data.EnableRadioTypeA.IsNull() {
		return false
	}
	if !data.EnableRadioTypeB.IsNull() {
		return false
	}
	if !data.EnableRadioType6GHz.IsNull() {
		return false
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeADataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeAChannelWidth.IsNull() {
		return false
	}
	if !data.RadioTypeAPreamblePuncture.IsNull() {
		return false
	}
	if !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		return false
	}
	if !data.RadioTypeACustomRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeAMaxRadioClients.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeBDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeCDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull

// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates.
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func certificateDER(value string) ([]byte, error) {
	value = strings.TrimSpace(value)
	block, rest := pem.Decode([]byte(value))
	if !strings.HasPrefix(value, "-----BEGIN CERTIFICATE-----") || block == nil || block.Type != "CERTIFICATE" || len(block.Headers) != 0 || len(bytes.TrimSpace(rest)) != 0 || strings.Count(value, "-----BEGIN ") != 1 {
		return nil, fmt.Errorf("provide exactly one PEM-encoded X.509 certificate, without a private key or certificate chain")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("invalid X.509 certificate: %w", err)
	}
	return cert.Raw, nil
}

type certificateValidator struct{}

func CertificatePEM() validator.String { return certificateValidator{} }

func (certificateValidator) Description(context.Context) string {
	return "Value must contain exactly one PEM-encoded X.509 certificate."
}

func (v certificateValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (certificateValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if _, err := certificateDER(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid certificate", err.Error())
	}
}

type certificateReplace struct{ fingerprintAttribute string }

// CertificateRequiresReplace compares public certificate identity, not PEM formatting.
// After terraform import, the PEM is absent from state; compare against ISE's fingerprint.
func CertificateRequiresReplace(fingerprintAttribute string) planmodifier.String {
	return certificateReplace{fingerprintAttribute: fingerprintAttribute}
}

func (certificateReplace) Description(context.Context) string {
	return "Requires replacement if the X.509 certificate DER identity changes."
}

func (m certificateReplace) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m certificateReplace) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}
	planned, err := certificateDER(req.PlanValue.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid certificate", err.Error())
		return
	}
	if !req.StateValue.IsNull() {
		previous, err := certificateDER(req.StateValue.ValueString())
		resp.RequiresReplace = err != nil || !bytes.Equal(previous, planned)
		return
	}
	var fingerprint types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(m.fingerprintAttribute), &fingerprint)...)
	if resp.Diagnostics.HasError() {
		return
	}
	normalize := strings.NewReplacer(":", "", " ", "", "-", "")
	wanted := strings.ToLower(normalize.Replace(fingerprint.ValueString()))
	sum := sha256.Sum256(planned)
	resp.RequiresReplace = wanted != hex.EncodeToString(sum[:])
}

type certificateFingerprintState struct{ certificateAttribute string }

// CertificateFingerprintUseState keeps the output known only when the planned
// certificate hashes to the current fingerprint. Replacement inputs remain unknown.
func CertificateFingerprintUseState(certificateAttribute string) planmodifier.String {
	return certificateFingerprintState{certificateAttribute: certificateAttribute}
}

func (certificateFingerprintState) Description(context.Context) string {
	return "Preserve the fingerprint when the certificate DER identity is unchanged."
}

func (m certificateFingerprintState) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m certificateFingerprintState) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	var certificate types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root(m.certificateAttribute), &certificate)...)
	if resp.Diagnostics.HasError() || certificate.IsNull() || certificate.IsUnknown() {
		return
	}
	der, err := certificateDER(certificate.ValueString())
	if err != nil {
		return
	} // The certificate input validator reports invalid PEM.
	sum := sha256.Sum256(der)
	current := strings.ToLower(strings.NewReplacer(":", "", " ", "", "-", "").Replace(req.StateValue.ValueString()))
	if current == hex.EncodeToString(sum[:]) {
		resp.PlanValue = req.StateValue
	}
}

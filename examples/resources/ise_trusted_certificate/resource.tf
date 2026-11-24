resource "ise_trusted_certificate" "example" {
  name                                   = "NaC Test CA"
  description                            = "Root CA for ISE authentication"
  certificate_pem                        = file("ca.pem")
  trust_for_ise_auth                     = true
  trust_for_client_auth                  = true
  trust_for_certificate_based_admin_auth = false
  trust_for_cisco_services_auth          = false
  allow_basic_constraint_ca_false        = false
  allow_multiple_cn_cert                 = false
  allow_out_of_date_cert                 = false
  allow_sha1_certificates                = false
  validate_certificate_extensions        = true
}

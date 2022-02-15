// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WellKnown WellKnown WellKnown represents important OpenID Connect discovery metadata
//
// It includes links to several endpoints (e.g. /oauth2/token) and exposes information on supported signature algorithms
// among others.
//
// swagger:model wellKnown
type WellKnown struct {

	// URL of the OP's OAuth 2.0 Authorization Endpoint.
	// Example: https://playground.ory.sh/ory-hydra/public/oauth2/auth
	// Required: true
	AuthorizationEndpoint *string `json:"authorization_endpoint"`

	// Boolean value specifying whether the OP can pass a sid (session ID) Claim in the Logout Token to identify the RP
	// session with the OP. If supported, the sid Claim is also included in ID Tokens issued by the OP
	BackchannelLogoutSessionSupported bool `json:"backchannel_logout_session_supported,omitempty"`

	// Boolean value specifying whether the OP supports back-channel logout, with true indicating support.
	BackchannelLogoutSupported bool `json:"backchannel_logout_supported,omitempty"`

	// Boolean value specifying whether the OP supports use of the claims parameter, with true indicating support.
	ClaimsParameterSupported bool `json:"claims_parameter_supported,omitempty"`

	// JSON array containing a list of the Claim Names of the Claims that the OpenID Provider MAY be able to supply
	// values for. Note that for privacy or other reasons, this might not be an exhaustive list.
	ClaimsSupported []string `json:"claims_supported"`

	// JSON array containing a list of Proof Key for Code Exchange (PKCE) [RFC7636] code challenge methods supported
	// by this authorization server.
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`

	// URL at the OP to which an RP can perform a redirect to request that the End-User be logged out at the OP.
	EndSessionEndpoint string `json:"end_session_endpoint,omitempty"`

	// Boolean value specifying whether the OP can pass iss (issuer) and sid (session ID) query parameters to identify
	// the RP session with the OP when the frontchannel_logout_uri is used. If supported, the sid Claim is also
	// included in ID Tokens issued by the OP.
	FrontchannelLogoutSessionSupported bool `json:"frontchannel_logout_session_supported,omitempty"`

	// Boolean value specifying whether the OP supports HTTP-based logout, with true indicating support.
	FrontchannelLogoutSupported bool `json:"frontchannel_logout_supported,omitempty"`

	// JSON array containing a list of the OAuth 2.0 Grant Type values that this OP supports.
	GrantTypesSupported []string `json:"grant_types_supported"`

	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the OP for the ID Token
	// to encode the Claims in a JWT.
	// Required: true
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`

	// URL using the https scheme with no query or fragment component that the OP asserts as its IssuerURL Identifier.
	// If IssuerURL discovery is supported , this value MUST be identical to the issuer value returned
	// by WebFinger. This also MUST be identical to the iss Claim value in ID Tokens issued from this IssuerURL.
	// Example: https://playground.ory.sh/ory-hydra/public/
	// Required: true
	Issuer *string `json:"issuer"`

	// URL of the OP's JSON Web Key Set [JWK] document. This contains the signing key(s) the RP uses to validate
	// signatures from the OP. The JWK Set MAY also contain the Server's encryption key(s), which are used by RPs
	// to encrypt requests to the Server. When both signing and encryption keys are made available, a use (Key Use)
	// parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage.
	// Although some algorithms allow the same key to be used for both signatures and encryption, doing so is
	// NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of
	// keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
	// Example: https://playground.ory.sh/ory-hydra/public/.well-known/jwks.json
	// Required: true
	JwksURI *string `json:"jwks_uri"`

	// URL of the OP's Dynamic Client Registration Endpoint.
	// Example: https://playground.ory.sh/ory-hydra/admin/client
	RegistrationEndpoint string `json:"registration_endpoint,omitempty"`

	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the OP for Request Objects,
	// which are described in Section 6.1 of OpenID Connect Core 1.0 [OpenID.Core]. These algorithms are used both when
	// the Request Object is passed by value (using the request parameter) and when it is passed by reference
	// (using the request_uri parameter).
	RequestObjectSigningAlgValuesSupported []string `json:"request_object_signing_alg_values_supported"`

	// Boolean value specifying whether the OP supports use of the request parameter, with true indicating support.
	RequestParameterSupported bool `json:"request_parameter_supported,omitempty"`

	// Boolean value specifying whether the OP supports use of the request_uri parameter, with true indicating support.
	RequestURIParameterSupported bool `json:"request_uri_parameter_supported,omitempty"`

	// Boolean value specifying whether the OP requires any request_uri values used to be pre-registered
	// using the request_uris registration parameter.
	RequireRequestURIRegistration bool `json:"require_request_uri_registration,omitempty"`

	// JSON array containing a list of the OAuth 2.0 response_mode values that this OP supports.
	ResponseModesSupported []string `json:"response_modes_supported"`

	// JSON array containing a list of the OAuth 2.0 response_type values that this OP supports. Dynamic OpenID
	// Providers MUST support the code, id_token, and the token id_token Response Type values.
	// Required: true
	ResponseTypesSupported []string `json:"response_types_supported"`

	// URL of the authorization server's OAuth 2.0 revocation endpoint.
	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`

	// SON array containing a list of the OAuth 2.0 [RFC6749] scope values that this server supports. The server MUST
	// support the openid scope value. Servers MAY choose not to advertise some supported scope values even when this parameter is used
	ScopesSupported []string `json:"scopes_supported"`

	// JSON array containing a list of the Subject Identifier types that this OP supports. Valid types include
	// pairwise and public.
	// Required: true
	SubjectTypesSupported []string `json:"subject_types_supported"`

	// URL of the OP's OAuth 2.0 Token Endpoint
	// Example: https://playground.ory.sh/ory-hydra/public/oauth2/token
	// Required: true
	TokenEndpoint *string `json:"token_endpoint"`

	// JSON array containing a list of Client Authentication methods supported by this Token Endpoint. The options are
	// client_secret_post, client_secret_basic, client_secret_jwt, and private_key_jwt, as described in Section 9 of OpenID Connect Core 1.0
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`

	// URL of the OP's UserInfo Endpoint.
	UserinfoEndpoint string `json:"userinfo_endpoint,omitempty"`

	// JSON array containing a list of the JWS [JWS] signing algorithms (alg values) [JWA] supported by the UserInfo Endpoint to encode the Claims in a JWT [JWT].
	UserinfoSigningAlgValuesSupported []string `json:"userinfo_signing_alg_values_supported"`
}

// Validate validates this well known
func (m *WellKnown) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizationEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenSigningAlgValuesSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwksURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTypesSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectTypesSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WellKnown) validateAuthorizationEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("authorization_endpoint", "body", m.AuthorizationEndpoint); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateIDTokenSigningAlgValuesSupported(formats strfmt.Registry) error {

	if err := validate.Required("id_token_signing_alg_values_supported", "body", m.IDTokenSigningAlgValuesSupported); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateJwksURI(formats strfmt.Registry) error {

	if err := validate.Required("jwks_uri", "body", m.JwksURI); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateResponseTypesSupported(formats strfmt.Registry) error {

	if err := validate.Required("response_types_supported", "body", m.ResponseTypesSupported); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateSubjectTypesSupported(formats strfmt.Registry) error {

	if err := validate.Required("subject_types_supported", "body", m.SubjectTypesSupported); err != nil {
		return err
	}

	return nil
}

func (m *WellKnown) validateTokenEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("token_endpoint", "body", m.TokenEndpoint); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this well known based on context it is used
func (m *WellKnown) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WellKnown) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WellKnown) UnmarshalBinary(b []byte) error {
	var res WellKnown
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

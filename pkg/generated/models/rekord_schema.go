// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RekordSchema Rekor Schema
//
// Schema for Rekord object
//
// swagger:model rekordSchema
type RekordSchema struct {

	// data
	// Required: true
	Data *RekordSchemaData `json:"data"`

	// Arbitrary content to be included in the verifiable entry in the transparency log
	ExtraData interface{} `json:"extraData,omitempty"`

	// signature
	// Required: true
	Signature *RekordSchemaSignature `json:"signature"`
}

// Validate validates this rekord schema
func (m *RekordSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchema) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RekordSchema) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rekord schema based on the context it is used
func (m *RekordSchema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchema) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RekordSchema) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {
		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordSchema) UnmarshalBinary(b []byte) error {
	var res RekordSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordSchemaData Information about the content associated with the entry
//
// swagger:model RekordSchemaData
type RekordSchemaData struct {

	// Specifies the content inline within the document
	Content string `json:"content,omitempty"`

	// hash
	Hash *RekordSchemaDataHash `json:"hash,omitempty"`

	// Specifies the location of the content; if this is specified, a hash value must also be provided
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord schema data
func (m *RekordSchemaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchemaData) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *RekordSchemaData) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rekord schema data based on the context it is used
func (m *RekordSchemaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchemaData) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {
		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordSchemaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordSchemaData) UnmarshalBinary(b []byte) error {
	var res RekordSchemaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordSchemaDataHash Specifies the hash algorithm and value for the content
//
// swagger:model RekordSchemaDataHash
type RekordSchemaDataHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the content
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rekord schema data hash
func (m *RekordSchemaDataHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rekordSchemaDataHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rekordSchemaDataHashTypeAlgorithmPropEnum = append(rekordSchemaDataHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// RekordSchemaDataHashAlgorithmSha256 captures enum value "sha256"
	RekordSchemaDataHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *RekordSchemaDataHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rekordSchemaDataHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RekordSchemaDataHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("data"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *RekordSchemaDataHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rekord schema data hash based on context it is used
func (m *RekordSchemaDataHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RekordSchemaDataHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordSchemaDataHash) UnmarshalBinary(b []byte) error {
	var res RekordSchemaDataHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordSchemaSignature Information about the detached signature associated with the entry
//
// swagger:model RekordSchemaSignature
type RekordSchemaSignature struct {

	// Specifies the content of the signature inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// Specifies the format of the signature
	// Enum: [pgp]
	Format string `json:"format,omitempty"`

	// public key
	PublicKey *RekordSchemaSignaturePublicKey `json:"publicKey,omitempty"`

	// Specifies the location of the signature
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord schema signature
func (m *RekordSchemaSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rekordSchemaSignatureTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pgp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rekordSchemaSignatureTypeFormatPropEnum = append(rekordSchemaSignatureTypeFormatPropEnum, v)
	}
}

const (

	// RekordSchemaSignatureFormatPgp captures enum value "pgp"
	RekordSchemaSignatureFormatPgp string = "pgp"
)

// prop value enum
func (m *RekordSchemaSignature) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rekordSchemaSignatureTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RekordSchemaSignature) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("signature"+"."+"format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *RekordSchemaSignature) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

func (m *RekordSchemaSignature) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("signature"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rekord schema signature based on the context it is used
func (m *RekordSchemaSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchemaSignature) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {
		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RekordSchemaSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordSchemaSignature) UnmarshalBinary(b []byte) error {
	var res RekordSchemaSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RekordSchemaSignaturePublicKey The public key that can verify the signature
//
// swagger:model RekordSchemaSignaturePublicKey
type RekordSchemaSignaturePublicKey struct {

	// Specifies the content of the public key inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// Specifies the location of the public key
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rekord schema signature public key
func (m *RekordSchemaSignaturePublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RekordSchemaSignaturePublicKey) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("signature"+"."+"publicKey"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rekord schema signature public key based on context it is used
func (m *RekordSchemaSignaturePublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RekordSchemaSignaturePublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RekordSchemaSignaturePublicKey) UnmarshalBinary(b []byte) error {
	var res RekordSchemaSignaturePublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

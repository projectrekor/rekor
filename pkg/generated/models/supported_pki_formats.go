// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SupportedPKIFormats This represents the tokens that indicate the format of the PKI artifacts supported by the server
//
// swagger:model SupportedPKIFormats
type SupportedPKIFormats string

const (

	// SupportedPKIFormatsPgp captures enum value "pgp"
	SupportedPKIFormatsPgp SupportedPKIFormats = "pgp"
)

// for schema
var supportedPkiFormatsEnum []interface{}

func init() {
	var res []SupportedPKIFormats
	if err := json.Unmarshal([]byte(`["pgp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportedPkiFormatsEnum = append(supportedPkiFormatsEnum, v)
	}
}

func (m SupportedPKIFormats) validateSupportedPKIFormatsEnum(path, location string, value SupportedPKIFormats) error {
	if err := validate.EnumCase(path, location, value, supportedPkiFormatsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this supported PKI formats
func (m SupportedPKIFormats) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSupportedPKIFormatsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this supported PKI formats based on context it is used
func (m SupportedPKIFormats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

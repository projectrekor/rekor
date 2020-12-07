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

// LogEntry log entry
//
// swagger:model LogEntry
type LogEntry map[string]LogEntryAnon

// Validate validates this log entry
func (m LogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this log entry based on the context it is used
func (m LogEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// LogEntryAnon log entry anon
//
// swagger:model LogEntryAnon
type LogEntryAnon struct {

	// body
	// Required: true
	Body interface{} `json:"body"`

	// log index
	// Required: true
	// Minimum: 1
	LogIndex *int64 `json:"logIndex"`
}

// Validate validates this log entry anon
func (m *LogEntryAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogEntryAnon) validateBody(formats strfmt.Registry) error {

	if m.Body == nil {
		return errors.Required("body", "body", nil)
	}

	return nil
}

func (m *LogEntryAnon) validateLogIndex(formats strfmt.Registry) error {

	if err := validate.Required("logIndex", "body", m.LogIndex); err != nil {
		return err
	}

	if err := validate.MinimumInt("logIndex", "body", int64(*m.LogIndex), 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this log entry anon based on context it is used
func (m *LogEntryAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogEntryAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogEntryAnon) UnmarshalBinary(b []byte) error {
	var res LogEntryAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

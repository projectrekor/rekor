// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchLogQuery search log query
//
// swagger:model SearchLogQuery
type SearchLogQuery struct {
	entriesField []ProposedEntry

	// entry u UI ds
	EntryUUIDs []string `json:"entryUUIDs"`

	// log indexes
	LogIndexes []int64 `json:"logIndexes"`
}

// Entries gets the entries of this base type
func (m *SearchLogQuery) Entries() []ProposedEntry {
	return m.entriesField
}

// SetEntries sets the entries of this base type
func (m *SearchLogQuery) SetEntries(val []ProposedEntry) {
	m.entriesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SearchLogQuery) UnmarshalJSON(raw []byte) error {
	var data struct {
		Entries json.RawMessage `json:"entries"`

		EntryUUIDs []string `json:"entryUUIDs"`

		LogIndexes []int64 `json:"logIndexes"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propEntries []ProposedEntry
	if string(data.Entries) != "null" {
		entries, err := UnmarshalProposedEntrySlice(bytes.NewBuffer(data.Entries), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propEntries = entries
	}

	var result SearchLogQuery

	// entries
	result.entriesField = propEntries

	// entryUUIDs
	result.EntryUUIDs = data.EntryUUIDs

	// logIndexes
	result.LogIndexes = data.LogIndexes

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SearchLogQuery) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		EntryUUIDs []string `json:"entryUUIDs"`

		LogIndexes []int64 `json:"logIndexes"`
	}{

		EntryUUIDs: m.EntryUUIDs,

		LogIndexes: m.LogIndexes,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Entries []ProposedEntry `json:"entries"`
	}{

		Entries: m.entriesField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this search log query
func (m *SearchLogQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogIndexes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchLogQuery) validateEntries(formats strfmt.Registry) error {
	if swag.IsZero(m.Entries()) { // not required
		return nil
	}

	for i := 0; i < len(m.Entries()); i++ {

		if err := m.entriesField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entries" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SearchLogQuery) validateLogIndexes(formats strfmt.Registry) error {
	if swag.IsZero(m.LogIndexes) { // not required
		return nil
	}

	for i := 0; i < len(m.LogIndexes); i++ {

		if err := validate.MinimumInt("logIndexes"+"."+strconv.Itoa(i), "body", int64(m.LogIndexes[i]), 1, false); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this search log query based on the context it is used
func (m *SearchLogQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchLogQuery) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entries()); i++ {

		if err := m.entriesField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entries" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchLogQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchLogQuery) UnmarshalBinary(b []byte) error {
	var res SearchLogQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

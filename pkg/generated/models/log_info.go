// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
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
//

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

// LogInfo log info
//
// swagger:model LogInfo
type LogInfo struct {

	// The current hash value stored at the root of the merkle tree
	// Required: true
	// Pattern: ^[0-9a-fA-F]{64}$
	RootHash *string `json:"rootHash"`

	// The current signed tree head
	// Required: true
	SignedTreeHead *string `json:"signedTreeHead"`

	// The current number of nodes in the merkle tree
	// Required: true
	// Minimum: 1
	TreeSize *int64 `json:"treeSize"`
}

// Validate validates this log info
func (m *LogInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRootHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignedTreeHead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTreeSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogInfo) validateRootHash(formats strfmt.Registry) error {

	if err := validate.Required("rootHash", "body", m.RootHash); err != nil {
		return err
	}

	if err := validate.Pattern("rootHash", "body", *m.RootHash, `^[0-9a-fA-F]{64}$`); err != nil {
		return err
	}

	return nil
}

func (m *LogInfo) validateSignedTreeHead(formats strfmt.Registry) error {

	if err := validate.Required("signedTreeHead", "body", m.SignedTreeHead); err != nil {
		return err
	}

	return nil
}

func (m *LogInfo) validateTreeSize(formats strfmt.Registry) error {

	if err := validate.Required("treeSize", "body", m.TreeSize); err != nil {
		return err
	}

	if err := validate.MinimumInt("treeSize", "body", *m.TreeSize, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this log info based on context it is used
func (m *LogInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogInfo) UnmarshalBinary(b []byte) error {
	var res LogInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

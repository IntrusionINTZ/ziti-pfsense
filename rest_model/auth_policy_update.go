// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthPolicyUpdate auth policy update
//
// swagger:model authPolicyUpdate
type AuthPolicyUpdate struct {
	AuthPolicyCreate
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AuthPolicyUpdate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AuthPolicyCreate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AuthPolicyCreate = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AuthPolicyUpdate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AuthPolicyCreate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this auth policy update
func (m *AuthPolicyUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this auth policy update based on the context it is used
func (m *AuthPolicyUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AuthPolicyCreate
	if err := m.AuthPolicyCreate.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

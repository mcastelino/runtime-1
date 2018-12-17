// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package client_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MachineConfiguration Describes the number of vCPUs, memory size, Hyperthreading capabilities and the CPU template.
// swagger:model MachineConfiguration
type MachineConfiguration struct {

	// cpu template
	CPUTemplate CPUTemplate `json:"cpu_template,omitempty"`

	// Flag for enabling/disabling Hyperthreading
	HtEnabled bool `json:"ht_enabled,omitempty"`

	// Memory size of VM
	MemSizeMib int64 `json:"mem_size_mib,omitempty"`

	// Number of vCPUs (either 1 or an even number)
	VcpuCount int64 `json:"vcpu_count,omitempty"`
}

// Validate validates this machine configuration
func (m *MachineConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineConfiguration) validateCPUTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUTemplate) { // not required
		return nil
	}

	if err := m.CPUTemplate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cpu_template")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineConfiguration) UnmarshalBinary(b []byte) error {
	var res MachineConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

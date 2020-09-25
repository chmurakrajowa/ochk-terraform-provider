// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EdgeClusterGetResponse EdgeClusterGetResponse
//
// swagger:model EdgeClusterGetResponse
type EdgeClusterGetResponse struct {

	// edge cluster instance
	EdgeClusterInstance *EdgeClusterInstance `json:"edgeClusterInstance,omitempty"`

	// messages
	Messages string `json:"messages,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this edge cluster get response
func (m *EdgeClusterGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeClusterInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterGetResponse) validateEdgeClusterInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeClusterInstance) { // not required
		return nil
	}

	if m.EdgeClusterInstance != nil {
		if err := m.EdgeClusterInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeClusterInstance")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeClusterGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterGetResponse) UnmarshalBinary(b []byte) error {
	var res EdgeClusterGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
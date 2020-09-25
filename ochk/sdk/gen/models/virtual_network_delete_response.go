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

// VirtualNetworkDeleteResponse VirtualNetworkDeleteResponse
//
// swagger:model VirtualNetworkDeleteResponse
type VirtualNetworkDeleteResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// request instance
	RequestInstance *RequestInstance `json:"requestInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// virtual network instance
	VirtualNetworkInstance *VirtualNetworkInstance `json:"virtualNetworkInstance,omitempty"`
}

// Validate validates this virtual network delete response
func (m *VirtualNetworkDeleteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualNetworkDeleteResponse) validateRequestInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestInstance) { // not required
		return nil
	}

	if m.RequestInstance != nil {
		if err := m.RequestInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestInstance")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualNetworkDeleteResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualNetworkDeleteResponse) validateVirtualNetworkInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkInstance) { // not required
		return nil
	}

	if m.VirtualNetworkInstance != nil {
		if err := m.VirtualNetworkInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetworkInstance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualNetworkDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualNetworkDeleteResponse) UnmarshalBinary(b []byte) error {
	var res VirtualNetworkDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
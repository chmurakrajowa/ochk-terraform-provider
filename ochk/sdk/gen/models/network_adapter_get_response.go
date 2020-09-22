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

// NetworkAdapterGetResponse NetworkAdapterGetResponse
//
// swagger:model NetworkAdapterGetResponse
type NetworkAdapterGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// network adapter
	NetworkAdapter *NetworkAdapterInstance `json:"networkAdapter,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this network adapter get response
func (m *NetworkAdapterGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkAdapter(formats); err != nil {
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

func (m *NetworkAdapterGetResponse) validateNetworkAdapter(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkAdapter) { // not required
		return nil
	}

	if m.NetworkAdapter != nil {
		if err := m.NetworkAdapter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkAdapter")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkAdapterGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAdapterGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAdapterGetResponse) UnmarshalBinary(b []byte) error {
	var res NetworkAdapterGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

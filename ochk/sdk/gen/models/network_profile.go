// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkProfile NetworkProfile
//
// swagger:model NetworkProfile
type NetworkProfile struct {

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// network profile Id
	NetworkProfileID string `json:"networkProfileId,omitempty"`
}

// Validate validates this network profile
func (m *NetworkProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProfile) UnmarshalBinary(b []byte) error {
	var res NetworkProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

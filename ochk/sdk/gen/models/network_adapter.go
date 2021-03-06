// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkAdapter NetworkAdapter
//
// swagger:model NetworkAdapter
type NetworkAdapter struct {

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// network adapter Id
	NetworkAdapterID string `json:"networkAdapterId,omitempty"`
}

// Validate validates this network adapter
func (m *NetworkAdapter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAdapter) UnmarshalBinary(b []byte) error {
	var res NetworkAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

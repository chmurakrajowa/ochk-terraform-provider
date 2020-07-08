// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecurityPolicyListResponse security policy list response
//
// swagger:model SecurityPolicyListResponse
type SecurityPolicyListResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// security policy collection
	SecurityPolicyCollection []*SecurityPolicy `json:"securityPolicyCollection"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	Timestamp *Timestamp `json:"timestamp,omitempty"`
}

// Validate validates this security policy list response
func (m *SecurityPolicyListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityPolicyCollection(formats); err != nil {
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

func (m *SecurityPolicyListResponse) validateSecurityPolicyCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityPolicyCollection) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityPolicyCollection); i++ {
		if swag.IsZero(m.SecurityPolicyCollection[i]) { // not required
			continue
		}

		if m.SecurityPolicyCollection[i] != nil {
			if err := m.SecurityPolicyCollection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityPolicyCollection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SecurityPolicyListResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if m.Timestamp != nil {
		if err := m.Timestamp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestamp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityPolicyListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityPolicyListResponse) UnmarshalBinary(b []byte) error {
	var res SecurityPolicyListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KeyInstance KeyInstance
//
// swagger:model KeyInstance
type KeyInstance struct {

	// activation date
	// Format: date-time
	ActivationDate *strfmt.DateTime `json:"activationDate,omitempty"`

	// algorithm
	Algorithm string `json:"algorithm,omitempty"`

	// created at
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt,omitempty"`

	// default i v
	DefaultIV string `json:"defaultIV,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// key usage list
	KeyUsageList []string `json:"keyUsageList"`

	// material
	Material string `json:"material,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// object type
	ObjectType string `json:"objectType,omitempty"`

	// revocation reason
	// Enum: [AffiliationChanged CACompromise CessationOfOperation KeyCompromise PrivilegeWithdrawn Superseded Unspecified]
	RevocationReason string `json:"revocationReason,omitempty"`

	// sha1 fingerprint
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`

	// sha256 fingerprint
	Sha256Fingerprint string `json:"sha256Fingerprint,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this key instance
func (m *KeyInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyUsageList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevocationReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyInstance) validateActivationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("activationDate", "body", "date-time", m.ActivationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KeyInstance) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var keyInstanceKeyUsageListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CERTIFICATE_SIGN","CONTENT_COMMITMENT","CRL_SIGN","DECRYPT","DERIVE_KEY","ENCRYPT","EXPORT_KEY","FPE_DECRYPT","FPE_ENCRYPT","GENERATE_CRYPTOGRAM","GENERATE_MAC","KEY_AGREEMENT","SIGN","TRANSLATE_DECRYPT","TRANSLATE_ENCRYPT","TRANSLATE_UNWRAP","TRANSLATE_WRAP","UNWRAP_KEY","VALIDATE_CRYPTOGRAM","VERIFY","VERIFY_MAC","WRAP_KEY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keyInstanceKeyUsageListItemsEnum = append(keyInstanceKeyUsageListItemsEnum, v)
	}
}

func (m *KeyInstance) validateKeyUsageListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, keyInstanceKeyUsageListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KeyInstance) validateKeyUsageList(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyUsageList) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyUsageList); i++ {

		// value enum
		if err := m.validateKeyUsageListItemsEnum("keyUsageList"+"."+strconv.Itoa(i), "body", m.KeyUsageList[i]); err != nil {
			return err
		}

	}

	return nil
}

var keyInstanceTypeRevocationReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AffiliationChanged","CACompromise","CessationOfOperation","KeyCompromise","PrivilegeWithdrawn","Superseded","Unspecified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keyInstanceTypeRevocationReasonPropEnum = append(keyInstanceTypeRevocationReasonPropEnum, v)
	}
}

const (

	// KeyInstanceRevocationReasonAffiliationChanged captures enum value "AffiliationChanged"
	KeyInstanceRevocationReasonAffiliationChanged string = "AffiliationChanged"

	// KeyInstanceRevocationReasonCACompromise captures enum value "CACompromise"
	KeyInstanceRevocationReasonCACompromise string = "CACompromise"

	// KeyInstanceRevocationReasonCessationOfOperation captures enum value "CessationOfOperation"
	KeyInstanceRevocationReasonCessationOfOperation string = "CessationOfOperation"

	// KeyInstanceRevocationReasonKeyCompromise captures enum value "KeyCompromise"
	KeyInstanceRevocationReasonKeyCompromise string = "KeyCompromise"

	// KeyInstanceRevocationReasonPrivilegeWithdrawn captures enum value "PrivilegeWithdrawn"
	KeyInstanceRevocationReasonPrivilegeWithdrawn string = "PrivilegeWithdrawn"

	// KeyInstanceRevocationReasonSuperseded captures enum value "Superseded"
	KeyInstanceRevocationReasonSuperseded string = "Superseded"

	// KeyInstanceRevocationReasonUnspecified captures enum value "Unspecified"
	KeyInstanceRevocationReasonUnspecified string = "Unspecified"
)

// prop value enum
func (m *KeyInstance) validateRevocationReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, keyInstanceTypeRevocationReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KeyInstance) validateRevocationReason(formats strfmt.Registry) error {

	if swag.IsZero(m.RevocationReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateRevocationReasonEnum("revocationReason", "body", m.RevocationReason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyInstance) UnmarshalBinary(b []byte) error {
	var res KeyInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

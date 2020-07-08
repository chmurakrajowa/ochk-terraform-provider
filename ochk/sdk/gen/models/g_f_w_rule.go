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

// GFWRule g f w rule
//
// swagger:model GFWRule
type GFWRule struct {

	// action
	// Enum: [ALLOW REJECT DROP]
	Action string `json:"action,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation date
	CreationDate *Timestamp `json:"creationDate,omitempty"`

	// default services
	DefaultServices []*ServiceInstance `json:"defaultServices"`

	// destination
	Destination []*SecurityGroup `json:"destination"`

	// destination excluded
	DestinationExcluded bool `json:"destinationExcluded,omitempty"`

	// direction
	// Enum: [IN_OUT IN OUT]
	Direction string `json:"direction,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// ip protocol
	// Enum: [IPV4_IPV6 IPV4 IPV6]
	IPProtocol string `json:"ipProtocol,omitempty"`

	// logged
	Logged bool `json:"logged,omitempty"`

	// modification date
	ModificationDate *Timestamp `json:"modificationDate,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// parent path
	ParentPath string `json:"parentPath,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// profile
	Profile []*ContextProfileInstance `json:"profile"`

	// protection
	Protection *Protection `json:"protection,omitempty"`

	// relative path
	RelativePath string `json:"relativePath,omitempty"`

	// resource type
	ResourceType *ResourceType `json:"resourceType,omitempty"`

	// rule Id
	RuleID string `json:"ruleId,omitempty"`

	// scope
	Scope []string `json:"scope"`

	// sequence number
	SequenceNumber int64 `json:"sequenceNumber,omitempty"`

	// source
	Source []*SecurityGroup `json:"source"`

	// sources excluded
	SourcesExcluded bool `json:"sourcesExcluded,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this g f w rule
func (m *GFWRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gFWRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","REJECT","DROP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gFWRuleTypeActionPropEnum = append(gFWRuleTypeActionPropEnum, v)
	}
}

const (

	// GFWRuleActionALLOW captures enum value "ALLOW"
	GFWRuleActionALLOW string = "ALLOW"

	// GFWRuleActionREJECT captures enum value "REJECT"
	GFWRuleActionREJECT string = "REJECT"

	// GFWRuleActionDROP captures enum value "DROP"
	GFWRuleActionDROP string = "DROP"
)

// prop value enum
func (m *GFWRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gFWRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GFWRule) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *GFWRule) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if m.CreationDate != nil {
		if err := m.CreationDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creationDate")
			}
			return err
		}
	}

	return nil
}

func (m *GFWRule) validateDefaultServices(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultServices) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultServices); i++ {
		if swag.IsZero(m.DefaultServices[i]) { // not required
			continue
		}

		if m.DefaultServices[i] != nil {
			if err := m.DefaultServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("defaultServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GFWRule) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	for i := 0; i < len(m.Destination); i++ {
		if swag.IsZero(m.Destination[i]) { // not required
			continue
		}

		if m.Destination[i] != nil {
			if err := m.Destination[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destination" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var gFWRuleTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_OUT","IN","OUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gFWRuleTypeDirectionPropEnum = append(gFWRuleTypeDirectionPropEnum, v)
	}
}

const (

	// GFWRuleDirectionINOUT captures enum value "IN_OUT"
	GFWRuleDirectionINOUT string = "IN_OUT"

	// GFWRuleDirectionIN captures enum value "IN"
	GFWRuleDirectionIN string = "IN"

	// GFWRuleDirectionOUT captures enum value "OUT"
	GFWRuleDirectionOUT string = "OUT"
)

// prop value enum
func (m *GFWRule) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gFWRuleTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GFWRule) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var gFWRuleTypeIPProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPV4_IPV6","IPV4","IPV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gFWRuleTypeIPProtocolPropEnum = append(gFWRuleTypeIPProtocolPropEnum, v)
	}
}

const (

	// GFWRuleIPProtocolIPV4IPV6 captures enum value "IPV4_IPV6"
	GFWRuleIPProtocolIPV4IPV6 string = "IPV4_IPV6"

	// GFWRuleIPProtocolIPV4 captures enum value "IPV4"
	GFWRuleIPProtocolIPV4 string = "IPV4"

	// GFWRuleIPProtocolIPV6 captures enum value "IPV6"
	GFWRuleIPProtocolIPV6 string = "IPV6"
)

// prop value enum
func (m *GFWRule) validateIPProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gFWRuleTypeIPProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GFWRule) validateIPProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.IPProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPProtocolEnum("ipProtocol", "body", m.IPProtocol); err != nil {
		return err
	}

	return nil
}

func (m *GFWRule) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if m.ModificationDate != nil {
		if err := m.ModificationDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modificationDate")
			}
			return err
		}
	}

	return nil
}

func (m *GFWRule) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	for i := 0; i < len(m.Profile); i++ {
		if swag.IsZero(m.Profile[i]) { // not required
			continue
		}

		if m.Profile[i] != nil {
			if err := m.Profile[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profile" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GFWRule) validateProtection(formats strfmt.Registry) error {

	if swag.IsZero(m.Protection) { // not required
		return nil
	}

	if m.Protection != nil {
		if err := m.Protection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection")
			}
			return err
		}
	}

	return nil
}

func (m *GFWRule) validateResourceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceType")
			}
			return err
		}
	}

	return nil
}

func (m *GFWRule) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	for i := 0; i < len(m.Source); i++ {
		if swag.IsZero(m.Source[i]) { // not required
			continue
		}

		if m.Source[i] != nil {
			if err := m.Source[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("source" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GFWRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GFWRule) UnmarshalBinary(b []byte) error {
	var res GFWRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

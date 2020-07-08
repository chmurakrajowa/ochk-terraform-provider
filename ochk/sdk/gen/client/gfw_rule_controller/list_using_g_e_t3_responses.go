// Code generated by go-swagger; DO NOT EDIT.

package gfw_rule_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// ListUsingGET3Reader is a Reader for the ListUsingGET3 structure.
type ListUsingGET3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsingGET3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsingGET3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUsingGET3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUsingGET3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUsingGET3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUsingGET3OK creates a ListUsingGET3OK with default headers values
func NewListUsingGET3OK() *ListUsingGET3OK {
	return &ListUsingGET3OK{}
}

/*ListUsingGET3OK handles this case with default header values.

OK
*/
type ListUsingGET3OK struct {
	Payload *models.GFWRuleListResponse
}

func (o *ListUsingGET3OK) Error() string {
	return fmt.Sprintf("[GET /nsx/fw/gp/{gatewayPolicyId}/rules][%d] listUsingGET3OK  %+v", 200, o.Payload)
}

func (o *ListUsingGET3OK) GetPayload() *models.GFWRuleListResponse {
	return o.Payload
}

func (o *ListUsingGET3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GFWRuleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsingGET3Unauthorized creates a ListUsingGET3Unauthorized with default headers values
func NewListUsingGET3Unauthorized() *ListUsingGET3Unauthorized {
	return &ListUsingGET3Unauthorized{}
}

/*ListUsingGET3Unauthorized handles this case with default header values.

Unauthorized
*/
type ListUsingGET3Unauthorized struct {
}

func (o *ListUsingGET3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /nsx/fw/gp/{gatewayPolicyId}/rules][%d] listUsingGET3Unauthorized ", 401)
}

func (o *ListUsingGET3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGET3Forbidden creates a ListUsingGET3Forbidden with default headers values
func NewListUsingGET3Forbidden() *ListUsingGET3Forbidden {
	return &ListUsingGET3Forbidden{}
}

/*ListUsingGET3Forbidden handles this case with default header values.

Forbidden
*/
type ListUsingGET3Forbidden struct {
}

func (o *ListUsingGET3Forbidden) Error() string {
	return fmt.Sprintf("[GET /nsx/fw/gp/{gatewayPolicyId}/rules][%d] listUsingGET3Forbidden ", 403)
}

func (o *ListUsingGET3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGET3NotFound creates a ListUsingGET3NotFound with default headers values
func NewListUsingGET3NotFound() *ListUsingGET3NotFound {
	return &ListUsingGET3NotFound{}
}

/*ListUsingGET3NotFound handles this case with default header values.

Not Found
*/
type ListUsingGET3NotFound struct {
}

func (o *ListUsingGET3NotFound) Error() string {
	return fmt.Sprintf("[GET /nsx/fw/gp/{gatewayPolicyId}/rules][%d] listUsingGET3NotFound ", 404)
}

func (o *ListUsingGET3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

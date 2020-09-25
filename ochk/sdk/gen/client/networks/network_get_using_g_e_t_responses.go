// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NetworkGetUsingGETReader is a Reader for the NetworkGetUsingGET structure.
type NetworkGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNetworkGetUsingGETOK creates a NetworkGetUsingGETOK with default headers values
func NewNetworkGetUsingGETOK() *NetworkGetUsingGETOK {
	return &NetworkGetUsingGETOK{}
}

/*NetworkGetUsingGETOK handles this case with default header values.

OK
*/
type NetworkGetUsingGETOK struct {
	Payload *models.NetworkGetResponse
}

func (o *NetworkGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/networks/{networkId}][%d] networkGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *NetworkGetUsingGETOK) GetPayload() *models.NetworkGetResponse {
	return o.Payload
}

func (o *NetworkGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkGetUsingGETBadRequest creates a NetworkGetUsingGETBadRequest with default headers values
func NewNetworkGetUsingGETBadRequest() *NetworkGetUsingGETBadRequest {
	return &NetworkGetUsingGETBadRequest{}
}

/*NetworkGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type NetworkGetUsingGETBadRequest struct {
}

func (o *NetworkGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/networks/{networkId}][%d] networkGetUsingGETBadRequest ", 400)
}

func (o *NetworkGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkGetUsingGETNotFound creates a NetworkGetUsingGETNotFound with default headers values
func NewNetworkGetUsingGETNotFound() *NetworkGetUsingGETNotFound {
	return &NetworkGetUsingGETNotFound{}
}

/*NetworkGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type NetworkGetUsingGETNotFound struct {
}

func (o *NetworkGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /subtenants/{subtenantId}/networks/{networkId}][%d] networkGetUsingGETNotFound ", 404)
}

func (o *NetworkGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
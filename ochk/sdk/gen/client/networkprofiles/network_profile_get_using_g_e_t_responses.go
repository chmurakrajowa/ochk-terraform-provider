// Code generated by go-swagger; DO NOT EDIT.

package networkprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NetworkProfileGetUsingGETReader is a Reader for the NetworkProfileGetUsingGET structure.
type NetworkProfileGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkProfileGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkProfileGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkProfileGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNetworkProfileGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNetworkProfileGetUsingGETOK creates a NetworkProfileGetUsingGETOK with default headers values
func NewNetworkProfileGetUsingGETOK() *NetworkProfileGetUsingGETOK {
	return &NetworkProfileGetUsingGETOK{}
}

/*NetworkProfileGetUsingGETOK handles this case with default header values.

OK
*/
type NetworkProfileGetUsingGETOK struct {
	Payload *models.NetworkProfileGetResponse
}

func (o *NetworkProfileGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /vra/networkprofiles/{networkProfileId}][%d] networkProfileGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *NetworkProfileGetUsingGETOK) GetPayload() *models.NetworkProfileGetResponse {
	return o.Payload
}

func (o *NetworkProfileGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkProfileGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkProfileGetUsingGETBadRequest creates a NetworkProfileGetUsingGETBadRequest with default headers values
func NewNetworkProfileGetUsingGETBadRequest() *NetworkProfileGetUsingGETBadRequest {
	return &NetworkProfileGetUsingGETBadRequest{}
}

/*NetworkProfileGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type NetworkProfileGetUsingGETBadRequest struct {
}

func (o *NetworkProfileGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /vra/networkprofiles/{networkProfileId}][%d] networkProfileGetUsingGETBadRequest ", 400)
}

func (o *NetworkProfileGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkProfileGetUsingGETNotFound creates a NetworkProfileGetUsingGETNotFound with default headers values
func NewNetworkProfileGetUsingGETNotFound() *NetworkProfileGetUsingGETNotFound {
	return &NetworkProfileGetUsingGETNotFound{}
}

/*NetworkProfileGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type NetworkProfileGetUsingGETNotFound struct {
}

func (o *NetworkProfileGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /vra/networkprofiles/{networkProfileId}][%d] networkProfileGetUsingGETNotFound ", 404)
}

func (o *NetworkProfileGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package ip_collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// IPCollectionUpdateUsingPUTReader is a Reader for the IPCollectionUpdateUsingPUT structure.
type IPCollectionUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPCollectionUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPCollectionUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIPCollectionUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPCollectionUpdateUsingPUTOK creates a IPCollectionUpdateUsingPUTOK with default headers values
func NewIPCollectionUpdateUsingPUTOK() *IPCollectionUpdateUsingPUTOK {
	return &IPCollectionUpdateUsingPUTOK{}
}

/*IPCollectionUpdateUsingPUTOK handles this case with default header values.

OK
*/
type IPCollectionUpdateUsingPUTOK struct {
	Payload *models.IPCollectionUpdateResponse
}

func (o *IPCollectionUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] ipCollectionUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *IPCollectionUpdateUsingPUTOK) GetPayload() *models.IPCollectionUpdateResponse {
	return o.Payload
}

func (o *IPCollectionUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPCollectionUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPCollectionUpdateUsingPUTBadRequest creates a IPCollectionUpdateUsingPUTBadRequest with default headers values
func NewIPCollectionUpdateUsingPUTBadRequest() *IPCollectionUpdateUsingPUTBadRequest {
	return &IPCollectionUpdateUsingPUTBadRequest{}
}

/*IPCollectionUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type IPCollectionUpdateUsingPUTBadRequest struct {
}

func (o *IPCollectionUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipcs/{ipCollectionId}][%d] ipCollectionUpdateUsingPUTBadRequest ", 400)
}

func (o *IPCollectionUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
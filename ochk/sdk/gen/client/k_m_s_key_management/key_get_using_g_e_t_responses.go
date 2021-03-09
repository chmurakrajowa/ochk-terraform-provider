// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// KeyGetUsingGETReader is a Reader for the KeyGetUsingGET structure.
type KeyGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeyGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyGetUsingGETOK creates a KeyGetUsingGETOK with default headers values
func NewKeyGetUsingGETOK() *KeyGetUsingGETOK {
	return &KeyGetUsingGETOK{}
}

/*KeyGetUsingGETOK handles this case with default header values.

OK
*/
type KeyGetUsingGETOK struct {
	Payload *models.KeyGetResponse
}

func (o *KeyGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /kms/key/{id}][%d] keyGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *KeyGetUsingGETOK) GetPayload() *models.KeyGetResponse {
	return o.Payload
}

func (o *KeyGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyGetUsingGETBadRequest creates a KeyGetUsingGETBadRequest with default headers values
func NewKeyGetUsingGETBadRequest() *KeyGetUsingGETBadRequest {
	return &KeyGetUsingGETBadRequest{}
}

/*KeyGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyGetUsingGETBadRequest struct {
}

func (o *KeyGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /kms/key/{id}][%d] keyGetUsingGETBadRequest ", 400)
}

func (o *KeyGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKeyGetUsingGETNotFound creates a KeyGetUsingGETNotFound with default headers values
func NewKeyGetUsingGETNotFound() *KeyGetUsingGETNotFound {
	return &KeyGetUsingGETNotFound{}
}

/*KeyGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type KeyGetUsingGETNotFound struct {
}

func (o *KeyGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /kms/key/{id}][%d] keyGetUsingGETNotFound ", 404)
}

func (o *KeyGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
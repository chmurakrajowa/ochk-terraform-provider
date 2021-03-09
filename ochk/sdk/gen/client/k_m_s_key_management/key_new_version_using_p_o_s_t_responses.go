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

// KeyNewVersionUsingPOSTReader is a Reader for the KeyNewVersionUsingPOST structure.
type KeyNewVersionUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyNewVersionUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeyNewVersionUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewKeyNewVersionUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeyNewVersionUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyNewVersionUsingPOSTOK creates a KeyNewVersionUsingPOSTOK with default headers values
func NewKeyNewVersionUsingPOSTOK() *KeyNewVersionUsingPOSTOK {
	return &KeyNewVersionUsingPOSTOK{}
}

/*KeyNewVersionUsingPOSTOK handles this case with default header values.

OK
*/
type KeyNewVersionUsingPOSTOK struct {
	Payload *models.CreateNewKmsKeyVersionResponse
}

func (o *KeyNewVersionUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *KeyNewVersionUsingPOSTOK) GetPayload() *models.CreateNewKmsKeyVersionResponse {
	return o.Payload
}

func (o *KeyNewVersionUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNewKmsKeyVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyNewVersionUsingPOSTCreated creates a KeyNewVersionUsingPOSTCreated with default headers values
func NewKeyNewVersionUsingPOSTCreated() *KeyNewVersionUsingPOSTCreated {
	return &KeyNewVersionUsingPOSTCreated{}
}

/*KeyNewVersionUsingPOSTCreated handles this case with default header values.

New key version has been created.
*/
type KeyNewVersionUsingPOSTCreated struct {
	Payload *models.CreateNewKmsKeyVersionResponse
}

func (o *KeyNewVersionUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *KeyNewVersionUsingPOSTCreated) GetPayload() *models.CreateNewKmsKeyVersionResponse {
	return o.Payload
}

func (o *KeyNewVersionUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNewKmsKeyVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyNewVersionUsingPOSTBadRequest creates a KeyNewVersionUsingPOSTBadRequest with default headers values
func NewKeyNewVersionUsingPOSTBadRequest() *KeyNewVersionUsingPOSTBadRequest {
	return &KeyNewVersionUsingPOSTBadRequest{}
}

/*KeyNewVersionUsingPOSTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type KeyNewVersionUsingPOSTBadRequest struct {
}

func (o *KeyNewVersionUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /kms/key/{id}/versions][%d] keyNewVersionUsingPOSTBadRequest ", 400)
}

func (o *KeyNewVersionUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
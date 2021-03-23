// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// VcsVirtualMachineUpdateUsingPUTReader is a Reader for the VcsVirtualMachineUpdateUsingPUT structure.
type VcsVirtualMachineUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcsVirtualMachineUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcsVirtualMachineUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewVcsVirtualMachineUpdateUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVcsVirtualMachineUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVcsVirtualMachineUpdateUsingPUTOK creates a VcsVirtualMachineUpdateUsingPUTOK with default headers values
func NewVcsVirtualMachineUpdateUsingPUTOK() *VcsVirtualMachineUpdateUsingPUTOK {
	return &VcsVirtualMachineUpdateUsingPUTOK{}
}

/*VcsVirtualMachineUpdateUsingPUTOK handles this case with default header values.

OK
*/
type VcsVirtualMachineUpdateUsingPUTOK struct {
	Payload *models.VcsVirtualMachineUpdateResponse
}

func (o *VcsVirtualMachineUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *VcsVirtualMachineUpdateUsingPUTOK) GetPayload() *models.VcsVirtualMachineUpdateResponse {
	return o.Payload
}

func (o *VcsVirtualMachineUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsVirtualMachineUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineUpdateUsingPUTCreated creates a VcsVirtualMachineUpdateUsingPUTCreated with default headers values
func NewVcsVirtualMachineUpdateUsingPUTCreated() *VcsVirtualMachineUpdateUsingPUTCreated {
	return &VcsVirtualMachineUpdateUsingPUTCreated{}
}

/*VcsVirtualMachineUpdateUsingPUTCreated handles this case with default header values.

Entity has been updated
*/
type VcsVirtualMachineUpdateUsingPUTCreated struct {
	Payload *models.VcsVirtualMachineUpdateResponse
}

func (o *VcsVirtualMachineUpdateUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineUpdateUsingPUTCreated  %+v", 201, o.Payload)
}

func (o *VcsVirtualMachineUpdateUsingPUTCreated) GetPayload() *models.VcsVirtualMachineUpdateResponse {
	return o.Payload
}

func (o *VcsVirtualMachineUpdateUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsVirtualMachineUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcsVirtualMachineUpdateUsingPUTBadRequest creates a VcsVirtualMachineUpdateUsingPUTBadRequest with default headers values
func NewVcsVirtualMachineUpdateUsingPUTBadRequest() *VcsVirtualMachineUpdateUsingPUTBadRequest {
	return &VcsVirtualMachineUpdateUsingPUTBadRequest{}
}

/*VcsVirtualMachineUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type VcsVirtualMachineUpdateUsingPUTBadRequest struct {
}

func (o *VcsVirtualMachineUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /vcs/virtual-machines/{virtualMachineId}][%d] vcsVirtualMachineUpdateUsingPUTBadRequest ", 400)
}

func (o *VcsVirtualMachineUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

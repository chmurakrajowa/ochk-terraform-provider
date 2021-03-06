// Code generated by go-swagger; DO NOT EDIT.

package custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CustomServiceCreateUsingPUT(params *CustomServiceCreateUsingPUTParams) (*CustomServiceCreateUsingPUTOK, *CustomServiceCreateUsingPUTCreated, error)

	CustomServiceDeleteUsingDELETE(params *CustomServiceDeleteUsingDELETEParams) (*CustomServiceDeleteUsingDELETEOK, *CustomServiceDeleteUsingDELETECreated, error)

	CustomServiceGetUsingGET(params *CustomServiceGetUsingGETParams) (*CustomServiceGetUsingGETOK, error)

	CustomServiceListUsingGET(params *CustomServiceListUsingGETParams) (*CustomServiceListUsingGETOK, error)

	CustomServiceUpdateUsingPUT(params *CustomServiceUpdateUsingPUTParams) (*CustomServiceUpdateUsingPUTOK, *CustomServiceUpdateUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CustomServiceCreateUsingPUT creates

  Create custom service in NSX-T
*/
func (a *Client) CustomServiceCreateUsingPUT(params *CustomServiceCreateUsingPUTParams) (*CustomServiceCreateUsingPUTOK, *CustomServiceCreateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomServiceCreateUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customServiceCreateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/network/custom-services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomServiceCreateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CustomServiceCreateUsingPUTOK:
		return value, nil, nil
	case *CustomServiceCreateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for custom_services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomServiceDeleteUsingDELETE deletes

  Delete custom service from NSX-T
*/
func (a *Client) CustomServiceDeleteUsingDELETE(params *CustomServiceDeleteUsingDELETEParams) (*CustomServiceDeleteUsingDELETEOK, *CustomServiceDeleteUsingDELETECreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomServiceDeleteUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customServiceDeleteUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/network/custom-services/{serviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomServiceDeleteUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CustomServiceDeleteUsingDELETEOK:
		return value, nil, nil
	case *CustomServiceDeleteUsingDELETECreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for custom_services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomServiceGetUsingGET gets

  Get custom service from NSX-T
*/
func (a *Client) CustomServiceGetUsingGET(params *CustomServiceGetUsingGETParams) (*CustomServiceGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomServiceGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customServiceGetUsingGET",
		Method:             "GET",
		PathPattern:        "/network/custom-services/{serviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomServiceGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomServiceGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for customServiceGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomServiceListUsingGET lists

  List custom services from NSX-T
*/
func (a *Client) CustomServiceListUsingGET(params *CustomServiceListUsingGETParams) (*CustomServiceListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomServiceListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customServiceListUsingGET",
		Method:             "GET",
		PathPattern:        "/network/custom-services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomServiceListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomServiceListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for customServiceListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CustomServiceUpdateUsingPUT updates

  Update custom service from NSX-T
*/
func (a *Client) CustomServiceUpdateUsingPUT(params *CustomServiceUpdateUsingPUTParams) (*CustomServiceUpdateUsingPUTOK, *CustomServiceUpdateUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomServiceUpdateUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customServiceUpdateUsingPUT",
		Method:             "PUT",
		PathPattern:        "/network/custom-services/{serviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomServiceUpdateUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CustomServiceUpdateUsingPUTOK:
		return value, nil, nil
	case *CustomServiceUpdateUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for custom_services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

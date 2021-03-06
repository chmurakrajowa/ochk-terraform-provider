// Code generated by go-swagger; DO NOT EDIT.

package reservations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reservations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reservations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ReservationGetUsingGET(params *ReservationGetUsingGETParams) (*ReservationGetUsingGETOK, error)

	ReservationListUsingGET(params *ReservationListUsingGETParams) (*ReservationListUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ReservationGetUsingGET gets

  Get reservation
*/
func (a *Client) ReservationGetUsingGET(params *ReservationGetUsingGETParams) (*ReservationGetUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReservationGetUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reservationGetUsingGET",
		Method:             "GET",
		PathPattern:        "/subtenants/{subtenantId}/reservations/{reservationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReservationGetUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReservationGetUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reservationGetUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReservationListUsingGET lists

  List reservations
*/
func (a *Client) ReservationListUsingGET(params *ReservationListUsingGETParams) (*ReservationListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReservationListUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reservationListUsingGET",
		Method:             "GET",
		PathPattern:        "/subtenants/{subtenantId}/reservations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReservationListUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReservationListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for reservationListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

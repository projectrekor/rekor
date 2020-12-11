// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tlog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tlog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetLogInfo(params *GetLogInfoParams) (*GetLogInfoOK, error)

	GetLogProof(params *GetLogProofParams) (*GetLogProofOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLogInfo gets information about the current state of the transparency log

  Returns the current root hash and size of the merkle tree used to store the log entries.
*/
func (a *Client) GetLogInfo(params *GetLogInfoParams) (*GetLogInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLogInfo",
		Method:             "GET",
		PathPattern:        "/api/v1/log",
		ProducesMediaTypes: []string{"application/json", "application/xml", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLogInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLogInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLogProof gets information required to generate a consistency proof for the transparency log

  Returns a list of hashes for specified tree sizes that can be used to confirm the consistency of the transparency log
*/
func (a *Client) GetLogProof(params *GetLogProofParams) (*GetLogProofOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogProofParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLogProof",
		Method:             "GET",
		PathPattern:        "/api/v1/log/proof",
		ProducesMediaTypes: []string{"application/json", "application/xml", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLogProofReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLogProofOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogProofDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

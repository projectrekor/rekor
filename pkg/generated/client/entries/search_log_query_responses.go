// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
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
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// SearchLogQueryReader is a Reader for the SearchLogQuery structure.
type SearchLogQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLogQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLogQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLogQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSearchLogQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchLogQueryOK creates a SearchLogQueryOK with default headers values
func NewSearchLogQueryOK() *SearchLogQueryOK {
	return &SearchLogQueryOK{}
}

/* SearchLogQueryOK describes a response with status code 200, with default header values.

Returns zero or more entries from the transparency log, according to how many were included in request query
*/
type SearchLogQueryOK struct {
	Payload []models.LogEntry
}

func (o *SearchLogQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/log/entries/retrieve][%d] searchLogQueryOK  %+v", 200, o.Payload)
}
func (o *SearchLogQueryOK) GetPayload() []models.LogEntry {
	return o.Payload
}

func (o *SearchLogQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLogQueryBadRequest creates a SearchLogQueryBadRequest with default headers values
func NewSearchLogQueryBadRequest() *SearchLogQueryBadRequest {
	return &SearchLogQueryBadRequest{}
}

/* SearchLogQueryBadRequest describes a response with status code 400, with default header values.

The content supplied to the server was invalid
*/
type SearchLogQueryBadRequest struct {
	Payload *models.Error
}

func (o *SearchLogQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/log/entries/retrieve][%d] searchLogQueryBadRequest  %+v", 400, o.Payload)
}
func (o *SearchLogQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchLogQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLogQueryDefault creates a SearchLogQueryDefault with default headers values
func NewSearchLogQueryDefault(code int) *SearchLogQueryDefault {
	return &SearchLogQueryDefault{
		_statusCode: code,
	}
}

/* SearchLogQueryDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type SearchLogQueryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the search log query default response
func (o *SearchLogQueryDefault) Code() int {
	return o._statusCode
}

func (o *SearchLogQueryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/log/entries/retrieve][%d] searchLogQuery default  %+v", o._statusCode, o.Payload)
}
func (o *SearchLogQueryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchLogQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLogProofHandlerFunc turns a function with the right signature into a get log proof handler
type GetLogProofHandlerFunc func(GetLogProofParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogProofHandlerFunc) Handle(params GetLogProofParams) middleware.Responder {
	return fn(params)
}

// GetLogProofHandler interface for that can handle valid get log proof params
type GetLogProofHandler interface {
	Handle(GetLogProofParams) middleware.Responder
}

// NewGetLogProof creates a new http.Handler for the get log proof operation
func NewGetLogProof(ctx *middleware.Context, handler GetLogProofHandler) *GetLogProof {
	return &GetLogProof{Context: ctx, Handler: handler}
}

/*GetLogProof swagger:route GET /api/v1/log/proof tlog getLogProof

Get information required to generate a consistency proof for the transparency log

Returns a list of hashes for specified tree sizes that can be used to confirm the consistency of the transparency log

*/
type GetLogProof struct {
	Context *middleware.Context
	Handler GetLogProofHandler
}

func (o *GetLogProof) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLogProofParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

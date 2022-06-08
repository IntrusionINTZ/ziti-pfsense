// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DisassociateIdentitysServiceConfigsHandlerFunc turns a function with the right signature into a disassociate identitys service configs handler
type DisassociateIdentitysServiceConfigsHandlerFunc func(DisassociateIdentitysServiceConfigsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DisassociateIdentitysServiceConfigsHandlerFunc) Handle(params DisassociateIdentitysServiceConfigsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DisassociateIdentitysServiceConfigsHandler interface for that can handle valid disassociate identitys service configs params
type DisassociateIdentitysServiceConfigsHandler interface {
	Handle(DisassociateIdentitysServiceConfigsParams, interface{}) middleware.Responder
}

// NewDisassociateIdentitysServiceConfigs creates a new http.Handler for the disassociate identitys service configs operation
func NewDisassociateIdentitysServiceConfigs(ctx *middleware.Context, handler DisassociateIdentitysServiceConfigsHandler) *DisassociateIdentitysServiceConfigs {
	return &DisassociateIdentitysServiceConfigs{Context: ctx, Handler: handler}
}

/* DisassociateIdentitysServiceConfigs swagger:route DELETE /identities/{id}/service-configs Identity disassociateIdentitysServiceConfigs

Remove associated service configs from a specific identity

Remove service configs from a specific identity

*/
type DisassociateIdentitysServiceConfigs struct {
	Context *middleware.Context
	Handler DisassociateIdentitysServiceConfigsHandler
}

func (o *DisassociateIdentitysServiceConfigs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDisassociateIdentitysServiceConfigsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

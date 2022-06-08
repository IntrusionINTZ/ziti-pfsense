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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListEdgeRouterServiceEdgeRouterPoliciesOKCode is the HTTP code returned for type ListEdgeRouterServiceEdgeRouterPoliciesOK
const ListEdgeRouterServiceEdgeRouterPoliciesOKCode int = 200

/*ListEdgeRouterServiceEdgeRouterPoliciesOK A list of service policies

swagger:response listEdgeRouterServiceEdgeRouterPoliciesOK
*/
type ListEdgeRouterServiceEdgeRouterPoliciesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListServicePoliciesEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServiceEdgeRouterPoliciesOK creates ListEdgeRouterServiceEdgeRouterPoliciesOK with default headers values
func NewListEdgeRouterServiceEdgeRouterPoliciesOK() *ListEdgeRouterServiceEdgeRouterPoliciesOK {

	return &ListEdgeRouterServiceEdgeRouterPoliciesOK{}
}

// WithPayload adds the payload to the list edge router service edge router policies o k response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesOK) WithPayload(payload *rest_model.ListServicePoliciesEnvelope) *ListEdgeRouterServiceEdgeRouterPoliciesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router service edge router policies o k response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesOK) SetPayload(payload *rest_model.ListServicePoliciesEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServiceEdgeRouterPoliciesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServiceEdgeRouterPoliciesUnauthorizedCode is the HTTP code returned for type ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized
const ListEdgeRouterServiceEdgeRouterPoliciesUnauthorizedCode int = 401

/*ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listEdgeRouterServiceEdgeRouterPoliciesUnauthorized
*/
type ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServiceEdgeRouterPoliciesUnauthorized creates ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized with default headers values
func NewListEdgeRouterServiceEdgeRouterPoliciesUnauthorized() *ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized {

	return &ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized{}
}

// WithPayload adds the payload to the list edge router service edge router policies unauthorized response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router service edge router policies unauthorized response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServiceEdgeRouterPoliciesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEdgeRouterServiceEdgeRouterPoliciesNotFoundCode is the HTTP code returned for type ListEdgeRouterServiceEdgeRouterPoliciesNotFound
const ListEdgeRouterServiceEdgeRouterPoliciesNotFoundCode int = 404

/*ListEdgeRouterServiceEdgeRouterPoliciesNotFound The requested resource does not exist

swagger:response listEdgeRouterServiceEdgeRouterPoliciesNotFound
*/
type ListEdgeRouterServiceEdgeRouterPoliciesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListEdgeRouterServiceEdgeRouterPoliciesNotFound creates ListEdgeRouterServiceEdgeRouterPoliciesNotFound with default headers values
func NewListEdgeRouterServiceEdgeRouterPoliciesNotFound() *ListEdgeRouterServiceEdgeRouterPoliciesNotFound {

	return &ListEdgeRouterServiceEdgeRouterPoliciesNotFound{}
}

// WithPayload adds the payload to the list edge router service edge router policies not found response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListEdgeRouterServiceEdgeRouterPoliciesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list edge router service edge router policies not found response
func (o *ListEdgeRouterServiceEdgeRouterPoliciesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEdgeRouterServiceEdgeRouterPoliciesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

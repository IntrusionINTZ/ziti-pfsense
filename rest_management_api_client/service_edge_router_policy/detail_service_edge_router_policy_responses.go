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

package service_edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// DetailServiceEdgeRouterPolicyReader is a Reader for the DetailServiceEdgeRouterPolicy structure.
type DetailServiceEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailServiceEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailServiceEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailServiceEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailServiceEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailServiceEdgeRouterPolicyOK creates a DetailServiceEdgeRouterPolicyOK with default headers values
func NewDetailServiceEdgeRouterPolicyOK() *DetailServiceEdgeRouterPolicyOK {
	return &DetailServiceEdgeRouterPolicyOK{}
}

/* DetailServiceEdgeRouterPolicyOK describes a response with status code 200, with default header values.

A single service edge router policy
*/
type DetailServiceEdgeRouterPolicyOK struct {
	Payload *rest_model.DetailServiceEdgePolicyEnvelope
}

func (o *DetailServiceEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}][%d] detailServiceEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *DetailServiceEdgeRouterPolicyOK) GetPayload() *rest_model.DetailServiceEdgePolicyEnvelope {
	return o.Payload
}

func (o *DetailServiceEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailServiceEdgePolicyEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailServiceEdgeRouterPolicyUnauthorized creates a DetailServiceEdgeRouterPolicyUnauthorized with default headers values
func NewDetailServiceEdgeRouterPolicyUnauthorized() *DetailServiceEdgeRouterPolicyUnauthorized {
	return &DetailServiceEdgeRouterPolicyUnauthorized{}
}

/* DetailServiceEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailServiceEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailServiceEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}][%d] detailServiceEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailServiceEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailServiceEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailServiceEdgeRouterPolicyNotFound creates a DetailServiceEdgeRouterPolicyNotFound with default headers values
func NewDetailServiceEdgeRouterPolicyNotFound() *DetailServiceEdgeRouterPolicyNotFound {
	return &DetailServiceEdgeRouterPolicyNotFound{}
}

/* DetailServiceEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailServiceEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailServiceEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /service-edge-router-policies/{id}][%d] detailServiceEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *DetailServiceEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailServiceEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

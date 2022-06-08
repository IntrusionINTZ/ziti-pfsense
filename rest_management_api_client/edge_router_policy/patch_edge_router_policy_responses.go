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

package edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// PatchEdgeRouterPolicyReader is a Reader for the PatchEdgeRouterPolicy structure.
type PatchEdgeRouterPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEdgeRouterPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEdgeRouterPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchEdgeRouterPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchEdgeRouterPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchEdgeRouterPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchEdgeRouterPolicyOK creates a PatchEdgeRouterPolicyOK with default headers values
func NewPatchEdgeRouterPolicyOK() *PatchEdgeRouterPolicyOK {
	return &PatchEdgeRouterPolicyOK{}
}

/* PatchEdgeRouterPolicyOK describes a response with status code 200, with default header values.

The patch request was successful and the resource has been altered
*/
type PatchEdgeRouterPolicyOK struct {
	Payload *rest_model.Empty
}

func (o *PatchEdgeRouterPolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /edge-router-policies/{id}][%d] patchEdgeRouterPolicyOK  %+v", 200, o.Payload)
}
func (o *PatchEdgeRouterPolicyOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *PatchEdgeRouterPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterPolicyBadRequest creates a PatchEdgeRouterPolicyBadRequest with default headers values
func NewPatchEdgeRouterPolicyBadRequest() *PatchEdgeRouterPolicyBadRequest {
	return &PatchEdgeRouterPolicyBadRequest{}
}

/* PatchEdgeRouterPolicyBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type PatchEdgeRouterPolicyBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchEdgeRouterPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /edge-router-policies/{id}][%d] patchEdgeRouterPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *PatchEdgeRouterPolicyBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterPolicyUnauthorized creates a PatchEdgeRouterPolicyUnauthorized with default headers values
func NewPatchEdgeRouterPolicyUnauthorized() *PatchEdgeRouterPolicyUnauthorized {
	return &PatchEdgeRouterPolicyUnauthorized{}
}

/* PatchEdgeRouterPolicyUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type PatchEdgeRouterPolicyUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchEdgeRouterPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /edge-router-policies/{id}][%d] patchEdgeRouterPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchEdgeRouterPolicyUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEdgeRouterPolicyNotFound creates a PatchEdgeRouterPolicyNotFound with default headers values
func NewPatchEdgeRouterPolicyNotFound() *PatchEdgeRouterPolicyNotFound {
	return &PatchEdgeRouterPolicyNotFound{}
}

/* PatchEdgeRouterPolicyNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type PatchEdgeRouterPolicyNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *PatchEdgeRouterPolicyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /edge-router-policies/{id}][%d] patchEdgeRouterPolicyNotFound  %+v", 404, o.Payload)
}
func (o *PatchEdgeRouterPolicyNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *PatchEdgeRouterPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

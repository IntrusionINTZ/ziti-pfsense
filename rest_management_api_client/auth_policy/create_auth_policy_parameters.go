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

package auth_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// NewCreateAuthPolicyParams creates a new CreateAuthPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAuthPolicyParams() *CreateAuthPolicyParams {
	return &CreateAuthPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthPolicyParamsWithTimeout creates a new CreateAuthPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateAuthPolicyParamsWithTimeout(timeout time.Duration) *CreateAuthPolicyParams {
	return &CreateAuthPolicyParams{
		timeout: timeout,
	}
}

// NewCreateAuthPolicyParamsWithContext creates a new CreateAuthPolicyParams object
// with the ability to set a context for a request.
func NewCreateAuthPolicyParamsWithContext(ctx context.Context) *CreateAuthPolicyParams {
	return &CreateAuthPolicyParams{
		Context: ctx,
	}
}

// NewCreateAuthPolicyParamsWithHTTPClient creates a new CreateAuthPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAuthPolicyParamsWithHTTPClient(client *http.Client) *CreateAuthPolicyParams {
	return &CreateAuthPolicyParams{
		HTTPClient: client,
	}
}

/* CreateAuthPolicyParams contains all the parameters to send to the API endpoint
   for the create auth policy operation.

   Typically these are written to a http.Request.
*/
type CreateAuthPolicyParams struct {

	/* AuthPolicy.

	   An Auth Policy to create
	*/
	AuthPolicy *rest_model.AuthPolicyCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create auth policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthPolicyParams) WithDefaults() *CreateAuthPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create auth policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAuthPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create auth policy params
func (o *CreateAuthPolicyParams) WithTimeout(timeout time.Duration) *CreateAuthPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create auth policy params
func (o *CreateAuthPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create auth policy params
func (o *CreateAuthPolicyParams) WithContext(ctx context.Context) *CreateAuthPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create auth policy params
func (o *CreateAuthPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create auth policy params
func (o *CreateAuthPolicyParams) WithHTTPClient(client *http.Client) *CreateAuthPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create auth policy params
func (o *CreateAuthPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthPolicy adds the authPolicy to the create auth policy params
func (o *CreateAuthPolicyParams) WithAuthPolicy(authPolicy *rest_model.AuthPolicyCreate) *CreateAuthPolicyParams {
	o.SetAuthPolicy(authPolicy)
	return o
}

// SetAuthPolicy adds the authPolicy to the create auth policy params
func (o *CreateAuthPolicyParams) SetAuthPolicy(authPolicy *rest_model.AuthPolicyCreate) {
	o.AuthPolicy = authPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AuthPolicy != nil {
		if err := r.SetBodyParam(o.AuthPolicy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

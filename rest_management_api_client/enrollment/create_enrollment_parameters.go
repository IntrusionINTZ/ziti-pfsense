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

package enrollment

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

// NewCreateEnrollmentParams creates a new CreateEnrollmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEnrollmentParams() *CreateEnrollmentParams {
	return &CreateEnrollmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEnrollmentParamsWithTimeout creates a new CreateEnrollmentParams object
// with the ability to set a timeout on a request.
func NewCreateEnrollmentParamsWithTimeout(timeout time.Duration) *CreateEnrollmentParams {
	return &CreateEnrollmentParams{
		timeout: timeout,
	}
}

// NewCreateEnrollmentParamsWithContext creates a new CreateEnrollmentParams object
// with the ability to set a context for a request.
func NewCreateEnrollmentParamsWithContext(ctx context.Context) *CreateEnrollmentParams {
	return &CreateEnrollmentParams{
		Context: ctx,
	}
}

// NewCreateEnrollmentParamsWithHTTPClient creates a new CreateEnrollmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEnrollmentParamsWithHTTPClient(client *http.Client) *CreateEnrollmentParams {
	return &CreateEnrollmentParams{
		HTTPClient: client,
	}
}

/* CreateEnrollmentParams contains all the parameters to send to the API endpoint
   for the create enrollment operation.

   Typically these are written to a http.Request.
*/
type CreateEnrollmentParams struct {

	/* Enrollment.

	   An enrollment to create
	*/
	Enrollment *rest_model.EnrollmentCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create enrollment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEnrollmentParams) WithDefaults() *CreateEnrollmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create enrollment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEnrollmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create enrollment params
func (o *CreateEnrollmentParams) WithTimeout(timeout time.Duration) *CreateEnrollmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create enrollment params
func (o *CreateEnrollmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create enrollment params
func (o *CreateEnrollmentParams) WithContext(ctx context.Context) *CreateEnrollmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create enrollment params
func (o *CreateEnrollmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create enrollment params
func (o *CreateEnrollmentParams) WithHTTPClient(client *http.Client) *CreateEnrollmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create enrollment params
func (o *CreateEnrollmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnrollment adds the enrollment to the create enrollment params
func (o *CreateEnrollmentParams) WithEnrollment(enrollment *rest_model.EnrollmentCreate) *CreateEnrollmentParams {
	o.SetEnrollment(enrollment)
	return o
}

// SetEnrollment adds the enrollment to the create enrollment params
func (o *CreateEnrollmentParams) SetEnrollment(enrollment *rest_model.EnrollmentCreate) {
	o.Enrollment = enrollment
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEnrollmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Enrollment != nil {
		if err := r.SetBodyParam(o.Enrollment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

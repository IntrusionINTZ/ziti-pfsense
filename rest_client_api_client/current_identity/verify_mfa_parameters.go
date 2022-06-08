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

package current_identity

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

// NewVerifyMfaParams creates a new VerifyMfaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyMfaParams() *VerifyMfaParams {
	return &VerifyMfaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyMfaParamsWithTimeout creates a new VerifyMfaParams object
// with the ability to set a timeout on a request.
func NewVerifyMfaParamsWithTimeout(timeout time.Duration) *VerifyMfaParams {
	return &VerifyMfaParams{
		timeout: timeout,
	}
}

// NewVerifyMfaParamsWithContext creates a new VerifyMfaParams object
// with the ability to set a context for a request.
func NewVerifyMfaParamsWithContext(ctx context.Context) *VerifyMfaParams {
	return &VerifyMfaParams{
		Context: ctx,
	}
}

// NewVerifyMfaParamsWithHTTPClient creates a new VerifyMfaParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyMfaParamsWithHTTPClient(client *http.Client) *VerifyMfaParams {
	return &VerifyMfaParams{
		HTTPClient: client,
	}
}

/* VerifyMfaParams contains all the parameters to send to the API endpoint
   for the verify mfa operation.

   Typically these are written to a http.Request.
*/
type VerifyMfaParams struct {

	/* MfaValidation.

	   An MFA validation request
	*/
	MfaValidation *rest_model.MfaCode

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify mfa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyMfaParams) WithDefaults() *VerifyMfaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify mfa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyMfaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify mfa params
func (o *VerifyMfaParams) WithTimeout(timeout time.Duration) *VerifyMfaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify mfa params
func (o *VerifyMfaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify mfa params
func (o *VerifyMfaParams) WithContext(ctx context.Context) *VerifyMfaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify mfa params
func (o *VerifyMfaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify mfa params
func (o *VerifyMfaParams) WithHTTPClient(client *http.Client) *VerifyMfaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify mfa params
func (o *VerifyMfaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMfaValidation adds the mfaValidation to the verify mfa params
func (o *VerifyMfaParams) WithMfaValidation(mfaValidation *rest_model.MfaCode) *VerifyMfaParams {
	o.SetMfaValidation(mfaValidation)
	return o
}

// SetMfaValidation adds the mfaValidation to the verify mfa params
func (o *VerifyMfaParams) SetMfaValidation(mfaValidation *rest_model.MfaCode) {
	o.MfaValidation = mfaValidation
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyMfaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.MfaValidation != nil {
		if err := r.SetBodyParam(o.MfaValidation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

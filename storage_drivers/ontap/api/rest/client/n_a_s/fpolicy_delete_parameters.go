// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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
)

// NewFpolicyDeleteParams creates a new FpolicyDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyDeleteParams() *FpolicyDeleteParams {
	return &FpolicyDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyDeleteParamsWithTimeout creates a new FpolicyDeleteParams object
// with the ability to set a timeout on a request.
func NewFpolicyDeleteParamsWithTimeout(timeout time.Duration) *FpolicyDeleteParams {
	return &FpolicyDeleteParams{
		timeout: timeout,
	}
}

// NewFpolicyDeleteParamsWithContext creates a new FpolicyDeleteParams object
// with the ability to set a context for a request.
func NewFpolicyDeleteParamsWithContext(ctx context.Context) *FpolicyDeleteParams {
	return &FpolicyDeleteParams{
		Context: ctx,
	}
}

// NewFpolicyDeleteParamsWithHTTPClient creates a new FpolicyDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyDeleteParamsWithHTTPClient(client *http.Client) *FpolicyDeleteParams {
	return &FpolicyDeleteParams{
		HTTPClient: client,
	}
}

/* FpolicyDeleteParams contains all the parameters to send to the API endpoint
   for the fpolicy delete operation.

   Typically these are written to a http.Request.
*/
type FpolicyDeleteParams struct {

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyDeleteParams) WithDefaults() *FpolicyDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy delete params
func (o *FpolicyDeleteParams) WithTimeout(timeout time.Duration) *FpolicyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy delete params
func (o *FpolicyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy delete params
func (o *FpolicyDeleteParams) WithContext(ctx context.Context) *FpolicyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy delete params
func (o *FpolicyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy delete params
func (o *FpolicyDeleteParams) WithHTTPClient(client *http.Client) *FpolicyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy delete params
func (o *FpolicyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSVMUUIDPathParameter adds the svmUUID to the fpolicy delete params
func (o *FpolicyDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *FpolicyDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fpolicy delete params
func (o *FpolicyDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

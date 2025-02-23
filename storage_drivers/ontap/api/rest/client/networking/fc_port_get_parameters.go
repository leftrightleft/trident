// Code generated by go-swagger; DO NOT EDIT.

package networking

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
	"github.com/go-openapi/swag"
)

// NewFcPortGetParams creates a new FcPortGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcPortGetParams() *FcPortGetParams {
	return &FcPortGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcPortGetParamsWithTimeout creates a new FcPortGetParams object
// with the ability to set a timeout on a request.
func NewFcPortGetParamsWithTimeout(timeout time.Duration) *FcPortGetParams {
	return &FcPortGetParams{
		timeout: timeout,
	}
}

// NewFcPortGetParamsWithContext creates a new FcPortGetParams object
// with the ability to set a context for a request.
func NewFcPortGetParamsWithContext(ctx context.Context) *FcPortGetParams {
	return &FcPortGetParams{
		Context: ctx,
	}
}

// NewFcPortGetParamsWithHTTPClient creates a new FcPortGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcPortGetParamsWithHTTPClient(client *http.Client) *FcPortGetParams {
	return &FcPortGetParams{
		HTTPClient: client,
	}
}

/* FcPortGetParams contains all the parameters to send to the API endpoint
   for the fc port get operation.

   Typically these are written to a http.Request.
*/
type FcPortGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   The unique identifier for the FC port.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc port get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortGetParams) WithDefaults() *FcPortGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc port get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcPortGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fc port get params
func (o *FcPortGetParams) WithTimeout(timeout time.Duration) *FcPortGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc port get params
func (o *FcPortGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc port get params
func (o *FcPortGetParams) WithContext(ctx context.Context) *FcPortGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc port get params
func (o *FcPortGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc port get params
func (o *FcPortGetParams) WithHTTPClient(client *http.Client) *FcPortGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc port get params
func (o *FcPortGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the fc port get params
func (o *FcPortGetParams) WithFieldsQueryParameter(fields []string) *FcPortGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the fc port get params
func (o *FcPortGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the fc port get params
func (o *FcPortGetParams) WithUUIDPathParameter(uuid string) *FcPortGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the fc port get params
func (o *FcPortGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *FcPortGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFcPortGet binds the parameter fields
func (o *FcPortGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

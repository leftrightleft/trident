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
	"github.com/go-openapi/swag"
)

// NewCifsSessionGetParams creates a new CifsSessionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSessionGetParams() *CifsSessionGetParams {
	return &CifsSessionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSessionGetParamsWithTimeout creates a new CifsSessionGetParams object
// with the ability to set a timeout on a request.
func NewCifsSessionGetParamsWithTimeout(timeout time.Duration) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		timeout: timeout,
	}
}

// NewCifsSessionGetParamsWithContext creates a new CifsSessionGetParams object
// with the ability to set a context for a request.
func NewCifsSessionGetParamsWithContext(ctx context.Context) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		Context: ctx,
	}
}

// NewCifsSessionGetParamsWithHTTPClient creates a new CifsSessionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSessionGetParamsWithHTTPClient(client *http.Client) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		HTTPClient: client,
	}
}

/* CifsSessionGetParams contains all the parameters to send to the API endpoint
   for the cifs session get operation.

   Typically these are written to a http.Request.
*/
type CifsSessionGetParams struct {

	/* ConnectionID.

	   Unique identifier for the SMB connection.
	*/
	ConnectionIDPathParameter int64

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Identifier.

	   Unique identifier for the SMB session.
	*/
	IdentifierPathParameter int64

	/* NodeUUID.

	   Node UUID.
	*/
	NodeUUIDPathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs session get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionGetParams) WithDefaults() *CifsSessionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs session get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs session get params
func (o *CifsSessionGetParams) WithTimeout(timeout time.Duration) *CifsSessionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs session get params
func (o *CifsSessionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs session get params
func (o *CifsSessionGetParams) WithContext(ctx context.Context) *CifsSessionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs session get params
func (o *CifsSessionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs session get params
func (o *CifsSessionGetParams) WithHTTPClient(client *http.Client) *CifsSessionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs session get params
func (o *CifsSessionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionIDPathParameter adds the connectionID to the cifs session get params
func (o *CifsSessionGetParams) WithConnectionIDPathParameter(connectionID int64) *CifsSessionGetParams {
	o.SetConnectionIDPathParameter(connectionID)
	return o
}

// SetConnectionIDPathParameter adds the connectionId to the cifs session get params
func (o *CifsSessionGetParams) SetConnectionIDPathParameter(connectionID int64) {
	o.ConnectionIDPathParameter = connectionID
}

// WithFieldsQueryParameter adds the fields to the cifs session get params
func (o *CifsSessionGetParams) WithFieldsQueryParameter(fields []string) *CifsSessionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs session get params
func (o *CifsSessionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIdentifierPathParameter adds the identifier to the cifs session get params
func (o *CifsSessionGetParams) WithIdentifierPathParameter(identifier int64) *CifsSessionGetParams {
	o.SetIdentifierPathParameter(identifier)
	return o
}

// SetIdentifierPathParameter adds the identifier to the cifs session get params
func (o *CifsSessionGetParams) SetIdentifierPathParameter(identifier int64) {
	o.IdentifierPathParameter = identifier
}

// WithNodeUUIDPathParameter adds the nodeUUID to the cifs session get params
func (o *CifsSessionGetParams) WithNodeUUIDPathParameter(nodeUUID string) *CifsSessionGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the cifs session get params
func (o *CifsSessionGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs session get params
func (o *CifsSessionGetParams) WithSVMUUIDPathParameter(svmUUID string) *CifsSessionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs session get params
func (o *CifsSessionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSessionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", swag.FormatInt64(o.ConnectionIDPathParameter)); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param identifier
	if err := r.SetPathParam("identifier", swag.FormatInt64(o.IdentifierPathParameter)); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsSessionGet binds the parameter fields
func (o *CifsSessionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

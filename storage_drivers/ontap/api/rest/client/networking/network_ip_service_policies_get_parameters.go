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

// NewNetworkIPServicePoliciesGetParams creates a new NetworkIPServicePoliciesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPServicePoliciesGetParams() *NetworkIPServicePoliciesGetParams {
	return &NetworkIPServicePoliciesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPServicePoliciesGetParamsWithTimeout creates a new NetworkIPServicePoliciesGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPServicePoliciesGetParamsWithTimeout(timeout time.Duration) *NetworkIPServicePoliciesGetParams {
	return &NetworkIPServicePoliciesGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPServicePoliciesGetParamsWithContext creates a new NetworkIPServicePoliciesGetParams object
// with the ability to set a context for a request.
func NewNetworkIPServicePoliciesGetParamsWithContext(ctx context.Context) *NetworkIPServicePoliciesGetParams {
	return &NetworkIPServicePoliciesGetParams{
		Context: ctx,
	}
}

// NewNetworkIPServicePoliciesGetParamsWithHTTPClient creates a new NetworkIPServicePoliciesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPServicePoliciesGetParamsWithHTTPClient(client *http.Client) *NetworkIPServicePoliciesGetParams {
	return &NetworkIPServicePoliciesGetParams{
		HTTPClient: client,
	}
}

/* NetworkIPServicePoliciesGetParams contains all the parameters to send to the API endpoint
   for the network ip service policies get operation.

   Typically these are written to a http.Request.
*/
type NetworkIPServicePoliciesGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceNameQueryParameter *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUIDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* Services.

	   Filter by services
	*/
	ServicesQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip service policies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPServicePoliciesGetParams) WithDefaults() *NetworkIPServicePoliciesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip service policies get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPServicePoliciesGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NetworkIPServicePoliciesGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithTimeout(timeout time.Duration) *NetworkIPServicePoliciesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithContext(ctx context.Context) *NetworkIPServicePoliciesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithHTTPClient(client *http.Client) *NetworkIPServicePoliciesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithFieldsQueryParameter(fields []string) *NetworkIPServicePoliciesGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIpspaceNameQueryParameter adds the ipspaceName to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithIpspaceNameQueryParameter(ipspaceName *string) *NetworkIPServicePoliciesGetParams {
	o.SetIpspaceNameQueryParameter(ipspaceName)
	return o
}

// SetIpspaceNameQueryParameter adds the ipspaceName to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetIpspaceNameQueryParameter(ipspaceName *string) {
	o.IpspaceNameQueryParameter = ipspaceName
}

// WithIpspaceUUIDQueryParameter adds the ipspaceUUID to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithIpspaceUUIDQueryParameter(ipspaceUUID *string) *NetworkIPServicePoliciesGetParams {
	o.SetIpspaceUUIDQueryParameter(ipspaceUUID)
	return o
}

// SetIpspaceUUIDQueryParameter adds the ipspaceUuid to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetIpspaceUUIDQueryParameter(ipspaceUUID *string) {
	o.IpspaceUUIDQueryParameter = ipspaceUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NetworkIPServicePoliciesGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithNameQueryParameter(name *string) *NetworkIPServicePoliciesGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithOrderByQueryParameter(orderBy []string) *NetworkIPServicePoliciesGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NetworkIPServicePoliciesGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NetworkIPServicePoliciesGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithScopeQueryParameter(scope *string) *NetworkIPServicePoliciesGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithServicesQueryParameter adds the services to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithServicesQueryParameter(services *string) *NetworkIPServicePoliciesGetParams {
	o.SetServicesQueryParameter(services)
	return o
}

// SetServicesQueryParameter adds the services to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetServicesQueryParameter(services *string) {
	o.ServicesQueryParameter = services
}

// WithSVMNameQueryParameter adds the svmName to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithSVMNameQueryParameter(svmName *string) *NetworkIPServicePoliciesGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NetworkIPServicePoliciesGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) WithUUIDQueryParameter(uuid *string) *NetworkIPServicePoliciesGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the network ip service policies get params
func (o *NetworkIPServicePoliciesGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPServicePoliciesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IpspaceNameQueryParameter != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceNameQueryParameter != nil {
			qrIpspaceName = *o.IpspaceNameQueryParameter
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUIDQueryParameter != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUIDQueryParameter != nil {
			qrIpspaceUUID = *o.IpspaceUUIDQueryParameter
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.ServicesQueryParameter != nil {

		// query param services
		var qrServices string

		if o.ServicesQueryParameter != nil {
			qrServices = *o.ServicesQueryParameter
		}
		qServices := qrServices
		if qServices != "" {

			if err := r.SetQueryParam("services", qServices); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetworkIPServicePoliciesGet binds the parameter fields
func (o *NetworkIPServicePoliciesGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNetworkIPServicePoliciesGet binds the parameter order_by
func (o *NetworkIPServicePoliciesGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}

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

// NewKerberosRealmCollectionGetParams creates a new KerberosRealmCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKerberosRealmCollectionGetParams() *KerberosRealmCollectionGetParams {
	return &KerberosRealmCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKerberosRealmCollectionGetParamsWithTimeout creates a new KerberosRealmCollectionGetParams object
// with the ability to set a timeout on a request.
func NewKerberosRealmCollectionGetParamsWithTimeout(timeout time.Duration) *KerberosRealmCollectionGetParams {
	return &KerberosRealmCollectionGetParams{
		timeout: timeout,
	}
}

// NewKerberosRealmCollectionGetParamsWithContext creates a new KerberosRealmCollectionGetParams object
// with the ability to set a context for a request.
func NewKerberosRealmCollectionGetParamsWithContext(ctx context.Context) *KerberosRealmCollectionGetParams {
	return &KerberosRealmCollectionGetParams{
		Context: ctx,
	}
}

// NewKerberosRealmCollectionGetParamsWithHTTPClient creates a new KerberosRealmCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewKerberosRealmCollectionGetParamsWithHTTPClient(client *http.Client) *KerberosRealmCollectionGetParams {
	return &KerberosRealmCollectionGetParams{
		HTTPClient: client,
	}
}

/* KerberosRealmCollectionGetParams contains all the parameters to send to the API endpoint
   for the kerberos realm collection get operation.

   Typically these are written to a http.Request.
*/
type KerberosRealmCollectionGetParams struct {

	/* AdServerAddress.

	   Filter by ad_server.address
	*/
	AdServerAddressQueryParameter *string

	/* AdServerName.

	   Filter by ad_server.name
	*/
	AdServerNameQueryParameter *string

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* EncryptionTypes.

	   Filter by encryption_types
	*/
	EncryptionTypesQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* KdcIP.

	   Filter by kdc.ip
	*/
	KdcIPQueryParameter *string

	/* KdcPort.

	   Filter by kdc.port
	*/
	KdcPortQueryParameter *int64

	/* KdcVendor.

	   Filter by kdc.vendor
	*/
	KdcVendorQueryParameter *string

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kerberos realm collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmCollectionGetParams) WithDefaults() *KerberosRealmCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kerberos realm collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosRealmCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := KerberosRealmCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithTimeout(timeout time.Duration) *KerberosRealmCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithContext(ctx context.Context) *KerberosRealmCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithHTTPClient(client *http.Client) *KerberosRealmCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdServerAddressQueryParameter adds the adServerAddress to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithAdServerAddressQueryParameter(adServerAddress *string) *KerberosRealmCollectionGetParams {
	o.SetAdServerAddressQueryParameter(adServerAddress)
	return o
}

// SetAdServerAddressQueryParameter adds the adServerAddress to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetAdServerAddressQueryParameter(adServerAddress *string) {
	o.AdServerAddressQueryParameter = adServerAddress
}

// WithAdServerNameQueryParameter adds the adServerName to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithAdServerNameQueryParameter(adServerName *string) *KerberosRealmCollectionGetParams {
	o.SetAdServerNameQueryParameter(adServerName)
	return o
}

// SetAdServerNameQueryParameter adds the adServerName to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetAdServerNameQueryParameter(adServerName *string) {
	o.AdServerNameQueryParameter = adServerName
}

// WithCommentQueryParameter adds the comment to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithCommentQueryParameter(comment *string) *KerberosRealmCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithEncryptionTypesQueryParameter adds the encryptionTypes to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithEncryptionTypesQueryParameter(encryptionTypes *string) *KerberosRealmCollectionGetParams {
	o.SetEncryptionTypesQueryParameter(encryptionTypes)
	return o
}

// SetEncryptionTypesQueryParameter adds the encryptionTypes to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetEncryptionTypesQueryParameter(encryptionTypes *string) {
	o.EncryptionTypesQueryParameter = encryptionTypes
}

// WithFieldsQueryParameter adds the fields to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithFieldsQueryParameter(fields []string) *KerberosRealmCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithKdcIPQueryParameter adds the kdcIP to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithKdcIPQueryParameter(kdcIP *string) *KerberosRealmCollectionGetParams {
	o.SetKdcIPQueryParameter(kdcIP)
	return o
}

// SetKdcIPQueryParameter adds the kdcIp to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetKdcIPQueryParameter(kdcIP *string) {
	o.KdcIPQueryParameter = kdcIP
}

// WithKdcPortQueryParameter adds the kdcPort to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithKdcPortQueryParameter(kdcPort *int64) *KerberosRealmCollectionGetParams {
	o.SetKdcPortQueryParameter(kdcPort)
	return o
}

// SetKdcPortQueryParameter adds the kdcPort to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetKdcPortQueryParameter(kdcPort *int64) {
	o.KdcPortQueryParameter = kdcPort
}

// WithKdcVendorQueryParameter adds the kdcVendor to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithKdcVendorQueryParameter(kdcVendor *string) *KerberosRealmCollectionGetParams {
	o.SetKdcVendorQueryParameter(kdcVendor)
	return o
}

// SetKdcVendorQueryParameter adds the kdcVendor to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetKdcVendorQueryParameter(kdcVendor *string) {
	o.KdcVendorQueryParameter = kdcVendor
}

// WithMaxRecordsQueryParameter adds the maxRecords to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *KerberosRealmCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithNameQueryParameter(name *string) *KerberosRealmCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *KerberosRealmCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *KerberosRealmCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *KerberosRealmCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *KerberosRealmCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *KerberosRealmCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the kerberos realm collection get params
func (o *KerberosRealmCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *KerberosRealmCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdServerAddressQueryParameter != nil {

		// query param ad_server.address
		var qrAdServerAddress string

		if o.AdServerAddressQueryParameter != nil {
			qrAdServerAddress = *o.AdServerAddressQueryParameter
		}
		qAdServerAddress := qrAdServerAddress
		if qAdServerAddress != "" {

			if err := r.SetQueryParam("ad_server.address", qAdServerAddress); err != nil {
				return err
			}
		}
	}

	if o.AdServerNameQueryParameter != nil {

		// query param ad_server.name
		var qrAdServerName string

		if o.AdServerNameQueryParameter != nil {
			qrAdServerName = *o.AdServerNameQueryParameter
		}
		qAdServerName := qrAdServerName
		if qAdServerName != "" {

			if err := r.SetQueryParam("ad_server.name", qAdServerName); err != nil {
				return err
			}
		}
	}

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.EncryptionTypesQueryParameter != nil {

		// query param encryption_types
		var qrEncryptionTypes string

		if o.EncryptionTypesQueryParameter != nil {
			qrEncryptionTypes = *o.EncryptionTypesQueryParameter
		}
		qEncryptionTypes := qrEncryptionTypes
		if qEncryptionTypes != "" {

			if err := r.SetQueryParam("encryption_types", qEncryptionTypes); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.KdcIPQueryParameter != nil {

		// query param kdc.ip
		var qrKdcIP string

		if o.KdcIPQueryParameter != nil {
			qrKdcIP = *o.KdcIPQueryParameter
		}
		qKdcIP := qrKdcIP
		if qKdcIP != "" {

			if err := r.SetQueryParam("kdc.ip", qKdcIP); err != nil {
				return err
			}
		}
	}

	if o.KdcPortQueryParameter != nil {

		// query param kdc.port
		var qrKdcPort int64

		if o.KdcPortQueryParameter != nil {
			qrKdcPort = *o.KdcPortQueryParameter
		}
		qKdcPort := swag.FormatInt64(qrKdcPort)
		if qKdcPort != "" {

			if err := r.SetQueryParam("kdc.port", qKdcPort); err != nil {
				return err
			}
		}
	}

	if o.KdcVendorQueryParameter != nil {

		// query param kdc.vendor
		var qrKdcVendor string

		if o.KdcVendorQueryParameter != nil {
			qrKdcVendor = *o.KdcVendorQueryParameter
		}
		qKdcVendor := qrKdcVendor
		if qKdcVendor != "" {

			if err := r.SetQueryParam("kdc.vendor", qKdcVendor); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamKerberosRealmCollectionGet binds the parameter fields
func (o *KerberosRealmCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamKerberosRealmCollectionGet binds the parameter order_by
func (o *KerberosRealmCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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

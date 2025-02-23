// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3GroupDeleteReader is a Reader for the S3GroupDelete structure.
type S3GroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3GroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3GroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3GroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3GroupDeleteOK creates a S3GroupDeleteOK with default headers values
func NewS3GroupDeleteOK() *S3GroupDeleteOK {
	return &S3GroupDeleteOK{}
}

/* S3GroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type S3GroupDeleteOK struct {
}

func (o *S3GroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3GroupDeleteOK ", 200)
}

func (o *S3GroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS3GroupDeleteDefault creates a S3GroupDeleteDefault with default headers values
func NewS3GroupDeleteDefault(code int) *S3GroupDeleteDefault {
	return &S3GroupDeleteDefault{
		_statusCode: code,
	}
}

/* S3GroupDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type S3GroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 group delete default response
func (o *S3GroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *S3GroupDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3_group_delete default  %+v", o._statusCode, o.Payload)
}
func (o *S3GroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3GroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// S3GroupGetReader is a Reader for the S3GroupGet structure.
type S3GroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3GroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3GroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3GroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3GroupGetOK creates a S3GroupGetOK with default headers values
func NewS3GroupGetOK() *S3GroupGetOK {
	return &S3GroupGetOK{}
}

/* S3GroupGetOK describes a response with status code 200, with default header values.

OK
*/
type S3GroupGetOK struct {
	Payload *models.S3Group
}

func (o *S3GroupGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3GroupGetOK  %+v", 200, o.Payload)
}
func (o *S3GroupGetOK) GetPayload() *models.S3Group {
	return o.Payload
}

func (o *S3GroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3GroupGetDefault creates a S3GroupGetDefault with default headers values
func NewS3GroupGetDefault(code int) *S3GroupGetDefault {
	return &S3GroupGetDefault{
		_statusCode: code,
	}
}

/* S3GroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3GroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 group get default response
func (o *S3GroupGetDefault) Code() int {
	return o._statusCode
}

func (o *S3GroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups/{id}][%d] s3_group_get default  %+v", o._statusCode, o.Payload)
}
func (o *S3GroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3GroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LocalCifsGroupCreateReader is a Reader for the LocalCifsGroupCreate structure.
type LocalCifsGroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLocalCifsGroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupCreateCreated creates a LocalCifsGroupCreateCreated with default headers values
func NewLocalCifsGroupCreateCreated() *LocalCifsGroupCreateCreated {
	return &LocalCifsGroupCreateCreated{}
}

/* LocalCifsGroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LocalCifsGroupCreateCreated struct {
}

func (o *LocalCifsGroupCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/local-groups][%d] localCifsGroupCreateCreated ", 201)
}

func (o *LocalCifsGroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsGroupCreateDefault creates a LocalCifsGroupCreateDefault with default headers values
func NewLocalCifsGroupCreateDefault(code int) *LocalCifsGroupCreateDefault {
	return &LocalCifsGroupCreateDefault{
		_statusCode: code,
	}
}

/* LocalCifsGroupCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 262278     | Name is a required field. |
| 655399     | CIFS server must exist to create a local group. |
| 655660     | The operation is allowed only on data SVMs. |
| 655661     | The group name and description should not exceed 256 characters. |
| 655668     | The specified group name contains illegal characters. |
| 655675     | The local domain name specified in the group name does not exist. |
| 655677     | This operation does not allow for the creation of a group in the BUILTIN domain. |
| 655682     | The group name cannot be blank. |
| 655717     | The specified group name already exists. |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name. |

*/
type LocalCifsGroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs group create default response
func (o *LocalCifsGroupCreateDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsGroupCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/local-groups][%d] local_cifs_group_create default  %+v", o._statusCode, o.Payload)
}
func (o *LocalCifsGroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// ColumnFamilyNameKeyspaceGetReader is a Reader for the ColumnFamilyNameKeyspaceGet structure.
type ColumnFamilyNameKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyNameKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyNameKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyNameKeyspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyNameKeyspaceGetOK creates a ColumnFamilyNameKeyspaceGetOK with default headers values
func NewColumnFamilyNameKeyspaceGetOK() *ColumnFamilyNameKeyspaceGetOK {
	return &ColumnFamilyNameKeyspaceGetOK{}
}

/*ColumnFamilyNameKeyspaceGetOK handles this case with default header values.

ColumnFamilyNameKeyspaceGetOK column family name keyspace get o k
*/
type ColumnFamilyNameKeyspaceGetOK struct {
	Payload []string
}

func (o *ColumnFamilyNameKeyspaceGetOK) GetPayload() []string {
	return o.Payload
}

func (o *ColumnFamilyNameKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyNameKeyspaceGetDefault creates a ColumnFamilyNameKeyspaceGetDefault with default headers values
func NewColumnFamilyNameKeyspaceGetDefault(code int) *ColumnFamilyNameKeyspaceGetDefault {
	return &ColumnFamilyNameKeyspaceGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyNameKeyspaceGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyNameKeyspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family name keyspace get default response
func (o *ColumnFamilyNameKeyspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyNameKeyspaceGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyNameKeyspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyNameKeyspaceGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

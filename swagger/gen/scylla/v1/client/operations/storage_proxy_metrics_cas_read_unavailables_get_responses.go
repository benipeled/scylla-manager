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

// StorageProxyMetricsCasReadUnavailablesGetReader is a Reader for the StorageProxyMetricsCasReadUnavailablesGet structure.
type StorageProxyMetricsCasReadUnavailablesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasReadUnavailablesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasReadUnavailablesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsCasReadUnavailablesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsCasReadUnavailablesGetOK creates a StorageProxyMetricsCasReadUnavailablesGetOK with default headers values
func NewStorageProxyMetricsCasReadUnavailablesGetOK() *StorageProxyMetricsCasReadUnavailablesGetOK {
	return &StorageProxyMetricsCasReadUnavailablesGetOK{}
}

/*StorageProxyMetricsCasReadUnavailablesGetOK handles this case with default header values.

StorageProxyMetricsCasReadUnavailablesGetOK storage proxy metrics cas read unavailables get o k
*/
type StorageProxyMetricsCasReadUnavailablesGetOK struct {
	Payload interface{}
}

func (o *StorageProxyMetricsCasReadUnavailablesGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadUnavailablesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsCasReadUnavailablesGetDefault creates a StorageProxyMetricsCasReadUnavailablesGetDefault with default headers values
func NewStorageProxyMetricsCasReadUnavailablesGetDefault(code int) *StorageProxyMetricsCasReadUnavailablesGetDefault {
	return &StorageProxyMetricsCasReadUnavailablesGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsCasReadUnavailablesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsCasReadUnavailablesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics cas read unavailables get default response
func (o *StorageProxyMetricsCasReadUnavailablesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsCasReadUnavailablesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadUnavailablesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsCasReadUnavailablesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

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

// CacheServiceMetricsRowHitRateGetReader is a Reader for the CacheServiceMetricsRowHitRateGet structure.
type CacheServiceMetricsRowHitRateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsRowHitRateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsRowHitRateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsRowHitRateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsRowHitRateGetOK creates a CacheServiceMetricsRowHitRateGetOK with default headers values
func NewCacheServiceMetricsRowHitRateGetOK() *CacheServiceMetricsRowHitRateGetOK {
	return &CacheServiceMetricsRowHitRateGetOK{}
}

/*CacheServiceMetricsRowHitRateGetOK handles this case with default header values.

CacheServiceMetricsRowHitRateGetOK cache service metrics row hit rate get o k
*/
type CacheServiceMetricsRowHitRateGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsRowHitRateGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsRowHitRateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsRowHitRateGetDefault creates a CacheServiceMetricsRowHitRateGetDefault with default headers values
func NewCacheServiceMetricsRowHitRateGetDefault(code int) *CacheServiceMetricsRowHitRateGetDefault {
	return &CacheServiceMetricsRowHitRateGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsRowHitRateGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsRowHitRateGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics row hit rate get default response
func (o *CacheServiceMetricsRowHitRateGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsRowHitRateGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsRowHitRateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsRowHitRateGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

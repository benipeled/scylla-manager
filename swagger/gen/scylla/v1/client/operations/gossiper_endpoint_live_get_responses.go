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

// GossiperEndpointLiveGetReader is a Reader for the GossiperEndpointLiveGet structure.
type GossiperEndpointLiveGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GossiperEndpointLiveGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGossiperEndpointLiveGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGossiperEndpointLiveGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGossiperEndpointLiveGetOK creates a GossiperEndpointLiveGetOK with default headers values
func NewGossiperEndpointLiveGetOK() *GossiperEndpointLiveGetOK {
	return &GossiperEndpointLiveGetOK{}
}

/*GossiperEndpointLiveGetOK handles this case with default header values.

GossiperEndpointLiveGetOK gossiper endpoint live get o k
*/
type GossiperEndpointLiveGetOK struct {
	Payload []string
}

func (o *GossiperEndpointLiveGetOK) GetPayload() []string {
	return o.Payload
}

func (o *GossiperEndpointLiveGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGossiperEndpointLiveGetDefault creates a GossiperEndpointLiveGetDefault with default headers values
func NewGossiperEndpointLiveGetDefault(code int) *GossiperEndpointLiveGetDefault {
	return &GossiperEndpointLiveGetDefault{
		_statusCode: code,
	}
}

/*GossiperEndpointLiveGetDefault handles this case with default header values.

internal server error
*/
type GossiperEndpointLiveGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the gossiper endpoint live get default response
func (o *GossiperEndpointLiveGetDefault) Code() int {
	return o._statusCode
}

func (o *GossiperEndpointLiveGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GossiperEndpointLiveGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *GossiperEndpointLiveGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

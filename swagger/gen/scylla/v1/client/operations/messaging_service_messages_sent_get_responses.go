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

// MessagingServiceMessagesSentGetReader is a Reader for the MessagingServiceMessagesSentGet structure.
type MessagingServiceMessagesSentGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesSentGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMessagingServiceMessagesSentGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMessagingServiceMessagesSentGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMessagingServiceMessagesSentGetOK creates a MessagingServiceMessagesSentGetOK with default headers values
func NewMessagingServiceMessagesSentGetOK() *MessagingServiceMessagesSentGetOK {
	return &MessagingServiceMessagesSentGetOK{}
}

/*MessagingServiceMessagesSentGetOK handles this case with default header values.

MessagingServiceMessagesSentGetOK messaging service messages sent get o k
*/
type MessagingServiceMessagesSentGetOK struct {
	Payload []*models.MessageCounter
}

func (o *MessagingServiceMessagesSentGetOK) GetPayload() []*models.MessageCounter {
	return o.Payload
}

func (o *MessagingServiceMessagesSentGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMessagingServiceMessagesSentGetDefault creates a MessagingServiceMessagesSentGetDefault with default headers values
func NewMessagingServiceMessagesSentGetDefault(code int) *MessagingServiceMessagesSentGetDefault {
	return &MessagingServiceMessagesSentGetDefault{
		_statusCode: code,
	}
}

/*MessagingServiceMessagesSentGetDefault handles this case with default header values.

internal server error
*/
type MessagingServiceMessagesSentGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the messaging service messages sent get default response
func (o *MessagingServiceMessagesSentGetDefault) Code() int {
	return o._statusCode
}

func (o *MessagingServiceMessagesSentGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *MessagingServiceMessagesSentGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *MessagingServiceMessagesSentGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

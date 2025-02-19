// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigDataFileDirectoriesReader is a Reader for the FindConfigDataFileDirectories structure.
type FindConfigDataFileDirectoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDataFileDirectoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDataFileDirectoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDataFileDirectoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDataFileDirectoriesOK creates a FindConfigDataFileDirectoriesOK with default headers values
func NewFindConfigDataFileDirectoriesOK() *FindConfigDataFileDirectoriesOK {
	return &FindConfigDataFileDirectoriesOK{}
}

/*FindConfigDataFileDirectoriesOK handles this case with default header values.

Config value
*/
type FindConfigDataFileDirectoriesOK struct {
	Payload []string
}

func (o *FindConfigDataFileDirectoriesOK) GetPayload() []string {
	return o.Payload
}

func (o *FindConfigDataFileDirectoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDataFileDirectoriesDefault creates a FindConfigDataFileDirectoriesDefault with default headers values
func NewFindConfigDataFileDirectoriesDefault(code int) *FindConfigDataFileDirectoriesDefault {
	return &FindConfigDataFileDirectoriesDefault{
		_statusCode: code,
	}
}

/*FindConfigDataFileDirectoriesDefault handles this case with default header values.

unexpected error
*/
type FindConfigDataFileDirectoriesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config data file directories default response
func (o *FindConfigDataFileDirectoriesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDataFileDirectoriesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDataFileDirectoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDataFileDirectoriesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

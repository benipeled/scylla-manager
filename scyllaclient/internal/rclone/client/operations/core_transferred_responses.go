// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/rclone/models"
)

// CoreTransferredReader is a Reader for the CoreTransferred structure.
type CoreTransferredReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreTransferredReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCoreTransferredOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewCoreTransferredNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCoreTransferredInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCoreTransferredOK creates a CoreTransferredOK with default headers values
func NewCoreTransferredOK() *CoreTransferredOK {
	return &CoreTransferredOK{}
}

/*CoreTransferredOK handles this case with default header values.

Completed transfers
*/
type CoreTransferredOK struct {
	Payload *CoreTransferredOKBody
}

func (o *CoreTransferredOK) Error() string {
	return fmt.Sprintf("[POST /core/transferred][%d] coreTransferredOK  %+v", 200, o.Payload)
}

func (o *CoreTransferredOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CoreTransferredOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreTransferredNotFound creates a CoreTransferredNotFound with default headers values
func NewCoreTransferredNotFound() *CoreTransferredNotFound {
	return &CoreTransferredNotFound{}
}

/*CoreTransferredNotFound handles this case with default header values.

Not found
*/
type CoreTransferredNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CoreTransferredNotFound) Error() string {
	return fmt.Sprintf("[POST /core/transferred][%d] coreTransferredNotFound  %+v", 404, o.Payload)
}

func (o *CoreTransferredNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreTransferredInternalServerError creates a CoreTransferredInternalServerError with default headers values
func NewCoreTransferredInternalServerError() *CoreTransferredInternalServerError {
	return &CoreTransferredInternalServerError{}
}

/*CoreTransferredInternalServerError handles this case with default header values.

Server error
*/
type CoreTransferredInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CoreTransferredInternalServerError) Error() string {
	return fmt.Sprintf("[POST /core/transferred][%d] coreTransferredInternalServerError  %+v", 500, o.Payload)
}

func (o *CoreTransferredInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CoreTransferredOKBody core transferred o k body
swagger:model CoreTransferredOKBody
*/
type CoreTransferredOKBody struct {

	// transferred
	Transferred []*models.Transfer `json:"transferred"`
}

// Validate validates this core transferred o k body
func (o *CoreTransferredOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTransferred(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CoreTransferredOKBody) validateTransferred(formats strfmt.Registry) error {

	if swag.IsZero(o.Transferred) { // not required
		return nil
	}

	for i := 0; i < len(o.Transferred); i++ {
		if swag.IsZero(o.Transferred[i]) { // not required
			continue
		}

		if o.Transferred[i] != nil {
			if err := o.Transferred[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coreTransferredOK" + "." + "transferred" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CoreTransferredOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CoreTransferredOKBody) UnmarshalBinary(b []byte) error {
	var res CoreTransferredOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

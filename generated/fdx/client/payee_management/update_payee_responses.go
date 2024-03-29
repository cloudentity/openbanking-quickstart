// Code generated by go-swagger; DO NOT EDIT.

package payee_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// UpdatePayeeReader is a Reader for the UpdatePayee structure.
type UpdatePayeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePayeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePayeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePayeeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePayeeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdatePayeeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePayeeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewUpdatePayeeNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdatePayeeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePayeeOK creates a UpdatePayeeOK with default headers values
func NewUpdatePayeeOK() *UpdatePayeeOK {
	return &UpdatePayeeOK{}
}

/* UpdatePayeeOK describes a response with status code 200, with default header values.

Ok
*/
type UpdatePayeeOK struct {
	Payload *models.Payeeentity
}

func (o *UpdatePayeeOK) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeOK  %+v", 200, o.Payload)
}
func (o *UpdatePayeeOK) GetPayload() *models.Payeeentity {
	return o.Payload
}

func (o *UpdatePayeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Payeeentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeBadRequest creates a UpdatePayeeBadRequest with default headers values
func NewUpdatePayeeBadRequest() *UpdatePayeeBadRequest {
	return &UpdatePayeeBadRequest{}
}

/* UpdatePayeeBadRequest describes a response with status code 400, with default header values.

Input sent by client does not satisfy API specification
*/
type UpdatePayeeBadRequest struct {
	Payload *models.Error
}

func (o *UpdatePayeeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeBadRequest  %+v", 400, o.Payload)
}
func (o *UpdatePayeeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePayeeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeNotFound creates a UpdatePayeeNotFound with default headers values
func NewUpdatePayeeNotFound() *UpdatePayeeNotFound {
	return &UpdatePayeeNotFound{}
}

/* UpdatePayeeNotFound describes a response with status code 404, with default header values.

Payee with provided ID was not found
*/
type UpdatePayeeNotFound struct {
	Payload *models.Error
}

func (o *UpdatePayeeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeNotFound  %+v", 404, o.Payload)
}
func (o *UpdatePayeeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePayeeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeConflict creates a UpdatePayeeConflict with default headers values
func NewUpdatePayeeConflict() *UpdatePayeeConflict {
	return &UpdatePayeeConflict{}
}

/* UpdatePayeeConflict describes a response with status code 409, with default header values.

Duplicate Request
*/
type UpdatePayeeConflict struct {
	Payload *models.Payeeentity
}

func (o *UpdatePayeeConflict) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeConflict  %+v", 409, o.Payload)
}
func (o *UpdatePayeeConflict) GetPayload() *models.Payeeentity {
	return o.Payload
}

func (o *UpdatePayeeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Payeeentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeInternalServerError creates a UpdatePayeeInternalServerError with default headers values
func NewUpdatePayeeInternalServerError() *UpdatePayeeInternalServerError {
	return &UpdatePayeeInternalServerError{}
}

/* UpdatePayeeInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type UpdatePayeeInternalServerError struct {
	Payload *models.Error
}

func (o *UpdatePayeeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdatePayeeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePayeeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeNotImplemented creates a UpdatePayeeNotImplemented with default headers values
func NewUpdatePayeeNotImplemented() *UpdatePayeeNotImplemented {
	return &UpdatePayeeNotImplemented{}
}

/* UpdatePayeeNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type UpdatePayeeNotImplemented struct {
	Payload *models.Error
}

func (o *UpdatePayeeNotImplemented) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeNotImplemented  %+v", 501, o.Payload)
}
func (o *UpdatePayeeNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePayeeNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePayeeServiceUnavailable creates a UpdatePayeeServiceUnavailable with default headers values
func NewUpdatePayeeServiceUnavailable() *UpdatePayeeServiceUnavailable {
	return &UpdatePayeeServiceUnavailable{}
}

/* UpdatePayeeServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type UpdatePayeeServiceUnavailable struct {
	Payload *models.Error
}

func (o *UpdatePayeeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /payees/{payeeId}][%d] updatePayeeServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UpdatePayeeServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePayeeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package personal_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// GetAccountContactReader is a Reader for the GetAccountContact structure.
type GetAccountContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAccountContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetAccountContactNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAccountContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountContactOK creates a GetAccountContactOK with default headers values
func NewGetAccountContactOK() *GetAccountContactOK {
	return &GetAccountContactOK{}
}

/* GetAccountContactOK describes a response with status code 200, with default header values.

Details used to verify an account
*/
type GetAccountContactOK struct {
	Payload *models.AccountContactentity
}

func (o *GetAccountContactOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/contact][%d] getAccountContactOK  %+v", 200, o.Payload)
}
func (o *GetAccountContactOK) GetPayload() *models.AccountContactentity {
	return o.Payload
}

func (o *GetAccountContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountContactentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountContactNotFound creates a GetAccountContactNotFound with default headers values
func NewGetAccountContactNotFound() *GetAccountContactNotFound {
	return &GetAccountContactNotFound{}
}

/* GetAccountContactNotFound describes a response with status code 404, with default header values.

Account with id not found
*/
type GetAccountContactNotFound struct {
	Payload *models.Error
}

func (o *GetAccountContactNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/contact][%d] getAccountContactNotFound  %+v", 404, o.Payload)
}
func (o *GetAccountContactNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountContactInternalServerError creates a GetAccountContactInternalServerError with default headers values
func NewGetAccountContactInternalServerError() *GetAccountContactInternalServerError {
	return &GetAccountContactInternalServerError{}
}

/* GetAccountContactInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type GetAccountContactInternalServerError struct {
	Payload *models.Error
}

func (o *GetAccountContactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/contact][%d] getAccountContactInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountContactInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountContactNotImplemented creates a GetAccountContactNotImplemented with default headers values
func NewGetAccountContactNotImplemented() *GetAccountContactNotImplemented {
	return &GetAccountContactNotImplemented{}
}

/* GetAccountContactNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type GetAccountContactNotImplemented struct {
	Payload *models.Error
}

func (o *GetAccountContactNotImplemented) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/contact][%d] getAccountContactNotImplemented  %+v", 501, o.Payload)
}
func (o *GetAccountContactNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountContactNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountContactServiceUnavailable creates a GetAccountContactServiceUnavailable with default headers values
func NewGetAccountContactServiceUnavailable() *GetAccountContactServiceUnavailable {
	return &GetAccountContactServiceUnavailable{}
}

/* GetAccountContactServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type GetAccountContactServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetAccountContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/contact][%d] getAccountContactServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAccountContactServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
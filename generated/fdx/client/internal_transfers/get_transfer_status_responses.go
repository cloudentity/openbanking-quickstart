// Code generated by go-swagger; DO NOT EDIT.

package internal_transfers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// GetTransferStatusReader is a Reader for the GetTransferStatus structure.
type GetTransferStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransferStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransferStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTransferStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTransferStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetTransferStatusNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTransferStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransferStatusOK creates a GetTransferStatusOK with default headers values
func NewGetTransferStatusOK() *GetTransferStatusOK {
	return &GetTransferStatusOK{}
}

/* GetTransferStatusOK describes a response with status code 200, with default header values.

Current status of the requested transfer
*/
type GetTransferStatusOK struct {
	Payload *models.TransferStatusentity
}

func (o *GetTransferStatusOK) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}/status][%d] getTransferStatusOK  %+v", 200, o.Payload)
}
func (o *GetTransferStatusOK) GetPayload() *models.TransferStatusentity {
	return o.Payload
}

func (o *GetTransferStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransferStatusentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransferStatusNotFound creates a GetTransferStatusNotFound with default headers values
func NewGetTransferStatusNotFound() *GetTransferStatusNotFound {
	return &GetTransferStatusNotFound{}
}

/* GetTransferStatusNotFound describes a response with status code 404, with default header values.

Transfer with id was not found
*/
type GetTransferStatusNotFound struct {
	Payload *models.Error
}

func (o *GetTransferStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}/status][%d] getTransferStatusNotFound  %+v", 404, o.Payload)
}
func (o *GetTransferStatusNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransferStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransferStatusInternalServerError creates a GetTransferStatusInternalServerError with default headers values
func NewGetTransferStatusInternalServerError() *GetTransferStatusInternalServerError {
	return &GetTransferStatusInternalServerError{}
}

/* GetTransferStatusInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type GetTransferStatusInternalServerError struct {
	Payload *models.Error
}

func (o *GetTransferStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}/status][%d] getTransferStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTransferStatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransferStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransferStatusNotImplemented creates a GetTransferStatusNotImplemented with default headers values
func NewGetTransferStatusNotImplemented() *GetTransferStatusNotImplemented {
	return &GetTransferStatusNotImplemented{}
}

/* GetTransferStatusNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type GetTransferStatusNotImplemented struct {
	Payload *models.Error
}

func (o *GetTransferStatusNotImplemented) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}/status][%d] getTransferStatusNotImplemented  %+v", 501, o.Payload)
}
func (o *GetTransferStatusNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransferStatusNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransferStatusServiceUnavailable creates a GetTransferStatusServiceUnavailable with default headers values
func NewGetTransferStatusServiceUnavailable() *GetTransferStatusServiceUnavailable {
	return &GetTransferStatusServiceUnavailable{}
}

/* GetTransferStatusServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type GetTransferStatusServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetTransferStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}/status][%d] getTransferStatusServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetTransferStatusServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransferStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

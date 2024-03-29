// Code generated by go-swagger; DO NOT EDIT.

package money_movement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// GetAccountPaymentNetworksReader is a Reader for the GetAccountPaymentNetworks structure.
type GetAccountPaymentNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountPaymentNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountPaymentNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAccountPaymentNetworksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountPaymentNetworksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetAccountPaymentNetworksNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAccountPaymentNetworksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountPaymentNetworksOK creates a GetAccountPaymentNetworksOK with default headers values
func NewGetAccountPaymentNetworksOK() *GetAccountPaymentNetworksOK {
	return &GetAccountPaymentNetworksOK{}
}

/* GetAccountPaymentNetworksOK describes a response with status code 200, with default header values.

Information required to execute a payment transaction against this account
*/
type GetAccountPaymentNetworksOK struct {
	Payload *models.Arrayofaccountpaymentnetworks
}

func (o *GetAccountPaymentNetworksOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/payment-networks][%d] getAccountPaymentNetworksOK  %+v", 200, o.Payload)
}
func (o *GetAccountPaymentNetworksOK) GetPayload() *models.Arrayofaccountpaymentnetworks {
	return o.Payload
}

func (o *GetAccountPaymentNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Arrayofaccountpaymentnetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPaymentNetworksNotFound creates a GetAccountPaymentNetworksNotFound with default headers values
func NewGetAccountPaymentNetworksNotFound() *GetAccountPaymentNetworksNotFound {
	return &GetAccountPaymentNetworksNotFound{}
}

/* GetAccountPaymentNetworksNotFound describes a response with status code 404, with default header values.

Account with id not found
*/
type GetAccountPaymentNetworksNotFound struct {
	Payload *models.Error
}

func (o *GetAccountPaymentNetworksNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/payment-networks][%d] getAccountPaymentNetworksNotFound  %+v", 404, o.Payload)
}
func (o *GetAccountPaymentNetworksNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountPaymentNetworksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPaymentNetworksInternalServerError creates a GetAccountPaymentNetworksInternalServerError with default headers values
func NewGetAccountPaymentNetworksInternalServerError() *GetAccountPaymentNetworksInternalServerError {
	return &GetAccountPaymentNetworksInternalServerError{}
}

/* GetAccountPaymentNetworksInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type GetAccountPaymentNetworksInternalServerError struct {
	Payload *models.Error
}

func (o *GetAccountPaymentNetworksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/payment-networks][%d] getAccountPaymentNetworksInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountPaymentNetworksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountPaymentNetworksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPaymentNetworksNotImplemented creates a GetAccountPaymentNetworksNotImplemented with default headers values
func NewGetAccountPaymentNetworksNotImplemented() *GetAccountPaymentNetworksNotImplemented {
	return &GetAccountPaymentNetworksNotImplemented{}
}

/* GetAccountPaymentNetworksNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type GetAccountPaymentNetworksNotImplemented struct {
	Payload *models.Error
}

func (o *GetAccountPaymentNetworksNotImplemented) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/payment-networks][%d] getAccountPaymentNetworksNotImplemented  %+v", 501, o.Payload)
}
func (o *GetAccountPaymentNetworksNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountPaymentNetworksNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountPaymentNetworksServiceUnavailable creates a GetAccountPaymentNetworksServiceUnavailable with default headers values
func NewGetAccountPaymentNetworksServiceUnavailable() *GetAccountPaymentNetworksServiceUnavailable {
	return &GetAccountPaymentNetworksServiceUnavailable{}
}

/* GetAccountPaymentNetworksServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type GetAccountPaymentNetworksServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetAccountPaymentNetworksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/payment-networks][%d] getAccountPaymentNetworksServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAccountPaymentNetworksServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountPaymentNetworksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

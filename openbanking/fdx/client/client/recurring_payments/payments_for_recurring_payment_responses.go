// Code generated by go-swagger; DO NOT EDIT.

package recurring_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
)

// PaymentsForRecurringPaymentReader is a Reader for the PaymentsForRecurringPayment structure.
type PaymentsForRecurringPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentsForRecurringPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentsForRecurringPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPaymentsForRecurringPaymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentsForRecurringPaymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPaymentsForRecurringPaymentNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPaymentsForRecurringPaymentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentsForRecurringPaymentOK creates a PaymentsForRecurringPaymentOK with default headers values
func NewPaymentsForRecurringPaymentOK() *PaymentsForRecurringPaymentOK {
	return &PaymentsForRecurringPaymentOK{}
}

/* PaymentsForRecurringPaymentOK describes a response with status code 200, with default header values.

Ok
*/
type PaymentsForRecurringPaymentOK struct {
	Payload *models.Paymentsentity
}

func (o *PaymentsForRecurringPaymentOK) Error() string {
	return fmt.Sprintf("[GET /recurring-payments/{recurringPaymentId}/payments][%d] paymentsForRecurringPaymentOK  %+v", 200, o.Payload)
}
func (o *PaymentsForRecurringPaymentOK) GetPayload() *models.Paymentsentity {
	return o.Payload
}

func (o *PaymentsForRecurringPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Paymentsentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsForRecurringPaymentNotFound creates a PaymentsForRecurringPaymentNotFound with default headers values
func NewPaymentsForRecurringPaymentNotFound() *PaymentsForRecurringPaymentNotFound {
	return &PaymentsForRecurringPaymentNotFound{}
}

/* PaymentsForRecurringPaymentNotFound describes a response with status code 404, with default header values.

A recurring payment with provided ID was not found
*/
type PaymentsForRecurringPaymentNotFound struct {
	Payload *models.Error
}

func (o *PaymentsForRecurringPaymentNotFound) Error() string {
	return fmt.Sprintf("[GET /recurring-payments/{recurringPaymentId}/payments][%d] paymentsForRecurringPaymentNotFound  %+v", 404, o.Payload)
}
func (o *PaymentsForRecurringPaymentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PaymentsForRecurringPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsForRecurringPaymentInternalServerError creates a PaymentsForRecurringPaymentInternalServerError with default headers values
func NewPaymentsForRecurringPaymentInternalServerError() *PaymentsForRecurringPaymentInternalServerError {
	return &PaymentsForRecurringPaymentInternalServerError{}
}

/* PaymentsForRecurringPaymentInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type PaymentsForRecurringPaymentInternalServerError struct {
	Payload *models.Error
}

func (o *PaymentsForRecurringPaymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /recurring-payments/{recurringPaymentId}/payments][%d] paymentsForRecurringPaymentInternalServerError  %+v", 500, o.Payload)
}
func (o *PaymentsForRecurringPaymentInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PaymentsForRecurringPaymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsForRecurringPaymentNotImplemented creates a PaymentsForRecurringPaymentNotImplemented with default headers values
func NewPaymentsForRecurringPaymentNotImplemented() *PaymentsForRecurringPaymentNotImplemented {
	return &PaymentsForRecurringPaymentNotImplemented{}
}

/* PaymentsForRecurringPaymentNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type PaymentsForRecurringPaymentNotImplemented struct {
	Payload *models.Error
}

func (o *PaymentsForRecurringPaymentNotImplemented) Error() string {
	return fmt.Sprintf("[GET /recurring-payments/{recurringPaymentId}/payments][%d] paymentsForRecurringPaymentNotImplemented  %+v", 501, o.Payload)
}
func (o *PaymentsForRecurringPaymentNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *PaymentsForRecurringPaymentNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsForRecurringPaymentServiceUnavailable creates a PaymentsForRecurringPaymentServiceUnavailable with default headers values
func NewPaymentsForRecurringPaymentServiceUnavailable() *PaymentsForRecurringPaymentServiceUnavailable {
	return &PaymentsForRecurringPaymentServiceUnavailable{}
}

/* PaymentsForRecurringPaymentServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type PaymentsForRecurringPaymentServiceUnavailable struct {
	Payload *models.Error
}

func (o *PaymentsForRecurringPaymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /recurring-payments/{recurringPaymentId}/payments][%d] paymentsForRecurringPaymentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PaymentsForRecurringPaymentServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PaymentsForRecurringPaymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

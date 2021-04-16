// Code generated by go-swagger; DO NOT EDIT.

package scheduled_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/openbanking/accountinformation/models"
)

// GetAccountsAccountIDScheduledPaymentsReader is a Reader for the GetAccountsAccountIDScheduledPayments structure.
type GetAccountsAccountIDScheduledPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDScheduledPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDScheduledPaymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDScheduledPaymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDScheduledPaymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDScheduledPaymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDScheduledPaymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDScheduledPaymentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDScheduledPaymentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDScheduledPaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDScheduledPaymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDScheduledPaymentsOK creates a GetAccountsAccountIDScheduledPaymentsOK with default headers values
func NewGetAccountsAccountIDScheduledPaymentsOK() *GetAccountsAccountIDScheduledPaymentsOK {
	return &GetAccountsAccountIDScheduledPaymentsOK{}
}

/* GetAccountsAccountIDScheduledPaymentsOK describes a response with status code 200, with default header values.

Scheduled Payments Read
*/
type GetAccountsAccountIDScheduledPaymentsOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadScheduledPayment3
}

func (o *GetAccountsAccountIDScheduledPaymentsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDScheduledPaymentsOK) GetPayload() *models.OBReadScheduledPayment3 {
	return o.Payload
}

func (o *GetAccountsAccountIDScheduledPaymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadScheduledPayment3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsBadRequest creates a GetAccountsAccountIDScheduledPaymentsBadRequest with default headers values
func NewGetAccountsAccountIDScheduledPaymentsBadRequest() *GetAccountsAccountIDScheduledPaymentsBadRequest {
	return &GetAccountsAccountIDScheduledPaymentsBadRequest{}
}

/* GetAccountsAccountIDScheduledPaymentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDScheduledPaymentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDScheduledPaymentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDScheduledPaymentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDScheduledPaymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsUnauthorized creates a GetAccountsAccountIDScheduledPaymentsUnauthorized with default headers values
func NewGetAccountsAccountIDScheduledPaymentsUnauthorized() *GetAccountsAccountIDScheduledPaymentsUnauthorized {
	return &GetAccountsAccountIDScheduledPaymentsUnauthorized{}
}

/* GetAccountsAccountIDScheduledPaymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDScheduledPaymentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDScheduledPaymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsUnauthorized ", 401)
}

func (o *GetAccountsAccountIDScheduledPaymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsForbidden creates a GetAccountsAccountIDScheduledPaymentsForbidden with default headers values
func NewGetAccountsAccountIDScheduledPaymentsForbidden() *GetAccountsAccountIDScheduledPaymentsForbidden {
	return &GetAccountsAccountIDScheduledPaymentsForbidden{}
}

/* GetAccountsAccountIDScheduledPaymentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDScheduledPaymentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDScheduledPaymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDScheduledPaymentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDScheduledPaymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsNotFound creates a GetAccountsAccountIDScheduledPaymentsNotFound with default headers values
func NewGetAccountsAccountIDScheduledPaymentsNotFound() *GetAccountsAccountIDScheduledPaymentsNotFound {
	return &GetAccountsAccountIDScheduledPaymentsNotFound{}
}

/* GetAccountsAccountIDScheduledPaymentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAccountsAccountIDScheduledPaymentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDScheduledPaymentsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsNotFound ", 404)
}

func (o *GetAccountsAccountIDScheduledPaymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsMethodNotAllowed creates a GetAccountsAccountIDScheduledPaymentsMethodNotAllowed with default headers values
func NewGetAccountsAccountIDScheduledPaymentsMethodNotAllowed() *GetAccountsAccountIDScheduledPaymentsMethodNotAllowed {
	return &GetAccountsAccountIDScheduledPaymentsMethodNotAllowed{}
}

/* GetAccountsAccountIDScheduledPaymentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDScheduledPaymentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDScheduledPaymentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDScheduledPaymentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsNotAcceptable creates a GetAccountsAccountIDScheduledPaymentsNotAcceptable with default headers values
func NewGetAccountsAccountIDScheduledPaymentsNotAcceptable() *GetAccountsAccountIDScheduledPaymentsNotAcceptable {
	return &GetAccountsAccountIDScheduledPaymentsNotAcceptable{}
}

/* GetAccountsAccountIDScheduledPaymentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDScheduledPaymentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDScheduledPaymentsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDScheduledPaymentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsTooManyRequests creates a GetAccountsAccountIDScheduledPaymentsTooManyRequests with default headers values
func NewGetAccountsAccountIDScheduledPaymentsTooManyRequests() *GetAccountsAccountIDScheduledPaymentsTooManyRequests {
	return &GetAccountsAccountIDScheduledPaymentsTooManyRequests{}
}

/* GetAccountsAccountIDScheduledPaymentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDScheduledPaymentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDScheduledPaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDScheduledPaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertInt64(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "int64", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDScheduledPaymentsInternalServerError creates a GetAccountsAccountIDScheduledPaymentsInternalServerError with default headers values
func NewGetAccountsAccountIDScheduledPaymentsInternalServerError() *GetAccountsAccountIDScheduledPaymentsInternalServerError {
	return &GetAccountsAccountIDScheduledPaymentsInternalServerError{}
}

/* GetAccountsAccountIDScheduledPaymentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDScheduledPaymentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDScheduledPaymentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/scheduled-payments][%d] getAccountsAccountIdScheduledPaymentsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDScheduledPaymentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDScheduledPaymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

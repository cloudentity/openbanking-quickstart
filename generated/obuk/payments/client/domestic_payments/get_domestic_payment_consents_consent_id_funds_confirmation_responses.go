// Code generated by go-swagger; DO NOT EDIT.

package domestic_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/models"
)

// GetDomesticPaymentConsentsConsentIDFundsConfirmationReader is a Reader for the GetDomesticPaymentConsentsConsentIDFundsConfirmation structure.
type GetDomesticPaymentConsentsConsentIDFundsConfirmationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationOK creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationOK with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationOK() *GetDomesticPaymentConsentsConsentIDFundsConfirmationOK {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationOK{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationOK describes a response with status code 200, with default header values.

Domestic Payment Consents Read
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteFundsConfirmationResponse1
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationOK) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationOK  %+v", 200, o.Payload)
}
func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationOK) GetPayload() *models.OBWriteFundsConfirmationResponse1 {
	return o.Payload
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBWriteFundsConfirmationResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest() *GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationBadRequest  %+v", 400, o.Payload)
}
func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized() *GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationUnauthorized ", 401)
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden() *GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationForbidden  %+v", 403, o.Payload)
}
func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed() *GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationMethodNotAllowed ", 405)
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable() *GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationNotAcceptable ", 406)
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests() *GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationTooManyRequests ", 429)
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError creates a GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError with default headers values
func NewGetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError() *GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError {
	return &GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError{}
}

/* GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /domestic-payment-consents/{ConsentId}/funds-confirmation][%d] getDomesticPaymentConsentsConsentIdFundsConfirmationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticPaymentConsentsConsentIDFundsConfirmationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// CreateDomesticPaymentsReader is a Reader for the CreateDomesticPayments structure.
type CreateDomesticPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDomesticPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDomesticPaymentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDomesticPaymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDomesticPaymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDomesticPaymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateDomesticPaymentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateDomesticPaymentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDomesticPaymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDomesticPaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDomesticPaymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDomesticPaymentsCreated creates a CreateDomesticPaymentsCreated with default headers values
func NewCreateDomesticPaymentsCreated() *CreateDomesticPaymentsCreated {
	return &CreateDomesticPaymentsCreated{}
}

/* CreateDomesticPaymentsCreated describes a response with status code 201, with default header values.

Domestic Payments Created
*/
type CreateDomesticPaymentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteDomesticResponse5
}

func (o *CreateDomesticPaymentsCreated) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsCreated  %+v", 201, o.Payload)
}
func (o *CreateDomesticPaymentsCreated) GetPayload() *models.OBWriteDomesticResponse5 {
	return o.Payload
}

func (o *CreateDomesticPaymentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteDomesticResponse5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDomesticPaymentsBadRequest creates a CreateDomesticPaymentsBadRequest with default headers values
func NewCreateDomesticPaymentsBadRequest() *CreateDomesticPaymentsBadRequest {
	return &CreateDomesticPaymentsBadRequest{}
}

/* CreateDomesticPaymentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDomesticPaymentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticPaymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDomesticPaymentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticPaymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticPaymentsUnauthorized creates a CreateDomesticPaymentsUnauthorized with default headers values
func NewCreateDomesticPaymentsUnauthorized() *CreateDomesticPaymentsUnauthorized {
	return &CreateDomesticPaymentsUnauthorized{}
}

/* CreateDomesticPaymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDomesticPaymentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticPaymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsUnauthorized ", 401)
}

func (o *CreateDomesticPaymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticPaymentsForbidden creates a CreateDomesticPaymentsForbidden with default headers values
func NewCreateDomesticPaymentsForbidden() *CreateDomesticPaymentsForbidden {
	return &CreateDomesticPaymentsForbidden{}
}

/* CreateDomesticPaymentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDomesticPaymentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticPaymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateDomesticPaymentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticPaymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticPaymentsMethodNotAllowed creates a CreateDomesticPaymentsMethodNotAllowed with default headers values
func NewCreateDomesticPaymentsMethodNotAllowed() *CreateDomesticPaymentsMethodNotAllowed {
	return &CreateDomesticPaymentsMethodNotAllowed{}
}

/* CreateDomesticPaymentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateDomesticPaymentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticPaymentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsMethodNotAllowed ", 405)
}

func (o *CreateDomesticPaymentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticPaymentsNotAcceptable creates a CreateDomesticPaymentsNotAcceptable with default headers values
func NewCreateDomesticPaymentsNotAcceptable() *CreateDomesticPaymentsNotAcceptable {
	return &CreateDomesticPaymentsNotAcceptable{}
}

/* CreateDomesticPaymentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateDomesticPaymentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticPaymentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsNotAcceptable ", 406)
}

func (o *CreateDomesticPaymentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticPaymentsUnsupportedMediaType creates a CreateDomesticPaymentsUnsupportedMediaType with default headers values
func NewCreateDomesticPaymentsUnsupportedMediaType() *CreateDomesticPaymentsUnsupportedMediaType {
	return &CreateDomesticPaymentsUnsupportedMediaType{}
}

/* CreateDomesticPaymentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateDomesticPaymentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticPaymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsUnsupportedMediaType ", 415)
}

func (o *CreateDomesticPaymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateDomesticPaymentsTooManyRequests creates a CreateDomesticPaymentsTooManyRequests with default headers values
func NewCreateDomesticPaymentsTooManyRequests() *CreateDomesticPaymentsTooManyRequests {
	return &CreateDomesticPaymentsTooManyRequests{}
}

/* CreateDomesticPaymentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateDomesticPaymentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateDomesticPaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsTooManyRequests ", 429)
}

func (o *CreateDomesticPaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDomesticPaymentsInternalServerError creates a CreateDomesticPaymentsInternalServerError with default headers values
func NewCreateDomesticPaymentsInternalServerError() *CreateDomesticPaymentsInternalServerError {
	return &CreateDomesticPaymentsInternalServerError{}
}

/* CreateDomesticPaymentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateDomesticPaymentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateDomesticPaymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /domestic-payments][%d] createDomesticPaymentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDomesticPaymentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateDomesticPaymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package domestic_standing_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/openbanking/paymentinitiation/models"
)

// GetDomesticStandingOrderConsentsConsentIDReader is a Reader for the GetDomesticStandingOrderConsentsConsentID structure.
type GetDomesticStandingOrderConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomesticStandingOrderConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomesticStandingOrderConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomesticStandingOrderConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomesticStandingOrderConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDomesticStandingOrderConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDomesticStandingOrderConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetDomesticStandingOrderConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetDomesticStandingOrderConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDomesticStandingOrderConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDomesticStandingOrderConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomesticStandingOrderConsentsConsentIDOK creates a GetDomesticStandingOrderConsentsConsentIDOK with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDOK() *GetDomesticStandingOrderConsentsConsentIDOK {
	return &GetDomesticStandingOrderConsentsConsentIDOK{}
}

/* GetDomesticStandingOrderConsentsConsentIDOK describes a response with status code 200, with default header values.

Domestic Standing Order Consents Read
*/
type GetDomesticStandingOrderConsentsConsentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteDomesticStandingOrderConsentResponse6
}

func (o *GetDomesticStandingOrderConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *GetDomesticStandingOrderConsentsConsentIDOK) GetPayload() *models.OBWriteDomesticStandingOrderConsentResponse6 {
	return o.Payload
}

func (o *GetDomesticStandingOrderConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteDomesticStandingOrderConsentResponse6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomesticStandingOrderConsentsConsentIDBadRequest creates a GetDomesticStandingOrderConsentsConsentIDBadRequest with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDBadRequest() *GetDomesticStandingOrderConsentsConsentIDBadRequest {
	return &GetDomesticStandingOrderConsentsConsentIDBadRequest{}
}

/* GetDomesticStandingOrderConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDomesticStandingOrderConsentsConsentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticStandingOrderConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetDomesticStandingOrderConsentsConsentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticStandingOrderConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticStandingOrderConsentsConsentIDUnauthorized creates a GetDomesticStandingOrderConsentsConsentIDUnauthorized with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDUnauthorized() *GetDomesticStandingOrderConsentsConsentIDUnauthorized {
	return &GetDomesticStandingOrderConsentsConsentIDUnauthorized{}
}

/* GetDomesticStandingOrderConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDomesticStandingOrderConsentsConsentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticStandingOrderConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdUnauthorized ", 401)
}

func (o *GetDomesticStandingOrderConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticStandingOrderConsentsConsentIDForbidden creates a GetDomesticStandingOrderConsentsConsentIDForbidden with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDForbidden() *GetDomesticStandingOrderConsentsConsentIDForbidden {
	return &GetDomesticStandingOrderConsentsConsentIDForbidden{}
}

/* GetDomesticStandingOrderConsentsConsentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDomesticStandingOrderConsentsConsentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticStandingOrderConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetDomesticStandingOrderConsentsConsentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticStandingOrderConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticStandingOrderConsentsConsentIDNotFound creates a GetDomesticStandingOrderConsentsConsentIDNotFound with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDNotFound() *GetDomesticStandingOrderConsentsConsentIDNotFound {
	return &GetDomesticStandingOrderConsentsConsentIDNotFound{}
}

/* GetDomesticStandingOrderConsentsConsentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDomesticStandingOrderConsentsConsentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticStandingOrderConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdNotFound ", 404)
}

func (o *GetDomesticStandingOrderConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticStandingOrderConsentsConsentIDMethodNotAllowed creates a GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDMethodNotAllowed() *GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed {
	return &GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed{}
}

/* GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdMethodNotAllowed ", 405)
}

func (o *GetDomesticStandingOrderConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticStandingOrderConsentsConsentIDNotAcceptable creates a GetDomesticStandingOrderConsentsConsentIDNotAcceptable with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDNotAcceptable() *GetDomesticStandingOrderConsentsConsentIDNotAcceptable {
	return &GetDomesticStandingOrderConsentsConsentIDNotAcceptable{}
}

/* GetDomesticStandingOrderConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetDomesticStandingOrderConsentsConsentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticStandingOrderConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdNotAcceptable ", 406)
}

func (o *GetDomesticStandingOrderConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetDomesticStandingOrderConsentsConsentIDTooManyRequests creates a GetDomesticStandingOrderConsentsConsentIDTooManyRequests with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDTooManyRequests() *GetDomesticStandingOrderConsentsConsentIDTooManyRequests {
	return &GetDomesticStandingOrderConsentsConsentIDTooManyRequests{}
}

/* GetDomesticStandingOrderConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDomesticStandingOrderConsentsConsentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetDomesticStandingOrderConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdTooManyRequests ", 429)
}

func (o *GetDomesticStandingOrderConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDomesticStandingOrderConsentsConsentIDInternalServerError creates a GetDomesticStandingOrderConsentsConsentIDInternalServerError with default headers values
func NewGetDomesticStandingOrderConsentsConsentIDInternalServerError() *GetDomesticStandingOrderConsentsConsentIDInternalServerError {
	return &GetDomesticStandingOrderConsentsConsentIDInternalServerError{}
}

/* GetDomesticStandingOrderConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDomesticStandingOrderConsentsConsentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetDomesticStandingOrderConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /domestic-standing-order-consents/{ConsentId}][%d] getDomesticStandingOrderConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDomesticStandingOrderConsentsConsentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetDomesticStandingOrderConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

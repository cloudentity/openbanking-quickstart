// Code generated by go-swagger; DO NOT EDIT.

package file_payments

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

// GetFilePaymentsFilePaymentIDReader is a Reader for the GetFilePaymentsFilePaymentID structure.
type GetFilePaymentsFilePaymentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentsFilePaymentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentsFilePaymentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilePaymentsFilePaymentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilePaymentsFilePaymentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentsFilePaymentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilePaymentsFilePaymentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFilePaymentsFilePaymentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFilePaymentsFilePaymentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentsFilePaymentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFilePaymentsFilePaymentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentsFilePaymentIDOK creates a GetFilePaymentsFilePaymentIDOK with default headers values
func NewGetFilePaymentsFilePaymentIDOK() *GetFilePaymentsFilePaymentIDOK {
	return &GetFilePaymentsFilePaymentIDOK{}
}

/* GetFilePaymentsFilePaymentIDOK describes a response with status code 200, with default header values.

File Payments Read
*/
type GetFilePaymentsFilePaymentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteFileResponse3
}

func (o *GetFilePaymentsFilePaymentIDOK) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDOK) GetPayload() *models.OBWriteFileResponse3 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWriteFileResponse3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDBadRequest creates a GetFilePaymentsFilePaymentIDBadRequest with default headers values
func NewGetFilePaymentsFilePaymentIDBadRequest() *GetFilePaymentsFilePaymentIDBadRequest {
	return &GetFilePaymentsFilePaymentIDBadRequest{}
}

/* GetFilePaymentsFilePaymentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFilePaymentsFilePaymentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDUnauthorized creates a GetFilePaymentsFilePaymentIDUnauthorized with default headers values
func NewGetFilePaymentsFilePaymentIDUnauthorized() *GetFilePaymentsFilePaymentIDUnauthorized {
	return &GetFilePaymentsFilePaymentIDUnauthorized{}
}

/* GetFilePaymentsFilePaymentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFilePaymentsFilePaymentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdUnauthorized ", 401)
}

func (o *GetFilePaymentsFilePaymentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDForbidden creates a GetFilePaymentsFilePaymentIDForbidden with default headers values
func NewGetFilePaymentsFilePaymentIDForbidden() *GetFilePaymentsFilePaymentIDForbidden {
	return &GetFilePaymentsFilePaymentIDForbidden{}
}

/* GetFilePaymentsFilePaymentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFilePaymentsFilePaymentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDNotFound creates a GetFilePaymentsFilePaymentIDNotFound with default headers values
func NewGetFilePaymentsFilePaymentIDNotFound() *GetFilePaymentsFilePaymentIDNotFound {
	return &GetFilePaymentsFilePaymentIDNotFound{}
}

/* GetFilePaymentsFilePaymentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetFilePaymentsFilePaymentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdNotFound ", 404)
}

func (o *GetFilePaymentsFilePaymentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDMethodNotAllowed creates a GetFilePaymentsFilePaymentIDMethodNotAllowed with default headers values
func NewGetFilePaymentsFilePaymentIDMethodNotAllowed() *GetFilePaymentsFilePaymentIDMethodNotAllowed {
	return &GetFilePaymentsFilePaymentIDMethodNotAllowed{}
}

/* GetFilePaymentsFilePaymentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFilePaymentsFilePaymentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdMethodNotAllowed ", 405)
}

func (o *GetFilePaymentsFilePaymentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDNotAcceptable creates a GetFilePaymentsFilePaymentIDNotAcceptable with default headers values
func NewGetFilePaymentsFilePaymentIDNotAcceptable() *GetFilePaymentsFilePaymentIDNotAcceptable {
	return &GetFilePaymentsFilePaymentIDNotAcceptable{}
}

/* GetFilePaymentsFilePaymentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetFilePaymentsFilePaymentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdNotAcceptable ", 406)
}

func (o *GetFilePaymentsFilePaymentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDTooManyRequests creates a GetFilePaymentsFilePaymentIDTooManyRequests with default headers values
func NewGetFilePaymentsFilePaymentIDTooManyRequests() *GetFilePaymentsFilePaymentIDTooManyRequests {
	return &GetFilePaymentsFilePaymentIDTooManyRequests{}
}

/* GetFilePaymentsFilePaymentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFilePaymentsFilePaymentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdTooManyRequests ", 429)
}

func (o *GetFilePaymentsFilePaymentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDInternalServerError creates a GetFilePaymentsFilePaymentIDInternalServerError with default headers values
func NewGetFilePaymentsFilePaymentIDInternalServerError() *GetFilePaymentsFilePaymentIDInternalServerError {
	return &GetFilePaymentsFilePaymentIDInternalServerError{}
}

/* GetFilePaymentsFilePaymentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFilePaymentsFilePaymentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}][%d] getFilePaymentsFilePaymentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

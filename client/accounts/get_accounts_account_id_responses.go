// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/models"
)

// GetAccountsAccountIDReader is a Reader for the GetAccountsAccountID structure.
type GetAccountsAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDOK creates a GetAccountsAccountIDOK with default headers values
func NewGetAccountsAccountIDOK() *GetAccountsAccountIDOK {
	return &GetAccountsAccountIDOK{}
}

/* GetAccountsAccountIDOK describes a response with status code 200, with default header values.

Accounts Read
*/
type GetAccountsAccountIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadAccount6
}

func (o *GetAccountsAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDOK) GetPayload() *models.OBReadAccount6 {
	return o.Payload
}

func (o *GetAccountsAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadAccount6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDBadRequest creates a GetAccountsAccountIDBadRequest with default headers values
func NewGetAccountsAccountIDBadRequest() *GetAccountsAccountIDBadRequest {
	return &GetAccountsAccountIDBadRequest{}
}

/* GetAccountsAccountIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDUnauthorized creates a GetAccountsAccountIDUnauthorized with default headers values
func NewGetAccountsAccountIDUnauthorized() *GetAccountsAccountIDUnauthorized {
	return &GetAccountsAccountIDUnauthorized{}
}

/* GetAccountsAccountIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdUnauthorized ", 401)
}

func (o *GetAccountsAccountIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDForbidden creates a GetAccountsAccountIDForbidden with default headers values
func NewGetAccountsAccountIDForbidden() *GetAccountsAccountIDForbidden {
	return &GetAccountsAccountIDForbidden{}
}

/* GetAccountsAccountIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDMethodNotAllowed creates a GetAccountsAccountIDMethodNotAllowed with default headers values
func NewGetAccountsAccountIDMethodNotAllowed() *GetAccountsAccountIDMethodNotAllowed {
	return &GetAccountsAccountIDMethodNotAllowed{}
}

/* GetAccountsAccountIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDNotAcceptable creates a GetAccountsAccountIDNotAcceptable with default headers values
func NewGetAccountsAccountIDNotAcceptable() *GetAccountsAccountIDNotAcceptable {
	return &GetAccountsAccountIDNotAcceptable{}
}

/* GetAccountsAccountIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDTooManyRequests creates a GetAccountsAccountIDTooManyRequests with default headers values
func NewGetAccountsAccountIDTooManyRequests() *GetAccountsAccountIDTooManyRequests {
	return &GetAccountsAccountIDTooManyRequests{}
}

/* GetAccountsAccountIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDInternalServerError creates a GetAccountsAccountIDInternalServerError with default headers values
func NewGetAccountsAccountIDInternalServerError() *GetAccountsAccountIDInternalServerError {
	return &GetAccountsAccountIDInternalServerError{}
}

/* GetAccountsAccountIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}][%d] getAccountsAccountIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

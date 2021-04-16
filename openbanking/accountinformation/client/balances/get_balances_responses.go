// Code generated by go-swagger; DO NOT EDIT.

package balances

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

// GetBalancesReader is a Reader for the GetBalances structure.
type GetBalancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBalancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBalancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBalancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBalancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBalancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBalancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetBalancesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetBalancesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBalancesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBalancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBalancesOK creates a GetBalancesOK with default headers values
func NewGetBalancesOK() *GetBalancesOK {
	return &GetBalancesOK{}
}

/* GetBalancesOK describes a response with status code 200, with default header values.

Balances Read
*/
type GetBalancesOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadBalance1
}

func (o *GetBalancesOK) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesOK  %+v", 200, o.Payload)
}
func (o *GetBalancesOK) GetPayload() *models.OBReadBalance1 {
	return o.Payload
}

func (o *GetBalancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadBalance1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBalancesBadRequest creates a GetBalancesBadRequest with default headers values
func NewGetBalancesBadRequest() *GetBalancesBadRequest {
	return &GetBalancesBadRequest{}
}

/* GetBalancesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBalancesBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBalancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesBadRequest  %+v", 400, o.Payload)
}
func (o *GetBalancesBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBalancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBalancesUnauthorized creates a GetBalancesUnauthorized with default headers values
func NewGetBalancesUnauthorized() *GetBalancesUnauthorized {
	return &GetBalancesUnauthorized{}
}

/* GetBalancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBalancesUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBalancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesUnauthorized ", 401)
}

func (o *GetBalancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetBalancesForbidden creates a GetBalancesForbidden with default headers values
func NewGetBalancesForbidden() *GetBalancesForbidden {
	return &GetBalancesForbidden{}
}

/* GetBalancesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBalancesForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBalancesForbidden) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesForbidden  %+v", 403, o.Payload)
}
func (o *GetBalancesForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBalancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBalancesNotFound creates a GetBalancesNotFound with default headers values
func NewGetBalancesNotFound() *GetBalancesNotFound {
	return &GetBalancesNotFound{}
}

/* GetBalancesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBalancesNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBalancesNotFound) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesNotFound ", 404)
}

func (o *GetBalancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetBalancesMethodNotAllowed creates a GetBalancesMethodNotAllowed with default headers values
func NewGetBalancesMethodNotAllowed() *GetBalancesMethodNotAllowed {
	return &GetBalancesMethodNotAllowed{}
}

/* GetBalancesMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetBalancesMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBalancesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesMethodNotAllowed ", 405)
}

func (o *GetBalancesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetBalancesNotAcceptable creates a GetBalancesNotAcceptable with default headers values
func NewGetBalancesNotAcceptable() *GetBalancesNotAcceptable {
	return &GetBalancesNotAcceptable{}
}

/* GetBalancesNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetBalancesNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBalancesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesNotAcceptable ", 406)
}

func (o *GetBalancesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetBalancesTooManyRequests creates a GetBalancesTooManyRequests with default headers values
func NewGetBalancesTooManyRequests() *GetBalancesTooManyRequests {
	return &GetBalancesTooManyRequests{}
}

/* GetBalancesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetBalancesTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBalancesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesTooManyRequests ", 429)
}

func (o *GetBalancesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBalancesInternalServerError creates a GetBalancesInternalServerError with default headers values
func NewGetBalancesInternalServerError() *GetBalancesInternalServerError {
	return &GetBalancesInternalServerError{}
}

/* GetBalancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBalancesInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBalancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /balances][%d] getBalancesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBalancesInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBalancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

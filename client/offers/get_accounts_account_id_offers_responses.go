// Code generated by go-swagger; DO NOT EDIT.

package offers

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

// GetAccountsAccountIDOffersReader is a Reader for the GetAccountsAccountIDOffers structure.
type GetAccountsAccountIDOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDOffersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDOffersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDOffersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDOffersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDOffersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDOffersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDOffersOK creates a GetAccountsAccountIDOffersOK with default headers values
func NewGetAccountsAccountIDOffersOK() *GetAccountsAccountIDOffersOK {
	return &GetAccountsAccountIDOffersOK{}
}

/* GetAccountsAccountIDOffersOK describes a response with status code 200, with default header values.

Offers Read
*/
type GetAccountsAccountIDOffersOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadOffer1
}

func (o *GetAccountsAccountIDOffersOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDOffersOK) GetPayload() *models.OBReadOffer1 {
	return o.Payload
}

func (o *GetAccountsAccountIDOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadOffer1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDOffersBadRequest creates a GetAccountsAccountIDOffersBadRequest with default headers values
func NewGetAccountsAccountIDOffersBadRequest() *GetAccountsAccountIDOffersBadRequest {
	return &GetAccountsAccountIDOffersBadRequest{}
}

/* GetAccountsAccountIDOffersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDOffersBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDOffersBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDOffersBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDOffersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDOffersUnauthorized creates a GetAccountsAccountIDOffersUnauthorized with default headers values
func NewGetAccountsAccountIDOffersUnauthorized() *GetAccountsAccountIDOffersUnauthorized {
	return &GetAccountsAccountIDOffersUnauthorized{}
}

/* GetAccountsAccountIDOffersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDOffersUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDOffersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersUnauthorized ", 401)
}

func (o *GetAccountsAccountIDOffersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDOffersForbidden creates a GetAccountsAccountIDOffersForbidden with default headers values
func NewGetAccountsAccountIDOffersForbidden() *GetAccountsAccountIDOffersForbidden {
	return &GetAccountsAccountIDOffersForbidden{}
}

/* GetAccountsAccountIDOffersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDOffersForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDOffersForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDOffersForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDOffersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDOffersNotFound creates a GetAccountsAccountIDOffersNotFound with default headers values
func NewGetAccountsAccountIDOffersNotFound() *GetAccountsAccountIDOffersNotFound {
	return &GetAccountsAccountIDOffersNotFound{}
}

/* GetAccountsAccountIDOffersNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAccountsAccountIDOffersNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDOffersNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersNotFound ", 404)
}

func (o *GetAccountsAccountIDOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDOffersMethodNotAllowed creates a GetAccountsAccountIDOffersMethodNotAllowed with default headers values
func NewGetAccountsAccountIDOffersMethodNotAllowed() *GetAccountsAccountIDOffersMethodNotAllowed {
	return &GetAccountsAccountIDOffersMethodNotAllowed{}
}

/* GetAccountsAccountIDOffersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDOffersMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDOffersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDOffersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDOffersNotAcceptable creates a GetAccountsAccountIDOffersNotAcceptable with default headers values
func NewGetAccountsAccountIDOffersNotAcceptable() *GetAccountsAccountIDOffersNotAcceptable {
	return &GetAccountsAccountIDOffersNotAcceptable{}
}

/* GetAccountsAccountIDOffersNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDOffersNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDOffersNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDOffersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDOffersTooManyRequests creates a GetAccountsAccountIDOffersTooManyRequests with default headers values
func NewGetAccountsAccountIDOffersTooManyRequests() *GetAccountsAccountIDOffersTooManyRequests {
	return &GetAccountsAccountIDOffersTooManyRequests{}
}

/* GetAccountsAccountIDOffersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDOffersTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDOffersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDOffersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDOffersInternalServerError creates a GetAccountsAccountIDOffersInternalServerError with default headers values
func NewGetAccountsAccountIDOffersInternalServerError() *GetAccountsAccountIDOffersInternalServerError {
	return &GetAccountsAccountIDOffersInternalServerError{}
}

/* GetAccountsAccountIDOffersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDOffersInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/offers][%d] getAccountsAccountIdOffersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDOffersInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

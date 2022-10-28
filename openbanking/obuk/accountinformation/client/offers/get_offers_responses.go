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

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
)

// GetOffersReader is a Reader for the GetOffers structure.
type GetOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOffersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOffersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOffersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetOffersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetOffersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOffersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOffersOK creates a GetOffersOK with default headers values
func NewGetOffersOK() *GetOffersOK {
	return &GetOffersOK{}
}

/* GetOffersOK describes a response with status code 200, with default header values.

Offers Read
*/
type GetOffersOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadOffer1
}

func (o *GetOffersOK) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersOK  %+v", 200, o.Payload)
}
func (o *GetOffersOK) GetPayload() *models.OBReadOffer1 {
	return o.Payload
}

func (o *GetOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOffersBadRequest creates a GetOffersBadRequest with default headers values
func NewGetOffersBadRequest() *GetOffersBadRequest {
	return &GetOffersBadRequest{}
}

/* GetOffersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetOffersBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetOffersBadRequest) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersBadRequest  %+v", 400, o.Payload)
}
func (o *GetOffersBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetOffersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOffersUnauthorized creates a GetOffersUnauthorized with default headers values
func NewGetOffersUnauthorized() *GetOffersUnauthorized {
	return &GetOffersUnauthorized{}
}

/* GetOffersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOffersUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetOffersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersUnauthorized ", 401)
}

func (o *GetOffersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetOffersForbidden creates a GetOffersForbidden with default headers values
func NewGetOffersForbidden() *GetOffersForbidden {
	return &GetOffersForbidden{}
}

/* GetOffersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOffersForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetOffersForbidden) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersForbidden  %+v", 403, o.Payload)
}
func (o *GetOffersForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetOffersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOffersNotFound creates a GetOffersNotFound with default headers values
func NewGetOffersNotFound() *GetOffersNotFound {
	return &GetOffersNotFound{}
}

/* GetOffersNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetOffersNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetOffersNotFound) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersNotFound ", 404)
}

func (o *GetOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetOffersMethodNotAllowed creates a GetOffersMethodNotAllowed with default headers values
func NewGetOffersMethodNotAllowed() *GetOffersMethodNotAllowed {
	return &GetOffersMethodNotAllowed{}
}

/* GetOffersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetOffersMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetOffersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersMethodNotAllowed ", 405)
}

func (o *GetOffersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetOffersNotAcceptable creates a GetOffersNotAcceptable with default headers values
func NewGetOffersNotAcceptable() *GetOffersNotAcceptable {
	return &GetOffersNotAcceptable{}
}

/* GetOffersNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetOffersNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetOffersNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersNotAcceptable ", 406)
}

func (o *GetOffersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetOffersTooManyRequests creates a GetOffersTooManyRequests with default headers values
func NewGetOffersTooManyRequests() *GetOffersTooManyRequests {
	return &GetOffersTooManyRequests{}
}

/* GetOffersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetOffersTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetOffersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersTooManyRequests ", 429)
}

func (o *GetOffersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOffersInternalServerError creates a GetOffersInternalServerError with default headers values
func NewGetOffersInternalServerError() *GetOffersInternalServerError {
	return &GetOffersInternalServerError{}
}

/* GetOffersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOffersInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /offers][%d] getOffersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOffersInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
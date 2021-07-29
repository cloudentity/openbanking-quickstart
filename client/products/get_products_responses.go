// Code generated by go-swagger; DO NOT EDIT.

package products

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

// GetProductsReader is a Reader for the GetProducts structure.
type GetProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProductsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProductsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProductsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProductsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetProductsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetProductsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetProductsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProductsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProductsOK creates a GetProductsOK with default headers values
func NewGetProductsOK() *GetProductsOK {
	return &GetProductsOK{}
}

/* GetProductsOK describes a response with status code 200, with default header values.

Products Read
*/
type GetProductsOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadProduct2
}

func (o *GetProductsOK) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsOK  %+v", 200, o.Payload)
}
func (o *GetProductsOK) GetPayload() *models.OBReadProduct2 {
	return o.Payload
}

func (o *GetProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadProduct2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsBadRequest creates a GetProductsBadRequest with default headers values
func NewGetProductsBadRequest() *GetProductsBadRequest {
	return &GetProductsBadRequest{}
}

/* GetProductsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetProductsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetProductsBadRequest) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsBadRequest  %+v", 400, o.Payload)
}
func (o *GetProductsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetProductsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProductsUnauthorized creates a GetProductsUnauthorized with default headers values
func NewGetProductsUnauthorized() *GetProductsUnauthorized {
	return &GetProductsUnauthorized{}
}

/* GetProductsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProductsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetProductsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsUnauthorized ", 401)
}

func (o *GetProductsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetProductsForbidden creates a GetProductsForbidden with default headers values
func NewGetProductsForbidden() *GetProductsForbidden {
	return &GetProductsForbidden{}
}

/* GetProductsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProductsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetProductsForbidden) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsForbidden  %+v", 403, o.Payload)
}
func (o *GetProductsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetProductsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProductsNotFound creates a GetProductsNotFound with default headers values
func NewGetProductsNotFound() *GetProductsNotFound {
	return &GetProductsNotFound{}
}

/* GetProductsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProductsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetProductsNotFound) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsNotFound ", 404)
}

func (o *GetProductsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetProductsMethodNotAllowed creates a GetProductsMethodNotAllowed with default headers values
func NewGetProductsMethodNotAllowed() *GetProductsMethodNotAllowed {
	return &GetProductsMethodNotAllowed{}
}

/* GetProductsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetProductsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetProductsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsMethodNotAllowed ", 405)
}

func (o *GetProductsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetProductsNotAcceptable creates a GetProductsNotAcceptable with default headers values
func NewGetProductsNotAcceptable() *GetProductsNotAcceptable {
	return &GetProductsNotAcceptable{}
}

/* GetProductsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetProductsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetProductsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsNotAcceptable ", 406)
}

func (o *GetProductsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetProductsTooManyRequests creates a GetProductsTooManyRequests with default headers values
func NewGetProductsTooManyRequests() *GetProductsTooManyRequests {
	return &GetProductsTooManyRequests{}
}

/* GetProductsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetProductsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetProductsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsTooManyRequests ", 429)
}

func (o *GetProductsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProductsInternalServerError creates a GetProductsInternalServerError with default headers values
func NewGetProductsInternalServerError() *GetProductsInternalServerError {
	return &GetProductsInternalServerError{}
}

/* GetProductsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProductsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetProductsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProductsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetProductsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

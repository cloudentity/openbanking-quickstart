// Code generated by go-swagger; DO NOT EDIT.

package beneficiaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/models"
)

// GetAccountsAccountIDBeneficiariesReader is a Reader for the GetAccountsAccountIDBeneficiaries structure.
type GetAccountsAccountIDBeneficiariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDBeneficiariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDBeneficiariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDBeneficiariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDBeneficiariesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDBeneficiariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDBeneficiariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDBeneficiariesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDBeneficiariesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDBeneficiariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDBeneficiariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDBeneficiariesOK creates a GetAccountsAccountIDBeneficiariesOK with default headers values
func NewGetAccountsAccountIDBeneficiariesOK() *GetAccountsAccountIDBeneficiariesOK {
	return &GetAccountsAccountIDBeneficiariesOK{}
}

/* GetAccountsAccountIDBeneficiariesOK describes a response with status code 200, with default header values.

Beneficiaries Read
*/
type GetAccountsAccountIDBeneficiariesOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadBeneficiary5
}

func (o *GetAccountsAccountIDBeneficiariesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDBeneficiariesOK) GetPayload() *models.OBReadBeneficiary5 {
	return o.Payload
}

func (o *GetAccountsAccountIDBeneficiariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadBeneficiary5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDBeneficiariesBadRequest creates a GetAccountsAccountIDBeneficiariesBadRequest with default headers values
func NewGetAccountsAccountIDBeneficiariesBadRequest() *GetAccountsAccountIDBeneficiariesBadRequest {
	return &GetAccountsAccountIDBeneficiariesBadRequest{}
}

/* GetAccountsAccountIDBeneficiariesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDBeneficiariesBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDBeneficiariesBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDBeneficiariesBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDBeneficiariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDBeneficiariesUnauthorized creates a GetAccountsAccountIDBeneficiariesUnauthorized with default headers values
func NewGetAccountsAccountIDBeneficiariesUnauthorized() *GetAccountsAccountIDBeneficiariesUnauthorized {
	return &GetAccountsAccountIDBeneficiariesUnauthorized{}
}

/* GetAccountsAccountIDBeneficiariesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDBeneficiariesUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDBeneficiariesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesUnauthorized ", 401)
}

func (o *GetAccountsAccountIDBeneficiariesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDBeneficiariesForbidden creates a GetAccountsAccountIDBeneficiariesForbidden with default headers values
func NewGetAccountsAccountIDBeneficiariesForbidden() *GetAccountsAccountIDBeneficiariesForbidden {
	return &GetAccountsAccountIDBeneficiariesForbidden{}
}

/* GetAccountsAccountIDBeneficiariesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDBeneficiariesForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDBeneficiariesForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDBeneficiariesForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDBeneficiariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDBeneficiariesNotFound creates a GetAccountsAccountIDBeneficiariesNotFound with default headers values
func NewGetAccountsAccountIDBeneficiariesNotFound() *GetAccountsAccountIDBeneficiariesNotFound {
	return &GetAccountsAccountIDBeneficiariesNotFound{}
}

/* GetAccountsAccountIDBeneficiariesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAccountsAccountIDBeneficiariesNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDBeneficiariesNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesNotFound ", 404)
}

func (o *GetAccountsAccountIDBeneficiariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDBeneficiariesMethodNotAllowed creates a GetAccountsAccountIDBeneficiariesMethodNotAllowed with default headers values
func NewGetAccountsAccountIDBeneficiariesMethodNotAllowed() *GetAccountsAccountIDBeneficiariesMethodNotAllowed {
	return &GetAccountsAccountIDBeneficiariesMethodNotAllowed{}
}

/* GetAccountsAccountIDBeneficiariesMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDBeneficiariesMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDBeneficiariesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDBeneficiariesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDBeneficiariesNotAcceptable creates a GetAccountsAccountIDBeneficiariesNotAcceptable with default headers values
func NewGetAccountsAccountIDBeneficiariesNotAcceptable() *GetAccountsAccountIDBeneficiariesNotAcceptable {
	return &GetAccountsAccountIDBeneficiariesNotAcceptable{}
}

/* GetAccountsAccountIDBeneficiariesNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDBeneficiariesNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDBeneficiariesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDBeneficiariesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDBeneficiariesTooManyRequests creates a GetAccountsAccountIDBeneficiariesTooManyRequests with default headers values
func NewGetAccountsAccountIDBeneficiariesTooManyRequests() *GetAccountsAccountIDBeneficiariesTooManyRequests {
	return &GetAccountsAccountIDBeneficiariesTooManyRequests{}
}

/* GetAccountsAccountIDBeneficiariesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDBeneficiariesTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDBeneficiariesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDBeneficiariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDBeneficiariesInternalServerError creates a GetAccountsAccountIDBeneficiariesInternalServerError with default headers values
func NewGetAccountsAccountIDBeneficiariesInternalServerError() *GetAccountsAccountIDBeneficiariesInternalServerError {
	return &GetAccountsAccountIDBeneficiariesInternalServerError{}
}

/* GetAccountsAccountIDBeneficiariesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDBeneficiariesInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDBeneficiariesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/beneficiaries][%d] getAccountsAccountIdBeneficiariesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDBeneficiariesInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDBeneficiariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

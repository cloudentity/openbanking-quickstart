// Code generated by go-swagger; DO NOT EDIT.

package banking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/cdr/models"
)

// ListDirectDebitsReader is a Reader for the ListDirectDebits structure.
type ListDirectDebitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDirectDebitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDirectDebitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListDirectDebitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListDirectDebitsOK creates a ListDirectDebitsOK with default headers values
func NewListDirectDebitsOK() *ListDirectDebitsOK {
	return &ListDirectDebitsOK{}
}

/* ListDirectDebitsOK describes a response with status code 200, with default header values.

Success
*/
type ListDirectDebitsOK struct {

	/* An [RFC4122](https://tools.ietf.org/html/rfc4122) UUID used as a correlation id. If provided, the data holder must play back this value in the x-fapi-interaction-id response header. If not provided a [RFC4122] UUID value is required to be provided in the response header to track the interaction.
	 */
	XFapiInteractionID string

	/* The [version](#response-headers) of the API end point that the data holder has responded with.
	 */
	Xv string

	Payload *models.ResponseBankingDirectDebitAuthorisationList
}

func (o *ListDirectDebitsOK) Error() string {
	return fmt.Sprintf("[GET /banking/accounts/{accountId}/direct-debits][%d] listDirectDebitsOK  %+v", 200, o.Payload)
}
func (o *ListDirectDebitsOK) GetPayload() *models.ResponseBankingDirectDebitAuthorisationList {
	return o.Payload
}

func (o *ListDirectDebitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-v
	hdrXv := response.GetHeader("x-v")

	if hdrXv != "" {
		o.Xv = hdrXv
	}

	o.Payload = new(models.ResponseBankingDirectDebitAuthorisationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDirectDebitsBadRequest creates a ListDirectDebitsBadRequest with default headers values
func NewListDirectDebitsBadRequest() *ListDirectDebitsBadRequest {
	return &ListDirectDebitsBadRequest{}
}

/* ListDirectDebitsBadRequest describes a response with status code 400, with default header values.

The following error codes MUST be supported:<br/><ul class="error-code-list"><li>[400 - Invalid Field](#error-400-field-invalid)</li><li>[400 - Invalid Page Size](#error-400-field-invalid-page-size)</li><li>[400 - Invalid Version](#error-400-header-invalid-version)</li><li>[404 - Unavailable Banking Account](#error-404-authorisation-unavailable-banking-account)</li><li>[404 - Invalid Banking Account](#error-404-authorisation-invalid-banking-account)</li><li>[406 - Unsupported Version](#error-406-header-unsupported-version)</li><li>[422 - Invalid Page](#error-422-field-invalid-page)</li></ul>
*/
type ListDirectDebitsBadRequest struct {

	/* An [RFC4122](https://tools.ietf.org/html/rfc4122) UUID used as a correlation id. If provided, the data holder must play back this value in the x-fapi-interaction-id response header. If not provided a [RFC4122] UUID value is required to be provided in the response header to track the interaction.
	 */
	XFapiInteractionID string

	Payload *models.ResponseErrorListV2
}

func (o *ListDirectDebitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /banking/accounts/{accountId}/direct-debits][%d] listDirectDebitsBadRequest  %+v", 400, o.Payload)
}
func (o *ListDirectDebitsBadRequest) GetPayload() *models.ResponseErrorListV2 {
	return o.Payload
}

func (o *ListDirectDebitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseErrorListV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

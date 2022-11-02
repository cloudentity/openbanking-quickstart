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

// ListScheduledPaymentsSpecificAccountsReader is a Reader for the ListScheduledPaymentsSpecificAccounts structure.
type ListScheduledPaymentsSpecificAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScheduledPaymentsSpecificAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScheduledPaymentsSpecificAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListScheduledPaymentsSpecificAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListScheduledPaymentsSpecificAccountsOK creates a ListScheduledPaymentsSpecificAccountsOK with default headers values
func NewListScheduledPaymentsSpecificAccountsOK() *ListScheduledPaymentsSpecificAccountsOK {
	return &ListScheduledPaymentsSpecificAccountsOK{}
}

/* ListScheduledPaymentsSpecificAccountsOK describes a response with status code 200, with default header values.

Success
*/
type ListScheduledPaymentsSpecificAccountsOK struct {

	/* An [RFC4122](https://tools.ietf.org/html/rfc4122) UUID used as a correlation id. If provided, the data holder must play back this value in the x-fapi-interaction-id response header. If not provided a [RFC4122] UUID value is required to be provided in the response header to track the interaction.
	 */
	XFapiInteractionID string

	/* The [version](#response-headers) of the API end point that the data holder has responded with.
	 */
	Xv string

	Payload *models.ResponseBankingScheduledPaymentsList
}

func (o *ListScheduledPaymentsSpecificAccountsOK) Error() string {
	return fmt.Sprintf("[POST /banking/payments/scheduled][%d] listScheduledPaymentsSpecificAccountsOK  %+v", 200, o.Payload)
}
func (o *ListScheduledPaymentsSpecificAccountsOK) GetPayload() *models.ResponseBankingScheduledPaymentsList {
	return o.Payload
}

func (o *ListScheduledPaymentsSpecificAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponseBankingScheduledPaymentsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScheduledPaymentsSpecificAccountsBadRequest creates a ListScheduledPaymentsSpecificAccountsBadRequest with default headers values
func NewListScheduledPaymentsSpecificAccountsBadRequest() *ListScheduledPaymentsSpecificAccountsBadRequest {
	return &ListScheduledPaymentsSpecificAccountsBadRequest{}
}

/* ListScheduledPaymentsSpecificAccountsBadRequest describes a response with status code 400, with default header values.

The following error codes MUST be supported:<br/><ul class="error-code-list"><li>[400 - Invalid Field](#error-400-field-invalid)</li><li>[400 - Invalid Page Size](#error-400-field-invalid-page-size)</li><li>[400 - Invalid Version](#error-400-header-invalid-version)</li><li>[406 - Unsupported Version](#error-406-header-unsupported-version)</li><li>[422 - Invalid Page](#error-422-field-invalid-page)</li><li>[422 - Unavailable Banking Account](#error-422-authorisation-unavailable-banking-account)</li><li>[422 - Invalid Banking Account](#error-422-authorisation-invalid-banking-account)</li></ul>
*/
type ListScheduledPaymentsSpecificAccountsBadRequest struct {

	/* An [RFC4122](https://tools.ietf.org/html/rfc4122) UUID used as a correlation id. If provided, the data holder must play back this value in the x-fapi-interaction-id response header. If not provided a [RFC4122] UUID value is required to be provided in the response header to track the interaction.
	 */
	XFapiInteractionID string

	Payload *models.ResponseErrorListV2
}

func (o *ListScheduledPaymentsSpecificAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /banking/payments/scheduled][%d] listScheduledPaymentsSpecificAccountsBadRequest  %+v", 400, o.Payload)
}
func (o *ListScheduledPaymentsSpecificAccountsBadRequest) GetPayload() *models.ResponseErrorListV2 {
	return o.Payload
}

func (o *ListScheduledPaymentsSpecificAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
// Code generated by go-swagger; DO NOT EDIT.

package pagamentos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/openbankingbr/payments/models"
)

// PaymentsPostConsentsReader is a Reader for the PaymentsPostConsents structure.
type PaymentsPostConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentsPostConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPaymentsPostConsentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentsPostConsentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentsPostConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentsPostConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentsPostConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPaymentsPostConsentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewPaymentsPostConsentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPaymentsPostConsentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPaymentsPostConsentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPaymentsPostConsentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentsPostConsentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPaymentsPostConsentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentsPostConsentsCreated creates a PaymentsPostConsentsCreated with default headers values
func NewPaymentsPostConsentsCreated() *PaymentsPostConsentsCreated {
	return &PaymentsPostConsentsCreated{}
}

/* PaymentsPostConsentsCreated describes a response with status code 201, with default header values.

Consentimento de pagamento criado com sucesso.
*/
type PaymentsPostConsentsCreated struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponsePaymentConsent
}

func (o *PaymentsPostConsentsCreated) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsCreated  %+v", 201, o.Payload)
}
func (o *PaymentsPostConsentsCreated) GetPayload() *models.OpenbankingBrasilResponsePaymentConsent {
	return o.Payload
}

func (o *PaymentsPostConsentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponsePaymentConsent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsBadRequest creates a PaymentsPostConsentsBadRequest with default headers values
func NewPaymentsPostConsentsBadRequest() *PaymentsPostConsentsBadRequest {
	return &PaymentsPostConsentsBadRequest{}
}

/* PaymentsPostConsentsBadRequest describes a response with status code 400, with default header values.

A requisio foi malformada, omitindo atributos obrigatrios, seja no payload ou atravs de atributos na URL.
*/
type PaymentsPostConsentsBadRequest struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsBadRequest  %+v", 400, o.Payload)
}
func (o *PaymentsPostConsentsBadRequest) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsUnauthorized creates a PaymentsPostConsentsUnauthorized with default headers values
func NewPaymentsPostConsentsUnauthorized() *PaymentsPostConsentsUnauthorized {
	return &PaymentsPostConsentsUnauthorized{}
}

/* PaymentsPostConsentsUnauthorized describes a response with status code 401, with default header values.

Cabealho de autenticao ausente/invlido ou token invlido
*/
type PaymentsPostConsentsUnauthorized struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsUnauthorized  %+v", 401, o.Payload)
}
func (o *PaymentsPostConsentsUnauthorized) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsForbidden creates a PaymentsPostConsentsForbidden with default headers values
func NewPaymentsPostConsentsForbidden() *PaymentsPostConsentsForbidden {
	return &PaymentsPostConsentsForbidden{}
}

/* PaymentsPostConsentsForbidden describes a response with status code 403, with default header values.

O token tem escopo incorreto ou uma poltica de segurana foi violada
*/
type PaymentsPostConsentsForbidden struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsForbidden) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsForbidden  %+v", 403, o.Payload)
}
func (o *PaymentsPostConsentsForbidden) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsNotFound creates a PaymentsPostConsentsNotFound with default headers values
func NewPaymentsPostConsentsNotFound() *PaymentsPostConsentsNotFound {
	return &PaymentsPostConsentsNotFound{}
}

/* PaymentsPostConsentsNotFound describes a response with status code 404, with default header values.

O recurso solicitado no existe ou no foi implementado
*/
type PaymentsPostConsentsNotFound struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsNotFound) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsNotFound  %+v", 404, o.Payload)
}
func (o *PaymentsPostConsentsNotFound) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsMethodNotAllowed creates a PaymentsPostConsentsMethodNotAllowed with default headers values
func NewPaymentsPostConsentsMethodNotAllowed() *PaymentsPostConsentsMethodNotAllowed {
	return &PaymentsPostConsentsMethodNotAllowed{}
}

/* PaymentsPostConsentsMethodNotAllowed describes a response with status code 405, with default header values.

O consumidor tentou acessar o recurso com um mtodo no suportado
*/
type PaymentsPostConsentsMethodNotAllowed struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *PaymentsPostConsentsMethodNotAllowed) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsNotAcceptable creates a PaymentsPostConsentsNotAcceptable with default headers values
func NewPaymentsPostConsentsNotAcceptable() *PaymentsPostConsentsNotAcceptable {
	return &PaymentsPostConsentsNotAcceptable{}
}

/* PaymentsPostConsentsNotAcceptable describes a response with status code 406, with default header values.

A solicitao continha um cabealho Accept diferente dos tipos de mdia permitidos ou um conjunto de caracteres diferente de UTF-8
*/
type PaymentsPostConsentsNotAcceptable struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsNotAcceptable  %+v", 406, o.Payload)
}
func (o *PaymentsPostConsentsNotAcceptable) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsUnsupportedMediaType creates a PaymentsPostConsentsUnsupportedMediaType with default headers values
func NewPaymentsPostConsentsUnsupportedMediaType() *PaymentsPostConsentsUnsupportedMediaType {
	return &PaymentsPostConsentsUnsupportedMediaType{}
}

/* PaymentsPostConsentsUnsupportedMediaType describes a response with status code 415, with default header values.

O formato do payload no  um formato suportado.
*/
type PaymentsPostConsentsUnsupportedMediaType struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *PaymentsPostConsentsUnsupportedMediaType) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsUnprocessableEntity creates a PaymentsPostConsentsUnprocessableEntity with default headers values
func NewPaymentsPostConsentsUnprocessableEntity() *PaymentsPostConsentsUnprocessableEntity {
	return &PaymentsPostConsentsUnprocessableEntity{}
}

/* PaymentsPostConsentsUnprocessableEntity describes a response with status code 422, with default header values.

A solicitao foi bem formada, mas no pde ser processada devido  lgica de negcios especfica da solicitao.
*/
type PaymentsPostConsentsUnprocessableEntity struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilUnprocessableEntityConsents1
}

func (o *PaymentsPostConsentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PaymentsPostConsentsUnprocessableEntity) GetPayload() *models.OpenbankingBrasilUnprocessableEntityConsents1 {
	return o.Payload
}

func (o *PaymentsPostConsentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilUnprocessableEntityConsents1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsTooManyRequests creates a PaymentsPostConsentsTooManyRequests with default headers values
func NewPaymentsPostConsentsTooManyRequests() *PaymentsPostConsentsTooManyRequests {
	return &PaymentsPostConsentsTooManyRequests{}
}

/* PaymentsPostConsentsTooManyRequests describes a response with status code 429, with default header values.

A operao foi recusada, pois muitas solicitaes foram feitas dentro de um determinado perodo ou o limite global de requisies concorrentes foi atingido
*/
type PaymentsPostConsentsTooManyRequests struct {

	/* Cabealho que indica o tempo (em segundos) que o cliente dever aguardar para realizar uma nova tentativa de chamada. Este cabealho dever estar presente quando o cdigo HTTP de retorno for 429 Too many requests.

	 */
	RetryAfter string

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsTooManyRequests  %+v", 429, o.Payload)
}
func (o *PaymentsPostConsentsTooManyRequests) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		o.RetryAfter = hdrRetryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsInternalServerError creates a PaymentsPostConsentsInternalServerError with default headers values
func NewPaymentsPostConsentsInternalServerError() *PaymentsPostConsentsInternalServerError {
	return &PaymentsPostConsentsInternalServerError{}
}

/* PaymentsPostConsentsInternalServerError describes a response with status code 500, with default header values.

Ocorreu um erro no gateway da API ou no microsservio
*/
type PaymentsPostConsentsInternalServerError struct {

	/* Um UUID RFC4122 usado como um ID de correlao. O transmissor deve usar o mesmo valor recebido na requisio para o cabealho de resposta recebido na requisio, caso no tenha sido fornecido, deve se usar um UUID RFC4122.

	 */
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilResponseError
}

func (o *PaymentsPostConsentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsentsInternalServerError  %+v", 500, o.Payload)
}
func (o *PaymentsPostConsentsInternalServerError) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentsPostConsentsDefault creates a PaymentsPostConsentsDefault with default headers values
func NewPaymentsPostConsentsDefault(code int) *PaymentsPostConsentsDefault {
	return &PaymentsPostConsentsDefault{
		_statusCode: code,
	}
}

/* PaymentsPostConsentsDefault describes a response with status code -1, with default header values.

Erro inesperado.
*/
type PaymentsPostConsentsDefault struct {
	_statusCode int

	Payload *models.OpenbankingBrasilResponseError
}

// Code gets the status code for the payments post consents default response
func (o *PaymentsPostConsentsDefault) Code() int {
	return o._statusCode
}

func (o *PaymentsPostConsentsDefault) Error() string {
	return fmt.Sprintf("[POST /consents][%d] paymentsPostConsents default  %+v", o._statusCode, o.Payload)
}
func (o *PaymentsPostConsentsDefault) GetPayload() *models.OpenbankingBrasilResponseError {
	return o.Payload
}

func (o *PaymentsPostConsentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

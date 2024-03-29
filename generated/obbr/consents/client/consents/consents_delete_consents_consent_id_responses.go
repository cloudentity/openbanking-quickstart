// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/obbr/consents/models"
)

// ConsentsDeleteConsentsConsentIDReader is a Reader for the ConsentsDeleteConsentsConsentID structure.
type ConsentsDeleteConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsentsDeleteConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConsentsDeleteConsentsConsentIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConsentsDeleteConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewConsentsDeleteConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConsentsDeleteConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConsentsDeleteConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewConsentsDeleteConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewConsentsDeleteConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewConsentsDeleteConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConsentsDeleteConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewConsentsDeleteConsentsConsentIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 529:
		result := NewConsentsDeleteConsentsConsentIDStatus529()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConsentsDeleteConsentsConsentIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsentsDeleteConsentsConsentIDNoContent creates a ConsentsDeleteConsentsConsentIDNoContent with default headers values
func NewConsentsDeleteConsentsConsentIDNoContent() *ConsentsDeleteConsentsConsentIDNoContent {
	return &ConsentsDeleteConsentsConsentIDNoContent{}
}

/* ConsentsDeleteConsentsConsentIDNoContent describes a response with status code 204, with default header values.

Consentimento revogado com sucesso.
*/
type ConsentsDeleteConsentsConsentIDNoContent struct {

	/* Um UID [RFC4122](https://tools.ietf.org/html/rfc4122) usado como um ID de correlao.
	Se fornecido, o transmissor deve "reproduzir" esse valor no cabealho de resposta.

	*/
	XFapiInteractionID string
}

func (o *ConsentsDeleteConsentsConsentIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdNoContent ", 204)
}

func (o *ConsentsDeleteConsentsConsentIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDBadRequest creates a ConsentsDeleteConsentsConsentIDBadRequest with default headers values
func NewConsentsDeleteConsentsConsentIDBadRequest() *ConsentsDeleteConsentsConsentIDBadRequest {
	return &ConsentsDeleteConsentsConsentIDBadRequest{}
}

/* ConsentsDeleteConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

A requisio foi malformada, omitindo atributos obrigatrios, seja no payload ou atravs de atributos na URL.
*/
type ConsentsDeleteConsentsConsentIDBadRequest struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDBadRequest) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDUnauthorized creates a ConsentsDeleteConsentsConsentIDUnauthorized with default headers values
func NewConsentsDeleteConsentsConsentIDUnauthorized() *ConsentsDeleteConsentsConsentIDUnauthorized {
	return &ConsentsDeleteConsentsConsentIDUnauthorized{}
}

/* ConsentsDeleteConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Cabealho de autenticao ausente/invlido ou token invlido
*/
type ConsentsDeleteConsentsConsentIDUnauthorized struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdUnauthorized  %+v", 401, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDUnauthorized) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDForbidden creates a ConsentsDeleteConsentsConsentIDForbidden with default headers values
func NewConsentsDeleteConsentsConsentIDForbidden() *ConsentsDeleteConsentsConsentIDForbidden {
	return &ConsentsDeleteConsentsConsentIDForbidden{}
}

/* ConsentsDeleteConsentsConsentIDForbidden describes a response with status code 403, with default header values.

O token tem escopo incorreto ou uma poltica de segurana foi violada
*/
type ConsentsDeleteConsentsConsentIDForbidden struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDForbidden) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDNotFound creates a ConsentsDeleteConsentsConsentIDNotFound with default headers values
func NewConsentsDeleteConsentsConsentIDNotFound() *ConsentsDeleteConsentsConsentIDNotFound {
	return &ConsentsDeleteConsentsConsentIDNotFound{}
}

/* ConsentsDeleteConsentsConsentIDNotFound describes a response with status code 404, with default header values.

O recurso solicitado no existe ou no foi implementado
*/
type ConsentsDeleteConsentsConsentIDNotFound struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdNotFound  %+v", 404, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDNotFound) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDMethodNotAllowed creates a ConsentsDeleteConsentsConsentIDMethodNotAllowed with default headers values
func NewConsentsDeleteConsentsConsentIDMethodNotAllowed() *ConsentsDeleteConsentsConsentIDMethodNotAllowed {
	return &ConsentsDeleteConsentsConsentIDMethodNotAllowed{}
}

/* ConsentsDeleteConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

O consumidor tentou acessar o recurso com um mtodo no suportado
*/
type ConsentsDeleteConsentsConsentIDMethodNotAllowed struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDMethodNotAllowed) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDNotAcceptable creates a ConsentsDeleteConsentsConsentIDNotAcceptable with default headers values
func NewConsentsDeleteConsentsConsentIDNotAcceptable() *ConsentsDeleteConsentsConsentIDNotAcceptable {
	return &ConsentsDeleteConsentsConsentIDNotAcceptable{}
}

/* ConsentsDeleteConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

A solicitao continha um cabealho Accept diferente dos tipos de mdia permitidos ou um conjunto de caracteres diferente de UTF-8
*/
type ConsentsDeleteConsentsConsentIDNotAcceptable struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdNotAcceptable  %+v", 406, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDNotAcceptable) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDTooManyRequests creates a ConsentsDeleteConsentsConsentIDTooManyRequests with default headers values
func NewConsentsDeleteConsentsConsentIDTooManyRequests() *ConsentsDeleteConsentsConsentIDTooManyRequests {
	return &ConsentsDeleteConsentsConsentIDTooManyRequests{}
}

/* ConsentsDeleteConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

A operao foi recusada, pois muitas solicitaes foram feitas dentro de um determinado perodo ou o limite global de requisies concorrentes foi atingido
*/
type ConsentsDeleteConsentsConsentIDTooManyRequests struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDTooManyRequests) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDInternalServerError creates a ConsentsDeleteConsentsConsentIDInternalServerError with default headers values
func NewConsentsDeleteConsentsConsentIDInternalServerError() *ConsentsDeleteConsentsConsentIDInternalServerError {
	return &ConsentsDeleteConsentsConsentIDInternalServerError{}
}

/* ConsentsDeleteConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Ocorreu um erro no gateway da API ou no microsservio
*/
type ConsentsDeleteConsentsConsentIDInternalServerError struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDInternalServerError) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDGatewayTimeout creates a ConsentsDeleteConsentsConsentIDGatewayTimeout with default headers values
func NewConsentsDeleteConsentsConsentIDGatewayTimeout() *ConsentsDeleteConsentsConsentIDGatewayTimeout {
	return &ConsentsDeleteConsentsConsentIDGatewayTimeout{}
}

/* ConsentsDeleteConsentsConsentIDGatewayTimeout describes a response with status code 504, with default header values.

GATEWAY TIMEOUT - A requisio no foi atendida dentro do tempo limite estabelecido
*/
type ConsentsDeleteConsentsConsentIDGatewayTimeout struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDGatewayTimeout) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDStatus529 creates a ConsentsDeleteConsentsConsentIDStatus529 with default headers values
func NewConsentsDeleteConsentsConsentIDStatus529() *ConsentsDeleteConsentsConsentIDStatus529 {
	return &ConsentsDeleteConsentsConsentIDStatus529{}
}

/* ConsentsDeleteConsentsConsentIDStatus529 describes a response with status code 529, with default header values.

O site est sobrecarregado e a operao foi recusada, pois foi atingido o limite mximo de TPS global, neste momento.
*/
type ConsentsDeleteConsentsConsentIDStatus529 struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsDeleteConsentsConsentIDStatus529) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentIdStatus529  %+v", 529, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDStatus529) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDStatus529) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsDeleteConsentsConsentIDDefault creates a ConsentsDeleteConsentsConsentIDDefault with default headers values
func NewConsentsDeleteConsentsConsentIDDefault(code int) *ConsentsDeleteConsentsConsentIDDefault {
	return &ConsentsDeleteConsentsConsentIDDefault{
		_statusCode: code,
	}
}

/* ConsentsDeleteConsentsConsentIDDefault describes a response with status code -1, with default header values.

Erro inesperado.
*/
type ConsentsDeleteConsentsConsentIDDefault struct {
	_statusCode int

	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

// Code gets the status code for the consents delete consents consent Id default response
func (o *ConsentsDeleteConsentsConsentIDDefault) Code() int {
	return o._statusCode
}

func (o *ConsentsDeleteConsentsConsentIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /consents/{consentId}][%d] consentsDeleteConsentsConsentId default  %+v", o._statusCode, o.Payload)
}
func (o *ConsentsDeleteConsentsConsentIDDefault) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsDeleteConsentsConsentIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

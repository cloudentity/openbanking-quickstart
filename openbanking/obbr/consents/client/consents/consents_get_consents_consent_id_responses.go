// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/consents/models"
)

// ConsentsGetConsentsConsentIDReader is a Reader for the ConsentsGetConsentsConsentID structure.
type ConsentsGetConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsentsGetConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsentsGetConsentsConsentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConsentsGetConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewConsentsGetConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConsentsGetConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConsentsGetConsentsConsentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewConsentsGetConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewConsentsGetConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewConsentsGetConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConsentsGetConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewConsentsGetConsentsConsentIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 529:
		result := NewConsentsGetConsentsConsentIDStatus529()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConsentsGetConsentsConsentIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsentsGetConsentsConsentIDOK creates a ConsentsGetConsentsConsentIDOK with default headers values
func NewConsentsGetConsentsConsentIDOK() *ConsentsGetConsentsConsentIDOK {
	return &ConsentsGetConsentsConsentIDOK{}
}

/* ConsentsGetConsentsConsentIDOK describes a response with status code 200, with default header values.

Consentimento consultado com sucesso.
*/
type ConsentsGetConsentsConsentIDOK struct {

	/* Um UID [RFC4122](https://tools.ietf.org/html/rfc4122) usado como um ID de correlao.
	Se fornecido, o transmissor deve "reproduzir" esse valor no cabealho de resposta.

	*/
	XFapiInteractionID string

	Payload *models.OpenbankingBrasilConsentV2ResponseConsentRead
}

func (o *ConsentsGetConsentsConsentIDOK) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdOK  %+v", 200, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDOK) GetPayload() *models.OpenbankingBrasilConsentV2ResponseConsentRead {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseConsentRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDBadRequest creates a ConsentsGetConsentsConsentIDBadRequest with default headers values
func NewConsentsGetConsentsConsentIDBadRequest() *ConsentsGetConsentsConsentIDBadRequest {
	return &ConsentsGetConsentsConsentIDBadRequest{}
}

/* ConsentsGetConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

A requisio foi malformada, omitindo atributos obrigatrios, seja no payload ou atravs de atributos na URL.
*/
type ConsentsGetConsentsConsentIDBadRequest struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDBadRequest) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDUnauthorized creates a ConsentsGetConsentsConsentIDUnauthorized with default headers values
func NewConsentsGetConsentsConsentIDUnauthorized() *ConsentsGetConsentsConsentIDUnauthorized {
	return &ConsentsGetConsentsConsentIDUnauthorized{}
}

/* ConsentsGetConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Cabealho de autenticao ausente/invlido ou token invlido
*/
type ConsentsGetConsentsConsentIDUnauthorized struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdUnauthorized  %+v", 401, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDUnauthorized) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDForbidden creates a ConsentsGetConsentsConsentIDForbidden with default headers values
func NewConsentsGetConsentsConsentIDForbidden() *ConsentsGetConsentsConsentIDForbidden {
	return &ConsentsGetConsentsConsentIDForbidden{}
}

/* ConsentsGetConsentsConsentIDForbidden describes a response with status code 403, with default header values.

O token tem escopo incorreto ou uma poltica de segurana foi violada
*/
type ConsentsGetConsentsConsentIDForbidden struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDForbidden) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDNotFound creates a ConsentsGetConsentsConsentIDNotFound with default headers values
func NewConsentsGetConsentsConsentIDNotFound() *ConsentsGetConsentsConsentIDNotFound {
	return &ConsentsGetConsentsConsentIDNotFound{}
}

/* ConsentsGetConsentsConsentIDNotFound describes a response with status code 404, with default header values.

O recurso solicitado no existe ou no foi implementado
*/
type ConsentsGetConsentsConsentIDNotFound struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdNotFound  %+v", 404, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDNotFound) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDMethodNotAllowed creates a ConsentsGetConsentsConsentIDMethodNotAllowed with default headers values
func NewConsentsGetConsentsConsentIDMethodNotAllowed() *ConsentsGetConsentsConsentIDMethodNotAllowed {
	return &ConsentsGetConsentsConsentIDMethodNotAllowed{}
}

/* ConsentsGetConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

O consumidor tentou acessar o recurso com um mtodo no suportado
*/
type ConsentsGetConsentsConsentIDMethodNotAllowed struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDMethodNotAllowed) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDNotAcceptable creates a ConsentsGetConsentsConsentIDNotAcceptable with default headers values
func NewConsentsGetConsentsConsentIDNotAcceptable() *ConsentsGetConsentsConsentIDNotAcceptable {
	return &ConsentsGetConsentsConsentIDNotAcceptable{}
}

/* ConsentsGetConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

A solicitao continha um cabealho Accept diferente dos tipos de mdia permitidos ou um conjunto de caracteres diferente de UTF-8
*/
type ConsentsGetConsentsConsentIDNotAcceptable struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdNotAcceptable  %+v", 406, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDNotAcceptable) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDTooManyRequests creates a ConsentsGetConsentsConsentIDTooManyRequests with default headers values
func NewConsentsGetConsentsConsentIDTooManyRequests() *ConsentsGetConsentsConsentIDTooManyRequests {
	return &ConsentsGetConsentsConsentIDTooManyRequests{}
}

/* ConsentsGetConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

A operao foi recusada, pois muitas solicitaes foram feitas dentro de um determinado perodo ou o limite global de requisies concorrentes foi atingido
*/
type ConsentsGetConsentsConsentIDTooManyRequests struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDTooManyRequests) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDInternalServerError creates a ConsentsGetConsentsConsentIDInternalServerError with default headers values
func NewConsentsGetConsentsConsentIDInternalServerError() *ConsentsGetConsentsConsentIDInternalServerError {
	return &ConsentsGetConsentsConsentIDInternalServerError{}
}

/* ConsentsGetConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Ocorreu um erro no gateway da API ou no microsservio
*/
type ConsentsGetConsentsConsentIDInternalServerError struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDInternalServerError) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDGatewayTimeout creates a ConsentsGetConsentsConsentIDGatewayTimeout with default headers values
func NewConsentsGetConsentsConsentIDGatewayTimeout() *ConsentsGetConsentsConsentIDGatewayTimeout {
	return &ConsentsGetConsentsConsentIDGatewayTimeout{}
}

/* ConsentsGetConsentsConsentIDGatewayTimeout describes a response with status code 504, with default header values.

GATEWAY TIMEOUT - A requisio no foi atendida dentro do tempo limite estabelecido
*/
type ConsentsGetConsentsConsentIDGatewayTimeout struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDGatewayTimeout) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDStatus529 creates a ConsentsGetConsentsConsentIDStatus529 with default headers values
func NewConsentsGetConsentsConsentIDStatus529() *ConsentsGetConsentsConsentIDStatus529 {
	return &ConsentsGetConsentsConsentIDStatus529{}
}

/* ConsentsGetConsentsConsentIDStatus529 describes a response with status code 529, with default header values.

O site est sobrecarregado e a operao foi recusada, pois foi atingido o limite mximo de TPS global, neste momento.
*/
type ConsentsGetConsentsConsentIDStatus529 struct {
	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

func (o *ConsentsGetConsentsConsentIDStatus529) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentIdStatus529  %+v", 529, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDStatus529) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDStatus529) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsentsGetConsentsConsentIDDefault creates a ConsentsGetConsentsConsentIDDefault with default headers values
func NewConsentsGetConsentsConsentIDDefault(code int) *ConsentsGetConsentsConsentIDDefault {
	return &ConsentsGetConsentsConsentIDDefault{
		_statusCode: code,
	}
}

/* ConsentsGetConsentsConsentIDDefault describes a response with status code -1, with default header values.

Erro inesperado.
*/
type ConsentsGetConsentsConsentIDDefault struct {
	_statusCode int

	Payload *models.OpenbankingBrasilConsentV2ResponseError
}

// Code gets the status code for the consents get consents consent Id default response
func (o *ConsentsGetConsentsConsentIDDefault) Code() int {
	return o._statusCode
}

func (o *ConsentsGetConsentsConsentIDDefault) Error() string {
	return fmt.Sprintf("[GET /consents/{consentId}][%d] consentsGetConsentsConsentId default  %+v", o._statusCode, o.Payload)
}
func (o *ConsentsGetConsentsConsentIDDefault) GetPayload() *models.OpenbankingBrasilConsentV2ResponseError {
	return o.Payload
}

func (o *ConsentsGetConsentsConsentIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenbankingBrasilConsentV2ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

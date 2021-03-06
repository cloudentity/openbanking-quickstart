// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenbankingBrasilResponsePixPaymentData ResponsePixPaymentData
//
// Objeto contendo dados do pagamento e da conta do recebedor (creditor).
//
// swagger:model OpenbankingBrasilResponsePixPaymentData
type OpenbankingBrasilResponsePixPaymentData struct {

	// Identificador nico do consentimento criado para a iniciao de pagamento solicitada. Dever ser um URN - Uniform Resource Name.
	// Um URN, conforme definido na [RFC8141](https://tools.ietf.org/html/rfc8141)  um Uniform Resource
	// Identifier - URI - que  atribudo sob o URI scheme "urn" e um namespace URN especfico, com a inteno de que o URN
	// seja um identificador de recurso persistente e independente da localizao.
	// Considerando a string urn:bancoex:C1DD33123 como exemplo para consentId temos:
	// - o namespace(urn)
	// - o identificador associado ao namespace da instituio transnmissora (bancoex)
	// - o identificador especfico dentro do namespace (C1DD33123).
	// Informaes mais detalhadas sobre a construo de namespaces devem ser consultadas na [RFC8141](https://tools.ietf.org/html/rfc8141).
	// Example: urn:bancoex:C1DD33123
	// Required: true
	// Max Length: 256
	// Pattern: ^urn:[a-zA-Z0-9][a-zA-Z0-9-]{0,31}:[a-zA-Z0-9()+,\-.:=@;$_!*'%\/?#]+$
	ConsentID string `json:"consentId"`

	// Data e hora em que o recurso foi criado.
	// Uma string com data e hora conforme especificao RFC-3339,
	// sempre com a utilizao de timezone UTC(UTC time format).
	// Example: 2020-07-21T08:30:00Z
	// Required: true
	// Format: date-time
	CreationDateTime strfmt.DateTime `json:"creationDateTime"`

	// creditor account
	// Required: true
	CreditorAccount *OpenbankingBrasilCreditorAccount `json:"creditorAccount"`

	// Deve ser preenchido no formato padro ExxxxxxxxyyyyMMddHHmmkkkkkkkkkkk (32 caracteres; case sensitive, isso , diferencia letras maisculas e minsculas), sendo:
	//  E  fixo (1 caractere);
	//  xxxxxxxx  identificao do agente que gerou o EndToEndId, podendo ser: o ISPB do participante direto ou o ISPB do participante indireto (8 caracteres numricos [0-9]);
	//  yyyyMMddHHmm  data, hora e minuto (12 caracteres), seguindo o horrio UTC, da submisso da ordem de pagamento, caso a liquidao seja prioritria, ou prevista para o envio da ordem ao sistema de liquidao, caso seja realizado um agendamento. Para ordens prioritrias e no prioritrias, aceita-se o preenchimento, pelo agente que gerou o EndToEndId, com uma tolerncia mxima de 12 horas, para o futuro e para o passado, em relao ao horrio efetivo de processamento da ordem pelo SPI;
	//  kkkkkkkkkkk  sequencial criado pelo agente que gerou o EndToEndId (11 caracteres alfanumricos [a-z/A-Z/0-9]). Deve ser nico dentro de cada yyyyMMddHHmm.
	// Admite-se que o EndToEndId seja gerado pelo participante direto, pelo participante indireto ou pelo iniciador de pagamento.
	// Ele deve ser nico, no podendo ser repetido em qualquer outra operao enviada ao SPI.
	// Example: E9040088820210128000800123873170
	// Required: true
	// Max Length: 32
	// Min Length: 32
	// Pattern: ^([E])([0-9]{8})([0-9]{4})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])(2[0-3]|[01][0-9])([0-5][0-9])([a-zA-Z0-9]{11})$
	EndToEndID string `json:"endToEndId"`

	// local instrument
	// Required: true
	LocalInstrument *OpenbankingBrasilEnumLocalInstrument `json:"localInstrument"`

	// payment
	// Required: true
	Payment *OpenbankingBrasilPaymentPix `json:"payment"`

	// Cdigo ou identificador nico informado pela instituio detentora da conta para representar
	// a iniciao de pagamento individual. O `paymentId` deve ser diferente do `endToEndId`.
	// Este  o identificador que dever ser utilizado na consulta ao status da iniciao de pagamento efetuada.
	// Example: TXpRMU9UQTROMWhZV2xSU1FUazJSMDl
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9][a-zA-Z0-9\-]{0,99}$
	PaymentID string `json:"paymentId"`

	// Chave cadastrada no DICT pertencente ao recebedor. Os tipos de chaves podem ser: telefone, e-mail, cpf/cnpj ou chave aleatria.
	// No caso de telefone celular deve ser informado no padro E.1641.
	// Para e-mail deve ter o formato xxxxxxxx@xxxxxxx.xxx(.xx) e no mximo 77 caracteres.
	// No caso de CPF dever ser informado com 11 nmeros, sem pontos ou traos.
	// Para o caso de CNPJ dever ser informado com 14 nmeros, sem pontos ou traos.
	// No caso de chave aleatria deve ser informado o UUID gerado pelo DICT, conforme formato especificado na RFC41223.
	// [Restrio] Obrigatrio quando o campo localInstrument for igual a DICT.
	// Example: 12345678901
	// Max Length: 77
	// Pattern: [\w\W\s]*
	Proxy *string `json:"proxy,omitempty"`

	// rejection reason
	RejectionReason OpenbankingBrasilRejectionReason `json:"rejectionReason,omitempty"`

	// Deve ser preenchido sempre que o usurio pagador inserir alguma informao adicional em um pagamento, a ser enviada ao recebedor.
	// Example: Pagamento da nota RSTO035-002.
	// Max Length: 140
	// Pattern: [\w\W\s]*
	RemittanceInformation *string `json:"remittanceInformation,omitempty"`

	// status
	// Required: true
	Status *OpenbankingBrasilStatus1 `json:"status"`

	// Data e hora da ltima atualizao da iniciao de pagamento.
	// Uma string com data e hora conforme especificao RFC-3339,
	// sempre com a utilizao de timezone UTC(UTC time format).
	// Example: 2020-07-21T08:30:00Z
	// Required: true
	// Format: date-time
	StatusUpdateDateTime strfmt.DateTime `json:"statusUpdateDateTime"`
}

// Validate validates this openbanking brasil response pix payment data
func (m *OpenbankingBrasilResponsePixPaymentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndToEndID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectionReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemittanceInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusUpdateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateConsentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("consentId", "body", m.ConsentID); err != nil {
		return err
	}

	if err := validate.MaxLength("consentId", "body", m.ConsentID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("consentId", "body", m.ConsentID, `^urn:[a-zA-Z0-9][a-zA-Z0-9-]{0,31}:[a-zA-Z0-9()+,\-.:=@;$_!*'%\/?#]+$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateCreationDateTime(formats strfmt.Registry) error {

	if err := validate.Required("creationDateTime", "body", strfmt.DateTime(m.CreationDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDateTime", "body", "date-time", m.CreationDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateCreditorAccount(formats strfmt.Registry) error {

	if err := validate.Required("creditorAccount", "body", m.CreditorAccount); err != nil {
		return err
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateEndToEndID(formats strfmt.Registry) error {

	if err := validate.RequiredString("endToEndId", "body", m.EndToEndID); err != nil {
		return err
	}

	if err := validate.MinLength("endToEndId", "body", m.EndToEndID, 32); err != nil {
		return err
	}

	if err := validate.MaxLength("endToEndId", "body", m.EndToEndID, 32); err != nil {
		return err
	}

	if err := validate.Pattern("endToEndId", "body", m.EndToEndID, `^([E])([0-9]{8})([0-9]{4})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])(2[0-3]|[01][0-9])([0-5][0-9])([a-zA-Z0-9]{11})$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateLocalInstrument(formats strfmt.Registry) error {

	if err := validate.Required("localInstrument", "body", m.LocalInstrument); err != nil {
		return err
	}

	if err := validate.Required("localInstrument", "body", m.LocalInstrument); err != nil {
		return err
	}

	if m.LocalInstrument != nil {
		if err := m.LocalInstrument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localInstrument")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validatePayment(formats strfmt.Registry) error {

	if err := validate.Required("payment", "body", m.Payment); err != nil {
		return err
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("paymentId", "body", m.PaymentID); err != nil {
		return err
	}

	if err := validate.MinLength("paymentId", "body", m.PaymentID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("paymentId", "body", m.PaymentID, 100); err != nil {
		return err
	}

	if err := validate.Pattern("paymentId", "body", m.PaymentID, `^[a-zA-Z0-9][a-zA-Z0-9\-]{0,99}$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if err := validate.MaxLength("proxy", "body", *m.Proxy, 77); err != nil {
		return err
	}

	if err := validate.Pattern("proxy", "body", *m.Proxy, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateRejectionReason(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectionReason) { // not required
		return nil
	}

	if err := m.RejectionReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rejectionReason")
		}
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateRemittanceInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.RemittanceInformation) { // not required
		return nil
	}

	if err := validate.MaxLength("remittanceInformation", "body", *m.RemittanceInformation, 140); err != nil {
		return err
	}

	if err := validate.Pattern("remittanceInformation", "body", *m.RemittanceInformation, `[\w\W\s]*`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) validateStatusUpdateDateTime(formats strfmt.Registry) error {

	if err := validate.Required("statusUpdateDateTime", "body", strfmt.DateTime(m.StatusUpdateDateTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("statusUpdateDateTime", "body", "date-time", m.StatusUpdateDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openbanking brasil response pix payment data based on the context it is used
func (m *OpenbankingBrasilResponsePixPaymentData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreditorAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalInstrument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRejectionReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) contextValidateCreditorAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) contextValidateLocalInstrument(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalInstrument != nil {
		if err := m.LocalInstrument.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localInstrument")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) contextValidateRejectionReason(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RejectionReason.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rejectionReason")
		}
		return err
	}

	return nil
}

func (m *OpenbankingBrasilResponsePixPaymentData) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingBrasilResponsePixPaymentData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingBrasilResponsePixPaymentData) UnmarshalBinary(b []byte) error {
	var res OpenbankingBrasilResponsePixPaymentData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

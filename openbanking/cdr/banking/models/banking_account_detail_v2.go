// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BankingAccountDetailV2 BankingAccountDetailV2
//
// swagger:model BankingAccountDetailV2
type BankingAccountDetailV2 struct {

	// A unique ID of the account adhering to the standards for ID permanence
	// Required: true
	AccountID *string `json:"accountId"`

	// The unmasked account number for the account. Should not be supplied if the account number is a PAN requiring PCI compliance. Is expected to be formatted as digits only with leading zeros included and no punctuation or spaces
	AccountNumber string `json:"accountNumber,omitempty"`

	// The addresses for the account to be used for correspondence
	Addresses []*CommonPhysicalAddress `json:"addresses"`

	// The unmasked BSB for the account. Is expected to be formatted as digits only with leading zeros included and no punctuation or spaces
	Bsb string `json:"bsb,omitempty"`

	// Optional field to indicate if this account is part of a bundle that is providing additional benefit for to the customer
	BundleName string `json:"bundleName,omitempty"`

	// Date that the account was created (if known)
	CreationDate string `json:"creationDate,omitempty"`

	// credit card
	CreditCard *BankingCreditCardAccount `json:"creditCard,omitempty"`

	// current rate to calculate interest earned being applied to deposit balances as it stands at the time of the API call
	DepositRate string `json:"depositRate,omitempty"`

	// Fully described deposit rates for this account based on the equivalent structure in Product Reference
	DepositRates []*BankingProductDepositRate `json:"depositRates"`

	// The display name of the account as defined by the bank. This should not incorporate account numbers or PANs. If it does the values should be masked according to the rules of the MaskedAccountString common type.
	// Required: true
	DisplayName *string `json:"displayName"`

	// Array of features of the account based on the equivalent structure in Product Reference with the following additional field
	Features []*Feature `json:"features"`

	// Fees and charges applicable to the account based on the equivalent structure in Product Reference
	Fees []*BankingProductFee `json:"fees"`

	// Flag indicating that the customer associated with the authorisation is an owner of the account. Does not indicate sole ownership, however. If not present then 'true' is assumed
	// Example: true
	IsOwned *bool `json:"isOwned,omitempty"`

	// The current rate to calculate interest payable being applied to lending balances as it stands at the time of the API call
	LendingRate string `json:"lendingRate,omitempty"`

	// Fully described deposit rates for this account based on the equivalent structure in Product Reference
	LendingRates []*BankingProductLendingRateV2 `json:"lendingRates"`

	// loan
	Loan *BankingLoanAccountV2 `json:"loan,omitempty"`

	// A masked version of the account. Whether BSB/Account Number, Credit Card PAN or another number
	// Required: true
	MaskedNumber *string `json:"maskedNumber"`

	// A customer supplied nick name for the account
	Nickname string `json:"nickname,omitempty"`

	// open status
	OpenStatus OpenStatus `json:"openStatus,omitempty"`

	// product category
	// Required: true
	ProductCategory *BankingProductCategory `json:"productCategory"`

	// The unique identifier of the account as defined by the data holder (akin to model number for the account)
	// Required: true
	ProductName *string `json:"productName"`

	// specific account u type
	SpecificAccountUType SpecificAccountUType `json:"specificAccountUType,omitempty"`

	// term deposit
	TermDeposit []*BankingTermDepositAccount `json:"termDeposit"`
}

// Validate validates this banking account detail v2
func (m *BankingAccountDetailV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepositRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLendingRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaskedNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecificAccountUType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingAccountDetailV2) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) validateCreditCard(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditCard) { // not required
		return nil
	}

	if m.CreditCard != nil {
		if err := m.CreditCard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditCard")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) validateDepositRates(formats strfmt.Registry) error {
	if swag.IsZero(m.DepositRates) { // not required
		return nil
	}

	for i := 0; i < len(m.DepositRates); i++ {
		if swag.IsZero(m.DepositRates[i]) { // not required
			continue
		}

		if m.DepositRates[i] != nil {
			if err := m.DepositRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("depositRates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) validateFees(formats strfmt.Registry) error {
	if swag.IsZero(m.Fees) { // not required
		return nil
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) validateLendingRates(formats strfmt.Registry) error {
	if swag.IsZero(m.LendingRates) { // not required
		return nil
	}

	for i := 0; i < len(m.LendingRates); i++ {
		if swag.IsZero(m.LendingRates[i]) { // not required
			continue
		}

		if m.LendingRates[i] != nil {
			if err := m.LendingRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lendingRates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) validateLoan(formats strfmt.Registry) error {
	if swag.IsZero(m.Loan) { // not required
		return nil
	}

	if m.Loan != nil {
		if err := m.Loan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loan")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) validateMaskedNumber(formats strfmt.Registry) error {

	if err := validate.Required("maskedNumber", "body", m.MaskedNumber); err != nil {
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateOpenStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenStatus) { // not required
		return nil
	}

	if err := m.OpenStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("openStatus")
		}
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateProductCategory(formats strfmt.Registry) error {

	if err := validate.Required("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	if err := validate.Required("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	if m.ProductCategory != nil {
		if err := m.ProductCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productCategory")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) validateProductName(formats strfmt.Registry) error {

	if err := validate.Required("productName", "body", m.ProductName); err != nil {
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateSpecificAccountUType(formats strfmt.Registry) error {
	if swag.IsZero(m.SpecificAccountUType) { // not required
		return nil
	}

	if err := m.SpecificAccountUType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("specificAccountUType")
		}
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) validateTermDeposit(formats strfmt.Registry) error {
	if swag.IsZero(m.TermDeposit) { // not required
		return nil
	}

	for i := 0; i < len(m.TermDeposit); i++ {
		if swag.IsZero(m.TermDeposit[i]) { // not required
			continue
		}

		if m.TermDeposit[i] != nil {
			if err := m.TermDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("termDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this banking account detail v2 based on the context it is used
func (m *BankingAccountDetailV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDepositRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLendingRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecificAccountUType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTermDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingAccountDetailV2) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateCreditCard(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditCard != nil {
		if err := m.CreditCard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditCard")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateDepositRates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DepositRates); i++ {

		if m.DepositRates[i] != nil {
			if err := m.DepositRates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("depositRates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Features); i++ {

		if m.Features[i] != nil {
			if err := m.Features[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fees); i++ {

		if m.Fees[i] != nil {
			if err := m.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateLendingRates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LendingRates); i++ {

		if m.LendingRates[i] != nil {
			if err := m.LendingRates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lendingRates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateLoan(ctx context.Context, formats strfmt.Registry) error {

	if m.Loan != nil {
		if err := m.Loan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loan")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateOpenStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OpenStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("openStatus")
		}
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateProductCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductCategory != nil {
		if err := m.ProductCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productCategory")
			}
			return err
		}
	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateSpecificAccountUType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SpecificAccountUType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("specificAccountUType")
		}
		return err
	}

	return nil
}

func (m *BankingAccountDetailV2) contextValidateTermDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TermDeposit); i++ {

		if m.TermDeposit[i] != nil {
			if err := m.TermDeposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("termDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingAccountDetailV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingAccountDetailV2) UnmarshalBinary(b []byte) error {
	var res BankingAccountDetailV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
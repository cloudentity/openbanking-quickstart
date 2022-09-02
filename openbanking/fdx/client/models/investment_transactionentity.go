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

// InvestmentTransactionentity InvestmentTransactionentity
//
// Specific transaction information
//
// swagger:model InvestmentTransactionentity
type InvestmentTransactionentity struct {

	// Corresponds to AccountId in Account
	// Max Length: 256
	AccountID string `json:"accountId,omitempty"`

	// Accrued interest
	AccruedInterest float64 `json:"accruedInterest,omitempty"`

	// The amount of money in the account currency
	Amount float64 `json:"amount,omitempty"`

	// Transaction category, preferably MCC or SIC.
	Category string `json:"category,omitempty"`

	// Transaction commission
	Commission float64 `json:"commission,omitempty"`

	// Confirmation number of the transaction
	ConfirmationNumber string `json:"confirmationNumber,omitempty"`

	// debit credit memo
	DebitCreditMemo DebitCreditMemo2 `json:"debitCreditMemo,omitempty"`

	// The description of the transaction
	Description string `json:"description,omitempty"`

	// Full precision unit number, unlimited digits after decimal point
	DigitalUnits string `json:"digitalUnits,omitempty"`

	// Cash value for bonds
	FaceValue float64 `json:"faceValue,omitempty"`

	// Fees applied to the trade
	Fees float64 `json:"fees,omitempty"`

	// Array of FI-specific attributes
	FiAttributes []*FIAttributeentity `json:"fiAttributes"`

	// The amount of money in the foreign currency
	ForeignAmount float64 `json:"foreignAmount,omitempty"`

	// foreign currency
	ForeignCurrency ISO4217Code3 `json:"foreignCurrency,omitempty"`

	// Cash for fractional units (used for stock splits)
	FractionalCash float64 `json:"fractionalCash,omitempty"`

	// For sales
	Gain float64 `json:"gain,omitempty"`

	// Array of Image Identifiers (unique to this transaction) used to retrieve Images of check or transaction receipt
	ImageIds []string `json:"imageIds"`

	// income type
	IncomeType IncomeType2 `json:"incomeType,omitempty"`

	// inv401k source
	Inv401kSource Investment401kSourceType2 `json:"inv401kSource,omitempty"`

	// Breakdown of the transaction details
	LineItem []*LineItementity `json:"lineItem"`

	// Links (unique to this transaction) used to retrieve images of checks or transaction receipts
	Links []*HATEOASLink `json:"links"`

	// Load on the transaction
	Load float64 `json:"load,omitempty"`

	// For 401k accounts only. This indicates the transaction was due to a loan or a loan repayment
	LoanID string `json:"loanId,omitempty"`

	// How much loan pre-payment is interest
	LoanInterest float64 `json:"loanInterest,omitempty"`

	// How much loan pre-payment is principal
	LoanPrincipal float64 `json:"loanPrincipal,omitempty"`

	// Portion of unit price that is attributed to the dealer markup
	Markup float64 `json:"markup,omitempty"`

	// Secondary transaction description
	// Max Length: 255
	Memo string `json:"memo,omitempty"`

	// Number of shares after split
	NewUnits float64 `json:"newUnits,omitempty"`

	// Number of shares before split
	OldUnits float64 `json:"oldUnits,omitempty"`

	// The date for the 401k transaction was obtained in payroll
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	PayrollDate strfmt.Date `json:"payrollDate,omitempty"`

	// Indicates amount withheld due to a penalty
	Penalty float64 `json:"penalty,omitempty"`

	// position type
	PositionType PositionType1 `json:"positionType,omitempty"`

	// The date and time that the transaction was posted to the account. If not provided then TransactionTimestamp can be used as PostedTimeStamp
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	PostedTimestamp strfmt.DateTime `json:"postedTimestamp,omitempty"`

	// Unit purchase price
	Price float64 `json:"price,omitempty"`

	// Indicates this buy was made using prior year's contribution
	PriorYearContrib bool `json:"priorYearContrib,omitempty"`

	// A tracking reference identifier
	Reference string `json:"reference,omitempty"`

	// For reverse postings, the identity of the transaction being reversed. For the correction transaction, the identity of the reversing post. For credit card posting transactions, the identity of the authorization transaction
	// Max Length: 256
	ReferenceTransactionID string `json:"referenceTransactionId,omitempty"`

	// reward
	Reward *TransactionRewardentity2 `json:"reward,omitempty"`

	// Running balance of the position
	RunningBalance float64 `json:"runningBalance,omitempty"`

	// Unique identifier of security
	SecurityID string `json:"securityId,omitempty"`

	// security Id type
	SecurityIDType SecurityIDType1 `json:"securityIdType,omitempty"`

	// security type
	SecurityType SecurityType2 `json:"securityType,omitempty"`

	// Required for stock, mutual funds. Number of shares (with decimals). Negative numbers indicate securities are being removed from the account
	Shares float64 `json:"shares,omitempty"`

	// Split ratio denominator
	SplitRatioDenominator float64 `json:"splitRatioDenominator,omitempty"`

	// Split ratio numerator
	SplitRatioNumerator float64 `json:"splitRatioNumerator,omitempty"`

	// State tax withholding
	StateWithholding float64 `json:"stateWithholding,omitempty"`

	// status
	Status TransactionStatus2 `json:"status,omitempty"`

	// sub account fund
	SubAccountFund SubAccountType2 `json:"subAccountFund,omitempty"`

	// sub account sec
	SubAccountSec SubAccountType1 `json:"subAccountSec,omitempty"`

	// Transaction category detail
	SubCategory string `json:"subCategory,omitempty"`

	// Ticker symbol
	Symbol string `json:"symbol,omitempty"`

	// Tax-exempt transaction
	TaxExempt bool `json:"taxExempt,omitempty"`

	// Taxes on the trade
	Taxes float64 `json:"taxes,omitempty"`

	// Long term persistent identity of the transaction (unique to account)
	// Max Length: 256
	TransactionID string `json:"transactionId,omitempty"`

	// transaction reason
	TransactionReason TransactionReason2 `json:"transactionReason,omitempty"`

	// The date and time that the transaction was added to the server backend systems
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	TransactionTimestamp strfmt.DateTime `json:"transactionTimestamp,omitempty"`

	// transaction type
	TransactionType InvestmentTransactionType2 `json:"transactionType,omitempty"`

	// transfer action
	TransferAction TransferAction `json:"transferAction,omitempty"`

	// Price per commonly-quoted unit. Does not include markup/markdown, unitprice. Share price for stocks, mutual funds, and others. Percentage of par for bonds. Per share (not contract) for options
	UnitPrice float64 `json:"unitPrice,omitempty"`

	// unit type
	UnitType UnitType1 `json:"unitType,omitempty"`

	// For security-based actions other than stock splits, quantity. Shares for stocks, mutual funds, and others. Face value for bonds. Contracts for options
	Units float64 `json:"units,omitempty"`

	// Federal tax withholding
	Withholding float64 `json:"withholding,omitempty"`
}

// Validate validates this investment transactionentity
func (m *InvestmentTransactionentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebitCreditMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncomeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInv401kSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayrollDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReward(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityIDType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAccountFund(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAccountSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvestmentTransactionentity) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("accountId", "body", m.AccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateDebitCreditMemo(formats strfmt.Registry) error {
	if swag.IsZero(m.DebitCreditMemo) { // not required
		return nil
	}

	if err := m.DebitCreditMemo.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debitCreditMemo")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debitCreditMemo")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateFiAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.FiAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.FiAttributes); i++ {
		if swag.IsZero(m.FiAttributes[i]) { // not required
			continue
		}

		if m.FiAttributes[i] != nil {
			if err := m.FiAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) validateForeignCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.ForeignCurrency) { // not required
		return nil
	}

	if err := m.ForeignCurrency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("foreignCurrency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("foreignCurrency")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateIncomeType(formats strfmt.Registry) error {
	if swag.IsZero(m.IncomeType) { // not required
		return nil
	}

	if err := m.IncomeType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("incomeType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("incomeType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateInv401kSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Inv401kSource) { // not required
		return nil
	}

	if err := m.Inv401kSource.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("inv401kSource")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("inv401kSource")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateLineItem(formats strfmt.Registry) error {
	if swag.IsZero(m.LineItem) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItem); i++ {
		if swag.IsZero(m.LineItem[i]) { // not required
			continue
		}

		if m.LineItem[i] != nil {
			if err := m.LineItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lineItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) validateMemo(formats strfmt.Registry) error {
	if swag.IsZero(m.Memo) { // not required
		return nil
	}

	if err := validate.MaxLength("memo", "body", m.Memo, 255); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validatePayrollDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PayrollDate) { // not required
		return nil
	}

	if err := validate.FormatOf("payrollDate", "body", "date", m.PayrollDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validatePositionType(formats strfmt.Registry) error {
	if swag.IsZero(m.PositionType) { // not required
		return nil
	}

	if err := m.PositionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("positionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("positionType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validatePostedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("postedTimestamp", "body", "date-time", m.PostedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateReferenceTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceTransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("referenceTransactionId", "body", m.ReferenceTransactionID, 256); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateReward(formats strfmt.Registry) error {
	if swag.IsZero(m.Reward) { // not required
		return nil
	}

	if m.Reward != nil {
		if err := m.Reward.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward")
			}
			return err
		}
	}

	return nil
}

func (m *InvestmentTransactionentity) validateSecurityIDType(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityIDType) { // not required
		return nil
	}

	if err := m.SecurityIDType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("securityIdType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("securityIdType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateSecurityType(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityType) { // not required
		return nil
	}

	if err := m.SecurityType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("securityType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("securityType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateSubAccountFund(formats strfmt.Registry) error {
	if swag.IsZero(m.SubAccountFund) { // not required
		return nil
	}

	if err := m.SubAccountFund.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subAccountFund")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subAccountFund")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateSubAccountSec(formats strfmt.Registry) error {
	if swag.IsZero(m.SubAccountSec) { // not required
		return nil
	}

	if err := m.SubAccountSec.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subAccountSec")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subAccountSec")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("transactionId", "body", m.TransactionID, 256); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateTransactionReason(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionReason) { // not required
		return nil
	}

	if err := m.TransactionReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionReason")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionReason")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateTransactionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("transactionTimestamp", "body", "date-time", m.TransactionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateTransactionType(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionType) { // not required
		return nil
	}

	if err := m.TransactionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateTransferAction(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferAction) { // not required
		return nil
	}

	if err := m.TransferAction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transferAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transferAction")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) validateUnitType(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitType) { // not required
		return nil
	}

	if err := m.UnitType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unitType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unitType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this investment transactionentity based on the context it is used
func (m *InvestmentTransactionentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDebitCreditMemo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForeignCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncomeType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInv401kSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLineItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePositionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReward(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityIDType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubAccountFund(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubAccountSec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactionReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransferAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnitType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvestmentTransactionentity) contextValidateDebitCreditMemo(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DebitCreditMemo.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("debitCreditMemo")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("debitCreditMemo")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateFiAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FiAttributes); i++ {

		if m.FiAttributes[i] != nil {
			if err := m.FiAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fiAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateForeignCurrency(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ForeignCurrency.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("foreignCurrency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("foreignCurrency")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateIncomeType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IncomeType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("incomeType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("incomeType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateInv401kSource(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Inv401kSource.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("inv401kSource")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("inv401kSource")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateLineItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineItem); i++ {

		if m.LineItem[i] != nil {
			if err := m.LineItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lineItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidatePositionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PositionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("positionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("positionType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateReward(ctx context.Context, formats strfmt.Registry) error {

	if m.Reward != nil {
		if err := m.Reward.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward")
			}
			return err
		}
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateSecurityIDType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecurityIDType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("securityIdType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("securityIdType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateSecurityType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecurityType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("securityType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("securityType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateSubAccountFund(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SubAccountFund.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subAccountFund")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subAccountFund")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateSubAccountSec(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SubAccountSec.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subAccountSec")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("subAccountSec")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateTransactionReason(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransactionReason.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionReason")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionReason")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateTransactionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransactionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactionType")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateTransferAction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransferAction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transferAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transferAction")
		}
		return err
	}

	return nil
}

func (m *InvestmentTransactionentity) contextValidateUnitType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UnitType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unitType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unitType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvestmentTransactionentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvestmentTransactionentity) UnmarshalBinary(b []byte) error {
	var res InvestmentTransactionentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
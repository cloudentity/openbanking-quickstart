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

// DepositAccountentity DepositAccountentity
//
// Information for a deposit account type
//
// swagger:model DepositAccountentity
type DepositAccountentity struct {

	// account category
	AccountCategory AccountCategorytype `json:"accountCategory,omitempty"`

	// Long-term persistent identity of the account, though not an account number. This identity must be unique to the owning institution
	// Max Length: 256
	AccountID string `json:"accountId,omitempty"`

	// Full account number for the end user's handle for the account at the owning institution
	AccountNumber string `json:"accountNumber,omitempty"`

	// Account display number for the end user's handle at the owning institution. This is to be displayed by the Interface Provider
	AccountNumberDisplay string `json:"accountNumberDisplay,omitempty"`

	// account type
	AccountType AccountType2 `json:"accountType,omitempty"`

	// Annual Percentage Yield
	AnnualPercentageYield float64 `json:"annualPercentageYield,omitempty"`

	// Balance of funds available for use
	AvailableBalance float64 `json:"availableBalance,omitempty"`

	// As-of date of balances
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	BalanceAsOf strfmt.DateTime `json:"balanceAsOf,omitempty"`

	// balance type
	BalanceType BalanceType2 `json:"balanceType,omitempty"`

	// bill pay status
	BillPayStatus AccountBillPayStatus2 `json:"billPayStatus,omitempty"`

	// contact
	Contact *AccountContactentity2 `json:"contact,omitempty"`

	// currency
	Currency *Currencyentity4 `json:"currency,omitempty"`

	// Balance of funds in account
	CurrentBalance float64 `json:"currentBalance,omitempty"`

	// Description of account
	Description string `json:"description,omitempty"`

	// error
	Error *Error1 `json:"error,omitempty"`

	// Array of Financial institution-specific attributes
	FiAttributes []*FIAttributeentity `json:"fiAttributes"`

	// Interest Rate of Account
	InterestRate float64 `json:"interestRate,omitempty"`

	// Date of account's interest rate
	// Example: 2021-07-15T14:46:41.375Z
	// Format: date-time
	InterestRateAsOf strfmt.DateTime `json:"interestRateAsOf,omitempty"`

	// interest rate type
	InterestRateType InterestRateType2 `json:"interestRateType,omitempty"`

	// YTD Interest
	InterestYtd float64 `json:"interestYtd,omitempty"`

	// Date that last transaction occurred on account
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	LastActivityDate strfmt.Date `json:"lastActivityDate,omitempty"`

	// The line of business, such as consumer, consumer joint, small business, corporate, etc.
	LineOfBusiness string `json:"lineOfBusiness,omitempty"`

	// Maturity date for CDs
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	MaturityDate strfmt.Date `json:"maturityDate,omitempty"`

	// MICR Number
	// Max Length: 64
	MicrNumber string `json:"micrNumber,omitempty"`

	// Name given by the user. Used in UIs to assist in account selection
	Nickname string `json:"nickname,omitempty"`

	// Day's opening fund balance
	OpeningDayBalance float64 `json:"openingDayBalance,omitempty"`

	// Long-term persistent identity of the parent account. This is used to group accounts
	// Max Length: 256
	ParentAccountID string `json:"parentAccountId,omitempty"`

	// Previous Interest Rate of Account
	PriorInterestRate float64 `json:"priorInterestRate,omitempty"`

	// Marketed product name for this account. Used in UIs to assist in account selection
	ProductName string `json:"productName,omitempty"`

	// Long-term persistent identity of rewards program associated with this account
	// Max Length: 256
	RewardProgramID string `json:"rewardProgramId,omitempty"`

	// Routing transit number (RTN) associated with account number at owning institution
	RoutingTransitNumber string `json:"routingTransitNumber,omitempty"`

	// status
	Status AccountStatus2 `json:"status,omitempty"`

	// Term of CD in months
	Term int32 `json:"term,omitempty"`

	// Transactions on the deposit account
	Transactions []*DepositTransactionentity `json:"transactions"`

	// Default is false. If present and true, a call to retrieve transactions will not return any further details about this account. This is an optimization that allows an FDX API server to return transactions and account details in a single call
	TransactionsIncluded bool `json:"transactionsIncluded,omitempty"`

	// Account is eligible for incoming transfers
	TransferIn bool `json:"transferIn,omitempty"`

	// Account is eligible for outgoing transfers
	TransferOut bool `json:"transferOut,omitempty"`
}

// Validate validates this deposit accountentity
func (m *DepositAccountentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanceAsOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillPayStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterestRateAsOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterestRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActivityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaturityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMicrNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepositAccountentity) validateAccountCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountCategory) { // not required
		return nil
	}

	if err := m.AccountCategory.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountCategory")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountCategory")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("accountId", "body", m.AccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateBalanceAsOf(formats strfmt.Registry) error {
	if swag.IsZero(m.BalanceAsOf) { // not required
		return nil
	}

	if err := validate.FormatOf("balanceAsOf", "body", "date-time", m.BalanceAsOf.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateBalanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.BalanceType) { // not required
		return nil
	}

	if err := m.BalanceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("balanceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("balanceType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateBillPayStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.BillPayStatus) { // not required
		return nil
	}

	if err := m.BillPayStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("billPayStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("billPayStatus")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateContact(formats strfmt.Registry) error {
	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) validateFiAttributes(formats strfmt.Registry) error {
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

func (m *DepositAccountentity) validateInterestRateAsOf(formats strfmt.Registry) error {
	if swag.IsZero(m.InterestRateAsOf) { // not required
		return nil
	}

	if err := validate.FormatOf("interestRateAsOf", "body", "date-time", m.InterestRateAsOf.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateInterestRateType(formats strfmt.Registry) error {
	if swag.IsZero(m.InterestRateType) { // not required
		return nil
	}

	if err := m.InterestRateType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interestRateType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interestRateType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateLastActivityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActivityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastActivityDate", "body", "date", m.LastActivityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateMaturityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.MaturityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("maturityDate", "body", "date", m.MaturityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateMicrNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.MicrNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("micrNumber", "body", m.MicrNumber, 64); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateParentAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentAccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("parentAccountId", "body", m.ParentAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateRewardProgramID(formats strfmt.Registry) error {
	if swag.IsZero(m.RewardProgramID) { // not required
		return nil
	}

	if err := validate.MaxLength("rewardProgramId", "body", m.RewardProgramID, 256); err != nil {
		return err
	}

	return nil
}

func (m *DepositAccountentity) validateStatus(formats strfmt.Registry) error {
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

func (m *DepositAccountentity) validateTransactions(formats strfmt.Registry) error {
	if swag.IsZero(m.Transactions) { // not required
		return nil
	}

	for i := 0; i < len(m.Transactions); i++ {
		if swag.IsZero(m.Transactions[i]) { // not required
			continue
		}

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this deposit accountentity based on the context it is used
func (m *DepositAccountentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBalanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBillPayStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterestRateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DepositAccountentity) contextValidateAccountCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountCategory.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountCategory")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountCategory")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) contextValidateBalanceType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BalanceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("balanceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("balanceType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) contextValidateBillPayStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BillPayStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("billPayStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("billPayStatus")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) contextValidateContact(ctx context.Context, formats strfmt.Registry) error {

	if m.Contact != nil {
		if err := m.Contact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if m.Currency != nil {
		if err := m.Currency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *DepositAccountentity) contextValidateFiAttributes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DepositAccountentity) contextValidateInterestRateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InterestRateType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interestRateType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interestRateType")
		}
		return err
	}

	return nil
}

func (m *DepositAccountentity) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DepositAccountentity) contextValidateTransactions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Transactions); i++ {

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DepositAccountentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepositAccountentity) UnmarshalBinary(b []byte) error {
	var res DepositAccountentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

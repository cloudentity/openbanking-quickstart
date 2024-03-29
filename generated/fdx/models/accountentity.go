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

// Accountentity Accountentity
//
// An abstract account entity that concrete account entities extend
//
// swagger:model Accountentity
type Accountentity struct {

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

	// balance type
	BalanceType BalanceType2 `json:"balanceType,omitempty"`

	// bill pay status
	BillPayStatus AccountBillPayStatus2 `json:"billPayStatus,omitempty"`

	// contact
	Contact *AccountContactentity2 `json:"contact,omitempty"`

	// currency
	Currency *Currencyentity4 `json:"currency,omitempty"`

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

	// Date that last transaction occurred on account
	// Example: 2021-07-15T00:00:00.000Z
	// Format: date
	LastActivityDate strfmt.Date `json:"lastActivityDate,omitempty"`

	// The line of business, such as consumer, consumer joint, small business, corporate, etc.
	LineOfBusiness string `json:"lineOfBusiness,omitempty"`

	// MICR Number
	// Max Length: 64
	MicrNumber string `json:"micrNumber,omitempty"`

	// Name given by the user. Used in UIs to assist in account selection
	Nickname string `json:"nickname,omitempty"`

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

	// Default is false. If present and true, a call to retrieve transactions will not return any further details about this account. This is an optimization that allows an FDX API server to return transactions and account details in a single call
	TransactionsIncluded bool `json:"transactionsIncluded,omitempty"`

	// Account is eligible for incoming transfers
	TransferIn bool `json:"transferIn,omitempty"`

	// Account is eligible for outgoing transfers
	TransferOut bool `json:"transferOut,omitempty"`
}

// Validate validates this accountentity
func (m *Accountentity) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Accountentity) validateAccountCategory(formats strfmt.Registry) error {
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

func (m *Accountentity) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("accountId", "body", m.AccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateAccountType(formats strfmt.Registry) error {
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

func (m *Accountentity) validateBalanceType(formats strfmt.Registry) error {
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

func (m *Accountentity) validateBillPayStatus(formats strfmt.Registry) error {
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

func (m *Accountentity) validateContact(formats strfmt.Registry) error {
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

func (m *Accountentity) validateCurrency(formats strfmt.Registry) error {
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

func (m *Accountentity) validateError(formats strfmt.Registry) error {
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

func (m *Accountentity) validateFiAttributes(formats strfmt.Registry) error {
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

func (m *Accountentity) validateInterestRateAsOf(formats strfmt.Registry) error {
	if swag.IsZero(m.InterestRateAsOf) { // not required
		return nil
	}

	if err := validate.FormatOf("interestRateAsOf", "body", "date-time", m.InterestRateAsOf.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateInterestRateType(formats strfmt.Registry) error {
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

func (m *Accountentity) validateLastActivityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActivityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastActivityDate", "body", "date", m.LastActivityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateMicrNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.MicrNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("micrNumber", "body", m.MicrNumber, 64); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateParentAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentAccountID) { // not required
		return nil
	}

	if err := validate.MaxLength("parentAccountId", "body", m.ParentAccountID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateRewardProgramID(formats strfmt.Registry) error {
	if swag.IsZero(m.RewardProgramID) { // not required
		return nil
	}

	if err := validate.MaxLength("rewardProgramId", "body", m.RewardProgramID, 256); err != nil {
		return err
	}

	return nil
}

func (m *Accountentity) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this accountentity based on the context it is used
func (m *Accountentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Accountentity) contextValidateAccountCategory(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateBalanceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateBillPayStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateContact(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateFiAttributes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateInterestRateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Accountentity) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Accountentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Accountentity) UnmarshalBinary(b []byte) error {
	var res Accountentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

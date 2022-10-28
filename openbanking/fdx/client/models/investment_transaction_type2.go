// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// InvestmentTransactionType2 InvestmentTransactionType2
//
// PURCHASED, SOLD, PURCHASEDTOCOVER, ADJUSTMENT, PURCHASETOOPEN, PURCHASETOCLOSE, SOLDTOOPEN, SOLDTOCLOSE, INTEREST, MARGININTEREST, REINVESTOFINCOME, RETURNOFCAPITAL, TRANSFER, CONTRIBUTION, FEE, OPTIONEXERCISE, OPTIONEXPIRATION, DIVIDEND, DIVIDENDREINVEST, SPLIT, CLOSURE, INCOME, EXPENSE, CLOSUREOPT, INVEXPENSE, JRNLSEC, JRNLFUND, OTHER, DIV, SRVCHG, DEP, DEPOSIT, ATM, POS, XFER, CHECK, PAYMENT, CASH, DIRECTDEP, DIRECTDEBIT, REPEATPMT
//
// swagger:model InvestmentTransactionType2
type InvestmentTransactionType2 string

func NewInvestmentTransactionType2(value InvestmentTransactionType2) *InvestmentTransactionType2 {
	v := value
	return &v
}

const (

	// InvestmentTransactionType2ADJUSTMENT captures enum value "ADJUSTMENT"
	InvestmentTransactionType2ADJUSTMENT InvestmentTransactionType2 = "ADJUSTMENT"

	// InvestmentTransactionType2AIRDROP captures enum value "AIRDROP"
	InvestmentTransactionType2AIRDROP InvestmentTransactionType2 = "AIRDROP"

	// InvestmentTransactionType2ATM captures enum value "ATM"
	InvestmentTransactionType2ATM InvestmentTransactionType2 = "ATM"

	// InvestmentTransactionType2CASH captures enum value "CASH"
	InvestmentTransactionType2CASH InvestmentTransactionType2 = "CASH"

	// InvestmentTransactionType2CHECK captures enum value "CHECK"
	InvestmentTransactionType2CHECK InvestmentTransactionType2 = "CHECK"

	// InvestmentTransactionType2CLOSURE captures enum value "CLOSURE"
	InvestmentTransactionType2CLOSURE InvestmentTransactionType2 = "CLOSURE"

	// InvestmentTransactionType2CLOSUREOPT captures enum value "CLOSUREOPT"
	InvestmentTransactionType2CLOSUREOPT InvestmentTransactionType2 = "CLOSUREOPT"

	// InvestmentTransactionType2CONTRIBUTION captures enum value "CONTRIBUTION"
	InvestmentTransactionType2CONTRIBUTION InvestmentTransactionType2 = "CONTRIBUTION"

	// InvestmentTransactionType2DEP captures enum value "DEP"
	InvestmentTransactionType2DEP InvestmentTransactionType2 = "DEP"

	// InvestmentTransactionType2DEPOSIT captures enum value "DEPOSIT"
	InvestmentTransactionType2DEPOSIT InvestmentTransactionType2 = "DEPOSIT"

	// InvestmentTransactionType2DIRECTDEBIT captures enum value "DIRECTDEBIT"
	InvestmentTransactionType2DIRECTDEBIT InvestmentTransactionType2 = "DIRECTDEBIT"

	// InvestmentTransactionType2DIRECTDEP captures enum value "DIRECTDEP"
	InvestmentTransactionType2DIRECTDEP InvestmentTransactionType2 = "DIRECTDEP"

	// InvestmentTransactionType2DIV captures enum value "DIV"
	InvestmentTransactionType2DIV InvestmentTransactionType2 = "DIV"

	// InvestmentTransactionType2DIVIDEND captures enum value "DIVIDEND"
	InvestmentTransactionType2DIVIDEND InvestmentTransactionType2 = "DIVIDEND"

	// InvestmentTransactionType2DIVIDENDREINVEST captures enum value "DIVIDENDREINVEST"
	InvestmentTransactionType2DIVIDENDREINVEST InvestmentTransactionType2 = "DIVIDENDREINVEST"

	// InvestmentTransactionType2EXPENSE captures enum value "EXPENSE"
	InvestmentTransactionType2EXPENSE InvestmentTransactionType2 = "EXPENSE"

	// InvestmentTransactionType2FEE captures enum value "FEE"
	InvestmentTransactionType2FEE InvestmentTransactionType2 = "FEE"

	// InvestmentTransactionType2FORKED captures enum value "FORKED"
	InvestmentTransactionType2FORKED InvestmentTransactionType2 = "FORKED"

	// InvestmentTransactionType2INCOME captures enum value "INCOME"
	InvestmentTransactionType2INCOME InvestmentTransactionType2 = "INCOME"

	// InvestmentTransactionType2INTEREST captures enum value "INTEREST"
	InvestmentTransactionType2INTEREST InvestmentTransactionType2 = "INTEREST"

	// InvestmentTransactionType2INVEXPENSE captures enum value "INVEXPENSE"
	InvestmentTransactionType2INVEXPENSE InvestmentTransactionType2 = "INVEXPENSE"

	// InvestmentTransactionType2JRNLFUND captures enum value "JRNLFUND"
	InvestmentTransactionType2JRNLFUND InvestmentTransactionType2 = "JRNLFUND"

	// InvestmentTransactionType2JRNLSEC captures enum value "JRNLSEC"
	InvestmentTransactionType2JRNLSEC InvestmentTransactionType2 = "JRNLSEC"

	// InvestmentTransactionType2MARGININTEREST captures enum value "MARGININTEREST"
	InvestmentTransactionType2MARGININTEREST InvestmentTransactionType2 = "MARGININTEREST"

	// InvestmentTransactionType2MINED captures enum value "MINED"
	InvestmentTransactionType2MINED InvestmentTransactionType2 = "MINED"

	// InvestmentTransactionType2OPTIONEXERCISE captures enum value "OPTIONEXERCISE"
	InvestmentTransactionType2OPTIONEXERCISE InvestmentTransactionType2 = "OPTIONEXERCISE"

	// InvestmentTransactionType2OPTIONEXPIRATION captures enum value "OPTIONEXPIRATION"
	InvestmentTransactionType2OPTIONEXPIRATION InvestmentTransactionType2 = "OPTIONEXPIRATION"

	// InvestmentTransactionType2OTHER captures enum value "OTHER"
	InvestmentTransactionType2OTHER InvestmentTransactionType2 = "OTHER"

	// InvestmentTransactionType2PAYMENT captures enum value "PAYMENT"
	InvestmentTransactionType2PAYMENT InvestmentTransactionType2 = "PAYMENT"

	// InvestmentTransactionType2POS captures enum value "POS"
	InvestmentTransactionType2POS InvestmentTransactionType2 = "POS"

	// InvestmentTransactionType2PURCHASED captures enum value "PURCHASED"
	InvestmentTransactionType2PURCHASED InvestmentTransactionType2 = "PURCHASED"

	// InvestmentTransactionType2PURCHASEDTOCOVER captures enum value "PURCHASEDTOCOVER"
	InvestmentTransactionType2PURCHASEDTOCOVER InvestmentTransactionType2 = "PURCHASEDTOCOVER"

	// InvestmentTransactionType2PURCHASETOCLOSE captures enum value "PURCHASETOCLOSE"
	InvestmentTransactionType2PURCHASETOCLOSE InvestmentTransactionType2 = "PURCHASETOCLOSE"

	// InvestmentTransactionType2PURCHASETOOPEN captures enum value "PURCHASETOOPEN"
	InvestmentTransactionType2PURCHASETOOPEN InvestmentTransactionType2 = "PURCHASETOOPEN"

	// InvestmentTransactionType2REINVESTOFINCOME captures enum value "REINVESTOFINCOME"
	InvestmentTransactionType2REINVESTOFINCOME InvestmentTransactionType2 = "REINVESTOFINCOME"

	// InvestmentTransactionType2REPEATPMT captures enum value "REPEATPMT"
	InvestmentTransactionType2REPEATPMT InvestmentTransactionType2 = "REPEATPMT"

	// InvestmentTransactionType2RETURNOFCAPITAL captures enum value "RETURNOFCAPITAL"
	InvestmentTransactionType2RETURNOFCAPITAL InvestmentTransactionType2 = "RETURNOFCAPITAL"

	// InvestmentTransactionType2SOLD captures enum value "SOLD"
	InvestmentTransactionType2SOLD InvestmentTransactionType2 = "SOLD"

	// InvestmentTransactionType2SOLDTOCLOSE captures enum value "SOLDTOCLOSE"
	InvestmentTransactionType2SOLDTOCLOSE InvestmentTransactionType2 = "SOLDTOCLOSE"

	// InvestmentTransactionType2SOLDTOOPEN captures enum value "SOLDTOOPEN"
	InvestmentTransactionType2SOLDTOOPEN InvestmentTransactionType2 = "SOLDTOOPEN"

	// InvestmentTransactionType2SPLIT captures enum value "SPLIT"
	InvestmentTransactionType2SPLIT InvestmentTransactionType2 = "SPLIT"

	// InvestmentTransactionType2SRVCHG captures enum value "SRVCHG"
	InvestmentTransactionType2SRVCHG InvestmentTransactionType2 = "SRVCHG"

	// InvestmentTransactionType2STAKED captures enum value "STAKED"
	InvestmentTransactionType2STAKED InvestmentTransactionType2 = "STAKED"

	// InvestmentTransactionType2TRANSFER captures enum value "TRANSFER"
	InvestmentTransactionType2TRANSFER InvestmentTransactionType2 = "TRANSFER"

	// InvestmentTransactionType2WITHDRAWAL captures enum value "WITHDRAWAL"
	InvestmentTransactionType2WITHDRAWAL InvestmentTransactionType2 = "WITHDRAWAL"

	// InvestmentTransactionType2XFER captures enum value "XFER"
	InvestmentTransactionType2XFER InvestmentTransactionType2 = "XFER"
)

// for schema
var investmentTransactionType2Enum []interface{}

func init() {
	var res []InvestmentTransactionType2
	if err := json.Unmarshal([]byte(`["ADJUSTMENT","AIRDROP","ATM","CASH","CHECK","CLOSURE","CLOSUREOPT","CONTRIBUTION","DEP","DEPOSIT","DIRECTDEBIT","DIRECTDEP","DIV","DIVIDEND","DIVIDENDREINVEST","EXPENSE","FEE","FORKED","INCOME","INTEREST","INVEXPENSE","JRNLFUND","JRNLSEC","MARGININTEREST","MINED","OPTIONEXERCISE","OPTIONEXPIRATION","OTHER","PAYMENT","POS","PURCHASED","PURCHASEDTOCOVER","PURCHASETOCLOSE","PURCHASETOOPEN","REINVESTOFINCOME","REPEATPMT","RETURNOFCAPITAL","SOLD","SOLDTOCLOSE","SOLDTOOPEN","SPLIT","SRVCHG","STAKED","TRANSFER","WITHDRAWAL","XFER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		investmentTransactionType2Enum = append(investmentTransactionType2Enum, v)
	}
}

func (m InvestmentTransactionType2) validateInvestmentTransactionType2Enum(path, location string, value InvestmentTransactionType2) error {
	if err := validate.EnumCase(path, location, value, investmentTransactionType2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this investment transaction type2
func (m InvestmentTransactionType2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvestmentTransactionType2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this investment transaction type2 based on context it is used
func (m InvestmentTransactionType2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
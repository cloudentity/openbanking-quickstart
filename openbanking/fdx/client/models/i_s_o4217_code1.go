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

// ISO4217Code1 ISO4217Code1
//
// ISO 4217 currency code
//
// swagger:model ISO4217Code1
type ISO4217Code1 string

func NewISO4217Code1(value ISO4217Code1) *ISO4217Code1 {
	v := value
	return &v
}

const (

	// ISO4217Code1AED captures enum value "AED"
	ISO4217Code1AED ISO4217Code1 = "AED"

	// ISO4217Code1AFN captures enum value "AFN"
	ISO4217Code1AFN ISO4217Code1 = "AFN"

	// ISO4217Code1ALL captures enum value "ALL"
	ISO4217Code1ALL ISO4217Code1 = "ALL"

	// ISO4217Code1AMD captures enum value "AMD"
	ISO4217Code1AMD ISO4217Code1 = "AMD"

	// ISO4217Code1ANG captures enum value "ANG"
	ISO4217Code1ANG ISO4217Code1 = "ANG"

	// ISO4217Code1AOA captures enum value "AOA"
	ISO4217Code1AOA ISO4217Code1 = "AOA"

	// ISO4217Code1ARS captures enum value "ARS"
	ISO4217Code1ARS ISO4217Code1 = "ARS"

	// ISO4217Code1AUD captures enum value "AUD"
	ISO4217Code1AUD ISO4217Code1 = "AUD"

	// ISO4217Code1AWG captures enum value "AWG"
	ISO4217Code1AWG ISO4217Code1 = "AWG"

	// ISO4217Code1AZN captures enum value "AZN"
	ISO4217Code1AZN ISO4217Code1 = "AZN"

	// ISO4217Code1BAM captures enum value "BAM"
	ISO4217Code1BAM ISO4217Code1 = "BAM"

	// ISO4217Code1BBD captures enum value "BBD"
	ISO4217Code1BBD ISO4217Code1 = "BBD"

	// ISO4217Code1BDT captures enum value "BDT"
	ISO4217Code1BDT ISO4217Code1 = "BDT"

	// ISO4217Code1BGN captures enum value "BGN"
	ISO4217Code1BGN ISO4217Code1 = "BGN"

	// ISO4217Code1BHD captures enum value "BHD"
	ISO4217Code1BHD ISO4217Code1 = "BHD"

	// ISO4217Code1BIF captures enum value "BIF"
	ISO4217Code1BIF ISO4217Code1 = "BIF"

	// ISO4217Code1BMD captures enum value "BMD"
	ISO4217Code1BMD ISO4217Code1 = "BMD"

	// ISO4217Code1BND captures enum value "BND"
	ISO4217Code1BND ISO4217Code1 = "BND"

	// ISO4217Code1BOB captures enum value "BOB"
	ISO4217Code1BOB ISO4217Code1 = "BOB"

	// ISO4217Code1BOV captures enum value "BOV"
	ISO4217Code1BOV ISO4217Code1 = "BOV"

	// ISO4217Code1BRL captures enum value "BRL"
	ISO4217Code1BRL ISO4217Code1 = "BRL"

	// ISO4217Code1BSD captures enum value "BSD"
	ISO4217Code1BSD ISO4217Code1 = "BSD"

	// ISO4217Code1BTN captures enum value "BTN"
	ISO4217Code1BTN ISO4217Code1 = "BTN"

	// ISO4217Code1BWP captures enum value "BWP"
	ISO4217Code1BWP ISO4217Code1 = "BWP"

	// ISO4217Code1BYR captures enum value "BYR"
	ISO4217Code1BYR ISO4217Code1 = "BYR"

	// ISO4217Code1BZD captures enum value "BZD"
	ISO4217Code1BZD ISO4217Code1 = "BZD"

	// ISO4217Code1CAD captures enum value "CAD"
	ISO4217Code1CAD ISO4217Code1 = "CAD"

	// ISO4217Code1CDF captures enum value "CDF"
	ISO4217Code1CDF ISO4217Code1 = "CDF"

	// ISO4217Code1CHE captures enum value "CHE"
	ISO4217Code1CHE ISO4217Code1 = "CHE"

	// ISO4217Code1CHF captures enum value "CHF"
	ISO4217Code1CHF ISO4217Code1 = "CHF"

	// ISO4217Code1CHW captures enum value "CHW"
	ISO4217Code1CHW ISO4217Code1 = "CHW"

	// ISO4217Code1CLF captures enum value "CLF"
	ISO4217Code1CLF ISO4217Code1 = "CLF"

	// ISO4217Code1CLP captures enum value "CLP"
	ISO4217Code1CLP ISO4217Code1 = "CLP"

	// ISO4217Code1CNY captures enum value "CNY"
	ISO4217Code1CNY ISO4217Code1 = "CNY"

	// ISO4217Code1COP captures enum value "COP"
	ISO4217Code1COP ISO4217Code1 = "COP"

	// ISO4217Code1COU captures enum value "COU"
	ISO4217Code1COU ISO4217Code1 = "COU"

	// ISO4217Code1CRC captures enum value "CRC"
	ISO4217Code1CRC ISO4217Code1 = "CRC"

	// ISO4217Code1CUC captures enum value "CUC"
	ISO4217Code1CUC ISO4217Code1 = "CUC"

	// ISO4217Code1CUP captures enum value "CUP"
	ISO4217Code1CUP ISO4217Code1 = "CUP"

	// ISO4217Code1CVE captures enum value "CVE"
	ISO4217Code1CVE ISO4217Code1 = "CVE"

	// ISO4217Code1CZK captures enum value "CZK"
	ISO4217Code1CZK ISO4217Code1 = "CZK"

	// ISO4217Code1DJF captures enum value "DJF"
	ISO4217Code1DJF ISO4217Code1 = "DJF"

	// ISO4217Code1DKK captures enum value "DKK"
	ISO4217Code1DKK ISO4217Code1 = "DKK"

	// ISO4217Code1DOP captures enum value "DOP"
	ISO4217Code1DOP ISO4217Code1 = "DOP"

	// ISO4217Code1DZD captures enum value "DZD"
	ISO4217Code1DZD ISO4217Code1 = "DZD"

	// ISO4217Code1EGP captures enum value "EGP"
	ISO4217Code1EGP ISO4217Code1 = "EGP"

	// ISO4217Code1ERN captures enum value "ERN"
	ISO4217Code1ERN ISO4217Code1 = "ERN"

	// ISO4217Code1ETB captures enum value "ETB"
	ISO4217Code1ETB ISO4217Code1 = "ETB"

	// ISO4217Code1EUR captures enum value "EUR"
	ISO4217Code1EUR ISO4217Code1 = "EUR"

	// ISO4217Code1FJD captures enum value "FJD"
	ISO4217Code1FJD ISO4217Code1 = "FJD"

	// ISO4217Code1FKP captures enum value "FKP"
	ISO4217Code1FKP ISO4217Code1 = "FKP"

	// ISO4217Code1GBP captures enum value "GBP"
	ISO4217Code1GBP ISO4217Code1 = "GBP"

	// ISO4217Code1GEL captures enum value "GEL"
	ISO4217Code1GEL ISO4217Code1 = "GEL"

	// ISO4217Code1GHS captures enum value "GHS"
	ISO4217Code1GHS ISO4217Code1 = "GHS"

	// ISO4217Code1GIP captures enum value "GIP"
	ISO4217Code1GIP ISO4217Code1 = "GIP"

	// ISO4217Code1GMD captures enum value "GMD"
	ISO4217Code1GMD ISO4217Code1 = "GMD"

	// ISO4217Code1GNF captures enum value "GNF"
	ISO4217Code1GNF ISO4217Code1 = "GNF"

	// ISO4217Code1GTQ captures enum value "GTQ"
	ISO4217Code1GTQ ISO4217Code1 = "GTQ"

	// ISO4217Code1GYD captures enum value "GYD"
	ISO4217Code1GYD ISO4217Code1 = "GYD"

	// ISO4217Code1HKD captures enum value "HKD"
	ISO4217Code1HKD ISO4217Code1 = "HKD"

	// ISO4217Code1HNL captures enum value "HNL"
	ISO4217Code1HNL ISO4217Code1 = "HNL"

	// ISO4217Code1HRK captures enum value "HRK"
	ISO4217Code1HRK ISO4217Code1 = "HRK"

	// ISO4217Code1HTG captures enum value "HTG"
	ISO4217Code1HTG ISO4217Code1 = "HTG"

	// ISO4217Code1HUF captures enum value "HUF"
	ISO4217Code1HUF ISO4217Code1 = "HUF"

	// ISO4217Code1IDR captures enum value "IDR"
	ISO4217Code1IDR ISO4217Code1 = "IDR"

	// ISO4217Code1ILS captures enum value "ILS"
	ISO4217Code1ILS ISO4217Code1 = "ILS"

	// ISO4217Code1INR captures enum value "INR"
	ISO4217Code1INR ISO4217Code1 = "INR"

	// ISO4217Code1IQD captures enum value "IQD"
	ISO4217Code1IQD ISO4217Code1 = "IQD"

	// ISO4217Code1IRR captures enum value "IRR"
	ISO4217Code1IRR ISO4217Code1 = "IRR"

	// ISO4217Code1ISK captures enum value "ISK"
	ISO4217Code1ISK ISO4217Code1 = "ISK"

	// ISO4217Code1JMD captures enum value "JMD"
	ISO4217Code1JMD ISO4217Code1 = "JMD"

	// ISO4217Code1JOD captures enum value "JOD"
	ISO4217Code1JOD ISO4217Code1 = "JOD"

	// ISO4217Code1JPY captures enum value "JPY"
	ISO4217Code1JPY ISO4217Code1 = "JPY"

	// ISO4217Code1KES captures enum value "KES"
	ISO4217Code1KES ISO4217Code1 = "KES"

	// ISO4217Code1KGS captures enum value "KGS"
	ISO4217Code1KGS ISO4217Code1 = "KGS"

	// ISO4217Code1KHR captures enum value "KHR"
	ISO4217Code1KHR ISO4217Code1 = "KHR"

	// ISO4217Code1KMF captures enum value "KMF"
	ISO4217Code1KMF ISO4217Code1 = "KMF"

	// ISO4217Code1KPW captures enum value "KPW"
	ISO4217Code1KPW ISO4217Code1 = "KPW"

	// ISO4217Code1KRW captures enum value "KRW"
	ISO4217Code1KRW ISO4217Code1 = "KRW"

	// ISO4217Code1KWD captures enum value "KWD"
	ISO4217Code1KWD ISO4217Code1 = "KWD"

	// ISO4217Code1KYD captures enum value "KYD"
	ISO4217Code1KYD ISO4217Code1 = "KYD"

	// ISO4217Code1KZT captures enum value "KZT"
	ISO4217Code1KZT ISO4217Code1 = "KZT"

	// ISO4217Code1LAK captures enum value "LAK"
	ISO4217Code1LAK ISO4217Code1 = "LAK"

	// ISO4217Code1LBP captures enum value "LBP"
	ISO4217Code1LBP ISO4217Code1 = "LBP"

	// ISO4217Code1LKR captures enum value "LKR"
	ISO4217Code1LKR ISO4217Code1 = "LKR"

	// ISO4217Code1LRD captures enum value "LRD"
	ISO4217Code1LRD ISO4217Code1 = "LRD"

	// ISO4217Code1LSL captures enum value "LSL"
	ISO4217Code1LSL ISO4217Code1 = "LSL"

	// ISO4217Code1LYD captures enum value "LYD"
	ISO4217Code1LYD ISO4217Code1 = "LYD"

	// ISO4217Code1MAD captures enum value "MAD"
	ISO4217Code1MAD ISO4217Code1 = "MAD"

	// ISO4217Code1MDL captures enum value "MDL"
	ISO4217Code1MDL ISO4217Code1 = "MDL"

	// ISO4217Code1MGA captures enum value "MGA"
	ISO4217Code1MGA ISO4217Code1 = "MGA"

	// ISO4217Code1MKD captures enum value "MKD"
	ISO4217Code1MKD ISO4217Code1 = "MKD"

	// ISO4217Code1MMK captures enum value "MMK"
	ISO4217Code1MMK ISO4217Code1 = "MMK"

	// ISO4217Code1MNT captures enum value "MNT"
	ISO4217Code1MNT ISO4217Code1 = "MNT"

	// ISO4217Code1MOP captures enum value "MOP"
	ISO4217Code1MOP ISO4217Code1 = "MOP"

	// ISO4217Code1MRO captures enum value "MRO"
	ISO4217Code1MRO ISO4217Code1 = "MRO"

	// ISO4217Code1MUR captures enum value "MUR"
	ISO4217Code1MUR ISO4217Code1 = "MUR"

	// ISO4217Code1MVR captures enum value "MVR"
	ISO4217Code1MVR ISO4217Code1 = "MVR"

	// ISO4217Code1MWK captures enum value "MWK"
	ISO4217Code1MWK ISO4217Code1 = "MWK"

	// ISO4217Code1MXN captures enum value "MXN"
	ISO4217Code1MXN ISO4217Code1 = "MXN"

	// ISO4217Code1MXV captures enum value "MXV"
	ISO4217Code1MXV ISO4217Code1 = "MXV"

	// ISO4217Code1MYR captures enum value "MYR"
	ISO4217Code1MYR ISO4217Code1 = "MYR"

	// ISO4217Code1MZN captures enum value "MZN"
	ISO4217Code1MZN ISO4217Code1 = "MZN"

	// ISO4217Code1NAD captures enum value "NAD"
	ISO4217Code1NAD ISO4217Code1 = "NAD"

	// ISO4217Code1NGN captures enum value "NGN"
	ISO4217Code1NGN ISO4217Code1 = "NGN"

	// ISO4217Code1NIO captures enum value "NIO"
	ISO4217Code1NIO ISO4217Code1 = "NIO"

	// ISO4217Code1NOK captures enum value "NOK"
	ISO4217Code1NOK ISO4217Code1 = "NOK"

	// ISO4217Code1NPR captures enum value "NPR"
	ISO4217Code1NPR ISO4217Code1 = "NPR"

	// ISO4217Code1NZD captures enum value "NZD"
	ISO4217Code1NZD ISO4217Code1 = "NZD"

	// ISO4217Code1OMR captures enum value "OMR"
	ISO4217Code1OMR ISO4217Code1 = "OMR"

	// ISO4217Code1PAB captures enum value "PAB"
	ISO4217Code1PAB ISO4217Code1 = "PAB"

	// ISO4217Code1PEN captures enum value "PEN"
	ISO4217Code1PEN ISO4217Code1 = "PEN"

	// ISO4217Code1PGK captures enum value "PGK"
	ISO4217Code1PGK ISO4217Code1 = "PGK"

	// ISO4217Code1PHP captures enum value "PHP"
	ISO4217Code1PHP ISO4217Code1 = "PHP"

	// ISO4217Code1PKR captures enum value "PKR"
	ISO4217Code1PKR ISO4217Code1 = "PKR"

	// ISO4217Code1PLN captures enum value "PLN"
	ISO4217Code1PLN ISO4217Code1 = "PLN"

	// ISO4217Code1PYG captures enum value "PYG"
	ISO4217Code1PYG ISO4217Code1 = "PYG"

	// ISO4217Code1QAR captures enum value "QAR"
	ISO4217Code1QAR ISO4217Code1 = "QAR"

	// ISO4217Code1RON captures enum value "RON"
	ISO4217Code1RON ISO4217Code1 = "RON"

	// ISO4217Code1RSD captures enum value "RSD"
	ISO4217Code1RSD ISO4217Code1 = "RSD"

	// ISO4217Code1RUB captures enum value "RUB"
	ISO4217Code1RUB ISO4217Code1 = "RUB"

	// ISO4217Code1RWF captures enum value "RWF"
	ISO4217Code1RWF ISO4217Code1 = "RWF"

	// ISO4217Code1SAR captures enum value "SAR"
	ISO4217Code1SAR ISO4217Code1 = "SAR"

	// ISO4217Code1SBD captures enum value "SBD"
	ISO4217Code1SBD ISO4217Code1 = "SBD"

	// ISO4217Code1SCR captures enum value "SCR"
	ISO4217Code1SCR ISO4217Code1 = "SCR"

	// ISO4217Code1SDG captures enum value "SDG"
	ISO4217Code1SDG ISO4217Code1 = "SDG"

	// ISO4217Code1SEK captures enum value "SEK"
	ISO4217Code1SEK ISO4217Code1 = "SEK"

	// ISO4217Code1SGD captures enum value "SGD"
	ISO4217Code1SGD ISO4217Code1 = "SGD"

	// ISO4217Code1SHP captures enum value "SHP"
	ISO4217Code1SHP ISO4217Code1 = "SHP"

	// ISO4217Code1SLL captures enum value "SLL"
	ISO4217Code1SLL ISO4217Code1 = "SLL"

	// ISO4217Code1SOS captures enum value "SOS"
	ISO4217Code1SOS ISO4217Code1 = "SOS"

	// ISO4217Code1SRD captures enum value "SRD"
	ISO4217Code1SRD ISO4217Code1 = "SRD"

	// ISO4217Code1SSP captures enum value "SSP"
	ISO4217Code1SSP ISO4217Code1 = "SSP"

	// ISO4217Code1STD captures enum value "STD"
	ISO4217Code1STD ISO4217Code1 = "STD"

	// ISO4217Code1SVC captures enum value "SVC"
	ISO4217Code1SVC ISO4217Code1 = "SVC"

	// ISO4217Code1SYP captures enum value "SYP"
	ISO4217Code1SYP ISO4217Code1 = "SYP"

	// ISO4217Code1SZL captures enum value "SZL"
	ISO4217Code1SZL ISO4217Code1 = "SZL"

	// ISO4217Code1THB captures enum value "THB"
	ISO4217Code1THB ISO4217Code1 = "THB"

	// ISO4217Code1TJS captures enum value "TJS"
	ISO4217Code1TJS ISO4217Code1 = "TJS"

	// ISO4217Code1TMT captures enum value "TMT"
	ISO4217Code1TMT ISO4217Code1 = "TMT"

	// ISO4217Code1TND captures enum value "TND"
	ISO4217Code1TND ISO4217Code1 = "TND"

	// ISO4217Code1TOP captures enum value "TOP"
	ISO4217Code1TOP ISO4217Code1 = "TOP"

	// ISO4217Code1TRY captures enum value "TRY"
	ISO4217Code1TRY ISO4217Code1 = "TRY"

	// ISO4217Code1TTD captures enum value "TTD"
	ISO4217Code1TTD ISO4217Code1 = "TTD"

	// ISO4217Code1TWD captures enum value "TWD"
	ISO4217Code1TWD ISO4217Code1 = "TWD"

	// ISO4217Code1TZS captures enum value "TZS"
	ISO4217Code1TZS ISO4217Code1 = "TZS"

	// ISO4217Code1UAH captures enum value "UAH"
	ISO4217Code1UAH ISO4217Code1 = "UAH"

	// ISO4217Code1UGX captures enum value "UGX"
	ISO4217Code1UGX ISO4217Code1 = "UGX"

	// ISO4217Code1USD captures enum value "USD"
	ISO4217Code1USD ISO4217Code1 = "USD"

	// ISO4217Code1USN captures enum value "USN"
	ISO4217Code1USN ISO4217Code1 = "USN"

	// ISO4217Code1UYI captures enum value "UYI"
	ISO4217Code1UYI ISO4217Code1 = "UYI"

	// ISO4217Code1UYU captures enum value "UYU"
	ISO4217Code1UYU ISO4217Code1 = "UYU"

	// ISO4217Code1UZS captures enum value "UZS"
	ISO4217Code1UZS ISO4217Code1 = "UZS"

	// ISO4217Code1VEF captures enum value "VEF"
	ISO4217Code1VEF ISO4217Code1 = "VEF"

	// ISO4217Code1VND captures enum value "VND"
	ISO4217Code1VND ISO4217Code1 = "VND"

	// ISO4217Code1VUV captures enum value "VUV"
	ISO4217Code1VUV ISO4217Code1 = "VUV"

	// ISO4217Code1WST captures enum value "WST"
	ISO4217Code1WST ISO4217Code1 = "WST"

	// ISO4217Code1XAF captures enum value "XAF"
	ISO4217Code1XAF ISO4217Code1 = "XAF"

	// ISO4217Code1XAG captures enum value "XAG"
	ISO4217Code1XAG ISO4217Code1 = "XAG"

	// ISO4217Code1XAU captures enum value "XAU"
	ISO4217Code1XAU ISO4217Code1 = "XAU"

	// ISO4217Code1XBA captures enum value "XBA"
	ISO4217Code1XBA ISO4217Code1 = "XBA"

	// ISO4217Code1XBB captures enum value "XBB"
	ISO4217Code1XBB ISO4217Code1 = "XBB"

	// ISO4217Code1XBC captures enum value "XBC"
	ISO4217Code1XBC ISO4217Code1 = "XBC"

	// ISO4217Code1XBD captures enum value "XBD"
	ISO4217Code1XBD ISO4217Code1 = "XBD"

	// ISO4217Code1XCD captures enum value "XCD"
	ISO4217Code1XCD ISO4217Code1 = "XCD"

	// ISO4217Code1XDR captures enum value "XDR"
	ISO4217Code1XDR ISO4217Code1 = "XDR"

	// ISO4217Code1XOF captures enum value "XOF"
	ISO4217Code1XOF ISO4217Code1 = "XOF"

	// ISO4217Code1XPD captures enum value "XPD"
	ISO4217Code1XPD ISO4217Code1 = "XPD"

	// ISO4217Code1XPF captures enum value "XPF"
	ISO4217Code1XPF ISO4217Code1 = "XPF"

	// ISO4217Code1XPT captures enum value "XPT"
	ISO4217Code1XPT ISO4217Code1 = "XPT"

	// ISO4217Code1XSU captures enum value "XSU"
	ISO4217Code1XSU ISO4217Code1 = "XSU"

	// ISO4217Code1XTS captures enum value "XTS"
	ISO4217Code1XTS ISO4217Code1 = "XTS"

	// ISO4217Code1XUA captures enum value "XUA"
	ISO4217Code1XUA ISO4217Code1 = "XUA"

	// ISO4217Code1XXX captures enum value "XXX"
	ISO4217Code1XXX ISO4217Code1 = "XXX"

	// ISO4217Code1YER captures enum value "YER"
	ISO4217Code1YER ISO4217Code1 = "YER"

	// ISO4217Code1ZAR captures enum value "ZAR"
	ISO4217Code1ZAR ISO4217Code1 = "ZAR"

	// ISO4217Code1ZMW captures enum value "ZMW"
	ISO4217Code1ZMW ISO4217Code1 = "ZMW"

	// ISO4217Code1ZWL captures enum value "ZWL"
	ISO4217Code1ZWL ISO4217Code1 = "ZWL"
)

// for schema
var iSO4217Code1Enum []interface{}

func init() {
	var res []ISO4217Code1
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BOV","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHE","CHF","CHW","CLF","CLP","CNY","COP","COU","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","INR","IQD","IRR","ISK","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MXV","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SRD","SSP","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","USD","USN","UYI","UYU","UZS","VEF","VND","VUV","WST","XAF","XAG","XAU","XBA","XBB","XBC","XBD","XCD","XDR","XOF","XPD","XPF","XPT","XSU","XTS","XUA","XXX","YER","ZAR","ZMW","ZWL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iSO4217Code1Enum = append(iSO4217Code1Enum, v)
	}
}

func (m ISO4217Code1) validateISO4217Code1Enum(path, location string, value ISO4217Code1) error {
	if err := validate.EnumCase(path, location, value, iSO4217Code1Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this i s o4217 code1
func (m ISO4217Code1) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateISO4217Code1Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this i s o4217 code1 based on context it is used
func (m ISO4217Code1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
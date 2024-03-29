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

// ISO4217Code2 ISO4217Code2
//
// Original ISO 4217 currency code
//
// swagger:model ISO4217Code2
type ISO4217Code2 string

func NewISO4217Code2(value ISO4217Code2) *ISO4217Code2 {
	v := value
	return &v
}

const (

	// ISO4217Code2AED captures enum value "AED"
	ISO4217Code2AED ISO4217Code2 = "AED"

	// ISO4217Code2AFN captures enum value "AFN"
	ISO4217Code2AFN ISO4217Code2 = "AFN"

	// ISO4217Code2ALL captures enum value "ALL"
	ISO4217Code2ALL ISO4217Code2 = "ALL"

	// ISO4217Code2AMD captures enum value "AMD"
	ISO4217Code2AMD ISO4217Code2 = "AMD"

	// ISO4217Code2ANG captures enum value "ANG"
	ISO4217Code2ANG ISO4217Code2 = "ANG"

	// ISO4217Code2AOA captures enum value "AOA"
	ISO4217Code2AOA ISO4217Code2 = "AOA"

	// ISO4217Code2ARS captures enum value "ARS"
	ISO4217Code2ARS ISO4217Code2 = "ARS"

	// ISO4217Code2AUD captures enum value "AUD"
	ISO4217Code2AUD ISO4217Code2 = "AUD"

	// ISO4217Code2AWG captures enum value "AWG"
	ISO4217Code2AWG ISO4217Code2 = "AWG"

	// ISO4217Code2AZN captures enum value "AZN"
	ISO4217Code2AZN ISO4217Code2 = "AZN"

	// ISO4217Code2BAM captures enum value "BAM"
	ISO4217Code2BAM ISO4217Code2 = "BAM"

	// ISO4217Code2BBD captures enum value "BBD"
	ISO4217Code2BBD ISO4217Code2 = "BBD"

	// ISO4217Code2BDT captures enum value "BDT"
	ISO4217Code2BDT ISO4217Code2 = "BDT"

	// ISO4217Code2BGN captures enum value "BGN"
	ISO4217Code2BGN ISO4217Code2 = "BGN"

	// ISO4217Code2BHD captures enum value "BHD"
	ISO4217Code2BHD ISO4217Code2 = "BHD"

	// ISO4217Code2BIF captures enum value "BIF"
	ISO4217Code2BIF ISO4217Code2 = "BIF"

	// ISO4217Code2BMD captures enum value "BMD"
	ISO4217Code2BMD ISO4217Code2 = "BMD"

	// ISO4217Code2BND captures enum value "BND"
	ISO4217Code2BND ISO4217Code2 = "BND"

	// ISO4217Code2BOB captures enum value "BOB"
	ISO4217Code2BOB ISO4217Code2 = "BOB"

	// ISO4217Code2BOV captures enum value "BOV"
	ISO4217Code2BOV ISO4217Code2 = "BOV"

	// ISO4217Code2BRL captures enum value "BRL"
	ISO4217Code2BRL ISO4217Code2 = "BRL"

	// ISO4217Code2BSD captures enum value "BSD"
	ISO4217Code2BSD ISO4217Code2 = "BSD"

	// ISO4217Code2BTN captures enum value "BTN"
	ISO4217Code2BTN ISO4217Code2 = "BTN"

	// ISO4217Code2BWP captures enum value "BWP"
	ISO4217Code2BWP ISO4217Code2 = "BWP"

	// ISO4217Code2BYR captures enum value "BYR"
	ISO4217Code2BYR ISO4217Code2 = "BYR"

	// ISO4217Code2BZD captures enum value "BZD"
	ISO4217Code2BZD ISO4217Code2 = "BZD"

	// ISO4217Code2CAD captures enum value "CAD"
	ISO4217Code2CAD ISO4217Code2 = "CAD"

	// ISO4217Code2CDF captures enum value "CDF"
	ISO4217Code2CDF ISO4217Code2 = "CDF"

	// ISO4217Code2CHE captures enum value "CHE"
	ISO4217Code2CHE ISO4217Code2 = "CHE"

	// ISO4217Code2CHF captures enum value "CHF"
	ISO4217Code2CHF ISO4217Code2 = "CHF"

	// ISO4217Code2CHW captures enum value "CHW"
	ISO4217Code2CHW ISO4217Code2 = "CHW"

	// ISO4217Code2CLF captures enum value "CLF"
	ISO4217Code2CLF ISO4217Code2 = "CLF"

	// ISO4217Code2CLP captures enum value "CLP"
	ISO4217Code2CLP ISO4217Code2 = "CLP"

	// ISO4217Code2CNY captures enum value "CNY"
	ISO4217Code2CNY ISO4217Code2 = "CNY"

	// ISO4217Code2COP captures enum value "COP"
	ISO4217Code2COP ISO4217Code2 = "COP"

	// ISO4217Code2COU captures enum value "COU"
	ISO4217Code2COU ISO4217Code2 = "COU"

	// ISO4217Code2CRC captures enum value "CRC"
	ISO4217Code2CRC ISO4217Code2 = "CRC"

	// ISO4217Code2CUC captures enum value "CUC"
	ISO4217Code2CUC ISO4217Code2 = "CUC"

	// ISO4217Code2CUP captures enum value "CUP"
	ISO4217Code2CUP ISO4217Code2 = "CUP"

	// ISO4217Code2CVE captures enum value "CVE"
	ISO4217Code2CVE ISO4217Code2 = "CVE"

	// ISO4217Code2CZK captures enum value "CZK"
	ISO4217Code2CZK ISO4217Code2 = "CZK"

	// ISO4217Code2DJF captures enum value "DJF"
	ISO4217Code2DJF ISO4217Code2 = "DJF"

	// ISO4217Code2DKK captures enum value "DKK"
	ISO4217Code2DKK ISO4217Code2 = "DKK"

	// ISO4217Code2DOP captures enum value "DOP"
	ISO4217Code2DOP ISO4217Code2 = "DOP"

	// ISO4217Code2DZD captures enum value "DZD"
	ISO4217Code2DZD ISO4217Code2 = "DZD"

	// ISO4217Code2EGP captures enum value "EGP"
	ISO4217Code2EGP ISO4217Code2 = "EGP"

	// ISO4217Code2ERN captures enum value "ERN"
	ISO4217Code2ERN ISO4217Code2 = "ERN"

	// ISO4217Code2ETB captures enum value "ETB"
	ISO4217Code2ETB ISO4217Code2 = "ETB"

	// ISO4217Code2EUR captures enum value "EUR"
	ISO4217Code2EUR ISO4217Code2 = "EUR"

	// ISO4217Code2FJD captures enum value "FJD"
	ISO4217Code2FJD ISO4217Code2 = "FJD"

	// ISO4217Code2FKP captures enum value "FKP"
	ISO4217Code2FKP ISO4217Code2 = "FKP"

	// ISO4217Code2GBP captures enum value "GBP"
	ISO4217Code2GBP ISO4217Code2 = "GBP"

	// ISO4217Code2GEL captures enum value "GEL"
	ISO4217Code2GEL ISO4217Code2 = "GEL"

	// ISO4217Code2GHS captures enum value "GHS"
	ISO4217Code2GHS ISO4217Code2 = "GHS"

	// ISO4217Code2GIP captures enum value "GIP"
	ISO4217Code2GIP ISO4217Code2 = "GIP"

	// ISO4217Code2GMD captures enum value "GMD"
	ISO4217Code2GMD ISO4217Code2 = "GMD"

	// ISO4217Code2GNF captures enum value "GNF"
	ISO4217Code2GNF ISO4217Code2 = "GNF"

	// ISO4217Code2GTQ captures enum value "GTQ"
	ISO4217Code2GTQ ISO4217Code2 = "GTQ"

	// ISO4217Code2GYD captures enum value "GYD"
	ISO4217Code2GYD ISO4217Code2 = "GYD"

	// ISO4217Code2HKD captures enum value "HKD"
	ISO4217Code2HKD ISO4217Code2 = "HKD"

	// ISO4217Code2HNL captures enum value "HNL"
	ISO4217Code2HNL ISO4217Code2 = "HNL"

	// ISO4217Code2HRK captures enum value "HRK"
	ISO4217Code2HRK ISO4217Code2 = "HRK"

	// ISO4217Code2HTG captures enum value "HTG"
	ISO4217Code2HTG ISO4217Code2 = "HTG"

	// ISO4217Code2HUF captures enum value "HUF"
	ISO4217Code2HUF ISO4217Code2 = "HUF"

	// ISO4217Code2IDR captures enum value "IDR"
	ISO4217Code2IDR ISO4217Code2 = "IDR"

	// ISO4217Code2ILS captures enum value "ILS"
	ISO4217Code2ILS ISO4217Code2 = "ILS"

	// ISO4217Code2INR captures enum value "INR"
	ISO4217Code2INR ISO4217Code2 = "INR"

	// ISO4217Code2IQD captures enum value "IQD"
	ISO4217Code2IQD ISO4217Code2 = "IQD"

	// ISO4217Code2IRR captures enum value "IRR"
	ISO4217Code2IRR ISO4217Code2 = "IRR"

	// ISO4217Code2ISK captures enum value "ISK"
	ISO4217Code2ISK ISO4217Code2 = "ISK"

	// ISO4217Code2JMD captures enum value "JMD"
	ISO4217Code2JMD ISO4217Code2 = "JMD"

	// ISO4217Code2JOD captures enum value "JOD"
	ISO4217Code2JOD ISO4217Code2 = "JOD"

	// ISO4217Code2JPY captures enum value "JPY"
	ISO4217Code2JPY ISO4217Code2 = "JPY"

	// ISO4217Code2KES captures enum value "KES"
	ISO4217Code2KES ISO4217Code2 = "KES"

	// ISO4217Code2KGS captures enum value "KGS"
	ISO4217Code2KGS ISO4217Code2 = "KGS"

	// ISO4217Code2KHR captures enum value "KHR"
	ISO4217Code2KHR ISO4217Code2 = "KHR"

	// ISO4217Code2KMF captures enum value "KMF"
	ISO4217Code2KMF ISO4217Code2 = "KMF"

	// ISO4217Code2KPW captures enum value "KPW"
	ISO4217Code2KPW ISO4217Code2 = "KPW"

	// ISO4217Code2KRW captures enum value "KRW"
	ISO4217Code2KRW ISO4217Code2 = "KRW"

	// ISO4217Code2KWD captures enum value "KWD"
	ISO4217Code2KWD ISO4217Code2 = "KWD"

	// ISO4217Code2KYD captures enum value "KYD"
	ISO4217Code2KYD ISO4217Code2 = "KYD"

	// ISO4217Code2KZT captures enum value "KZT"
	ISO4217Code2KZT ISO4217Code2 = "KZT"

	// ISO4217Code2LAK captures enum value "LAK"
	ISO4217Code2LAK ISO4217Code2 = "LAK"

	// ISO4217Code2LBP captures enum value "LBP"
	ISO4217Code2LBP ISO4217Code2 = "LBP"

	// ISO4217Code2LKR captures enum value "LKR"
	ISO4217Code2LKR ISO4217Code2 = "LKR"

	// ISO4217Code2LRD captures enum value "LRD"
	ISO4217Code2LRD ISO4217Code2 = "LRD"

	// ISO4217Code2LSL captures enum value "LSL"
	ISO4217Code2LSL ISO4217Code2 = "LSL"

	// ISO4217Code2LYD captures enum value "LYD"
	ISO4217Code2LYD ISO4217Code2 = "LYD"

	// ISO4217Code2MAD captures enum value "MAD"
	ISO4217Code2MAD ISO4217Code2 = "MAD"

	// ISO4217Code2MDL captures enum value "MDL"
	ISO4217Code2MDL ISO4217Code2 = "MDL"

	// ISO4217Code2MGA captures enum value "MGA"
	ISO4217Code2MGA ISO4217Code2 = "MGA"

	// ISO4217Code2MKD captures enum value "MKD"
	ISO4217Code2MKD ISO4217Code2 = "MKD"

	// ISO4217Code2MMK captures enum value "MMK"
	ISO4217Code2MMK ISO4217Code2 = "MMK"

	// ISO4217Code2MNT captures enum value "MNT"
	ISO4217Code2MNT ISO4217Code2 = "MNT"

	// ISO4217Code2MOP captures enum value "MOP"
	ISO4217Code2MOP ISO4217Code2 = "MOP"

	// ISO4217Code2MRO captures enum value "MRO"
	ISO4217Code2MRO ISO4217Code2 = "MRO"

	// ISO4217Code2MUR captures enum value "MUR"
	ISO4217Code2MUR ISO4217Code2 = "MUR"

	// ISO4217Code2MVR captures enum value "MVR"
	ISO4217Code2MVR ISO4217Code2 = "MVR"

	// ISO4217Code2MWK captures enum value "MWK"
	ISO4217Code2MWK ISO4217Code2 = "MWK"

	// ISO4217Code2MXN captures enum value "MXN"
	ISO4217Code2MXN ISO4217Code2 = "MXN"

	// ISO4217Code2MXV captures enum value "MXV"
	ISO4217Code2MXV ISO4217Code2 = "MXV"

	// ISO4217Code2MYR captures enum value "MYR"
	ISO4217Code2MYR ISO4217Code2 = "MYR"

	// ISO4217Code2MZN captures enum value "MZN"
	ISO4217Code2MZN ISO4217Code2 = "MZN"

	// ISO4217Code2NAD captures enum value "NAD"
	ISO4217Code2NAD ISO4217Code2 = "NAD"

	// ISO4217Code2NGN captures enum value "NGN"
	ISO4217Code2NGN ISO4217Code2 = "NGN"

	// ISO4217Code2NIO captures enum value "NIO"
	ISO4217Code2NIO ISO4217Code2 = "NIO"

	// ISO4217Code2NOK captures enum value "NOK"
	ISO4217Code2NOK ISO4217Code2 = "NOK"

	// ISO4217Code2NPR captures enum value "NPR"
	ISO4217Code2NPR ISO4217Code2 = "NPR"

	// ISO4217Code2NZD captures enum value "NZD"
	ISO4217Code2NZD ISO4217Code2 = "NZD"

	// ISO4217Code2OMR captures enum value "OMR"
	ISO4217Code2OMR ISO4217Code2 = "OMR"

	// ISO4217Code2PAB captures enum value "PAB"
	ISO4217Code2PAB ISO4217Code2 = "PAB"

	// ISO4217Code2PEN captures enum value "PEN"
	ISO4217Code2PEN ISO4217Code2 = "PEN"

	// ISO4217Code2PGK captures enum value "PGK"
	ISO4217Code2PGK ISO4217Code2 = "PGK"

	// ISO4217Code2PHP captures enum value "PHP"
	ISO4217Code2PHP ISO4217Code2 = "PHP"

	// ISO4217Code2PKR captures enum value "PKR"
	ISO4217Code2PKR ISO4217Code2 = "PKR"

	// ISO4217Code2PLN captures enum value "PLN"
	ISO4217Code2PLN ISO4217Code2 = "PLN"

	// ISO4217Code2PYG captures enum value "PYG"
	ISO4217Code2PYG ISO4217Code2 = "PYG"

	// ISO4217Code2QAR captures enum value "QAR"
	ISO4217Code2QAR ISO4217Code2 = "QAR"

	// ISO4217Code2RON captures enum value "RON"
	ISO4217Code2RON ISO4217Code2 = "RON"

	// ISO4217Code2RSD captures enum value "RSD"
	ISO4217Code2RSD ISO4217Code2 = "RSD"

	// ISO4217Code2RUB captures enum value "RUB"
	ISO4217Code2RUB ISO4217Code2 = "RUB"

	// ISO4217Code2RWF captures enum value "RWF"
	ISO4217Code2RWF ISO4217Code2 = "RWF"

	// ISO4217Code2SAR captures enum value "SAR"
	ISO4217Code2SAR ISO4217Code2 = "SAR"

	// ISO4217Code2SBD captures enum value "SBD"
	ISO4217Code2SBD ISO4217Code2 = "SBD"

	// ISO4217Code2SCR captures enum value "SCR"
	ISO4217Code2SCR ISO4217Code2 = "SCR"

	// ISO4217Code2SDG captures enum value "SDG"
	ISO4217Code2SDG ISO4217Code2 = "SDG"

	// ISO4217Code2SEK captures enum value "SEK"
	ISO4217Code2SEK ISO4217Code2 = "SEK"

	// ISO4217Code2SGD captures enum value "SGD"
	ISO4217Code2SGD ISO4217Code2 = "SGD"

	// ISO4217Code2SHP captures enum value "SHP"
	ISO4217Code2SHP ISO4217Code2 = "SHP"

	// ISO4217Code2SLL captures enum value "SLL"
	ISO4217Code2SLL ISO4217Code2 = "SLL"

	// ISO4217Code2SOS captures enum value "SOS"
	ISO4217Code2SOS ISO4217Code2 = "SOS"

	// ISO4217Code2SRD captures enum value "SRD"
	ISO4217Code2SRD ISO4217Code2 = "SRD"

	// ISO4217Code2SSP captures enum value "SSP"
	ISO4217Code2SSP ISO4217Code2 = "SSP"

	// ISO4217Code2STD captures enum value "STD"
	ISO4217Code2STD ISO4217Code2 = "STD"

	// ISO4217Code2SVC captures enum value "SVC"
	ISO4217Code2SVC ISO4217Code2 = "SVC"

	// ISO4217Code2SYP captures enum value "SYP"
	ISO4217Code2SYP ISO4217Code2 = "SYP"

	// ISO4217Code2SZL captures enum value "SZL"
	ISO4217Code2SZL ISO4217Code2 = "SZL"

	// ISO4217Code2THB captures enum value "THB"
	ISO4217Code2THB ISO4217Code2 = "THB"

	// ISO4217Code2TJS captures enum value "TJS"
	ISO4217Code2TJS ISO4217Code2 = "TJS"

	// ISO4217Code2TMT captures enum value "TMT"
	ISO4217Code2TMT ISO4217Code2 = "TMT"

	// ISO4217Code2TND captures enum value "TND"
	ISO4217Code2TND ISO4217Code2 = "TND"

	// ISO4217Code2TOP captures enum value "TOP"
	ISO4217Code2TOP ISO4217Code2 = "TOP"

	// ISO4217Code2TRY captures enum value "TRY"
	ISO4217Code2TRY ISO4217Code2 = "TRY"

	// ISO4217Code2TTD captures enum value "TTD"
	ISO4217Code2TTD ISO4217Code2 = "TTD"

	// ISO4217Code2TWD captures enum value "TWD"
	ISO4217Code2TWD ISO4217Code2 = "TWD"

	// ISO4217Code2TZS captures enum value "TZS"
	ISO4217Code2TZS ISO4217Code2 = "TZS"

	// ISO4217Code2UAH captures enum value "UAH"
	ISO4217Code2UAH ISO4217Code2 = "UAH"

	// ISO4217Code2UGX captures enum value "UGX"
	ISO4217Code2UGX ISO4217Code2 = "UGX"

	// ISO4217Code2USD captures enum value "USD"
	ISO4217Code2USD ISO4217Code2 = "USD"

	// ISO4217Code2USN captures enum value "USN"
	ISO4217Code2USN ISO4217Code2 = "USN"

	// ISO4217Code2UYI captures enum value "UYI"
	ISO4217Code2UYI ISO4217Code2 = "UYI"

	// ISO4217Code2UYU captures enum value "UYU"
	ISO4217Code2UYU ISO4217Code2 = "UYU"

	// ISO4217Code2UZS captures enum value "UZS"
	ISO4217Code2UZS ISO4217Code2 = "UZS"

	// ISO4217Code2VEF captures enum value "VEF"
	ISO4217Code2VEF ISO4217Code2 = "VEF"

	// ISO4217Code2VND captures enum value "VND"
	ISO4217Code2VND ISO4217Code2 = "VND"

	// ISO4217Code2VUV captures enum value "VUV"
	ISO4217Code2VUV ISO4217Code2 = "VUV"

	// ISO4217Code2WST captures enum value "WST"
	ISO4217Code2WST ISO4217Code2 = "WST"

	// ISO4217Code2XAF captures enum value "XAF"
	ISO4217Code2XAF ISO4217Code2 = "XAF"

	// ISO4217Code2XAG captures enum value "XAG"
	ISO4217Code2XAG ISO4217Code2 = "XAG"

	// ISO4217Code2XAU captures enum value "XAU"
	ISO4217Code2XAU ISO4217Code2 = "XAU"

	// ISO4217Code2XBA captures enum value "XBA"
	ISO4217Code2XBA ISO4217Code2 = "XBA"

	// ISO4217Code2XBB captures enum value "XBB"
	ISO4217Code2XBB ISO4217Code2 = "XBB"

	// ISO4217Code2XBC captures enum value "XBC"
	ISO4217Code2XBC ISO4217Code2 = "XBC"

	// ISO4217Code2XBD captures enum value "XBD"
	ISO4217Code2XBD ISO4217Code2 = "XBD"

	// ISO4217Code2XCD captures enum value "XCD"
	ISO4217Code2XCD ISO4217Code2 = "XCD"

	// ISO4217Code2XDR captures enum value "XDR"
	ISO4217Code2XDR ISO4217Code2 = "XDR"

	// ISO4217Code2XOF captures enum value "XOF"
	ISO4217Code2XOF ISO4217Code2 = "XOF"

	// ISO4217Code2XPD captures enum value "XPD"
	ISO4217Code2XPD ISO4217Code2 = "XPD"

	// ISO4217Code2XPF captures enum value "XPF"
	ISO4217Code2XPF ISO4217Code2 = "XPF"

	// ISO4217Code2XPT captures enum value "XPT"
	ISO4217Code2XPT ISO4217Code2 = "XPT"

	// ISO4217Code2XSU captures enum value "XSU"
	ISO4217Code2XSU ISO4217Code2 = "XSU"

	// ISO4217Code2XTS captures enum value "XTS"
	ISO4217Code2XTS ISO4217Code2 = "XTS"

	// ISO4217Code2XUA captures enum value "XUA"
	ISO4217Code2XUA ISO4217Code2 = "XUA"

	// ISO4217Code2XXX captures enum value "XXX"
	ISO4217Code2XXX ISO4217Code2 = "XXX"

	// ISO4217Code2YER captures enum value "YER"
	ISO4217Code2YER ISO4217Code2 = "YER"

	// ISO4217Code2ZAR captures enum value "ZAR"
	ISO4217Code2ZAR ISO4217Code2 = "ZAR"

	// ISO4217Code2ZMW captures enum value "ZMW"
	ISO4217Code2ZMW ISO4217Code2 = "ZMW"

	// ISO4217Code2ZWL captures enum value "ZWL"
	ISO4217Code2ZWL ISO4217Code2 = "ZWL"
)

// for schema
var iSO4217Code2Enum []interface{}

func init() {
	var res []ISO4217Code2
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BOV","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHE","CHF","CHW","CLF","CLP","CNY","COP","COU","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","INR","IQD","IRR","ISK","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MXV","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SRD","SSP","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","USD","USN","UYI","UYU","UZS","VEF","VND","VUV","WST","XAF","XAG","XAU","XBA","XBB","XBC","XBD","XCD","XDR","XOF","XPD","XPF","XPT","XSU","XTS","XUA","XXX","YER","ZAR","ZMW","ZWL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iSO4217Code2Enum = append(iSO4217Code2Enum, v)
	}
}

func (m ISO4217Code2) validateISO4217Code2Enum(path, location string, value ISO4217Code2) error {
	if err := validate.EnumCase(path, location, value, iSO4217Code2Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this i s o4217 code2
func (m ISO4217Code2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateISO4217Code2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this i s o4217 code2 based on context it is used
func (m ISO4217Code2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

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

// ISO3166CountryCode ISO3166CountryCode
//
// ISO 3166 Codes for the representation of names of countries and their subdivisions
//
// swagger:model ISO3166CountryCode
type ISO3166CountryCode string

func NewISO3166CountryCode(value ISO3166CountryCode) *ISO3166CountryCode {
	v := value
	return &v
}

const (

	// ISO3166CountryCodeAD captures enum value "AD"
	ISO3166CountryCodeAD ISO3166CountryCode = "AD"

	// ISO3166CountryCodeAE captures enum value "AE"
	ISO3166CountryCodeAE ISO3166CountryCode = "AE"

	// ISO3166CountryCodeAF captures enum value "AF"
	ISO3166CountryCodeAF ISO3166CountryCode = "AF"

	// ISO3166CountryCodeAG captures enum value "AG"
	ISO3166CountryCodeAG ISO3166CountryCode = "AG"

	// ISO3166CountryCodeAI captures enum value "AI"
	ISO3166CountryCodeAI ISO3166CountryCode = "AI"

	// ISO3166CountryCodeAL captures enum value "AL"
	ISO3166CountryCodeAL ISO3166CountryCode = "AL"

	// ISO3166CountryCodeAM captures enum value "AM"
	ISO3166CountryCodeAM ISO3166CountryCode = "AM"

	// ISO3166CountryCodeAN captures enum value "AN"
	ISO3166CountryCodeAN ISO3166CountryCode = "AN"

	// ISO3166CountryCodeAO captures enum value "AO"
	ISO3166CountryCodeAO ISO3166CountryCode = "AO"

	// ISO3166CountryCodeAQ captures enum value "AQ"
	ISO3166CountryCodeAQ ISO3166CountryCode = "AQ"

	// ISO3166CountryCodeAR captures enum value "AR"
	ISO3166CountryCodeAR ISO3166CountryCode = "AR"

	// ISO3166CountryCodeAS captures enum value "AS"
	ISO3166CountryCodeAS ISO3166CountryCode = "AS"

	// ISO3166CountryCodeAT captures enum value "AT"
	ISO3166CountryCodeAT ISO3166CountryCode = "AT"

	// ISO3166CountryCodeAU captures enum value "AU"
	ISO3166CountryCodeAU ISO3166CountryCode = "AU"

	// ISO3166CountryCodeAW captures enum value "AW"
	ISO3166CountryCodeAW ISO3166CountryCode = "AW"

	// ISO3166CountryCodeAX captures enum value "AX"
	ISO3166CountryCodeAX ISO3166CountryCode = "AX"

	// ISO3166CountryCodeAZ captures enum value "AZ"
	ISO3166CountryCodeAZ ISO3166CountryCode = "AZ"

	// ISO3166CountryCodeBA captures enum value "BA"
	ISO3166CountryCodeBA ISO3166CountryCode = "BA"

	// ISO3166CountryCodeBB captures enum value "BB"
	ISO3166CountryCodeBB ISO3166CountryCode = "BB"

	// ISO3166CountryCodeBD captures enum value "BD"
	ISO3166CountryCodeBD ISO3166CountryCode = "BD"

	// ISO3166CountryCodeBE captures enum value "BE"
	ISO3166CountryCodeBE ISO3166CountryCode = "BE"

	// ISO3166CountryCodeBF captures enum value "BF"
	ISO3166CountryCodeBF ISO3166CountryCode = "BF"

	// ISO3166CountryCodeBG captures enum value "BG"
	ISO3166CountryCodeBG ISO3166CountryCode = "BG"

	// ISO3166CountryCodeBH captures enum value "BH"
	ISO3166CountryCodeBH ISO3166CountryCode = "BH"

	// ISO3166CountryCodeBI captures enum value "BI"
	ISO3166CountryCodeBI ISO3166CountryCode = "BI"

	// ISO3166CountryCodeBJ captures enum value "BJ"
	ISO3166CountryCodeBJ ISO3166CountryCode = "BJ"

	// ISO3166CountryCodeBM captures enum value "BM"
	ISO3166CountryCodeBM ISO3166CountryCode = "BM"

	// ISO3166CountryCodeBN captures enum value "BN"
	ISO3166CountryCodeBN ISO3166CountryCode = "BN"

	// ISO3166CountryCodeBO captures enum value "BO"
	ISO3166CountryCodeBO ISO3166CountryCode = "BO"

	// ISO3166CountryCodeBR captures enum value "BR"
	ISO3166CountryCodeBR ISO3166CountryCode = "BR"

	// ISO3166CountryCodeBS captures enum value "BS"
	ISO3166CountryCodeBS ISO3166CountryCode = "BS"

	// ISO3166CountryCodeBT captures enum value "BT"
	ISO3166CountryCodeBT ISO3166CountryCode = "BT"

	// ISO3166CountryCodeBV captures enum value "BV"
	ISO3166CountryCodeBV ISO3166CountryCode = "BV"

	// ISO3166CountryCodeBW captures enum value "BW"
	ISO3166CountryCodeBW ISO3166CountryCode = "BW"

	// ISO3166CountryCodeBY captures enum value "BY"
	ISO3166CountryCodeBY ISO3166CountryCode = "BY"

	// ISO3166CountryCodeBZ captures enum value "BZ"
	ISO3166CountryCodeBZ ISO3166CountryCode = "BZ"

	// ISO3166CountryCodeCA captures enum value "CA"
	ISO3166CountryCodeCA ISO3166CountryCode = "CA"

	// ISO3166CountryCodeCC captures enum value "CC"
	ISO3166CountryCodeCC ISO3166CountryCode = "CC"

	// ISO3166CountryCodeCD captures enum value "CD"
	ISO3166CountryCodeCD ISO3166CountryCode = "CD"

	// ISO3166CountryCodeCF captures enum value "CF"
	ISO3166CountryCodeCF ISO3166CountryCode = "CF"

	// ISO3166CountryCodeCG captures enum value "CG"
	ISO3166CountryCodeCG ISO3166CountryCode = "CG"

	// ISO3166CountryCodeCH captures enum value "CH"
	ISO3166CountryCodeCH ISO3166CountryCode = "CH"

	// ISO3166CountryCodeCI captures enum value "CI"
	ISO3166CountryCodeCI ISO3166CountryCode = "CI"

	// ISO3166CountryCodeCK captures enum value "CK"
	ISO3166CountryCodeCK ISO3166CountryCode = "CK"

	// ISO3166CountryCodeCL captures enum value "CL"
	ISO3166CountryCodeCL ISO3166CountryCode = "CL"

	// ISO3166CountryCodeCM captures enum value "CM"
	ISO3166CountryCodeCM ISO3166CountryCode = "CM"

	// ISO3166CountryCodeCN captures enum value "CN"
	ISO3166CountryCodeCN ISO3166CountryCode = "CN"

	// ISO3166CountryCodeCO captures enum value "CO"
	ISO3166CountryCodeCO ISO3166CountryCode = "CO"

	// ISO3166CountryCodeCR captures enum value "CR"
	ISO3166CountryCodeCR ISO3166CountryCode = "CR"

	// ISO3166CountryCodeCS captures enum value "CS"
	ISO3166CountryCodeCS ISO3166CountryCode = "CS"

	// ISO3166CountryCodeCU captures enum value "CU"
	ISO3166CountryCodeCU ISO3166CountryCode = "CU"

	// ISO3166CountryCodeCV captures enum value "CV"
	ISO3166CountryCodeCV ISO3166CountryCode = "CV"

	// ISO3166CountryCodeCX captures enum value "CX"
	ISO3166CountryCodeCX ISO3166CountryCode = "CX"

	// ISO3166CountryCodeCY captures enum value "CY"
	ISO3166CountryCodeCY ISO3166CountryCode = "CY"

	// ISO3166CountryCodeCZ captures enum value "CZ"
	ISO3166CountryCodeCZ ISO3166CountryCode = "CZ"

	// ISO3166CountryCodeDE captures enum value "DE"
	ISO3166CountryCodeDE ISO3166CountryCode = "DE"

	// ISO3166CountryCodeDJ captures enum value "DJ"
	ISO3166CountryCodeDJ ISO3166CountryCode = "DJ"

	// ISO3166CountryCodeDK captures enum value "DK"
	ISO3166CountryCodeDK ISO3166CountryCode = "DK"

	// ISO3166CountryCodeDM captures enum value "DM"
	ISO3166CountryCodeDM ISO3166CountryCode = "DM"

	// ISO3166CountryCodeDO captures enum value "DO"
	ISO3166CountryCodeDO ISO3166CountryCode = "DO"

	// ISO3166CountryCodeDZ captures enum value "DZ"
	ISO3166CountryCodeDZ ISO3166CountryCode = "DZ"

	// ISO3166CountryCodeEC captures enum value "EC"
	ISO3166CountryCodeEC ISO3166CountryCode = "EC"

	// ISO3166CountryCodeEE captures enum value "EE"
	ISO3166CountryCodeEE ISO3166CountryCode = "EE"

	// ISO3166CountryCodeEG captures enum value "EG"
	ISO3166CountryCodeEG ISO3166CountryCode = "EG"

	// ISO3166CountryCodeEH captures enum value "EH"
	ISO3166CountryCodeEH ISO3166CountryCode = "EH"

	// ISO3166CountryCodeER captures enum value "ER"
	ISO3166CountryCodeER ISO3166CountryCode = "ER"

	// ISO3166CountryCodeES captures enum value "ES"
	ISO3166CountryCodeES ISO3166CountryCode = "ES"

	// ISO3166CountryCodeET captures enum value "ET"
	ISO3166CountryCodeET ISO3166CountryCode = "ET"

	// ISO3166CountryCodeFI captures enum value "FI"
	ISO3166CountryCodeFI ISO3166CountryCode = "FI"

	// ISO3166CountryCodeFJ captures enum value "FJ"
	ISO3166CountryCodeFJ ISO3166CountryCode = "FJ"

	// ISO3166CountryCodeFK captures enum value "FK"
	ISO3166CountryCodeFK ISO3166CountryCode = "FK"

	// ISO3166CountryCodeFM captures enum value "FM"
	ISO3166CountryCodeFM ISO3166CountryCode = "FM"

	// ISO3166CountryCodeFO captures enum value "FO"
	ISO3166CountryCodeFO ISO3166CountryCode = "FO"

	// ISO3166CountryCodeFR captures enum value "FR"
	ISO3166CountryCodeFR ISO3166CountryCode = "FR"

	// ISO3166CountryCodeGA captures enum value "GA"
	ISO3166CountryCodeGA ISO3166CountryCode = "GA"

	// ISO3166CountryCodeGB captures enum value "GB"
	ISO3166CountryCodeGB ISO3166CountryCode = "GB"

	// ISO3166CountryCodeGD captures enum value "GD"
	ISO3166CountryCodeGD ISO3166CountryCode = "GD"

	// ISO3166CountryCodeGE captures enum value "GE"
	ISO3166CountryCodeGE ISO3166CountryCode = "GE"

	// ISO3166CountryCodeGF captures enum value "GF"
	ISO3166CountryCodeGF ISO3166CountryCode = "GF"

	// ISO3166CountryCodeGG captures enum value "GG"
	ISO3166CountryCodeGG ISO3166CountryCode = "GG"

	// ISO3166CountryCodeGH captures enum value "GH"
	ISO3166CountryCodeGH ISO3166CountryCode = "GH"

	// ISO3166CountryCodeGI captures enum value "GI"
	ISO3166CountryCodeGI ISO3166CountryCode = "GI"

	// ISO3166CountryCodeGL captures enum value "GL"
	ISO3166CountryCodeGL ISO3166CountryCode = "GL"

	// ISO3166CountryCodeGM captures enum value "GM"
	ISO3166CountryCodeGM ISO3166CountryCode = "GM"

	// ISO3166CountryCodeGN captures enum value "GN"
	ISO3166CountryCodeGN ISO3166CountryCode = "GN"

	// ISO3166CountryCodeGP captures enum value "GP"
	ISO3166CountryCodeGP ISO3166CountryCode = "GP"

	// ISO3166CountryCodeGQ captures enum value "GQ"
	ISO3166CountryCodeGQ ISO3166CountryCode = "GQ"

	// ISO3166CountryCodeGR captures enum value "GR"
	ISO3166CountryCodeGR ISO3166CountryCode = "GR"

	// ISO3166CountryCodeGS captures enum value "GS"
	ISO3166CountryCodeGS ISO3166CountryCode = "GS"

	// ISO3166CountryCodeGT captures enum value "GT"
	ISO3166CountryCodeGT ISO3166CountryCode = "GT"

	// ISO3166CountryCodeGU captures enum value "GU"
	ISO3166CountryCodeGU ISO3166CountryCode = "GU"

	// ISO3166CountryCodeGW captures enum value "GW"
	ISO3166CountryCodeGW ISO3166CountryCode = "GW"

	// ISO3166CountryCodeGY captures enum value "GY"
	ISO3166CountryCodeGY ISO3166CountryCode = "GY"

	// ISO3166CountryCodeHK captures enum value "HK"
	ISO3166CountryCodeHK ISO3166CountryCode = "HK"

	// ISO3166CountryCodeHM captures enum value "HM"
	ISO3166CountryCodeHM ISO3166CountryCode = "HM"

	// ISO3166CountryCodeHN captures enum value "HN"
	ISO3166CountryCodeHN ISO3166CountryCode = "HN"

	// ISO3166CountryCodeHR captures enum value "HR"
	ISO3166CountryCodeHR ISO3166CountryCode = "HR"

	// ISO3166CountryCodeHT captures enum value "HT"
	ISO3166CountryCodeHT ISO3166CountryCode = "HT"

	// ISO3166CountryCodeHU captures enum value "HU"
	ISO3166CountryCodeHU ISO3166CountryCode = "HU"

	// ISO3166CountryCodeID captures enum value "ID"
	ISO3166CountryCodeID ISO3166CountryCode = "ID"

	// ISO3166CountryCodeIE captures enum value "IE"
	ISO3166CountryCodeIE ISO3166CountryCode = "IE"

	// ISO3166CountryCodeIL captures enum value "IL"
	ISO3166CountryCodeIL ISO3166CountryCode = "IL"

	// ISO3166CountryCodeIM captures enum value "IM"
	ISO3166CountryCodeIM ISO3166CountryCode = "IM"

	// ISO3166CountryCodeIN captures enum value "IN"
	ISO3166CountryCodeIN ISO3166CountryCode = "IN"

	// ISO3166CountryCodeIO captures enum value "IO"
	ISO3166CountryCodeIO ISO3166CountryCode = "IO"

	// ISO3166CountryCodeIQ captures enum value "IQ"
	ISO3166CountryCodeIQ ISO3166CountryCode = "IQ"

	// ISO3166CountryCodeIR captures enum value "IR"
	ISO3166CountryCodeIR ISO3166CountryCode = "IR"

	// ISO3166CountryCodeIS captures enum value "IS"
	ISO3166CountryCodeIS ISO3166CountryCode = "IS"

	// ISO3166CountryCodeIT captures enum value "IT"
	ISO3166CountryCodeIT ISO3166CountryCode = "IT"

	// ISO3166CountryCodeJE captures enum value "JE"
	ISO3166CountryCodeJE ISO3166CountryCode = "JE"

	// ISO3166CountryCodeJM captures enum value "JM"
	ISO3166CountryCodeJM ISO3166CountryCode = "JM"

	// ISO3166CountryCodeJO captures enum value "JO"
	ISO3166CountryCodeJO ISO3166CountryCode = "JO"

	// ISO3166CountryCodeJP captures enum value "JP"
	ISO3166CountryCodeJP ISO3166CountryCode = "JP"

	// ISO3166CountryCodeKE captures enum value "KE"
	ISO3166CountryCodeKE ISO3166CountryCode = "KE"

	// ISO3166CountryCodeKG captures enum value "KG"
	ISO3166CountryCodeKG ISO3166CountryCode = "KG"

	// ISO3166CountryCodeKH captures enum value "KH"
	ISO3166CountryCodeKH ISO3166CountryCode = "KH"

	// ISO3166CountryCodeKI captures enum value "KI"
	ISO3166CountryCodeKI ISO3166CountryCode = "KI"

	// ISO3166CountryCodeKM captures enum value "KM"
	ISO3166CountryCodeKM ISO3166CountryCode = "KM"

	// ISO3166CountryCodeKN captures enum value "KN"
	ISO3166CountryCodeKN ISO3166CountryCode = "KN"

	// ISO3166CountryCodeKP captures enum value "KP"
	ISO3166CountryCodeKP ISO3166CountryCode = "KP"

	// ISO3166CountryCodeKR captures enum value "KR"
	ISO3166CountryCodeKR ISO3166CountryCode = "KR"

	// ISO3166CountryCodeKW captures enum value "KW"
	ISO3166CountryCodeKW ISO3166CountryCode = "KW"

	// ISO3166CountryCodeKY captures enum value "KY"
	ISO3166CountryCodeKY ISO3166CountryCode = "KY"

	// ISO3166CountryCodeKZ captures enum value "KZ"
	ISO3166CountryCodeKZ ISO3166CountryCode = "KZ"

	// ISO3166CountryCodeLA captures enum value "LA"
	ISO3166CountryCodeLA ISO3166CountryCode = "LA"

	// ISO3166CountryCodeLB captures enum value "LB"
	ISO3166CountryCodeLB ISO3166CountryCode = "LB"

	// ISO3166CountryCodeLC captures enum value "LC"
	ISO3166CountryCodeLC ISO3166CountryCode = "LC"

	// ISO3166CountryCodeLI captures enum value "LI"
	ISO3166CountryCodeLI ISO3166CountryCode = "LI"

	// ISO3166CountryCodeLK captures enum value "LK"
	ISO3166CountryCodeLK ISO3166CountryCode = "LK"

	// ISO3166CountryCodeLR captures enum value "LR"
	ISO3166CountryCodeLR ISO3166CountryCode = "LR"

	// ISO3166CountryCodeLS captures enum value "LS"
	ISO3166CountryCodeLS ISO3166CountryCode = "LS"

	// ISO3166CountryCodeLT captures enum value "LT"
	ISO3166CountryCodeLT ISO3166CountryCode = "LT"

	// ISO3166CountryCodeLU captures enum value "LU"
	ISO3166CountryCodeLU ISO3166CountryCode = "LU"

	// ISO3166CountryCodeLV captures enum value "LV"
	ISO3166CountryCodeLV ISO3166CountryCode = "LV"

	// ISO3166CountryCodeLY captures enum value "LY"
	ISO3166CountryCodeLY ISO3166CountryCode = "LY"

	// ISO3166CountryCodeMA captures enum value "MA"
	ISO3166CountryCodeMA ISO3166CountryCode = "MA"

	// ISO3166CountryCodeMC captures enum value "MC"
	ISO3166CountryCodeMC ISO3166CountryCode = "MC"

	// ISO3166CountryCodeMD captures enum value "MD"
	ISO3166CountryCodeMD ISO3166CountryCode = "MD"

	// ISO3166CountryCodeMG captures enum value "MG"
	ISO3166CountryCodeMG ISO3166CountryCode = "MG"

	// ISO3166CountryCodeMH captures enum value "MH"
	ISO3166CountryCodeMH ISO3166CountryCode = "MH"

	// ISO3166CountryCodeMK captures enum value "MK"
	ISO3166CountryCodeMK ISO3166CountryCode = "MK"

	// ISO3166CountryCodeML captures enum value "ML"
	ISO3166CountryCodeML ISO3166CountryCode = "ML"

	// ISO3166CountryCodeMM captures enum value "MM"
	ISO3166CountryCodeMM ISO3166CountryCode = "MM"

	// ISO3166CountryCodeMN captures enum value "MN"
	ISO3166CountryCodeMN ISO3166CountryCode = "MN"

	// ISO3166CountryCodeMO captures enum value "MO"
	ISO3166CountryCodeMO ISO3166CountryCode = "MO"

	// ISO3166CountryCodeMP captures enum value "MP"
	ISO3166CountryCodeMP ISO3166CountryCode = "MP"

	// ISO3166CountryCodeMQ captures enum value "MQ"
	ISO3166CountryCodeMQ ISO3166CountryCode = "MQ"

	// ISO3166CountryCodeMR captures enum value "MR"
	ISO3166CountryCodeMR ISO3166CountryCode = "MR"

	// ISO3166CountryCodeMS captures enum value "MS"
	ISO3166CountryCodeMS ISO3166CountryCode = "MS"

	// ISO3166CountryCodeMT captures enum value "MT"
	ISO3166CountryCodeMT ISO3166CountryCode = "MT"

	// ISO3166CountryCodeMU captures enum value "MU"
	ISO3166CountryCodeMU ISO3166CountryCode = "MU"

	// ISO3166CountryCodeMV captures enum value "MV"
	ISO3166CountryCodeMV ISO3166CountryCode = "MV"

	// ISO3166CountryCodeMW captures enum value "MW"
	ISO3166CountryCodeMW ISO3166CountryCode = "MW"

	// ISO3166CountryCodeMX captures enum value "MX"
	ISO3166CountryCodeMX ISO3166CountryCode = "MX"

	// ISO3166CountryCodeMY captures enum value "MY"
	ISO3166CountryCodeMY ISO3166CountryCode = "MY"

	// ISO3166CountryCodeMZ captures enum value "MZ"
	ISO3166CountryCodeMZ ISO3166CountryCode = "MZ"

	// ISO3166CountryCodeNA captures enum value "NA"
	ISO3166CountryCodeNA ISO3166CountryCode = "NA"

	// ISO3166CountryCodeNC captures enum value "NC"
	ISO3166CountryCodeNC ISO3166CountryCode = "NC"

	// ISO3166CountryCodeNE captures enum value "NE"
	ISO3166CountryCodeNE ISO3166CountryCode = "NE"

	// ISO3166CountryCodeNF captures enum value "NF"
	ISO3166CountryCodeNF ISO3166CountryCode = "NF"

	// ISO3166CountryCodeNG captures enum value "NG"
	ISO3166CountryCodeNG ISO3166CountryCode = "NG"

	// ISO3166CountryCodeNI captures enum value "NI"
	ISO3166CountryCodeNI ISO3166CountryCode = "NI"

	// ISO3166CountryCodeNL captures enum value "NL"
	ISO3166CountryCodeNL ISO3166CountryCode = "NL"

	// ISO3166CountryCodeNO captures enum value "NO"
	ISO3166CountryCodeNO ISO3166CountryCode = "NO"

	// ISO3166CountryCodeNP captures enum value "NP"
	ISO3166CountryCodeNP ISO3166CountryCode = "NP"

	// ISO3166CountryCodeNR captures enum value "NR"
	ISO3166CountryCodeNR ISO3166CountryCode = "NR"

	// ISO3166CountryCodeNU captures enum value "NU"
	ISO3166CountryCodeNU ISO3166CountryCode = "NU"

	// ISO3166CountryCodeNZ captures enum value "NZ"
	ISO3166CountryCodeNZ ISO3166CountryCode = "NZ"

	// ISO3166CountryCodeOM captures enum value "OM"
	ISO3166CountryCodeOM ISO3166CountryCode = "OM"

	// ISO3166CountryCodePA captures enum value "PA"
	ISO3166CountryCodePA ISO3166CountryCode = "PA"

	// ISO3166CountryCodePE captures enum value "PE"
	ISO3166CountryCodePE ISO3166CountryCode = "PE"

	// ISO3166CountryCodePF captures enum value "PF"
	ISO3166CountryCodePF ISO3166CountryCode = "PF"

	// ISO3166CountryCodePG captures enum value "PG"
	ISO3166CountryCodePG ISO3166CountryCode = "PG"

	// ISO3166CountryCodePH captures enum value "PH"
	ISO3166CountryCodePH ISO3166CountryCode = "PH"

	// ISO3166CountryCodePK captures enum value "PK"
	ISO3166CountryCodePK ISO3166CountryCode = "PK"

	// ISO3166CountryCodePL captures enum value "PL"
	ISO3166CountryCodePL ISO3166CountryCode = "PL"

	// ISO3166CountryCodePM captures enum value "PM"
	ISO3166CountryCodePM ISO3166CountryCode = "PM"

	// ISO3166CountryCodePN captures enum value "PN"
	ISO3166CountryCodePN ISO3166CountryCode = "PN"

	// ISO3166CountryCodePR captures enum value "PR"
	ISO3166CountryCodePR ISO3166CountryCode = "PR"

	// ISO3166CountryCodePS captures enum value "PS"
	ISO3166CountryCodePS ISO3166CountryCode = "PS"

	// ISO3166CountryCodePT captures enum value "PT"
	ISO3166CountryCodePT ISO3166CountryCode = "PT"

	// ISO3166CountryCodePW captures enum value "PW"
	ISO3166CountryCodePW ISO3166CountryCode = "PW"

	// ISO3166CountryCodePY captures enum value "PY"
	ISO3166CountryCodePY ISO3166CountryCode = "PY"

	// ISO3166CountryCodeQA captures enum value "QA"
	ISO3166CountryCodeQA ISO3166CountryCode = "QA"

	// ISO3166CountryCodeRE captures enum value "RE"
	ISO3166CountryCodeRE ISO3166CountryCode = "RE"

	// ISO3166CountryCodeRO captures enum value "RO"
	ISO3166CountryCodeRO ISO3166CountryCode = "RO"

	// ISO3166CountryCodeRU captures enum value "RU"
	ISO3166CountryCodeRU ISO3166CountryCode = "RU"

	// ISO3166CountryCodeRW captures enum value "RW"
	ISO3166CountryCodeRW ISO3166CountryCode = "RW"

	// ISO3166CountryCodeSA captures enum value "SA"
	ISO3166CountryCodeSA ISO3166CountryCode = "SA"

	// ISO3166CountryCodeSB captures enum value "SB"
	ISO3166CountryCodeSB ISO3166CountryCode = "SB"

	// ISO3166CountryCodeSC captures enum value "SC"
	ISO3166CountryCodeSC ISO3166CountryCode = "SC"

	// ISO3166CountryCodeSD captures enum value "SD"
	ISO3166CountryCodeSD ISO3166CountryCode = "SD"

	// ISO3166CountryCodeSE captures enum value "SE"
	ISO3166CountryCodeSE ISO3166CountryCode = "SE"

	// ISO3166CountryCodeSG captures enum value "SG"
	ISO3166CountryCodeSG ISO3166CountryCode = "SG"

	// ISO3166CountryCodeSH captures enum value "SH"
	ISO3166CountryCodeSH ISO3166CountryCode = "SH"

	// ISO3166CountryCodeSI captures enum value "SI"
	ISO3166CountryCodeSI ISO3166CountryCode = "SI"

	// ISO3166CountryCodeSJ captures enum value "SJ"
	ISO3166CountryCodeSJ ISO3166CountryCode = "SJ"

	// ISO3166CountryCodeSK captures enum value "SK"
	ISO3166CountryCodeSK ISO3166CountryCode = "SK"

	// ISO3166CountryCodeSL captures enum value "SL"
	ISO3166CountryCodeSL ISO3166CountryCode = "SL"

	// ISO3166CountryCodeSM captures enum value "SM"
	ISO3166CountryCodeSM ISO3166CountryCode = "SM"

	// ISO3166CountryCodeSN captures enum value "SN"
	ISO3166CountryCodeSN ISO3166CountryCode = "SN"

	// ISO3166CountryCodeSO captures enum value "SO"
	ISO3166CountryCodeSO ISO3166CountryCode = "SO"

	// ISO3166CountryCodeSR captures enum value "SR"
	ISO3166CountryCodeSR ISO3166CountryCode = "SR"

	// ISO3166CountryCodeST captures enum value "ST"
	ISO3166CountryCodeST ISO3166CountryCode = "ST"

	// ISO3166CountryCodeSV captures enum value "SV"
	ISO3166CountryCodeSV ISO3166CountryCode = "SV"

	// ISO3166CountryCodeSY captures enum value "SY"
	ISO3166CountryCodeSY ISO3166CountryCode = "SY"

	// ISO3166CountryCodeSZ captures enum value "SZ"
	ISO3166CountryCodeSZ ISO3166CountryCode = "SZ"

	// ISO3166CountryCodeTC captures enum value "TC"
	ISO3166CountryCodeTC ISO3166CountryCode = "TC"

	// ISO3166CountryCodeTD captures enum value "TD"
	ISO3166CountryCodeTD ISO3166CountryCode = "TD"

	// ISO3166CountryCodeTF captures enum value "TF"
	ISO3166CountryCodeTF ISO3166CountryCode = "TF"

	// ISO3166CountryCodeTG captures enum value "TG"
	ISO3166CountryCodeTG ISO3166CountryCode = "TG"

	// ISO3166CountryCodeTH captures enum value "TH"
	ISO3166CountryCodeTH ISO3166CountryCode = "TH"

	// ISO3166CountryCodeTJ captures enum value "TJ"
	ISO3166CountryCodeTJ ISO3166CountryCode = "TJ"

	// ISO3166CountryCodeTK captures enum value "TK"
	ISO3166CountryCodeTK ISO3166CountryCode = "TK"

	// ISO3166CountryCodeTL captures enum value "TL"
	ISO3166CountryCodeTL ISO3166CountryCode = "TL"

	// ISO3166CountryCodeTM captures enum value "TM"
	ISO3166CountryCodeTM ISO3166CountryCode = "TM"

	// ISO3166CountryCodeTN captures enum value "TN"
	ISO3166CountryCodeTN ISO3166CountryCode = "TN"

	// ISO3166CountryCodeTO captures enum value "TO"
	ISO3166CountryCodeTO ISO3166CountryCode = "TO"

	// ISO3166CountryCodeTR captures enum value "TR"
	ISO3166CountryCodeTR ISO3166CountryCode = "TR"

	// ISO3166CountryCodeTT captures enum value "TT"
	ISO3166CountryCodeTT ISO3166CountryCode = "TT"

	// ISO3166CountryCodeTV captures enum value "TV"
	ISO3166CountryCodeTV ISO3166CountryCode = "TV"

	// ISO3166CountryCodeTW captures enum value "TW"
	ISO3166CountryCodeTW ISO3166CountryCode = "TW"

	// ISO3166CountryCodeTZ captures enum value "TZ"
	ISO3166CountryCodeTZ ISO3166CountryCode = "TZ"

	// ISO3166CountryCodeUA captures enum value "UA"
	ISO3166CountryCodeUA ISO3166CountryCode = "UA"

	// ISO3166CountryCodeUG captures enum value "UG"
	ISO3166CountryCodeUG ISO3166CountryCode = "UG"

	// ISO3166CountryCodeUM captures enum value "UM"
	ISO3166CountryCodeUM ISO3166CountryCode = "UM"

	// ISO3166CountryCodeUS captures enum value "US"
	ISO3166CountryCodeUS ISO3166CountryCode = "US"

	// ISO3166CountryCodeUY captures enum value "UY"
	ISO3166CountryCodeUY ISO3166CountryCode = "UY"

	// ISO3166CountryCodeUZ captures enum value "UZ"
	ISO3166CountryCodeUZ ISO3166CountryCode = "UZ"

	// ISO3166CountryCodeVA captures enum value "VA"
	ISO3166CountryCodeVA ISO3166CountryCode = "VA"

	// ISO3166CountryCodeVC captures enum value "VC"
	ISO3166CountryCodeVC ISO3166CountryCode = "VC"

	// ISO3166CountryCodeVE captures enum value "VE"
	ISO3166CountryCodeVE ISO3166CountryCode = "VE"

	// ISO3166CountryCodeVG captures enum value "VG"
	ISO3166CountryCodeVG ISO3166CountryCode = "VG"

	// ISO3166CountryCodeVI captures enum value "VI"
	ISO3166CountryCodeVI ISO3166CountryCode = "VI"

	// ISO3166CountryCodeVN captures enum value "VN"
	ISO3166CountryCodeVN ISO3166CountryCode = "VN"

	// ISO3166CountryCodeVU captures enum value "VU"
	ISO3166CountryCodeVU ISO3166CountryCode = "VU"

	// ISO3166CountryCodeWF captures enum value "WF"
	ISO3166CountryCodeWF ISO3166CountryCode = "WF"

	// ISO3166CountryCodeWS captures enum value "WS"
	ISO3166CountryCodeWS ISO3166CountryCode = "WS"

	// ISO3166CountryCodeYE captures enum value "YE"
	ISO3166CountryCodeYE ISO3166CountryCode = "YE"

	// ISO3166CountryCodeYT captures enum value "YT"
	ISO3166CountryCodeYT ISO3166CountryCode = "YT"

	// ISO3166CountryCodeZA captures enum value "ZA"
	ISO3166CountryCodeZA ISO3166CountryCode = "ZA"

	// ISO3166CountryCodeZM captures enum value "ZM"
	ISO3166CountryCodeZM ISO3166CountryCode = "ZM"

	// ISO3166CountryCodeZW captures enum value "ZW"
	ISO3166CountryCodeZW ISO3166CountryCode = "ZW"
)

// for schema
var iSO3166CountryCodeEnum []interface{}

func init() {
	var res []ISO3166CountryCode
	if err := json.Unmarshal([]byte(`["AD","AE","AF","AG","AI","AL","AM","AN","AO","AQ","AR","AS","AT","AU","AW","AX","AZ","BA","BB","BD","BE","BF","BG","BH","BI","BJ","BM","BN","BO","BR","BS","BT","BV","BW","BY","BZ","CA","CC","CD","CF","CG","CH","CI","CK","CL","CM","CN","CO","CR","CS","CU","CV","CX","CY","CZ","DE","DJ","DK","DM","DO","DZ","EC","EE","EG","EH","ER","ES","ET","FI","FJ","FK","FM","FO","FR","GA","GB","GD","GE","GF","GG","GH","GI","GL","GM","GN","GP","GQ","GR","GS","GT","GU","GW","GY","HK","HM","HN","HR","HT","HU","ID","IE","IL","IM","IN","IO","IQ","IR","IS","IT","JE","JM","JO","JP","KE","KG","KH","KI","KM","KN","KP","KR","KW","KY","KZ","LA","LB","LC","LI","LK","LR","LS","LT","LU","LV","LY","MA","MC","MD","MG","MH","MK","ML","MM","MN","MO","MP","MQ","MR","MS","MT","MU","MV","MW","MX","MY","MZ","NA","NC","NE","NF","NG","NI","NL","NO","NP","NR","NU","NZ","OM","PA","PE","PF","PG","PH","PK","PL","PM","PN","PR","PS","PT","PW","PY","QA","RE","RO","RU","RW","SA","SB","SC","SD","SE","SG","SH","SI","SJ","SK","SL","SM","SN","SO","SR","ST","SV","SY","SZ","TC","TD","TF","TG","TH","TJ","TK","TL","TM","TN","TO","TR","TT","TV","TW","TZ","UA","UG","UM","US","UY","UZ","VA","VC","VE","VG","VI","VN","VU","WF","WS","YE","YT","ZA","ZM","ZW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iSO3166CountryCodeEnum = append(iSO3166CountryCodeEnum, v)
	}
}

func (m ISO3166CountryCode) validateISO3166CountryCodeEnum(path, location string, value ISO3166CountryCode) error {
	if err := validate.EnumCase(path, location, value, iSO3166CountryCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this i s o3166 country code
func (m ISO3166CountryCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateISO3166CountryCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this i s o3166 country code based on context it is used
func (m ISO3166CountryCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
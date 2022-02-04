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
)

// BankingProductAdditionalInformationV2 BankingProductAdditionalInformationV2
//
// Object that contains links to additional information on specific topics
//
// swagger:model BankingProductAdditionalInformationV2
type BankingProductAdditionalInformationV2 struct {

	// An array of additional bundles for the product, if applicable. To be treated as secondary documents to the `bundleUri`. Only to be used if there is a primary `bundleUri`.
	AdditionalBundleUris []*BankingProductAdditionalInformationV2AdditionalInformationUris `json:"additionalBundleUris"`

	// An array of additional eligibility rules and criteria for the product, if applicable. To be treated as secondary documents to the `eligibilityUri`. Only to be used if there is a primary `eligibilityUri`.
	AdditionalEligibilityUris []*BankingProductAdditionalInformationV2AdditionalInformationUris `json:"additionalEligibilityUris"`

	// An array of additional fees, pricing, discounts, exemptions and bonuses for the product, if applicable. To be treated as secondary documents to the `feesAndPricingUri`. Only to be used if there is a primary `feesAndPricingUri`.
	AdditionalFeesAndPricingUris []*BankingProductAdditionalInformationV2AdditionalInformationUris `json:"additionalFeesAndPricingUris"`

	// An array of additional general overviews for the product or features of the product, if applicable. To be treated as secondary documents to the `overviewUri`. Only to be used if there is a primary `overviewUri`.
	AdditionalOverviewUris []*BankingProductAdditionalInformationV2AdditionalInformationUris `json:"additionalOverviewUris"`

	// An array of additional terms and conditions for the product, if applicable. To be treated as secondary documents to the `termsUri`. Only to be used if there is a primary `termsUri`.
	AdditionalTermsUris []*BankingProductAdditionalInformationV2AdditionalInformationUris `json:"additionalTermsUris"`

	// Description of a bundle that this product can be part of. Mandatory if `additionalBundleUris` includes one or more supporting documents.
	BundleURI string `json:"bundleUri,omitempty"`

	// Eligibility rules and criteria for the product. Mandatory if `additionalEligibilityUris` includes one or more supporting documents.
	EligibilityURI string `json:"eligibilityUri,omitempty"`

	// Description of fees, pricing, discounts, exemptions and bonuses for the product. Mandatory if `additionalFeesAndPricingUris` includes one or more supporting documents.
	FeesAndPricingURI string `json:"feesAndPricingUri,omitempty"`

	// General overview of the product. Mandatory if `additionalOverviewUris` includes one or more supporting documents.
	OverviewURI string `json:"overviewUri,omitempty"`

	// Terms and conditions for the product. Mandatory if `additionalTermsUris` includes one or more supporting documents.
	TermsURI string `json:"termsUri,omitempty"`
}

// Validate validates this banking product additional information v2
func (m *BankingProductAdditionalInformationV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalBundleUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalEligibilityUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalFeesAndPricingUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalOverviewUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditionalTermsUris(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductAdditionalInformationV2) validateAdditionalBundleUris(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalBundleUris) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalBundleUris); i++ {
		if swag.IsZero(m.AdditionalBundleUris[i]) { // not required
			continue
		}

		if m.AdditionalBundleUris[i] != nil {
			if err := m.AdditionalBundleUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalBundleUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) validateAdditionalEligibilityUris(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalEligibilityUris) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalEligibilityUris); i++ {
		if swag.IsZero(m.AdditionalEligibilityUris[i]) { // not required
			continue
		}

		if m.AdditionalEligibilityUris[i] != nil {
			if err := m.AdditionalEligibilityUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalEligibilityUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) validateAdditionalFeesAndPricingUris(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalFeesAndPricingUris) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalFeesAndPricingUris); i++ {
		if swag.IsZero(m.AdditionalFeesAndPricingUris[i]) { // not required
			continue
		}

		if m.AdditionalFeesAndPricingUris[i] != nil {
			if err := m.AdditionalFeesAndPricingUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalFeesAndPricingUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) validateAdditionalOverviewUris(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalOverviewUris) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalOverviewUris); i++ {
		if swag.IsZero(m.AdditionalOverviewUris[i]) { // not required
			continue
		}

		if m.AdditionalOverviewUris[i] != nil {
			if err := m.AdditionalOverviewUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalOverviewUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) validateAdditionalTermsUris(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalTermsUris) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalTermsUris); i++ {
		if swag.IsZero(m.AdditionalTermsUris[i]) { // not required
			continue
		}

		if m.AdditionalTermsUris[i] != nil {
			if err := m.AdditionalTermsUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalTermsUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this banking product additional information v2 based on the context it is used
func (m *BankingProductAdditionalInformationV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalBundleUris(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionalEligibilityUris(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionalFeesAndPricingUris(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionalOverviewUris(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditionalTermsUris(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankingProductAdditionalInformationV2) contextValidateAdditionalBundleUris(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalBundleUris); i++ {

		if m.AdditionalBundleUris[i] != nil {
			if err := m.AdditionalBundleUris[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalBundleUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) contextValidateAdditionalEligibilityUris(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalEligibilityUris); i++ {

		if m.AdditionalEligibilityUris[i] != nil {
			if err := m.AdditionalEligibilityUris[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalEligibilityUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) contextValidateAdditionalFeesAndPricingUris(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalFeesAndPricingUris); i++ {

		if m.AdditionalFeesAndPricingUris[i] != nil {
			if err := m.AdditionalFeesAndPricingUris[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalFeesAndPricingUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) contextValidateAdditionalOverviewUris(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalOverviewUris); i++ {

		if m.AdditionalOverviewUris[i] != nil {
			if err := m.AdditionalOverviewUris[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalOverviewUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BankingProductAdditionalInformationV2) contextValidateAdditionalTermsUris(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalTermsUris); i++ {

		if m.AdditionalTermsUris[i] != nil {
			if err := m.AdditionalTermsUris[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalTermsUris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankingProductAdditionalInformationV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankingProductAdditionalInformationV2) UnmarshalBinary(b []byte) error {
	var res BankingProductAdditionalInformationV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
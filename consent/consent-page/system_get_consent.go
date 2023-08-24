package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
)

func GetOBBRPaymentsSystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	var (
		response *obbrModels.GetOBBRCustomerPaymentConsentSystemOK
		err      error
	)

	if response, err = client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	); err != nil {
		return OBBRConsentWrapper{}, err
	}

	if response.Payload.CustomerPaymentConsent != nil && response.Payload.CustomerPaymentConsent.ConsentID != "" {
		return OBBRConsentWrapper{response, OBBRPaymentsV1SystemConsent{response.Payload.CustomerPaymentConsent}}, nil
	}

	if response.Payload.CustomerPaymentConsentV2 != nil && response.Payload.CustomerPaymentConsentV2.ConsentID != "" {
		return OBBRConsentWrapper{response, OBBRPaymentsV2SystemConsent{response.Payload.CustomerPaymentConsentV2}}, nil
	}

	if response.Payload.CustomerPaymentConsentV3 != nil && response.Payload.CustomerPaymentConsentV3.ConsentID != "" {
		return OBBRConsentWrapper{response, OBBRPaymentsV3SystemConsent{response.Payload.CustomerPaymentConsentV3}}, nil
	}

	return OBBRConsentWrapper{}, errors.New("system get consent response was empty")
}

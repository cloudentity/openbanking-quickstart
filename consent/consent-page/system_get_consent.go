package main

import (
	"github.com/gin-gonic/gin"

	acpclient "github.com/cloudentity/acp-client-go"
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
)

type SystemConsentRetriever func(*gin.Context, acpclient.Client, LoginRequest) (OBBRConsentWrapper, error)

func GetOBBRPaymentsV1SystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	response, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	)
	return OBBRConsentWrapper{
		Version:                     V1,
		OBBRPaymentsV1SystemConsent: OBBRPaymentsV1SystemConsent{response},
	}, err
}

func GetOBBRPaymentsV2SystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	response, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	)
	return OBBRConsentWrapper{
		Version:                     V2,
		OBBRPaymentsV2SystemConsent: OBBRPaymentsV2SystemConsent{response},
	}, err
}

func GetOBBRPaymentsV3SystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	response, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV3(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV3ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	)
	return OBBRConsentWrapper{
		Version:                     V3,
		OBBRPaymentsV3SystemConsent: OBBRPaymentsV3SystemConsent{response},
	}, err
}

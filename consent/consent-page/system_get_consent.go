package main

import (
	acpclient "github.com/cloudentity/acp-client-go"
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/client/c_o_n_s_e_n_t_p_a_g_e"
	"github.com/gin-gonic/gin"
)

type SystemConsentRetriever func(*gin.Context, acpclient.Client, LoginRequest) (OBBRConsentWrapper, error)

func GetOBBRPaymentsV1SystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	response, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystem(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	)
	return OBBRConsentWrapper{version: V1, v1: response}, err
}

func GetOBBRPaymentsV2SystemConsent(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) (OBBRConsentWrapper, error) {
	response, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV2(
		obbrModels.NewGetOBBRCustomerPaymentConsentSystemV2ParamsWithContext(c).
			WithLogin(loginRequest.ID),
		nil,
	)
	return OBBRConsentWrapper{version: V2, v2: response}, err
}

// func OBBRPaymentsV3SystemConsentRetriever(c *gin.Context, client acpclient.Client, loginRequest LoginRequest) error {
// 	_, err := client.Obbr.Consentpage.GetOBBRCustomerPaymentConsentSystemV3(
// 		obbrModels.NewGetOBBRCustomerPaymentConsentSystemParamsWithContext(c).
// 			WithLogin(loginRequest.ID),
// 		nil,
// 	)
// 	return err
// }

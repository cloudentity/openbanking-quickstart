package main

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	cdr "github.com/cloudentity/acp-client-go/clients/openbanking/client/c_d_r"
	obbr "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_b_r"
	obuk "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_u_k"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

func (s *Server) OBUKIntrospectAccountsToken(c *gin.Context) (*obuk.OpenbankingAccountAccessConsentIntrospectOKBody, error) {
	var (
		introspectionResponse *obuk.OpenbankingAccountAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.Openbankinguk.OpenbankingAccountAccessConsentIntrospect(
		obuk.NewOpenbankingAccountAccessConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		return nil, err
	}

	if !introspectionResponse.Payload.Active {
		return nil, errors.New("access token is not active")
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) OBUKIntrospectPaymentsToken(c *gin.Context) (*obuk.OpenbankingDomesticPaymentConsentIntrospectOKBody, error) {
	var (
		introspectionResponse *obuk.OpenbankingDomesticPaymentConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.Openbankinguk.OpenbankingDomesticPaymentConsentIntrospect(
		obuk.NewOpenbankingDomesticPaymentConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		logrus.WithError(err).Errorf("failed to introspect token %s", token)
		return nil, err
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) OBBRIntrospectAccountsToken(c *gin.Context) (*obModels.IntrospectOBBRDataAccessConsentResponse, error) {
	var (
		introspectionResponse *obbr.ObbrDataAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.Openbankingbr.ObbrDataAccessConsentIntrospect(
		obbr.NewObbrDataAccessConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		return nil, err
	}

	if !introspectionResponse.Payload.Active {
		return nil, errors.New("access token is not active")
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) OBBRIntrospectPaymentsToken(c *gin.Context) (*obModels.IntrospectOBBRPaymentConsentResponse, error) {
	var (
		introspectionResponse *obbr.ObbrPaymentConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.Openbankingbr.ObbrPaymentConsentIntrospect(
		obbr.NewObbrPaymentConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		return nil, err
	}

	if !introspectionResponse.Payload.Active {
		return nil, errors.New("access token is not active")
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) CDRIntrospectAccountsToken(c *gin.Context) (*cdr.CdrConsentIntrospectOKBody, error) {
	var (
		introspectResponse *cdr.CdrConsentIntrospectOK
		err                error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectResponse, err = s.Client.Openbanking.Cdr.CdrConsentIntrospect(
		cdr.NewCdrConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		return nil, err
	}

	if !introspectResponse.Payload.Active {
		return nil, errors.New("access token is not active")
	}

	return introspectResponse.Payload, nil
}

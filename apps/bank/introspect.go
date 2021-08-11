package main

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	acpClient "github.com/cloudentity/acp-client-go/models"
)

func (s *Server) OBUKIntrospectAccountsToken(c *gin.Context) (*acpClient.IntrospectOpenbankingAccountAccessConsentResponse, error) {
	var (
		introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.OpenbankingAccountAccessConsentIntrospect(
		openbanking.NewOpenbankingAccountAccessConsentIntrospectParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithAid(s.Client.ServerID).
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

func (s *Server) OBUKIntrospectPaymentsToken(c *gin.Context) (*acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse, error) {
	var (
		introspectionResponse *openbanking.OpenbankingDomesticPaymentConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.OpenbankingDomesticPaymentConsentIntrospect(
		openbanking.NewOpenbankingDomesticPaymentConsentIntrospectParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithAid(s.Client.ServerID).
			WithToken(&token),
		nil,
	); err != nil {
		logrus.WithError(err).Errorf("failed to introspect token %s", token)
		return nil, err
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) OBBRIntrospectAccountsToken(c *gin.Context) (*acpClient.IntrospectOBBRDataAccessConsentResponse, error) {
	var (
		introspectionResponse *openbanking.ObbrDataAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.ObbrDataAccessConsentIntrospect(
		openbanking.NewObbrDataAccessConsentIntrospectParamsWithContext(c).
			WithTid(s.Client.TenantID).
			WithAid(s.Client.ServerID).
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

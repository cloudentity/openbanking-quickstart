package main

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	cdr "github.com/cloudentity/acp-client-go/clients/cdr/client/c_d_r"
	fdx "github.com/cloudentity/acp-client-go/clients/fdx/client/f_d_x"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	obbr "github.com/cloudentity/acp-client-go/clients/obbr/client/o_b_b_r"
	obbrModels "github.com/cloudentity/acp-client-go/clients/obbr/models"
	obuk "github.com/cloudentity/acp-client-go/clients/obuk/client/o_b_u_k"

	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

func (s *Server) OBUKIntrospectAccountsToken(c *gin.Context) (*obuk.OpenbankingAccountAccessConsentIntrospectOKBody, error) {
	var (
		introspectionResponse *obuk.OpenbankingAccountAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Obuk.Obuk.OpenbankingAccountAccessConsentIntrospect(
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

	if introspectionResponse, err = s.Client.Obuk.Obuk.OpenbankingDomesticPaymentConsentIntrospect(
		obuk.NewOpenbankingDomesticPaymentConsentIntrospectParamsWithContext(c).
			WithToken(&token),
		nil,
	); err != nil {
		logrus.WithError(err).Errorf("failed to introspect token %s", token)
		return nil, err
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) OBBRIntrospectAccountsToken(c *gin.Context) (*obbrModels.IntrospectOBBRDataAccessConsentResponse, error) {
	var (
		introspectionResponse *obbr.ObbrDataAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Obbr.Obbr.ObbrDataAccessConsentIntrospect(
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

func (s *Server) OBBRIntrospectPaymentsToken(c *gin.Context) (*obbrModels.IntrospectOBBRPaymentConsentResponse, error) {
	var (
		introspectionResponse *obbr.ObbrPaymentConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Obbr.Obbr.ObbrPaymentConsentIntrospect(
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
	if introspectResponse, err = s.Client.Cdr.Cdr.CdrConsentIntrospect(
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

func (s *Server) FDXIntrospectAccountsToken(c *gin.Context) (*fdx.FdxConsentIntrospectOKBody, error) {
	var (
		introspectResponse *fdx.FdxConsentIntrospectOK
		err                error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	if introspectResponse, err = s.Client.Fdx.Fdx.FdxConsentIntrospect(
		fdx.NewFdxConsentIntrospectParamsWithContext(c).
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

func (s *Server) GenericIntrospectAccountsToken(c *gin.Context) (*oauth2Models.IntrospectResponse, error) {
	var (
		introspectionResponse *oauth2.IntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Oauth2.Oauth2.Introspect(
		oauth2.NewIntrospectParamsWithContext(c).
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

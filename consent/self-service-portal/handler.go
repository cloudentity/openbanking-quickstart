package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	o2Params "github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"

	obCommon "github.com/cloudentity/acp-client-go/clients/openbanking/client/openbanking_common"
	obModels "github.com/cloudentity/acp-client-go/clients/openbanking/models"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LogoURI   string `json:"logo_uri"`
	ClientURI string `json:"client_uri"`
}

type ClientConsents struct {
	Client
	Consents []obModels.ConsentWithClient `json:"consents"`
}

type ConsentsResponse struct {
	ClientConsents []ClientConsents `json:"client_consents"`
	Accounts       InternalAccounts `json:"accounts"`
}

func (s *Server) ListConsents() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			consentsByAccounts *ConsentsAndAccounts
			clientToConsents   = map[string][]obModels.ConsentWithClient{}
			clients            = map[string]Client{}
			res                = []ClientConsents{}
			err                error
		)

		if consentsByAccounts, err = s.FetchConsents(c); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		logrus.Infof("consents by accounts %v", consentsByAccounts)

		for _, c := range consentsByAccounts.Consents {
			if _, ok := clients[c.Client.ID]; !ok {
				clients[c.Client.ID] = Client{
					ID:        c.Client.ID,
					Name:      c.Client.Name,
					LogoURI:   c.Client.LogoURI,
					ClientURI: c.Client.ClientURI,
				}
			}

			clientToConsents[c.Client.ID] = append(clientToConsents[c.Client.ID], *c)
		}

		for _, x := range clients {
			res = append(res, ClientConsents{
				Client:   x,
				Consents: clientToConsents[x.ID],
			})
		}

		c.JSON(http.StatusOK, &ConsentsResponse{ClientConsents: res, Accounts: consentsByAccounts.Accounts})
	}
}

type ConsentsAndAccounts struct {
	Consents []*obModels.ConsentWithClient `json:"consents"`
	Accounts InternalAccounts              `json:"accounts"`
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                 = c.Param("id")
			consentsByAccounts *ConsentsAndAccounts
			canBeRevoked       bool
			err                error
		)

		if consentsByAccounts, err = s.FetchConsents(c); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		for _, c := range consentsByAccounts.Consents {
			if c.ConsentID == id {
				canBeRevoked = true
				break
			}
		}

		if !canBeRevoked {
			Error(c, APIError{
				Code:    http.StatusUnauthorized,
				Message: "user is not authorized to revoke this consent",
			})
			return
		}

		if _, err = s.Client.Openbanking.OpenbankingCommon.RevokeOpenbankingConsent(
			obCommon.NewRevokeOpenbankingConsentParamsWithContext(c).
				WithAid(s.Config.SystemClientsServerID).
				WithConsentID(id),
			nil,
		); err != nil {
			Error(c, ToAPIError(err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

var (
	ErrTokenNotActive = errors.New("token is not active")
	ErrTokenMissing   = errors.New("token is missing")
)

func (s *Server) FetchConsents(c *gin.Context) (*ConsentsAndAccounts, error) {
	var (
		accounts       InternalAccounts
		response       *obCommon.ListOBConsentsOK
		introspectResp *o2Params.IntrospectOK
		err            error
		types          []string
		ok             bool
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if token == "" {
		return nil, ErrTokenMissing
	}

	if types, ok = c.GetQueryArray("types"); !ok {
		types = nil
	}

	if introspectResp, err = s.IntrospectClient.Oauth2.Oauth2.Introspect(o2Params.NewIntrospectParamsWithContext(c).
		WithToken(&token), nil); err != nil {
		return nil, err
	}

	if !introspectResp.Payload.Active {
		return nil, ErrTokenNotActive
	}

	if accounts, err = s.BankClient.GetInternalAccounts(introspectResp.Payload.Sub); err != nil {
		return nil, fmt.Errorf("failed to get accounts from bank: %w", err)
	}

	logrus.Infof("accounts %v", accounts)

	accountIDs := make([]string, len(accounts.Accounts))
	for i, a := range accounts.Accounts {
		accountIDs[i] = a.ID
	}

	if response, err = s.Client.Openbanking.OpenbankingCommon.ListOBConsents(
		obCommon.NewListOBConsentsParamsWithContext(c).
			WithAid(s.Config.SystemClientsServerID).
			WithConsentsRequest(&obModels.ConsentsRequest{
				Accounts: accountIDs,
				Types:    types,
			}),
		nil,
	); err != nil {
		return nil, err
	}

	logrus.Infof("list ob consents response %v", response.Payload)

	return &ConsentsAndAccounts{
		Consents: response.Payload.Consents,
		Accounts: accounts,
	}, nil
}

func ToAPIError(err error) APIError {
	if errors.Is(err, ErrTokenNotActive) {
		return APIError{
			http.StatusUnauthorized,
			"token is not valid",
			nil,
		}
	} else if errors.Is(err, ErrTokenMissing) {
		return APIError{
			http.StatusUnauthorized,
			"token is missing",
			nil,
		}
	}

	return APIError{
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		err,
	}
}

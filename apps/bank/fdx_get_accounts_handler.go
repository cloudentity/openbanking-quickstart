package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"

	fdxModels "github.com/cloudentity/acp-client-go/clients/fdx/client/f_d_x"
)

type FDXGetAccountsHandler struct {
	*Server
	introspectionResponse *fdxModels.FdxConsentIntrospectOKBody
}

func NewFDXGetAccountsHandler(server *Server) GetEndpointLogic {
	return &FDXGetAccountsHandler{Server: server}
}

func (h *FDXGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	var err error

	if h.introspectionResponse, err = h.FDXIntrospectAccountsToken(c); err != nil {
		return ErrBadRequest.WithMessage("failed to introspect token")
	}

	return nil
}

func (h *FDXGetAccountsHandler) MapError(_ *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetAccountsHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(h.Config.Port)))
	return NewFDXAccountsResponse(data.FDXAccounts, self), nil
}

func (h *FDXGetAccountsHandler) Validate(_ *gin.Context) *Error {
	scopes := strings.Split(h.introspectionResponse.Scope, " ")
	if !has(scopes, "fdx:accountdetailed:read") {
		return ErrForbidden.WithMessage("token has no fdx:accountdetailed:read scope granted")
	}

	return nil
}

func (h *FDXGetAccountsHandler) GetUserIdentifier(_ *gin.Context) string {
	return h.introspectionResponse.Sub
}

func (h *FDXGetAccountsHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	var (
		jsonStr []byte
		err     error
	)

	var depositAccounts []models.AccountWithDetailsentity
	for _, acct := range data.FDXAccounts.Accounts {
		var depositAccount models.AccountWithDetailsentity
		if jsonStr, err = json.Marshal(acct); err != nil {
			continue
		}

		if err = json.Unmarshal(jsonStr, &depositAccount); err != nil {
			continue
		}
		depositAccounts = append(depositAccounts, depositAccount)
	}

	var filteredData BankUserData
	for _, a := range depositAccounts {
		for _, c := range h.introspectionResponse.FdxConsent.Resources {
			if c.ID == a.DepositAccount.AccountID {
				filteredData.FDXAccounts.Accounts = append(filteredData.FDXAccounts.Accounts, a)
			}
		}
	}

	return filteredData
}

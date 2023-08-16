package main

import (
	obbrAccountModels "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
	"github.com/cloudentity/openbanking-quickstart/generated/obbr/consents/models"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /internal/accounts bank generic getInternalAccountsRequest
//
// get all accounts for user
//
// Security:
//
//	defaultcc: accounts
//
// Responses:
//
//	200: ResponseAccountList
//	404: OpenbankingBrasilResponseError
type GenericGetAccountsInternalHandler struct {
	*Server
}

func NewGenericGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &GenericGetAccountsInternalHandler{Server: server}
}

func (h *GenericGetAccountsInternalHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	return nil
}

func (h *GenericGetAccountsInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = GenericMapError(c, err)
	return
}

func (h *GenericGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewGenericAccountsResponse(data.GenericAccounts), nil
}

func (h *GenericGetAccountsInternalHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *GenericGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *GenericGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

func GenericMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = err.Code, models.OpenbankingBrasilConsentV2ResponseError{
		Errors: []*models.OpenbankingBrasilConsentV2Error{
			{
				Detail: err.Message,
			},
		},
	}
	return
}

func NewGenericAccountsResponse(accounts []obbrAccountModels.AccountData) obbrAccountModels.ResponseAccountList {
	accountPointers := []*obbrAccountModels.AccountData{}
	for _, account := range accounts {
		a := account
		accountPointers = append(accountPointers, &a)
	}

	return obbrAccountModels.ResponseAccountList{
		Data: accountPointers,
	}
}

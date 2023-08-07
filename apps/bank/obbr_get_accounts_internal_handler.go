package main

import (
	"github.com/gin-gonic/gin"
)

// swagger:route GET /internal/accounts bank br getInternalAccountsRequest
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
type OBBRGetAccountsInternalHandler struct {
	*Server
}

func NewOBBRGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &OBBRGetAccountsInternalHandler{Server: server}
}

func (h *OBBRGetAccountsInternalHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	return nil
}

func (h *OBBRGetAccountsInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = OBBRMapError(c, err)
	return
}

func (h *OBBRGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewOBBRAccountsResponse(data.OBBRAccounts), nil
}

func (h *OBBRGetAccountsInternalHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *OBBRGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Query("id")
}

func (h *OBBRGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

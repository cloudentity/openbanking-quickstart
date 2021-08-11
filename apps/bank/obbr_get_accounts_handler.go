package main

import (
	"github.com/gin-gonic/gin"
)

type OBBRGetAccountsHandler struct {
	*Server
	// introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func NewOBBRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &OBBRGetAccountsHandler{Server: server}
}

func (h *OBBRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	return nil
}

func (h *OBBRGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = 500, nil
	return
}

func (h *OBBRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	return nil
}

func (h *OBBRGetAccountsHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *OBBRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return ""
}

func (h *OBBRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

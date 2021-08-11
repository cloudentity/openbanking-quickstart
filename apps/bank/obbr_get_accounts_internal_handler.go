package main

import (
	"github.com/gin-gonic/gin"
)

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

func (h *OBBRGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	return NewOBBRAccountsResponse(data.Accounts.OBBR)
}

func (h *OBBRGetAccountsInternalHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *OBBRGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.Param("sub")
}

func (h *OBBRGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

package main

import (
	"github.com/gin-gonic/gin"
)

type CDRGetAccountsInternalHandler struct {
	*Server
}

func NewCDRGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &CDRGetAccountsInternalHandler{Server: server}
}

func (h *CDRGetAccountsInternalHandler) SetIntrospectionResponse(_ *gin.Context) *Error {
	// should try introspecting for this at some point
	return nil
}

func (h *CDRGetAccountsInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetAccountsInternalHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRAccountsResponse(data.CDRAccounts), nil
}

func (h *CDRGetAccountsInternalHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *CDRGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.PostForm("customer_id")
}

func (h *CDRGetAccountsInternalHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}

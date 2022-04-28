package main

import (
	"github.com/gin-gonic/gin"
)

type CDRGetAccountsHandler struct {
	*Server
}

func NewCDRGetAccountsHandler(server *Server) GetEndpointLogic {
	return &CDRGetAccountsHandler{Server: server}
}

func (h *CDRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	// should try introspecting for this at some point
	return nil
}

func (h *CDRGetAccountsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRAccountsResponse(data.CDRAccounts), nil
}

func (h *CDRGetAccountsHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *CDRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return c.PostForm("bfb689fb-7745-45b9-bbaa-b21e00072447")
}

func (h *CDRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

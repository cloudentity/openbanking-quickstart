package main

import (
	"github.com/gin-gonic/gin"
)

type CDRGetTransactionsHandler struct {
	*Server
}

func NewCDRGetTransactionsHandler(server *Server) GetEndpointLogic {
	return &CDRGetTransactionsHandler{Server: server}
}

func (h *CDRGetTransactionsHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	// should try introspecting for this at some point
	return nil
}

func (h *CDRGetTransactionsHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetTransactionsHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRTransactionsResponse(data.CDRTransactions), nil
}

func (h *CDRGetTransactionsHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *CDRGetTransactionsHandler) GetUserIdentifier(c *gin.Context) string {
	return "bfb689fb-7745-45b9-bbaa-b21e00072447"
}

func (h *CDRGetTransactionsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

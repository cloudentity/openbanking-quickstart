package main

import (
	"github.com/gin-gonic/gin"
)

type CDRGetBalancesHandler struct {
	*Server
}

func NewCDRGetBalancesHandler(server *Server) GetEndpointLogic {
	return &CDRGetBalancesHandler{Server: server}
}

func (h *CDRGetBalancesHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	// should try introspecting for this at some point
	return nil
}

func (h *CDRGetBalancesHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = CDRMapError(c, err)
	return
}

func (h *CDRGetBalancesHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	return NewCDRBalancesResponse(data.CDRBalances), nil
}

func (h *CDRGetBalancesHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *CDRGetBalancesHandler) GetUserIdentifier(c *gin.Context) string {
	return "bfb689fb-7745-45b9-bbaa-b21e00072447"
}

func (h *CDRGetBalancesHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

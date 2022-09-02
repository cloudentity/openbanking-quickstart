package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type FDXGetAccountsInternalHandler struct {
	*Server
}

func NewFDXGetAccountsInternalHandler(server *Server) GetEndpointLogic {
	return &FDXGetAccountsInternalHandler{Server: server}
}

func (h *FDXGetAccountsInternalHandler) SetIntrospectionResponse(c *gin.Context) *Error {
	// should try introspecting for this at some point
	return nil
}

func (h *FDXGetAccountsInternalHandler) MapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetAccountsInternalHandler) BuildResponse(c *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/internal/accounts", strconv.Itoa(h.Config.Port)))
	return NewFDXAccountsResponse(data.FDXAccounts, self), nil
}

func (h *FDXGetAccountsInternalHandler) Validate(c *gin.Context) *Error {
	return nil
}

func (h *FDXGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return "user"
}

func (h *FDXGetAccountsInternalHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

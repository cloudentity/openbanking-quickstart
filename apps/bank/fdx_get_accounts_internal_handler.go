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

func (h *FDXGetAccountsInternalHandler) SetIntrospectionResponse(_ *gin.Context) *Error {
	return nil
}

func (h *FDXGetAccountsInternalHandler) MapError(_ *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = FDXMapError(err)
	return
}

func (h *FDXGetAccountsInternalHandler) BuildResponse(_ *gin.Context, data BankUserData) (interface{}, *Error) {
	self := strfmt.URI(fmt.Sprintf("http://localhost:%s/internal/accounts", strconv.Itoa(h.Config.Port)))
	return NewFDXAccountsResponse(data.FDXAccounts, self), nil
}

func (h *FDXGetAccountsInternalHandler) Validate(_ *gin.Context) *Error {
	return nil
}

func (h *FDXGetAccountsInternalHandler) GetUserIdentifier(c *gin.Context) string {
	return c.PostForm("user_id")
}

func (h *FDXGetAccountsInternalHandler) Filter(_ *gin.Context, data BankUserData) BankUserData {
	return data
}

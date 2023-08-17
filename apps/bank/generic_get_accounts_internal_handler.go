package main

import (
	"github.com/gin-gonic/gin"
)

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

func GenericCDRMapError(c *gin.Context, err *Error) (code int, resp interface{}) {
	code, resp = 400, nil
	return
}

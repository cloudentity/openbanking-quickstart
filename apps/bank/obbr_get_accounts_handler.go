package main

import (
	acpClient "github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
)

type OBBRGetAccountsHandler struct {
	*Server
	introspectionResponse *acpClient.IntrospectOBBRDataAccessConsentResponse
}

func (h *OBBRGetAccountsHandler) SetIntrospectionResponse(c *gin.Context) error {
	return nil
}

func (h *OBBRGetAccountsHandler) MapError(c *gin.Context, err error) interface{} {
	return nil
}

func (h *OBBRGetAccountsHandler) BuildResponse(c *gin.Context, data BankUserData) interface{} {
	return nil
}

func (h *OBBRGetAccountsHandler) Validate(c *gin.Context) error {
	return nil
}

func (h *OBBRGetAccountsHandler) GetUserIdentifier(c *gin.Context) string {
	return ""
}

func (h *OBBRGetAccountsHandler) Filter(c *gin.Context, data BankUserData) BankUserData {
	return data
}

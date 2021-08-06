package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	acpClient "github.com/cloudentity/acp-client-go/models"
)

type DomesticPaymentStatus string

const (
	AcceptedSettlementInProcess DomesticPaymentStatus = "AcceptedSettlementInProcess"
	AcceptedSettlementCompleted DomesticPaymentStatus = "AcceptedSettlementCompleted"
)

func (s *Server) CreateDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
			paymentRequest        *paymentModels.OBWriteDomestic2
			err                   error
			errAlreadyExists      ErrAlreadyExists
		)

		if err = json.NewDecoder(c.Request.Body).Decode(&paymentRequest); err != nil {
			msg := "unable to decode domestic payments request object"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		if introspectionResponse, err = s.IntrospectPaymentsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "payments") {
			msg := "token has no payments scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if introspectionResponse.Status != "Authorised" {
			msg := "domestic payment consent does not have status authorised"
			c.JSON(http.StatusUnprocessableEntity, models.OBError1{
				Message: &msg,
			})
			return
		}

		if paymentRequest.Data.Initiation == nil {
			msg := "initiation data not present in request"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		if introspectionResponse.Initiation == nil {
			msg := "initiation data not present in introspection response"
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		if !initiationsAreEqual(*paymentRequest.Data.Initiation, *introspectionResponse.Initiation) {
			msg := "request initiation does not match consent initiation"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		if paymentRequest.Risk == nil {
			msg := "no risk data in payment request"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		paymentRisk := paymentRequest.Risk
		consentRisk := &paymentModels.OBRisk1{}

		if err = copier.Copy(consentRisk, introspectionResponse); err != nil {
			msg := "internal error"
			logrus.WithError(err).Error("field copying failed")
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		if !reflect.DeepEqual(paymentRisk, consentRisk) {
			msg := "risk validation failed"
			logrus.Errorf(msg)
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		id := uuid.New().String()
		status := string(AcceptedSettlementInProcess)
		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments/%s", strconv.Itoa(s.Config.Port), id))
		response := paymentModels.OBWriteDomesticResponse5{
			Data: &paymentModels.OBWriteDomesticResponse5Data{
				DomesticPaymentID:    &id,
				ConsentID:            &introspectionResponse.ConsentID,
				Status:               &status,
				Charges:              []*paymentModels.OBWriteDomesticResponse5DataChargesItems0{},
				CreationDateTime:     newDateTimePtr(time.Now()),
				StatusUpdateDateTime: newDateTimePtr(time.Now()),
				Initiation:           toDomesticResponse5DataInitiation(introspectionResponse.Initiation),
			},
			Links: &paymentModels.Links{
				Self: &self,
			},
		}

		// create resource
		if err = s.Storage.CreateDomesticPayment(introspectionResponse.Sub, response); err != nil {
			msg := err.Error()
			if errors.As(err, &errAlreadyExists) {
				c.JSON(http.StatusConflict, models.OBError1{
					Message: &msg,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		// add to payment queue worker
		// s.PaymentQueue.queue <- response

		c.PureJSON(http.StatusCreated, response)
	}
}

func (s *Server) GetDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			payment               paymentModels.OBWriteDomesticResponse5
			err                   error
			domesticPaymentID     = c.Param("DomesticPaymentId")
			introspectionResponse *acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse
			errNotFound           ErrNotFound
		)

		if introspectionResponse, err = s.IntrospectPaymentsToken(c); err != nil {
			return
		}

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "payments") {
			msg := "token has no payments scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if payment, err = s.Storage.GetDomesticPayment(introspectionResponse.Sub, domesticPaymentID); err != nil {
			if errors.As(err, &errNotFound) {
				msg := "domestic payment id not found"
				c.JSON(http.StatusNotFound, models.OBErrorResponse1{
					Message: &msg,
				})
				return
			}
			msg := "retrieving domestic payment id failed"
			c.JSON(http.StatusInternalServerError, models.OBError1{
				Message: &msg,
			})
			return
		}

		c.PureJSON(http.StatusOK, payment)
	}
}

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			accounts              []models.OBAccount6
			err                   error
		)

		if introspectionResponse, err = s.IntrospectAccountsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadAccountsBasic") {
			msg := "ReadAccountsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if accounts, err = s.Storage.GetAccounts(introspectionResponse.Sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		filteredAccounts := []models.OBAccount6{}

		for _, account := range accounts {
			if has(introspectionResponse.AccountIDs, string(*account.AccountID)) {
				if !has(grantedPermissions, "ReadAccountsDetail") {
					account.Account = []*models.OBAccount6AccountItems0{}
				}

				filteredAccounts = append(filteredAccounts, account)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewAccountsResponse(filteredAccounts, self))
	}
}

func (s *Server) InternalGetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts []models.OBAccount6
			sub      = c.Param("sub")
			err      error
		)

		if accounts, err = s.Storage.GetAccounts(sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewAccountsResponse(accounts, self))
	}
}

func NewAccountsResponse(accounts []models.OBAccount6, self strfmt.URI) models.OBReadAccount6 {
	accountsPointers := make([]*models.OBAccount6, len(accounts))

	for i, a := range accounts {
		account := a
		accountsPointers[i] = &account
	}

	return models.OBReadAccount6{
		Data: &models.OBReadAccount6Data{
			Account: accountsPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(accounts)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			balances              []models.OBReadBalance1DataBalanceItems0
			err                   error
		)

		if introspectionResponse, err = s.IntrospectAccountsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadBalances") {
			msg := "ReadBalances permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if balances, err = s.Storage.GetBalances(introspectionResponse.Sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		filteredBalances := []models.OBReadBalance1DataBalanceItems0{}
		for _, balance := range balances {
			if has(introspectionResponse.AccountIDs, string(*balance.AccountID)) {
				filteredBalances = append(filteredBalances, balance)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewBalancesResponse(filteredBalances, self))
	}
}

func (s *Server) InternalGetBalances() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			balances []models.OBReadBalance1DataBalanceItems0
			sub      = c.Param("sub")
			err      error
		)

		if balances, err = s.Storage.GetBalances(sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(s.Config.Port)))
		c.PureJSON(http.StatusOK, NewBalancesResponse(balances, self))
	}
}

func NewBalancesResponse(balances []models.OBReadBalance1DataBalanceItems0, self strfmt.URI) models.OBReadBalance1 {
	balancesPointers := make([]*models.OBReadBalance1DataBalanceItems0, len(balances))

	for i, b := range balances {
		balance := b
		balancesPointers[i] = &balance
	}

	return models.OBReadBalance1{
		Data: &models.OBReadBalance1Data{
			Balance: balancesPointers,
		},
		Meta: &models.Meta{
			TotalPages: int32(len(balances)),
		},
		Links: &models.Links{
			Self: &self,
		},
	}
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userTransactions      []models.OBTransaction6
			err                   error
		)

		if introspectionResponse, err = s.IntrospectAccountsToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadTransactionsBasic") {
			msg := "ReadTransactionsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if userTransactions, err = s.Storage.GetTransactions(introspectionResponse.Sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		transactions := []*models.OBTransaction6{}

		for _, transaction := range userTransactions {
			t := transaction
			if has(introspectionResponse.AccountIDs, string(*t.AccountID)) {
				if !has(grantedPermissions, "ReadTransactionsDetail") {
					t.TransactionInformation = ""
					t.Balance = &models.OBTransactionCashBalance{}
					t.MerchantDetails = &models.OBMerchantDetails1{}
					t.CreditorAgent = &models.OBBranchAndFinancialInstitutionIdentification61{}
					t.CreditorAccount = &models.OBCashAccount60{}
					t.DebtorAgent = &models.OBBranchAndFinancialInstitutionIdentification62{}
					t.DebtorAccount = &models.OBCashAccount61{}
				}

				transactions = append(transactions, &t)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(s.Config.Port)))

		response := models.OBReadTransaction6{
			Data: &models.OBReadDataTransaction6{
				Transaction: transactions,
			},
			Meta: &models.Meta{
				TotalPages: int32(len(transactions)),
			},
			Links: &models.Links{
				Self: &self,
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}

func (s *Server) IntrospectAccountsToken(c *gin.Context) (*acpClient.IntrospectOpenbankingAccountAccessConsentResponse, error) {
	var (
		introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.OpenbankingAccountAccessConsentIntrospect(
		openbanking.NewOpenbankingAccountAccessConsentIntrospectParams().
			WithTid(s.Client.TenantID).
			WithAid(s.Client.ServerID).
			WithToken(&token),
		nil,
	); err != nil {
		return nil, err
	}

	if !introspectionResponse.Payload.Active {
		return nil, errors.New("access token is not active")
	}

	return introspectionResponse.Payload, nil
}

func (s *Server) IntrospectPaymentsToken(c *gin.Context) (*acpClient.IntrospectOpenbankingDomesticPaymentConsentResponse, error) {
	var (
		introspectionResponse *openbanking.OpenbankingDomesticPaymentConsentIntrospectOK
		err                   error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if introspectionResponse, err = s.Client.Openbanking.OpenbankingDomesticPaymentConsentIntrospect(
		openbanking.NewOpenbankingDomesticPaymentConsentIntrospectParams().
			WithTid(s.Client.TenantID).
			WithAid(s.Client.ServerID).
			WithToken(&token),
		nil,
	); err != nil {
		logrus.WithError(err).Errorf("failed to introspect token %s", token)
		return nil, err
	}

	return introspectionResponse.Payload, nil
}

func has(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func newDateTimePtr(t time.Time) *strfmt.DateTime {
	str := strfmt.DateTime(t)
	return &str
}

func initiationsAreEqual(initiation1, initiation2 interface{}) bool {
	var (
		initiation1Bytes []byte
		initiation2Bytes []byte
		err              error
	)
	if initiation1Bytes, err = json.Marshal(initiation1); err != nil {
		return false
	}
	if initiation2Bytes, err = json.Marshal(initiation2); err != nil {
		return false
	}
	return bytes.Equal(initiation1Bytes, initiation2Bytes)
}

func toDomesticResponse5DataInitiation(initiation *acpClient.OBWriteDomesticConsentResponse5DataInitiation) *paymentModels.OBWriteDomesticResponse5DataInitiation {
	var (
		initiationBytes []byte
		err             error
		ret             paymentModels.OBWriteDomesticResponse5DataInitiation
	)

	if initiationBytes, err = json.Marshal(*initiation); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(initiationBytes, &ret); err != nil {
		panic(err)
	}

	return &ret
}

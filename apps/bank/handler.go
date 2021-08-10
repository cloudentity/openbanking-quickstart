package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EndpointLogicCommon interface {
	SetIntrospectionResponse(*gin.Context) error
	Validate(*gin.Context) error
	MapError(*gin.Context, error) interface{}
	GetUserIdentifier(*gin.Context) string
}

type GetEndpointLogic interface {
	BuildResponse(*gin.Context, BankUserData) interface{}
	Filter(*gin.Context, BankUserData) BankUserData
	EndpointLogicCommon
}

func (s *Server) Get(h GetEndpointLogic) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			data BankUserData
			err  error
		)

		if err = h.SetIntrospectionResponse(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		if err = h.Validate(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		if data, err = s.Storage.Get(h.GetUserIdentifier(c)); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		filtered := h.Filter(c, data)
		c.PureJSON(http.StatusOK, h.BuildResponse(c, filtered))
	}
}

type CreateEndpointLogic interface {
	SetRequest(*gin.Context) error
	BuildResource(*gin.Context) interface{}
	StoreResource(*gin.Context, string, interface{}) (interface{}, error)
	ResourceAlreadyExists(*gin.Context, string, interface{}) bool
	EndpointLogicCommon
}

func (s *Server) Post(h CreateEndpointLogic) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			resource interface{}
			stored   interface{}
			err      error
		)

		if err = h.SetRequest(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		if err = h.SetIntrospectionResponse(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		if err = h.Validate(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		resource = h.BuildResource(c)

		if h.ResourceAlreadyExists(c, h.GetUserIdentifier(c), resource) {
			c.PureJSON(http.StatusConflict, h.MapError(c, errors.New("resource already exists")))
		}

		if stored, err = h.StoreResource(c, h.GetUserIdentifier(c), resource); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		c.PureJSON(http.StatusCreated, stored)
	}
}

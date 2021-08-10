package main

import (
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
	CreateResource(*gin.Context, string) (interface{}, error)
	EndpointLogicCommon
}

func (s *Server) Post(h CreateEndpointLogic) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			created interface{}
			err     error
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

		if created, err = h.CreateResource(c, h.GetUserIdentifier(c)); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(c, err))
			return
		}

		c.PureJSON(http.StatusCreated, created)
	}
}

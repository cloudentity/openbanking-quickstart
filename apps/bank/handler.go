package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EndpointLogicCommon interface {
	SetIntrospectionResponse(*gin.Context) error
	Validate(*gin.Context) error
	MapError(*gin.Context, error) (int, interface{})
	GetUserIdentifier(*gin.Context) string
}

type GetEndpointLogicFactory func(server *Server) GetEndpointLogic

type GetEndpointLogic interface {
	BuildResponse(*gin.Context, BankUserData) interface{}
	Filter(*gin.Context, BankUserData) BankUserData
	EndpointLogicCommon
}

func (s *Server) Get(factory GetEndpointLogicFactory) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			h    = factory(s)
			data BankUserData
			err  error
		)

		if err = h.SetIntrospectionResponse(c); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.Validate(c); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if data, err = s.Storage.Get(h.GetUserIdentifier(c)); err != nil {
			_, errResp := h.MapError(c, err)
			c.PureJSON(http.StatusNotFound, errResp)
			return
		}

		filtered := h.Filter(c, data)
		c.PureJSON(http.StatusOK, h.BuildResponse(c, filtered))
	}
}

type CreateEndpointLogicFactory func(*Server) CreateEndpointLogic

type CreateEndpointLogic interface {
	SetRequest(*gin.Context) error
	CreateResource(*gin.Context, string) (interface{}, error)
	EndpointLogicCommon
}

func (s *Server) Post(factory CreateEndpointLogicFactory) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			h       = factory(s)
			created interface{}
			err     error
		)

		if err = h.SetRequest(c); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.SetIntrospectionResponse(c); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.Validate(c); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if created, err = h.CreateResource(c, h.GetUserIdentifier(c)); err != nil {
			code, errResp := h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		c.PureJSON(http.StatusCreated, created)
	}
}

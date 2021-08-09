package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EndpointLogic interface {
	SetRequest(c *gin.Context) error
	SetIntrospectionResponse(c *gin.Context) error
	MapError(err error) interface{}
	BuildResponse(BankUserData) interface{}
	Filter(data BankUserData) BankUserData
	Validate() error
	GetUserIdentifier(c *gin.Context) string
}

func (s *Server) Get(h EndpointLogic) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			data BankUserData
			err  error
		)

		if err = h.SetIntrospectionResponse(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		if err = h.Validate(); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		if data, err = s.Storage.Get(h.GetUserIdentifier(c)); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		filtered := h.Filter(data)
		c.PureJSON(http.StatusOK, h.BuildResponse(filtered))
	}
}

/*func (s *Server) Post(h EndpointLogic) func(*gin.Context) {
	return (c *gin.Context) {
		var (
			data BankUserData
			err error
		)


		if err = h.SetRequest(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		if err = h.SetIntrospectionResponse(c); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		if err = h.Validate(); err != nil {
			c.PureJSON(http.StatusBadRequest, h.MapError(err))
			return
		}

		// write

		// get


		c.PureJSON(http.StatusOK, h.BuildResponse(filtered))
	}
}*/

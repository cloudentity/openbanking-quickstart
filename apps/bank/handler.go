package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EndpointLogicCommon interface {
	SetIntrospectionResponse(*gin.Context) *Error
	Validate(*gin.Context) *Error
	MapError(*gin.Context, *Error) (int, interface{})
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
			h       = factory(s)
			data    BankUserData
			err     *Error
			e       error
			code    int
			errResp interface{}
		)

		logrus.Infof("XXX get %+v", c.Request.URL.Path)

		if err = h.SetIntrospectionResponse(c); err != nil {
			logrus.Info(fmt.Sprintf("err: %+v", err))
			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.Validate(c); err != nil {
			logrus.Info(fmt.Sprintf("err: %+v", err))

			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if data, e = s.Storage.Get(h.GetUserIdentifier(c)); e != nil {
			logrus.Info(fmt.Sprintf("err: %+v", err))

			code, errResp = h.MapError(c, ErrNotFound)
			c.PureJSON(code, errResp)
			return
		}

		logrus.Infof("XXX data %+v", data)

		filtered := h.Filter(c, data)

		logrus.Infof("XXX data filtered %+v", filtered)

		resp := h.BuildResponse(c, filtered)

		logrus.Infof("XXX resp: %+v", resp)

		c.PureJSON(http.StatusOK, resp)
	}
}

type CreateEndpointLogicFactory func(*Server) CreateEndpointLogic

type CreateEndpointLogic interface {
	SetRequest(*gin.Context) *Error
	CreateResource(*gin.Context, string) (interface{}, *Error)
	EndpointLogicCommon
}

func (s *Server) Post(factory CreateEndpointLogicFactory) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			h       = factory(s)
			created interface{}
			err     *Error
			code    int
			errResp interface{}
		)

		if err = h.SetRequest(c); err != nil {
			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.SetIntrospectionResponse(c); err != nil {
			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if err = h.Validate(c); err != nil {
			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		if created, err = h.CreateResource(c, h.GetUserIdentifier(c)); err != nil {
			code, errResp = h.MapError(c, err)
			c.PureJSON(code, errResp)
			return
		}

		c.PureJSON(http.StatusCreated, created)
	}
}

package main

import (
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
	var (
		h       = factory(s)
		handler = func(c *gin.Context, f func() (interface{}, *Error)) {
			var (
				resp interface{}
				err  *Error
			)
			if resp, err = f(); err != nil {
				code, errResp := h.MapError(c, err)
				logrus.WithField("response", errResp).Warnf("GET %s", c.FullPath())
				c.PureJSON(code, errResp)
				return
			}

			logrus.WithField("response", resp).Debugf("GET %s", c.FullPath())
			c.PureJSON(http.StatusOK, resp)
		}
	)

	return func(c *gin.Context) {
		handler(c, func() (interface{}, *Error) {
			var (
				data BankUserData
				err  *Error
				e    error
			)

			if err = h.SetIntrospectionResponse(c); err != nil {
				return nil, err
			}

			if err = h.Validate(c); err != nil {
				return nil, err
			}

			if data, e = s.Storage.Get(h.GetUserIdentifier(c)); e != nil {
				return nil, err
			}

			logrus.WithField("data", data).Debug("pulled data from database")

			filtered := h.Filter(c, data)
			response := h.BuildResponse(c, filtered)

			return response, nil
		})
	}
}

type CreateEndpointLogicFactory func(*Server) CreateEndpointLogic

type CreateEndpointLogic interface {
	SetRequest(*gin.Context) *Error
	CreateResource(*gin.Context, string) (interface{}, *Error)
	EndpointLogicCommon
}

func (s *Server) Post(factory CreateEndpointLogicFactory) func(*gin.Context) {
	var (
		h       = factory(s)
		handler = func(c *gin.Context, f func() (interface{}, *Error)) {
			var (
				resp interface{}
				err  *Error
			)
			if resp, err = f(); err != nil {
				code, errResp := h.MapError(c, err)
				logrus.WithField("response", errResp).Warnf("POST %s", c.FullPath())
				c.PureJSON(code, errResp)
				return
			}

			logrus.WithField("response", resp).Infof("POST %s", c.FullPath())
			c.PureJSON(http.StatusCreated, resp)
		}
	)

	return func(c *gin.Context) {
		handler(c, func() (interface{}, *Error) {
			var (
				created interface{}
				err     *Error
			)

			if err = h.SetRequest(c); err != nil {
				return nil, err
			}

			if err = h.SetIntrospectionResponse(c); err != nil {
				return nil, err
			}

			if err = h.Validate(c); err != nil {
				return nil, err
			}

			if created, err = h.CreateResource(c, h.GetUserIdentifier(c)); err != nil {
				return nil, err
			}

			return created, nil
		})
	}
}

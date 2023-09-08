package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Server) RenderInvalidRequestError(c *gin.Context, err error) {
	message := s.Trans.T("invalidRequest")

	s.RenderError(c, http.StatusBadRequest, message, err)
}

func (s *Server) RenderInternalServerError(c *gin.Context, err error) {
	message := s.Trans.T("internalServerError")

	s.RenderError(c, http.StatusInternalServerError, message, err)
}

func (s *Server) RenderError(c *gin.Context, statusCode int, msg interface{}, err error) {
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
	}

	c.HTML(statusCode, "error.tmpl", gin.H{
		"bank_logo": s.Config.BankLogo,
		"msg":       msg,
	})
}

func (s *Server) Render(c *gin.Context, templateName string, data map[string]interface{}) {
	data["bank_logo"] = s.Config.BankLogo
	c.HTML(http.StatusOK, templateName, data)
}

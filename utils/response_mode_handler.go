package utils

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type ResponseData struct {
	State            string `json:"state"`
	Code             string `json:"code"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (r *ResponseData) Valid() error {
	return nil
}

func HandleAuthResponseMode(r *http.Request) (ResponseData, error) {
	if r.URL.Query().Get("response") != "" {
		return GetResponseDataFromJWT(r)
	}

	if r.URL.Query().Get("code") != "" {
		return GetResponseDataFromQuery(r)
	}

	return ResponseData{}, errors.New("unable to determine response mode")
}

func GetResponseDataFromQuery(r *http.Request) (ResponseData, error) {
	return ResponseData{
		Code:             r.URL.Query().Get("code"),
		State:            r.URL.Query().Get("state"),
		Error:            r.URL.Query().Get("error"),
		ErrorDescription: r.URL.Query().Get("error_description"),
	}, nil
}

func GetResponseDataFromJWT(r *http.Request) (ResponseData, error) {
	var (
		responseData ResponseData
		token        = r.URL.Query().Get("response")
		parser       jwt.Parser
		err          error
	)

	if _, _, err = parser.ParseUnverified(token, &responseData); err != nil {
		return ResponseData{}, err
	}

	return responseData, nil
}

package shared

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/square/go-jose.v2"
)

type ResponseData struct {
	State            string `json:"state"`
	Code             string `json:"code"`
	Error            string `json:"error"`
	ErrorCause       string `json:"error_cause"`
	ErrorDescription string `json:"error_description"`
	TraceID          string `json:"trace_id"`
}

func (r *ResponseData) Valid() error {
	return nil
}

func HandleAuthResponseMode(r *http.Request, verificationKey jose.JSONWebKey) (ResponseData, error) {
	query := r.URL.Query()

	if query.Has("response") {
		return GetResponseDataFromJWT(r, verificationKey)
	}

	if query.Has("code") || query.Has("error") {
		return GetResponseDataFromQuery(r)
	}

	return ResponseData{}, errors.New("unable to determine response mode")
}

func GetResponseDataFromQuery(r *http.Request) (ResponseData, error) {
	return ResponseData{
		Code:             r.URL.Query().Get("code"),
		State:            r.URL.Query().Get("state"),
		Error:            r.URL.Query().Get("error"),
		ErrorCause:       r.URL.Query().Get("error_cause"),
		ErrorDescription: r.URL.Query().Get("error_description"),
		TraceID:          r.URL.Query().Get("trace_id"),
	}, nil
}

func GetResponseDataFromJWT(r *http.Request, key jose.JSONWebKey) (ResponseData, error) {
	var (
		responseData ResponseData
		token        = r.URL.Query().Get("response")
		parser       jwt.Parser
		keyFunc      = func(t *jwt.Token) (interface{}, error) {
			return key.Key, nil
		}
		err error
	)

	if _, err = parser.ParseWithClaims(token, &responseData, keyFunc); err != nil {
		return ResponseData{}, err
	}

	return responseData, nil
}

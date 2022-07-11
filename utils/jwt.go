package utils

import (
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

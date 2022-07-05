package utils

import "net/http"

func GetResponseDataFromQuery(r *http.Request) (ResponseData, error) {
	return ResponseData{
		Code:             r.URL.Query().Get("code"),
		State:            r.URL.Query().Get("state"),
		Error:            r.URL.Query().Get("error"),
		ErrorDescription: r.URL.Query().Get("error_description"),
	}, nil
}

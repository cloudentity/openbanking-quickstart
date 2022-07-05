package utils

import "net/http"

type CallbackURLResponseModeParser func(*http.Request) (ResponseData, error)

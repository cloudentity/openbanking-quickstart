package main

import (
	"errors"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/sirupsen/logrus"
)

type HTTPRuntime struct {
	*client.Runtime
}

func (r *HTTPRuntime) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	logrus.WithField("operation", operation).Info("submitted operation")
	var selectedMediaType string
	firstMediaType := operation.ConsumesMediaTypes[0]
	for _, mediaType := range operation.ConsumesMediaTypes {
		if _, ok := r.Producers[mediaType]; ok {
			selectedMediaType = mediaType
			break
		}
	}

	if selectedMediaType == "" && firstMediaType != runtime.MultipartFormMime && firstMediaType != runtime.URLencodedFormMime {
		return nil, errors.New("failed to select producer")
	} else if selectedMediaType == "" {
		selectedMediaType = firstMediaType
	}

	operation.ConsumesMediaTypes = []string{selectedMediaType}

	if _, ok := r.Consumers[selectedMediaType]; ok {
		// content type selected for msg creation has priority for response
		operation.ProducesMediaTypes = []string{selectedMediaType}
	} else {
		for _, mediaType := range operation.ProducesMediaTypes {
			if _, ok := r.Consumers[mediaType]; ok {
				selectedMediaType = mediaType
				break
			}
		}

		if selectedMediaType == "" {
			return nil, errors.New("failed to select consumer")
		}

		operation.ProducesMediaTypes = []string{selectedMediaType}
	}

	return r.Runtime.Submit(operation)
}

func NewHTTPRuntimeWithClient(host, basePath string, schemes []string, c *http.Client) *HTTPRuntime {
	return &HTTPRuntime{
		client.NewWithClient(host, basePath, schemes, c),
	}
}

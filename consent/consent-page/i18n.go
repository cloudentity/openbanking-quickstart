package main

import (
	"html/template"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"
)

type Trans struct {
	*i18n.Localizer
}

func NewTranslations(bundle *i18n.Bundle, lang string) *Trans {
	return &Trans{i18n.NewLocalizer(bundle, lang)}
}

type Option func(i interface{}) interface{}

func AsHTML(i interface{}) interface{} {
	return template.HTML(i.(string))
}

func (t *Trans) T(ID string, options ...Option) interface{} {
	var result interface{}
	message := t.MustLocalize(&i18n.LocalizeConfig{
		MessageID: ID,
		DefaultMessage: &i18n.Message{
			ID: ID,
			Other: ID,
		},
	})

	logrus.WithField("message", message).
		Debugf("tried to find message %s", ID)

	result = message

	for _, opt := range options {
		result = opt(result)
	}

	return result
}

func (t *Trans) TD(ID string, data map[string]interface{}, options ...Option) interface{} {
	var result interface{}
	message := t.MustLocalize(&i18n.LocalizeConfig{
		MessageID: ID,
		TemplateData: data,
		DefaultMessage: &i18n.Message{
			ID: ID,
			Other: ID,
		},
	})

	logrus.WithField("message", message).
		WithField("data", data).
		Debugf("tried to find message %s", ID)

	result = message

	for _, opt := range options {
		result = opt(result)
	}

	return result
}

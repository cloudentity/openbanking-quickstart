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

func (t *Trans) T(id string) interface{} {
	message := t.MustLocalize(&i18n.LocalizeConfig{
		MessageID: id,
		DefaultMessage: &i18n.Message{
			ID:    id,
			Other: id,
		},
	})

	logrus.WithField("message", message).
		Debugf("tried to find message %s", id)

	return template.HTML(message) // nolint
}

func (t *Trans) TD(id string, data map[string]interface{}) interface{} {
	message := t.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
		DefaultMessage: &i18n.Message{
			ID:    id,
			Other: id,
		},
	})

	logrus.WithField("message", message).
		WithField("data", data).
		Debugf("tried to find message %s", id)

	return template.HTML(message) // nolint
}

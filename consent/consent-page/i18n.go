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

func (t *Trans) T(ID string) interface{} {
	message := t.MustLocalize(&i18n.LocalizeConfig{
		MessageID: ID,
		DefaultMessage: &i18n.Message{
			ID: ID,
			Other: ID,
		},
	})

	logrus.WithField("message", message).
		Debugf("tried to find message %s", ID)

	return template.HTML(message)
}

func (t *Trans) TD(ID string, data map[string]interface{}) interface{} {
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

	return template.HTML(message)
}

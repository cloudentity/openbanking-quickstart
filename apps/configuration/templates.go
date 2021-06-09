package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	sprig "github.com/Masterminds/sprig/v3"
)

type YamlFile map[string]interface{}

type Templates struct {
	m map[string][]byte
}

func LoadTemplates(dirs []string, variablesFile *string) (Templates, error) {
	var (
		templates = Templates{
			m: map[string][]byte{},
		}
		files     []os.FileInfo
		bs        []byte
		variables = YamlFile{}
		t         *template.Template
		err       error
	)

	if variablesFile != nil && *variablesFile != "" {
		if bs, err = ioutil.ReadFile(*variablesFile); err != nil {
			return templates, err
		}

		if err = yaml.Unmarshal(bs, &variables); err != nil {
			return templates, errors.Wrapf(err, "failed to unmarshal env file to yaml")
		}

		logrus.Debugf("variables: %+v", variables)
	}

	for _, d := range dirs {
		if files, err = ioutil.ReadDir(d); err != nil {
			return templates, errors.Wrapf(err, "failed to read templates dir: %s", d)
		}

		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".tmpl") {
				logrus.Debugf("file skipped: %+v", f.Name())
				continue
			}

			file := fmt.Sprintf("%s/%s", d, f.Name())
			logrus.Debugf("read file: %+v", file)

			if bs, err = ioutil.ReadFile(file); err != nil {
				return templates, errors.Wrapf(err, "failed to read template: %s", file)
			}

			if t, err = template.New(file).Funcs(FuncMap()).Parse(string(bs)); err != nil {
				return templates, errors.Wrapf(err, "failed to parse template: %s", file)
			}

			var buf bytes.Buffer
			if err = t.Execute(&buf, variables); err != nil {
				return templates, errors.Wrapf(err, "failed to render template: %s", file)
			}

			templates.m[file] = buf.Bytes()
		}
	}

	return templates, nil
}

func (t Templates) Merge() (YamlFile, error) {
	var (
		yamlFile = YamlFile{}
		err      error
	)

	for t, v := range t.m {
		var tmp YamlFile

		if err = yaml.Unmarshal(v, &tmp); err != nil {
			return yamlFile, errors.Wrapf(err, "failed to unmarshal template: %s to yaml", t)
		}

		logrus.Debugf("merge file: %s", t)

		if err = mergo.Merge(&yamlFile, &tmp, mergo.WithAppendSlice); err != nil {
			return yamlFile, errors.Wrapf(err, "failed to merge template: %s", t)
		}
	}

	logrus.Debugf("final yaml file: \n %s", yamlFile)

	return yamlFile, nil
}

func (y YamlFile) ToJSON() ([]byte, error) {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&y); err != nil {
		return bs, errors.Wrapf(err, "failed to marshal yaml to json")
	}

	return bs, nil
}

func FuncMap() template.FuncMap {
	extra := template.FuncMap{
		"toYaml":     toYAML,
		"readFile":   readFile,
		"jwksEncode": jwksEncode,
	}

	// merge with sprig
	f := sprig.TxtFuncMap()
	for k, v := range extra {
		f[k] = v
	}

	return f
}

func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// ignore
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

func readFile(v string) string {
	data, err := ioutil.ReadFile(v)
	if err != nil {
		// ignore
		return ""
	}
	return string(data)
}

func jwksEncode(v string) string {
	pemblock, _ := pem.Decode([]byte(v))
	if pemblock == nil {
		return ""
	}

	parsed, err := x509.ParsePKCS8PrivateKey(pemblock.Bytes)
	if err != nil {
		return ""
	}

	key, err := jwk.New(parsed)
	if err != nil {
		return ""
	}

	output, _ := yaml.Marshal(key)
	return string(output)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
)

type YamlFile map[string]interface{}

type Templates struct {
	m map[string][]byte
}

func LoadTemplates(dir string, variablesFile *string) (Templates, error) {
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

	if variablesFile != nil {
		if bs, err = ioutil.ReadFile(*variablesFile); err != nil {
			return templates, err
		}

		if err = yaml.Unmarshal(bs, &variables); err != nil {
			return templates, err
		}
	}

	if files, err = ioutil.ReadDir(dir); err != nil {
		return templates, err
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".tmpl") {
			continue
		}

		file := fmt.Sprintf("%s/%s", dir, f.Name())
		if bs, err = ioutil.ReadFile(file); err != nil {
			return templates, err
		}

		if t, err = template.New(file).Parse(string(bs)); err != nil {
			return templates, err
		}

		var buf bytes.Buffer
		if err = t.Execute(&buf, variables); err != nil {
			return templates, err
		}

		templates.m[file] = buf.Bytes()
	}

	return templates, nil
}

func (t Templates) Merge() (YamlFile, error) {
	var (
		yamlFile = YamlFile{}
		err      error
	)

	for _, v := range t.m {
		var tmp YamlFile

		if err = yaml.Unmarshal(v, &tmp); err != nil {
			return yamlFile, err
		}

		if err = mergo.Merge(&yamlFile, &tmp, mergo.WithAppendSlice); err != nil {
			return yamlFile, err
		}
	}

	return yamlFile, nil
}

func (y YamlFile) ToJSON() ([]byte, error) {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&y); err != nil {
		return bs, err
	}

	return bs, nil
}

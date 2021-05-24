package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ImportConfiguration(iss *url.URL, client *http.Client, body []byte, mode string) error {
	var (
		req  *http.Request
		resp *http.Response
		bs   []byte
		err  error
	)

	parts := strings.Split(iss.Path, "/")
	if len(parts) < 3 {
		return fmt.Errorf("tenant/server must be present in the issuer url")
	}

	tenant := parts[1]
	server := parts[2]

	if tenant == "system" {
		iss.Path = "/api/system/configuration"
	} else {
		iss.Path = fmt.Sprintf("/api/system/%s/configuration", tenant)
	}

	if server != "system" {
		return fmt.Errorf("system server must be used in the issuer url")
	}

	if mode != "" {
		iss.RawQuery = fmt.Sprintf("mode=%s", mode)
	}

	if req, err = http.NewRequest("PUT", iss.String(), bytes.NewBuffer(body)); err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")

	if resp, err = client.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		if bs, err = ioutil.ReadAll(resp.Body); err != nil {
			return err
		}

		return fmt.Errorf("import endpoint: %s returned invalid status code: %d, body: %s", iss.String(), resp.StatusCode, bs)
	}

	return nil
}

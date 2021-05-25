package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const systemTenant = "system"

func ImportConfiguration(tenantURL *url.URL, tenant *string, client *http.Client, body []byte, mode string) error {
	var (
		req  *http.Request
		resp *http.Response
		bs   []byte
		err  error
	)

	if *tenant == systemTenant {
		tenantURL.Path = "/api/system/configuration"
	} else {
		tenantURL.Path = fmt.Sprintf("/api/system/%s/configuration", *tenant)
	}

	if mode != "" {
		tenantURL.RawQuery = fmt.Sprintf("mode=%s", mode)
	}

	if req, err = http.NewRequest("PUT", tenantURL.String(), bytes.NewBuffer(body)); err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")

	if resp, err = client.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		if bs, err = ioutil.ReadAll(resp.Body); err != nil {
			return err
		}

		return fmt.Errorf("import endpoint: %s returned invalid status code: %d, body: %s", tenantURL.String(), resp.StatusCode, bs)
	}

	return nil
}

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	TenantID          string `env:"TENANT"`
	AdminClientID     string `env:"ADMIN_CLIENT_ID"`
	AdminClientSecret string `env:"ADMIN_CLIENT_SECRET"`
}

func main() {
	var (
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
		config       Config
	)

	if config, err = LoadConfig(); err != nil {
		log.Fatalf("failed to load env %+v", err)
	}

	tenantURLRaw = fmt.Sprintf("https://%s.authz.cloudentity.io", config.TenantID)

	if tURL, err = url.Parse(tenantURLRaw); err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // nolint
			},
		},
	}

	cc := clientcredentials.Config{
		ClientID:     "c79lsrgh5kre3dfd8kf0",
		ClientSecret: "S4DYjFEowDmEKfwbXOtR-mqaHWuIae2Mt4i-6KimZYQ",
		TokenURL:     fmt.Sprintf("%s/%s/%s/oauth2/token", tURL.String(), config.TenantID, "admin"),
	}

	client := cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, httpClient))

	workspaceIDs := []string{
		"openbanking_brasil",
		"openbanking",
		"bank-admins",
		"bank-customers",
		"financroo",
	}

	for _, wid := range workspaceIDs {
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), config.TenantID, wid), nil); err != nil {
			log.Fatalf("failed to create server delete request: %v", err)
		}
		if response, err = doRequest(client, request); err != nil {
			log.Fatalf("failed to delete server: %v", err)
		}
		response.Body.Close()
	}

	clientIDs := []string{
		"c79lsrgh5kre3dfd8kf0",
	}

	for _, cid := range clientIDs {
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/clients/%s", tURL.String(), config.TenantID, cid), nil); err != nil {
			log.Fatalf("failed to create client delete request")
		}
		if response, err = doRequest(client, request); err != nil {
			log.Fatalf("failed to delete client: %v", err)
		}
		response.Body.Close()
	}
}

func doRequest(client *http.Client, request *http.Request) (response *http.Response, err error) {
	if response, err = client.Do(request); err != nil {
		return response, err
	}

	if response.StatusCode != http.StatusNoContent {
		return response, fmt.Errorf("expected response code %d, got %d", http.StatusNoContent, response.StatusCode)
	}

	return response, nil
}

func LoadConfig() (config Config, err error) {
	err = env.Parse(&config)
	return config, err
}

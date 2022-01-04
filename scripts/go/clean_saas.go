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
		ClientID:     config.AdminClientID,
		ClientSecret: config.AdminClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/%s/oauth2/token", tURL.String(), config.TenantID, "admin"),
	}

	client := cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, httpClient))

	var (
		workspaceIDs = []string{
			"openbanking_brasil",
			"openbanking",
			"bank-admins",
			"bank-customers",
			"financroo",
		}
	)

	for _, wid := range workspaceIDs {
		func() {
			if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), config.TenantID, wid), nil); err != nil {
				log.Fatalf("failed to create delete http request: %v", err)
			}

			if response, err = client.Do(request); err != nil {
				log.Fatalf("http request failed: %+v", err)
			}

			if response.StatusCode != http.StatusNoContent {
				log.Fatalf("expected response code %d, got %d", http.StatusNoContent, response.StatusCode)
			}
			defer response.Body.Close()
		}()
	}

}

func LoadConfig() (config Config, err error) {
	err = env.Parse(&config)
	return config, err
}

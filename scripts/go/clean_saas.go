package main

import (
	"context"
	"crypto/tls"
	"flag"
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
	AdminClientID     string `env:"CLEANUP_CLIENT_ID"`
	AdminClientSecret string `env:"CLEANUP_CLIENT_SECRET"`
}

func main() {
	var (
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
		config       Config
		workspaceIDs []string
	)

	if config, err = LoadConfig(); err != nil {
		log.Fatalf("failed to load env %+v", err)
	}

	tenantURLRaw = fmt.Sprintf("https://%s.us.authz.cloudentity.io", config.TenantID)

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

	obSpec := flag.String("spec", "none", "Openbanking quickstart specification type")
	flag.Parse()

	switch *obSpec {
	case "obuk":
		workspaceIDs = []string{
			"openbanking",
			"bank-customers",
		}
	case "obbr":
		workspaceIDs = []string{
			"openbanking_brasil",
			"bank-customers",
		}
	case "cdr":
		workspaceIDs = []string{
			"cdr",
			"bank-customers",
		}
	default:
		log.Fatalf("The openbanking specification flag '-spec=%s' is not supported", *obSpec)
	}

	for _, wid := range workspaceIDs {
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), config.TenantID, wid), nil); err != nil {
			log.Fatalf("failed to setup delete server '%s' request: %v", wid, err)
		}
		if response, err = doRequest(client, request); err != nil {
			log.Fatalf("failed to delete server '%s': %v", wid, err)
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

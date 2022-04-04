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

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	obSpec            = flag.String("spec", "none", "Openbanking quickstart specification type")
	tenantID          = flag.String("tenant", "none", "Openbanking SaaS tenant ID")
	adminClientID     = flag.String("cid", "none", "Openbanking SaaS admin client ID")
	adminClientSecret = flag.String("csec", "none", "Openbanking SaaS admin client secret")
)

func main() {
	flag.Parse()

	var (
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
		workspaceIDs []string
	)

	tenantURLRaw = fmt.Sprintf("https://%s.us.authz.cloudentity.io", *tenantID)

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
		ClientID:     *adminClientID,
		ClientSecret: *adminClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/%s/oauth2/token", tURL.String(), *tenantID, "admin"),
	}

	client := cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, httpClient))

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
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), *tenantID, wid), nil); err != nil {
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

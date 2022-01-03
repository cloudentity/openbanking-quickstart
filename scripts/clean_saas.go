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

func main() {
	var (
		clientID     = flag.String("client_id", "", "admin client id")
		clientSecret = flag.String("client_secret", "", "admin client secret")
		tenantID     = flag.String("tenant_id", "", "saas tenant id")
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
	)

	flag.Parse()

	tenantURLRaw = fmt.Sprintf("https://%s.authz.cloudentity.io", *tenantID)

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
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/%s/oauth2/token", tURL.String(), *tenantID, "admin"),
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
			if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), *tenantID, wid), nil); err != nil {
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

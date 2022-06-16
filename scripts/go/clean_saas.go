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
	tenantID          = flag.String("tenant", "none", "Openbanking SaaS tenant ID")
	adminClientID     = flag.String("cid", "none", "Openbanking SaaS admin client ID")
	adminClientSecret = flag.String("csec", "none", "Openbanking SaaS admin client secret")

	openbankingClientsIDs = []string{"buc3b1hhuc714r78env0", "bv2fe0tpfc67lmeti340", "bv0ocudfotn6edhsiu7g"}
	openbankingServersIDs = []string{"cdr", "fdx", "openbanking", "openbanking_brasil", "bank-customers"}
)

func main() {
	flag.Parse()

	var (
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
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

	for _, sid := range openbankingServersIDs {
		fmt.Printf("INFO: Trying to delete server with ID: '%s'\n", sid)
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), *tenantID, sid), nil); err != nil {
			log.Fatalf("ERROR: Failed to setup delete server '%s' request: %v", sid, err)
		}

		if response, err = doRequest(client, request, http.StatusNoContent); err != nil {
			log.Fatalf("ERROR: Failed to delete server '%s': %v", sid, err)
		}

		response.Body.Close()
	}

	for _, cid := range openbankingClientsIDs {
		fmt.Printf("INFO: Trying to delete client with ID: '%s'\n", cid)
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/clients/%s", tURL.String(), *tenantID, cid), nil); err != nil {
			log.Fatalf("ERROR: Failed to setup delete client '%s' request: %v", cid, err)
		}

		if response, err = doRequest(client, request, http.StatusNoContent); err != nil {
			log.Fatalf("ERROR: Failed to delete client '%s': %v", cid, err)
		}

		response.Body.Close()
	}
}

func doRequest(client *http.Client, request *http.Request, statusCode int) (response *http.Response, err error) {
	if response, err = client.Do(request); err != nil {
		return response, err
	}

	if http.StatusNotFound == response.StatusCode {
		fmt.Printf("INFO: The response finished with status code '%d'\n", response.StatusCode)
		return response, nil
	} else if response.StatusCode != statusCode {
		fmt.Printf("INFO: The response finished with status code '%d'\n", response.StatusCode)
	}
	fmt.Printf("INFO: The response finished with status code '%d'\n", response.StatusCode)
	return response, nil
}

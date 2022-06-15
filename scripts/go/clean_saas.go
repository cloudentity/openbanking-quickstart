package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	tenantID              = flag.String("tenant", "none", "Openbanking SaaS tenant ID")
	adminClientID         = flag.String("cid", "none", "Openbanking SaaS admin client ID")
	adminClientSecret     = flag.String("csec", "none", "Openbanking SaaS admin client secret")
	
	defaultServicesIDs    = []string{"system", "admin", "default", "developer", "demo"}
	openbankingClientsIDs = []string{"buc3b1hhuc714r78env0", "bv2fe0tpfc67lmeti340", "bv0ocudfotn6edhsiu7g"}
)

type Server struct {
	ID string `json:"id"`
}
type ServersResponse struct {
	Servers []Server `json:"servers"`
}

func main() {
	flag.Parse()

	var (
		request      *http.Request
		response     *http.Response
		err          error
		tURL         *url.URL
		tenantURLRaw string
		serversIDs []string
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

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/api/admin/%s/servers", tURL.String(), *tenantID), nil); err != nil {
		log.Fatalf("failed to setup get servers request: %v", err)
	}

	if response, err = doRequest(client, request, http.StatusOK); err != nil {
		log.Fatalf("failed to get servers: %v", err)
	}

	if serversIDs, err = getCurrentServersIDs(response); err != nil {
		log.Fatalf("failed to get servers IDs: %v", err)
	}

	serversIDs = getServersIDsToDelete(serversIDs)

	for _, sid := range serversIDs {
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/servers/%s", tURL.String(), *tenantID, sid), nil); err != nil {
			log.Fatalf("failed to setup delete server '%s' request: %v", sid, err)
		}

		if response, err = doRequest(client, request, http.StatusNoContent); err != nil {
			log.Fatalf("failed to delete server '%s': %v", sid, err)
		}

		response.Body.Close()
		fmt.Printf("INFO: server with ID: '%s' was successfully removed\n", sid)
	}

	for _, cid := range openbankingClientsIDs {
		if request, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/%s/clients/%s", tURL.String(), *tenantID, cid), nil); err != nil {
			log.Fatalf("failed to setup delete client '%s' request: %v", cid, err)
		}

		if response, err = doRequest(client, request, http.StatusNoContent); err != nil {
			log.Fatalf("failed to delete client '%s': %v", cid, err)
		}

		response.Body.Close()
		fmt.Printf("INFO: clientwith ID: '%s' was successfully removed\n", cid)
	}

}

func doRequest(client *http.Client, request *http.Request, statusCode int) (response *http.Response, err error) {
	if response, err = client.Do(request); err != nil {
		return response, err
	}

	if response.StatusCode != statusCode {
		return response, fmt.Errorf("expected response code %d, got %d", statusCode, response.StatusCode)
	}

	return response, nil
}

func getCurrentServersIDs(response *http.Response) (IDs []string, err error) {
	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		log.Fatalf("failed to get body: %v", err)
	}

	response.Body.Close()

	var serversResponse ServersResponse
	json.Unmarshal([]byte(body), &serversResponse)

	for _, server := range serversResponse.Servers {
		IDs = append(IDs, server.ID)
	}

	return IDs, nil
}

func getServersIDsToDelete(actualIDs []string) (expectedIDs []string) {
	tmpServersIDsMap := make(map[string]bool)
	for _, i := range actualIDs {
		tmpServersIDsMap[i] = true
	}

	for _, i := range defaultServicesIDs {
		delete(tmpServersIDsMap, i)
	}

	for key := range tmpServersIDsMap {
		expectedIDs = append(expectedIDs, key)
	}

	return expectedIDs
}

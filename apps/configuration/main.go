package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	tenantURL     = flag.String("tenant-url", "https://localhost:8443/system", "tenant url")
	tenant        = flag.String("tenant", "default", "tenant id")
	clientID      = flag.String("client-id", "system", "client id")
	clientSecret  = flag.String("client-secret", "n8HF35qzZkmsukHJzzz9LnN8m9Mf97uq", "client secret")
	templatesDir  = flag.String("templates-dir", "", "templates directory in yaml format")
	variablesFile = flag.String("variables-file", "", "optional variables file in yaml format that can be used in templates")
	httpTimeout   = flag.Duration("http-timeout", time.Second*10, "http timeout")
	httpInsecure  = flag.Bool("http-insecure", true, "http insecure connection")
	importMode    = flag.String("import-mode", "update", "how acp should behave in case of conflicts, possible options: fail | ignore | update")
)

const systemServer = "system"

func initialize() {
	flag.Parse()

	if *templatesDir == "" {
		log.Printf("templates dir must be set, see usage below")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	initialize()

	var (
		err       error
		templates Templates
		yamlFile  YamlFile
		body      []byte
		tURL      *url.URL
	)

	if tURL, err = url.Parse(*tenantURL); err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: *httpTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: *httpInsecure, // nolint
			},
		},
	}

	cc := clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/oauth2/token", tURL.String(), systemServer),
	}

	client := cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, httpClient))

	if templates, err = LoadTemplates(*templatesDir, variablesFile); err != nil {
		log.Fatal(err)
	}

	if yamlFile, err = templates.Merge(); err != nil {
		log.Fatal(err)
	}

	if body, err = yamlFile.ToJSON(); err != nil {
		log.Fatal(err)
	}

	if err = ImportConfiguration(tURL, tenant, client, body, *importMode); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Configuration imported")
}

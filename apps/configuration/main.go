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
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	tenantURL     = flag.String("tenant-url", "https://localhost:8443/system", "tenant url")
	tenant        = flag.String("tenant", "default", "tenant id")
	clientID      = flag.String("client-id", "system", "client id")
	clientSecret  = flag.String("client-secret", "n8HF35qzZkmsukHJzzz9LnN8m9Mf97uq", "client secret")
	templatesDirs = flag.String("templates-dirs", "", "comma separated path to directories containing import templates with tmpl extension (yaml format)")
	variablesFile = flag.String("variables-file", "", "optional variables file in yaml format that can be used in templates")
	httpTimeout   = flag.Duration("http-timeout", time.Second*10, "http timeout")
	httpInsecure  = flag.Bool("http-insecure", true, "http insecure connection")
	importMode    = flag.String("import-mode", "update", "how acp should behave in case of conflicts, possible options: fail | ignore | update")
	verbose       = flag.Bool("verbose", false, "show verbose logs")
	target        = flag.String("target", "system", "where the data should be imported, one of: system | identity")
)

const systemServer = "system"

func initialize() {
	flag.Parse()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if *templatesDirs == "" {
		logrus.Info("templates dirs must be set, see usage below")
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
				InsecureSkipVerify: *httpInsecure, //nolint
			},
		},
	}

	cc := clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/oauth2/token", tURL.String(), systemServer),
	}

	client := cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, httpClient))

	dirs := strings.Split(*templatesDirs, ",")

	if templates, err = LoadTemplates(dirs, variablesFile); err != nil {
		log.Fatal(err)
	}

	if yamlFile, err = templates.Merge(); err != nil {
		log.Fatal(err)
	}

	if body, err = yamlFile.ToJSON(); err != nil {
		log.Fatal(err)
	}

	if err = ImportConfiguration(tURL, tenant, client, body, *importMode, *target); err != nil {
		log.Fatal(err)
	}

	logrus.Info("Configuration imported")
}

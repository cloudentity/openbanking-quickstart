# About

A configuration app is a helper tool that allows importing data to the ACP.
It's useful when you want to add a new application, change idp, or openbanking consent page.


# Prerequisites

The configuration app uses ACP [system import API]( https://docs.authorization.cloudentity.com/api/system/#operation/importConfiguration)
This endpoint must be explicitly enabled on the ACP side (see ACP configuration reference guide: https://docs.authorization.cloudentity.com/reference/configuration/)

To be able to authenticate to the import endpoint you need an oauth client created in the system workspace.
This client should be able to get token using client credentials flow and issue scope: `manage_configuration`.
When you run acp with flag responsible for creating a default tenant, this client is created automatically with a default set of credentials.

# How it works

The configuration app scans provided directories and look for import template files.
Each file is a [golang template](https://golang.org/pkg/text/template)
The configuration app renders each template which should produce a valid yaml.
Next, all yamls are concatenated and send as a body to ACP [import API](https://docs.authorization.cloudentity.com/api/system/#operation/importConfiguration)
You can optionally specify a path to variables file which could be used in the templates.

# Sample usage

You can run a docker:

```
docker run cloudentity/openbanking-quickstart-configuration:latest
```

and then a help will be disabled with all the available options.

Or you can preview `configuration` docker container in the [docker-compose](../../docker-compose.yaml)

# How to add a new template

Create a new file with the `.tmpl` extension under directory which is used by the configuration app.
In case of openbanking quickstart this directory in located here: [Templates dir](../../data/imports)

# How to define a new variable

You can add new variables to [variables file](../../data/variables.yaml).
The file itself is not a golang template and should be a valid yaml file.

# How to use variable in the template

Assumming that you have defined a new variable under in [variables.yaml](../../data/variables.yaml):

```
key=value
```

Then in the template you can use it as follows:

``` sh
{{ .key }}
```

# How to add a new application

To add a new application you need to define a new template and use it in (the configuration app)[../apps/configuration/README.md].

First, take a look at clients schema in the import configuration request body: https://docs.authorization.cloudentity.com/api/system/operation/importConfiguration

If you want to create application which uses client credentials flow the template may look like this: 

```
clients:
- tenant_id: default
  authorization_server_id: openbanking
  client_id: test
  client_name: client credentials test
  client_secret: xMPBmv62z3Jt1S4sWl2qRhOhEGPVZ9EcujGL7Xy0-E1
  grant_types:
  - client_credentials
  token_endpoint_auth_method: client_secret_basic
```

If you you would like to create web application, you could use:

```
clients:
- tenant_id: default
  authorization_server_id: openbanking
  client_id: web
  client_name: web client test
  client_secret: -TlfoycUiE0qas-XUBFDfTxMlhHTCjVxOF6pLrWZbQA
  token_endpoint_auth_method: client_secret_basic
  redirect_uris:
  - {{ .example_web_client_url }}/callback
```

Please note that `example_web_client_url` must be defined in the variables file. 

# How to change IDP

To change the IDP of the existing workspace (in this case `openbanking`) you need to define a new template and use it in (the configuration app)[../apps/configuration/README.md].

First, take a look at idp schema in the ACP import configuration request body: https://docs.authorization.cloudentity.com/api/system/#operation/importConfiguration

Below you can find a sample template if you would like to change the IDP to `custom`:

```
idps:
- tenant_id: default
  authorization_server_id: openbanking
  id: bugkgai3g9kregtu04u0 # id of existing idp
  name: IDP name
  method: custom # idp type
  settings:
    custom: # idp custom settings
      login_url: "http://example.com"
  credentials:
```

Please note that depending on the IDP type that you want to use, the import template will look slightly different.
For instance, if you set idp method to `azure`, then you need to provide `settings.azure` and `credentials.azure` configuration.

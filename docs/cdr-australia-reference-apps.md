# Integrate CDR reference apps with Cloudentity Authorization Platform

In this article we will integrate [CDR reference applications](https://github.com/ConsumerDataRight) with [Cloudentity authorization platform](https://cloudentity.com/)
to provide CDR data authorization management capabilities for `Data Holder` and `Data recipient` platforms to be CDR compatible.

**Cloudentity authorization platform** is an [Australia CDR(Based on FAPI 1 Advanced Final profile) OpenID certfied platform](https://openid.net/certification/#FAPI_OPs). Cloudentity's advanced authorization platform is also a major leading platform
in other open data initiatives world wide like **OpenBanking UK**, **OpenBanking Brazil** etc and is also certified in those initiatives.

Cloudentity authorization platform can easily augment and enable Australian CDR Data Holders to be CDR compatible in a seamless fashion. Cloudentity authorization platform(ACP) can provide below functionalities:
* Dynamic Client Registration for Data recipients
* Act as authorization server to issue accessTokens complaint to CDR specification
* Provide token grant, revoke, arrangement revoke capabilities
* Consent grant/deny/management capabilities and consent ledger API's
* Centralized dynamic policy administration, management for access
* Advanced analytics on data recipient to holder integrations, usage in relation to data authorization and access requests.
* Customizable integration pattern to have custom Consent UI apps that can integrate with Data holders and Cloudentity Authorization platform to drive user experience for consent collection and management.

![Cloudentity-cdr-reference-apps-intgration](./Cloudentity-cdr-reference-apps-intgration.jpeg)

Cloudentity authorization platform is available as
* SaaS offering
* On prem deployment offering

In this article we will be walking through connecting to Cloudentity authorization platform via both the options. You may choose and proceed with one of the option for quick verification of the platform capabilities.


For the demonstration in this articles, we will use following software components:

* `mock-registry` - https://github.com/ConsumerDataRight/mock-register
* `mock-data-holder` - https://github.com/cloudentity/mock-data-holder
* `mock-data-recipient` - https://github.com/cloudentity/mock-data-recipient

Above reference applications have a built in Identity Server implementation but we will be swapping that out with a more powerful, flexible and dynamic platform offered by Cloudentity.

## Git clone CDR openbanking quickstart

```
git clone git@github.com:cloudentity/openbanking-quickstart.git
```

## Pick an platform deployment style

Before proceeding further, decide on whether you want to utilize Cloudentity platform in SaaS(preferred) or deploy
Cloudentity platform in a machine locally.

. ACP SaaS

Attached are the components that will be deployed locally and rest will be on SaaS platform offered by Cloudentity

![Cloudentity-cdr-reference-apps-saas](./cdr-app-ce-saas.jpeg)

. ACP local

All attached components will be deployed on local network.


![Cloudentity-cdr-reference-apps-saas](./cdr-app-ce-local.jpeg)

> ## NOTE
> For local deployment, use a domain like `authorization.cloudentity.com` and make a host entry in local file system for
> `127.0.0.1 authorization.cloudentity.com`
>
> To use the Cloudentity SaaS platform, [signup for a free tenant](https://authz.cloudentity.io/) at 
> https://authz.cloudentity.io/ and utilize the tenant url and endpoints in below configuration

### Configure mock-data-recipient with Cloudentity authorization Platform

As you see in above attached reference diagram, the data holder is augmented with Cloudentity authorization platform 
and to configure the `mock-data-recipient`, following configuration nedds to be overriden in the application.
Override below settings in the file under this location https://github.com/cloudentity/openbanking-quickstart/blob/master/mount/cdr/recipient.json

. ACP SaaS (sample)

```
{
    ..
  "MockDataRecipient": {
      ..
    "DataHolder": {
      "infosecBaseUri": "https://authorization.cloudentity.com:8443/default/cdr",
      "jwksUri": "https://authorization.cloudentity.com:8443/default/cdr/.well-known/jwks.json",
      "registrationEndpoint": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/register"
    }
  }
}
```

. ACP Local

```
{
    ..
  "MockDataRecipient": {
      ..
    "DataHolder": {
      "infosecBaseUri": "https://authorization.cloudentity.com:8443/default/cdr",
      "jwksUri": "https://authorization.cloudentity.com:8443/default/cdr/.well-known/jwks.json",
      "registrationEndpoint": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/register"
    }
  }
}
```




### Configure mock-data-holder with Cloudentity authorization Platform

As you see in above attached reference diagram, the data holder is augmented with Cloudentity authorization platform 
and to configure the `mock-data-recipient`, following configuration nedds to be overriden in the application at
 https://github.com/cloudentity/openbanking-quickstart/blob/master/mount/cdr/holder.json

. ACP SaaS (Sample)
```
{
  "IssuerUri": "https://authorization.cloudentity.com:8443/default/cdr",
  "JwksUri": "https://authorization.cloudentity.com:8443/default/cdr/.well-known/jwks.json",
  "AuthorizeUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/authorize",
  "TokenUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/token",
  "IntrospectionUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/introspect",
  "UserinfoUri": "https://authorization.cloudentity.com:8443/default/cdr/userinfo",
  "RegisterUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/register",
  "ParUri": "https://authorization.cloudentity.com:8443/default/cdr/par",
  "RevocationUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/revoke",
  "ArrangementRevocationUri": "https://authorization.cloudentity.com:8443/default/cdr/arrangements/revoke"
  "Registration": {
    "AudienceUri": "https://authorization.cloudentity.com:8443/default/cdr"
  }
  ..
}

```

. ACP local

```
{
  "IssuerUri": "https://authorization.cloudentity.com:8443/default/cdr",
  "JwksUri": "https://authorization.cloudentity.com:8443/default/cdr/.well-known/jwks.json",
  "AuthorizeUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/authorize",
  "TokenUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/token",
  "IntrospectionUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/introspect",
  "UserinfoUri": "https://authorization.cloudentity.com:8443/default/cdr/userinfo",
  "RegisterUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/register",
  "ParUri": "https://authorization.cloudentity.com:8443/default/cdr/par",
  "RevocationUri": "https://authorization.cloudentity.com:8443/default/cdr/oauth2/revoke",
  "ArrangementRevocationUri": "https://authorization.cloudentity.com:8443/default/cdr/arrangements/revoke"
  "Registration": {
    "AudienceUri": "https://authorization.cloudentity.com:8443/default/cdr"
  }
  ..
}

```

### Deploy the docker containers

Based on the deployment style

. ACP SaaS


```
make run-cdr-apps-with-saas
```

. ACP local

```
make run-cdr-apps-with-acp-local
```

Once this is deployed, following url's can be used

## Appliation URL and credentials
- ACP admin portal: `https://authorization.cloudentity.com:8443/app/default/admin` `admin / admin`
- CDR Mock recipient app: `https://localhost:9001/`

### Data Recipient app

Now that we have all the applications deployed and connected to Cloudentity authorization platform. Let's check the
reference application.

The reference `mock-data-recipient` application has screens has integrations to talk to `mock-registry` and `data-holder` components.

Let's verify the application settings are accurate to configured to talk to Cloudentity authorization platform

Navigate to `https://localhost:9001/settings`. The highlighted setting should be pointing to `Cloudentity authorization platform

![Cloudentity-cdr-reference-apps-saas](./sample-mock-data-recipient-settings.png)

Now let's load the data holder registrations seeded from registry onto this app. Click on `Refresh` on `Discover Data Holders`

![Cloudentity-cdr-reference-apps-saas](./load-dh-from-registry.png)


#### Register a dynamic client application

This action will register a client application in target `Cloudentity Authorization Platform`. First let's register
the application in the `mock-data-recipient` app.

![Cloudentity-cdr-reference-apps-saas](./dcr.png)

![Cloudentity-cdr-reference-apps-saas](./dcr-app-registerd.png)


Navigate to `Cloudentity ACP` tenant and choose `CDR` workspace and naviaget to `Applications`. You will see a client 
app that is registered and all appropriate CDR scopes assigned to the application.

![Cloudentity-cdr-reference-apps-saas](./dcr-acp-client-app.png)

![Cloudentity-cdr-reference-apps-saas](./dcr-consents.png)

#### Fetch consent from user by registered client app above

Let's first construct a consent URL, navigate manually to Consent screen in data recipient app and choose the

![Cloudentity-cdr-reference-apps-saas](./consent-request-construct.png)

![Cloudentity-cdr-reference-apps-saas](./consent-request-url.png)

Now copy the URL from above screen and paste into a browser. This call will initiate an authorization consent request
to `Cloudentity authorization` platform.  The consent app is served based on the Cloudentity authorization server setting, for example based on default setting, it will redirect user to `http://localhost:7080`

![Cloudentity-cdr-reference-apps-saas](./consent-setup-acp.png)

This is sample consent app and this  will prompt for authentication with the configured Identity provider(in this case
we have configured a static Idp by default - use `user/p@ssword!` as credentials) and then the consent screen.

![Cloudentity-cdr-reference-apps-saas](./sample-consent-app.png)

Once user consents to the application, user is redirected back to the `mock-data-recipient` app.

![Cloudentity-cdr-reference-apps-saas](./dr-consent-callback.png)


### Manage consents

To manage all the consents and view the accessToken, idToken and to see and revoke all the consents, the data recipient app
provides some UI interfaces.

![Cloudentity-cdr-reference-apps-saas](./dr-consents.png)

### Invoke CDR API's

Now that we have the `cdr-arrangement-id` and associated `accessToken`, we can invoke the API's exposed by the `data-holder`
using the accessToken..

![Cloudentity-cdr-reference-apps-saas](./cdr-api.png)

Sample curl:

```

curl -kv -X  'GET' \
  'https://datarecipient.mock:9001/data-sharing/proxy/cds-au/v1/banking/accounts?open-status=ALL&page=1&page-size=2' \
  -H 'accept: application/json' \
  -H 'x-v: 1' \
  -H 'x-inject-cdr-arrangement-id: c7mf3n75ajjjteq461ig' \
  -H 'x-fapi-auth-date: Fri, 01 Jan 2021 00:00:00 GMT' \
  -H 'Authorization: Bearer eyJhbGciOiJQUzI1NiIsImtpZCI6IjEzNjEwMDEwMzU1MDc2MjU3ODA2OTI1Mzc0MjM1OTkyNDM5OTk5MiIsInR5cCI6IkpXVCJ9.eyJhY3IiOiJ1cm46Y2RzLmF1OmNkcjozIiwiYWlkIjoiY2RyIiwiYW1yIjpbInB3ZCJdLCJhdWQiOlsiYzdtZjAxbjVhampqdGVxNDYwbjAiLCJzcGlmZmU6Ly9hdXRob3JpemF0aW9uLmNsb3VkZW50aXR5LmNvbS9kZWZhdWx0L2Nkci9jN21lb3J2NWFqamp0ZXE0NXZpZyIsInNwaWZmZTovL2F1dGhvcml6YXRpb24uY2xvdWRlbnRpdHkuY29tL2RlZmF1bHQvY2RyL2M3bWVvcnY1YWpqanRlcTQ1dnJnIiwic3BpZmZlOi8vYXV0aG9yaXphdGlvbi5jbG91ZGVudGl0eS5jb20vZGVmYXVsdC9jZHIvYzdtZW9ydjVhampqdGVxNDV2bmciXSwiZXhwIjoxNjQyOTE4OTg0LCJpYXQiOjE2NDI5MTgzODMsImlkcCI6ImJ1Z2tnYWkzZzlrcmVndHUwNHUwIiwiaXNzIjoiaHR0cHM6Ly9hdXRob3JpemF0aW9uLmNsb3VkZW50aXR5LmNvbTo4NDQzL2RlZmF1bHQvY2RyIiwianRpIjoiMTE1Y2UwN2QtMWNiMi00MmY2LWEyZmYtZTkzZjdlODRkZjE3IiwibmJmIjoxNjQyOTE4MzgzLCJzY3AiOlsib3BlbmlkIiwicHJvZmlsZSIsImJhbms6YWNjb3VudHMuYmFzaWM6cmVhZCIsImJhbms6YWNjb3VudHMuZGV0YWlsOnJlYWQiLCJiYW5rOnRyYW5zYWN0aW9uczpyZWFkIiwiY29tbW9uOmN1c3RvbWVyLmJhc2ljOnJlYWQiLCJpbnRyb3NwZWN0X3Rva2VucyIsInJldm9rZV90b2tlbnMiLCJvZmZsaW5lX2FjY2VzcyJdLCJzdCI6InBhaXJ3aXNlIiwic3ViIjoiOWVkN2JkMTEzMjZkMjdlMDRlYTFhNjM2ZDgzMzM0ZDgyY2Y4NzFkNzgwNGRkMDA1OGVmYzI3ZmU0NWVmYjllMSIsInRpZCI6ImRlZmF1bHQifQ.S74W5FR0raKqSMtr9zjyx8Mxgyl7EIyAt4LaMxbMFhOH0aTPNWLF1WHBsOrfMSNZQtAMvkSw9t_4SAmdlUknt7R9LJA5xK7swA7uIpt-nQnzTxaoFLwpsaI4B71nrBoZIQfjn8V6LyxzO7lJqAj_kG_JCPQmWXrI63S653aUEmYjgWBpvslhSVbhaB21QpoI1h16VXqLukS1L8VVozac3eTOe8eBBf1_Yna4AzjbN1OkVraeszSgnLme7AJ91wAmt6KeiY3FXPF-g1B_8QQ-v3wYoRudtB6LFrBqGhqKmvpVX4WdwaTTL9cgqYn0wzmCmHaQy82lniAXN6m4YbFnKA'

```

### Cleanup

If you want to clean up all the locally deployed resources, use the following command. 

```
make clean
```







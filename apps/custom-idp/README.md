# custom-idp

This application demonstrates a Custom Identity Provider that integrates with an external OpenID Connect authorization provider.

Briefly, we are going to create a Custom Identity Provider (Custom IDP) in ACP,
that delegates authentication and authorization to an external OpenID Connect (OIDC) provider.
The process is documented in [Connecting with a custom IDP](https://docs.authorization.cloudentity.com/guides/workspace_admin/connect/custom_idp/).
For convenience in the demonstration, ACP will also play the role of the external OIDC provider,
but it is intended to represent a pre-existing OIDC provider that you wish to integrate with ACP.

## Prerequisites

To run this demo you will need Docker, including the docker-compose command.

If you are running ACP on your local host, you must also define `host.docker.internal` as an alias to `127.0.0.1` in your `/etc/hosts` file:
```
127.0.0.1	host.docker.internal
```

## Configuring ACP Resources

We will create two resources in ACP, both of which are represented by OAuth2 clients:
* Custom Identity Provider (IDP)
* External OpenID Connect (OIDC) Provider

## Create the Custom IDP

In the Default Workspace Overview, select **Identities** in the left-hand menu, and click **CREATE IDENTITY**.
Choose **Custom IDP**, click **Next**, and enter the following properies:

- Name: _Custom-IDP_
- Login URL: `https://127.0.0.1:8080/login`

In the _Custom-IDP_'s _Attributes_ tab, add the following attributes:

| Name         | Display Name | Type   |
|--------------|--------------|--------|
| access_token | Access Token | string |
| id_token     | ID Token     | string |
|              |              |        |

Then, go to the Data Lineage page, and copy each attribute to the **Authentication Context**,
and from there to the _Demo Portal_ box in the **Applications** column.

From the _Custom-IDP_'s _Configuration_ tab, copy the *CLIENT_ID*, *CLIENT_SECRET* and *ISSUER_URL* settings
into the `.env` file. These values are used in the `docker-compose.yaml` file.

```
CLIENT_ID     = c7o7a1f4dqkttnd4356g
CLIENT_SECRET = h8udwSBfgfGkwUkHVsrlQjPAKDM9Ng6w995HPh_j1BY
ISSUER_URL    = https://host.docker.internal:8443/default/system
```

## Create the External OIDC Provider

We are going to create the external OIDC provider as a Web Server Application in its own workspace.

1. In ACP, select **View All Workspaces** in the upper-left menu, and then click **+ Create New**.

2. Select **Consumer Applications and Services** and click **NEXT** at the lower-right.

3. In **Display Name** enter "External Authorizer" and click **NEXT** at the lower-right.

4. Select **Sandbox IDP** and click **NEXT** at the lower-right.

5. In **Name** enter "External Authorizer", enter a user name and password for the initial user. Add as many more users as you wish, and click **NEXT** at the lower-right.

6. For **Developer Portal** click **Skip**.

7. Select the workspace "External Authorizer", select **Settings** on the left-nav, and in the **Authorization** tab, disable both "Enforce PKSE" options, and click **Save Changes**.

8. Select **Applications** on the left-nav, and click **Create Application**

9. In **Name** enter "My Example Web App", select **Web Server Application**, then click **Create**.

10. In the **Scopes** tab, select **Profile** and enable **OpenID**.

11. In the **Overview** tab, enter `https://host.docker.internal:8080/callback` in the **Redirect URL** and click **Save Changes**.

12. Copy the **CLIENT ID** and **CLIENT_SECRET** into the `.env` file:
```
OIDC_CLIENT_ID		= c7ob5jf4dqkttnd43ilg
OIDC_CLIENT_SECRET	= w7GLTQWNVwVu2D9VeNzZDMj-2tDLbxtaEx690NwcD4A
```
13. Copy the **TOKEN** URL to the `OIDC_ISSUER_URL`, removing the `/oauth2/token` suffix:
```
OIDC_ISSUER_URL		= https://host.docker.internal:8443/default/external-authorizer
```

## Running the Demonstration

### Build the Docker image
```
    make docker-build
```

### Start the application

You can now start the Docker container via:
```
	make up
```

You can tail the container logs with:
```
    make logs
```

### Stop the application

When you wish to stop the container:
```
    make down
```

### Interacting with the Demonstration

1. In ACP, click **View All Workspaces** on the upper-left menu.

2. Under the Default Workspace, click "Demo application" in the drop-down menu. This will open the Demo Application in a new tab or window.

3. Click **LOGIN TO DEMO APP**

4. Click **Custom-IDP**

5. In the **External Authorizer**, enter the username and password that your configured in Step 5 above.

6. In the consent page for My Example Web App, click **ALLOW ACCESS**

7. Enjoy the Demo Application

8. Click **SIGN OUT** on the upper-right.


## Parameters

This application obtains its configuration from environment variables:

| Variable           | Required | Default            | Description                                                |
|--------------------|----------|--------------------|------------------------------------------------------------|
| CLIENT_ID          | req      |                    | The Client ID from your custom IDP settings                |
| CLIENT_SECRET      | req      |                    | The Client Secret from your custom IDP settings            |
| ISSUER_URL         | req      |                    | The Issuer URL from your custom IDP settings               |
| CERT_FILE          | req      |                    | Path to the TLS public certificate PEM file                |
| KEY_FILE           | req      |                    | Path to the TLS private key PEM file                       |
| ROOT_CA            | opt      |                    | Path to the Certificate Authority cert PEM for ACP         |
| FAILURE_URL        | req      |                    | URL to redirect user to in case of failure.                |
| GIN_MODE           | opt      | debug              | Sets log level for gin-gonic. Use 'release' for production |
| LOG_LEVEL          | opt      | info               | Sets the level of detail in log output                     |
| PORT               | opt      | 8080               | TCP port where this service will listen for connections    |
| TIMEOUT            | opt      | 5 sec              | Timeout for connections to ACP                             |
| OIDC_AUTH_METHOD   | opt      | client_secret_post | Client Authentication Method(*)                            |
| OIDC_CLIENT_ID     | req      |                    | The Client ID from the external OIDC provider              |
| OIDC_CLIENT_SECRET | req      |                    | The Client Secret from the external OIDC provider          |
| OIDC_ISSUER_URL    | req      |                    | The Issuer URL from the external OIDC provider             |
| OIDC_REDIRECT_URL  | req      |                    | The URL of this server's /callback endpoint                |
| OIDC_CA_PATH       | opt      |                    | Certificate Authority cert for the OIDC provider           |

_*_ Refer to [Client secret authentication methods](https://docs.authorization.cloudentity.com/features/oauth/client_auth/client_secrets/)

The files `docker-compose.yaml` and `.env` file in this directory provide an example
how these environment variables can be used to configure this application.

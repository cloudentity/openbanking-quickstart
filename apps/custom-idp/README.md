# custom-idp

This application demonstrates a Custom Identity Provider that integrates an external identity provider with ACP.

## Parameters

This application obtains its configuration from environment variables:

| Variable      | Required | Default | Description                                                |
|---------------|----------|---------|------------------------------------------------------------|
| CLIENT_ID     | req      |         | The Client ID from your custom IDP settings                |
| CLIENT_SECRET | req      |         | The Client Secret from your custom IDP settings            |
| ISSUER_URL    | req      |         | The Issuer URL from your custom IDP settings               |
| CERT_FILE     | req      |         | Path to the TLS public certificate PEM file                |
| KEY_FILE      | req      |         | Path to the TLS private key PEM file                       |
| ROOT_CA       | req      |         | Path to the root (Certificate Authority) cert PEM for ACP  |
| FAILURE_URL   | req      |         | URL to redirect user to in case of failure.                |
| GIN_MODE      | opt      | debug   | Sets log level for gin-gonic. Use 'release' for production |
| LOG_LEVEL     | opt      | info    | Sets the level of detail in log output                     |
| PORT          | opt      | 8080    | TCP port where this service will listen for connections    |
| TIMEOUT       | opt      | 5 sec   | Timeout for connections to ACP                             |

The files `docker-compose.yaml` and `.env` file in this directory provide an example
how these environment variables can be used to configure this application.

## Prerequisites

To run this demo you will need Docker, including the docker-compose command.

## Create a Custom IDP in ACP

In the Workspace Overview, select **Identities** in the left-hand menu, and click **CREATE IDENTITY**.
Choose **Custom IDP**, click **Next**, and enter the following properies:

- Name: _Custom-IDP_
- Login URL: `https://127.0.0.1:8080/login`

In the _Custom-IDP_'s _Attributes_ tab, add the following authentication context attributes:

| Source | Name            | Variable     | Type            |
| ------ | --------------- | ------------ | --------------- |
| Root   | Phone number    | phone_number | string          |

Then in the _Custom-IDP_'s _Mappings_ tab, set up the following mappings and click **Save Mappings**:

| Source Name     | Target_Name                   |
| --------------- | ----------------------------- |
| Phone Number    | Phone                         |

From the _Configuration_ tab, copy the *CLIENT_ID*, *CLIENT_SECRET* and *ISSUER_URL* settings
into the `.env` file. These values are used in the `docker-compose.yaml` file.
If your issuer URL uses localhost, you should replace that with `host.docker.internal`:

```
CLIENT_ID=c5iuh8neau86enmgnmvg
CLIENT_SECRET=w61w03Dcdmsmzwiyq6ykderpTR7WHYA9EhvK4LfPVQw
ISSUER_URL=https://host.docker.internal:8443/default/system
```

## Build the Docker image
```
    make docker-build
```

## Start the custom-idp container using docker-compose

You can now start the mks-session Docker container via:
```
	make up
```

You can tail the container logs with:
```
    make logs
```

When you wish to stop the container:
```
    make down
```

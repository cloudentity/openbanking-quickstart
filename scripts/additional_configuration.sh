#!/bin/bash
set -e

# configuration
ACP_URL=https://authorization.cloudentity.com:8443
ACP_MTLS_URL=https://authorization.cloudentity.com:8443
TENANT=default
SERVER=cdr
DATA_RECIPIENT_URL=https://datarecipient.mock:9001
MOCK_REGISTER_URL=https://mock-register:7000
MOCK_REGISTER_MTLS_URL=https://mock-register:7001

# do not modify below
URL=${ACP_URL}/${TENANT}/${SERVER}
MTLS_URL=${ACP_MTLS_URL}/${TENANT}/${SERVER}

configure_cdr() {
cat <<EOF > ./mount/cdr/holder.json
{
  "ConnectionStrings": {
    "IdentityServerStoreDatabase": "Data Source=/tmp/idsvr.db",
    "ResourceDatabase": "Data Source=/tmp/mdh.db"
  },
  "IssuerUri": "${URL}",
  "JwksUri": "${URL}/.well-known/jwks.json",
  "AuthorizeUri": "${URL}/oauth2/authorize",
  "TokenUri": "${MTLS_URL}/oauth2/token",
  "IntrospectionUri": "${URL}/oauth2/introspect",
  "UserinfoUri": "${URL}/userinfo",
  "RegisterUri": "${MTLS_URL}/oauth2/register",
  "ParUri": "${URL}/par",
  "RevocationUri": "${URL}/oauth2/revoke",
  "ArrangementRevocationUri": "${URL}/arrangements/revoke",
  "Register": {
    "SsaJwksUri": "${MOCK_REGISTER_URL}/cdr-register/v1/jwks"
  },
  "Registration": {
    "AudienceUri": "${URL}"
  },
  "Logging": {
    "LogLevel": {
      "Default": "Debug",
      "Microsoft": "Debug",
      "Microsoft.Hosting.Lifetime": "Debug"
    }
  },
  "Serilog": {
    "Using": [ "Serilog.Sinks.Console", "Serilog.Sinks.File" ],
    "MinimumLevel": "Debug",
    "WriteTo": [
      { "Name": "Console" },
      {
        "Name": "File",
        "Args": {
          "path": "/tmp/cdr-mdh-identityserver.log",
          "rollingInterval": "Day",
          "rollOnFileSizeLimit": true,
          "fileSizeLimitBytes": "1000000"
        }
      }
    ],
    "Enrich": [ "FromLogContext", "WithMachineName", "WithThreadId" ],
    "Properties": {
      "Application": "CDR.DataHolder.IdentityServer"
    }
  }
}
EOF

cat <<EOF > ./mount/cdr/mock-data-holder/resource-api-appsettings.json
{
    "AcpCDRWorkspace": "${SERVER}",
    "AcpManagementBaseURI": "${ACP_URL}/${TENANT}/system",
    "AcpManagementClientID": "buc3b1hhuc714r78env0",
    "AcpManagementClientSecret": "PBV7q0akoP603rZbU0EFdxbhZ-djxF7FIVwyKaLnBYU",
	"IdentityServerIssuerUri": "${URL}",
	"IdentityServerUrl": "${URL}",
    "AccessTokenIntrospectionEndpoint": "${URL}/oauth2/introspect",
    "ConnectionStrings": {
        "DefaultConnection": "Data Source=/tmp/mdh.db"
    },
    "Logging": {
        "LogLevel": {
            "Default": "Debug",
            "Microsoft": "Debug",
            "Microsoft.Hosting.Lifetime": "Debug"
        }
    },
    "Serilog": {
        "Using": [ "Serilog.Sinks.Console", "Serilog.Sinks.File" ],
        "MinimumLevel": "Debug",
        "WriteTo": [
            { "Name": "Console" },
            {
                "Name": "File",
                "Args": { "path": "/tmp/cdr-mdh-resource-api.log" }
            }
        ],
        "Enrich": [ "FromLogContext", "WithMachineName", "WithThreadId" ],
        "Properties": {
            "Application": "CDR.DataHolder.Resource.API"
        }
    }
}
EOF

cat <<EOF > ./mount/cdr/recipient.json
{
  "ConnectionStrings": {
    "DefaultConnection": "Data Source=/tmp/mdr.db"
  },
  "Logging": {
    "LogLevel": {
      "Default": "Debug",
      "Microsoft": "Warning",
      "Microsoft.Hosting.Lifetime": "Information"
    }
  },
  "Serilog": {
    "Using": [ "Serilog.Sinks.Console", "Serilog.Sinks.File" ],
    "MinimumLevel": "Verbose",
    "WriteTo": [
      { "Name": "Console" },
      {
        "Name": "File",
        "Args": { "path": "/tmp/cdr-mdr-web.log" }
      }
    ],
    "Enrich": [ "FromLogContext", "WithMachineName", "WithThreadId" ],
    "Properties": {
      "Application": "CDR.DataRecipient.Web"
    }
  },
  "ConsumerDataStandardsSwagger": "https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_banking.json",
  "MockDataRecipient": {
    "Register": {
      "tlsBaseUri": "${MOCK_REGISTER_URL}",
      "mtlsBaseUri": "${MOCK_REGISTER_MTLS_URL}",
      "oidcDiscoveryUri": "${MOCK_REGISTER_URL}/idp/.well-known/openid-configuration",
      "tokenEndpoint": "${MOCK_REGISTER_MTLS_URL}/idp/connect/token"
    },
    "SoftwareProduct": {
      "scope": "openid profile bank:accounts.basic:read bank:accounts.detail:read bank:transactions:read common:customer.basic:read introspect_tokens revoke_tokens",
      "jwksUri": "${DATA_RECIPIENT_URL}/jwks",
      "redirectUris": "${DATA_RECIPIENT_URL}/consent/callback,${DATA_RECIPIENT_URL}/consent/callback2",
      "recipientBaseUri": "${DATA_RECIPIENT_URL}"
    },
    "DataHolder": {
      "infosecBaseUri": "${URL}",
      "resourceBaseUri": "https://mock-data-holder:8002",
      "publicBaseUri": "https://mock-data-holder:8000",
      "oidcDiscoveryUri": "${URL}/.well-known/openid-configuration",
      "jwksUri": "${URL}/.well-known/jwks.json",
      "registrationEndpoint": "${MTLS_URL}/oauth2/register"
    }
  }
}
EOF


RS_FILE="./mount/cdr/registry-seed.json"
DATA_RECIPIENT_URL_ESCAPED=$(echo $DATA_RECIPIENT_URL | sed 's;/;\\/;g')
sed -i.bak "s/https:\/\/datarecipient.mock:9001/${DATA_RECIPIENT_URL_ESCAPED}/g" $RS_FILE && rm -f "${RS_FILE}.bak"

URL_ESCAPED=$(echo $URL | sed 's;/;\\/;g')
sed -i.bak "s/https:\/\/authorization.cloudentity.com:8443\/default\/cdr/${URL_ESCAPED}/g" $RS_FILE && rm -f "${RS_FILE}.bak"

}

for ACTION in "$@"
do
  case "$ACTION" in
  cdr)
    configure_cdr
    ;;
  *)
    exit
    ;;
  esac
done

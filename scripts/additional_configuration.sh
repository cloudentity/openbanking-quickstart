#!/bin/bash
set -e

# configuration
export ACP_URL=https://authorization.cloudentity.com:8443
export ACP_MTLS_URL=https://authorization.cloudentity.com:8443
export TENANT=default
export SERVER=cdr
export DATA_RECIPIENT_URL=https://datarecipient.mock:9001
export MOCK_REGISTER_URL=https://mock-register:7000
export MOCK_REGISTER_MTLS_URL=https://mock-register:7001

# do not modify below
export URL=${ACP_URL}/${TENANT}/${SERVER}
export MTLS_URL=${ACP_MTLS_URL}/${TENANT}/${SERVER}

configure_cdr() {
    envsubst < ./data/cdr/mock-apps/holder.template > ./mount/cdr/holder.json
    envsubst < ./data/cdr/mock-apps/recipient.template > ./mount/cdr/recipient.json
    envsubst < ./data/cdr/mock-apps/registry-seed.template > ./mount/cdr/registry-seed.json
    envsubst < ./data/cdr/mock-apps/resource-api-appsettings.template > ./mount/cdr/holder-resource-api-appsettings.json
    if [[ "$1" == "saas" ]]; then ./scripts/override_variables.sh cdr_register_url "https://tpp1.cdr.cloudentity-se.com:7000"; fi
}

override_server() {
    ./scripts/override_env.sh SERVER $1
    ./scripts/override_variables.sh server_id $1
}

for ACTION in "$@"
do
  env=$2
  case "$ACTION" in
  obuk)
    override_server openbanking
    ;;
  obbr)
    override_server openbanking_brasil
    ;;
  cdr)
    override_server cdr
    configure_cdr $env
    ;;
  fdx)
    override_server fdx
    ;;
  *)
    exit
    ;;
  esac
done

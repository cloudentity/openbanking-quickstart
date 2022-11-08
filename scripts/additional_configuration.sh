#!/bin/bash
set -ex

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
    if [[ "$1" == "saas" ]]; then ./scripts/override_variables.sh cdr_adr_validation_enabled false; fi
}

override_server() {
    ./scripts/override_env.sh SERVER $1
    ./scripts/override_variables.sh server_id $1
}

override_client_ids() {
    ./scripts/override_env.sh DEVELOPER_TPP_CLIENT_ID $1
    ./scripts/override_variables.sh developer_tpp_client_id $1

    ./scripts/override_env.sh FINANCROO_TPP_CLIENT_ID $2
    ./scripts/override_variables.sh financroo_tpp_client_id $2
    
    ./scripts/override_env.sh BANK_CLIENT_ID $3
    ./scripts/override_variables.sh bank_client_id $3

    ./scripts/override_env.sh CONSENT_PAGE_CLIENT_ID $4
    ./scripts/override_variables.sh consent_page_client_id $4

    ./scripts/override_env.sh INTERNAL_BANK_CLIENT_ID $5
    ./scripts/override_variables.sh internal_bank_client_id $5
}

for ACTION in "$@"
do
  env=$2
  case "$ACTION" in
  obuk)
    override_server openbanking
    override_client_ids obuk-developer-tpp obuk-financroo-tpp obuk-bank obuk-consent-page obuk-internal-bank-client
    ;;
  obbr)
    override_server openbanking_brasil
    override_client_ids obbr-developer-tpp obbr-financroo-tpp obbr-bank obbr-consent-page obbr-internal-bank-client
    ;;
  cdr)
    override_server cdr
    override_client_ids cdr-developer-tpp cdr-financroo-tpp cdr-bank cdr-consent-page cdr-internal-bank-client
    configure_cdr $env
    ;;
  fdx)
    override_server fdx
    override_client_ids fdx-developer-tpp fdx-financroo-tpp fdx-bank fdx-consent-page fdx-internal-bank-client
    ;;
  *)
    exit
    ;;
  esac
done


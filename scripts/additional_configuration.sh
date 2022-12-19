#!/bin/bash
set -ex

configure_prefix() {
    local prefix=$1
    if [[ "$prefix" != "" ]];
        then echo "${prefix,,}-";
    else
        echo "";
    fi
}

# configuration
export ACP_URL=https://authorization.cloudentity.com:8443
export ACP_MTLS_URL=https://authorization.cloudentity.com:8443
export TENANT=default

# cdr configuration
export SERVER=$(configure_prefix $3)cdr
export SYSTEM_BANK_CLIENT_ID=$(configure_prefix $3)buc3b1hhuc714r78env0
export DATA_RECIPIENT_URL=https://mock-data-recipient:9001
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
    if [[ "$1" == "saas" ]]; 
        then ./scripts/override_variables.sh cdr_adr_validation_enabled false;
    else
        ./scripts/override_variables.sh cdr_adr_validation_enabled true;
    fi
}

override_server() {
    ./scripts/override_env.sh SERVER $1
    ./scripts/override_variables.sh server_id $1

    ./scripts/override_env.sh BANK_CUSTOMERS_SERVER $2
    ./scripts/override_variables.sh bank_customers_server_id $2
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

    ./scripts/override_env.sh CONSENT_SELF_SERVICE_CLIENT_ID $6
    ./scripts/override_variables.sh consent_self_service_client_id $6

    ./scripts/override_env.sh CONSENT_SELF_SERVICE_BACKEND_CLIENT_ID $7
    ./scripts/override_variables.sh consent_self_service_backend_client_id $7

    ./scripts/override_env.sh SYSTEM_BANK_CLIENT_ID $8
    ./scripts/override_variables.sh system_bank_client_id $8

    ./scripts/override_env.sh SYSTEM_ADMIN_CONSENT_CLIENT_ID $9
    ./scripts/override_variables.sh system_admin_consent_client_id $9
}

for ACTION in "$@"
do
  env=$2
  prf=$(configure_prefix $3)
  case "$ACTION" in
  obuk)
    override_server "$prf"openbanking "$prf"bank-customers
    override_client_ids "$prf"obuk-developer-tpp "$prf"obuk-financroo-tpp "$prf"obuk-bank "$prf"obuk-consent-page "$prf"obuk-internal-bank-client "$prf"bv0nab0mekk67nekvq7g "$prf"bv2dkff8mll9cf6pvd6g "$prf"buc3b1hhuc714r78env0 "$prf"bv2fe0tpfc67lmeti340
    ;;
  obbr)
    override_server "$prf"openbanking_brasil "$prf"bank-customers 
    override_client_ids "$prf"obbr-developer-tpp "$prf"obbr-financroo-tpp "$prf"obbr-bank "$prf"obbr-consent-page "$prf"obbr-internal-bank-client "$prf"bv0nab0mekk67nekvq7g "$prf"bv2dkff8mll9cf6pvd6g "$prf"buc3b1hhuc714r78env0 "$prf"bv2fe0tpfc67lmeti340
    ;;
  cdr)
    override_server "$prf"cdr "$prf"bank-customers
    override_client_ids "$prf"cdr-developer-tpp "$prf"cdr-financroo-tpp "$prf"cdr-bank "$prf"cdr-consent-page "$prf"cdr-internal-bank-client "$prf"bv0nab0mekk67nekvq7g "$prf"bv2dkff8mll9cf6pvd6g "$prf"buc3b1hhuc714r78env0 "$prf"bv2fe0tpfc67lmeti340
    configure_cdr $env
    ;;
  fdx)
    override_server "$prf"fdx "$prf"bank-customers
    override_client_ids "$prf"fdx-developer-tpp "$prf"fdx-financroo-tpp "$prf"fdx-bank "$prf"fdx-consent-page "$prf"fdx-internal-bank-client "$prf"bv0nab0mekk67nekvq7g "$prf"bv2dkff8mll9cf6pvd6g "$prf"buc3b1hhuc714r78env0 "$prf"bv2fe0tpfc67lmeti340
    ;;
  *)
    exit
    ;;
  esac
done


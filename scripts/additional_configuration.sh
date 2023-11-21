#!/bin/bash
set -ex

# prefix
export PREFIX=$(echo $3 | tr [:upper:] [:lower:])

configure_prefix() {
    local id=$1
    if [[ "$PREFIX" != "" ]];
        then echo "$PREFIX-$id";
    else
        echo "$id";
    fi
}

# configuration
export ACP_URL=https://authorization.cloudentity.com:8443
export ACP_MTLS_URL=https://authorization.cloudentity.com:8443
export TENANT=default

# cdr configuration
export SERVER=$(configure_prefix "cdr")
export SYSTEM_BANK_CLIENT_ID=$(configure_prefix "buc3b1hhuc714r78env0")
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
  arrEnv=("SERVER" "BANK_CUSTOMERS_SERVER")
  arrVar=("server_id" "bank_customers_server_id")
  arrIds=($@)

  for i in "${!arrIds[@]}"; do
    ./scripts/override_env.sh "${arrEnv[i]}" $(configure_prefix "${arrIds[i]}")
    ./scripts/override_variables.sh "${arrVar[i]}" $(configure_prefix "${arrIds[i]}")
  done
}

override_client_ids() {
  arrEnv=("DEVELOPER_TPP_CLIENT_ID" "FINANCROO_TPP_CLIENT_ID" "BANK_CLIENT_ID" \
  "CONSENT_PAGE_CLIENT_ID" "INTERNAL_BANK_CLIENT_ID" "CONSENT_SELF_SERVICE_CLIENT_ID" \
  "CONSENT_SELF_SERVICE_BACKEND_CLIENT_ID" "SYSTEM_BANK_CLIENT_ID" "SYSTEM_ADMIN_CONSENT_CLIENT_ID")

  arrVar=("developer_tpp_client_id" "financroo_tpp_client_id" "bank_client_id" \
  "consent_page_client_id" "internal_bank_client_id" "consent_self_service_client_id" \
  "consent_self_service_backend_client_id" "system_bank_client_id" "system_admin_consent_client_id")

  arrIds=($@)

  for i in "${!arrIds[@]}"; do
    ./scripts/override_env.sh "${arrEnv[i]}" $(configure_prefix "${arrIds[i]}")
    ./scripts/override_variables.sh "${arrVar[i]}" $(configure_prefix "${arrIds[i]}")
  done
}

for ACTION in "$@"
do
  env=$2
  system_clients="bv0nab0mekk67nekvq7g bv2dkff8mll9cf6pvd6g buc3b1hhuc714r78env0 bv2fe0tpfc67lmeti340"
  case "$ACTION" in
  obuk)
    override_server "openbanking" "bank-customers"
    override_client_ids "obuk-developer-tpp" "obuk-financroo-tpp" "obuk-bank" "obuk-consent-page" "obuk-internal-bank-client" $system_clients
    ./scripts/override_variables.sh  "server_profile" "openbanking_uk"
    ;;
  obbr)
    override_server "openbanking_brasil" "bank-customers" "openbanking_br"
    override_client_ids "obbr-developer-tpp" "obbr-financroo-tpp" "obbr-bank" "obbr-consent-page" "obbr-internal-bank-client" $system_clients
    ./scripts/override_variables.sh  "server_profile" "openbanking_br"
    ;;
  cdr)
    override_server "cdr" "bank-customers" "cdr_australia_fapi_rw"
    override_client_ids "cdr-developer-tpp" "cdr-financroo-tpp" "cdr-bank" "cdr-consent-page" "cdr-internal-bank-client" $system_clients
    ./scripts/override_variables.sh  "server_profile" "cdr_australia_fapi_rw"
    configure_cdr $env
    ;;
  fdx)
    override_server "fdx" "bank-customers" "fdx"
    override_client_ids "fdx-developer-tpp" "fdx-financroo-tpp" "fdx-bank" "fdx-consent-page" "fdx-internal-bank-client" $system_clients
    ./scripts/override_variables.sh  "server_profile" "fdx"
    ;;
  generic)
    #override_server "gobank" "bank-customers" "gobank"
    #override_client_ids "gobank-developer-tpp" "gobank-financroo-tpp" "gobank-bank" "gobank-consent-page" "gobank-internal-bank-client" $system_clients
    ./scripts/override_variables.sh  "server_profile" "fapi_20_security"
    ;;
  *)
    exit
    ;;
  esac
done


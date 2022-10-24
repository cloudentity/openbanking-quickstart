#!/bin/bash

set -e

if [[ ! -v SAAS_TENANT_ID ]]; then
    echo "SAAS_TENANT_ID is not set"
    exit 1
else
    ./scripts/override_env.sh TENANT ${SAAS_TENANT_ID}
fi

if [[ ! -v SAAS_CLIENT_ID ]]; then
    echo "SAAS_CLIENT_ID is not set"
    exit 1
else
    ./scripts/override_env.sh CONFIGURATION_CLIENT_ID ${SAAS_CLIENT_ID}
fi

if [[ ! -v SAAS_CLIENT_SECRET ]]; then
    echo "SAAS_CLIENT_SECRET is not set"
    exit 1
else
    ./scripts/override_env.sh CONFIGURATION_CLIENT_SECRET ${SAAS_CLIENT_SECRET}
fi


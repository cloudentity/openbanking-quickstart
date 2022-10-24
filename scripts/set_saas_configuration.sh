#!/bin/bash

set -e

#!/bin/bash

vars=( SAAS_TENANT_ID SAAS_CLIENT_ID SAAS_CLIENT_SECRET )
vars2=( TENANT CONFIGURATION_CLIENT_ID CONFIGURATION_CLIENT_SECRET )


for index in ${!vars[*]}; do 
    if [[ ! -v ${vars[$index]} ]]; then
        echo "${vars[$index]} is not set"
        exit 1
    else
    ./scripts/override_env.sh ${vars2[$index]} $(printenv ${vars[$index]})
    fi
done

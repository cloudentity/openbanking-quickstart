#!/bin/bash

set -e

enabled=true && [[ $1 == "disable" ]] && enabled=false
scheme=https && [[ $1 == "disable" ]] && scheme=http

dockerfiles=(
    docker-compose.cdr.yaml
    docker-compose.obuk.yaml 
    docker-compose.obbr.yaml
)

for file in "${dockerfiles[@]}"
do 
    sed -i '' 's/\(ENABLE_TLS_SERVER=\)\(.*\)$/\1'$enabled'/' $file
    sed -i '' 's/\(UI_URL=\)\(.*\)\(:\/\/.*\)$/\1'$scheme'\3/' $file
done

sed -i '' 's/\(financroo_tpp_url: \)\(.*\)\(:\/\/.*\)$/\1'$scheme'\3/' data/variables.yaml

#!/bin/bash

set -e

function generateClient()
{
rm -rf $1/*
docker-compose -f docker-compose.acp.local.yaml exec runner sh -c \
"swagger generate client \
    -f $3 \
    -A $2 \
    -t $1"
}

case $1 in

  obbr)
    generateClient "./generated/obbr/consents" "consents" "api/obbr/consents.yaml"
    generateClient "./generated/obbr/payments" "payments" "api/obbr/payments.yaml"
    generateClient "./generated/obbr/accounts" "accounts" "api/obbr/accounts.yaml"
    ;;

  obuk)
    generateClient "./generated/obuk/payments" "payments" "api/obuk/payments.yaml"
    generateClient "./generated/obuk/accounts" "accounts" "api/obuk/accounts.yaml"
    ;;

  cdr)
    generateClient "./generated/cdr" "banking" "api/cdr/cdr.yaml"
    ;;

  fdx)
    generateClient "./generated/fdx" "client" "api/fdx/fdx.yaml"
    ;;

  *)
    echo -n "unknown"
    ;;
esac

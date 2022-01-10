#!/bin/bash

sed -i'.bak' "s|^$1=.*|$1=$2|g" .env-local && rm .env-local.bak
sed -i'.bak' "s|^$1=.*|$1=$2|g" .env-saas && rm .env-saas.bak

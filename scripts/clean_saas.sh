#!/bin/bash
set -ex 
grep -v '^#' .env-saas
export $(grep -v '^#' .env | xargs)
go run ./scripts/go/clean_saas.go -spec=$1

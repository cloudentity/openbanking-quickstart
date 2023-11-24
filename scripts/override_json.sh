#!/bin/bash

tmp=$(mktemp)
path=mount/financroo-tpp/banks.json
jq "$1 = \"$2\"" $path > "$tmp" && mv "$tmp" $path

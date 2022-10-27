#!/bin/bash

set -e

enabled=true && [[ $1 == "disable" ]] && enabled=false

./scripts/override_env.sh ENABLE_MFA $enabled

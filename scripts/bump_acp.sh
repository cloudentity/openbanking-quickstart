#!/bin/bash

latest_tag=$(docker images --format '{{.Tag}}' docker.cloudentity.io/acp | head -n 1)
./scripts/override_env.sh ACP_VERSION $latest_tag

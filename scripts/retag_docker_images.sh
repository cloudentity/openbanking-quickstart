#!/bin/bash

set -x

while IFS= read image; do
    docker tag $image ${image/cloudentity\//docker.cloudentity.io\/}
done

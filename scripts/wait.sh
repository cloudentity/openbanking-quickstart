#!/bin/bash
set -e

pendingServices=0

function checkServices() {
    pendingServices=$(docker ps --format "{{.Status}}" | grep -E "starting|unhealthy" | wc -l)
}

checkServices

while [ $pendingServices -gt 0 ]; do
    echo "Waiting for $pendingServices services to become healthy..."
    docker ps -a
    sleep 1
    checkServices
done

echo "All services are healthy"

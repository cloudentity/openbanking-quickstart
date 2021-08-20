# How to run quickstart apps with oidf conformance suit with OBB payments consents
1. Start docker compose:
    ```
    docker-compose -f docker-compose.yaml -f docker-compose.brasil.yaml -f conformance/docker-compose.fapi.yaml up -d 
    ```
1. Add `127.0.0.1 fapi-test` to your hosts config
1. Go to https://authorization.cloudentity.com:8443/app/default/admin and disable `PKCE` enforcement in your workspace settings
1. Go to https://fapi-test:8444 and use config from `conformance/obb-oidf-payments.json` to set up a new test plan

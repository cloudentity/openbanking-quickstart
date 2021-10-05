# How to run quickstart apps with OIDF and OBBR conformance suites 
1. Start quickstart using:
    ```
    make obbr cmd="up -d" 
    ```
2. Add `127.0.0.1 fapi-test` to your hosts config
3. Go to https://fapi-test:8444 and use config from `conformance/obb-oidf-payments.json` to set up a new test plan
4. Alternatively you can go to https://fapi-test to run obbr functional conformance suite.
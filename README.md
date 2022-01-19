# openbanking-quickstart
Openbanking Quickstart

## How to run environment locally
### Set up MFA
By default, MFA is turned off.
1. Run `make enable-mfa` to turn MFA on.
2. Run `make disable-mfa` to turn MFA off.

### Set up ACP locally
Run `make run-dev` to set up ACP.
> #### Note
> By default, MFA is turned off.
> To run environment with MFA run `make enable-mfa`.

## How to run Cypress test locally
Run `make run-tests` to open Cypress GUI with tests scenarios.
To run tests:
- single test suite - click on intended test suite name `{TEST_NAME}.ts`
- full set - click on `Run integration spec`

> #### Note
> Tests will be executed accordingly to the set load for `MFA`- see **Set up MFA**. 

## How to use Quickstart with SAAS
### How to start up quickstart with SAAS
- Enable Admin workspace and System workspace on SAAS Tenant
- Go to "System" workspace and create a new application with `manage_configuration` scope and `client_credentials` grant type
- Configure the `.env-saas` file.
- Run `make run-apps-with-saas`
### How to clean up SAAS tenant
- Run `make-clean-saas`. This will turn off all services and delete all extra workspaces from tenant.

## How to Run Quickstart with Different Open Banking Specifications
- If you have any of the quickstart applications running, turn them all off first with `make clean`
- To run in Open Banking UK mode: run `make enable-spec-obuk run` if running locally. Run `make enable-spec-obuk run-apps-with-saas` if running with saas
- To run in Open Banking Brasil mode: run `make enable-spec-obbr run` if running locally. Run `make enable-spec-obbr run-apps-with-saas` if running with saas
## Credentials
- ACP admin portal: `https://localhost:8443/app/default/admin` `admin / admin`
- Developer TPP: `https://localhost:8090` `user | user2 | user3 / p@ssw0rd!`
- Financroo TPP: `https://localhost:8091` `test / p@ssw0rd!`
- Consent self service portal: `https://localhost:8085` `user | user2 | user3 / p@ssw0rd!`
- Consent admin portal: `https://localhost:8086` `admin / p@ssw0rd!`

## Licenses
- [Bank](apps/bank/LICENSE) - Apache 2.0
- [Developer TPP](apps/developer-tpp/LICENSE) - Apache 2.0
- [Financroo TPP](apps/financroo-tpp/LICENSE) - Apache 2.0
- [Configuration app](apps/configuration/LICENSE) - Apache 2.0
- [Consent self service portal](consent/self-service-portal/LICENSE) - Cloudentity
- [Consent admin portal](consent/admin-portal/LICENSE) - Cloudentity
- [Consent page](consent/consent-page/LICENSE) - Cloudentity

## Release process
- `git checkout -b release/VERSION`
- `make set-version`
- update CHANGELOG.md
- `git add . && git commit -m 'Release VERSION' && git push`
- verify if github action build and pushed released images to public docker hub
- `git tag -a VERSION && git push --tags`

## FAQ
### I've deployed my quickstart apps under a domain that's not "localhost", but after I run the configuration job, the client application redirect urls are all still localhost.
-  Under `data/variables/yaml` edit the values for `consent_self_service_portal_url`, ` consent_admin_portal_url`, `consent_page_url`, `developer_tpp_url`, and `financroo_tpp_url`.
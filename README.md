# openbanking-quickstart
Openbanking Quickstart


## How to Run Quickstart
### Prequisite configuration
| ACP type      | Associated Environment File | Additional Configuration needed? 
| ----------- | ----------- | --------|
| Local      | .env-local   | No, everything is good to go out of the box. You shouldn't need to touch this file|
| SAAS   | .env-saas        | [Yes, running with saas takes a few extra steps] (#saas-configuration-instructions) |

> #### Additionally, quickstart can be used to showcase flows with MFA
> By default, MFA is turned off.
> 1. Run `make enable-mfa` to turn MFA on.
> 2. Run `make disable-mfa` to turn MFA off.

### Makefile targets 
The current types of runtimes are currently supported: 
1. Open Banking UK with local ACP instance: `make run-obuk-local`
2. Open Banking UK with SAAS ACP: `make run-obuk-saas`
3. Open Banking Brasil with local ACP instance: `make run-obbr-local`
4. Open Banking Brasil with SAAS ACP: `make run-obbr-saas`
5. CDR with local ACP instance: `make run-cdr-local`


## How to run Cypress Tests
| Mode      | Instructions |
| ----------- | ----------- |
| With Cypress UI      | Run `make run-tests` to open Cypress GUI with tests scenarios. Then click on the intended test suite name `{TEST_NAME}.ts`    |
| Headless (commandline)   | Run one of the following: `run-obuk-tests-headless`, `run-obbr-tests-headless`, `run-cdr-tests-headless`        |


## Configuring your ACP SAAS Tenant for Quickstart {#saas-configuration-instructions}
- Enable System workspace on SAAS Tenant (you will most likely need to contact someone at Cloudentity to do this for you)
- Go to "System" workspace and create a new application with `manage_configuration` scope and `client_credentials` grant type
- Configure the `.env-saas` file with your saas tenant id.
- Run one of the saas related makefile targets


## Credentials
- ACP admin portal: `https://authorization.cloudentity.com:8443` `admin / admin`
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

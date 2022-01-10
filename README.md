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
- Configure the `.env-saas` with correct parameters.
  - There are two optional fields which are not crucial for quickstart to work with SAAS. These are `ADMIN_CLIENT_ID` and `ADMIN_CLIENT_SECRET`, which can be left unconfigured
- Run `make run-apps-with-saas`
### How to clean up SAAS tenant afterwards
The quickstart project creates a lot of extra workspaces to showcase its capabilities. To quickly cleanup your SAAS tenant after turning off quickstart, do the following:
- Navigate to admin workspace in SAAS and create an admin application with `client_credentials` grant flow enabled. 
- Copy the client id and client secret from the UI into the `ADMIN_CLIENT_ID` and `ADMIN_CLIENT_SECRET` fields in the `.env-saas` file.
- Run `make clean-saas`, and all of the workspaces created by quickstart will be deleted. 
- Or if you don't want to do this, you can just manually delete all the workspaces via the UI. 

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

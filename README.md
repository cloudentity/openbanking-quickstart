# openbanking-quickstart

This repository demonstrates Cloudentity's capabilities for handling Openbanking scenarios.

You can find here a sample implementation of TPP, consent-page, bank, admin and self-service apps.

The following specifications are currently supported:

* Openbanking UK
* Openbanking Brasil
* CDR
* FDX (partial support)

For more details, please visit our [Openbanking Guides](https://docs.authorization.cloudentity.com/guides/ob_guides/).

### Prequisite configuration

#### Required software

- `make`
- `docker` > 20.10.2
- `docker-compose` > 1.29.0

For Windows you need to use WSL.

#### Deployment mode

There are two ways you can run quickstart.

If you want to set up everything locally then you need to contact sales@cloudentity.com to get credentials to be able to
download ACP docker image.

The other option is to register and use ACP in SaaS: https://authz.cloudentity.io

| ACP type | Associated Environment File | Additional Configuration needed?                                                              |
|----------|-----------------------------|-----------------------------------------------------------------------------------------------|
| Local    | .env-local                  | No, everything is good to go out of the box. You shouldn't need to touch this file            |
| SAAS     | .env-saas                   | <a href="#saas-configuration-instructions">Yes, running with saas takes a few extra steps</a> |

> #### Additionally, quickstart can be used to showcase flows with MFA
> By default, MFA is turned off.
> 1. Run `make enable-mfa` to turn MFA on.
> 2. Run `make disable-mfa` to turn MFA off.

<br/>
<h3 id="makefile-targets"> Makefile Targets</h3>

The current types of runtimes are currently supported:

1. Open Banking UK with local ACP instance: `make run-obuk-local`
2. Open Banking UK with SAAS ACP: `make run-obuk-saas`
3. Open Banking Brasil with local ACP instance: `make run-obbr-local`
4. Open Banking Brasil with SAAS ACP: `make run-obbr-saas`
5. CDR with local ACP instance: `make run-cdr-local`

<br/>

## How to run Cypress Tests

| Mode                   | Instructions                                                                                                               |
|------------------------|----------------------------------------------------------------------------------------------------------------------------|
| With Cypress UI        | Run `make run-tests` to open Cypress GUI with tests scenarios. Then click on the intended test suite name `{TEST_NAME}.ts` |
| Headless (commandline) | Run one of the following: `run-obuk-tests-headless`, `run-obbr-tests-headless`, `run-cdr-tests-headless`                   |

<br/>

<h2 id="saas-configuration-instructions">Configuring your ACP SAAS Tenant For Quickstart</h3>

1. Enable System workspace on your ACP SAAS Tenant (you will most likely need to contact someone at Cloudentity to do
   this for you)
2. Go to "System" workspace and create a new application with `manage_configuration` scope and `client_credentials`
   grant type
3. Configure the `.env-saas` file with the following:
    * Saas tenant id
    * id of the client you created in step 2
    * secret of the client you created in step 2
4. Now you are ready to run one of the saas related <a href="#makefile-targets"> makefile targets</a>

## How to Run Quickstart

The current types of runtimes are currently supported:

1. Open Banking UK with local ACP instance: `make run-obuk-local`
2. Open Banking UK with SAAS ACP: `make run-obuk-saas`
3. Open Banking Brasil with local ACP instance: `make run-obbr-local`
4. Open Banking Brasil with SAAS ACP: `make run-obbr-saas`
5. CDR with local ACP instance: `make run-cdr-local`
6. FDX with local ACP instance: `make run-fdx-local`
7. FDX with SAAS ACP: `make run-fdx-saas`

> #### Additionally, quickstart can be used to showcase flows with MFA
> By default, MFA is turned off.
> 1. Run `make enable-mfa` to turn MFA on.
> 2. Run `make disable-mfa` to turn MFA off.

To clean up the environment execute: `make clean`.
This step is also necessary if you want to switch between different specs.

> ### Using Hypr
>To run using Hypr Passwordless see [Hypr Passwordless Setup](docs/how_to_use_hypr.md)

## What to do next

Once you can run the quickstart, you can visit our sample apps:

## Credentials

| Name                        | Url                                          | Credentials                    | Availability          |
|-----------------------------|----------------------------------------------|--------------------------------|-----------------------|
| Financroo TPP               | `https://localhost:8091`                     | `test / p@ssw0rd!`             |                       |                                          
| Developer TPP               | `https://localhost:8090`                     | `user user2 user3 / p@ssw0rd!` |                       | 
| Consent self service portal | `https://localhost:8085`                     | `user user2 user3 / p@ssw0rd!` |                       | 
| Consent admin portal        | `https://localhost:8086`                     | `admin / p@ssw0rd!`            |                       |
| ACP admin portal            | `https://authorization.cloudentity.com:8443` | `admin / admin`                | only local deployment |

## How to run Tests

| Mode                   | Instructions                                                                                                               |
|------------------------|----------------------------------------------------------------------------------------------------------------------------|
| With Cypress UI        | Run `make run-tests` to open Cypress GUI with tests scenarios. Then click on the intended test suite name `{TEST_NAME}.ts` |
| Headless (commandline) | Run one of the following: `run-obuk-tests-headless`, `run-obbr-tests-headless`, `run-cdr-tests-headless`                   |

## Licenses

- [Bank](apps/bank/LICENSE) - Apache 2.0
- [Developer TPP](apps/developer-tpp/LICENSE) - Apache 2.0
- [Financroo TPP](apps/financroo-tpp/LICENSE) - Apache 2.0
- [Configuration app](apps/configuration/LICENSE) - Apache 2.0
- [Consent self-service portal](consent/self-service-portal/LICENSE) - Apache 2.0
- [Consent admin portal](consent/admin-portal/LICENSE) - Apache 2.0
- [Consent page](consent/consent-page/LICENSE) - Apache 2.0 

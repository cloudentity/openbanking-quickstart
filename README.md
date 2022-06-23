# openbanking-quickstart

This repository demonstrates Cloudentity's capabilities for handling Openbanking scenarios.

You can find here a sample implementation of TPP, consent-page, bank, admin and self service apps.

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

If you want to set up everything locally then you need to contact sales@cloudentity.com to get credentials to be able to download ACP docker image.

The other option is to register and use ACP in SaaS: https://authz.cloudentity.io

| ACP type      | Associated Environment File | Additional Configuration needed?
| ----------- | ----------- | --------|
| Local      | .env-local   | No, everything is good to go out of the box. You shouldn't need to touch this file|
| SAAS   | .env-saas        | <a href="#saas-configuration-instructions">Yes, running with saas takes a few extra steps</a>|

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
| Mode      | Instructions |
| ----------- | ----------- |
| With Cypress UI      | Run `make run-tests` to open Cypress GUI with tests scenarios. Then click on the intended test suite name `{TEST_NAME}.ts`    |
| Headless (commandline)   | Run one of the following: `run-obuk-tests-headless`, `run-obbr-tests-headless`, `run-cdr-tests-headless`        |

<br/>

<h2 id="saas-configuration-instructions">Configuring your ACP SAAS Tenant For Quickstart</h3>

1. Enable System workspace on your ACP SAAS Tenant (you will most likely need to contact someone at Cloudentity to do this for you)
2. Go to "System" workspace and create a new application with `manage_configuration` scope and `client_credentials` grant type
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

To cleanup the environment execute: `make clean`.
This step is also necessary if you want to switch between different specs.

## What to do next

Once you can run the quickstart, you can visit our sample apps:

<h2>Configure QuickStart for Hypr Passwordless</h2>

### Registering a User and Device with Hypr

If using Hypr passwordless you will need to have a username and a registered device. Download either the [iOS Hypr app](https://apps.apple.com/us/app/hypr/id1343368858) or the [Android Hypr app](https://play.google.com/store/apps/details?id=com.hypr.one&hl=en_US&gl=US). To register a device with your username, you can generate a magic link by performing the following and replacing the placeholder values with your own:

```bash
curl --request POST \
  --url https://demo.gethypr.com/rp/api/versioned/magiclink \
  --header 'Authorization: Bearer <your Hypr token>' \
  --header 'Content-Type: application/json' \
  --data '{
  "username": "user",
  "email": "example@example.com",
  "firstname": "",
  "lastname": "",
  "message": "",
  "secondsValid": "6000",
  "hyprServerUrl": "https://demo.gethypr.com"
}'
```

This will return a response with a magic link(`webLink`). Go to the `webLink` on your desktop in a browser and choose the device method to register. If registering a mobile device, this will be a QR code which you can then scan with the Hypr mobile app. Once scanned your username and device are registered.

### Environment Variables for Hypr
In QuickStart set the environment variables in `.env-local`. The required environment variables are:
- HYPR_TOKEN - API token provided by Hypr
- ENABLE_MFA - must be set to `true`

Additionally, if using a different Hypr tenant and App ID set the following environment variables in `.env-local`:
- HYPR_BASE_URL - URL for your Hypr tenant
- HYPR_APP_ID - App ID for Hypr

Quickstart is now enabled to work with Hypr Passwordless. Connect Go Bank as before and you will be prompted to login on your Hypr enabled device.
## Credentials
- ACP admin portal: `https://authorization.cloudentity.com:8443` `admin / admin`
- Developer TPP: `https://localhost:8090` `user | user2 | user3 / p@ssw0rd!`
- Financroo TPP: `https://localhost:8091` `test / p@ssw0rd!`
- Developer TPP: `https://localhost:8090` `user | user2 | user3 / p@ssw0rd!`
- Consent self service portal: `https://localhost:8085` `user | user2 | user3 / p@ssw0rd!`
- Consent admin portal: `https://localhost:8086` `admin / p@ssw0rd!`
- ACP admin portal: `https://authorization.cloudentity.com:8443` `admin / admin` (only local deployment)

## How to run Tests
| Mode      | Instructions |
| ----------- | ----------- |
| With Cypress UI      | Run `make run-tests` to open Cypress GUI with tests scenarios. Then click on the intended test suite name `{TEST_NAME}.ts`    |
| Headless (commandline)   | Run one of the following: `run-obuk-tests-headless`, `run-obbr-tests-headless`, `run-cdr-tests-headless`        |

## Licenses
- [Bank](apps/bank/LICENSE) - Apache 2.0
- [Developer TPP](apps/developer-tpp/LICENSE) - Apache 2.0
- [Financroo TPP](apps/financroo-tpp/LICENSE) - Apache 2.0
- [Configuration app](apps/configuration/LICENSE) - Apache 2.0
- [Consent self service portal](consent/self-service-portal/LICENSE) - Cloudentity
- [Consent admin portal](consent/admin-portal/LICENSE) - Cloudentity
- [Consent page](consent/consent-page/LICENSE) - Cloudentity

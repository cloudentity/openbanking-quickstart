## [1.10.0]
- release stable version of openbanking apps

## [1.11.0]
- makes it easier to customize consent-page using:
    - CSS API
    - i18n
    - generic HTTP mfa_handler

## [1.11.1]
- minor fixes

## [1.11.2-rc1]
- support for OBBR consents

## [1.11.2-rc2]
- return access_denied when user rejects the consent 

## [1.11.2-rc3]
- support for OBBR bank API
- extended templates model with `ctx AuthenticationContext`
- configurable `BANK_ID_CLAIM` (default value `sub`) that is used to choose claim that is used to load data from bank
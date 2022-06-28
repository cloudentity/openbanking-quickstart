<h2>Configure QuickStart for Hypr Passwordless Provider</h2>

### Using Hypr

If using Hypr passwordless you will need to have a username and a registered device. Download either the [iOS Hypr app](https://apps.apple.com/us/app/hypr/id1343368858) or the [Android Hypr app](https://play.google.com/store/apps/details?id=com.hypr.one&hl=en_US&gl=US). To register a device with your username, you can generate a magic link by performing the following and replacing the placeholder values with your own:

```bash
curl --request POST \
  --url https://demo.gethypr.com/rp/api/versioned/magiclink \
  --header 'Authorization: Bearer <your Hypr token>' \
  --header 'Content-Type: application/json' \
  --data '{
  "username": "<your user name>",
  "email": "<your email>",
  "firstname": "",
  "lastname": "",
  "message": "",
  "secondsValid": "6000",
  "hyprServerUrl": "https://hypr-tenant-name.your-hypr-domain.com"
}'
```

>By default, quickstart creates a username and password for the IDP. When adding your device and registering a user you must add the username that was registered to the IDP in SaaS.

This will return a response with a magic link(`webLink`). Go to the `webLink` on your desktop in a browser and choose the device method to register. If registering a mobile device, this will be a QR code which you can then scan with the Hypr mobile app. Once scanned your username and device are registered.

### Environment Variables for Hypr

In QuickStart set the environment variables in `.env-saas`. The required environment variables are:
- MFA_PROVIDER - set to `hypr`
- ENABLE_MFA - must be set to `true`
- MFA_PROVIDER_CONFIG_PATH - path to the Hypr mfa config: defaults to:`../config_hypr.yaml`

### Configuration for Hypr

Additionally, in `data/config_hypr.yaml` replace the placeholder values the following:
- HYPR_TOKEN - your Hypr API Token
- HYPR_BASE_URL - your hypr tenant base url which must match the application registered in the previous step
- HYPR_APP_ID - your hypr application ID

Quickstart is now enabled to work with Hypr Passwordless. Start Financroo and add Go Bank and you will be prompted to authorize Financroo on your Hypr enabled device by running:
```bash
make run-obuk-saas
```
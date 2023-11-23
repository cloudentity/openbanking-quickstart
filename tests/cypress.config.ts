import { defineConfig } from 'cypress'

export default defineConfig({
  viewportWidth: 1920,
  viewportHeight: 1080,
  chromeWebSecurity: false,
  env: {
    consent_self_service_url: 'https://localhost:8085',
    consent_admin_url: 'https://localhost:8086',
    tpp_url: 'https://localhost:8090',
    financroo_url: 'https://localhost:8091',
    mock_data_recipient_url: 'https://localhost:9001',
  },
  projectId: 'icwsy8',
  defaultCommandTimeout: 10000,
  e2e: {
    // We've imported your old cypress plugins here.
    // You may want to clean this up later by importing these.
    setupNodeEvents(on, config) {
      return require('./cypress/plugins/index.js')(on, config)
    },
    specPattern: 'cypress/e2e/**/*.{js,jsx,ts,tsx}',
  },
})

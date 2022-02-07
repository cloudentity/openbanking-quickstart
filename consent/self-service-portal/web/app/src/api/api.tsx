import {base, baseWithCustomBaseUrl} from './api-base';

export const api = {
    userinfo: (authorizationServerURL, tenantId, authorizationServerId) => baseWithCustomBaseUrl('/', authorizationServerURL).get({url: `/${tenantId}/${authorizationServerId}/userinfo`}),
    getConsents: () => base.get({url: `/consents`}),
    deleteConsent: ({id, consent_type}) => base.delete({url: `/consents/${id}`, query: {consent_type: consent_type}}),
};

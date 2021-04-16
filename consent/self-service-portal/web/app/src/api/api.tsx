import {base, baseWithCustomBaseUrl} from './api-base';

export const api = {
    userinfo: (authorizationServerURL, tenantId, authorizationServerId) => baseWithCustomBaseUrl('/', authorizationServerURL).get({url: `/${tenantId}/${authorizationServerId}/userinfo`}),
    getConsents: () => base.get({url: `/consents`}),
    deleteConsent: ({id}) => base.delete({url: `/consents/${id}`, query: {}}),
};

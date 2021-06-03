import { base, baseWithCustomBaseUrl } from "./api-base";
//import {clientMockRes} from "../components/clientsMockRes";

export const api = {
  userinfo: (
    authorizationServerURL: string,
    tenantId: string,
    authorizationServerId: string
  ) =>
    baseWithCustomBaseUrl("/", authorizationServerURL).get({
      url: `/${tenantId}/${authorizationServerId}/userinfo`,
    }),
  getClients: () => base.get({ url: `/clients` }),
  // getClients: () => Promise.resolve(clientMockRes),
  deleteClient: ({ id }: { id: string }) =>
    base.delete({ url: `/clients/${id}`, query: {} }),
  getConsents: () => base.get({ url: `/consents` }),
  deleteConsent: ({ id }: { id: string }) =>
    base.delete({ url: `/consents/${id}`, query: {} }),
};

import { availableConstentTypesJoined } from "../components/utils";
import { base, baseWithCustomBaseUrl } from "./api-base";

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
  deleteClient: ({ id }: { id: string }) =>
    base.delete({
      url: `/clients/${id}`,
      query: { consent_types: availableConstentTypesJoined },
    }),
  getConsents: () => base.get({ url: `/consents` }),
  deleteConsent: ({ id, consent_type }: { id: string, consent_type: string }) =>
    base.delete({ url: `/consents/${id}`, query: {consent_type: consent_type} }),
};

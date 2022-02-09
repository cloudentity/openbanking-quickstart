import React, { useEffect, useState } from "react";
import { Container, makeStyles, Theme, Typography } from "@material-ui/core";

import PageToolbar from "../PageToolbar";
import Progress from "../Progress";
import { api } from "../../api/api";
import noAccountEmptyState from "../no-accounts-empty-state.svg";
import iconShield from "../../assets/icon-shield.svg";
import ClientsList from "../ClientsList";
import Subheader from "../Subheader";
import SearchInput from "../SearchInput";
import { ClientType } from "../utils";

const useStyles = makeStyles((theme: Theme) => ({
  subtitle: {
    ...theme.custom.body1,
    maxWidth: 588,
    margin: "0 auto",
  },
}));

export const searchTabs = (onSearch: (searchText: string) => void) => [
  {
    key: "account",
    label: "Account",
    content: (
      <div>
        <SearchInput
          placeholder="Search by account number"
          onSearch={onSearch}
        />
      </div>
    ),
  },
];

interface PropTypes {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

export default function ThirdPartyProvidersView({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: PropTypes) {
  const [isProgress, setProgress] = useState(true);
  const [clients, setClients] = useState<ClientType[] | []>([]);
  const classes = useStyles();

  useEffect(() => {
    setProgress(true);
    api
      .getClients()
      .then(({ clients }: { clients: ClientType[] }) => {
        setClients(clients || []);
      })
      .catch((err) => console.log(err))
      .finally(() => setProgress(false));
  }, []);

  // const handleRevokeConsent = (id: string) => {
  //   setProgress(true);
  //   api
  //     .deleteConsent({ id })
  //     .then(api.getClients)
  //     .then((res) => setClients(res.clients || []))
  //     .catch((err) => console.log(err))
  //     .finally(() => setProgress(false));
  // };

  const handleRevokeClient = (id: string, provider_type: string) => {
    setProgress(true);
    api
      .deleteClient({ id, provider_type })
      .then(api.getClients)
      .then((res) => {
        console.log(res)
        setClients(res.clients || [])
      })
      .catch((err) => console.log(err))
      .finally(() => setProgress(false));
  };

  return (
    <div>
      <PageToolbar
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />
      <div style={{ position: "relative" }}>
        {isProgress && <Progress />}
        {!isProgress && (
          <>
            {clients.length === 0 && (
              <div style={{ textAlign: "center", marginTop: 128 }}>
                <Typography variant="h3" style={{ color: "#626576" }}>
                  No authorized 3rd party Applications
                </Typography>
                <img
                  src={noAccountEmptyState}
                  style={{ marginTop: 64 }}
                  alt="empty state"
                />
              </div>
            )}
            {clients.length > 0 && (
              <>
                <Subheader
                  title="Third party providers (TPPs)"
                  icon={iconShield}
                >
                  <div className={classes.subtitle}>
                    Go Bank has granted API access to the following third-party
                    applications. Revoke access to those you no longer trust.
                  </div>
                </Subheader>
                <Container>
                  <ClientsList
                    clients={clients}
                    onRevokeClient={handleRevokeClient}
                  />
                </Container>
              </>
            )}
          </>
        )}
      </div>
    </div>
  );
}

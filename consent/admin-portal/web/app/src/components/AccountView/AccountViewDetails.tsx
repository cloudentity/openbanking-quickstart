import React, { useEffect, useState } from "react";
import { makeStyles, Theme } from "@material-ui/core";
import { useHistory, useLocation, useParams } from "react-router";
import ArrowBack from "@material-ui/icons/ArrowBack";
import IconButton from "@material-ui/core/IconButton";

import CustomTabs from "../CustomTabs";
import PageToolbar from "../PageToolbar";
import Progress from "../Progress";
import Subheader from "../Subheader";
import { ClientType, handleSearch, searchTabs } from "../utils";
import AccountClientCard from "./AccountClientCard";
import { api } from "../../api/api";
import ConsentTabs from "./ConsentTabs";

const useStyles = makeStyles((theme: Theme) => ({
  subtitle: {
    ...theme.custom.body1,
  },
  subHeaderToolbar: {
    backgroundColor: "#002D4C",
    height: 80,
    maxWidth: 850,
    margin: "0 auto",
    display: "flex",
  },
  header: {
    ...theme.custom.heading3,
    borderBottom: "solid 1px #ECECEC",
    paddingBottom: 24,
  },
  container: {
    maxWidth: 850,
    margin: "32px auto",
  },
  accountInfo: {
    backgroundColor: "white",
    boxShadow: "0px 0px 0px 1px #ECECEC",
    padding: 24,
    borderRadius: 4,
    marginBottom: 24,
  },
  back: {
    ...theme.custom.label,
    color: "white",
    marginLeft: 13,
  },
}));

interface PropTypes {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

interface LocationState {
  accounts: { [accountId: string]: string[] };
}

export default function AccountViewDetails({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: PropTypes) {
  const { id, clientId } = useParams<Record<string, string | undefined>>();
  const history = useHistory();
  const { state } = useLocation<LocationState>();
  const classes = useStyles();
  const [isProgress, setProgress] = useState(true);
  const [clients, setClients] = useState<ClientType[]>();
  const [client, setClient] = useState<ClientType>();
  const [consents, setConsents] = useState<any>();

  useEffect(() => {
    fetchClients();
  }, []);

  const fetchClients = () => {
    setProgress(true);
    api
      .getClients()
      .then(({ clients }: { clients: ClientType[] }) => {
        setClients(clients);
      })
      .catch((err) => console.log(err))
      .finally(() => setProgress(false));
  };

  useEffect(() => {
    if (clients && clientId && id) {
      const found = clients.find((v) => v.client_id === clientId);
      if (found) {
        setClient(found);
        const consents = found.consents?.filter((v) =>
          v.account_ids?.includes(id)
        );
        setConsents(consents);
      }
    }
  }, [clients, clientId, id]);

  const handleRevokeClient = (id: string) => {
    setProgress(true);
    api
      .deleteClient({ id })
      .then(fetchClients)
      .then((res) => setClients(res.clients || []))
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
            <Subheader
              title=""
              containerStyle={{
                backgroundColor: "transparent",
                padding: 0,
                textAlign: "left",
              }}
              contentStyle={{ maxWidth: "100%" }}
            >
              <div style={{ backgroundColor: "#002D4C" }}>
                <div className={classes.subHeaderToolbar}>
                  <IconButton
                    style={{ padding: 4 }}
                    onClick={() => {
                      if (id) {
                        handleSearch(id)(history, state?.accounts);
                      }
                    }}
                  >
                    <ArrowBack fontSize="small" style={{ color: "white" }} />
                    <div className={classes.back}>Back</div>
                  </IconButton>
                </div>
              </div>
              <div
                style={{
                  maxWidth: 590,
                  margin: "0 auto",
                  position: "relative",
                  top: -32,
                }}
              >
                <CustomTabs
                  tabs={searchTabs(
                    (searchText) =>
                      handleSearch(searchText)(history, state?.accounts),
                    id
                  )}
                />
              </div>
            </Subheader>
            <div className={classes.container}>
              <AccountClientCard
                key={client?.client_id}
                client={client}
                onRevokeClient={handleRevokeClient}
              />
              <ConsentTabs consents={consents} />
            </div>
          </>
        )}
      </div>
    </div>
  );
}

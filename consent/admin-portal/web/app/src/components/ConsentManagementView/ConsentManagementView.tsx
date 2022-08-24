import React, { useEffect, useState } from "react";
import { makeStyles } from "tss-react/mui";
import Typography from "@mui/material/Typography";
import { useNavigate } from "react-router-dom";
import mergeWith from "lodash/mergeWith";
import uniq from "lodash/uniq";

import PageToolbar from "../PageToolbar";
import Progress from "../Progress";
import { api } from "../../api/api";
import noAccountEmptyState from "../no-accounts-empty-state.svg";
import Subheader from "../Subheader";
import CustomTabs from "../CustomTabs";
import { ClientType, handleSearch, searchTabs } from "../utils";
import iconShield from "../../assets/icon-shield-2.svg";
import background from "../../assets/background.svg";

const mergeCustomizer = (objValue, srcValue) => {
  if (Array.isArray(objValue)) {
    return objValue.concat(srcValue);
  }
};

const useStyles = makeStyles()(theme => ({
  subtitle: {
    ...theme.custom.body1,
  },
  container: {
    position: "relative",
    minHeight: "calc(100vh - 64px)",
  },
}));

interface Props {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

export default function ConsentManagementView({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: Props) {
  const [isProgress, setProgress] = useState(true);
  const [clients, setClients] = useState<ClientType[] | []>([]);
  const [accounts, setAccounts] = useState<any>([]);
  const { classes } = useStyles();
  const navigate = useNavigate();

  useEffect(() => {
    setProgress(true);
    api
      .getClients()
      .then(({ clients }: { clients: ClientType[] }) => {
        setClients(clients || []);
        const accountIdToClients = clients?.reduce((clientsAcc, client) => {
          const { client_id, consents } = client;
          return mergeWith(
            clientsAcc,
            (consents ?? [])
              .flatMap(v => v.account_ids)
              .reduce(
                (accountsIdsAcc, accountId) => ({
                  ...accountsIdsAcc,
                  [accountId]: uniq([
                    ...(accountsIdsAcc[accountId] || []),
                    client_id,
                  ]),
                }),
                {}
              ),
            [],
            mergeCustomizer
          );
        }, {});
        setAccounts(accountIdToClients);
      })
      .catch(err => console.log(err))
      .finally(() => setProgress(false));
  }, []);

  return (
    <div>
      <PageToolbar
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />
      <div className={classes.container}>
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
                  title="Consent management"
                  icon={iconShield}
                  containerStyle={{
                    minHeight: "calc(100vh - 64px)",
                    backgroundImage: `url(${background})`,
                    backgroundPosition: "left",
                    backgroundRepeat: "no-repeat",
                    backgroundSize: "contain",
                  }}
                >
                  <div id="search-content" className={classes.subtitle}>
                    Search and manage consents on behalf of bank members
                  </div>
                  <div style={{ marginTop: 32 }}>
                    <CustomTabs
                      tabs={searchTabs(searchText =>
                        handleSearch(searchText)(navigate, accounts)
                      )}
                    />
                  </div>
                </Subheader>
              </>
            )}
          </>
        )}
      </div>
    </div>
  );
}

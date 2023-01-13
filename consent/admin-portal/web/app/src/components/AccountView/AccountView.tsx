import React, { useEffect, useState } from "react";
import { makeStyles } from "tss-react/mui";
import { useNavigate, useLocation, useParams } from "react-router-dom";
import ArrowBack from "@mui/icons-material/ArrowBack";
import IconButton from "@mui/material/IconButton";

import CustomTabs from "../CustomTabs";
import PageToolbar from "../PageToolbar";
import Progress from "../Progress";
import Subheader from "../Subheader";
import { api } from "../../api/api";
import { ClientType, handleSearch, searchTabs } from "../utils";
import AccountClientsList from "./AccountClientsList";

const useStyles = makeStyles()(theme => ({
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
  info: {
    "& span": {
      fontWeight: 500,
    },
  },
  subheader: {
    ...theme.custom.heading6,
  },
  empty: {
    textAlign: "center",
    color: "gray",
  },
  back: {
    ...theme.custom.label,
    color: "white",
    marginLeft: 13,
  },
}));

interface Props {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

interface LocationState {
  clientIds: string[];
  accounts: { [accountId: string]: string[] };
}

export default function AccountView({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: Props) {
  const { id } = useParams<Record<string, string | undefined>>();
  const navigate = useNavigate();
  const location = useLocation();
  const state = location.state as LocationState | undefined;
  const [isProgress, setProgress] = useState(true);
  const { classes } = useStyles();
  const [clients, setClients] = useState<ClientType[] | []>([]);

  useEffect(() => {
    setProgress(true);
    api
      .getClients()
      .then(({ clients }: { clients: ClientType[] }) => {
        const found = clients.filter(v =>
          state?.clientIds?.includes(v.client_id)
        );
        if (found.length) {
          setClients(found);
        }
      })
      .catch(err => console.log(err))
      .finally(() => setProgress(false));
  }, [state?.clientIds]);

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
                    onClick={() => navigate("/")}
                    onMouseDown={() => navigate("/")}
                    size="large"
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
                  tabs={searchTabs(searchText => {
                    handleSearch(searchText)(navigate, state?.accounts);
                  }, id)}
                />
              </div>
            </Subheader>
            <div className={classes.container}>
              {state?.clientIds ? (
                <>
                  <div className={classes.info} id="search-results">
                    <span>{clients.length} application(s)</span> found for
                    account <span>#{id}</span>
                  </div>
                  <AccountClientsList
                    clients={clients}
                    accountId={id}
                    accounts={state?.accounts}
                  />
                </>
              ) : (
                <div className={classes.empty}>No account found</div>
              )}
            </div>
          </>
        )}
      </div>
    </div>
  );
}

import React, { useEffect, useState } from "react";
import PageToolbar from "./PageToolbar";
import Progress from "./Progress";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import Container from "@material-ui/core/Container";
import Chip from "@material-ui/core/Chip";
import { makeStyles } from "@material-ui/core/styles";

import { api } from "../api/api";
import noAccountEmptyState from "../assets/no-accounts-empty-state.svg";
import Subheader from "./Subheader";
import ApplicationSimpleCard from "./ApplicationSimpleCard";

const useStyles = makeStyles(theme => ({
  filterTitle: {
    //  ...theme.custom.label,
    marginBottom: 12,
  },
  filterChips: {
    marginBottom: 24,
    "& > div": {
      marginRight: 8,
      ...theme.custom.label,
    },
  },
}));

const activeChipStyle = {
  backgroundColor: "#002D4C",
  color: "white",
};

export default function Dashboard({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const [isProgress, setProgress] = useState(true);
  const [clientConsents, setClientConsents] = useState<any>([]);
  const classes = useStyles();
  const [filter, setFilter] = useState<"all" | "account" | "payment">("all");

  useEffect(() => {
    setProgress(true);
    api
      .getConsents()
      .then(res =>
        setClientConsents(res.client_consents ? res.client_consents : [])
      )
      .catch(err => console.log(err))
      .finally(() => setProgress(false));
  }, []);

  const filteredClientConsents =
    (filter === "account" &&
      clientConsents.filter(v =>
        v.consents.every(
          c =>
            c.type === "account_access" ||
            c.type === "cdr_arrangement" ||
            c.type === "consents"
        )
      )) ||
    (filter === "payment" &&
      clientConsents.filter(v =>
        v.consents.every(c => c.type === "domestic_payment")
      )) ||
    clientConsents;

  return (
    <div style={{ background: "#F7FAFF", minHeight: "100vh" }}>
      <PageToolbar
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />
      <div style={{ marginTop: 64, position: "relative" }}>
        {isProgress && <Progress />}

        {!isProgress && (
          <>
            {clientConsents.length > 0 && (
              <Subheader title="Third-party connected apps">
                You have provided these third party applications with access to
                your accounts.
                <br />
                You can revoke access for those that you no longer trust.
              </Subheader>
            )}
            <Container style={{ marginTop: 64 }}>
              <Grid container justify="center">
                <Grid item xs={8}>
                  {clientConsents.length === 0 ? (
                    <div style={{ textAlign: "center", marginTop: 64 }}>
                      <Typography
                        id="no-account-title"
                        variant="h3"
                        style={{ color: "#626576" }}
                      >
                        No connected accounts
                      </Typography>
                      <Typography
                        id="no-account-subtitle"
                        style={{ marginTop: 12, color: "#A0A3B5" }}
                      >
                        You haven’t connected any accounts yet to manage access
                      </Typography>
                      <img
                        src={noAccountEmptyState}
                        style={{ marginTop: 64 }}
                        alt="empty state"
                      />
                    </div>
                  ) : (
                    <div>
                      <div className={classes.filterTitle}>
                        Filter by permissions:
                      </div>
                      <div className={classes.filterChips}>
                        <Chip
                          label="All types"
                          id="all-types"
                          onClick={() => setFilter("all")}
                          style={filter === "all" ? activeChipStyle : {}}
                        />
                        <Chip
                          label="Account only"
                          id="account-only"
                          onClick={() => setFilter("account")}
                          style={filter === "account" ? activeChipStyle : {}}
                        />
                        <Chip
                          label="Payment only"
                          id="payment-only"
                          onClick={() => setFilter("payment")}
                          style={filter === "payment" ? activeChipStyle : {}}
                        />
                      </div>
                      {filteredClientConsents.map(clientConsent => (
                        <ApplicationSimpleCard
                          key={clientConsent.id}
                          client={clientConsent}
                        />
                      ))}
                    </div>
                  )}
                </Grid>
              </Grid>
            </Container>
          </>
        )}
      </div>
    </div>
  );
}

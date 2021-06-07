import React from "react";
import { makeStyles, Theme } from "@material-ui/core";

import ClientCard from "./ClientCard";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    maxWidth: 850,
    margin: "32px auto",
  },
  header: {
    ...theme.custom.heading3,
  },
  subheader: {
    ...theme.custom.body2,
    paddingBottom: 16,
    borderBottom: "1px solid #ECECEC",
    marginBottom: 24,
  },
}));

export default function ClientsList({
  clients,
  onRevokeClient,
  onRevokeConsent,
}) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <div className={classes.header}>All Connected Applications</div>
      <div className={classes.subheader}>
        Manage and revoke access to connected third-party applications
      </div>
      {clients
        .sort((a, b) => ("" + a.client_name).localeCompare(b.client_name))
        .map((client) => (
          <ClientCard
            key={client.client_id}
            client={client}
            onRevokeClient={onRevokeClient}
            onRevokeConsent={onRevokeConsent}
          />
        ))}
    </div>
  );
}

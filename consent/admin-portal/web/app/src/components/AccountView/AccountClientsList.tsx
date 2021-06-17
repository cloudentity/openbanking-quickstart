import React from "react";
import { makeStyles } from "@material-ui/core";

import AccountClientCard from "./AccountClientCard";

const useStyles = makeStyles(() => ({
  container: {
    maxWidth: 850,
    margin: "32px auto",
  },
}));

export default function AccountClientsList({ clients, accountId, accounts }) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      {clients
        .sort((a, b) =>
          String(a?.client_name ?? "").localeCompare(b?.client_name ?? "")
        )
        .map((client) => (
          <AccountClientCard
            key={client?.client_id}
            client={client}
            accountId={accountId}
            accounts={accounts}
          />
        ))}
    </div>
  );
}

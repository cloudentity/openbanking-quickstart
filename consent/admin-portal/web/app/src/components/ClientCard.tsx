import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";

import { ClientType, ConsentStatus, getDate } from "./utils";
import RevokeDrawer from "./ThirdPartyProvidersView/RevokeDrawer";

const useStyles = makeStyles()(theme => ({
  card: {
    background: "#FFFFFF",
    padding: "12px 24px",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    marginBottom: 16,
  },
  cardDisabled: {
    backgroundColor: "#ECECEC",
  },
  date: {
    ...theme.custom.caption,
  },
  cardName: {
    display: "flex",
    alignItems: "center",
  },
  avatar: {
    background: "#FCFCFF",
    border: "1px solid #626576",
    color: "#626576",
    width: 48,
    height: 48,
    marginRight: 12,
    borderRadius: 4,
  },
  name: {
    ...theme.custom.heading3,
  },
  buttonRoot: {
    ...theme.custom.body2,
    textTransform: "capitalize",
    color: "#002D4C",
    backgroundColor: "#FCFCFF",
    border: "1px solid #002D4C",
    padding: "8px 24px",
  },
  buttonRootDisabled: {
    backgroundColor: "#DC1B37",
    border: "1px solid #DC1B37",
    color: "white !important",
  },
}));

function getAuthorisedDate(client) {
  const accountAccessConsent = client?.consents?.find(v => {
    return (
      v.consent_type === "account_access" ||
      v.consent_type === "cdr_arrangement" ||
      v.consent_type === "consents"
    );
  });
  if (accountAccessConsent) {
    const date = accountAccessConsent?.created_at;
    return getDate(date);
  }
  return "N/A";
}

interface Props {
  client: ClientType;
  onRevokeClient: (id: string, provider_type: string) => void;
}

export default function ClientCard({ client, onRevokeClient }: Props) {
  const { cx, classes } = useStyles();
  const date = getAuthorisedDate(client);
  const [openDrawer, setOpenDrawer] = useState(false);
  const status = client?.mainStatus;
  return (
    <div
      className={cx(
        classes.card,
        status === ConsentStatus.Inactive && classes.cardDisabled
      )}
    >
      <div className={classes.cardName}>
        <Avatar variant="square" className={classes.avatar}>
          {client?.client_name[0]?.toUpperCase()}
        </Avatar>
        <div>
          <div className={classes.name}>{client?.client_name}</div>
          {date !== "N/A" && (
            <div className={classes.date}>Connected {date}</div>
          )}
        </div>
      </div>
      {status !== ConsentStatus.Inactive ? (
        <Button
          variant="outlined"
          classes={{ root: classes.buttonRoot }}
          onClick={() => setOpenDrawer(true)}
        >
          Revoke
        </Button>
      ) : (
        <Button
          classes={{
            root: cx(classes.buttonRoot, classes.buttonRootDisabled),
          }}
          disabled
        >
          Revoked
        </Button>
      )}

      {openDrawer && (
        <RevokeDrawer
          handleClose={() => setOpenDrawer(false)}
          onConfirm={() =>
            client?.client_id &&
            onRevokeClient &&
            onRevokeClient(client?.client_id, client?.provider_type)
          }
          client={client}
        />
      )}
    </div>
  );
}
